package facade

import (
	"runtime"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/spiegel-im-spiegel/amazon-item/product"
	"github.com/spiegel-im-spiegel/gocli/exitcode"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

var (
	cfgFile string //config file
)

//newRootCmd returns cobra.Command instance for root command
func newRootCmd(ui *rwi.RWI, args []string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   product.Name,
		Short: "Searching Amazon Items",
		Long:  "Searching Amazon Items, Powered by Product Advertising API",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("no command")
		},
	}
	rootCmd.SetArgs(args)
	rootCmd.SetOutput(ui.ErrorWriter())
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default $HOME/.paapi.yaml)")
	rootCmd.PersistentFlags().StringP("marketplace", "", "webservices.amazon.co.jp", "Marketplace")
	rootCmd.PersistentFlags().StringP("associate-tag", "", "", "Associate Tag")
	rootCmd.PersistentFlags().StringP("access-key", "", "", "Access Key ID")
	rootCmd.PersistentFlags().StringP("secret-key", "", "", "Secret Access Key")
	viper.BindPFlag("marketplace", rootCmd.PersistentFlags().Lookup("marketplace"))
	viper.BindPFlag("associate-tag", rootCmd.PersistentFlags().Lookup("associate-tag"))
	viper.BindPFlag("access-key", rootCmd.PersistentFlags().Lookup("access-key"))
	viper.BindPFlag("secret-key", rootCmd.PersistentFlags().Lookup("secret-key"))
	rootCmd.AddCommand(newVersionCmd(ui))
	rootCmd.AddCommand(newSearchCmd(ui))
	rootCmd.AddCommand(newLookupCmd(ui))
	rootCmd.AddCommand(newReviewCmd(ui))

	return rootCmd
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			panic(err)
		}
		// Search config in home directory with name ".paapi.yaml" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".paapi")
	}
	viper.AutomaticEnv() // read in environment variables that match
	viper.ReadInConfig() // If a config file is found, read it in.
}

//Execute is called from main function
func Execute(ui *rwi.RWI, args []string) (exit exitcode.ExitCode) {
	defer func() {
		//panic hundling
		if r := recover(); r != nil {
			ui.OutputErrln("Panic:", r)
			for depth := 0; ; depth++ {
				pc, src, line, ok := runtime.Caller(depth)
				if !ok {
					break
				}
				ui.OutputErrln(" ->", depth, ":", runtime.FuncForPC(pc).Name(), ":", src, ":", line)
			}
			exit = exitcode.Abnormal
		}
	}()

	//execution
	exit = exitcode.Normal
	if err := newRootCmd(ui, args).Execute(); err != nil {
		exit = exitcode.Abnormal
	}
	return
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
