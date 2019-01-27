package product

import (
	"encoding/xml"
	"errors"
	"time"

	amazonproduct "github.com/spiegel-im-spiegel/go-amazon-product-api" //replace from github.com/DDRBoxman/go-amazon-product-api (temporary)
)

//Lookup is class for ItemLookup method
type Lookup struct {
	PaAPI         API
	IDType        string
	ResponseGroup string
}

//LookupOptFunc is self-referential function for functional options pattern
type LookupOptFunc func(*Lookup)

//NewLookup returns Lookup instance
func NewLookup(api API, opts ...LookupOptFunc) *Lookup {
	l := &Lookup{PaAPI: api, IDType: "ASIN", ResponseGroup: "Images,ItemAttributes,Small"}
	for _, opt := range opts {
		opt(l)
	}
	return l
}

//WithIDTypeForLookup returns function for setting IdType
func WithIDTypeForLookup(s string) LookupOptFunc {
	return func(Lookup *Lookup) {
		if Lookup != nil {
			Lookup.IDType = s
		}
	}
}

//WithResponseGroupForLookup returns function for setting ResponseGroup
func WithResponseGroupForLookup(rg string) LookupOptFunc {
	return func(Lookup *Lookup) {
		if Lookup != nil {
			Lookup.ResponseGroup = rg
		}
	}
}

//ItemLookupXML returns result of ItemSearch API
func (l *Lookup) ItemLookupXML(id string) (string, error) {
	if l == nil {
		return "", nil
	}
	params := map[string]string{
		"IdType":        l.IDType,
		"ResponseGroup": l.ResponseGroup,
		"ItemId":        id,
	}
	return l.PaAPI.ItemLookupWithParams(params)
}

//ItemLookup returns result of ItemSearch API
func (l *Lookup) ItemLookup(id string) (*Result, error) {
	result, err := l.ItemLookupXML(id)
	if err != nil {
		return &Result{}, err
	}
	res := &amazonproduct.ItemLookupResponse{}
	if err := xml.Unmarshal([]byte(result), res); err != nil {
		return &Result{}, err
	}
	if !res.Items.Request.IsValid {
		return &Result{}, errors.New("IsValid")
	}
	r := NewResult(getAssociateTag(l.PaAPI), time.Now(), res.Items.Item)
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
