package main

import (
	"fmt"
	"os"
	"strings"
)

func totalIncome(sale []int) {
	sum := 0

	for i := 0; i < len(sale); i++ {
		sum = sale[i] + sum
	}
	fmt.Println("Total Income for the day is : ", sum)
	fmt.Println("    THANK YOU     ")
	os.Exit(0)
}
func Billrs(quantity int, order string) int {
	A := map[string]int{
		"C": 40,
		"D": 80,
		"T": 20,
		"I": 20,
		"V": 25,
		"B": 30,
		"P": 30,
		"M": 90,
		"H": 70,
		"F": 30,
		"J": 30,
		"L": 15,
		"S": 20,
	}
	billt := quantity * A[order]
	return billt
}

func main() {
	var sale = []int{}
	var order string
	var quan int
	count := 0
	for true {
		fmt.Println("      Welcome TO Our Resturant      ")
		fmt.Println("****************MENU LIST***************")
		fmt.Print("C->Coffee, Price->40rs\n D->Dosa, Price->80rs \n T->Tomato Soup, Price->20rs\n I->Idli ,Price->20rs\n V->Vada, Price->25rs.\n B->Chhole+Bature, Price->30rs\n P->Paneer Pakoda, Price->30rs\n M->Manchurian, Price->90rs\n H->Hakka Noodles, Price->70rs.\n F->French Fries, Price->30rs\n J->Jalebi, Price->30rs\n L->Lemonade, Price->15rs\n S->Spring Roll, Price->20rs\n")
		fmt.Println("****************************************")
		fmt.Println("****************************************")
		fmt.Println("Press END For closing the day.")
		fmt.Println("Place the order: ")
		fmt.Scan(&order)
		order = strings.ToUpper(order)

		if order == "END" {
			totalIncome(sale)
		} else {
			fmt.Scan(&quan)
		}
		bill := Billrs(quan, order)
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Println("Your Total bill: ", bill)
		fmt.Println("~~~~~~~~THANK YOU~~~~~~~~")
		sale = append(sale, bill)
		count++

	}
}
