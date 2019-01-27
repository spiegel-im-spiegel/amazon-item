package product

import (
	"net/http"

	amazonproduct "github.com/spiegel-im-spiegel/go-amazon-product-api" //replace from github.com/DDRBoxman/go-amazon-product-api (temporary)
)

//API is interface class for github.com/DDRBoxman/go-amazon-product-api
type API interface {
	ItemLookupWithParams(map[string]string) (string, error)
	ItemSearch(string, map[string]string) (string, error)
}

//APIOptFunc is self-referential function for functional options pattern
type APIOptFunc func(*amazonproduct.AmazonProductAPI)

//NewAPI returns API instance
func NewAPI(opts ...APIOptFunc) API {
	api := &amazonproduct.AmazonProductAPI{Client: &http.Client{}}
	for _, opt := range opts {
		opt(api)
	}
	return api
}

//WithMarketplace returns function for setting Marketplace
func WithMarketplace(mp string) APIOptFunc {
	return func(api *amazonproduct.AmazonProductAPI) {
		if api != nil {
			api.Host = mp
		}
	}
}

//WithAssociateTag returns function for setting Associate Tag
func WithAssociateTag(tag string) APIOptFunc {
	return func(api *amazonproduct.AmazonProductAPI) {
		if api != nil {
			api.AssociateTag = tag
		}
	}
}

//WithAccessKey returns function for setting Access Key
func WithAccessKey(key string) APIOptFunc {
	return func(api *amazonproduct.AmazonProductAPI) {
		if api != nil {
			api.AccessKey = key
		}
	}
}

//WithSecretKey returns function for setting Secret Access Key
func WithSecretKey(key string) APIOptFunc {
	return func(api *amazonproduct.AmazonProductAPI) {
		if api != nil {
			api.SecretKey = key
		}
	}
}

func getAssociateTag(api API) string {
	if paapi, ok := api.(*amazonproduct.AmazonProductAPI); ok {
		return paapi.AssociateTag
	}
	return ""
}

/* MIT License
 *
 * Copyright 2019 Spiegel
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
