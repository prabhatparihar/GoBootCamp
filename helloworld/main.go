package main

import (
	"fmt"
	"github.com/fatih/color"
)

func main() {
	//emoji := "😀 😃 😄 😁 😆 😅. 😂. 🤣. 🥲. 🥹. ☺️"
	////var str = "line \n line2 "
	//rawstr := `line1 \n
	//line2`
	////color.Red("Hello, World!", str)
	//fmt.Println(rawstr)
	//fmt.Println(emoji)

	//var (
	//	value    string
	//	stars    int
	//	comment  []string
	//	userName string
	//)
	//id := make(map[string]int)
	//id["prabhat"] = 2
	//id["prabhat1"] = 3
	//id["prabhat3"] = 5
	//fmt.Println("enter the stars  ")
	//fmt.Scanln(&stars)
	//fmt.Println("enter the comment  ")
	//comment = append(comment, "this is bad,")
	//comment = append(comment, "this is god,")
	//comment = append(comment, "this is okay")
	//fmt.Println("enter the username  ")
	//fmt.Scanln(&userName)
	//
	//fmt.Printf("thank you for the rating  the submitted rating is: %v \n the first comments is:  %v \n the second comments is:  %v \n the third comments is:  %v ", stars, comment[0], comment[1], comment[2])
	//fmt.Println("Enter the product Id to get the rating")
	//fmt.Scanln(&value)
	//_, exists := id[value]
	//if exists {
	//	fmt.Println("found the rating is ", id[value])
	//}

	type rating struct {
		id       string
		comment  string
		stars    int
		username string
	}

	var r rating

	fmt.Println("Enter the product id")
	fmt.Scanln(&r.id)
	fmt.Println("Enter the rating ")
	fmt.Scanln(&r.stars)
	fmt.Println("Enter the comment ")
	fmt.Scanln(&r.comment)
	fmt.Println("Enter the username ")
	fmt.Scanln(&r.username)
	if r.stars > 3 {
		color.Green("Thanks for the review ")
	} else {
		color.Red("will look into the feedback  ")
	}

	switch r.stars {
	case 1:
		fmt.Println("⭐")
	case 2:
		fmt.Println("⭐⭐")
	case 3:
		fmt.Println("⭐⭐⭐")
	case 4:
		fmt.Println("⭐⭐⭐⭐")
	case 5:
		fmt.Println("⭐⭐⭐⭐⭐")
	}

	fmt.Printf("The rating for product id is: %v \n  rating is :  %v \n  comment is %v", r.id, r.stars, r.comment)

}
