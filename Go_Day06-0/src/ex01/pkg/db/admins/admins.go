/*
Пакет admins отвечает за регистрацию и авторизацию пользователей
Сохраняет данные в бд Admins

Методы:

	SignUpAdmin
	SignInAdmin

Возвращают:

	Ошибку при некорректных входных данных или при некорректном соединении с бд
*/
package admins

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

// данные для подключения к бд
const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "Admins"
)

// данные пользователя
type AdminData struct {
	Login    string  `json:"name"`
	Password string  `json:"password"`
	db       *sql.DB `json:"db"`
}

// регистрация нового пользователя
func (a AdminData) SignUpAdmin() error {
	err := a.connectToPostgreSQL()
	if err != nil {
		return err
	}
	err = a.createAdminsTable()
	if err != nil {
		return err
	}

	if unique, err := a.loginUniqueness(); err != nil {
		return err
	} else {
		if !unique {
			return errors.New("Логин занят!")
		}
	}

	err = a.insertAdminsTable()
	if err != nil {
		return err
	}
	a.db.Close()
	return nil
}

// вход
func (a AdminData) SignInAdmin() error {
	err := a.connectToPostgreSQL()
	if err != nil {
		return err
	}
	err = a.createAdminsTable()
	if err != nil {
		return err
	}

	if unique, err := a.loginUniqueness(); err != nil {
		return err
	} else {
		if unique {
			return errors.New("Неверный логин!")
		}
	}

	a.db.Close()
	return nil
}

// подключение к базе данных
func (a *AdminData) connectToPostgreSQL() error {
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)
	var err error
	a.db, err = sql.Open("postgres", psqlConn)
	if err != nil {
		return err
	}

	err = a.db.Ping()
	if err != nil {
		return err
	}
	return nil
}

// создание таблицы
func (a AdminData) createAdminsTable() error {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS admins (
		id SERIAL PRIMARY KEY,
		login VARCHAR(50),
		password VARCHAR(50)
	)
	`
	_, err := a.db.Exec(createTableQuery)
	if err != nil {
		return err
	}

	return nil
}

// проверка логина на уникальность
func (a *AdminData) loginUniqueness() (bool, error) {
	loginUniquenessQuery := `
		SELECT login FROM admins
		WHERE login = $1
	`
	req, err := a.db.Exec(loginUniquenessQuery, a.Login)
	if err != nil {
		return false, err
	}

	n, err := req.RowsAffected()
	if err != nil {
		return false, err
	}

	return n == 0, nil
}

// добавление нового пользователя в таблицу
func (a AdminData) insertAdminsTable() error {
	insertDataQuery := `
		INSERT INTO admins (login, password) VALUES ($1, $2)
	`
	_, err := a.db.Exec(insertDataQuery, a.Login, a.Password)
	if err != nil {
		return err
	}
	return nil
}
