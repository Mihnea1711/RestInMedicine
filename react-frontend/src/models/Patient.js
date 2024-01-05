export class PatientData {
    constructor(idPatient, idUser, firstName, secondName, email, phoneNumber, cnp, birthDay, isActive) {
      this.idPatient = idPatient;
      this.idUser = idUser;
      this.firstName = firstName;
      this.secondName = secondName;
      this.email = email;
      this.phoneNumber = phoneNumber;
      this.cnp = cnp;
      this.birthDay = birthDay;
      this.isActive = isActive;
    }
}