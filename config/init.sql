CREATE DATABASE IF NOT EXISTS weather-db 
CHARACTER SET utf8mb4 
COLLATE utf8mb4_unicode_ci;

USE weather-db;

CREATE TABLE IF NOT EXISTS weather_queries (
    id INT AUTO_INCREMENT PRIMARY KEY,
    location VARCHAR(255) NOT NULL,
    service_1_temperature FLOAT NOT NULL,
    service_2_temperature FLOAT NOT NULL,
    request_count INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_location (location),
    INDEX idx_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- CREATE USER IF NOT EXISTS 'weather_user'@'%' IDENTIFIED BY 'weather_user_pass';
-- GRANT ALL PRIVILEGES ON weather-db.* TO 'weather_user'@'%';
-- FLUSH PRIVILEGES;

SELECT 'Database created successfully!' AS status;
SHOW TABLES;
DESCRIBE weather_queries;
