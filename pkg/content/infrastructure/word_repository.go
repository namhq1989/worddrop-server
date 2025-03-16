package infrastructure

import (
	"database/sql"
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
		w.ID, w.Word, w.Level, w.Definitions, w.PartsOfSpeech, w.Ipa, w.Audio,
		w.NounForm, w.VerbForm,
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

func (r WordRepository) FindByWord(ctx *appcontext.AppContext, word string) (*domain.Word, error) {
	var w = r.getTable()
	stmt := postgres.SELECT(
		w.ID, w.Word, w.Level, w.Definitions, w.PartsOfSpeech, w.Ipa, w.Audio,
		w.NounForm, w.VerbForm,
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
		w.ID, w.Word, w.Level, w.Definitions, w.PartsOfSpeech, w.Ipa, w.Audio,
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
		w.ID, w.Word, w.Level, w.Definitions, w.PartsOfSpeech, w.Ipa, w.Audio,
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
