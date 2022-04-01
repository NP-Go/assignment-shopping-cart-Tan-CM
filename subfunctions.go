package main

import "fmt"

// This file has the sub functions
// View Shopping List
func viewShoppingList(shopListMap map[string]shopItem, categoryMap map[int]string) {
	// show all shopping items category, item , qty and unit price
	// Get all key of a map of struct
	keys := make([]string, 0, len(shopListMap))
	for k, _ := range shopListMap {
		keys = append(keys, k)
	}
	//	fmt.Println(keys)

	for _, shopItem := range keys {
		fmt.Printf("Category: %s - Item: %s Quantity: %d Unit Cost: %.0f\n",
			shopItem, categoryMap[shopListMap[shopItem].category], shopListMap[shopItem].qty, shopListMap[shopItem].cost)
	}

}
