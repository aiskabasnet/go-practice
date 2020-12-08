package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	// initializers()

	db, err := gorm.Open(sqlite.Open("practice.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&Product{})

}

// func initializers() {
// 	//variables
// 	i := 10
// 	var j int = 11
// 	var k int
// 	k = 12

// 	//constants
// 	const a int = 14
// 	const b bool = true

// 	//arrays
// 	grades := [3]int{6, 7, 10}
// 	grades1 := [...]int{2, 3}

// 	//slice
// 	marks := []int{1, 2, 3, 4, 5, 6, 7, 8}
// 	marks1 := marks[:4]

// 	fmt.Println(i)
// 	fmt.Println(j)
// 	fmt.Println(k)
// 	fmt.Println(a)
// 	fmt.Println(b)

// 	grades1[0] = 5
// 	fmt.Println(grades)
// 	fmt.Println(grades1)
// 	fmt.Println(len(marks))
// 	fmt.Println(marks1)
// }
