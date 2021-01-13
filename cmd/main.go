package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"hotel/hotel/configs"
	"hotel/hotel/internal/app/room/serverRoom"

	"log"
	"net/http"
)

type database struct {
	db *sql.DB
}

// создаём структуру для нашего сервера
type server struct {
	data *sql.DB
}

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var s = &server{}

func main() {

	config, err := configs.NewConfig("../hotels.json")
	if err != nil {
		log.Fatal(err)
	}
	var myFlags arrayFlags
	defer s.data.Close()
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", config.User, config.Password, config.DBName)
	s.data, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	/*http.HandleFunc("/bookings/create", createBook)
	http.HandleFunc("/bookings/delete", deleteBook)
	http.HandleFunc("/bookings/list", listBook)

	http.HandleFunc("/rooms/delete", deleteRoom)
	http.HandleFunc("/rooms/list", listRoom)*/
	http.HandleFunc("/rooms/create", serverRoom.CreateRoom)
	log.Println("Success")
	err = http.ListenAndServe(config.Host, nil)
	if err != nil {
		log.Fatal(err)
	}
	flag.Var(&myFlags, "d", "New string for database")
	flag.Parse()
}
