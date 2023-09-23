package main

import (
	"fmt"
	"inventory/modules/inventory"
	"os"
)

func main() {
	inventory := inventory.Inventory{}
	for {
		fmt.Println("\nInventory Management System")
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
			var quantity int

			fmt.Print("Enter the item name: ")
			fmt.Scanln(&name)

			fmt.Print("Enter the quantity: ")
			fmt.Scanf("%d", &quantity)

			inventory.AddItem(name, quantity)
			fmt.Println("Item added to the inventory.")
		case 2:
			var id, newQuantity int

			fmt.Print("Enter the item ID to update: ")
			fmt.Scanf("%d", &id)

			fmt.Print("Enter the new quantity: ")
			fmt.Scanf("%d", &newQuantity)

			err := inventory.UpdateQuantity(id, newQuantity)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Quantity updated successfully.")
			}
		case 3:
			inventory.ListItems()
		case 4:
			os.Exit(0)
		default:
			fmt.Println("Invalid option. Please choose a valid option.")
		}
	}
}
