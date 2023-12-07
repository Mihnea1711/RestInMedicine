# Create a logger class for better/easier logging

# Fronend ? Angular ?

# TODOs
## create some filter routes using query params (reduce the number of endpoints by including query params and multiple filter option) and some other routes if necessary
### /api/medical_office/physicians/{id}/patients
### /api/medical_office/patients/{id}/physicians
### /api/medical_office/physicians?specialization=...
### /api/medical_office/physicians/?name=...
### /api/medical_office/patients/{id}?date=...&type=...
### etc

## have some more constraints in the project architecture
### O consultatie poate fi creata doar daca exista o programare corespunzatoare
### Odata creata o consultatie, un potential client cu rol de pacient va putea vizualiza rezultatele prin intermediul sub-resursei programare corespunzatoare
### o programare pentru care câmpurile id_pacient, id_doctor s, i data coincid cu cele din consultat, ie

## update the role permissions
### administrator aplicatie
#### poate controla utilizatorii aplicatiei si datele referitoare la acestia
#### poate crea utilizatori noi cu rol de doctor;
#### nu poate vedea toate informat, iile din aplicatie (de exemplu, nu poate vizualiza consultatiile);

### doctor 
#### poate vizualiza informatiile tuturor pacientilor activi;
#### poate controla informatiile legate de consultatiile oferite pacientilor sai

### pacient 
#### poate realiza programari noi la doctori diferiti;
#### poate vizualiza istoricul sau medical (consultatiile precedente) .

## Generate Doc for project
### Generati documentul descriptiv al serviciului RESTful dezvoltat.
### Acest document trebuie s˘a urmeze specificat, iile OpenAPI[17] si sa fie disponibil în format JSON, prin metoda HTTP GET.








