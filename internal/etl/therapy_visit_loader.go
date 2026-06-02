package etl

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"hospital-analytics-pipeline/internal/models"
)

func LoadTherapyVisitsCSV(path string) ([]models.TherapyVisit, error) {

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

	var visits []models.TherapyVisit

	for i, row := range rows {

		if i == 0 {
			continue
		}
		if len(row) < 5 {
			return nil, fmt.Errorf(
				"row %d: invalid row length",
				i+1,
			)
		}

		patientID, err := strconv.Atoi(row[0])
		if err != nil {
			return nil, fmt.Errorf(
				"row %d: invalid patient_id: %w",
				i+1,
				err,
			)
		}

		therapistID, err := strconv.Atoi(row[1])
		if err != nil {
			return nil, fmt.Errorf(
				"row %d: invalid therapist_id: %w",
				i+1,
				err,
			)
		}

		durationMinutes, err := strconv.Atoi(row[4])
		if err != nil {
			return nil, fmt.Errorf(
				"row %d: invalid duration_minutes: %w",
				i+1,
				err,
			)
		}

		visit := models.TherapyVisit{
			PatientID:       patientID,
			TherapistID:     therapistID,
			VisitDate:       row[2],
			VisitType:       row[3],
			DurationMinutes: durationMinutes,
		}

		visits = append(visits, visit)
	}

	return visits, nil
}
