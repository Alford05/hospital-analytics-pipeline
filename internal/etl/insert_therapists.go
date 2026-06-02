package etl

import (
	"database/sql"

	"hospital-analytics-pipeline/internal/models"
)

func InsertTherapists(
	db *sql.DB,
	therapists []models.Therapist,
) error {

	query := `
	INSERT INTO therapists (
		therapist_name,
		department,
		hire_date,
		active
	)
	VALUES (
		$1, $2, $3, $4
	)
	`

	for _, therapist := range therapists {
		_, err := db.Exec(
			query,
			therapist.TherapistName,
			therapist.Department,
			therapist.HireDate,
			therapist.Active,
		)

		if err != nil {
			return err
		}
	}

	return nil
}
