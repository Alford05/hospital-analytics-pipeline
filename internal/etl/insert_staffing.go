package etl

import (
	"database/sql"

	"hospital-analytics-pipeline/internal/models"
)

func InsertStaffing(
	db *sql.DB,
	staffing []models.Staffing,
) error {
	query := `
	INSERT INTO staffing (
	    therapist_id,
		shift_date,
		shift_hours,
		patient_load
	)
	VALUES (
	    $1, $2, $3, $4
	)
	`

	for _, staff := range staffing {
		_, err := db.Exec(
			query,
			staff.TherapistID,
			staff.ShiftDate,
			staff.ShiftHours,
			staff.PatientLoad,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
