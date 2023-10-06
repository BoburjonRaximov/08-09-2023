package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"playground/cpp-bootcamp/models"
	"playground/cpp-bootcamp/pkg/helper"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUser(db *pgxpool.Pool) *userRepo {
	return &userRepo{
		db: db,
	}
}
func (p *userRepo) Create(req models.CreateUser) (string, error) {
	fmt.Println("User create")
	id := uuid.NewString()

	query := `
	INSERT INTO 
		users(id,
			username,
			password,
			name,
			phone_number,
			age) 
	VALUES($1,$2,$3,$4,$5,$6)`
	_, err := p.db.Exec(context.Background(), query,
		id,
		req.Username,
		req.Password,
		req.Name,
		req.PhoneNumber,
		req.Age,
	)
	if err != nil {
		fmt.Println("error:", err.Error())
		return "", err
	}
	return id, nil
}

func (p *userRepo) Update(req models.User) (string, error) {
	query := `
	UPDATE 
		users
	SET 
		name=$2,
		phone_number=$3,
		age=$4,
		updated_at= NOW()
	WHERE
		 id=$1`
	resp, err := p.db.Exec(context.Background(), query,
		req.Id,
		req.Name,
		req.PhoneNumber,
		req.Age,
	)
	if err != nil {
		return "", err
	}
	if resp.RowsAffected() == 0 {
		return "", pgx.ErrNoRows
	}
	return "OK", nil
}
func (p *userRepo) Get(req models.RequestByID) (models.User, error) {
	var (
		name        sql.NullString
		phoneNumber sql.NullString
		age         sql.NullInt64
		createdAt   sql.NullString
		updatedAt   sql.NullString
	)
	s := `
	SELECT
		id,
		name,
		phone_number,
		age,
		created_at,
	    updated_at
	FROM 
		users
	WHERE 
		id=$1
	`
	user := models.User{}
	err := p.db.QueryRow(context.Background(), s, req.ID).Scan(
		&user.Id,
		&name,
		&phoneNumber,
		&age,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return user, err
	}
	return models.User{
		Id:          user.Id,
		Name:        name.String,
		PhoneNumber: phoneNumber.String,
		Age:         int(age.Int64),
		CreatedAt:   createdAt.String,
		UpdatedAt:   updatedAt.String,
	}, err
}

func (p *userRepo) GetAll(req models.GetAllUsersRequest) (*models.GetAllUsersResponse, error) {
	var (
		params      = make(map[string]interface{})
		filter      = " WHERE true "
		offsetQ     = " OFFSET 0 "
		limit       = " LIMIT 10 "
		offset      = (req.Page - 1) * req.Limit
		count       = 0
		name        sql.NullString
		phoneNumber sql.NullString
		age         sql.NullInt64
		createdAt   sql.NullString
		updatedAt   sql.NullString
	)
	s := `
	SELECT
		id,
		name,
		phone_number,
		age,
		created_at,
		updated_at
	FROM 
		users
	`
	countQ := `
	SELECT
		count(*)
	FROM 
		users
	`
	if req.Search != "" {
		filter += ` AND name ILIKE '%' || @search || '%' `
		params["search"] = req.Search
	}
	if req.Age > 0 {
		filter += ` AND age=@age `
		params["age"] = req.Age
	}
	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}
	if offset > 0 {
		offsetQ = fmt.Sprintf(" OFFSET %d", offset)
	}

	query := s + filter + limit + offsetQ
	countQ = countQ + filter
	c, pArr := helper.ReplaceQueryParams(countQ, params)
	err := p.db.QueryRow(context.Background(), c, pArr...).Scan(&count)
	if err != nil {
		return nil, err
	}
	q, pArr := helper.ReplaceQueryParams(query, params)
	rows, err := p.db.Query(context.Background(), q, pArr...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	resp := []models.User{}
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.Id,
			&name,
			&phoneNumber,
			&age,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, err
		}
		resp = append(resp, models.User{
			Id:          user.Id,
			Name:        name.String,
			PhoneNumber: phoneNumber.String,
			Age:         int(age.Int64),
			CreatedAt:   createdAt.String,
			UpdatedAt:   updatedAt.String,
		})
	}

	return &models.GetAllUsersResponse{Users: resp, Count: int32(count)}, nil
}

// getbyusername
func (p *userRepo) GetByUsername(req models.RequestByUsername) (models.User, error) {
	s := `
	SELECT
		username,
		password
	FROM
		users
	WHERE 
		username=$1
	`
	user := models.User{}
	err := p.db.QueryRow(context.Background(), s, req.Username).Scan(
		&user.Username,
		&user.Password,
	)
	if err != nil {
		return user, err
	}
	return user, err
}

func (p *userRepo) Delete(req models.RequestByID) (string, error) {
	return "", errors.New("not found")
}

