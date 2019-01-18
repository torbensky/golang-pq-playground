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
	User User `db:"user"`
	Thing
}

func main() {
	db, err := sqlx.Connect("postgres", "password=playground user=postgres dbname=postgres sslmode=disable")
	must(err)
	defer db.Close()
	must(db.Ping())

	squirrel_sqlx_selectBasicStructScanExample(db)
	sqlx_joinExample(db)
	sqlx_joinWhereExample(db)
}

func sqlx_joinWhereExample(db *sqlx.DB) {
	query := `SELECT 
		things.*,
		users.id "user.id",
		users.name "user.name"	
	FROM
		user_things things JOIN users ON things.user_id = users.id
	WHERE
		users.id = $1;`

	rows, err := db.Queryx(query, 1)
	must(err)
	defer rows.Close()

	var utv UserThingView
	for rows.Next() {
		must(rows.StructScan(&utv))
		fmt.Println(utv)
	}
}

func sqlx_joinExample(db *sqlx.DB) {
	query := `SELECT 
		things.*,
		users.id "user.id",
		users.name "user.name"
	FROM
		user_things things JOIN users ON things.user_id = users.id;`

	rows, err := db.Queryx(query)
	must(err)
	defer rows.Close()

	var utv UserThingView
	for rows.Next() {
		must(rows.StructScan(&utv))
		fmt.Println(utv)
	}
}

func squirrel_sqlx_selectBasicStructScanExample(db *sqlx.DB) {
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
