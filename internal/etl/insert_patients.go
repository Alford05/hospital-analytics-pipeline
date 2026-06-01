package etl

import (
	"database/sql"

	"hospital-analytics-pipeline/internal/models"
)

func InsertPatients(
	db *sql.DB,
	patients []models.Patient,
) error {

	query := `
	INSERT INTO patients (
		first_name,
		last_name,
		diagnosis,
		admit_date,
		discharge_date,
		readmitted,
		age,
		insurance_provider
	)
	VALUES (
		$1, $2, $3, $4,
		$5, $6, $7, $8
	)
	`

	for _, patient := range patients {
		_, err := db.Exec (
			query,
			patient.FirstName,
			patient.LastName,
			patient.Diagnosis,
			patient.AdmitDate,
			patient.DischargeDate,
			patient.Readmitted,
			patient.Age,
			patient.InsuranceProvider,
		)

		if err != nil {
			return err
		}
	}

	return nil
}
