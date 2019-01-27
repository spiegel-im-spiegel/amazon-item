package product

import (
	"encoding/xml"
	"strconv"
	"time"

	amazonproduct "github.com/spiegel-im-spiegel/go-amazon-product-api" //replace from github.com/DDRBoxman/go-amazon-product-api (temporary)
)

//Srch is class for ItemSearch method
type Srch struct {
	PaAPI         API
	SearchIndex   string
	ResponseGroup string
}

//SrchOptFunc is self-referential function for functional options pattern
type SrchOptFunc func(*Srch)

//NewSrch returns Srch instance
func NewSrch(api API, opts ...SrchOptFunc) *Srch {
	s := &Srch{PaAPI: api, SearchIndex: "All", ResponseGroup: "ItemAttributes,Small"}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

//WithSearchIndexForSrch returns function for setting SearchIndex
func WithSearchIndexForSrch(s string) SrchOptFunc {
	return func(srch *Srch) {
		if srch != nil {
			srch.SearchIndex = s
		}
	}
}

//WithResponseGroupForSrch returns function for setting ResponseGroup
func WithResponseGroupForSrch(rg string) SrchOptFunc {
	return func(srch *Srch) {
		if srch != nil {
			srch.ResponseGroup = rg
		}
	}
}

//ItemSearch returns result of ItemSearch API
func (srch *Srch) ItemSearch(keywords string) (*Result, error) {
	if srch == nil {
		return &Result{}, nil
	}
	params := map[string]string{
		"ResponseGroup": srch.ResponseGroup,
		"Keywords":      keywords,
	}
	items := []amazonproduct.Item{}
	for page := 1; page <= 5; page++ {
		params["ItemPage"] = strconv.FormatInt(int64(page), 10)
		result, err := srch.PaAPI.ItemSearch(srch.SearchIndex, params)
		if err != nil {
			return &Result{}, err
		}
		res := &amazonproduct.ItemSearchResponse{}
		if err := xml.Unmarshal([]byte(result), res); err != nil {
			return &Result{}, err
		}
		if !res.Items.Request.IsValid {
			break
		}
		items = append(items, res.Items.Items...)
	}
	r := NewResult(getAssociateTag(srch.PaAPI), time.Now(), items)
	return r.MakeURL(), nil
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
