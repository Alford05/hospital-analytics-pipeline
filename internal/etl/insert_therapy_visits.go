package etl

import (
	"database/sql"

	"hospital-analytics-pipeline/internal/models"
)

func InsertTherapyVisits(
	db *sql.DB,
	visits []models.TherapyVisit,
) error {

	query := `
	INSERT INTO therapy_visits (
		patient_id,
		therapist_id,
		visit_date,
		visit_type,
		duration_minutes
	)
	VALUES (
		$1, $2, $3, $4, $5
	)
	`

	for _, visit := range visits {
		_, err := db.Exec(
			query,
			visit.PatientID,
			visit.TherapistID,
			visit.VisitDate,
			visit.VisitType,
			visit.DurationMinutes,
		)

		if err != nil {
			return err
		}
	}

	return nil
}
