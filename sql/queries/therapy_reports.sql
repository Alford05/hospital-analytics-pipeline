1. Visits by type 
SELECT
    visit_type,
    COUNT(*) AS visit_count
FROM therapy_visits
GROUP BY visit_type
ORDER BY visit_count DESC;

2. Therapist productivity
SELECT
    t.therapist_name,
    COUNT(v.visit_id) AS total_visits,
    SUM(v.duration_minutes) AS total_minutes,
    ROUND(AVG(v.duration_minutes), 2) AS avg_visit_minutes
FROM therapists t
JOIN therapy_visits v
    ON t.therapist_id = v.therapist_id
GROUP BY t.therapist_name
ORDER BY total_minutes DESC;

3. Therapy time by diagnosis 
SELECT
    p.diagnosis,
    COUNT(v.visit_id) AS total_visits,
    SUM(v.duration_minutes) AS total_minutes
FROM patients p
JOIN therapy_visits v
    ON p.patient_id = v.patient_id
GROUP BY p.diagnosis
ORDER BY total_minutes DESC;

4. Average visits per patient 
SELECT
    ROUND(AVG(visit_count), 2) AS avg_visits_per_patient
FROM (
    SELECT
        patient_id,
        COUNT(*) AS visit_count
    FROM therapy_visits
    GROUP BY patient_id
) patient_visit_counts;

5. Therapist Workload vs Productivity
SELECT
    t.therapist_name,
    AVG(s.patient_load) AS avg_patient_load,
    COUNT(v.visit_id) AS total_visits,
    SUM(v.duration_minutes) AS total_treatment_minutes
FROM therapists t
JOIN staffing s
    ON t.therapist_id = s.therapist_id
LEFT JOIN therapy_visits v
    ON t.therapist_id = v.therapist_id
GROUP BY t.therapist_name
ORDER BY total_treatment_minutes DESC;

6. Documentation Delay vs Patient Load
SELECT
    t.therapist_name,
    AVG(s.patient_load) AS avg_patient_load,
    ROUND(AVG(d.completion_delay_minutes), 2) AS avg_doc_delay
FROM therapists t
JOIN staffing s
    ON t.therapist_id = s.therapist_id
JOIN documentation_metrics d
    ON t.therapist_id = d.therapist_id
GROUP BY t.therapist_name
ORDER BY avg_doc_delay DESC;

