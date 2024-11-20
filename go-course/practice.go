package main

import "fmt"

func main() {
	hobbies := [3]string{"dance", "anime", "game"}
	fmt.Println(hobbies)

	fmt.Println(hobbies[0])

	fmt.Println(hobbies[1:3])

	fmt.Println(hobbies[0:2])
	hobbies3 := hobbies[:2]
	fmt.Println(hobbies3)

	hobbies4 := hobbies3[:3]
	fmt.Println(hobbies4)

	goalhobbies := append(hobbies4, "gomaster", "cool")
	fmt.Println(goalhobbies)

	goalhobbies2 := append(goalhobbies, "gopro")
	fmt.Println(goalhobbies2)

	type Product struct {
		title string
		id    int
		price float64
	}

	listProduct := []Product{{"ice", 2, 43}, {"ice", 2, 43}}
	fmt.Println(listProduct)
	netProduct := Product{
		"ice",
		2,
		43,
	}
	listProduct = append(listProduct, netProduct)
	fmt.Println(listProduct)
}
