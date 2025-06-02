-- Users table
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    username VARCHAR(100) NOT NULL UNIQUE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- NAB history table
CREATE TABLE nab_history (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nab DECIMAL(10, 5) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- User investment table
CREATE TABLE user_investments (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    unit DECIMAL(20, 8) NOT NULL DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id)
);

-- Topup/withdrawal transactions
CREATE TABLE transactions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    type ENUM('TOPUP', 'WITHDRAW') NOT NULL,
    amount DECIMAL(20, 2) NOT NULL,
    unit DECIMAL(20, 8) NOT NULL,
    nab_at DECIMAL(10, 5) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id)
);