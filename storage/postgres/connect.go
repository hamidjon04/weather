package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func Connect()(*sql.DB, error){
	connector := fmt.Sprintf("user = %s port = %s host = %s dbname = %s password = %s sslmode = disable", "postgres", "5432", "localhost", "weather", "hamidjon4424")
	db, err := sql.Open("postgres", connector)
	if err != nil{
		return nil, err
	}
	return db, nil
}