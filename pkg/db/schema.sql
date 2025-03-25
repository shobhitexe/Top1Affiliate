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
    country TEXT,
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
    amount NUMERIC(18,8) NOT NULL,
    transaction_type VARCHAR(50) NOT NULL,
    transaction_sub_type VARCHAR(50),
    status VARCHAR(20) NOT NULL,
    transaction_date TIMESTAMP NOT NULL,
    lead_id INT NOT NULL,
    lead_guid UUID NOT NULL,
    affiliate_id TEXT NOT NULL,
    email TEXT NOT NULL,
    -- CONSTRAINT fk_affiliate_id FOREIGN KEY (affiliate_id) REFERENCES users(affiliate_id) ON DELETE SET NULL
);
