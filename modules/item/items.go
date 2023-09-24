package item

import "inventory/modules/utility"

// Item represents a product in the inventory.
type Item struct {
	ID         string
	Name       string
	Quantity   int
	Attributes string
	Type       ItemType
}

type ItemType struct {
	TypeId   int
	TypeKind string
}

var (
	Fruit         ItemType
	Miscellaneous ItemType
	Electronics   ItemType
	Garments      ItemType
	HouseHold     ItemType
	Grocery       ItemType
	Cosmetics     ItemType
)

func NewItemType(kind string) ItemType {
	return ItemType{
		TypeKind: kind,
	}
}

func GetItemType(kind int) ItemType {
	switch kind {
	case 1:
		return Electronics
	case 2:
		return Cosmetics
	case 3:
		return HouseHold
	case 4:
		return Grocery
	case 5:
		return Garments
	case 6:
		return Fruit
	default:
		return Miscellaneous
	}
}

func init() {
	Fruit.TypeId = utility.FRUIT
	Fruit.TypeKind = utility.FRUIT_

	Miscellaneous.TypeId = utility.MISCELLANEOUS
	Miscellaneous.TypeKind = utility.MISCELLANEOUS_

	Electronics.TypeId = utility.ELECTRONICS
	Electronics.TypeKind = utility.ELECTRONICS_

	Grocery.TypeId = utility.GROCERY
	Grocery.TypeKind = utility.GROCERY_

	HouseHold.TypeId = utility.HOUSEHOLD
	HouseHold.TypeKind = utility.HOUSEHOLD_

	Garments.TypeId = utility.GARMENTS
	Garments.TypeKind = utility.GARMENTS_

	Cosmetics.TypeId = utility.COSMETICS
	Cosmetics.TypeKind = utility.COSMETICS_

}
