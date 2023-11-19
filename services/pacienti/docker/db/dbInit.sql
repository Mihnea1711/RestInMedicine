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

