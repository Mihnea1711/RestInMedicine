-- Create User Table with a UNIQUE Username Constraint
CREATE TABLE IF NOT EXISTS User (
    IDUser INT AUTO_INCREMENT PRIMARY KEY,
    Username VARCHAR(255) NOT NULL UNIQUE, -- Add UNIQUE constraint
    Password VARCHAR(255) NOT NULL,
    Token VARCHAR(255) NOT NULL
);

-- Create Role Table
CREATE TABLE IF NOT EXISTS Role (
    IDRole INT AUTO_INCREMENT PRIMARY KEY,
    IDUser INT,
    Role VARCHAR(255) NOT NULL,
    FOREIGN KEY (IDUser) REFERENCES User(IDUser)
);

-- Add Users
INSERT INTO User (Username, Password, Token)
VALUES
    ('admin', 'admin_password', 'admin_token'),
    ('doctor', 'doctor_password', 'doctor_token'),
    ('patient', 'patient_password', 'patient_token');

-- Add Roles
INSERT INTO Role (IDUser, Role)
VALUES
    (1, 'admin'),
    (2, 'doctor'),
    (3, 'patient');
