export class DoctorData {
    constructor(idDoctor, idUser, firstName, secondName, email, phoneNumber, specialization, isActive) {
      this.idDoctor = idDoctor;
      this.idUser = idUser;
      this.firstName = firstName;
      this.secondName = secondName;
      this.email = email;
      this.phoneNumber = phoneNumber;
      this.specialization = specialization;
      this.isActive = isActive;
    }
}