package mwsproducts

import (
	"errors"
)

// isValidProductIDTypes validates if given product id type is valid
func isValidProductIDTypes(idType string) error {
	switch idType {
	case
		PIDTypesASIN,
		PIDTypesGCID,
		PIDTypesSKU,
		PIDTypesUPC,
		PIDTypesEAN,
		PIDTypesISBN,
		PIDTypesJAN:
		return nil
	default:
		return errors.New("invalid product id type")
	}
}

// isValidList validates if given product id list is valid
	func isValidList(idList []string, maxLen int) error {
	if len(idList) > maxLen {
		return errors.New("excess maximum id length")
	}

	return nil
}

// isValidItemContidions validates if given item condition is valid
func isValidItemContidions(itemCondition string) error {
	switch itemCondition {
	case
		ICClub,
		ICCollectible,
		ICNew,
		ICRefurbished,
		ICUsed:
		return nil
	default:
		return errors.New("invalid item condition")
	}
}
