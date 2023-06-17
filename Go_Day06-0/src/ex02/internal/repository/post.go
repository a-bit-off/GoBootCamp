package repository

import (
	"context"
	"ex02/pkg/db"
	"time"
)

type Post struct {
	Id        int       `json:"id" db:"id"`
	UserLogin string    `json:"userLogin" db:"userLogin"`
	Created   time.Time `json:"created" db:"created"`
	Header    string    `json:"header" db:"header"`
	Content   string    `json:"content" db:"content"`
}

func (r *Repository) CreateTablePosts(ctx context.Context) (err error) {
	_, err = r.pool.Exec(ctx, db.QueryCreateTablePosts)
	if err != nil {
		return
	}
	return
}

func (r *Repository) AddNewPost(ctx context.Context, userLogin, header, content string) (err error) {
	_, err = r.pool.Exec(ctx, db.QueryAddNewPost, userLogin, time.Now(), header, content)
	if err != nil {
		return
	}
	return
}

func (r *Repository) GetNPosts(ctx context.Context, page int) (posts []Post, err error) {
	// TODO некорректный запрос, не может вывести оставшиеся записи
	row, err := r.pool.Query(ctx, db.QueryGetNPosts, page)
	if err != nil {
		return
	}
	defer row.Close()

	for row.Next() {
		var p Post
		err = row.Scan(&p.UserLogin, &p.Created, &p.Header, &p.Content)
		if err != nil {
			return
		}
		posts = append(posts, p)
	}
	return
}

func (r *Repository) GetAllPostsCount(ctx context.Context) (count int, err error) {
	row := r.pool.QueryRow(ctx, db.QueryGetPostsCount)
	err = row.Scan(&count)
	return
}
