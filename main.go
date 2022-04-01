package main

import (
	"fmt"
	"strconv"
	"strings"
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
		"8.  For Test",
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
		default:
			fmt.Println("Wrong Entry")
		}
	}
}

func genShoppingReport(shopListMap map[string]shopItem, categoryMap map[int]string) {
	generateReport := []string{
		"Generate Report",
		"1.  Total Cost of each category.",
		"2.  List of item by category.",
		"3.  Main menu.",
		"",
		"Choose your report:",
	}
	var choice int

	// Generate Shopping List Report
	for _, item := range generateReport {
		fmt.Println(item)
	}
	fmt.Scanln(&choice)

	if choice <= 2 {
		if choice == 1 {
			// Total cost of each category
			genReportByCategory(shopListMap, categoryMap)
		} else {
			// list items by cost
			genReportByCost(shopListMap, categoryMap)
			fmt.Println("")
		}
	} else {
		// do nothing for 3 and above
		fmt.Println("This is for 3 - do nothing")
	}
}

// Total cost of each category
func genReportByCategory(shopListMap map[string]shopItem, categoryMap map[int]string) {

	// create slice based on the size of the category map
	categorySumTotal := make([]float64, len(categoryMap))

	// add cost of respective category
	for _, item := range shopListMap {
		categorySumTotal[item.category] += item.cost
		//		fmt.Println(item)
	}

	fmt.Println("Total cost by Category")
	for i, _ := range categorySumTotal {
		if i > 0 && i < len(categoryMap) {
			fmt.Printf("%s cost : %.0f\n", categoryMap[i], categorySumTotal[i])
		}
	}
}

// list items by cost
func genReportByCost(shopListMap map[string]shopItem, categoryMap map[int]string) {

	// Show all shopping items category, item , qty and unit price
	// Get all the keys (Item name) of a map of struct into a slice
	keys := make([]string, 0, len(shopListMap))
	for k := range shopListMap {
		keys = append(keys, k)
	}

	//	fmt.Println(keys)
	listByCategory := make([]string, 10, 15)

	for _, name := range keys {
		// get particular item
		item := shopListMap[name]
		tempStr := "Category: " + categoryMap[item.category] + " - Item: " + name + " Quantity: " +
			strconv.Itoa(item.qty) + " Unit Cost: " + strconv.Itoa(int(item.cost)) + "\n"

		listByCategory[item.category] = listByCategory[item.category] + tempStr
	}

	fmt.Println("List by Category")
	for i, v := range listByCategory {
		if i > 0 && i < len(categoryMap) { // skip the first one which is undefined
			fmt.Println(v)
		}
	}
}

// Note map is passed by reference and not by value
func addItem(shopListMap map[string]shopItem, categoryMap map[int]string) {
	var item shopItem
	var name string
	var category string
	fmt.Println("What is the name of your item?")
	fmt.Scanln(&name)
	// check to ensure no white spaces for name
	if (strings.TrimSpace(name)) != "" {
		fmt.Println("What category does it belong to?")
		fmt.Scanln(&category)
		if (strings.TrimSpace(category)) != "" {
			// read status of category against the map
			ok, key := category_value_present(categoryMap, category)

			if ok {
				item.category = key
				fmt.Println("How many units are there?")
				fmt.Scanln(&item.qty)
				fmt.Println("How much does it cost per unit?")
				fmt.Scanln(&item.cost)
				// add a new key to map
				shopListMap[name] = item
			} else {
				fmt.Println("Wrong Category Error !")
			}
		} else {
			fmt.Println("Blank Category Entry !")
		}
	} else {
		fmt.Println("Blank Name Entry !")
	}
}

func modifyItem(shopListMap map[string]shopItem, categoryMap map[int]string) {
	var item1, item2 shopItem
	//	var shop_list []shopItem
	var name, nameNew string
	//	var x int
	var category string

	fmt.Println("Modify Items")
	fmt.Println("Which item do you want to modify?")
	fmt.Scanln(&name)
	if (strings.TrimSpace(name)) != "" {

		item1 = shopListMap[name] //original item

		fmt.Printf("Current item name is %s - Category is %s - Quantity is %d - Unit Cost %.0f\n",
			name, categoryMap[item1.category], item1.qty, item1.cost)
		// check that only valid name is allowed
		if x, found := shopListMap[name]; found {
			item1 = x

			// make a copy before modification
			item2 = item1

			//			fmt.Println(item2)
			//			fmt.Println("Pause")
			//			fmt.Scanln()

			fmt.Println("Enter New Name, Enter for no change?")
			fmt.Scanln(&nameNew)
			// Restore to previous name for white spaces entry
			if (strings.TrimSpace(nameNew)) == "" {
				nameNew = name
			}
			fmt.Println("Enter New Category, Enter for no change?")
			fmt.Scanln(&category)
			//restore to old value if white spaces in entry
			if (strings.TrimSpace(category)) == "" {
				category = categoryMap[item1.category]
			}
			item2.category = reverseMap(categoryMap)[category] // set up with initial value first

			// scan all keys of the category map
			//		fmt.Printf("Category # = %d", item2.category)
			//		fmt.Scanln()

			// > 0 means valid category
			if item1.category > 0 {

				// white spaces will not be captured for int and float so no added processing for white space entry
				fmt.Println("Enter New Quantity, Enter for no change?")
				fmt.Scanln(&item2.qty)
				fmt.Println("Enter New Unit Cost, Enter for no change?")
				fmt.Scanln(&item1.cost)
				// update item to shopList slice
				if name != nameNew {
					delete(shopListMap, name)    // remove old key
					shopListMap[nameNew] = item1 // create new key
				} else {
					shopListMap[name] = item1 // update old key with new values
				}

				// Update change status to each field
				if item1.category == item2.category {
					fmt.Println("No changes to category made")
				}
				if item1.qty == item2.qty {
					fmt.Println("No changes to quantity made")
				}
				if item1.cost == item2.cost {
					fmt.Println("No changes to unit cost made")
				}
				if name == nameNew {
					fmt.Println("No changes to name made")
				} else { // Item does not exist
					fmt.Println("Name Updated")
				}
			} else {
				fmt.Println("Name does not exist")
			}
		} else {
			fmt.Println("Unknown Category")
		}
	} else {
		fmt.Println("Blank Name Entry !")
	}
}

func deleteItem(shopListMap map[string]shopItem) {

	//	var shop_list shopItem
	var name string
	var x int
	fmt.Println("Delete Item")
	fmt.Println("Enter item name to delete")
	fmt.Scanln(&name)
	// check for blank name
	if (strings.TrimSpace(name)) != "" {
		// check for valid existing name
		if map_key_present(shopListMap, name) {
			//			fmt.Println("Key is present")
			//			fmt.Printf(" found= %d\n", x)
			//			fmt.Scanln()

			// Remove Map x
			if x < len(shopListMap) {
				// delete item in slice with append of particular slice
				delete(shopListMap, name) // remove old key name
				fmt.Println("Deleted " + name)
			} else { // Item does not exist
				fmt.Println("Item not found.  Nothing to delete")
			}
		} else {
			fmt.Println("Key is not present")
		}
	} else {
		fmt.Println("Blank Name Entry !")
	}
}

func printCurrentField(shopListMap map[string]shopItem, categoryMap map[int]string) {
	// key all keys in the shopList Map in a slice
	keys := getAllKeys(shopListMap)
	fmt.Println(keys)

	// Get data field of each slice and display it
	for _, k := range keys {
		fmt.Println(k, "-", shopListMap[k])
	}

}

// Note slice is passed by reference and not by value
func addNewCategory(categoryMap map[int]string) {
	var newCategory string

	//	fmt.Println(len(categoryMap), categoryMap)
	fmt.Println("Add New Category Name")
	fmt.Println("What is the New Category Name to add?")
	fmt.Scanln(&newCategory)
	if (strings.TrimSpace(newCategory)) != "" {
		//		fmt.Printf("Length of Map = %d\n", len(categoryMap))

		for i := 1; i < len(categoryMap)+1; i++ {
			if i == len(categoryMap) {
				fmt.Printf("New Category: %s added at index %d\n", newCategory, i)
				categoryMap[i] = newCategory
				break
			} else {
				if newCategory == categoryMap[i] {
					fmt.Printf("Category: %s already exist at index %d !\n", newCategory, i)
					// break for loop if found
					break
				} else {
					// Category Not found, continue to scan category Map
				}
			}
		}

		// show updated category Map
		//fmt.Println(len(categoryMap), categoryMap)
	} else {
		fmt.Println("No Input Found!")
	}

}

// create a reverse map for category
func reverseMap(m map[int]string) map[string]int {
	n := make(map[string]int, len(m))
	for k, v := range m {
		n[v] = k
	}
	return n
}

// check if map has this key
func map_key_present(m map[string]shopItem, k string) (ok bool) {
	_, ok = m[k]
	return
}

// check if category map has this value
func category_value_present(m map[int]string, v string) (ok bool, index int) {
	// test the values in the map
	for key, value := range m {
		// check validity of category name
		if value == v {
			ok = true
			index = key
			break
		} else {
			ok = false
			index = -1 // invalid key
		}
	}
	return
}

// get all keys of the shopItemMap and return a slice
func getAllKeys(m map[string]shopItem) (s []string) {
	// test the values in the map
	//s := make([]string, len(m))

	for key, _ := range m {
		s = append(s, key)
	}
	return
}
