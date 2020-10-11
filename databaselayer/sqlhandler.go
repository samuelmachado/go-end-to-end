package databaselayer

import (
	"database/sql"
	"fmt"
	"log"
)

type SQLHandler struct {
	*sql.DB
}

//GetAvailableDynos is responsible por get dynos
func (handler *SQLHandler) GetAvailableDynos() ([]Animal, error) {
	return handler.sendQuery("select * from animais")
}

func (handler *SQLHandler) GetDynoByNickname(nickname string) (Animal, error) {
	row := handler.QueryRow(fmt.Sprintf("select * from animais where nickname = '%s'", nickname))
	a := Animal{}
	err := row.Scan(&a.ID, &a.AnimalType, &a.Nickname, &a.Zone, &a.Age)
	return a, err
}

func (handler *SQLHandler) GetDynosByType(dinoType string) ([]Animal, error) {
	return handle.sendQuery(fmt.Sprintf("select * from animais where Animal_type = '%s'", dinoType))
}

func (handler *SQLHandler) AddAnimal(a Animal) error {
	_, err := handler.Exec(fmt.Sprintf("insert into animais (Animal_type, nickname, zone, age) values ('%s', '%s', %d, %d)", a.AnimalType, a.Nickname, a.Zone, a.Age))
	return err
}
func (handler *SQLHandler) UpdadeAnimal(a Animal, name string) error {
	_, err := handler.Exec(fmt.Sprintf("Update animais set Animal_type='%s', nickname='%s', zone=%d, age=%d' where nickname = '%s'", a.AnimalType, a.Nickname, a.Zone, a.Age, name))
	return err
}
func (handler *SQLHandler) sendQuery(q string) ([]Animal, error) {
	Animals := []Animal{}
	rows, err := handler.Query(q)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		a := Animal{}
		err := rows.Scan(&a.ID, &a.AnimalType, &a.Nickname, &a.Zone, &a.Age)
		if err != nil {
			log.Println(err)
			continue
		}
		Animals = append(Animals, a)
	}
	return Animals, rows.Err()
}
