package models

import (
	"database/sql"
)

type ShortUrl struct {
	Path string
	Url  string
}

type ShortUrlModel struct {
	DB *sql.DB
}

func (model *ShortUrlModel) GetUrl(path string) (*ShortUrl, error) {
	query := `SELECT path, url FROM urls WHERE path=?`
	shortUrl := &ShortUrl{}
	err := model.DB.QueryRow(query, path).Scan(shortUrl)
	if err != nil {
		return nil, err
	}

	return shortUrl, nil
}
