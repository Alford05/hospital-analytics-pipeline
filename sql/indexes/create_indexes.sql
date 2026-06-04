CREATE INDEX idx_therapy_visits_patient_id
ON therapy_visits(patient_id);

CREATE INDEX idx_therapy_visits_therapist_id
ON therapy_visits(therapist_id);

CREATE INDEX idx_documentation_metrics_visit_id
ON documentation_metrics(visit_id);

CREATE INDEX idx_documentation_metrics_therapist_id
ON documentation_metrics(therapist_id);

CREATE INDEX idx_staffing_therapist_id
ON staffing(therapist_id);

CREATE INDEX idx_staffing_shift_date
ON staffing(shift_date);
