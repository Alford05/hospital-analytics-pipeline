# Hospital Analytics Pipeline

A healthcare-focused ETL and analytics project built with Go and PostgreSQL.

## Overview

Hospital Analytics Pipeline demonstrates the design and implementation of a backend data pipeline for healthcare operations analytics. The project extracts data from CSV files, performs validation and transformation, loads the data into a relational PostgreSQL database, and generates operational reports using SQL.

The project was inspired by real-world rehabilitation and acute care hospital workflows and focuses on therapist productivity, patient visits, staffing utilization, and documentation metrics.

## Technologies

* Go
* PostgreSQL
* Docker
* pgAdmin 4
* SQL
* Git / GitHub

## Features

### ETL Pipeline

Extract, Transform, and Load workflow implemented in Go.

Data sources:

* Patients
* Therapists
* Therapy Visits
* Documentation Metrics
* Staffing

Transformations include:

* Data validation
* Whitespace cleanup
* Standardized diagnosis values
* Standardized insurance provider values
* Visit type normalization
* Data quality checks

### Relational Database Design

Database tables:

* patients
* therapists
* therapy_visits
* documentation_metrics
* staffing

Features:

* Primary keys
* Foreign key relationships
* Indexed lookup columns
* Normalized schema

### Reporting

Examples of implemented reports:

* Therapist productivity
* Visits by diagnosis
* Documentation delay by therapist
* Patient load by therapist
* Staffing utilization
* Average visits per patient

### Performance Analysis

Performance tuning techniques explored:

* Index creation
* Query optimization
* EXPLAIN ANALYZE
* Query plan interpretation

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

## What I Learned

This project provided hands-on experience with:

* ETL design and implementation
* Relational database modeling
* SQL reporting and aggregation
* Foreign keys and indexing
* Query performance analysis
* Data validation and transformation
* Backend development using Go

## Future Improvements

Potential enhancements:

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
