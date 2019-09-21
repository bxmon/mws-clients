package consts

// Marketplace id constants
const (
	MIDBrazil             = "A2Q3Y263D00KWC"
	MIDCanada             = "A2EUQ1WTGCTBG2"
	MIDMexico             = "A1AM78C64UM0Y8"
	MIDUS                 = "ATVPDKIKX0DER"
	MIDUnitedArabEmirates = "A2VIGQ35RCS4UG"
	MIDGermany            = "A1PA6795UKMFR9"
	MIDSpain              = "A1RKKUPIHCS9HS"
	MIDFrance             = "A13V1IB3VIYZZH"
	MIDUK                 = "A1F83G8C2ARO7P"
	MIDIndia              = "A21TJRUUN4KGV"
	MIDItaly              = "APJ6JRA9NG5V4"
	MIDTurkey             = "A33AVAJ2PDY3EV"
	MIDAustralia          = "A39IBJ37TRP1C6"
	MIDJapan              = "A1VC38T7YXB528"
	MIDChina              = "AAHKV2X7AFYLW"
)

// Marketplace endpoint constants
const (
	MEPBrazil             = "mws.amazonservices.com"
	MEPCanada             = "mws.amazonservices.ca"
	MEPMexico             = "mws.amazonservices.com.mx"
	MEPUS                 = "mws.amazonservices.com"
	MEPUnitedArabEmirates = "mws.amazonservices.ae"
	MEPGermany            = "mws-eu.amazonservices.com"
	MEPSpain              = "mws-eu.amazonservices.com"
	MEPFrance             = "mws-eu.amazonservices.com"
	MEPUK                 = "mws-eu.amazonservices.com"
	MEPIndia              = "mws.amazonservices.in"
	MEPItaly              = "mws-eu.amazonservices.com"
	MEPTurkey             = "mws-eu.amazonservices.com"
	MEPAustralia          = "mws.amazonservices.com.au"
	MEPJapan              = "mws.amazonservices.jp"
	MEPChina              = "mws.amazonservices.com.cn"
)

// MWS client constants
const (
	MWSGet          = "GET"
	MWSPost         = "POST"
	MWSScheme       = "https"
	MWSProductsPath = "/Products/2011-10-01"
)

// MWS product identifier types constants
const (
	PIDTypesASIN = "ASIN"
	PIDTypesGCID = "GCID"
	PIDTypesSKU  = "SellerSKU"
	PIDTypesUPC  = "UPC"
	PIDTypesEAN  = "EAN"
	PIDTypesISBN = "ISBN"
	PIDTypesJAN  = "JAN"
)

// MWS product item condition constants
const (
	ICNew         = "New"
	ICUsed        = "Used"
	ICCollectible = "Collectible"
	ICRefurbished = "Refurbished"
	ICClub        = "Club"
)

// MWS param key constants
const (
	ParamKeyAccessKey   = "AWSAccessKeyId"
	ParamKeyAction      = "Action"
	ParamKeyAPIVersion  = "Version"
	ParamKeyAuthToken   = "MWSAuthToken"
	ParamKeySellerID    = "SellerId"
	ParamKeySignature   = "Signature"
	ParamKeySignVersion = "SignatureVersion"
	ParamKeySignMethod  = "SignatureMethod"
	ParamKeyTimestamp   = "Timestamp"
)

// MWS param val constants
const (
	ParamValSignVersion = "2"
	ParamValSignMethod  = "HmacSHA256"
	ParamValAPIVersion  = "2011-10-01"
)

// MWS product action contansts
const (
	GetMatchingProductForID       = "GetMatchingProductForId"
	GetCompetitivePricingForASIN  = "GetCompetitivePricingForASIN"
	GetLowestPricedOffersForASIN  = "GetLowestPricedOffersForASIN"
	GetLowestOfferListingsForASIN = "GetLowestOfferListingsForASIN"
)
