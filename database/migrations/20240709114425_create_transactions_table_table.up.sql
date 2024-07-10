CREATE TABLE transactions (
    id  INT NOT NULL  AUTO_INCREMENT  PRIMARY KEY,
    wallet_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    amount DECIMAL(15, 2) NOT NULL,
    type ENUM('debit', 'credit') NOT NULL,
    description TEXT,
    status BOOLEAN NOT NULL,
    reference VARCHAR(255) NOT NULL UNIQUE,
    transaction_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP  NULL,
    FOREIGN KEY (wallet_id) REFERENCES wallets(id) ON DELETE CASCADE
);