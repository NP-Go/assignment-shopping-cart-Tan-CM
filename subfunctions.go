package main

import (
	"fmt"
	"strconv"
	"strings"
)

// This file has the sub functions
// View Shopping List
func viewShoppingList(shopListMap map[string]shopItem, categorySlice []string) {
	// show all shopping items category, item , qty and unit price
	// Get all the keys of a map of struct
	keys := make([]string, 0, len(shopListMap))
	for k, _ := range shopListMap {
		keys = append(keys, k)
	}
	//	fmt.Println(keys)

	for _, nameItem := range keys {
		fmt.Printf("Category: %s - Item: %s Quantity: %d Unit Cost: %.0f\n",
			nameItem, categorySlice[shopListMap[nameItem].category], shopListMap[nameItem].qty, shopListMap[nameItem].cost)
	}

}

func genShoppingReport(shopListMap map[string]shopItem, categorySlice []string) {
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
			genReportByCategory(shopListMap, categorySlice)
		} else {
			// list items by cost
			genReportByCost(shopListMap, categorySlice)
			fmt.Println("")
		}
	} else {
		// do nothing for 3 and above
		fmt.Println("This is for 3 - do nothing")
	}
}

// Total cost of each category
func genReportByCategory(shopListMap map[string]shopItem, categorySlice []string) {

	// create slice based on the size of the category map
	categorySumTotal := make([]float64, len(categorySlice))

	// add cost of respective category
	for _, item := range shopListMap {
		categorySumTotal[item.category] += item.cost
		//		fmt.Println(item)
	}

	fmt.Println("Total cost by Category")
	for i, _ := range categorySumTotal {
		if i > 0 && i < len(categorySlice) {
			fmt.Printf("%s cost : %.0f\n", categorySlice[i], categorySumTotal[i])
		}
	}
}

// list items by cost
func genReportByCost(shopListMap map[string]shopItem, categorySlice []string) {

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
		tempStr := "Category: " + categorySlice[item.category] + " - Item: " + name + " Quantity: " +
			strconv.Itoa(item.qty) + " Unit Cost: " + strconv.Itoa(int(item.cost)) + "\n"

		listByCategory[item.category] = listByCategory[item.category] + tempStr
	}

	fmt.Println("List by Category")
	for i, v := range listByCategory {
		if i > 0 && i < len(categorySlice) { // skip the first one which is undefined
			fmt.Println(v)
		}
	}
}

// Note map is passed by reference and not by value
func addItem(shopListMap map[string]shopItem, categorySlice []string) {
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
			ok, key := category_value_present(categorySlice, category)

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

func modifyItem(shopListMap map[string]shopItem, categorySlice []string) {
	var item1, item2 shopItem
	var name, nameNew string
	var category string

	fmt.Println("Modify Items")
	fmt.Println("Which item do you want to modify?")
	fmt.Scanln(&name)
	if (strings.TrimSpace(name)) != "" {

		item1 = shopListMap[name] //original item

		fmt.Printf("Current item name is %s - Category is %s - Quantity is %d - Unit Cost %.0f\n",
			name, categorySlice[item1.category], item1.qty, item1.cost)
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
				category = categorySlice[item1.category]
			}

			item2.category = sliceIndex(categorySlice, category) // set up with initial value first
			// check valid category
			if item2.category != -1 {

				// white spaces will not be captured for int and float so no added processing for white space entry
				fmt.Println("Enter New Quantity, Enter for no change?")
				fmt.Scanln(&item1.qty)
				fmt.Println("Enter New Unit Cost, Enter for no change?")
				fmt.Scanln(&item1.cost)
				// update item to shopList map
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
				fmt.Println("Invalid Category")
			}
		} else {
			fmt.Println("Invalid Name")
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

func printCurrentField(shopListMap map[string]shopItem) {
	// key all keys in the shopList Map in a slice
	keys := getAllKeys(shopListMap)
	//	fmt.Println(keys)

	fmt.Println("Print Current Data.")
	// Get data field of each slice and display it
	for _, k := range keys {
		fmt.Println(k, "-", shopListMap[k])
	}
}

// slice is passed by value which is the pointer to the array
// Need to pass the pointer to the slice itself if modification to the slice is needed
func addNewOrModifyCategory(categorySlice *[]string) {
	var newCategory string

	//	fmt.Println(len(categoryMap), categoryMap)
	fmt.Println("Add New or Modify Category Name")
	fmt.Println("What is the New Category Name to add?")
	fmt.Scanln(&newCategory)
	if (strings.TrimSpace(newCategory)) != "" {
		//		fmt.Printf("Length of Map = %d\n", len(categoryMap))

		i := sliceIndex(*categorySlice, newCategory)
		if i == -1 {
			//	New category added to last slice
			*categorySlice = append(*categorySlice, newCategory)
			fmt.Printf("New Category: %s added at index %d\n", newCategory, len(*categorySlice)-1)
		} else {
			if i > 0 {
				fmt.Printf("Category: %s already exist at index %d !\n", newCategory, i)
				fmt.Println("Enter New Name to modify the category name or Press return to ignore")
				// Get new name for existing category
				fmt.Scanln(&newCategory)
				if (strings.TrimSpace(newCategory)) != "" {
					// Note the () is needed to get the dereferencing
					(*categorySlice)[i] = newCategory
				} else {
					// Exit and nothing with whitespaces
					// no change to category name
				}

			} else {
				// Category for undefine category
				fmt.Print("Undefined Category Found")
			}
		}
	} else {
		fmt.Println("No Input Found!")
	}
	//	fmt.Println(*categorySlice)
}

// show category for debug use only
func showCategory(categorySlice []string) {

	fmt.Println("Show Category")
	fmt.Println(categorySlice)
	fmt.Println(len(categorySlice))
}

func deleteCategory(shopListMap map[string]shopItem, categorySlice *[]string) {
	var newCategory string
	// Note when category is deleted
	// Problem with this operation
	// 1.  How to handle the shopping List which has the category to be deleted?
	// 2.  The association of the shopping list will need to be reassociated because of deletion of a category
	// Solution:
	// 1.  Disallow delete if any shopping list has this category in use
	// 2.  Re-associate the affect shopping list for category affected by the deletion

	//	fmt.Println(len(categoryMap), categoryMap)
	fmt.Println("Delete Category Name")
	fmt.Println("What is the Category Name to delete?")
	fmt.Scanln(&newCategory)
	if (strings.TrimSpace(newCategory)) != "" {

		i := sliceIndex(*categorySlice, newCategory)
		if i == -1 {
			fmt.Printf("Category: %s cannot be found\n", newCategory)

		} else {
			if i > 0 {
				// Step 1.  Before delete is done, need to check if the category is used in shopping list
				// if it is, deletion is disallowed
				if map_category_present(shopListMap, i) == false {
					// This approach maintains the original order of the slice with one element removed
					// Note ... variadic argument because the it is variable length element to unpack
					*categorySlice = append((*categorySlice)[:i], (*categorySlice)[i+1:]...)
					fmt.Printf("Category: %s delete\n", newCategory)
					// Step 2.  If deletion is category i is removed, shoplist cat index has to be adjusted
					// All category > i is reduced by 1 for every item is shopping list
					shiftDownCategoryBy1(shopListMap, i)
				} else {
					fmt.Println("Category in used, cannot be deleted")
				}
			} else {
				// Category for undefine category
				fmt.Print("Undefined Category Found")
			}
		}
	} else {
		fmt.Println("No Input Found!")
	}
	fmt.Println(*categorySlice)
}

// return the index given the name of the category
// return -1 , if index is not found
func sliceIndex(s []string, strToFind string) int {
	for i, v := range s {
		// end when there is a match, otherwise continue to search till the end
		if v == strToFind {
			return i
		}
	}
	return -1 // No index found
}

// check if map has this key
// retrun true if there is
func map_key_present(m map[string]shopItem, k string) (ok bool) {
	_, ok = m[k]
	return
}

// check if map has this category in the object
// return true is category is present
func map_category_present(m map[string]shopItem, k int) (ok bool) {
	// add cost of respective category
	for _, item := range m {
		if (item.category) == k {
			ok = true
			break
		}
	}
	return
}

// operation on category
// For all shoplist item with category > categoryKey, category = category -1
// Note:  Map value cannot be changed directly, it can only be re-created
func shiftDownCategoryBy1(m map[string]shopItem, categoryKey int) {
	// add cost of respective category1

	for k, item := range m {
		if (item.category) > categoryKey {
			// delete affect entry of affected key
			delete(m, k)
			// modify the category of the item
			item.category = item.category - 1
			// create a new item with the same key
			m[k] = item
		}
	}
}

// check if category slice has this value
func category_value_present(s []string, v string) (ok bool, index int) {
	// test the values in the map
	for i, value := range s {
		// check validity of category name
		if value == v {
			ok = true
			index = i
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
