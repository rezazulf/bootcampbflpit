package employee

type Employee struct {
	ID       int
	Name     string
	Age      int
	Division string
}

var Employees = []Employee{
	{ID: 1, Name: "Test", Age: 23, Division: "IT"},
	{ID: 2, Name: "Testing", Age: 20, Division: "Finance"},
	{ID: 3, Name: "Tested", Age: 24, Division: "IT"},
}
