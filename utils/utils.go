package utils

import (
	"errors"

	"github.com/bxmon/mws-products-client/consts"
)

// IsValidProductIDTypes validates if given product id type is valid
func IsValidProductIDTypes(idType string) error {
	switch idType {
	case
		consts.PIDTypesASIN,
		consts.PIDTypesGCID,
		consts.PIDTypesSKU,
		consts.PIDTypesUPC,
		consts.PIDTypesEAN,
		consts.PIDTypesISBN,
		consts.PIDTypesJAN:
		return nil
	default:
		return errors.New("invalid product id type")
	}
}

// IsValidList validates if given product id list is valid
func IsValidList(idList []string, maxLen int) error {
	if len(idList) > maxLen {
		return errors.New("excess maximum id length")
	}

	return nil
}

// IsValidItemContidions validates if given item condition is valid
func IsValidItemContidions(itemCondition string) error {
	switch itemCondition {
	case
		consts.ICClub,
		consts.ICCollectible,
		consts.ICNew,
		consts.ICRefurbished,
		consts.ICUsed:
		return nil
	default:
		return errors.New("invalid item condition")
	}
}
