CREATE TABLE IF NOT EXISTS programari (
    id_programare INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    id_pacient INT NOT NULL,
    id_doctor INT NOT NULL,
    data DATE NOT NULL,
    status ENUM('onorata', 'neprezentata', 'anulata') NOT NULL,
    FOREIGN KEY (id_pacient) REFERENCES pacient(id_pacient),
    FOREIGN KEY (id_doctor) REFERENCES doctor(id_doctor)
);


INSERT INTO programare (id_pacient, id_doctor, data, status) 
VALUES 
    (1, 1, '2023-11-10', 'onorata'),
    (2, 2, '2023-11-15', 'anulata');
