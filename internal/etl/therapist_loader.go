package etl

import (
	"encoding/csv"
	"os"

	"hospital-analytics-pipeline/internal/models"
)

func LoadTherapistsCSV(path string) ([]models.Therapist, error) {

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

	var therapists []models.Therapist

	for i, row := range rows {

		if i == 0 {
			continue
		}

		therapist := models.Therapist{
			TherapistName: row[0],
			Department:    row[1],
			HireDate:      row[2],
			Active:        row[3] == "true",
		}

		therapists = append(therapists, therapist)
	}

	return therapists, nil
}
