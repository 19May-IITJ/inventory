package inventory

import (
	"bufio"
	"fmt"
	"inventory/modules/id"
	"inventory/modules/item"
	"os"
)

// Inventory represents the inventory system.
type Inventory struct {
	ItemsByType map[string][]item.Item
}

// AddItem adds a new item to the inventory.
func (inv *Inventory) AddItem(name string, quantity, itemType int) {

	newItem := item.Item{
		ID:       id.GenerateUniqueID(),
		Name:     name,
		Quantity: quantity,
		Type:     item.GetItemType(itemType),
	}
	inv.ItemsByType[newItem.Type.TypeKind] = append(inv.ItemsByType[newItem.Type.TypeKind], newItem)
}

// UpdateQuantity updates the quantity of an existing item.
func (inv *Inventory) UpdateQuantity(itemtype int) error {
	inv.ListItems(itemtype)
	var (
		newQuantity int
		idOrName    string
	)
	fmt.Println("Enter the Item Id or name")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	idOrName = scanner.Text()

	fmt.Println("Enter the new quantity")
	fmt.Scanln(&newQuantity)

	t := item.GetItemType(itemtype).TypeKind

	items := inv.ItemsByType[t]

	if len(items) > 0 {
		for i, item := range items {
			if item.ID == idOrName || item.Name == idOrName {
				inv.ItemsByType[t][i].Quantity = newQuantity
				fmt.Printf("ID: %s, Name: %s, Quantity: %d\n", item.ID, item.Name, inv.ItemsByType[t][i].Quantity)
				return nil
			}

		}
		return fmt.Errorf("no item found in inventory for the id or Name")
	} else {
		return fmt.Errorf("no item found in inventory for the type")
	}
	return fmt.Errorf("item with ID %s not found", idOrName)
}

// ListItems lists all items in the inventory.
func (inv *Inventory) ListItems(types int) {
	fmt.Println("Inventory:")
	if types != 7 {
		t := item.GetItemType(types).TypeKind
		fmt.Printf("Type: %s\n", t)
		items := inv.ItemsByType[t]
		if len(items) > 0 {
			for _, item := range items {
				fmt.Printf("ID: %s, Name: %s, Quantity: %d\n", item.ID, item.Name, item.Quantity)

			}
		} else {
			fmt.Printf("No stock for the selected Item Category\n")
		}
	} else {
		for itemType, items := range inv.ItemsByType {
			fmt.Printf("Type: %s\n", itemType)
			for _, item := range items {
				fmt.Printf("ID: %s, Name: %s, Quantity: %d\n\n", item.ID, item.Name, item.Quantity)
			}
		}
	}
}
