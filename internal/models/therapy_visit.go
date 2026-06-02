package models

type TherapyVisit struct {
	PatientID       int
	TherapistID     int
	VisitDate       string
	VisitType       string
	DurationMinutes int
}
