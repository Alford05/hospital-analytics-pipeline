# Hospital Analytics Pipeline

A healthcare-focused ETL and analytics project built with Go and PostgreSQL.

## Overview

Hospital Analytics Pipeline demonstrates the design and implementation of a backend data pipeline for healthcare operations analytics. The project extracts data from CSV files, performs validation and transformation, loads the data into a relational PostgreSQL database, and generates operational reports using SQL.

The project was inspired by real-world rehabilitation and acute care hospital workflows and focuses on therapist productivity, patient visits, staffing utilization, and documentation metrics.

## Architecture

```mermaid
flowchart TD
    A[Raw CSV Files] --> B[Go Cleaning Layer]
    B --> C[Cleaned CSV Files]
    C --> D[Go ETL Loaders]
    D --> E[(PostgreSQL Database)]
    E --> F[SQL Reporting Queries]
    E --> G[Indexes + Performance Analysis]
    F --> H[Operational Analytics]
    G --> H
```

## Technologies

* Go
* PostgreSQL
* Docker
* pgAdmin 4
* SQL
* Git / GitHub

## Features

### ETL Pipeline (Go-based)
Built a full extract-transform-load pipeline in Go to process healthcare-style datasets.

- Ingests multiple CSV data sources (patients, therapists, visits, staffing, documentation metrics)
- Performs data validation and transformation during ingestion
- Standardizes inconsistent clinical-style fields (diagnosis, insurance, visit types)
- Cleans and normalizes raw data before database insertion
- Implements reusable loader and transformation functions

### Relational Data Model
Designed a normalized PostgreSQL schema representing hospital workflows.

- Patient-to-visit relationships
- Therapist assignment tracking
- Documentation compliance tracking
- Staffing and workload modeling
- Enforced referential integrity using foreign keys

### Analytics & Reporting Layer
Built SQL-based reporting queries for operational healthcare analytics.

- Therapist productivity reporting
- Patient visit volume analysis
- Documentation delay and compliance tracking
- Staffing utilization metrics
- Patient load distribution across therapists

### Performance & Optimization
Explored database performance tuning techniques:

- Indexed high-traffic query columns
- Analyzed query execution plans using EXPLAIN ANALYZE
- Optimized joins across multi-table reporting queries

## Project Structure

```text
hospital-analytics-pipeline/
├── cmd/
├── data/
│   ├── raw/
│   ├── cleaned/
│   └── sample_reports/
├── internal/
│   ├── db/
│   ├── etl/
│   ├── models/
│   └── utils/
├── sql/
├── docs/
├── docker/
└── README.md
```

## ETL Workflow

```text
Raw CSV Files
      ↓
Data Validation & Cleaning
      ↓
Cleaned CSV Files
      ↓
Go ETL Loaders
      ↓
PostgreSQL Database
      ↓
Reporting Queries
      ↓
Operational Analytics
```

## Database Schema

Core relationships:

```text
patients
    ↑
therapy_visits
    ↑
documentation_metrics

therapists
    ↑
therapy_visits

therapists
    ↑
staffing
```

## Database Relationships

```mermaid
erDiagram
    PATIENTS ||--o{ THERAPY_VISITS : has
    THERAPISTS ||--o{ THERAPY_VISITS : performs
    THERAPY_VISITS ||--o{ DOCUMENTATION_METRICS : has
    THERAPISTS ||--o{ DOCUMENTATION_METRICS : completes
    THERAPISTS ||--o{ STAFFING : scheduled_for

    PATIENTS {
        int patient_id
        string first_name
        string last_name
        string diagnosis
        date admit_date
        date discharge_date
        boolean readmitted
        int age
        string insurance_provider
    }

    THERAPISTS {
        int therapist_id
        string therapist_name
        string department
        date hire_date
        boolean active
    }

    THERAPY_VISITS {
        int visit_id
        int patient_id
        int therapist_id
        date visit_date
        string visit_type
        int duration_minutes
        boolean notes_completed
    }

    DOCUMENTATION_METRICS {
        int note_id
        int therapist_id
        int visit_id
        int completion_delay_minutes
        timestamp completed_at
    }

    STAFFING {
        int staffing_id
        int therapist_id
        date shift_date
        decimal shift_hours
        int patient_load
    }
```

## Skills Demonstrated

* ETL design and implementation
* Relational database modeling
* SQL reporting and aggregation
* Foreign keys and indexing
* Query performance analysis
* Data validation and transformation
* Backend development using Go

## Future Improvements

* Scheduled ETL jobs
* REST API for analytics reports
* Dashboard visualization layer
* Additional healthcare operational metrics
* Automated testing
* CI/CD pipeline

## Author

Daniel Alford

* GitHub: https://github.com/Alford05
* LinkedIn: https://linkedin.com/in/alford-daniel/
