package etl

import (
	"database/sql"

	"hospital-analytics-pipeline/internal/models"
)

func InsertDocumentationMetrics(
	db *sql.DB,
	metrics []models.DocumentationMetric,
) error {
	query := `
	INSERT INTO documentation_metrics (
		therapist_id,
		visit_id,
		completion_delay_minutes,
		completed_at
	)
	VALUES (
		$1, $2, $3, $4
	)
	`

	for _, metric := range metrics {
		_, err := db.Exec(
			query,
			metric.TherapistID,
			metric.VisitID,
			metric.CompletionDelayMinutes,
			metric.CompletedAt,
		)

		if err != nil {
			return err
		}
	}

	return nil
}
