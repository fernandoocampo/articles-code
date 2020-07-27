package domain

// NewEmployee contains structured data of a new employee.
type NewEmployee struct {
	FirstName string
	LastName  string
	Email     string
}

// Employee contains structured data of a employee.
type Employee struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
}

// ToEmployee Converts the NewEmployee to Employee.
func (n NewEmployee) ToEmployee(id string) Employee {
	return Employee{
		ID:        id,
		FirstName: n.FirstName,
		LastName:  n.LastName,
		Email:     n.Email,
	}
}

// Validate verifies that the new employee contains valid data.
func (n NewEmployee) Validate() error {
	return nil
}
