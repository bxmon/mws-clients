# Amazon Marketplace Web Services (MWS) API [![Build Status](https://travis-ci.org/ezkl/go-amazon-mws-api.png?branch=master)](https://travis-ci.org/ezkl/go-amazon-mws-api)

This Amazon MWS API client is based heavily on @DDRBoxman's [go-amazon-product-api](https://github.com/DDRBoxman/go-amazon-product-api).

## Documentation

You can view the auto-generated documentation at [http://godoc.org/github.com/ezkl/go-amazon-mws-api](http://godoc.org/github.com/ezkl/go-amazon-mws-api).

## Usage

```go
package main

import (
       "fmt"
       "github.com/ezkl/go-amazon-mws-api"
)

func main() {
       var api mwsproducts.MWSClient

       api.AccessKey = ""
       api.SecretKey = ""
       api.Host = "mws.amazonservices.com"
       api.MarketplaceID = "ATVPDKIKX0DER"
       api.SellerID = ""

       asins := []string{"0195019199"}

       result,err := api.GetLowestOfferListingsForASIN(asins)

       if (err != nil) {
           fmt.Println(err)
       }

       fmt.Println(result)
}
```
