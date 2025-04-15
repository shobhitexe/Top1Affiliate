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
	GetWalletDetails(ctx context.Context, id string) (*models.WalletDetails, error)
	UpdateWalletDetails(ctx context.Context, payload models.WalletDetails) error
}

type userStore struct {
	db *pgxpool.Pool
}

func NewUserStore(db *pgxpool.Pool) UserStore {
	return &userStore{db: db}
}

func (s *userStore) GetUserFromID(ctx context.Context, id string) (*models.User, error) {

	var user models.User

	query := `SELECT id, affiliate_id, password, name, commission, link FROM users WHERE affiliate_id = $1`

	if err := s.db.QueryRow(ctx, query, id).Scan(&user.ID, &user.AffiliateID, &user.Password, &user.Name, &user.Commission, &user.Link); err != nil {
		log.Println(err)
		return nil, err
	}

	return &user, nil
}

func (s *userStore) RequestPayout(ctx context.Context, payload models.RequestPayout) error {

	query := `INSERT INTO payouts (amount, payout_type, user_id, method) VALUES ($1 ,$2, $3, $4)`

	if _, err := s.db.Exec(ctx, query, payload.Amount, payload.Type, payload.ID, payload.Method); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *userStore) GetPayouts(ctx context.Context, id, from, to string) ([]models.Payouts, error) {

	var payouts []models.Payouts

	query := `SELECT id, amount, payout_type, status, method, TO_CHAR(created_at, 'DD/MM/YYYY') AS created_at_str
	FROM payouts WHERE user_id = $1 AND created_at BETWEEN $2 AND $3 
	ORDER BY created_at DESC`

	rows, err := s.db.Query(ctx, query, id, from, to)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var payout models.Payouts

		if err := rows.Scan(&payout.ID, &payout.Amount, &payout.Type, &payout.Status, &payout.Method, &payout.CreatedAt); err != nil {
			return nil, err
		}

		payouts = append(payouts, payout)

	}

	return payouts, nil
}

func (s *userStore) GetWalletDetails(ctx context.Context, id string) (*models.WalletDetails, error) {

	var wallet models.WalletDetails

	query := `SELECT iban_number, swift_code, bank_name, chain_name, wallet_address FROM wallet_details WHERE user_id = $1`

	if err := s.db.QueryRow(ctx, query, id).Scan(
		&wallet.IBAN,
		&wallet.Swift,
		&wallet.BankName,
		&wallet.ChainName,
		&wallet.WalletAddress,
	); err != nil {
		return nil, err
	}

	return &wallet, nil
}

func (s *userStore) UpdateWalletDetails(ctx context.Context, payload models.WalletDetails) error {

	query := `INSERT INTO wallet_details (
    user_id, iban_number, swift_code, bank_name, chain_name, wallet_address
) VALUES (
    $1, $2, $3, $4, $5, $6
)
ON CONFLICT (user_id) 
DO UPDATE SET 
    iban_number = EXCLUDED.iban_number,
    swift_code = EXCLUDED.swift_code,
    bank_name = EXCLUDED.bank_name,
    chain_name = EXCLUDED.chain_name,
    wallet_address = EXCLUDED.wallet_address`

	if _, err := s.db.Exec(ctx, query,
		payload.ID,
		payload.IBAN,
		payload.Swift,
		payload.BankName,
		payload.ChainName,
		payload.WalletAddress,
	); err != nil {
		return err
	}

	return nil
}
