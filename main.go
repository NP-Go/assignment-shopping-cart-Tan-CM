package main

import (
	"fmt"
)

// *****************************************
// 2 Apr 2022
// Tan Chor Ming
// Record of Time on this stupid Go Basic assignment
// Wed, 5
// Thurs 12
// Fri 8
// Sat 3
// Total 28 hours
// ****************************************
// Key Lessons Learnt about Go lang exercise
// All functions are passed by value.  It is the content of the value that determine what is the content
//	1.  Slice are passed to functions by value (which is the pointer of the array)  Slice is made up of array
//      This means that if the slice is to be changed like modify the element in the slice, the pointer to the slice is needed
//      for change to the actual category to happen.  See addNewOrModifyCategory(&categorySlice) and deleteCategory(shopListMap, &categorySlice)
//  2.  Map is passed to function by value (which is the pointer to the key)
//      if the key is new, a new key-value pair can be created with m[k] = item
//		However, the key-value pair , if the key to the value already exist, it has first to be delete before a new key-value is recreated
//      delete(m, k)
//		m[k] = item
//      Note that Map value cannot be changed, so if the same key is needed for a different value, it has to be re-created after deletion
//	3.  ... means variadic or multiple arguments, which can only be the last parameter of a function if there are several parameters
//		eg append((*categorySlice)[:i], (*categorySlice)[i+1:]...)  where the ... means it is a slice (variable length array)
//  4.  for k, v := range map type    ==> k = key, v = value for map
//          i, e := range slice type  ==> i = index, e = element for both slice or array
//  5.  append  -- overloaded built-in function, use to build slice of elements of variable size
//      for slice --> slices = append (slice,e)
//
//*****************************************
// What can be improved for this course?
// Development Consideration should be in page 1 instead of the last page
// I wasted so much time, because this consideration is so important for novice like me.
//*****************************************

// define structure for item
type shopItem struct {
	category int
	qty      int
	cost     float64
}

func main() {
	//insert code here
	// create map reference with dynamic size allocation

	categorySlice := []string{
		// initialise map
		"Undefined",
		"Household",
		"Food",
		"Drinks",
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
		"7.  Add New or Modify Category Name",
		"8.  View Category",
		"9.  Delete Category",
		"x.  To Exit",
		"Select your choice:",
	}

	for true {
		for _, item := range shoppingListApp {
			fmt.Println(item)
		}

		var choice string
		fmt.Scanln(&choice)
		//	fmt.Println(strconv.Itoa(choice))

		switch choice {
		case "1":
			viewShoppingList(shopListMap, categorySlice)
		case "2":
			genShoppingReport(shopListMap, categorySlice)
		case "3":
			addItem(shopListMap, categorySlice)
		case "4":
			modifyItem(shopListMap, categorySlice)
		case "5":
			deleteItem(shopListMap)
		case "6":
			printCurrentField(shopListMap)
		case "7":
			addNewOrModifyCategory(&categorySlice)
		case "8":
			showCategory(categorySlice)
		case "9":
			deleteCategory(shopListMap, &categorySlice)
		case "x":
			// break only work for switch but not applied to for loop
		default:
			fmt.Println("Wrong Entry")
		}
		if choice == "x" {
			break // use break here to break the loop
		}
	}
}
