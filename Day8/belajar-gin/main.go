package main

import (
	"belajar-gin/routers"

	_ "github.com/lib/pq"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "root"
// 	dbname   = "bootcamp"
// )

// var (
// 	db  *sql.DB
// 	err error
// )

func main() {
	// psqlInfo := fmt.Sprintf("host=%s port =%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// db, err = sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()

	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Sukses Konek Database")
	var PORT = ":3000"

	routers.StartServer().Run(PORT)
}
