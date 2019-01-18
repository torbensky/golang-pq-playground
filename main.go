package main

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

type Thing struct {
	ID          int    `db:"thing_id"`
	UserID      int    `db:"user_id"`
	Type        string `db:"type"`
	Description string `db:"description"`
}

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type UserThingView struct {
	User
	Things []Thing `db:"things"`
}

func main() {
	db, err := sqlx.Connect("postgres", "password=playground user=postgres dbname=postgres sslmode=disable")
	must(err)
	defer db.Close()
	must(db.Ping())

	query, _ := sq.Select("*").
		From("users").
		MustSql()

	rows, err := db.Queryx(query)
	must(err)
	defer rows.Close()

	var user User
	for rows.Next() {
		must(rows.StructScan(&user))
		fmt.Println(user)
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
