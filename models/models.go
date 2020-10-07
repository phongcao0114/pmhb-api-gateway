package models

type Employee struct {
	EmployeeID string  `json:"employee_id"`
	Position   string  `json:"position"`
	Name       Name    `json:"name"`
	Address    Address `json:"address"`
}

type Name struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Address struct {
	Ward     string `json:"ward"`
	District string `json:"district"`
	Province string `json:"province"`
}

func (e Employee) ID() string {
	return e.EmployeeID
}
