package repository

import (
	"context"
	"ex01/pkg/db"
)

type User struct {
	Id             int    `json:"id" db:"id"`
	Login          string `json:"login" db:"login"`
	HashedPassword string `json:"hashed_password" db:"hashed_password"`
	Name           string `json:"name" db:"name"`
	Surname        string `json:"surname" db:"surname"`
}

func (r *Repository) CreateTableUsers(ctx context.Context) (err error) {
	_, err = r.pool.Exec(ctx, db.QueryCreateTableUsers)
	if err != nil {
		return
	}
	return
}

func (r *Repository) Login(ctx context.Context, login, hashedPassword string) (u User, err error) {
	// делам поиск по бд
	row := r.pool.QueryRow(ctx, db.QueryFindUserByLoginAndPassword, login, hashedPassword)

	// сканируем данные, в случае если в бд будет отсутвовать user -> err != nil
	err = row.Scan(&u.Id, &u.Login, &u.Name, &u.Surname)
	return
}

func (r *Repository) AddNewUser(ctx context.Context, name, surname, login, hashedPassword string) (err error) {
	_, err = r.pool.Exec(ctx, db.QueryAddNewUser, login, hashedPassword, name, surname)
	if err != nil {
		return
	}
	return
}
