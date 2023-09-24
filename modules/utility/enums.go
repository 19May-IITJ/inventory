package utility

import "fmt"

const (
	MISCELLANEOUS = iota
	ELECTRONICS
	COSMETICS
	HOUSEHOLD
	GROCERY
	GARMENTS
	FRUIT
)
const (
	MISCELLANEOUS_ = "MISCELLANEOUS"
	ELECTRONICS_   = "ELECTRONICS"
	COSMETICS_     = "COSMETICS"
	HOUSEHOLD_     = "HOUSEHOLD"
	GROCERY_       = "GROCERY"
	GARMENTS_      = "GARMENTS"
	FRUIT_         = "FRUIT"
)

func DisplayItemCategory() {
	fmt.Println()
	fmt.Println(MISCELLANEOUS, "MISCELLANEOUS")
	fmt.Println(ELECTRONICS, "ELECTRONICS")
	fmt.Println(COSMETICS, "COSMETICS")
	fmt.Println(HOUSEHOLD, "HOUSEHOLD")
	fmt.Println(GROCERY, "GROCERY")
	fmt.Println(GARMENTS, "GARMENTS")
	fmt.Println(FRUIT, "FRUIT")
}



