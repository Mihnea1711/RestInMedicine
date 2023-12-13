-- DROP USER 'mihnea_pos'@'%';
CREATE DATABASE pdp_db;
CREATE USER 'mihnea_pos'@'%' IDENTIFIED BY 'mihnea_pos';
GRANT ALL PRIVILEGES ON pdp_db.* TO 'mihnea_pos'@'%';
FLUSH PRIVILEGES;

USE pdp_db;

CREATE TABLE IF NOT EXISTS patient (
    id_patient INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    id_user INT NOT NULL UNIQUE,
    first_name VARCHAR(255) NOT NULL,
    second_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone_number VARCHAR(20) NOT NULL,
    cnp CHAR(13) NOT NULL UNIQUE ,
    birth_day DATE NOT NULL,
    is_active BOOLEAN DEFAULT true
);

-- Inserting two sample patients
INSERT INTO patient (id_user, first_name, second_name, email, phone_number, cnp, birth_day, is_active)
VALUES
    (1, 'Popescu', 'Ion', 'ion.popescu@example.com', '0712345678', '1010100890123', '2000-01-01', true),
    (2, 'Ionescu', 'Ana', 'ana.ionescu@example.com', '0712345679', '9150290210987', '1990-02-15', false);


CREATE TABLE IF NOT EXISTS doctor (
    id_doctor INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    id_user INT NOT NULL UNIQUE,
    first_name VARCHAR(255) NOT NULL,
    second_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone_number VARCHAR(20) NOT NULL,
    specialization ENUM('Cardiology', 'Neurology', 'Orthopedics', 'Pediatrics', 'Dermatology', 'Radiology', 'Surgery') NOT NULL,
    is_active BOOLEAN DEFAULT true
);

-- Inserting two sample doctors
INSERT INTO doctor (id_user, first_name, second_name, email, phone_number, specialization)
VALUES
    (1, 'Popescu', 'Ion', 'ion.popescu@example.com', '+40123456789', 'Cardiology'),
    (2, 'Ionescu', 'Ana', 'ana.ionescu@example.com', '+40123456790', 'Neurology');


CREATE TABLE IF NOT EXISTS appointment (
    id_appointment INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    id_patient INT NOT NULL,
    id_doctor INT NOT NULL,
    date DATE NOT NULL,
    status ENUM('honored', 'scheduled', 'confirmed', 'not_present', 'canceled') NOT NULL,
    UNIQUE KEY unique_appointment (id_patient, id_doctor, date),
    FOREIGN KEY (id_patient) REFERENCES patient(id_patient),
    FOREIGN KEY (id_doctor) REFERENCES doctor(id_doctor)
);


INSERT INTO appointment (id_patient, id_doctor, date, status) 
VALUES 
    (1, 1, '2023-11-10', 'honored'),
    (2, 2, '2023-11-15', 'canceled');
