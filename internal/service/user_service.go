package service

import (
	"context"
	"errors"
	"top1affiliate/internal/models"
	"top1affiliate/internal/store"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	UserLogin(ctx context.Context, payload models.LoginRequest) (*models.User, error)
	RequestPayout(ctx context.Context, payload models.RequestPayout) error
	GetPayouts(ctx context.Context, id, from, to string) ([]models.Payouts, error)
	GetWalletDetails(ctx context.Context, id string) (*models.WalletDetails, error)
	UpdateWalletDetails(ctx context.Context, payload models.WalletDetails) error
}

type userService struct {
	store store.UserStore
}

func NewUserService(store store.UserStore) UserService {
	return &userService{store: store}
}

func (s *userService) UserLogin(ctx context.Context, payload models.LoginRequest) (*models.User, error) {

	user, err := s.store.GetUserFromID(ctx, payload.Login)

	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (s *userService) RequestPayout(ctx context.Context, payload models.RequestPayout) error {

	if err := s.store.RequestPayout(ctx, payload); err != nil {
		return err
	}

	return nil
}

func (s *userService) GetPayouts(ctx context.Context, id, from, to string) ([]models.Payouts, error) {

	p, err := s.store.GetPayouts(ctx, id, from, to)

	if err != nil {
		return nil, err
	}

	return p, nil

}

func (s *userService) GetWalletDetails(ctx context.Context, id string) (*models.WalletDetails, error) {

	w, err := s.store.GetWalletDetails(ctx, id)

	if err != nil {
		return nil, err
	}

	return w, err
}

func (s *userService) UpdateWalletDetails(ctx context.Context, payload models.WalletDetails) error {

	if err := s.store.UpdateWalletDetails(ctx, payload); err != nil {
		return err
	}

	return nil

}
