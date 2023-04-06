package service

import (
	"shopper"
	"shopper/pkg/repo"
)

type CompanyService struct {
	repo repo.Company
}

func (s *CompanyService) ModerateCompany(id int) error {
	return s.repo.ModerateCompany(id)
}

func NewCompanyService(repo repo.Company) *CompanyService {
	return &CompanyService{repo: repo}
}

func (s *CompanyService) CreateCompany(company shopper.Company, userId int) (int, error) {
	return s.repo.CreateCompany(company, userId)
}

func (s *CompanyService) GetCompanyById(id int) (shopper.Company, error) {
	return s.repo.GetCompany(id)
}

func (s *CompanyService) UpdateCompany(userId, companyId int, input shopper.UpdateCompanyInput) error {
	return s.repo.UpdateCompany(userId, companyId, input)
}
