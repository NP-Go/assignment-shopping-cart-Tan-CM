package main

import (
	"fmt"
)

// define structure for item
type shopItem struct {
	category int
	qty      int
	cost     float64
}

func main() {
	//insert code here
	// create map reference with dynamic size allocation
	categoryMap := map[int]string{
		// initialise map
		0: "Undefined",
		1: "Household",
		2: "Food",
		3: "Drinks",
	}

	// create map reference
	shopListMap := make(map[string]shopItem)
	// initialise map
	shopListMap["Cup"] = shopItem{1, 5, 3}
	shopListMap["Cake"] = shopItem{2, 3, 1}
	shopListMap["Sprite"] = shopItem{3, 5, 2}
	shopListMap["Fork"] = shopItem{1, 4, 3}
	shopListMap["Bread"] = shopItem{2, 2, 2}
	shopListMap["Plate"] = shopItem{1, 4, 3}
	shopListMap["Coke"] = shopItem{3, 5, 2}

	// can create pointer to shopList instead of using &shopList
	//	shopListPtr := new(shopList)

	shoppingListApp := []string{
		"Shopping List Applications",
		"==========================",
		"1.  View entire shopping List.",
		"2.  Generate Shopping List Report.",
		"3.  Add Items.",
		"4.  Modify Items.",
		"5.  Delete Item.",
		"6.  Print Current Data",
		"7.  Add New Category Name.",
		"8.  For Testing Only",
		"0.  To Exit",
		"Select your choice:",
	}

	for true {
		for _, item := range shoppingListApp {
			fmt.Println(item)
		}

		var choice int
		fmt.Scanln(&choice)
		//	fmt.Println(strconv.Itoa(choice))

		switch choice {
		case 1:
			viewShoppingList(shopListMap, categoryMap)
		case 2:
			genShoppingReport(shopListMap, categoryMap)
		case 3:
			addItem(shopListMap, categoryMap)
		case 4:
			modifyItem(shopListMap, categoryMap)
		case 5:
			deleteItem(shopListMap)
		case 6:
			printCurrentField(shopListMap, categoryMap)
		case 7:
			addNewCategory(categoryMap)
		case 8:
			fmt.Println("Category Map")
			fmt.Println(categoryMap)
			fmt.Println(len(categoryMap))
		case 0:
			// break only work for switch but not applied to for loop
		default:
			fmt.Println("Wrong Entry")
		}
		if choice == 0 {
			break // use break here to break the loop
		}
	}
}
