package store

import (
	"context"
	"fmt"
	"log"
	"time"
	"top1affiliate/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DataStore interface {
	SaveLeadsData(ctx context.Context, lead models.Leads) error
	Getstatistics(ctx context.Context, id string) ([]models.Statistics, error)
	GetEmailsOfLeads(ctx context.Context) ([]models.LeadsEmails, error)
	GetAllEmails(ctx context.Context) ([]models.LeadsEmails, error)
	SaveTransactions(ctx context.Context, transactions []models.Transaction, email, affiliateId string) error
	SaveTransactionsAndUpdateBalance(ctx context.Context, transactions []models.Transaction, email string, affiliateID string) error

	GetweeklyStats(ctx context.Context, id string) (*models.Stats, error)
	GetNetStats(ctx context.Context, id string) (*models.Stats, error)
	GetMonthlyStats(ctx context.Context, id string) (*models.Stats, error)

	GetTransactions(ctx context.Context, id, from, to string) ([]models.CommissionTxn, error)
	GetLatestFiveTransactions(ctx context.Context, id string) ([]models.CommissionTxn, error)
	GetLeaderboard(ctx context.Context) ([]models.Leaderboard, error)

	GetBalance(ctx context.Context, id string) (float64, error)
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

func (s *dataStore) Getstatistics(ctx context.Context, id string) ([]models.Statistics, error) {

	var leads []models.Statistics

	query := `SELECT 
    l.affiliate_id, 
    l.first_name, 
    l.last_name, 
    l.country, 
    TO_CHAR(l.registration_date, 'DD/MM/YYYY, HH12:MI:SS') AS registration_date_str,
    SUM(CASE WHEN t.transaction_type = 'Deposit' AND t.status = 'Complete' THEN t.amount ELSE 0 END) AS total_deposit,
	SUM(CASE WHEN t.transaction_type = 'Withdraw' THEN t.amount ELSE 0 END) AS total_withdrawal,
	COALESCE(SUM(t.commission_amount),0) AS total_commission
FROM leads l
LEFT JOIN transactions t ON t.email = l.email
WHERE l.affiliate_id = $1
GROUP BY l.affiliate_id, l.first_name, l.last_name, l.country, l.registration_date;`

	rows, err := s.db.Query(ctx, query, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var lead models.Statistics

		if err := rows.Scan(
			&lead.AffiliateID,
			&lead.FirstName,
			&lead.LastName,
			&lead.Country,
			&lead.RegistrationDate,
			&lead.Deposits,
			&lead.Withdrawals,
			&lead.Commissions,
		); err != nil {
			log.Println(err)
			return nil, err
		}

		leads = append(leads, lead)

	}

	return leads, nil
}

func (s *dataStore) GetEmailsOfLeads(ctx context.Context) ([]models.LeadsEmails, error) {

	var emails []models.LeadsEmails

	query := `SELECT u.affiliate_id, COALESCE(l.email, 'N/A') FROM users u
LEFT JOIN leads l ON l.affiliate_id = u.affiliate_id`

	rows, err := s.db.Query(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var email models.LeadsEmails

		if err := rows.Scan(&email.AffiliateID, &email.Email); err != nil {
			return nil, err
		}

		emails = append(emails, email)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return emails, nil

}

func (s *dataStore) GetAllEmails(ctx context.Context) ([]models.LeadsEmails, error) {

	var emails []models.LeadsEmails

	query := `SELECT affiliate_id, email FROM leads WHERE affiliate_id != ''`

	rows, err := s.db.Query(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var email models.LeadsEmails

		if err := rows.Scan(&email.AffiliateID, &email.Email); err != nil {
			return nil, err
		}

		emails = append(emails, email)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return emails, nil

}

func (s *dataStore) SaveTransactions(ctx context.Context, transactions []models.Transaction, email, affiliateId string) error {
	if len(transactions) == 0 {
		return nil
	}

	batch := &pgx.Batch{}

	query := `INSERT INTO transactions (
		transaction_id, amount, transaction_type, transaction_sub_type, status, 
		transaction_date, lead_id, lead_guid, affiliate_id, email
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	ON CONFLICT (transaction_id) DO NOTHING;`

	for _, txn := range transactions {
		batch.Queue(query,
			txn.TransactionID,
			txn.Amount,
			txn.TransactionType,
			txn.TransactionSubType,
			txn.Status,
			txn.TransactionDate,
			txn.LeadID,
			txn.LeadGUID,
			affiliateId,
			email,
		)
	}

	br := s.db.SendBatch(ctx, batch)
	defer br.Close()

	for range transactions {
		_, err := br.Exec()
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *dataStore) SaveTransactionsAndUpdateBalance(ctx context.Context, transactions []models.Transaction, email string, affiliateID string) error {
	if len(transactions) == 0 {
		return nil
	}

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	var commission float64

	if err := tx.QueryRow(ctx, `SELECT (commission / 100.0) AS net_commission FROM users WHERE affiliate_id = $1`, affiliateID).Scan(&commission); err != nil {
		return err
	}

	batch := &pgx.Batch{}
	query := `INSERT INTO transactions (
    transaction_id, amount, transaction_type, transaction_sub_type, status, 
    transaction_date, lead_id, lead_guid, affiliate_id, email, commission_amount
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, 
   	CASE WHEN $5 = 'Complete' THEN ROUND($2 * CAST($11 AS NUMERIC), 2) ELSE 0 END
)
ON CONFLICT (transaction_id) DO NOTHING 
RETURNING commission_amount
`

	var totalCommission float64

	for _, txn := range transactions {

		batch.Queue(query,
			txn.TransactionID,
			txn.Amount,
			txn.TransactionType,
			txn.TransactionSubType,
			txn.Status,
			txn.TransactionDate,
			txn.LeadID,
			txn.LeadGUID,
			affiliateID,
			email,
			commission,
		)
	}

	br := tx.SendBatch(ctx, batch)

	for range transactions {
		var insertedAmount float64

		err := br.QueryRow().Scan(&insertedAmount)
		if err == nil {

			totalCommission += insertedAmount

		} else if err != pgx.ErrNoRows {
			br.Close()
			return fmt.Errorf("error inserting transactions: %w", err)
		}
	}

	br.Close()

	log.Printf("Total new deposit amount for %s: %.2f", email, totalCommission)

	if totalCommission > 0 {
		_, err = tx.Exec(ctx, `UPDATE users SET balance = balance + $1 WHERE affiliate_id = $2`, totalCommission, affiliateID)
		if err != nil {
			return fmt.Errorf("error updating user balance: %w", err)
		}

		log.Printf("Updated balance for %s by %.2f", email, totalCommission)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	log.Printf("Successfully saved transactions and updated balance for %s", email)
	return nil
}

func (s *dataStore) GetweeklyStats(ctx context.Context, id string) (*models.Stats, error) {

	var stats models.Stats

	query := `WITH transaction_stats AS (
    SELECT 
        t.affiliate_id,
			ROUND(SUM(CASE WHEN t.transaction_type = 'Deposit' AND status = 'Complete' THEN COALESCE(t.amount, 0) ELSE 0 END), 2) AS total_deposits,
			ROUND(SUM(CASE WHEN t.transaction_type = 'Withdrawal' THEN COALESCE(t.amount, 0) ELSE 0 END), 2) AS total_withdrawals,
			SUM(commission_amount) AS total_commissions
    FROM transactions t
    LEFT JOIN users u ON t.affiliate_id = u.affiliate_id
    WHERE t.affiliate_id = $1 AND t.status = 'Complete'
    AND t.transaction_date >= date_trunc('week', NOW())::DATE
    GROUP BY t.affiliate_id
),
lead_stats AS (
    SELECT 
        l.affiliate_id, 
        COUNT(l.id) AS lead_count
    FROM leads l
    WHERE l.affiliate_id = $1
    AND l.registration_date >= date_trunc('week', NOW())::DATE
    GROUP BY l.affiliate_id
)

SELECT 
    COALESCE(ls.lead_count, 0) AS lead_count,
    COALESCE(ts.total_deposits, 0) AS total_deposits,
    COALESCE(ts.total_withdrawals, 0) AS total_withdrawals,
    COALESCE(ts.total_commissions, 0) AS total_commissions
FROM transaction_stats ts
FULL OUTER JOIN lead_stats ls ON ts.affiliate_id = ls.affiliate_id
`

	if err := s.db.QueryRow(ctx, query, id).Scan(
		&stats.Registrations,
		&stats.Deposits,
		&stats.Withdrawals,
		&stats.Commissions,
	); err != nil {
		log.Println(err)
		return nil, err
	}

	return &stats, nil

}

func (s *dataStore) GetNetStats(ctx context.Context, id string) (*models.Stats, error) {

	var stats models.Stats

	query := `WITH transaction_stats AS (
		SELECT 
			t.affiliate_id,
			ROUND(SUM(CASE WHEN t.transaction_type = 'Deposit' AND status = 'Complete' THEN COALESCE(t.amount, 0) ELSE 0 END), 2) AS total_deposits,
			ROUND(SUM(CASE WHEN t.transaction_type = 'Withdrawal' THEN COALESCE(t.amount, 0) ELSE 0 END), 2) AS total_withdrawals,
			SUM(commission_amount) AS total_commissions
		FROM transactions t
		LEFT JOIN users u ON t.affiliate_id = u.affiliate_id
		WHERE t.affiliate_id = $1 AND t.status = 'Complete'
		GROUP BY t.affiliate_id
	),
	lead_stats AS (
		SELECT 
			l.affiliate_id, 
			COUNT(l.id) AS lead_count
		FROM leads l
		WHERE l.affiliate_id = $1
		GROUP BY l.affiliate_id
	)
	
	SELECT 
		COALESCE(ls.lead_count, 0) AS lead_count,
		COALESCE(ts.total_deposits, 0) AS total_deposits,
		COALESCE(ts.total_withdrawals, 0) AS total_withdrawals,
		COALESCE(ts.total_commissions, 0) AS total_commissions
	FROM transaction_stats ts
	FULL OUTER JOIN lead_stats ls ON ts.affiliate_id = ls.affiliate_id
	`

	if err := s.db.QueryRow(ctx, query, id).Scan(
		&stats.Registrations,
		&stats.Deposits,
		&stats.Withdrawals,
		&stats.Commissions,
	); err != nil {
		log.Println(err)
		return nil, err
	}

	return &stats, nil

}

func (s *dataStore) GetMonthlyStats(ctx context.Context, id string) (*models.Stats, error) {

	var stats models.Stats

	query := `WITH transaction_stats AS (
    SELECT 
        t.affiliate_id,
			ROUND(SUM(CASE WHEN t.transaction_type = 'Deposit' AND status = 'Complete' THEN COALESCE(t.amount, 0) ELSE 0 END), 2) AS total_deposits,
			ROUND(SUM(CASE WHEN t.transaction_type = 'Withdrawal' THEN COALESCE(t.amount, 0) ELSE 0 END), 2) AS total_withdrawals,
			SUM(commission_amount) AS total_commissions
    FROM transactions t
    LEFT JOIN users u ON t.affiliate_id = u.affiliate_id
    WHERE t.affiliate_id = $1 AND t.status = 'Complete'
    AND t.transaction_date >= date_trunc('month', NOW())::DATE
    GROUP BY t.affiliate_id
),
lead_stats AS (
    SELECT 
        l.affiliate_id, 
        COUNT(l.id) AS lead_count
    FROM leads l
    WHERE l.affiliate_id = $1
    AND l.registration_date >= date_trunc('month', NOW())::DATE
    GROUP BY l.affiliate_id
)

SELECT 
    COALESCE(ls.lead_count, 0) AS lead_count,
    COALESCE(ts.total_deposits, 0) AS total_deposits,
    COALESCE(ts.total_withdrawals, 0) AS total_withdrawals,
    COALESCE(ts.total_commissions, 0) AS total_commissions
FROM transaction_stats ts
FULL OUTER JOIN lead_stats ls ON ts.affiliate_id = ls.affiliate_id
`

	if err := s.db.QueryRow(ctx, query, id).Scan(
		&stats.Registrations,
		&stats.Deposits,
		&stats.Withdrawals,
		&stats.Commissions,
	); err != nil {
		log.Println(err)
		return nil, err
	}

	return &stats, nil

}

func (s *dataStore) GetTransactions(ctx context.Context, id, from, to string) ([]models.CommissionTxn, error) {

	var txns []models.CommissionTxn

	query := `SELECT t.lead_id, l.first_name, l.country, t.email, 
	TO_CHAR(t.transaction_date, 'DD/MM/YYYY') AS txn_date_str, commission_amount, t.transaction_type 
FROM transactions t
LEFT JOIN leads l ON t.affiliate_id = l.affiliate_id
LEFT JOIN users u ON t.affiliate_id = u.affiliate_id
WHERE t.affiliate_id = $1
AND t.status = 'Complete'
AND t.transaction_date BETWEEN $2 AND $3
ORDER BY t.transaction_date DESC`

	rows, err := s.db.Query(ctx, query, id, from, to)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var txn models.CommissionTxn

		if err := rows.Scan(
			&txn.LeadID,
			&txn.Name,
			&txn.Country,
			&txn.Email,
			&txn.Date,
			&txn.Amount,
			&txn.TxnType,
		); err != nil {
			return nil, err
		}

		txns = append(txns, txn)
	}

	return txns, nil
}

func (s *dataStore) GetLatestFiveTransactions(ctx context.Context, id string) ([]models.CommissionTxn, error) {

	var txns []models.CommissionTxn

	query := `SELECT t.lead_id, l.first_name, l.country, t.email, 
	TO_CHAR(t.transaction_date, 'DD/MM/YYYY') AS txn_date_str, commission_amount, t.transaction_type 
FROM transactions t
LEFT JOIN leads l ON t.affiliate_id = l.affiliate_id
LEFT JOIN users u ON t.affiliate_id = u.affiliate_id
WHERE t.affiliate_id = $1
AND t.status = 'Complete'
AND t.transaction_type = 'Deposit'
ORDER BY t.transaction_date DESC
LIMIT 5`

	rows, err := s.db.Query(ctx, query, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var txn models.CommissionTxn

		if err := rows.Scan(
			&txn.LeadID,
			&txn.Name,
			&txn.Country,
			&txn.Email,
			&txn.Date,
			&txn.Amount,
			&txn.TxnType,
		); err != nil {
			return nil, err
		}

		txns = append(txns, txn)
	}

	return txns, nil
}

func (s *dataStore) GetLeaderboard(ctx context.Context) ([]models.Leaderboard, error) {

	var leaderboard []models.Leaderboard

	query := `SELECT 
    u.name, 
	u.country,
    ROUND(SUM(commission_amount),2) AS total_commissions
FROM transactions t
INNER JOIN users u ON u.affiliate_id = t.affiliate_id 
WHERE 
    t.status = 'Complete' 
    AND t.transaction_type = 'Deposit'
GROUP BY u.name, u.country
ORDER BY total_commissions DESC
LIMIT 50
`

	rows, err := s.db.Query(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var lb models.Leaderboard

		if err := rows.Scan(
			&lb.Name,
			&lb.Country,
			&lb.TotalCommissions,
		); err != nil {
			return nil, err
		}

		leaderboard = append(leaderboard, lb)

	}

	return leaderboard, nil
}

func (s *dataStore) GetBalance(ctx context.Context, id string) (float64, error) {

	var balance float64

	query := `SELECT balance FROM users WHERE affiliate_id = $1`

	if err := s.db.QueryRow(ctx, query, id).Scan(&balance); err != nil {
		return 0, err
	}

	return balance, nil

}
