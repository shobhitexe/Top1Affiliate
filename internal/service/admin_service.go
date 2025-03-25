package service

import (
	"context"
	"errors"
	"top1affiliate/internal/models"
	"top1affiliate/internal/store"

	"golang.org/x/crypto/bcrypt"
)

type AdminService interface {
	AdminLogin(ctx context.Context, payload models.LoginRequest) (*models.Admin, error)
	GetAffiliates(ctx context.Context) ([]models.User, error)
	GetAffiliate(ctx context.Context, id string) (*models.User, error)
	AddAffiliate(ctx context.Context, payload models.AddAffiliate) error
	EditAffiliate(ctx context.Context, payload models.EditAffiliate) error
	BlockAffiliate(ctx context.Context, id string) error
}

type adminSevice struct {
	store store.AdminStore
}

func NewAdminService(store store.AdminStore) AdminService {
	return &adminSevice{store: store}
}

func (s *adminSevice) AdminLogin(ctx context.Context, payload models.LoginRequest) (*models.Admin, error) {

	a, err := s.store.GetAdminFromUsername(ctx, payload.Login)

	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(payload.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return a, nil
}

func (s *adminSevice) GetAffiliates(ctx context.Context) ([]models.User, error) {

	a, err := s.store.GetAffiliates(ctx)

	if err != nil {
		return nil, err
	}

	return a, nil
}

func (s *adminSevice) GetAffiliate(ctx context.Context, id string) (*models.User, error) {

	a, err := s.store.GetAffiliate(ctx, id)

	if err != nil {
		return nil, err
	}

	return a, nil
}

func (s *adminSevice) AddAffiliate(ctx context.Context, payload models.AddAffiliate) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	payload.Password = string(hashedPassword)

	if err := s.store.AddAffiliate(ctx, payload); err != nil {
		return err
	}

	return nil
}

func (s *adminSevice) BlockAffiliate(ctx context.Context, id string) error {

	if err := s.store.BlockAffiliate(ctx, id); err != nil {
		return err
	}

	return nil
}

func (s *adminSevice) EditAffiliate(ctx context.Context, payload models.EditAffiliate) error {

	if err := s.store.EditAffiliate(ctx, payload); err != nil {
		return err
	}

	return nil

}
