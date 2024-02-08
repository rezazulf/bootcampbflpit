package main

import (
	"database/sql"
	"fmt"

	"api_test/repository"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "bootcamp"
)

var (
	db  *sql.DB
	err error
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port =%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Sukses Konek Database")

	// var employee = models.Employee{}

	// sqlStatement := `INSERT INTO employees (full_name, email, age, division) VALUES ($1, $2, $3, $4) Returning *`

	// err = db.QueryRow(sqlStatement, "Reza Zulfi", "reza1@mail.com", "23", "IT").Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("New Employee Data : %+v\n", employee)

	repository.GetEmployees()
	repository.UpdateEmployee()
	repository.DeleteEmployee()
}
