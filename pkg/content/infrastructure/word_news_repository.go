package infrastructure

import (
	"database/sql"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/go-utilities/uuid"
	"github.com/namhq1989/worddrop-server/internal/database"
	"github.com/namhq1989/worddrop-server/internal/database/gen/word_drop/public/model"
	"github.com/namhq1989/worddrop-server/internal/database/gen/word_drop/public/table"
	"github.com/namhq1989/worddrop-server/pkg/content/domain"
	"github.com/namhq1989/worddrop-server/pkg/content/infrastructure/mapping"
)

type WordNewsRepository struct {
	db *database.Database
}

func NewWordNewsRepository(db *database.Database) WordNewsRepository {
	r := WordNewsRepository{
		db: db,
	}

	return r
}

func (r WordNewsRepository) getDB() *sql.DB {
	return r.db.GetDB()
}

func (WordNewsRepository) getTable() *table.WordNewsTable {
	return table.WordNews
}

func (r WordNewsRepository) FindByWordID(ctx *appcontext.AppContext, wordID string) ([]domain.WordNews, error) {
	if !uuid.IsValidID(wordID) {
		return make([]domain.WordNews, 0), nil
	}

	var (
		wn = r.getTable()
	)

	stmt := postgres.SELECT(
		wn.ID, wn.Title, wn.Summary, wn.ImageURL, wn.SourceURL, wn.CreatedAt, wn.PublishedAt,
	).
		FROM(wn).
		WHERE(wn.WordID.EQ(postgres.String(wordID)))

	var (
		docs   = make([]model.WordNews, 0)
		result = make([]domain.WordNews, 0)
	)
	if err := stmt.QueryContext(ctx.Context(), r.getDB(), &docs); err != nil {
		if r.db.IsNoRowsError(err) {
			return nil, nil
		}
		return nil, err
	}

	var (
		mapper = mapping.WordNewsMapper{}
	)
	for _, doc := range docs {
		news, err := mapper.FromModelToDomain(doc)
		if err != nil {
			return nil, err
		}
		result = append(result, *news)
	}
	return result, nil
}

func (r WordNewsRepository) FindBySourceURLs(ctx *appcontext.AppContext, urls []string) ([]domain.WordNews, error) {
	var wn = r.getTable()
	stmt := postgres.SELECT(
		wn.ID, wn.SourceURL,
	).
		FROM(wn).
		WHERE(
			postgres.BoolExp(postgres.Raw("word_news.source_url = ANY($urls)", postgres.RawArgs{
				"$urls": urls,
			})))

	var (
		docs   = make([]model.WordNews, 0)
		result = make([]domain.WordNews, 0)
	)
	if err := stmt.QueryContext(ctx.Context(), r.getDB(), &docs); err != nil {
		if r.db.IsNoRowsError(err) {
			return nil, nil
		}
		return nil, err
	}

	var (
		mapper = mapping.WordNewsMapper{}
	)
	for _, doc := range docs {
		news, err := mapper.FromModelToDomain(doc)
		if err != nil {
			return nil, err
		}
		result = append(result, *news)
	}
	return result, nil
}

func (r WordNewsRepository) Create(ctx *appcontext.AppContext, news domain.WordNews) error {
	mapper := mapping.WordNewsMapper{}
	doc, err := mapper.FromDomainToModel(news)
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
