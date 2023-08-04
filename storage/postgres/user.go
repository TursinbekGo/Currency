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

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(ctx context.Context, req *models.CreateUser) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO "users"(id, user_name, password, updated_at)
		VALUES ($1, $2, $3, NOW())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.UserName,
		req.Password,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *userRepo) GetByID(ctx context.Context, req *models.UserPrimaryKey) (*models.User, error) {

	
	var (
		query string

		id        sql.NullString
		user_name sql.NullString
		password  sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
		where = " id = $1"
		
		
	)
	if req.UserName!=""{
		req.Id=req.UserName
		where =" user_name = $1"
	}else{
		
	}

	query = `
		SELECT
			id,
			user_name,
			password,
			created_at,
			updated_at
		FROM "users"
		WHERE 
	`+ where

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&user_name,
		&password,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.User{
		Id:        id.String,
		UserName:  user_name.String,
		Password:  password.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}, nil
}

func (r *userRepo) GetList(ctx context.Context, req *models.UserGetListRequest) (*models.UserGetListResponse, error) {

	var (
		resp   = &models.UserGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			user_name,
			password,
			created_at,
			updated_at
		FROM "users"
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
			id        sql.NullString
			user_name sql.NullString
			password  sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&user_name,
			&password,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Users = append(resp.Users, &models.User{
			Id:        id.String,
			UserName:  user_name.String,
			Password:  password.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return resp, nil
}

func (r *userRepo) Update(ctx context.Context, req *models.UpdateUser) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
		"users"
		SET
		user_name = :user_name,
		password = :password,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":        req.Id,
		"user_name": req.UserName,
		"password":  req.Password,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *userRepo) Delete(ctx context.Context, req *models.UserPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM  users WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
