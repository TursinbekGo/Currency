package postgres

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"database/sql"
	"fmt"

	uuid "github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type saleRepo struct {
	db *pgxpool.Pool
}

func NewSaleRepo(db *pgxpool.Pool) *saleRepo {
	return &saleRepo{
		db: db,
	}
}

func (r *saleRepo) Create(ctx context.Context, req *models.CreateSale) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO sale(id, user_id, total_price,total_count, updated_at)
		VALUES ($1, $2, $3, $4, NOW())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.UserID,
		req.TotalPrice,
		req.TotalCount,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *saleRepo) GetByID(ctx context.Context, req *models.SalePrimaryKey) (*models.Sale, error) {

	var (
		query string

		id          sql.NullString
		user_id     sql.NullString
		total_price sql.NullFloat64
		total_count sql.NullInt64
		createdAt   sql.NullString
		updatedAt   sql.NullString
	)

	query = `
		SELECT
			id,
			user_id,
			total_price,
			total_count,
			created_at,
			updated_at
		FROM sale
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&user_id,
		&total_price,
		&total_count,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.Sale{
		Id:         id.String,
		UserID:     user_id.String,
		TotalPrice: total_price.Float64,
		TotalCount: total_count.Int64,
		CreatedAt:  createdAt.String,
		UpdatedAt:  updatedAt.String,
	}, nil
}

func (r *saleRepo) GetList(ctx context.Context, req *models.SaleGetListRequest) (*models.SaleGetListResponse, error) {

	var (
		resp   = &models.SaleGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			user_id,
			total_price,
			total_count,
			created_at,
			updated_at
		FROM sale
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Search != "" {
		where += ` AND title ILIKE '%' || '` + req.Search + `' || '%'`
	}

	query += where + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id          sql.NullString
			user_id     sql.NullString
			total_price sql.NullFloat64
			total_count sql.NullInt64
			createdAt   sql.NullString
			updatedAt   sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&user_id,
			&total_price,
			&total_count,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Sales = append(resp.Sales, &models.Sale{
			Id:         id.String,
			UserID:     user_id.String,
			TotalPrice: total_price.Float64,
			TotalCount: total_count.Int64,
			CreatedAt:  createdAt.String,
			UpdatedAt:  updatedAt.String,
		})
	}

	return resp, nil
}

func (r *saleRepo) Update(ctx context.Context, req *models.UpdateSale) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
		sale
		SET
			total_price = :total_price,
			total_count = :total_count,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":          req.Id,
		"total_price": req.TotalPrice,
		"total_count": req.TotalCount,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *saleRepo) Delete(ctx context.Context, req *models.SalePrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM sale WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
