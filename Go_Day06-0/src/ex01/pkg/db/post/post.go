/*
Пакет post отвечает за добавление новых постов
Сохраняет данные в бд PostTable

Методы:

	NewPost

Возвращает:

	Ошибку при некорректных входных данных или при некорректном соединении с бд
*/
package post

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// данные для подключения к бд
const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "PostDB"
)

// данные поста
type Post struct {
	Header  string  `json:"header"`
	Content string  `json:"content"`
	db      *sql.DB `json:"db"`
}

// создание нового поста
func (p Post) NewPost() error {
	err := p.connectToPostgreSQL()
	if err != nil {
		return err
	}

	err = p.createPostTable()
	if err != nil {
		return err
	}

	err = p.insertPost()
	if err != nil {
		return err
	}

	return nil
}

func (p Post) GetPosts() ([]Post, error) {
	err := p.connectToPostgreSQL()
	if err != nil {
		return []Post{}, err
	}

	err = p.createPostTable()
	if err != nil {
		return []Post{}, err
	}

	getPostsQuery := `
		SELECT header, content FROM PostTable
	`

	req, err := p.db.Query(getPostsQuery)
	if err != nil {
		return []Post{}, err
	}
	defer req.Close()

	var result []Post
	for req.Next() {
		res := Post{}
		err := req.Scan(&res.Header, &res.Content)
		if err != nil {
			return []Post{}, err
		}
		result = append(result, res)
	}
	return result, nil
}

// подключение к базе данных
func (p *Post) connectToPostgreSQL() error {
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)
	var err error
	p.db, err = sql.Open("postgres", psqlConn)
	if err != nil {
		return err
	}

	err = p.db.Ping()
	if err != nil {
		return err
	}
	return nil
}

// создание таблицы
func (p Post) createPostTable() error {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS PostTable (
		id SERIAL PRIMARY KEY,
		header VARCHAR(50),
		content VARCHAR
	)
	`
	_, err := p.db.Exec(createTableQuery)
	if err != nil {
		return err
	}

	return nil
}

// добавление нового поста в таблицу
func (p Post) insertPost() error {
	insertDataQuery := `
		INSERT INTO PostTable (header, content) VALUES ($1, $2)
	`
	_, err := p.db.Exec(insertDataQuery, p.Header, p.Content)
	if err != nil {
		return err
	}
	return nil
}
