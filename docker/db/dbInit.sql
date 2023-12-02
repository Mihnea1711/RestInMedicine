-- DROP USER 'mihnea_pos'@'%';
CREATE DATABASE pdp_db;
CREATE USER 'mihnea_pos'@'%' IDENTIFIED BY 'mihnea_pos';
GRANT ALL PRIVILEGES ON pdp_db.* TO 'mihnea_pos'@'%';
FLUSH PRIVILEGES;

USE pdp_db;

CREATE TABLE IF NOT EXISTS patient (
    id_patient INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    id_user INT NOT NULL UNIQUE,
    nume VARCHAR(255) NOT NULL,
    prenume VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    telefon VARCHAR(20) NOT NULL,
    cnp CHAR(13) NOT NULL UNIQUE ,
    data_nasterii DATE NOT NULL,
    is_active BOOLEAN DEFAULT false
);

-- Inserting two sample patients
INSERT INTO patient (id_user, nume, prenume, email, telefon, cnp, data_nasterii, is_active)
VALUES
    (1, 'Popescu', 'Ion', 'ion.popescu@example.com', '0712345678', '1010100890123', '2000-01-01', false),
    (2, 'Ionescu', 'Ana', 'ana.ionescu@example.com', '0712345679', '9150290210987', '1990-02-15', false);


CREATE TABLE IF NOT EXISTS doctor (
    id_doctor INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    id_user INT NOT NULL UNIQUE,
    nume VARCHAR(255) NOT NULL,
    prenume VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    telefon VARCHAR(20) NOT NULL,
    specializare VARCHAR(255) NOT NULL
);

-- Inserting two sample doctors
INSERT INTO doctor (id_user, nume, prenume, email, telefon, specializare)
VALUES
    (1, 'Popescu', 'Ion', 'ion.popescu@example.com', '+40123456789', 'Cardiologist'),
    (2, 'Ionescu', 'Ana', 'ana.ionescu@example.com', '+40123456790', 'Neurologist');


CREATE TABLE IF NOT EXISTS appointment (
    id_programare INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    id_patient INT NOT NULL,
    id_doctor INT NOT NULL,
    date DATE NOT NULL,
    status ENUM('onorata', 'programata', 'confirmata', 'neprezentata', 'anulata') NOT NULL,
    UNIQUE KEY unique_appointment (id_patient, id_doctor, date)
    -- FOREIGN KEY (id_patient) REFERENCES pacient(id_patient),
    -- FOREIGN KEY (id_doctor) REFERENCES doctor(id_doctor)
);


INSERT INTO appointment (id_patient, id_doctor, date, status) 
VALUES 
    (1, 1, '2023-11-10', 'onorata'),
    (2, 2, '2023-11-15', 'anulata');
