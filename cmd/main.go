package main

import (
	"fmt"
	"hospital-analytics-pipeline/internal/db"
	"hospital-analytics-pipeline/internal/etl"
)

func main() {

	conn, err := db.Connect()

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	patients, err := etl.LoadPatientsCSV(
		"data/raw/patients.csv",
	)
	if err != nil {
		panic(err)
	}

	for _, patient := range patients {
        fmt.Printf(
            "Name: %s %s | Diagnosis: %s | Age: %d | Insurance: %s\n",
            patient.FirstName,
            patient.LastName,
            patient.Diagnosis,
            patient.Age,
            patient.InsuranceProvider,
        )
    }
}


