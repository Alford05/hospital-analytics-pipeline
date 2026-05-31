package etl

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"hospital-analytics-pipeline/internal/models"
)

func LoadPatientsCSV(path string) ([]models.Patient, error) {

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

	var patients []models.Patient

	for i, row := range rows {

		if i == 0 {
			continue
		}
		if len(row) < 8 {
			return nil, fmt.Errorf("invalid row length: %v", row)
		}

		age, err := strconv.Atoi(row[6])
		if err != nil {
			return nil, err
		}

		patient := models.Patient{
			FirstName:         row[0],
			LastName:          row[1],
			Diagnosis:         row[2],
			AdmitDate:         row[3],
			DischargeDate:     row[4],
			Readmitted:        row[5] == "true",
			Age:               age,
			InsuranceProvider: row[7],
		}

		patients = append(patients, patient)
	}

	return patients, nil
}
