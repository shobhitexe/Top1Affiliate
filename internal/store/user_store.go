package store

import (
	"context"
	"log"
	"top1affiliate/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserStore interface {
	GetUserFromID(ctx context.Context, id string) (*models.User, error)
	RequestPayout(ctx context.Context, payload models.RequestPayout) error
	GetPayouts(ctx context.Context, id, from, to string) ([]models.Payouts, error)
}

type userStore struct {
	db *pgxpool.Pool
}

func NewUserStore(db *pgxpool.Pool) UserStore {
	return &userStore{db: db}
}

func (s *userStore) GetUserFromID(ctx context.Context, id string) (*models.User, error) {

	var user models.User

	query := `SELECT id, affiliate_id, password, name FROM users WHERE affiliate_id = $1`

	if err := s.db.QueryRow(ctx, query, id).Scan(&user.ID, &user.AffiliateID, &user.Password, &user.Name); err != nil {
		log.Println(err)
		return nil, err
	}

	return &user, nil
}

func (s *userStore) RequestPayout(ctx context.Context, payload models.RequestPayout) error {

	query := `INSERT INTO payouts (amount, payout_type, user_id) VALUES ($1 ,$2, $3)`

	if _, err := s.db.Exec(ctx, query, payload.Amount, payload.Type, payload.ID); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *userStore) GetPayouts(ctx context.Context, id, from, to string) ([]models.Payouts, error) {

	var payouts []models.Payouts

	query := `SELECT amount, payout_type, status, TO_CHAR(created_at, 'DD/MM/YYYY') AS created_at_str
	FROM payouts WHERE user_id = $1 AND created_at BETWEEN $2 AND $3 
	ORDER BY created_at DESC`

	rows, err := s.db.Query(ctx, query, id, from, to)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var payout models.Payouts

		if err := rows.Scan(&payout.Amount, &payout.Type, &payout.Status, &payout.CreatedAt); err != nil {
			return nil, err
		}

		payouts = append(payouts, payout)

	}

	return payouts, nil
}
