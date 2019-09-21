package clients

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

	"github.com/bxmon/mws-products-client/consts"
)

// MWSClient defines amazon marketplace API client
type MWSClient struct {
	AccessKey string
	SecretKey string
	Host      string
	AuthToken string
	SellerID  string
}

// populateQueryParams build query params
func (client MWSClient) populateQueryParams(action string, params map[string]string) url.Values {
	values := url.Values{}

	values.Set(consts.ParamKeyAction, action)

	if client.AuthToken != "" {
		values.Set(consts.ParamKeyAuthToken, client.AuthToken)
	}

	values.Set(consts.ParamKeyAccessKey, client.AccessKey)
	values.Set(consts.ParamKeySellerID, client.SellerID)

	values.Set(consts.ParamKeyAPIVersion, consts.ParamValAPIVersion)
	values.Set(consts.ParamKeySignMethod, consts.ParamValSignMethod)
	values.Set(consts.ParamKeySignVersion, consts.ParamValSignVersion)

	// Set timestamp for request using RFC3339 format
	values.Set(consts.ParamKeyTimestamp, time.Now().UTC().Format(time.RFC3339))

	// Merge given params with url params
	for k, v := range params {
		values.Set(k, v)
	}

	return values
}

// buildUnsignedtQuery generates request url without signature
func (client MWSClient) buildUnsignedQuery(action, path string, params map[string]string) (*url.URL, error) {
	unsignedURL, err := url.Parse(client.Host)
	if err != nil {
		return nil, err
	}

	unsignedURL.Scheme = consts.MWSScheme
	unsignedURL.Host = client.Host
	unsignedURL.Path = path

	values := client.populateQueryParams(action, params)
	unsignedURL.RawQuery = values.Encode()

	return unsignedURL, nil
}

// signQueryRequest signs query request using secrect key
func (client MWSClient) signQueryRequest(origURL *url.URL, HTTPVerb string) (string, error) {
	escapeURL := strings.Replace(origURL.RawQuery, ",", "%2C", -1)
	escapeURL = strings.Replace(escapeURL, ":", "%3A", -1)

	params := strings.Split(escapeURL, "&")
	sort.Strings(params)
	sortedParams := strings.Join(params, "&")

	toSign := fmt.Sprintf("%s\n%s\n%s\n%s", HTTPVerb, origURL.Host, origURL.Path, sortedParams)

	hasher := hmac.New(sha256.New, []byte(client.SecretKey))
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
func (client MWSClient) buildSignedQuery(action, path, HTTPVerb string, params map[string]string) (string, error) {
	unsignedURL, err := client.buildUnsignedQuery(action, path, params)
	if err != nil {
		return "", err
	}

	signedURL, err := client.signQueryRequest(unsignedURL, HTTPVerb)
	if err != nil {
		return "", err
	}

	return signedURL, nil
}

// fetch make request to amazon marketplace service
func (client MWSClient) fetch(action, path, HTTPVerb string, params map[string]string) (string, error) {
	signedURL, err := client.buildSignedQuery(action, path, HTTPVerb, params)

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
