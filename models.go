package main

import (
	"database/sql"
)

type Report struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// GetReports retrieves all reports from the database.
func GetReports(db *sql.DB) ([]Report, error) {
	rows, err := db.Query("SELECT id, title, content FROM reports")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []Report
	for rows.Next() {
		var report Report
		if err := rows.Scan(&report.ID, &report.Title, &report.Content); err != nil {
			return nil, err
		}
		reports = append(reports, report)
	}
	return reports, nil
}

// CreateReport inserts a new report into the database.
func CreateReport(db *sql.DB, report *Report) error {
	result, err := db.Exec("INSERT INTO reports (title, content) VALUES (?, ?)", report.Title, report.Content)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	report.ID = int(id)
	return nil
}

// UpdateReport updates an existing report in the database.
func UpdateReport(db *sql.DB, report *Report) error {
	_, err := db.Exec("UPDATE reports SET title = ?, content = ? WHERE id = ?", report.Title, report.Content, report.ID)
	return err
}

// DeleteReport deletes a report from the database by ID.
func DeleteReport(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM reports WHERE id = ?", id)
	return err
}
