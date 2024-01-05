export class ConsultationData {
    constructor(idConsultation, idPatient, idDoctor, date, diagnostic, investigations) {
      this.idConsultation = idConsultation;
      this.idPatient = idPatient;
      this.idDoctor = idDoctor;
      this.date = date;
      this.diagnostic = diagnostic;
      this.investigations = investigations;
    }
}

// class Investigation {
//   constructor(idInvestigatie, name, processingTime, result) {
//     this.idInvestigatie = idInvestigatie;
    // this.name = name;
    // this.processingTime = processingTime;
    // this.result = result;
//   }
// }