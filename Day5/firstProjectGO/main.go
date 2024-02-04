package main

import "fmt"

func main() {
	var arr3 [4][5]string = [4][5]string{
		{"alvin", "arif", "reza", "rinaldi", "nina"},      //0
		{"noel", "dilla", "rosa", "juan michel", "teguh"}, //1
		{"septyan", "farras", "giyanda", "afin", "azwar"}, //2
		{"dionysius", "rifki", "raffi", "zain"},           //3
	}
	for i, x := range arr3 {
		for j, valueX := range x {
			fmt.Println(i, x, j, valueX)
		}
	}
	fmt.Println("====================================================")
	for i, x := range arr3 {
		for j, valueX := range x {
			if valueX == "reza" {
				fmt.Println(i, j, x, valueX)
			}
		}
	}
}

// monyet1 := "cemberut"
// monyet2 := "senyum"

// if strings.ToLower(monyet1) != strings.ToLower(monyet2) {
// 	fmt.Println("Tidak aman")
// } else {
// 	fmt.Println("aman")
// }

// waktu := 2
// berkokok := false

// if waktu < 24 {
// 	if waktu >= 1 && waktu == 3 && berkokok {
// 		fmt.Println("kesambet")
// 	} else if waktu > 3 && waktu < 7 && !berkokok {
// 		fmt.Println("aman")
// 	} else {
// 		fmt.Println("jam invalid")
// 	}
// }

// waktu := 1
// berkokok := true

// if waktu < 24 {
// 	if waktu >= 1 && waktu == 3 && berkokok {
// 		fmt.Println("kesambet")
// 	} else if waktu > 3 && waktu < 7 && berkokok {
// 		fmt.Println("aman")
// 	} else {
// 		fmt.Println("jam invalid")
// 	}
// }

// func isBerkokok(waktu int, berkokok bool) bool {
// 	if waktu < 24 {
// 		if waktu > 12 && waktu == 3 && berkokok {
// 			fmt.Println("kesambet")
// 		} else if waktu > 3 && waktu < 7 && !berkokok {
// 			fmt.Println("aman")
// 		} else {
// 			fmt.Println("aman")
// 		}
// 	}
// }
