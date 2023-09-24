package main

import (
	"bufio"
	"fmt"
	"inventory/modules/inventory"
	"inventory/modules/item"
	"inventory/modules/utility"
	"os"
)

func main() {
	inventory := inventory.Inventory{ItemsByType: make(map[string][]item.Item)}
	for {
		fmt.Println("\n****Inventory Management System****")
		fmt.Println("1. Add Item")
		fmt.Println("2. Update Quantity")
		fmt.Println("3. List Items")
		fmt.Println("4. Exit")
		fmt.Print("Select an option: ")

		var choice int
		_, err := fmt.Scanf("%d", &choice)

		if err != nil {
			fmt.Println("Invalid input. Please enter a valid option.")
			continue
		}

		switch choice {
		case 1:
			var name string
			var quantity, itemType int
			fmt.Print("Enter the item name: ")

			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			name = scanner.Text()

			fmt.Println(name)
			fmt.Print("Enter the item type from below category: ")
			utility.DisplayItemCategory()
			fmt.Scanln(&itemType)
			fmt.Print("Enter the quantity: ")
			fmt.Scanf("%d", &quantity)
			inventory.AddItem(name, quantity, itemType)
			fmt.Println("Item added to the inventory.")
		case 2:
			var (
				itemType int
			)

			fmt.Print("Enter the item type to update from below category: ")
			utility.DisplayItemCategory()
			fmt.Scanln(&itemType)

			err := inventory.UpdateQuantity(itemType)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Quantity updated successfully.")
			}

		case 3:
			fmt.Println("\nSelect Item Type ")
			utility.DisplayItemCategory()
			fmt.Println("7 All")
			fmt.Print("Select an option: ")
			var itemType int
			fmt.Scanln(&itemType)
			inventory.ListItems(itemType)
		case 4:
			os.Exit(0)
		default:
			fmt.Println("Invalid option. Please choose a valid option.")

		}
	}

}
