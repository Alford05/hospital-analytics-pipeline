package etl

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"hospital-analytics-pipeline/internal/models"
)

func LoadStaffingCSV(path string) ([]models.Staffing, error) {
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

	var staffing []models.Staffing

	for i, row := range rows {

		if i == 0 {
			continue
		}
		if len(row) < 4 {
			return nil, fmt.Errorf(
				"row %d: invalid row length",
				i+1,
			)
		}

		therapistID, err := strconv.Atoi(row[0])
		if err != nil {
			return nil, fmt.Errorf(
				"row %d: invalid therapist_id: %w",
				i+1,
				err,
			)
		}

		shiftHours, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			return nil, fmt.Errorf(
				"row %d: invalid shift hours: %w",
				i+1,
				err,
			)
		}

		patientLoad, err := strconv.Atoi(row[3])
		if err != nil {
			return nil, fmt.Errorf(
				"row %d: invalid patient load: %w",
				i+1,
				err,
			)
		}

		staff := models.Staffing{
			TherapistID: therapistID,
			ShiftDate:   row[1],
			ShiftHours:  shiftHours,
			PatientLoad: patientLoad,
		}

		staffing = append(staffing, staff)
	}

	return staffing, nil
}
