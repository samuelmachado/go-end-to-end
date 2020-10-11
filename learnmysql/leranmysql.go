package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type animal struct {
	id         int
	animalType string
	nickname   string
	zone       int
	age        int
}

func main() {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("select * from cursogo.animais where age > ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	animals := []animal{}
	for rows.Next() {
		a := animal{}
		err := rows.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age)
		if err != nil {
			log.Println(err)
			continue
		}
		animals = append(animals, a)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(animals)
	defer db.Close()

}
