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

type WordExampleRepository struct {
	db *database.Database
}

func NewWordExampleRepository(db *database.Database) WordExampleRepository {
	r := WordExampleRepository{
		db: db,
	}

	return r
}

func (r WordExampleRepository) getDB() *sql.DB {
	return r.db.GetDB()
}

func (WordExampleRepository) getTable() *table.WordExamplesTable {
	return table.WordExamples
}

func (r WordExampleRepository) FindByWordID(ctx *appcontext.AppContext, wordID string) ([]domain.WordExample, error) {
	if !uuid.IsValidID(wordID) {
		return make([]domain.WordExample, 0), nil
	}

	var (
		we = r.getTable()
	)

	stmt := postgres.SELECT(
		we.ID, we.Level, we.Example, we.Audio, we.MainWord,
	).
		FROM(we).
		WHERE(we.WordID.EQ(postgres.String(wordID)))

	var (
		docs   = make([]model.WordExamples, 0)
		result = make([]domain.WordExample, 0)
	)
	if err := stmt.QueryContext(ctx.Context(), r.getDB(), &docs); err != nil {
		if r.db.IsNoRowsError(err) {
			return nil, nil
		}
		return nil, err
	}

	var (
		mapper = mapping.WordExampleMapper{}
	)
	for _, doc := range docs {
		example, err := mapper.FromModelToDomain(doc)
		if err != nil {
			return nil, err
		}
		result = append(result, *example)
	}
	return result, nil
}

func (r WordExampleRepository) Create(ctx *appcontext.AppContext, example domain.WordExample) error {
	mapper := mapping.WordExampleMapper{}
	doc, err := mapper.FromDomainToModel(example)
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
