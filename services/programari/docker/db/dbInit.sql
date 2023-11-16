USE pdp_db;

CREATE TABLE IF NOT EXISTS appointment (
    id_programare INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    id_pacient INT NOT NULL,
    id_doctor INT NOT NULL,
    date DATE NOT NULL,
    status ENUM('onorata', 'neprezentata', 'anulata') NOT NULL
    -- FOREIGN KEY (id_pacient) REFERENCES pacient(id_pacient),
    -- FOREIGN KEY (id_doctor) REFERENCES doctor(id_doctor)
);


INSERT INTO appointment (id_pacient, id_doctor, date, status) 
VALUES 
    (1, 1, '2023-11-10', 'onorata'),
    (2, 2, '2023-11-15', 'anulata');
