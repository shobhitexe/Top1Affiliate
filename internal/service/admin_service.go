package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
	"top1affiliate/internal/models"
	"top1affiliate/internal/store"
	"top1affiliate/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

type AdminService interface {
	AdminLogin(ctx context.Context, payload models.LoginRequest) (*models.Admin, error)
	GetAffiliates(ctx context.Context, id string) ([]models.User, error)
	GetAffiliate(ctx context.Context, id string) (*models.User, error)
	AddAffiliate(ctx context.Context, payload models.AddAffiliate) error
	EditAffiliate(ctx context.Context, payload models.EditAffiliate) error
	BlockAffiliate(ctx context.Context, id string) error
	GetPayouts(ctx context.Context, typevar string) ([]models.Payouts, error)
	DeclinePayout(ctx context.Context, id string) error
	ApprovePayout(ctx context.Context, id string, amount float64) error
}

type adminSevice struct {
	store store.AdminStore
	utils utils.Utils
}

func NewAdminService(store store.AdminStore, utils utils.Utils) AdminService {
	return &adminSevice{store: store, utils: utils}
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

func (s *adminSevice) GetAffiliates(ctx context.Context, id string) ([]models.User, error) {

	a, err := s.store.GetAffiliates(ctx, id)

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

	if payload.Password == "" {
		if err := s.store.EditAffiliate(ctx, payload); err != nil {
			return err
		}
	} else {

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		payload.Password = string(hashedPassword)

		if err := s.store.EditAffiliateWithPassword(ctx, payload); err != nil {
			return err
		}

	}

	return nil

}

func (s *adminSevice) GetPayouts(ctx context.Context, typevar string) ([]models.Payouts, error) {

	p, err := s.store.GetPayouts(ctx, typevar)

	if err != nil {
		return nil, err
	}

	return p, nil

}

func (s *adminSevice) DeclinePayout(ctx context.Context, id string) error {

	if err := s.store.DeclinePayout(ctx, id); err != nil {
		return err
	}

	return nil
}

func (s *adminSevice) ApprovePayout(ctx context.Context, id string, amount float64) error {

	userId, err := s.store.ApprovePayout(ctx, id, amount)

	if err != nil {
		return err
	}

	go func() {

		gCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		user, err := s.store.GetAffiliate(gCtx, userId)

		if err != nil {
			log.Println(err)
			return
		}

		message := fmt.Sprintf(
			`Withdrawal successfull
	
	Crm ID:            %s
	Name:              %s
	Country:           %s
	Amount:            %.2f`,
			user.AffiliateID,
			user.Name,
			user.Country,
			amount,
		)

		if err := s.utils.SendNotificationToSlack(gCtx, models.WithdrawalSuccessfull, message); err != nil {
			log.Println(err)
			return
		}
	}()

	return nil
}
