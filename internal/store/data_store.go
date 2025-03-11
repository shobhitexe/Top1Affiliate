package store

import (
	"context"
	"log"
	"time"
	"top1affiliate/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DataStore interface {
	SaveLeadsData(ctx context.Context, lead models.Leads) error
	Getstatistics(ctx context.Context, id string) ([]models.Leads, error)
}

type dataStore struct {
	db *pgxpool.Pool
}

func NewDataStore(db *pgxpool.Pool) DataStore {
	return &dataStore{db: db}
}

func parseTimestamp(value string) *time.Time {
	if value == "" {
		return nil
	}
	t, err := time.Parse("2006-01-02T15:04:05", value)
	if err != nil {
		return nil
	}
	return &t
}

func (s *dataStore) SaveLeadsData(ctx context.Context, lead models.Leads) error {

	query := `
		INSERT INTO leads (
		id, first_name, last_name, last_login_date, lead_guid, country, city, sales_status,
		language, business_unit, domain_name, is_qualified, conversion_agent_id, retention_manager_id,
		vip_manager_id, closer_manager_id, affiliate_id, registration_date, account_creation_date,
		activation_date, fully_activation_date, deposited, original_lead_id, original_by_name_lead_id, email
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25
	)
	ON CONFLICT (id) DO UPDATE SET
	    id = EXCLUDED.id,
		first_name = EXCLUDED.first_name,
		last_name = EXCLUDED.last_name,
		last_login_date = EXCLUDED.last_login_date,
		lead_guid = EXCLUDED.lead_guid,
		country = EXCLUDED.country,
		city = EXCLUDED.city,
		sales_status = EXCLUDED.sales_status,
		language = EXCLUDED.language,
		business_unit = EXCLUDED.business_unit,
		domain_name = EXCLUDED.domain_name,
		is_qualified = EXCLUDED.is_qualified,
		conversion_agent_id = EXCLUDED.conversion_agent_id,
		retention_manager_id = EXCLUDED.retention_manager_id,
		vip_manager_id = EXCLUDED.vip_manager_id,
		closer_manager_id = EXCLUDED.closer_manager_id,
		affiliate_id = EXCLUDED.affiliate_id,
		registration_date = EXCLUDED.registration_date,
		account_creation_date = EXCLUDED.account_creation_date,
		activation_date = EXCLUDED.activation_date,
		fully_activation_date = EXCLUDED.fully_activation_date,
		deposited = EXCLUDED.deposited,
		original_lead_id = EXCLUDED.original_lead_id,
		original_by_name_lead_id = EXCLUDED.original_by_name_lead_id,
		email = EXCLUDED.email
	`

	if _, err := s.db.Exec(ctx, query,
		lead.ID,
		lead.FirstName,
		lead.LastName,
		parseTimestamp(lead.LastLoginDate),
		lead.LeadGuid,
		lead.Country,
		lead.City,
		lead.SalesStatus,
		lead.Language,
		lead.BusinessUnit,
		lead.DomainName,
		lead.IsQualified,
		lead.ConversionAgentID,
		lead.RetentionManagerID,
		lead.VIPManagerID,
		lead.CloserManagerID,
		lead.AffiliateID,
		parseTimestamp(lead.RegistrationDate),
		parseTimestamp(lead.AccountCreationDate),
		parseTimestamp(lead.ActivationDate),
		parseTimestamp(lead.FullyActivationDate),
		lead.Deposited,
		lead.OriginalLeadID,
		lead.OriginalByNameLeadID,
		lead.Email,
	); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *dataStore) Getstatistics(ctx context.Context, id string) ([]models.Leads, error) {

	var leads []models.Leads

	query := `SELECT affiliate_id, first_name, last_name, country, 
	TO_CHAR(registration_date, 'DD/MM/YYYY, HH12:MI:SS') AS registration_date_str 
	FROM leads WHERE affiliate_id = $1`

	rows, err := s.db.Query(ctx, query, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var lead models.Leads

		if err := rows.Scan(
			&lead.AffiliateID,
			&lead.FirstName,
			&lead.LastName,
			&lead.Country,
			&lead.RegistrationDate,
		); err != nil {
			log.Println(err)
			return nil, err
		}

		leads = append(leads, lead)

	}

	return leads, nil
}
