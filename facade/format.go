package facade

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
	"text/template"
)

func format(obj interface{}, tr io.Reader) (io.Reader, error) {
	buf := &bytes.Buffer{}
	if tr == nil {
		s, err := json.Marshal(obj)
		if err != nil {
			return buf, err
		}
		if _, err := buf.Write(s); err != nil {
			return buf, err
		}
	} else {
		tmpdata := &strings.Builder{}
		io.Copy(tmpdata, tr)
		t, err := template.New("Formatting").Parse(tmpdata.String())
		if err != nil {
			return buf, err
		}
		if err := t.Execute(buf, obj); err != nil {
			return buf, err
		}
	}
	return buf, nil
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
