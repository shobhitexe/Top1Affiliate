package store

import (
	"context"
	"log"
	"top1affiliate/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AdminStore interface {
	GetAdminFromUsername(ctx context.Context, username string) (*models.Admin, error)
	GetAffiliates(ctx context.Context, id string) ([]models.User, error)
	GetAffiliate(ctx context.Context, id string) (*models.User, error)
	AddAffiliate(ctx context.Context, payload models.AddAffiliate) error
	BlockAffiliate(ctx context.Context, id string) error
	EditAffiliate(ctx context.Context, payload models.EditAffiliate) error
	EditAffiliateWithPassword(ctx context.Context, payload models.EditAffiliate) error
	GetPayouts(ctx context.Context, typevar string) ([]models.Payouts, error)
	DeclinePayout(ctx context.Context, id string) error
	ApprovePayout(ctx context.Context, id string, amount float64) (string, error)
}

type adminStore struct {
	db *pgxpool.Pool
}

func NewAdminStore(db *pgxpool.Pool) AdminStore {
	return &adminStore{db: db}
}

func (s *adminStore) GetAdminFromUsername(ctx context.Context, username string) (*models.Admin, error) {

	var admin models.Admin

	query := `SELECT id, username, password FROM admins WHERE username = $1`

	if err := s.db.QueryRow(ctx, query, username).Scan(&admin.ID, &admin.Username, &admin.Password); err != nil {
		return nil, err
	}

	return &admin, nil
}

func (s *adminStore) GetAffiliates(ctx context.Context, id string) ([]models.User, error) {

	var affiliates []models.User

	query := `SELECT id, affiliate_id, name, commission, country, blocked FROM users
WHERE added_by = $1::integer
ORDER BY id DESC`

	rows, err := s.db.Query(ctx, query, id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {

		var affiliate models.User

		if err := rows.Scan(
			&affiliate.ID,
			&affiliate.AffiliateID,
			&affiliate.Name,
			&affiliate.Commission,
			&affiliate.Country,
			&affiliate.Blocked,
		); err != nil {
			return nil, err
		}

		affiliates = append(affiliates, affiliate)

	}

	return affiliates, nil
}

func (s *adminStore) GetAffiliate(ctx context.Context, id string) (*models.User, error) {

	var affiliate models.User

	query := `SELECT id, affiliate_id, name, balance, commission, country, blocked, client_link, sub_link FROM users WHERE id = $1`

	if err := s.db.QueryRow(ctx, query, id).Scan(&affiliate.ID,
		&affiliate.AffiliateID,
		&affiliate.Name,
		&affiliate.Balance,
		&affiliate.Commission,
		&affiliate.Country,
		&affiliate.Blocked,
		&affiliate.ClientLink,
		&affiliate.SubLink,
	); err != nil {

		return nil, err
	}

	return &affiliate, nil
}

func (s *adminStore) AddAffiliate(ctx context.Context, payload models.AddAffiliate) error {

	query := `INSERT INTO users (name, affiliate_id, password, commission, country, added_by, client_link, sub_link) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := s.db.Exec(ctx, query,
		payload.Name,
		payload.AffiliateID,
		payload.Password,
		payload.Commission,
		payload.Country,
		payload.AddedBy,
		payload.ClientLink,
		payload.SubLink,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *adminStore) BlockAffiliate(ctx context.Context, id string) error {

	query := `UPDATE users SET blocked = NOT blocked WHERE id = $1`

	if _, err := s.db.Exec(ctx, query, id); err != nil {
		return err
	}

	return nil

}

func (s *adminStore) EditAffiliate(ctx context.Context, payload models.EditAffiliate) error {

	query := `UPDATE users SET name = $1, country = $2, commission = $3, client_link = $4, sub_link = $5, balance = $6 WHERE id = $7`

	if _, err := s.db.Exec(ctx, query,
		payload.Name,
		payload.Country,
		payload.Commission,
		payload.ClientLink,
		payload.SubLink,
		payload.Balance,
		payload.ID,
	); err != nil {
		return err
	}

	return nil
}

func (s *adminStore) EditAffiliateWithPassword(ctx context.Context, payload models.EditAffiliate) error {

	query := `UPDATE users SET name = $1, country = $2, commission = $3, client_link = $4, sub_link = $5, balance = $6, password = $7 WHERE id = $8`

	if _, err := s.db.Exec(ctx, query,
		payload.Name,
		payload.Country,
		payload.Commission,
		payload.ClientLink,
		payload.SubLink,
		payload.Balance,
		payload.Password,
		payload.ID,
	); err != nil {
		return err
	}

	return nil
}

func (s *adminStore) GetPayouts(ctx context.Context, typevar string) ([]models.Payouts, error) {

	var payouts []models.Payouts

	query := `SELECT p.id, u.name, u.affiliate_id, p.amount, p.method, p.payout_type, p.status, 
       TO_CHAR(p.created_at, 'DD/MM/YYYY') AS created_at_str,
	   COALESCE(wd.iban_number , 'N/A') AS iban_number,
	   COALESCE(wd.swift_code , 'N/A') AS swift_code,
	   COALESCE(wd.bank_name , 'N/A') AS bank_name,
	   COALESCE(wd.chain_name , 'N/A') AS chain_name,
	   COALESCE(wd.wallet_address , 'N/A') AS wallet_address
FROM payouts p
LEFT JOIN users u ON u.id = p.user_id
LEFT JOIN wallet_details wd ON wd.user_id = p.user_id
WHERE LOWER(p.status) = LOWER($1)
ORDER BY p.created_at DESC
`

	rows, err := s.db.Query(ctx, query, typevar)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var payout models.Payouts

		if err := rows.Scan(
			&payout.ID,
			&payout.Name,
			&payout.AffiliateId,
			&payout.Amount,
			&payout.Method,
			&payout.Type,
			&payout.Status,
			&payout.CreatedAt,
			&payout.IBAN,
			&payout.SwiftCode,
			&payout.BankName,
			&payout.ChainName,
			&payout.WalletAddress,
		); err != nil {
			return nil, err
		}

		payouts = append(payouts, payout)

	}

	return payouts, nil
}

func (s *adminStore) DeclinePayout(ctx context.Context, id string) error {

	query := `UPDATE payouts SET status = 'REJECTED' WHERE id = $1`

	if _, err := s.db.Exec(ctx, query, id); err != nil {
		log.Println(err)
		return err
	}

	return nil

}

func (s *adminStore) ApprovePayout(ctx context.Context, id string, amount float64) (string, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		log.Println("Failed to start transaction:", err)
		return "", err
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	var userId string
	query1 := `UPDATE payouts SET status = 'PAID' WHERE id = $1 RETURNING user_id`
	if err = tx.QueryRow(ctx, query1, id).Scan(&userId); err != nil {
		log.Println("ApprovePayout query failed:", err)
		return "", err
	}

	query2 := `UPDATE users SET balance = balance - $2 WHERE id = $1`
	if _, err = tx.Exec(ctx, query2, userId, amount); err != nil {
		log.Println("DebitUser query failed:", err)
		return "", err
	}

	return userId, nil
}
