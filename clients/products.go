package clients

import (
	"fmt"

	"github.com/bxmon/mws-products-client/consts"
	"github.com/bxmon/mws-products-client/utils"
	"github.com/bxmon/mws-types/reqs"
)

// GetMatchingProductForID operation returns a list of products and their attributes,
// based on a list of product identifier values that you specify. Possible product
// identifiers are ASIN, GCID, SellerSKU, UPC, EAN, ISBN, and JAN.
//
// Restriction:
//    IdType values: ASIN, GCID, SellerSKU, UPC, EAN, ISBN, JAN
//    Maximum: Five Id values
//
// Throttling:
//    Maximum request quota: 20 requests
//    Restore rate: One request every five seconds
//    Hourly request quota: 720 requests per hour
func (client MWSClient) GetMatchingProductForID(request *reqs.MatchingProductParams) (string, error) {
	if err := utils.IsValidProductIDTypes(request.IDType); err != nil {
		return "", err
	}

	if err := utils.IsValidList(request.IDList, 5); err != nil {
		return "", err
	}

	params := make(map[string]string)

	// Format IdList.Id match with MWS query requirement
	for k, v := range request.IDList {
		key := fmt.Sprintf("IdList.Id.%d", (k + 1))
		params[key] = string(v)
	}

	params["IdType"] = request.IDType
	params["MarketplaceId"] = request.MarketplaceID

	resp, err := client.fetch(consts.GetMatchingProductForID, consts.MWSProductsPath, consts.MWSGet, params)
	if err != nil {
		return "", err
	}

	return resp, nil
}

// GetCompetitivePricingForASIN operation returns the current competitive pricing of a product,
// based on the ASIN and MarketplaceId that you specify. This operation returns pricing for active
// offer listings based on two pricing models: New Buy Box Price and Used Buy Box Price. These pricing
// models are equivalent to the main Buy Box Price and the subordinate Buy Box Price, respectively, on
// a detail page from an Amazon marketplace website. Note that products with active offer listings might
// not return either of these prices. This could happen, for example, if none of the sellers with offer
// listings for a product are qualified for the New Buy Box or the Used Buy Box.
//
// Throttling:
//    Maximum request quota: 20 requests
//    Restore rate: 10 items every second
//    Hourly request quota: 36000 requests per hour
func (client MWSClient) GetCompetitivePricingForASIN(request *reqs.CompetitivePricingParams) (string, error) {
	if err := utils.IsValidList(request.ASINList, 20); err != nil {
		return "", err
	}

	params := make(map[string]string)

	// Format ASINList.Id match with MWS query requirement
	for k, v := range request.ASINList {
		key := fmt.Sprintf("ASINList.ASIN.%d", (k + 1))
		params[key] = string(v)
	}

	params["MarketplaceId"] = request.MarketplaceID

	resp, err := client.fetch(consts.GetCompetitivePricingForASIN, consts.MWSProductsPath, consts.MWSGet, params)
	if err != nil {
		return "", err
	}

	return resp, nil
}

// GetLowestPricedOffersForASIN operation returns the top 20 offers for a given MarketplaceId, ASIN,
// and ItemCondition that you specify. The top 20 offers are determined by the lowest landed price,
// which is the price plus shipping minus Amazon Points. If multiple sellers are charging the same
// landed price, the results will be returned in random order.
//
// Throttling:
//    Maximum request quota: 10 requests
//    Restore rate: Five items every second
//    Hourly request quota: 200 requests per hour
func (client MWSClient) GetLowestPricedOffersForASIN(request *reqs.LowestPricedOffersParams) (string, error) {
	if err := utils.IsValidItemContidions(request.ItemCondition); err != nil {
		return "", err
	}

	params := make(map[string]string)

	params["ASIN"] = request.ASIN
	params["ItemCondition"] = request.ItemCondition
	params["MarketplaceId"] = request.MarketplaceID

	resp, err := client.fetch(consts.GetLowestPricedOffersForASIN, consts.MWSProductsPath, consts.MWSGet, params)
	if err != nil {
		return "", err
	}

	return resp, nil
}

// GetLowestOfferListingsForASIN returns pricing information for the lowest-price active offer listings
// for up to 20 products, based on ASIN.
//
// Throttling:
//    Maximum request quota: 20 requests
//    Restore rate: 10 items every second
//    Hourly request quota: 36000 requests per hour
func (client MWSClient) GetLowestOfferListingsForASIN(request *reqs.LowestOfferListingsParams) (string, error) {
	if err := utils.IsValidList(request.ASINList, 20); err != nil {
		return "", err
	}

	params := make(map[string]string)

	// itemCondition is optional
	if request.ItemCondition != "" {
		if err := utils.IsValidItemContidions(request.ItemCondition); err != nil {
			return "", err
		}

		params["ItemCondition"] = request.ItemCondition
	}

	// Format ASINList.Id match with MWS query requirement
	for k, v := range request.ASINList {
		key := fmt.Sprintf("ASINList.ASIN.%d", (k + 1))
		params[key] = string(v)
	}

	params["MarketplaceId"] = request.MarketplaceID

	resp, err := client.fetch(consts.GetLowestOfferListingsForASIN, consts.MWSProductsPath, consts.MWSGet, params)
	if err != nil {
		return "", err
	}

	return resp, nil
}
