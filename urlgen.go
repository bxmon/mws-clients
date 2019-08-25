package mwsproducts

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

// ProductsMWSAPI defines amazon marketplace web service product API
type ProductsMWSAPI struct {
	AccessKey     string
	SecretKey     string
	Host          string
	AuthToken     string
	SellerID      string
}

// populateQueryParams build query params
func (api ProductsMWSAPI) populateQueryParams(action string, params map[string]string) url.Values {
	values := url.Values{}

	values.Set(ParamKeyAction, action)

	if api.AuthToken != "" {
		values.Set(ParamKeyAuthToken, api.AuthToken)
	}

	values.Set(ParamKeyAccessKey, api.AccessKey)
	values.Set(ParamKeySellerID, api.SellerID)

	values.Set(ParamKeyAPIVersion, ParamValAPIVersion)
	values.Set(ParamKeySignMethod, ParamValSignMethod)
	values.Set(ParamKeySignVersion, ParamValSignVersion)

	// Set timestamp for request using RFC3339 format
	values.Set(ParamKeyTimestamp, time.Now().UTC().Format(time.RFC3339))

	// Merge given params with url params
	for k, v := range params {
		values.Set(k, v)
	}

	return values
}

// buildUnsignedtQuery generates request url without signature
func (api ProductsMWSAPI) buildUnsignedQuery(action, path string, params map[string]string) (*url.URL, error) {
	unsignedURL, err := url.Parse(api.Host)
	if err != nil {
		return nil, err
	}

	unsignedURL.Scheme = MWSScheme
	unsignedURL.Host = api.Host
	unsignedURL.Path = path

	values := api.populateQueryParams(action, params)
	unsignedURL.RawQuery = values.Encode()

	return unsignedURL, nil
}

// signQueryRequest signs query request using secrect key
func (api ProductsMWSAPI) signQueryRequest(origURL *url.URL, HTTPVerb string) (string, error) {
	escapeURL := strings.Replace(origURL.RawQuery, ",", "%2C", -1)
	escapeURL = strings.Replace(escapeURL, ":", "%3A", -1)

	params := strings.Split(escapeURL, "&")
	sort.Strings(params)
	sortedParams := strings.Join(params, "&")

	toSign := fmt.Sprintf("%s\n%s\n%s\n%s", HTTPVerb, origURL.Host, origURL.Path, sortedParams)

	hasher := hmac.New(sha256.New, []byte(api.SecretKey))
	_, err := hasher.Write([]byte(toSign))
	if err != nil {
		return "", err
	}

	hash := base64.StdEncoding.EncodeToString(hasher.Sum(nil))

	hash = url.QueryEscape(hash)

	newParams := fmt.Sprintf("%s&Signature=%s", sortedParams, hash)

	origURL.RawQuery = newParams

	return origURL.String(), nil
}

// buildSignedQuery generates request url with signature
func (api ProductsMWSAPI) buildSignedQuery(action, path, HTTPVerb string, params map[string]string) (string, error) {
	unsignedURL, err := api.buildUnsignedQuery(action, path, params)
	if err != nil {
		return "", err
	}

	signedURL, err := api.signQueryRequest(unsignedURL, HTTPVerb)
	if err != nil {
		return "", err
	}

	return signedURL, nil
}

// fetch make request to fetch data from amazon
func (api ProductsMWSAPI) fetch(action, path, HTTPVerb string, params map[string]string) (string, error) {
	signedURL, err := api.buildSignedQuery(action, path, HTTPVerb, params)

	resp, err := http.Get(signedURL)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
