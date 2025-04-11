package infrastructure

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/go-utilities/uuid"
	"github.com/namhq1989/worddrop-server/internal/database"
	"github.com/namhq1989/worddrop-server/internal/database/gen/word_drop/public/model"
	"github.com/namhq1989/worddrop-server/internal/database/gen/word_drop/public/table"
	apperrors "github.com/namhq1989/worddrop-server/internal/error"
	"github.com/namhq1989/worddrop-server/pkg/content/domain"
	"github.com/namhq1989/worddrop-server/pkg/content/infrastructure/mapping"
)

type WordRepository struct {
	db *database.Database
}

func NewWordRepository(db *database.Database) WordRepository {
	r := WordRepository{
		db: db,
	}

	return r
}

func (r WordRepository) getDB() *sql.DB {
	return r.db.GetDB()
}

func (WordRepository) getTable() *table.WordsTable {
	return table.Words
}

func (r WordRepository) FindWithFilter(ctx *appcontext.AppContext, filter domain.WordFilter) ([]domain.Word, error) {
	var (
		w = r.getTable()
	)

	whereStmt := w.LastFetchedAt.LT(postgres.TimestampzT(filter.Timestamp))
	stmt := postgres.SELECT(
		w.ID, w.Word, w.Level, w.Definitions, w.PartsOfSpeech, w.Ipa,
		w.NounForm, w.VerbForm, w.LastFetchedAt,
	).
		FROM(w).
		WHERE(whereStmt).
		LIMIT(filter.Limit).
		ORDER_BY(w.LastFetchedAt.DESC())

	var (
		docs   = make([]model.Words, 0)
		result = make([]domain.Word, 0)
	)
	if err := stmt.QueryContext(ctx.Context(), r.getDB(), &docs); err != nil {
		if r.db.IsNoRowsError(err) {
			return nil, nil
		}
		return nil, err
	}

	var (
		mapper = mapping.WordMapper{}
	)
	for _, doc := range docs {
		word, err := mapper.FromModelToDomain(doc)
		if err != nil {
			return nil, err
		}
		result = append(result, *word)
	}
	return result, nil
}

func (r WordRepository) FindNewWord(ctx *appcontext.AppContext, categories []string, levels []string, ts time.Time) (*domain.Word, error) {
	var (
		w  = r.getTable().AS("words")
		wn = table.WordNews.AS("wn")
	)

	whereStmt := wn.PublishedAt.GT_EQ(postgres.TimestampzT(ts))
	if len(levels) > 0 {
		inElements := make([]string, len(levels))
		for i, level := range levels {
			escapedLevel := strings.ReplaceAll(level, "'", "''")
			inElements[i] = fmt.Sprintf("'%s'", escapedLevel)
		}

		inClause := fmt.Sprintf("(%s)", strings.Join(inElements, ", "))
		levelsCondition := postgres.BoolExp(postgres.Raw(fmt.Sprintf("words.level IN %s", inClause)))
		whereStmt = whereStmt.AND(levelsCondition)
	}
	if len(categories) > 0 {
		arrayElements := make([]string, len(categories))
		for i, category := range categories {
			escapedCategory := strings.ReplaceAll(category, "'", "''")
			arrayElements[i] = fmt.Sprintf("'%s'", escapedCategory)
		}

		arrayString := fmt.Sprintf("ARRAY[%s]", strings.Join(arrayElements, ", "))
		categoriesCondition := postgres.BoolExp(postgres.Raw(fmt.Sprintf("wn.categories && %s", arrayString)))
		whereStmt = whereStmt.AND(categoriesCondition)
	}

	stmt := postgres.SELECT(
		w.ID, w.Word, w.Level, w.Definitions, w.PartsOfSpeech, w.Ipa,
		w.NounForm, w.VerbForm,
	).
		FROM(w.LEFT_JOIN(wn, wn.WordID.EQ(w.ID))).
		WHERE(whereStmt).
		ORDER_BY(postgres.Raw("RANDOM()")).
		LIMIT(1)

	var doc model.Words
	if err := stmt.QueryContext(ctx.Context(), r.getDB(), &doc); err != nil {
		if r.db.IsNoRowsError(err) {
			return nil, nil
		}
		return nil, err
	}

	var (
		mapper    = mapping.WordMapper{}
		result, _ = mapper.FromModelToDomain(doc)
	)
	return result, nil
}

func (r WordRepository) FindByWord(ctx *appcontext.AppContext, word string) (*domain.Word, error) {
	var w = r.getTable()
	stmt := postgres.SELECT(
		w.ID, w.Word, w.Level, w.Definitions, w.PartsOfSpeech, w.Ipa,
		w.NounForm, w.VerbForm, w.LastFetchedAt,
	).
		FROM(w).
		WHERE(w.Word.EQ(postgres.String(word)))

	var doc model.Words
	if err := stmt.QueryContext(ctx.Context(), r.getDB(), &doc); err != nil {
		if r.db.IsNoRowsError(err) {
			return nil, nil
		}
		return nil, err
	}

	var (
		mapper    = mapping.WordMapper{}
		result, _ = mapper.FromModelToDomain(doc)
	)
	return result, nil
}

func (r WordRepository) FindByID(ctx *appcontext.AppContext, wordID string) (*domain.Word, error) {
	if !uuid.IsValidID(wordID) {
		return nil, apperrors.Common.InvalidID
	}

	var w = r.getTable()
	stmt := postgres.SELECT(
		w.ID, w.Word, w.Level, w.Definitions, w.PartsOfSpeech, w.Ipa,
		w.NounForm, w.VerbForm,
	).
		FROM(w).
		WHERE(w.ID.EQ(postgres.String(wordID)))

	var doc model.Words
	if err := stmt.QueryContext(ctx.Context(), r.getDB(), &doc); err != nil {
		if r.db.IsNoRowsError(err) {
			return nil, nil
		}
		return nil, err
	}

	var (
		mapper    = mapping.WordMapper{}
		result, _ = mapper.FromModelToDomain(doc)
	)
	return result, nil
}

func (r WordRepository) FindSimilar(ctx *appcontext.AppContext, pos, level string, ts time.Time, limit int64) ([]domain.Word, error) {
	var (
		w = r.getTable()
	)

	whereStmt := w.Level.EQ(postgres.String(level)).
		AND(
			postgres.BoolExp(postgres.Raw("$pos = ANY(words.parts_of_speech)", postgres.RawArgs{
				"$pos": pos,
			}))).
		AND(w.LastFetchedAt.GT_EQ(postgres.TimestampzT(ts)))

	stmt := postgres.SELECT(
		w.ID, w.Word, w.Level, w.Definitions, w.PartsOfSpeech, w.Ipa,
		w.NounForm, w.VerbForm,
	).
		FROM(w).
		WHERE(whereStmt).
		LIMIT(limit).
		ORDER_BY(w.LastFetchedAt.DESC())

	var (
		docs   = make([]model.Words, 0)
		result = make([]domain.Word, 0)
	)
	if err := stmt.QueryContext(ctx.Context(), r.getDB(), &docs); err != nil {
		if r.db.IsNoRowsError(err) {
			return nil, nil
		}
		return nil, err
	}

	var (
		mapper = mapping.WordMapper{}
	)
	for _, doc := range docs {
		word, err := mapper.FromModelToDomain(doc)
		if err != nil {
			return nil, err
		}
		result = append(result, *word)
	}
	return result, nil
}

func (r WordRepository) Create(ctx *appcontext.AppContext, word domain.Word) error {
	mapper := mapping.WordMapper{}
	doc, err := mapper.FromDomainToModel(word)
	if err != nil {
		return err
	}

	stmt := r.getTable().INSERT(
		r.getTable().AllColumns,
	).
		MODEL(doc)

	_, err = stmt.ExecContext(ctx.Context(), r.getDB())
	if err != nil {
		if isDuplicated, duplicateErr := r.db.IsDuplicatedError(err); isDuplicated {
			err = duplicateErr
		}
	}
	return err
}

func (r WordRepository) Update(ctx *appcontext.AppContext, word domain.Word) error {
	mapper := mapping.WordMapper{}
	doc, err := mapper.FromDomainToModel(word)
	if err != nil {
		return err
	}

	stmt := r.getTable().UPDATE(
		r.getTable().AllColumns,
	).
		MODEL(doc).
		WHERE(
			r.getTable().ID.EQ(postgres.String(doc.ID)),
		)

	_, err = stmt.ExecContext(ctx.Context(), r.getDB())
	if err != nil {
		if isDuplicated, duplicateErr := r.db.IsDuplicatedError(err); isDuplicated {
			err = duplicateErr
		}
	}
	return err
}

func (r WordRepository) Delete(ctx *appcontext.AppContext, word domain.Word) error {
	stmt := r.getTable().
		DELETE().
		WHERE(r.getTable().ID.EQ(postgres.String(word.ID)))

	_, err := stmt.ExecContext(ctx.Context(), r.getDB())
	return err
}
