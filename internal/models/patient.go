package models

type Patient struct {
	FirstName         string
	LastName          string
	Diagnosis         string
	AdmitDate         string
	DischargeDate     string
	Readmitted        bool
	Age               int
	InsuranceProvider string
}
