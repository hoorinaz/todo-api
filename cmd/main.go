package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Todo struct {
	gorm.Model
	Title       string
	Description string
	DueDate     time.Time
	IsDone      bool
}

var todo Todo

func main() {

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=hoorzy dbname=hoorzy password=sahar67 sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//Automatically migrate your schema, to keep your schema update to date. ONLY CREATE TABELS.
	// db.AutoMigrate(&Todo{})
	// db.Create(&Todo{
	// 	Title:       "4th thing to do",
	// 	Description: "4th description",
	// 	IsDone:      false,
	// })

	// db.Model(&Todo{}).DropColumn("Update")

	db.First(&todo, 2)
	// db.Model(&todo).Update("Title", "22222")
	db.Delete(&todo)

	fmt.Println("connect to db")

	http.HandleFunc("/", viewTodo)
	http.HandleFunc("/edit", editTodo)
	http.ListenAndServe(":8080", nil)

}

func viewTodo(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "View Hoorie Nazati")
}

func editTodo(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Edit Hoorie Nazati")

}
