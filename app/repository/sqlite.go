package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteTodo struct {
	unimplemented
	db *sqlx.DB
}

func NewSqliteTodo() TodoStorager {
	db := initDB()
	return &SqliteTodo{db: db}
}

func initDB() *sqlx.DB {
	db, err := sqlx.Connect("sqlite3", "./todos.db")
	if err != nil {
		panic(err)
	}
	// создать таблицу todos, с полями согласно структуре entity.Todo
	// названия колонок таблицы присутствуют в теге db структуры entity.Todo
	//	CREATE TABLE IF NOT EXISTS
	schema := `
-- 		Создать таблицу todos
	`

	_, err = db.Exec(schema)
	if err != nil {
		panic(err)
	}

	return db
}
