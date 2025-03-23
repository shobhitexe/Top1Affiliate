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
