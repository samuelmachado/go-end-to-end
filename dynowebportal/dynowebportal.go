package dynowebportal

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// RunWebPortal starts running the dino web portal on address addr
func RunWebPortal(addr string) error {
	http.HandleFunc("/", foothandler)
	return http.ListenAndServe(addr, nil)

}

type animal struct {
	ID         string `json:"id"`
	Age        uint64 `json:"age"`
	AnimalType string `json:"animalType"`
	Nickname   string `json:"nickname"`
	Zone       uint64 `json:"zone"`
}

func foothandler(w http.ResponseWriter, r *http.Request) {
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
		err := rows.Scan(&a.ID, &a.AnimalType, &a.Nickname, &a.Zone, &a.Age)
		if err != nil {
			log.Println(err)
			continue
		}
		animals = append(animals, a)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(animals)

	// fmt.Println(animals)
	defer db.Close()
	// fmt.Fprintf(w, "welcome to the Dino web portal %s", r.RemoteAddr)
}
