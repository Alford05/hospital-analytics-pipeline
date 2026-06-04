package etl

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CleanTherapyVisitsCSV(inputPath string, outputPath string) error {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	reader := csv.NewReader(inputFile)

	rows, err := reader.ReadAll()
	if err != nil {
		return err
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	for i, row := range rows {
		if i == 0 {
			if err := writer.Write(row); err != nil {
				return err
			}
			continue
		}

		if len(row) < 5 {
			return fmt.Errorf("row %d: invalid row length", i+1)
		}

		patientID, err := strconv.Atoi(strings.TrimSpace(row[0]))
		if err != nil {
			return fmt.Errorf("row %d: invalid patient_id: %w", i+1, err)
		}

		therapistID, err := strconv.Atoi(strings.TrimSpace(row[1]))
		if err != nil {
			return fmt.Errorf("row %d: invalid therapist_id: %w", i+1, err)
		}

		visitDate := strings.TrimSpace(row[2])
		visitType := normalizeVisitType(row[3])

		durationMinutes, err := strconv.Atoi(strings.TrimSpace(row[4]))
		if err != nil {
			return fmt.Errorf("row %d: invalid duration_minutes: %w", i+1, err)
		}

		if durationMinutes <= 0 {
			return fmt.Errorf("row %d: duration_minutes must be greater than 0", i+1)
		}

		cleanedRow := []string{
			strconv.Itoa(patientID),
			strconv.Itoa(therapistID),
			visitDate,
			visitType,
			strconv.Itoa(durationMinutes),
		}

		if err := writer.Write(cleanedRow); err != nil {
			return err
		}
	}

	return nil
}

func normalizeVisitType(value string) string {
	cleaned := strings.ToLower(strings.TrimSpace(value))

	switch cleaned {
	case "eval", "evaluation":
		return "Evaluation"
	case "treat", "treatment", "tx":
		return "Treatment"
	case "dc", "d/c", "discharge":
		return "Discharge"
	default:
		return strings.Title(cleaned)
	}
}
