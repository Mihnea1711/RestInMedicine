CREATE TABLE IF NOT EXISTS pacient (
    cnp CHAR(13) PRIMARY KEY NOT NULL,
    id_user INT NOT NULL,
    nume VARCHAR(50) NOT NULL,
    prenume VARCHAR(50) NOT NULL,
    email VARCHAR(70) NOT NULL UNIQUE,
    telefon CHAR(10) NOT NULL,
    data_nasterii DATE NOT NULL,
    is_active BOOLEAN NOT NULL
);

-- Inserting two sample patients
INSERT INTO pacient (cnp, id_user, nume, prenume, email, telefon, data_nasterii, is_active)
VALUES
    ('1234567890123', 1, 'Popescu', 'Ion', 'ion.popescu@example.com', '0712345678', '2000-01-01', true),
    ('9876543210987', 2, 'Ionescu', 'Ana', 'ana.ionescu@example.com', '0712345679', '1990-02-15', true);
