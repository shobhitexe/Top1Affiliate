package store

import (
	"context"
	"log"
	"top1affiliate/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AdminStore interface {
	GetAdminFromUsername(ctx context.Context, username string) (*models.Admin, error)
	GetAffiliates(ctx context.Context) ([]models.User, error)
	GetAffiliate(ctx context.Context, id string) (*models.User, error)
	AddAffiliate(ctx context.Context, payload models.AddAffiliate) error
	BlockAffiliate(ctx context.Context, id string) error
	EditAffiliate(ctx context.Context, payload models.EditAffiliate) error
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

func (s *adminStore) GetAffiliates(ctx context.Context) ([]models.User, error) {

	var affiliates []models.User

	query := `SELECT id, affiliate_id, name, commission, country, blocked FROM users ORDER BY id DESC`

	rows, err := s.db.Query(ctx, query)

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

	query := `SELECT id, affiliate_id, name, commission, country, blocked FROM users WHERE id = $1`

	if err := s.db.QueryRow(ctx, query, id).Scan(&affiliate.ID,
		&affiliate.AffiliateID,
		&affiliate.Name,
		&affiliate.Commission,
		&affiliate.Country,
		&affiliate.Blocked,
	); err != nil {

		return nil, err
	}

	return &affiliate, nil
}

func (s *adminStore) AddAffiliate(ctx context.Context, payload models.AddAffiliate) error {

	query := `INSERT INTO users (name, affiliate_id, password, commission, country) VALUES ($1, $2, $3, $4, $5)`

	_, err := s.db.Exec(ctx, query, payload.Name, payload.AffiliateID, payload.Password, payload.Commission, payload.Country)

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

	log.Println(payload)

	query := `UPDATE users SET name = $1, country = $2, commission = $3 WHERE id = $4`

	if _, err := s.db.Exec(ctx, query, payload.Name, payload.Country, payload.Commission, payload.ID); err != nil {
		return err
	}

	return nil
}
