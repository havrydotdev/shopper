package repo

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"shopper"
	"strings"
)

type CompanyRepo struct {
	db *sqlx.DB
}

func NewCompanyRepo(db *sqlx.DB) *CompanyRepo {
	return &CompanyRepo{
		db: db,
	}
}

func (r *CompanyRepo) GetCompany(userId int) (shopper.Company, error) {
	var company shopper.Company
	query := fmt.Sprintf("SELECT c.name, c.logo, c.description, c.isverified FROM %s u INNER JOIN %s c on c.id = u.company_id WHERE u.id = $1", users, companies)
	err := r.db.Get(&company, query, userId)
	if err != nil {
		return shopper.Company{}, err
	}

	return company, nil
}

func (r *CompanyRepo) CreateCompany(company shopper.Company, userId int) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, description, logo) values ($1, $2, $3) RETURNING id", companies)
	row := r.db.QueryRow(query, company.Name, company.Description, company.Logo)
	if err := row.Scan(&id); err != nil {
		return -1, err
	}

	update := fmt.Sprintf("UPDATE %s u SET company_id = $1 WHERE u.id = $2", users)
	_, err := r.db.Exec(update, id, userId)

	return id, err
}

func (r *CompanyRepo) ModerateCompany(id int) error {
	update := fmt.Sprintf("UPDATE %s c SET isverified = true WHERE c.id = $1", companies)
	_, err := r.db.Exec(update, id)

	return err
}

func (r *CompanyRepo) UpdateCompany(userId, companyId int, input shopper.UpdateCompanyInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.Logo != nil {
		setValues = append(setValues, fmt.Sprintf("logo=$%d", argId))
		args = append(args, *input.Logo)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s c SET %s FROM %s u WHERE u.company_id = c.id AND u.id = $%d AND c.id = $%d", companies, setQuery, users, argId, argId+1)

	args = append(args, userId, companyId)

	res, err := r.db.Exec(query, args...)

	rows, err := res.RowsAffected()

	if rows == 0 {
		return errors.New("no rows affected")
	}

	return err
}
