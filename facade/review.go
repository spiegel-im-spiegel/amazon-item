package facade

import (
	"io"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/spiegel-im-spiegel/amazon-item/product"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

const MAX_STAR = 5

//Review is class of review data
type Review struct {
	Date        string
	Rating      int
	Star        [MAX_STAR]bool
	Description string
	Lookup      *product.Result
}

//newReviewCmd returns cobra.Command instance for show sub-command
func newReviewCmd(ui *rwi.RWI) *cobra.Command {
	reviewCmd := &cobra.Command{
		Use:   "review [flags] description",
		Short: "Make review data for Amazon item",
		Long:  "Make review data for Amazon item, lookup item by ItemLookup Method",
		RunE: func(cmd *cobra.Command, args []string) error {
			rev := &Review{Date: time.Now().Format("2006-01-02")}
			//oproperties for PA-API
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
			//Rating for Amazon product
			rating, err := cmd.Flags().GetInt("rating")
			if err != nil {
				return errors.Wrap(err, "--rating")
			}
			if rating > MAX_STAR {
				rating = MAX_STAR
			} else if rating < 0 {
				rating = 0
			}
			rev.Rating = rating
			for i := 0; i < MAX_STAR; i++ {
				if rating > i {
					rev.Star[i] = true
				}
			}
			//Date of review
			dt, err := cmd.Flags().GetString("review-date")
			if err != nil {
				return errors.Wrap(err, "--review-date")
			}
			if len(dt) > 0 {
				rev.Date = dt
			}
			//Template data
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

			//Description
			if len(args) > 1 {
				return errors.Wrap(os.ErrInvalid, strings.Join(args, " "))
			} else if len(args) == 1 {
				rev.Description = args[0]
			} else {
				w := &strings.Builder{}
				io.Copy(w, ui.Reader())
				rev.Description = w.String()
			}

			//lookup item
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
			item, err := lookup.ItemLookup(id)
			if err != nil {
				return err
			}
			//output
			rev.Lookup = item
			r, err := format(rev, tr)
			if err != nil {
				return err
			}
			ui.WriteFrom(r)
			return nil
		},
	}
	reviewCmd.Flags().StringP("response-group", "g", "Images,ItemAttributes,Small", "ResponseGroup")
	reviewCmd.Flags().StringP("item-id", "d", "", "ItemId")
	reviewCmd.Flags().StringP("id-type", "p", "ASIN", "IdType")
	reviewCmd.Flags().IntP("rating", "r", 0, "Rating of product")
	reviewCmd.Flags().StringP("review-date", "", "", "Date of review")
	reviewCmd.Flags().StringP("template", "t", "", "Template file")

	return reviewCmd
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
