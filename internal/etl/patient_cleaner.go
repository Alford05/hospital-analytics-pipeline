package etl

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CleanPatientsCSV(inputPath string, outputPath string) error {
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

		if len(row) < 8 {
			return fmt.Errorf("row %d: invalid row length", i+1)
		}

		firstName := strings.TrimSpace(row[0])
		lastName := strings.TrimSpace(row[1])
		diagnosis := strings.ToUpper(strings.TrimSpace(row[2]))
		admitDate := strings.TrimSpace(row[3])
		dischargeDate := strings.TrimSpace(row[4])
		readmitted := strings.ToLower(strings.TrimSpace(row[5]))
		age := strings.TrimSpace(row[6])
		insuranceProvider := normalizeInsurance(row[7])

		if firstName == "" || lastName == "" {
			return fmt.Errorf("row %d: missing patient name", i+1)
		}

		ageInt, err := strconv.Atoi(age)
		if err != nil {
			return fmt.Errorf("row %d: invalid age: %w", i+1, err)
		}

		if ageInt < 0 || ageInt > 120 {
			return fmt.Errorf("row %d: unrealistic age: %d", i+1, ageInt)
		}

		cleanedRow := []string{
			firstName,
			lastName,
			diagnosis,
			admitDate,
			dischargeDate,
			readmitted,
			strconv.Itoa(ageInt),
			insuranceProvider,
		}

		if err := writer.Write(cleanedRow); err != nil {
			return err
		}
	}

	return nil
}

func normalizeInsurance(value string) string {
	cleaned := strings.ToLower(strings.TrimSpace(value))

	switch cleaned {
	case "unitedhealth", "united health":
		return "UnitedHealth"
	case "blue cross", "bluecross":
		return "Blue Cross"
	case "medicare":
		return "Medicare"
	case "aetna":
		return "Aetna"
	case "cigna":
		return "Cigna"
	default:
		return strings.Title(cleaned)
	}
}
