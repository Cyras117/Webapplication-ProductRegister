package dbcon

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaBD() *sql.DB {
	//conexao := "user dbname password host sslmode"
	conexao := "user=postgres dbname=loja password=root host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
