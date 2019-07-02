package main

import (
	"fmt"
	"net/http"

	"github.com/GuilhermeCaruso/bellt"
	"github.com/wilsilva/api-example/domain"
)

func main() {
	domain.People = append(domain.People, domain.Person{ID: "1", FirstName: "William", LastName: "Silva", Age: 25})
	domain.People = append(domain.People, domain.Person{ID: "2", FirstName: "Inamacinho", LastName: "De Jesus", Age: 25})
	fmt.Println("Listen port :8000")
	router := bellt.NewRouter()
	router.HandleFunc("/api/person", domain.GetPeople, "GET")
	router.HandleFunc("/api/person/{id}", domain.GetPerson, "GET")
	http.ListenAndServe(":8000", nil)
}
