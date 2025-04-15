package store

import (
	"context"
	"fmt"
	"log"
	"math"
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
	SaveTransactionsAndUpdateBalanceWithdraw(ctx context.Context, transactions []models.Transaction, email, affiliateId string) error
	SaveTransactionsAndUpdateBalanceDeposit(ctx context.Context, transactions []models.Transaction, email string, affiliateID string) error

	GetweeklyStats(ctx context.Context, id string) (*models.Stats, error)
	GetNetStats(ctx context.Context, id string) (*models.Stats, error)
	GetMonthlyStats(ctx context.Context, id string) (*models.Stats, error)

	GetTransactions(ctx context.Context, id, from, to string) ([]models.CommissionTxn, error)
	GetLatestFiveTransactions(ctx context.Context, id string) ([]models.CommissionTxn, error)
	GetLeaderboard(ctx context.Context) ([]models.Leaderboard, error)

	GetBalance(ctx context.Context, id string) (float64, error)

	GetSubAffiliates(ctx context.Context, id string) ([]models.User, error)

	GetSubAffiliatePath(ctx context.Context, id string) ([]models.AffiliatePath, error)
	GetAllUsers(ctx context.Context, id string) ([]models.TreeNode, error)

	GetMonthlySalesOverview(ctx context.Context, id string) ([]models.MonthlySalesOverview, error)
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
GROUP BY l.affiliate_id, l.first_name, l.last_name, l.country, l.registration_date`

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

func (s *dataStore) GetweeklyStats(ctx context.Context, id string) (*models.Stats, error) {

	var stats models.Stats

	query := `WITH base_affiliate AS (
    SELECT $1::TEXT AS affiliate_id
),
transaction_stats AS (
    SELECT 
        c.affiliate_id,
        ROUND(SUM(CASE WHEN c.transaction_type = 'deposit' THEN COALESCE(c.amount, 0) ELSE 0 END), 2) AS total_deposits,
        ROUND(SUM(CASE WHEN c.transaction_type = 'withdraw' THEN COALESCE(c.amount, 0) ELSE 0 END), 2) AS total_withdrawals,
        SUM(commission_amount) AS total_commissions
    FROM commissions c
    WHERE c.original_affiliate_id = $1
      AND c.transaction_date >= date_trunc('week', NOW())::DATE
    GROUP BY c.affiliate_id
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
FROM base_affiliate ba
LEFT JOIN transaction_stats ts ON ba.affiliate_id = ts.affiliate_id
LEFT JOIN lead_stats ls ON ba.affiliate_id = ls.affiliate_id`

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

	query := `WITH base_affiliate AS (
    SELECT $1::TEXT AS affiliate_id
),
transaction_stats AS (
    SELECT 
        c.affiliate_id,
        ROUND(SUM(CASE WHEN c.transaction_type = 'deposit' THEN COALESCE(c.amount, 0) ELSE 0 END), 2) AS total_deposits,
        ROUND(SUM(CASE WHEN c.transaction_type = 'withdraw' THEN COALESCE(c.amount, 0) ELSE 0 END), 2) AS total_withdrawals,
        ROUND(SUM(
        CASE 
            WHEN c.transaction_type = 'deposit' THEN COALESCE(c.commission_amount, 0)
            WHEN c.transaction_type = 'withdraw' THEN -COALESCE(c.commission_amount, 0)
            ELSE 0
        END
    ), 2) AS total_commissions,
	COUNT(DISTINCT(lead_id)) AS ftds
    FROM commissions c
    WHERE c.original_affiliate_id = $1
    GROUP BY c.affiliate_id
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
    COALESCE(ts.total_commissions, 0) AS total_commissions,
	COALESCE(ts.ftds, 0) AS ftds
FROM base_affiliate ba
LEFT JOIN transaction_stats ts ON ba.affiliate_id = ts.affiliate_id
LEFT JOIN lead_stats ls ON ba.affiliate_id = ls.affiliate_id`

	if err := s.db.QueryRow(ctx, query, id).Scan(
		&stats.Registrations,
		&stats.Deposits,
		&stats.Withdrawals,
		&stats.Commissions,
		&stats.FTDS,
	); err != nil {
		log.Println(err)
		return nil, err
	}

	return &stats, nil

}

func (s *dataStore) GetMonthlyStats(ctx context.Context, id string) (*models.Stats, error) {

	var stats models.Stats

	query := `WITH base_affiliate AS (
    SELECT $1::TEXT AS affiliate_id
),
transaction_stats AS (
    SELECT 
        c.affiliate_id,
        ROUND(SUM(CASE WHEN c.transaction_type = 'deposit' THEN COALESCE(c.amount, 0) ELSE 0 END), 2) AS total_deposits,
        ROUND(SUM(CASE WHEN c.transaction_type = 'withdraw' THEN COALESCE(c.amount, 0) ELSE 0 END), 2) AS total_withdrawals,
        SUM(commission_amount) AS total_commissions
    FROM commissions c
    WHERE c.original_affiliate_id = $1
      AND c.transaction_date >= date_trunc('month', NOW())::DATE
    GROUP BY c.affiliate_id
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
FROM base_affiliate ba
LEFT JOIN transaction_stats ts ON ba.affiliate_id = ts.affiliate_id
LEFT JOIN lead_stats ls ON ba.affiliate_id = ls.affiliate_id`

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

	query := `SELECT 
    c.lead_id AS commission_id,
    CASE 
        WHEN c.commission_type = 'sub' THEN l.first_name || ' (sub)'
        ELSE l.first_name
    END AS first_name,
    l.country,
    TO_CHAR(c.transaction_date, 'DD/MM/YYYY') AS txn_date_str,
    c.commission_amount,
    CASE 
        WHEN c.transaction_type = 'deposit' THEN 'Deposit'
        ELSE 'Withdrawal'
    END AS transaction_type
FROM commissions c
LEFT JOIN leads l ON c.lead_id = l.id
WHERE c.affiliate_id = $1
AND c.transaction_date BETWEEN $2 AND $3
ORDER BY c.transaction_date DESC;
`

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
			&txn.Date,
			&txn.Amount,
			&txn.TxnType,
		); err != nil {

			log.Println(err)

			return nil, err
		}

		txns = append(txns, txn)
	}

	return txns, nil
}

func (s *dataStore) GetLatestFiveTransactions(ctx context.Context, id string) ([]models.CommissionTxn, error) {

	var txns []models.CommissionTxn

	query := `SELECT 
    c.lead_id AS commission_id,
    CASE 
        WHEN c.commission_type = 'sub' THEN l.first_name || ' (sub)'
        ELSE l.first_name
    END AS first_name,
    l.country,
    TO_CHAR(c.transaction_date, 'DD/MM/YYYY') AS txn_date_str,
    c.commission_amount,
    CASE 
        WHEN c.transaction_type = 'deposit' THEN 'Deposit'
        ELSE 'Withdrawal'
    END AS transaction_type
FROM commissions c
LEFT JOIN leads l ON c.lead_id = l.id
WHERE c.affiliate_id = $1
ORDER BY c.transaction_date DESC
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
    u.affiliate_id,
    u.name, 
    u.country,
    ROUND(SUM(
        CASE 
            WHEN c.transaction_type = 'deposit' THEN COALESCE(c.commission_amount, 0)
            WHEN c.transaction_type = 'withdraw' THEN -COALESCE(c.commission_amount, 0)
            ELSE 0
        END
    ), 2) AS total_commissions
FROM commissions c
INNER JOIN users u ON u.affiliate_id = c.affiliate_id 
GROUP BY u.affiliate_id, u.name, u.country
ORDER BY total_commissions DESC
LIMIT 50`

	rows, err := s.db.Query(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var lb models.Leaderboard

		if err := rows.Scan(
			&lb.AffiliateId,
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

func (s *dataStore) GetSubAffiliates(ctx context.Context, id string) ([]models.User, error) {

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

func (s *dataStore) GetAllUsers(ctx context.Context, id string) ([]models.TreeNode, error) {

	query := `WITH RECURSIVE users_hierarchy AS (
    SELECT id, affiliate_id, name, country, commission, added_by, 0 AS depth
    FROM users 
    WHERE id = $1

    UNION ALL

    SELECT u.id, u.affiliate_id, u.name, u.country, u.commission, u.added_by, uh.depth + 1
    FROM users u
    INNER JOIN users_hierarchy uh ON u.added_by = uh.id
    WHERE uh.depth < 100
)
SELECT id, affiliate_id, name, country, commission, added_by, depth
FROM users_hierarchy`

	rows, err := s.db.Query(ctx, query, id)

	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var users []models.TreeNode
	for rows.Next() {
		var u models.TreeNode
		if err := rows.Scan(&u.ID, &u.AffiliateID, &u.Name, &u.Country, &u.Commission, &u.AddedBy, &u.Depth); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		users = append(users, u)
	}

	return users, nil
}

func (s *dataStore) GetSubAffiliatePath(ctx context.Context, id string) ([]models.AffiliatePath, error) {

	var path []models.AffiliatePath

	query := `WITH RECURSIVE referral_path AS (
    SELECT id, name, added_by, 0 AS depth
    FROM users 
    WHERE id = $1
    UNION ALL
    SELECT u.id, u.name, u.added_by, rp.depth + 1
    FROM users u
    INNER JOIN referral_path rp ON u.id = rp.added_by
)
SELECT * FROM referral_path ORDER BY depth DESC
`

	rows, err := s.db.Query(ctx, query, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var p models.AffiliatePath

		if err := rows.Scan(&p.ID, &p.Name, &p.AddedBy, &p.Depth); err != nil {
			return nil, err
		}

		path = append(path, p)
	}

	return path, nil
}

func (s *dataStore) GetMonthlySalesOverview(ctx context.Context, id string) ([]models.MonthlySalesOverview, error) {

	var sales []models.MonthlySalesOverview

	query := `SELECT 
    TRIM(TO_CHAR(transaction_date, 'Month')) AS month,
    SUM(CASE WHEN transaction_type = 'deposit' THEN commission_amount ELSE 0 END) AS deposit,
    SUM(CASE WHEN transaction_type = 'withdraw' THEN amount ELSE 0 END) AS withdrawal,
	SUM(CASE WHEN transaction_type = 'deposit' THEN commission_amount ELSE -commission_amount END) AS commission
FROM commissions
WHERE affiliate_id = $1
GROUP BY TO_CHAR(transaction_date, 'Month'), EXTRACT(MONTH FROM transaction_date)
ORDER BY EXTRACT(MONTH FROM transaction_date)`

	rows, err := s.db.Query(ctx, query, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var s models.MonthlySalesOverview

		if err := rows.Scan(&s.Month, &s.Deposits, &s.Withdrawals, &s.Commission); err != nil {
			return nil, err
		}

		sales = append(sales, s)
	}

	return sales, nil

}

// For Processing Data From API

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

func (s *dataStore) SaveTransactionsAndUpdateBalanceWithdraw(ctx context.Context, transactions []models.Transaction, email string, affiliateID string) error {
	if len(transactions) == 0 {
		return nil
	}

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	var commissionRate float64
	err = tx.QueryRow(ctx, `SELECT (commission / 100.0) FROM users WHERE affiliate_id = $1`, affiliateID).Scan(&commissionRate)
	if err != nil {
		return fmt.Errorf("error fetching commission rate: %w", err)
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
	RETURNING commission_amount, amount, status, transaction_id`

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
			commissionRate,
		)
	}

	br := tx.SendBatch(ctx, batch)

	type txnResult struct {
		amount      float64
		commission  float64
		status      string
		transaction models.Transaction
		txnID       string
		txnDate     string
	}

	var results []txnResult

	for _, txn := range transactions {
		var amount, commission float64
		var status, txnID string

		err := br.QueryRow().Scan(&commission, &amount, &status, &txnID)
		if err != nil && err != pgx.ErrNoRows {
			br.Close()
			return fmt.Errorf("error reading batch result: %w", err)
		}

		absCommission := math.Abs(commission)

		if err == nil && status == "Complete" && absCommission > 0 {
			results = append(results, txnResult{
				amount:      amount,
				commission:  absCommission,
				status:      status,
				transaction: txn,
				txnID:       txnID,
				txnDate:     txn.TransactionDate,
			})
		}
	}

	br.Close()

	for _, r := range results {

		absAmount := math.Abs(r.amount)

		_, err := tx.Exec(ctx, `UPDATE users SET balance = balance - $1 WHERE affiliate_id = $2`, r.commission, affiliateID)
		if err != nil {
			return fmt.Errorf("error updating user balance: %w", err)
		}

		if err := s.RecordTransaction(ctx, absAmount, r.commission, "withdraw", affiliateID, r.txnID, "direct", affiliateID, r.txnDate, r.transaction.LeadID); err != nil {
			return fmt.Errorf("error recording commission: %w", err)
		}

		if err := s.distributeCommissionWithdraw(ctx, tx, absAmount, &affiliateID, r.txnID, r.transaction.LeadID, affiliateID, r.txnDate); err != nil {
			return fmt.Errorf("error distributing commission: %w", err)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	log.Printf("Successfully saved transactions and updated balance for %s", email)
	return nil
}

func (s *dataStore) distributeCommissionWithdraw(ctx context.Context, tx pgx.Tx, commissionAmount float64, parentID *string, txnId string, leadId int, originalAffiliateId, txnDate string) error {
	if parentID == nil || *parentID == "N/A" {
		return nil
	}

	var parentCommission float64
	var grandParentID *string

	err := tx.QueryRow(ctx, `
	SELECT 
		COALESCE(uu.commission / 100.0, 0) AS net_commission, 
		COALESCE(uu.affiliate_id, 'N/A') AS grandParentID
	FROM users u
	LEFT JOIN users uu ON uu.id = u.added_by 
	WHERE u.affiliate_id = $1`, *parentID).Scan(&parentCommission, &grandParentID)

	if err != nil {
		return err
	}

	parentCommissionAmount := commissionAmount * parentCommission

	if parentCommissionAmount > 0 {
		_, err = tx.Exec(ctx, `UPDATE users SET balance = balance - $1 WHERE affiliate_id = $2`, parentCommissionAmount, *grandParentID)
		if err != nil {
			return fmt.Errorf("error updating parent balance: %w", err)
		}

		if err := s.RecordTransaction(ctx, commissionAmount, parentCommissionAmount, "withdraw", *grandParentID, txnId, "sub", originalAffiliateId, txnDate, leadId); err != nil {
			return fmt.Errorf("error recording transaction: %w", err)

		}

		log.Printf("Updated parent %s balance by %.2f", *grandParentID, parentCommissionAmount)

		err = s.distributeCommissionWithdraw(ctx, tx, commissionAmount, grandParentID, txnId, leadId, originalAffiliateId, txnDate)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *dataStore) SaveTransactionsAndUpdateBalanceDeposit(ctx context.Context, transactions []models.Transaction, email string, affiliateID string) error {
	if len(transactions) == 0 {
		return nil
	}

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	var commissionRate float64
	err = tx.QueryRow(ctx, `SELECT (commission / 100.0) FROM users WHERE affiliate_id = $1`, affiliateID).Scan(&commissionRate)
	if err != nil {
		return fmt.Errorf("error fetching commission rate: %w", err)
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
	RETURNING commission_amount, amount, status, transaction_id`

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
			commissionRate,
		)
	}

	br := tx.SendBatch(ctx, batch)

	type txnResult struct {
		amount      float64
		commission  float64
		status      string
		transaction models.Transaction
		txnID       string
		txnDate     string
	}

	var results []txnResult

	for _, txn := range transactions {
		var amount, commission float64
		var status, txnID string

		err := br.QueryRow().Scan(&commission, &amount, &status, &txnID)
		if err != nil && err != pgx.ErrNoRows {
			br.Close()
			return fmt.Errorf("error reading batch result: %w", err)
		}

		if err == nil && status == "Complete" && commission > 0 {
			results = append(results, txnResult{
				amount:      amount,
				commission:  commission,
				status:      status,
				transaction: txn,
				txnID:       txnID,
				txnDate:     txn.TransactionDate,
			})
		}
	}

	br.Close()

	for _, r := range results {
		_, err := tx.Exec(ctx, `UPDATE users SET balance = balance + $1 WHERE affiliate_id = $2`, r.commission, affiliateID)
		if err != nil {
			return fmt.Errorf("error updating user balance: %w", err)
		}

		if err := s.RecordTransaction(ctx, r.amount, r.commission, "deposit", affiliateID, r.txnID, "direct", affiliateID, r.txnDate, r.transaction.LeadID); err != nil {
			return fmt.Errorf("error recording commission: %w", err)
		}

		if err := s.distributeCommissionDeposit(ctx, tx, r.amount, &affiliateID, r.txnID, r.transaction.LeadID, affiliateID, r.txnDate); err != nil {
			return fmt.Errorf("error distributing commission: %w", err)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	log.Printf("Successfully saved transactions and updated balance for %s", email)
	return nil
}

func (s *dataStore) distributeCommissionDeposit(ctx context.Context, tx pgx.Tx, commissionAmount float64, parentID *string, txnId string, leadId int, originalAffiliateId, txnDate string) error {
	if parentID == nil || *parentID == "N/A" {
		return nil
	}

	var parentCommission float64
	var grandParentID *string

	err := tx.QueryRow(ctx, `
	SELECT 
		COALESCE(uu.commission / 100.0, 0) AS net_commission, 
		COALESCE(uu.affiliate_id, 'N/A') AS grandParentID
	FROM users u
	LEFT JOIN users uu ON uu.id = u.added_by 
	WHERE u.affiliate_id = $1`, *parentID).Scan(&parentCommission, &grandParentID)

	if err != nil {
		return err
	}

	parentCommissionAmount := commissionAmount * parentCommission

	if parentCommissionAmount > 0 {
		_, err = tx.Exec(ctx, `UPDATE users SET balance = balance + $1 WHERE affiliate_id = $2`, parentCommissionAmount, *grandParentID)
		if err != nil {
			return fmt.Errorf("error updating parent balance: %w", err)
		}

		if err := s.RecordTransaction(ctx, commissionAmount, parentCommissionAmount, "deposit", *grandParentID, txnId, "sub", originalAffiliateId, txnDate, leadId); err != nil {
			return fmt.Errorf("error recording transaction: %w", err)
		}

		log.Printf("Updated parent %s balance by %.2f", *grandParentID, parentCommissionAmount)

		err = s.distributeCommissionDeposit(ctx, tx, commissionAmount, grandParentID, txnId, leadId, originalAffiliateId, txnDate)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *dataStore) RecordTransaction(ctx context.Context, amount, commission float64, txnType, affiliateId, txnId, commissiontype, originalAffiliateId, txnDate string, leadId int) error {

	query := `INSERT INTO commissions 
	(amount, commission_amount, transaction_type, lead_id, affiliate_id, txn_id, commission_type, original_affiliate_id, transaction_date) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	if _, err := s.db.Exec(ctx, query,
		amount,
		commission,
		txnType,
		leadId,
		affiliateId,
		txnId,
		commissiontype,
		originalAffiliateId,
		txnDate,
	); err != nil {

		log.Println(err)
		return err
	}

	return nil
}
