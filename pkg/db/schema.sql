CREATE TABLE IF NOT EXISTS admins (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL,
    password TEXT NOT NULL
)

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name TEXT,
    affiliate_id TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    commission INT NOT NULL DEFAULT 0,
    balance NUMERIC(18,2) NOT NULL DEFAULT 0,
    country TEXT,
    added_by INTEGER,
    client_link TEXT NOT NULL,
    sub_link TEXT NOT NULL,
    blocked BOOLEAN NOT NULL DEFAULT FALSE,
)

CREATE TABLE IF NOT EXISTS leads (
    id INTEGER PRIMARY KEY,
    first_name TEXT,
    last_name TEXT,
    updated TIMESTAMP,
    last_login_date TIMESTAMP,
    lead_guid UUID,
    country TEXT,
    city TEXT,
    time_zone TEXT,
    sales_status TEXT,
    language TEXT,
    business_unit TEXT,
    domain_name TEXT,
    is_qualified BOOLEAN,
    conversion_agent_id INT,
    retention_manager_id INT,
    vip_manager_id INT,
    closer_manager_id INT,
    conversion_agent_team TEXT,
    retention_manager_team TEXT,
    vip_manager_team TEXT,
    closer_manager_team TEXT,
    affiliate_id TEXT,
    affiliate_name TEXT,
    utm_campaign TEXT,
    utm_medium TEXT,
    utm_source TEXT,
    utm_term TEXT,
    referring_page TEXT,
    registration_date TIMESTAMP,
    account_creation_date TIMESTAMP,
    activation_date TIMESTAMP,
    fully_activation_date TIMESTAMP,
    sub_channel TEXT,
    channel_name TEXT,
    tl_name TEXT,
    tracking_link_id TEXT,
    deposited BOOLEAN,
    original_lead_id INT,
    original_by_name_lead_id INT,
    name_duplicates TEXT,
    email TEXT,
    offer_description TEXT,
    ip_address TEXT,
    landing_page TEXT
);



CREATE TABLE IF NOT EXISTS transactions (
    transaction_id INT PRIMARY KEY,
    amount NUMERIC(18,2) NOT NULL,
    transaction_type TEXT NOT NULL,
    transaction_sub_type TEXT,
    status TEXT NOT NULL,
    transaction_date TIMESTAMP NOT NULL,
    lead_id INT NOT NULL,
    lead_guid UUID NOT NULL,
    affiliate_id TEXT NOT NULL,
    email TEXT NOT NULL,
    commission_amount  NUMERIC(18,2) NOT NULL  DEFAULT 0
    -- CONSTRAINT fk_affiliate_id FOREIGN KEY (affiliate_id) REFERENCES users(affiliate_id) ON DELETE SET NULL
);


CREATE TABLE IF NOT EXISTS commissions (
    id SERIAL PRIMARY KEY,
    amount NUMERIC(18,2) NOT NULL,
    commission_amount NUMERIC(18,2) NOT NULL,
    transaction_type TEXT NOT NULL,
    lead_id INT NOT NULL,
    affiliate_id TEXT NOT NULL,
    original_affiliate_id TEXT NOT NULL,
    txn_id INT NOT NULL,
    commission_type TEXT NOT NULL CHECK(commission_type IN ('direct','sub')),
    transaction_date TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);


CREATE TABLE IF NOT EXISTS account_details (
    id SERIAL PRIMARY KEY,
    user INT NOT NULL,

)


CREATE TABLE IF NOT EXISTS payouts (
    id SERIAL PRIMARY KEY,
    amount INT NOT NULL,
    payout_type TEXT CHECK(payout_type IN ('payout','transfer')) NOT NULL,
    user_id INT NOT NULL,
    method TEXT NOT NULL,
    status TEXT CHECK(status IN ('PENDING','REJECTED','PAID')) NOT NULL DEFAULT 'PENDING',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
)


CREATE TABLE IF NOT EXISTS wallet_details (
    id SERIAL PRIMARY KEY,
    iban_number TEXT NOT NULL DEFAULT 'N/A',
    swift_code TEXT NOT NULL DEFAULT 'N/A',
    bank_name TEXT NOT NULL DEFAULT 'N/A',
    chain_name TEXT NOT NULL DEFAULT 'N/A',
    wallet_address TEXT NOT NULL DEFAULT 'N/A',
    user_id INT NOT NULL,

    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
)

ALTER TABLE wallet_details ADD CONSTRAINT unique_user_wallet UNIQUE (user_id);



CREATE INDEX idx_leads_affiliate_id ON leads(affiliate_id);

CREATE INDEX idx_transactions_email ON transactions(email);

CREATE INDEX idx_transactions_type_status_amount 
ON transactions(transaction_type, status, amount, commission_amount);

CREATE INDEX idx_transactions_deposits 
ON transactions(amount) 
WHERE transaction_type = 'Deposit' AND status = 'Complete';
