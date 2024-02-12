package main

import (
	"encoding/json"
	"fmt"
)

// type Employee struct {
// 	FullName string `json:"full_name`
// 	Email    string `json:"email"`
// 	Age      int    `json:"age"`
// }

type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {

	var object = []User{{"john wick", 23}, {"ethan hunt", 32}}

	var jsonData, err = json.Marshal(object)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var jsonString = string(jsonData)
	fmt.Println(jsonString)

	// var urlString = "https://www.developer.com:80/hello?name=airell&age=23"
	// var u, e = url.Parse(urlString)
	// if e != nil {
	// 	fmt.Println(e.Error())
	// 	return
	// }

	// fmt.Printf("url: %s\n", urlString)

	// fmt.Printf("protocol: %s\n", u.Scheme)
	// fmt.Printf("host: %s\n", u.Host)
	// fmt.Printf("path: %s\n", u.Path)

	// var name = u.Query()["name"][0]
	// var age = u.Query()["age"][0]
	// fmt.Printf("name: %s, age: %s\n", name, age)

	// var jsonString = `[
	// {
	// 	"full_name": "Airell Jordan",
	// 	"email": "airell@mail.com",
	// 	"age": 23
	// },
	// {
	// 	"full_name": "Ananda RHP",
	// 	"email": "ananda@mail.com",
	// 	"age": 23
	// }]
	// `

	// // var temp interface{}
	// var result []Employee
	// var err = json.Unmarshal([]byte(jsonString), &result)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// for i, v := range result {
	// 	fmt.Printf("Index %d: %+v\n", i+1, v)
	// }

	// var result = temp.(map[string]interface{})

	// fmt.Println("full_name:", result["full_name"])
	// fmt.Println("email:", result["email"])
	// fmt.Println("age:", result["age"])

	// fmt.Println("full_name:", result.FullName)
	// fmt.Println("email:", result.Email)
	// fmt.Println("age:", result.Age)
}
