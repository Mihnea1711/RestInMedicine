-- DROP USER 'mihnea_pos'@'%';
-- CREATE DATABASE idm_db;
-- CREATE USER 'mihnea_pos'@'%' IDENTIFIED BY 'mihnea_pos';
GRANT ALL PRIVILEGES ON idm_db.* TO 'mihnea_pos'@'%';
FLUSH PRIVILEGES;

USE idm_db;

-- Create User Table with a UNIQUE Username Constraint
CREATE TABLE IF NOT EXISTS User (
    IDUser INT AUTO_INCREMENT PRIMARY KEY,
    Username VARCHAR(255) NOT NULL UNIQUE,
    Password VARCHAR(255) NOT NULL
);

-- Create Role Table with CASCADE DELETE
CREATE TABLE IF NOT EXISTS Role (
    IDRole INT AUTO_INCREMENT PRIMARY KEY,
    IDUser INT,
    Role VARCHAR(255) NOT NULL,
    FOREIGN KEY (IDUser) REFERENCES User(IDUser) ON DELETE CASCADE
);

-- Create Trash Table
CREATE TABLE IF NOT EXISTS Trash (
    IDUser INT NOT NULL,
    Username VARCHAR(255) NOT NULL UNIQUE,
    Password VARCHAR(255) NOT NULL,
    Role VARCHAR(255) NOT NULL
);

-- Add Users
INSERT INTO User (Username, Password)
VALUES
    ('admin', 'admin_password'),
    ('doctor', 'doctor_password'),
    ('patient', 'patient_password');

-- Add Roles
INSERT INTO Role (IDUser, Role)
VALUES
    (1, 'admin'),
    (2, 'doctor'),
    (3, 'patient');
