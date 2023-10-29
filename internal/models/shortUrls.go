package models

import (
	"database/sql"
	"fmt"
)

type ShortUrl struct {
	Path string
	Url  string
}

type ShortUrlModel struct {
	DB *sql.DB
}

func (model *ShortUrlModel) GetUrl(path string) (*ShortUrl, error) {
	query := `SELECT original FROM urls WHERE Id=?`
	shortUrl := &ShortUrl{}
	row := model.DB.QueryRow(query, path)

	var original string
	err := row.Scan(&original)
	if err != nil {
		return nil, err
	}

	shortUrl.Path = original
	shortUrl.Url = path

	return shortUrl, nil
}

func (model *ShortUrlModel) CreatePath(url string) (*ShortUrl, error) {
	query := `INSERT INTO urls(original) VALUES(?);`
	res, err := model.DB.Exec(query, url)

	if err != nil {
		return nil, err
	}

	lastInsertId, err := res.LastInsertId()

	if err != nil {
		return nil, err
	}

	shortUrl := &ShortUrl{Url: fmt.Sprintf("https://localhost:3000/path/%d", lastInsertId), Path: url}

	return shortUrl, nil

}
