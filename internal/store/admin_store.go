package store

import (
	"context"
	"top1affiliate/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AdminStore interface {
	GetAdminFromUsername(ctx context.Context, username string) (*models.Admin, error)
	GetAffiliates(ctx context.Context) ([]models.User, error)
	GetAffiliate(ctx context.Context, id string) (*models.User, error)
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

	query := `SELECT id, affiliate_id, name, commission, country FROM users ORDER BY id DESC`

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
		); err != nil {
			return nil, err
		}

		affiliates = append(affiliates, affiliate)

	}

	return affiliates, nil
}

func (s *adminStore) GetAffiliate(ctx context.Context, id string) (*models.User, error) {

	var affiliate models.User

	query := `SELECT id, affiliate_id, name, commission, country FROM users WHERE id = $1`

	if err := s.db.QueryRow(ctx, query, id).Scan(&affiliate.ID,
		&affiliate.AffiliateID,
		&affiliate.Name,
		&affiliate.Commission,
		&affiliate.Country); err != nil {

		return nil, err
	}

	return &affiliate, nil
}
