package db

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/yunuscvlk/shellix-task/internal/models/sentence"
)

func (d *DB) Create(sentence string) sql.Result {
	stmt, err := d.DB.Prepare("INSERT INTO sentences(sentence) VALUES (?)")

	if err != nil {
		log.Println(err)

		return nil
	}

	result, err := stmt.Exec(sentence)

	if err != nil {
		log.Println(err)

		return nil
	}

	return result
}

func (d *DB) Read(id int, params map[string]string) *sentence.Model {
	rows, err := d.DB.Query("SELECT * from sentences WHERE id = ?", id)

	if err != nil {
		log.Println(err)

		return nil
	}

	defer rows.Close()

	for rows.Next() {
		row := new(sentence.Model)

		err = rows.Scan(&row.Id, &row.Sentence, &row.CreatedAt)

		if err != nil {
			log.Println(err)
	
			return nil
		}

		if params != nil && len(params) > 0 {
			for k, v := range params {
				key := fmt.Sprintf("${%s}", k)

				row.Sentence = strings.Replace(row.Sentence, key, v, -1)
			}
		}

		return row
	}

	return nil
}

func (d *DB) ReadAll(page int) []*sentence.Model {
	var (
		limit  int    = 10
		offset int    = 0
		query  string = "SELECT * from sentences"
	)

	if page > 0 {
		offset = (page - 1) * limit

		query += fmt.Sprintf(" LIMIT %d OFFSET %d", limit, offset)
	}

	rows, err := d.DB.Query(query)

	if err != nil {
		log.Println(err)

		return nil
	}

	defer rows.Close()

	var results []*sentence.Model

	for rows.Next() {
		row := new(sentence.Model)

		err = rows.Scan(&row.Id, &row.Sentence, &row.CreatedAt)

		if err != nil {
			log.Println(err)
	
			return nil
		}

		results = append(results, row)
	}

	return results
}

func (d *DB) Update(id int, sentence string) sql.Result {
	stmt, err := d.DB.Prepare("UPDATE sentences SET sentence = ? WHERE id = ?")

	if err != nil {
		log.Println(err)

		return nil
	}

	result, err := stmt.Exec(sentence, id)

	if err != nil {
		log.Println(err)

		return nil
	}

	return result
}

func (d *DB) Delete(id int) sql.Result {
	stmt, err := d.DB.Prepare("DELETE FROM sentences WHERE id = ?")

	if err != nil {
		log.Println(err)

		return nil
	}

	result, err := stmt.Exec(id)

	if err != nil {
		log.Println(err)

		return nil
	}

	return result
}

func (d *DB) DeleteOneDayOld() sql.Result {
	result, err := d.DB.Exec("DELETE FROM sentences WHERE created_at < datetime('now', '-1 day')")

	if err != nil {
		log.Println(err)

		return nil
	}
	
	return result
}