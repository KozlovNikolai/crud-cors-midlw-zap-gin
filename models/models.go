package models

type Person struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Employer struct {
	ID      int    `json:"id"`
	Company string `json:"company"`
	Person  Person `json:"person"`
}
