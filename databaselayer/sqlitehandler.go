// package databaselayer

// import (
// 	"database/sql"

// 	_ "github.com/mattn/go-sqlite3"
// )

// type SQLLiteHandler struct {
// 	*SQLHandler
// }

// func NewSQLiteHandler(connection string) (*SQLLiteHandler, error) {
// 	db, err := sql.Open("sqlite3", connection)
// 	return &SQLLiteHandler{
// 		SQLHandler: &SQLHandler{
// 			DB: db,
// 		},
// 	}, err
// }
