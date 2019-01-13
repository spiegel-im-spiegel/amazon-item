package facade

import (
	"io"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/spiegel-im-spiegel/amazon-item/product"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

//newSearchCmd returns cobra.Command instance for show sub-command
func newLookupCmd(ui *rwi.RWI) *cobra.Command {
	lookupCmd := &cobra.Command{
		Use:   "lookup [flags]",
		Short: "Lookup Amazon Item",
		Long:  "Lookup Amazon Item by ItemLookup Method",
		RunE: func(cmd *cobra.Command, args []string) error {
			//options
			id, err := cmd.Flags().GetString("item-id")
			if err != nil {
				return errors.Wrap(err, "--item-id")
			}
			if len(id) == 0 {
				return errors.Wrap(os.ErrInvalid, "No ItemId property")
			}
			t, err := cmd.Flags().GetString("id-type")
			if err != nil {
				return errors.Wrap(err, "--id-type")
			}
			rg, err := cmd.Flags().GetString("response-group")
			if err != nil {
				return errors.Wrap(err, "--response-group")
			}
			xml, err := cmd.Flags().GetBool("xml")
			if err != nil {
				return errors.Wrap(err, "--xml")
			}
			tf, err := cmd.Flags().GetString("template")
			if err != nil {
				return errors.Wrap(err, "--template")
			}
			var tr io.Reader
			if len(tf) > 0 && !xml {
				file, err2 := os.Open(tf)
				if err != nil {
					return err2
				}
				defer file.Close()
				tr = file
			}

			//searching
			lookup := product.NewLookup(
				product.NewAPI(
					product.WithMarketplace(viper.GetString("marketplace")),
					product.WithAssociateTag(viper.GetString("associate-tag")),
					product.WithAccessKey(viper.GetString("access-key")),
					product.WithSecretKey(viper.GetString("secret-key")),
				),
				product.WithIDTypeForLookup(t),
				product.WithResponseGroupForLookup(rg),
			)
			if xml {
				s, err := lookup.ItemLookupXML(id)
				if err != nil {
					return err
				}
				ui.Outputln(string(s))
			} else {
				item, err := lookup.ItemLookup(id)
				if err != nil {
					return err
				}
				//output
				r, err := format(item, tr)
				if err != nil {
					return err
				}
				ui.WriteFrom(r)
			}
			return nil
		},
	}
	lookupCmd.Flags().StringP("response-group", "g", "Images,ItemAttributes,Small", "ResponseGroup")
	lookupCmd.Flags().StringP("item-id", "d", "", "ItemId")
	lookupCmd.Flags().StringP("id-type", "p", "ASIN", "IdType")
	lookupCmd.Flags().StringP("template", "t", "", "Template file")
	lookupCmd.Flags().BoolP("xml", "x", false, "Output with XML format")

	return lookupCmd
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
