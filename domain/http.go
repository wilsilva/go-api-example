package domain

import (
	"encoding/json"
	"net/http"

	"github.com/GuilhermeCaruso/bellt"
)

//People data
var People []Person

// GetPeople http handler
func GetPeople(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(People)
}

//GetPerson http handler
func GetPerson(w http.ResponseWriter, r *http.Request) {
	rv := bellt.RouteVariables(r)
	idPerson := rv.GetVar("id")

	var foundPerson Person
	for _, person := range People {
		if person.ID == idPerson {
			foundPerson = person
			break
		}
	}

	json.NewEncoder(w).Encode(foundPerson)
}
