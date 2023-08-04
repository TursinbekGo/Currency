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

type SaleProductRepo struct {
	db *pgxpool.Pool
}

func NewSaleProductRepo(db *pgxpool.Pool) *SaleProductRepo {
	return &SaleProductRepo{
		db: db,
	}
}

func (r *SaleProductRepo) Create(ctx context.Context, req *models.CreateSaleProduct) (string, error) {
	trx, err := r.db.Begin(ctx)
	if err != nil {
		return "", nil
	}

	defer func() {
		if err != nil {
			trx.Rollback(ctx)
		} else {
			trx.Commit(ctx)
		}
	}()

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO sale_product(id, sale_id, product_id,discount,discount_type,product_name,product_price,price_with_discount,discount_price,count, updated_at)
		VALUES ($1, $2, $3, $4,$5,$6,$7,$8,$9,$10 ,NOW())
	`
	_, err = trx.Exec(ctx, query,
		id,
		req.SaleID,
		req.ProductID,
		req.Discount,
		req.DiscountType,
		req.ProductName,
		req.ProductPrice,
		req.PriceWithDiscount,
		req.DiscountPrice,
		req.Count,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *SaleProductRepo) GetByID(ctx context.Context, req *models.SaleProductPrimaryKey) (*models.SaleProduct, error) {

	var (
		query string

		id                  sql.NullString
		sale_id             sql.NullString
		product_id          sql.NullString
		discount            sql.NullInt64
		discount_type       sql.NullString
		product_name        sql.NullString
		product_price       sql.NullFloat64
		price_with_discount sql.NullFloat64
		discount_price      sql.NullFloat64
		count               sql.NullInt64
		createdAt           sql.NullString
		updatedAt           sql.NullString
	)

	query = `
		SELECT
			id,
			sale_id,
			product_id,
			discount,
			discount_type,
			product_name,
			product_price,
			price_with_discount,
			discount_price,
			count,
			created_at,
			updated_at
		FROM sale_product
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&sale_id,
		&product_id,
		&discount,
		&discount_type,
		&product_name,
		&product_price,
		&price_with_discount,
		&discount_price,
		&count,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.SaleProduct{
		Id:                id.String,
		SaleID:            sale_id.String,
		ProductID:         product_id.String,
		Discount:          discount.Int64,
		DiscountType:      discount_type.String,
		ProductName:       product_name.String,
		ProductPrice:      product_price.Float64,
		PriceWithDiscount: price_with_discount.Float64,
		DiscountPrice:     discount_price.Float64,
		Count:             count.Int64,
		CreatedAt:         createdAt.String,
		UpdatedAt:         updatedAt.String,
	}, nil
}

func (r *SaleProductRepo) GetList(ctx context.Context, req *models.SaleProductGetListRequest) (*models.SaleProductGetListResponse, error) {

	var (
		resp   = &models.SaleProductGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			sale_id,
			product_id,
			discount,
			discount_type,
			product_name,
			product_price,
			price_with_discount,
			discount_price,
			count,
			created_at,
			updated_at
		FROM sale_product
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
			id                  sql.NullString
			sale_id             sql.NullString
			product_id          sql.NullString
			discount            sql.NullInt64
			discount_type       sql.NullString
			product_name        sql.NullString
			product_price       sql.NullFloat64
			price_with_discount sql.NullFloat64
			discount_price      sql.NullFloat64
			count               sql.NullInt64
			createdAt           sql.NullString
			updatedAt           sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&sale_id,
			&product_id,
			&discount,
			&discount_type,
			&product_name,
			&product_price,
			&price_with_discount,
			&discount_price,
			&count,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.SaleProducts = append(resp.SaleProducts, &models.SaleProduct{
			Id:                id.String,
			SaleID:            sale_id.String,
			ProductID:         product_id.String,
			Discount:          discount.Int64,
			DiscountType:      discount_type.String,
			ProductName:       product_name.String,
			ProductPrice:      product_price.Float64,
			PriceWithDiscount: price_with_discount.Float64,
			DiscountPrice:     discount_price.Float64,
			Count:             count.Int64,
			CreatedAt:         createdAt.String,
			UpdatedAt:         updatedAt.String,
		})
	}

	return resp, nil
}

func (r *SaleProductRepo) Update(ctx context.Context, req *models.UpdateSaleProduct) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
		sale_product
		SET
		sale_id = :sale_id,
		product_id = :product_id,
		discount = :discount,
		discount_type = :discount_type,
		product_name = :product_name,
		product_price = :product_price,
		price_with_discount = :price_with_discount,
		discount_price = :discount_price,
		count = :count,
		updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"sale_id":             req.SaleID,
		"product_id":          req.ProductID,
		"discount":            req.Discount,
		"discount_type":       req.DiscountType,
		"product_name":        req.ProductName,
		"product_price":       req.ProductPrice,
		"price_with_discount": req.PriceWithDiscount,
		"discount_price":      req.DiscountPrice,
		"count":               req.Count,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *SaleProductRepo) Delete(ctx context.Context, req *models.SaleProductPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM sale_product WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
