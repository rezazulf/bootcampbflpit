package main

import "fmt"

type Organization struct {
	name   string
	alamat string
}
type Employee struct {
	nik          int
	name         string
	age          int
	divisi       string
	organization Organization
}

func main() {
	var pekerja = []Employee{
		{nik: 111111, name: "Muh", age: 23, divisi: "IT Department", organization: Organization{
			name: "PT. ABC Mantap", alamat: "Jl. Kemana Kita 12",
		}},
		{nik: 222222, name: "Reza", age: 28, divisi: "IT Department", organization: Organization{
			name: "PT. ABC Mantap", alamat: "Jl. Kemana Kita 12",
		}},
		{nik: 333333, name: "Zulfi", age: 30, divisi: "IT Department", organization: Organization{
			name: "PT. ABC Mantap", alamat: "Jl. Kemana Kita 12",
		}},
	}

	var palingTua Employee
	maxAge := 0
	for _, pekerjas := range pekerja {
		if pekerjas.age > maxAge {
			maxAge = pekerjas.age
			palingTua = pekerjas
		}
	}
	fmt.Printf("%+v", palingTua)
}
