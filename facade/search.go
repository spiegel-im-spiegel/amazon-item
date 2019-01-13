package facade

import (
	"io"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/spiegel-im-spiegel/amazon-item/product"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

//newSearchCmd returns cobra.Command instance for show sub-command
func newSearchCmd(ui *rwi.RWI) *cobra.Command {
	searchCmd := &cobra.Command{
		Use:   "search [flags] keyword",
		Short: "Search Amazon Items",
		Long:  "Search Amazon Items by ItemSearch Method",
		RunE: func(cmd *cobra.Command, args []string) error {
			//options
			si, err := cmd.Flags().GetString("search-index")
			if err != nil {
				return errors.Wrap(err, "--search-index")
			}
			rg, err := cmd.Flags().GetString("response-group")
			if err != nil {
				return errors.Wrap(err, "--response-group")
			}
			tf, err := cmd.Flags().GetString("template")
			if err != nil {
				return errors.Wrap(err, "--template")
			}
			var tr io.Reader
			if len(tf) > 0 {
				file, err2 := os.Open(tf)
				if err != nil {
					return err2
				}
				defer file.Close()
				tr = file
			}

			//keyword
			if len(args) == 0 {
				return errors.Wrap(os.ErrInvalid, "No Keyword property")
			} else if len(args) > 1 {
				return errors.Wrap(os.ErrInvalid, strings.Join(args, " "))
			}
			keyword := args[0]

			//searching
			srch := product.NewSrch(
				product.NewAPI(
					product.WithMarketplace(viper.GetString("marketplace")),
					product.WithAssociateTag(viper.GetString("associate-tag")),
					product.WithAccessKey(viper.GetString("access-key")),
					product.WithSecretKey(viper.GetString("secret-key")),
				),
				product.WithSearchIndexForSrch(si),
				product.WithResponseGroupForSrch(rg),
			)
			items, err := srch.ItemSearch(keyword)
			if err != nil {
				return err
			}

			//output
			r, err := format(items, tr)
			if err != nil {
				return err
			}
			ui.WriteFrom(r)
			return nil
		},
	}
	searchCmd.Flags().StringP("search-index", "s", "All", "SearchIndex")
	searchCmd.Flags().StringP("response-group", "g", "ItemAttributes,Small", "ResponseGroup")
	searchCmd.Flags().StringP("template", "t", "", "Template file")

	return searchCmd
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
