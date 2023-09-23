package inventory

import (
	"fmt"
	"inventory/modules/item"
)

// Inventory represents the inventory system.
type Inventory struct {
	Items []item.Item
}

// AddItem adds a new item to the inventory.
func (inv *Inventory) AddItem(name string, quantity int) {
	newItem := item.Item{
		ID:       len(inv.Items) + 1,
		Name:     name,
		Quantity: quantity,
	}
	inv.Items = append(inv.Items, newItem)
}

// UpdateQuantity updates the quantity of an existing item.
func (inv *Inventory) UpdateQuantity(id int, quantity int) error {
	for i, item := range inv.Items {
		if item.ID == id {
			inv.Items[i].Quantity = quantity
			return nil
		}
	}
	return fmt.Errorf("item with ID %d not found", id)
}

// ListItems lists all items in the inventory.
func (inv *Inventory) ListItems() {
	fmt.Println("Inventory:")
	for _, item := range inv.Items {
		fmt.Printf("ID: %d, Name: %s, Quantity: %d\n", item.ID, item.Name, item.Quantity)
	}
}
