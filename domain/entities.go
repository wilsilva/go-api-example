package domain

// Person struct
type Person struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"fistname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Age       int    `json:"age,omitempty"`
}
