package repository

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/marumaro-git/study-go-api/go-api/domain/model"
)

type DBAccessor struct {
	Db *sql.DB
}

func (da DBAccessor) GetRecord(r *http.Request) ([]model.Record, error) {
	query := "SELECT id, item, content, study_date, create_datetime FROM record WHERE 1 = 1"
	params := []interface{}{}

	// パラメータ取得
	item := r.URL.Query().Get("item")
	if item != "" {
		query += " AND item = ?"
		params = append(params, item)
	}

	studyDate := r.URL.Query().Get("study_date")
	if studyDate != "" {
		query += " AND study_date = ?"
		params = append(params, studyDate)
	}

	// データ取得
	rows, err := da.Db.Query(query, params...)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer rows.Close()

	var records []model.Record
	for rows.Next() {
		var record model.Record
		err := rows.Scan(&record.Id, &record.Item, &record.Content, &record.StudyDate, &record.CreatedAt)
		if err != nil {
			panic(err.Error())
		}
		records = append(records, record)
	}

	return records, nil
}
