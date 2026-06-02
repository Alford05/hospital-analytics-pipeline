package main

import (
	"database/sql"
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

	if err := loadPatients(conn); err != nil {
		panic(err)
	}

	if err := loadTherapists(conn); err != nil {
		panic(err)
	}

	if err := loadStaffing(conn); err != nil {
		panic(err)
	}

	if err := loadTherapyVisits(conn); err != nil {
		panic(err)
	}

	if err := loadDocumentationMetrics(conn); err != nil {
		panic(err)
	}

	fmt.Println("ETL completed successfully")

}

func loadPatients(conn *sql.DB) error {
	patients, err := etl.LoadPatientsCSV(
		"data/raw/patients.csv",
	)
	if err != nil {
		return err
	}
	if err := etl.InsertPatients(conn, patients); err != nil {
		return err
	}
	fmt.Printf(
		"Loaded %d patients successfully\n",
		len(patients),
	)
	return nil
}

func loadTherapists(conn *sql.DB) error {
	therapists, err := etl.LoadTherapistsCSV(
		"data/raw/therapists.csv",
	)
	if err != nil {
		return err
	}

	if err := etl.InsertTherapists(conn, therapists); err != nil {
		return err
	}

	fmt.Printf(
		"Loaded %d therapists successfully\n",
		len(therapists),
	)
	return nil
}

func loadStaffing(conn *sql.DB) error {
	staffing, err := etl.LoadStaffingCSV(
		"data/raw/staffing.csv",
	)
	if err != nil {
		return err
	}
	if err := etl.InsertStaffing(conn, staffing); err != nil {
		return err
	}
	fmt.Printf(
		"Loaded %d staffing records successfully\n",
		len(staffing),
	)
	return nil
}

func loadTherapyVisits(conn *sql.DB) error {
	visits, err := etl.LoadTherapyVisitsCSV(
		"data/raw/therapy_visits.csv",
	)
	if err != nil {
		return err
	}

	if err := etl.InsertTherapyVisits(conn, visits); err != nil {
		return err
	}

	fmt.Printf(
		"Loaded %d therapy visits successfully\n",
		len(visits),
	)
	return nil
}

func loadDocumentationMetrics(conn *sql.DB) error {
	metrics, err := etl.LoadDocumentationMetricsCSV(
		"data/raw/documentation_metrics.csv",
	)
	if err != nil {
		return err
	}

	if err := etl.InsertDocumentationMetrics(conn, metrics); err != nil {
		return err
	}

	fmt.Printf(
		"Loaded %d documentation metrics successfully\n",
		len(metrics),
	)

	return nil
}
