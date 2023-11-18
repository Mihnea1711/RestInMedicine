use idm_db;

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


-- Add Users
INSERT INTO User (Username, Password)
VALUES
    ('admin', 'admin_password'),
    ('doctor', 'doctor_password'),
    ('patient', 'pacient_password');

-- Add Roles
INSERT INTO Role (IDUser, Role)
VALUES
    (1, 'admin'),
    (2, 'doctor'),
    (3, 'pacient');
