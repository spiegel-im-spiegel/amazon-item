module github.com/spiegel-im-spiegel/amazon-item

require (
	github.com/DDRBoxman/go-amazon-product-api v0.0.0-20190113062856-6736abd38089
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/mitchellh/go-homedir v1.0.0
	github.com/pkg/errors v0.8.1
	github.com/spf13/cobra v0.0.3
	github.com/spf13/viper v1.3.1
	github.com/spiegel-im-spiegel/gocli v0.8.1
)

replace github.com/DDRBoxman/go-amazon-product-api v0.0.0-20190113062856-6736abd38089 => github.com/spiegel-im-spiegel/go-amazon-product-api v0.0.0-20190113075218-1369f41b1e57
