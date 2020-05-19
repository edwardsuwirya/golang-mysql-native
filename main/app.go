package main

import "fmt"

func main() {
	productWithCategory, err := getAllProducts()
	if err != nil {
		fmt.Println(err)
	}
	for _, p := range productWithCategory {
		fmt.Println(p)
	}
}
