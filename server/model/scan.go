package model

import "database/sql"

func ScanArticle(r *sql.Row) (Article, error) {
	var s Article
	if err := r.Scan(
		&s.ID,
		&s.Title,
		&s.PostedAt,
		&s.URL,
		&s.Text,
		&s.Company,
	); err != nil {
		return Article{}, err
	}
	return s, nil
}

func ScanArticles(rs *sql.Rows) ([]Article, error) {
	structs := []Article{}
	var err error
	for rs.Next() {
		var s Article
		if err = rs.Scan(
			&s.ID,
			&s.Title,
			&s.PostedAt,
			&s.URL,
			&s.Text,
			&s.Company,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanCompanies(rs *sql.Rows) ([]Company, error) {
	structs := []Company{}
	var err error
	for rs.Next() {
		var s Company
		if err = rs.Scan(
			&s.ID,
			&s.Company,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}
