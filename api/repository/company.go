package repository

import (
	"context"

	"github.com/Wanted-Linx/linx-backend/api/domain"
	"github.com/Wanted-Linx/linx-backend/api/ent"
	"github.com/Wanted-Linx/linx-backend/api/ent/company"
	"github.com/pkg/errors"
)

type companyRepository struct {
	db *ent.Client
}

func NewCompanyRepository(db *ent.Client) domain.CompanyRepository {
	return &companyRepository{db: db}
}

func (r *companyRepository) Save(reqCompany *ent.Company) (*ent.Company, error) {
	c, err := r.db.Company.Create().
		SetID(reqCompany.ID).
		SetName(reqCompany.Name).
		SetBusinessNumber(reqCompany.BusinessNumber).
		SetUser(reqCompany.Edges.User).
		SetNillableProfileImage(reqCompany.ProfileImage).
		SetNillableDescription(reqCompany.Description).
		SetNillableAddress(reqCompany.Address).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	user, err := c.QueryUser().First(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	c.Edges.User = user
	return c, nil
}

func (r *companyRepository) GetByID(companyID int, reqStudent *ent.Company) (*ent.Company, error) {
	c, err := r.db.Company.Query().
		Where(company.ID(companyID)).
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	return c, nil
}