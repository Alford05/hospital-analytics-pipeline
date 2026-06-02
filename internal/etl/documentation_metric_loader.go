package etl

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"hospital-analytics-pipeline/internal/models"
)

func LoadDocumentationMetricsCSV(path string) ([]models.DocumentationMetric, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var metrics []models.DocumentationMetric

	for i, row := range rows {
		if i == 0 {
			continue
		}

		if len(row) < 4 {
			return nil, fmt.Errorf("row %d: invalid row length", i+1)
		}

		therapistID, err := strconv.Atoi(row[0])
		if err != nil {
			return nil, fmt.Errorf("row %d: invalid therapist_id: %w", i+1, err)
		}

		visitID, err := strconv.Atoi(row[1])
		if err != nil {
			return nil, fmt.Errorf("row %d: invalid visit_id: %w", i+1, err)
		}

		delayMinutes, err := strconv.Atoi(row[2])
		if err != nil {
			return nil, fmt.Errorf("row %d: invalid completion_delay_minutes: %w", i+1, err)
		}

		metric := models.DocumentationMetric{
			TherapistID:            therapistID,
			VisitID:                visitID,
			CompletionDelayMinutes: delayMinutes,
			CompletedAt:            row[3],
		}

		metrics = append(metrics, metric)
	}

	return metrics, nil
}
