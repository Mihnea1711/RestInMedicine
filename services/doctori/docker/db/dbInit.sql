USE pdp_db;

CREATE TABLE IF NOT EXISTS doctor (
    id_doctor INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    id_user INT NOT NULL,
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
