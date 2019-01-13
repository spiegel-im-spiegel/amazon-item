package product

import (
	"fmt"
	"time"

	amazonproduct "github.com/DDRBoxman/go-amazon-product-api"
)

//Result is dataset for result of SearchItem
type Result struct {
	AssociateTag string
	Today        string //RFC3339 format
	Items        []amazonproduct.Item
	AppName      string
	AppVersion   string
	AppURL       string
}

//NewResult returns Result instance
func NewResult(associateTag string, today time.Time, items []amazonproduct.Item) *Result {
	return &Result{
		AssociateTag: associateTag,
		Today:        today.Format("2006-01-02"),
		Items:        items,
		AppName:      Name,
		AppVersion:   Version,
		AppURL:       URL,
	}
}

//MakeURL makes URL in Result instance
func (res *Result) MakeURL() *Result {
	if res == nil {
		return res
	}
	if len(res.Items) == 0 {
		return res
	}
	for i := 0; i < len(res.Items); i++ {
		res.Items[i].URL = fmt.Sprintf("https://www.amazon.co.jp/exec/obidos/ASIN/%s/%s", res.Items[i].ASIN, res.AssociateTag)
	}
	return res
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
