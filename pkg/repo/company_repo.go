package repo

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"shopper"
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
