# [amazon-item] Searching Amazon Items, Powered by PA-API

[![Build Status](https://travis-ci.org/spiegel-im-spiegel/amazon-item.svg?branch=master)](https://travis-ci.org/spiegel-im-spiegel/amazon-item)
[![GitHub license](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/amazon-item/master/LICENSE)
[![GitHub release](http://img.shields.io/github/release/spiegel-im-spiegel/amazon-item.svg)](https://github.com/spiegel-im-spiegel/amazon-item/releases/latest)

## Installing

```
$ export GO111MODULE=on
$ go mod init tools
$ go get github.com/spiegel-im-spiegel/amazon-item@latest
```

## Usage

```
$ amazon-item -h
Searching Amazon Items, Powered by Product Advertising API

Usage:
  amazon-item [flags]
  amazon-item [command]

Available Commands:
  help        Help about any command
  lookup      Lookup Amazon Item
  review      Make review data for Amazon item
  search      Search Amazon Items
  version     Print the version number

Flags:
      --access-key string      Access Key ID
      --associate-tag string   Associate Tag
      --config string          config file (default $HOME/.paapi.yaml)
  -h, --help                   help for amazon-item
      --marketplace string     Marketplace (default "webservices.amazon.co.jp")
      --secret-key string      Secret Access Key

Use "amazon-item [command] --help" for more information about a command.

$ amazon-item search -h
Search Amazon Items by ItemSearch Method

Usage:
  amazon-item search [flags] keyword

Flags:
  -h, --help                    help for search
  -g, --response-group string   ResponseGroup (default "ItemAttributes,Small")
  -s, --search-index string     SearchIndex (default "All")
  -t, --template string         Template file

Global Flags:
      --access-key string      Access Key ID
      --associate-tag string   Associate Tag
      --config string          config file (default $HOME/.paapi.yaml)
      --marketplace string     Marketplace (default "webservices.amazon.co.jp")
      --secret-key string      Secret Access Key

$ amazon-item lookup -h
Lookup Amazon Item by ItemLookup Method

Usage:
  amazon-item lookup [flags]

Flags:
  -h, --help                    help for lookup
  -p, --id-type string          IdType (default "ASIN")
  -d, --item-id string          ItemId
  -g, --response-group string   ResponseGroup (default "Images,ItemAttributes,Small")
  -t, --template string         Template file
  -x, --xml                     Output with XML format

Global Flags:
      --access-key string      Access Key ID
      --associate-tag string   Associate Tag
      --config string          config file (default $HOME/.paapi.yaml)
      --marketplace string     Marketplace (default "webservices.amazon.co.jp")
      --secret-key string      Secret Access Key

$ amazon-item review -h
Make review data for Amazon item, lookup item by ItemLookup Method

Usage:
  amazon-item review [flags] description

Flags:
  -h, --help                    help for review
  -p, --id-type string          IdType (default "ASIN")
  -d, --item-id string          ItemId
  -r, --rating int              Rating of product
  -g, --response-group string   ResponseGroup (default "Images,ItemAttributes,Small")
      --review-date string      Date of review
  -t, --template string         Template file

Global Flags:
      --access-key string      Access Key ID
      --associate-tag string   Associate Tag
      --config string          config file (default $HOME/.paapi.yaml)
      --marketplace string     Marketplace (default "webservices.amazon.co.jp")
      --secret-key string      Secret Access Key
```

### Search items

```
$ cat ~/.paapi.yaml
marketplace: webservices.amazon.co.jp
associate-tag: mytag-20
access-key: AKIAIOSFODNN7EXAMPLE
secret-key: 1234567890

$ cat template/item-list.md
| ASIN | Title | Author | Binding | Publisher | PublicationDate | URL |
| ---- | ----- | ------ | ------- | --------- | --------------- | --- |
{{ range .Items }}| {{ .ASIN }} | {{ .ItemAttributes.Title }} | {{ range .ItemAttributes.Author }} {{ . }}{{ end }} | {{ .ItemAttributes.Binding }} | {{ .ItemAttributes.Publisher }} | {{ .ItemAttributes.PublicationDate }} | {{ .URL }} |
{{ end }}

$ amazon-item search -t template/item-list.md 数学ガール+フェルマーの最終定理+kindle
| ASIN | Title | Author | Binding | Publisher | PublicationDate | URL |
| ---- | ----- | ------ | ------- | --------- | --------------- | --- |
| B00AXUD4EQ | 数学ガール　フェルマーの最終定理　1 (MFコミックス　フラッパーシリーズ) |  春日旬 | Kindle版 | KADOKAWA / メディアファクトリー | 2012-11-01 | https://www.amazon.co.jp/exec/obidos/ASIN/B00AXUD4EQ/mytag-20 |
| B00I8AT1CM | 数学ガール／フェルマーの最終定理 |  結城 浩 | Kindle版 | SBクリエイティブ | 2008-07-29 | https://www.amazon.co.jp/exec/obidos/ASIN/B00I8AT1CM/mytag-20 |
| B00DONBQAI | 数学ガール　フェルマーの最終定理　3 (MFコミックス　フラッパーシリーズ) |  春日 旬 | Kindle版 | KADOKAWA / メディアファクトリー | 2013-06-27 | https://www.amazon.co.jp/exec/obidos/ASIN/B00DONBQAI/mytag-20 |
| B00AXUD4F0 | 数学ガール　フェルマーの最終定理　2 (MFコミックス　フラッパーシリーズ) |  春日旬 | Kindle版 | KADOKAWA / メディアファクトリー | 2012-11-01 | https://www.amazon.co.jp/exec/obidos/ASIN/B00AXUD4F0/mytag-20 |
| B009T4TN7Y | 数学ガール　上 (MFコミックス　フラッパーシリーズ) |  日坂 水柯 結城 浩 | Kindle版 | KADOKAWA / メディアファクトリー | 2012-09-01 | https://www.amazon.co.jp/exec/obidos/ASIN/B009T4TN7Y/mytag-20 |
| B009T4TNB0 | 数学ガール　下 (MFコミックス　フラッパーシリーズ) |  日坂 水柯 結城 浩 | Kindle版 | KADOKAWA / メディアファクトリー | 2012-09-01 | https://www.amazon.co.jp/exec/obidos/ASIN/B009T4TNB0/mytag-20 |
| B0756XMQBN | 数学ガール フェルマーの最終定理 |  春日旬 春日 旬 | Kindle版 |  |  | https://www.amazon.co.jp/exec/obidos/ASIN/B0756XMQBN/mytag-20 |
| B00ZEIEY1E | [まとめ買い] 数学ガール　フェルマーの最終定理（コミックフラッパー） |  春日旬 春日 旬 | Kindle版 |  |  | https://www.amazon.co.jp/exec/obidos/ASIN/B00ZEIEY1E/mytag-20 |
```

### Lookup an item

```
$ cat ~/.paapi.yaml
marketplace: webservices.amazon.co.jp
associate-tag: mytag-20
access-key: AKIAIOSFODNN7EXAMPLE
secret-key: 1234567890

$ cat ~/.paapi.yaml
marketplace: webservices.amazon.co.jp
associate-tag: mytag-20
access-key: AKIAIOSFODNN7EXAMPLE
secret-key: 1234567890

$ cat template/lookup.md
{{ range .Items }}<div class="hreview">
  <div class="photo"><a class="item url" href="{{ .URL }}"><img src="{{ .MediumImage.URL }}" width="{{ .MediumImage.Width }}" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="{{ .URL }}">{{ .ItemAttributes.Title }}</a></dt>
    {{ if .ItemAttributes.Author }}<dd>{{ range $i, $v := .ItemAttributes.Author }}{{ if ne $i 0 }}, {{ end }}{{ $v }}{{ end }}</dd>{{ end }}{{ if .ItemAttributes.Creator }}
	<dd>{{ range $i, $v := .ItemAttributes.Creator }}{{ if ne $i 0 }}, {{ end }}{{ $v.Value }}{{ with $v.Role }} ({{ . }}){{ end }}{{ end }}</dd>{{ end }}
    <dd>{{ .ItemAttributes.Publisher }}{{ with .ItemAttributes.PublicationDate }} {{ . }}{{ end }}{{ with .ItemAttributes.ReleaseDate }} (Release {{ . }}){{ end }}</dd>
    <dd>{{ .ItemAttributes.ProductGroup }} {{ .ItemAttributes.Binding }}</dd>
    <dd>ASIN: {{ .ASIN }}{{ with .ItemAttributes.EAN }}, EAN: {{ . }}{{ end }}{{ with .ItemAttributes.EISBN }}, EISBN: {{ . }}{{ end }}</dd>
  </dl>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>mytag-20</a> on <abbr class="dtreviewed" title="{{ $.Today }}">{{ $.Today }}</abbr> (powered by <a href="{{ $.AppURL }}" >{{ $.AppName }}</a> {{ $.AppVersion }})</p>
</div>{{ end }}

$ amazon-item lookup -d B00I8AT1CM -t template/lookup.html
<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/B00I8AT1CM/mytag-20"><img src="https://images-fe.ssl-images-amazon.com/images/I/41vT2D6sERL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/exec/obidos/ASIN/B00I8AT1CM/mytag-20">数学ガール／フェルマーの最終定理</a></dt>
    <dd>結城 浩</dd>
    <dd>SBクリエイティブ 2008-07-29 (Release 2014-03-12)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B00I8AT1CM</dd>
  </dl>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>mytag-20</a> on <abbr class="dtreviewed" title="2019-01-27">2019-01-27</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.1)</p>
</div>
```

### Make review data

```
$ cat ~/.paapi.yaml
marketplace: webservices.amazon.co.jp
associate-tag: mytag-20
access-key: AKIAIOSFODNN7EXAMPLE
secret-key: 1234567890

$ cat ~/.paapi.yaml
marketplace: webservices.amazon.co.jp
associate-tag: mytag-20
access-key: AKIAIOSFODNN7EXAMPLE
secret-key: 1234567890

$ cat template/lookup.md
{{ range .Lookup.Items }}<div class="hreview">
  <div class="photo"><a class="item url" href="{{ .URL }}"><img src="{{ .MediumImage.URL }}" width="{{ .MediumImage.Width }}" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="{{ .URL }}">{{ .ItemAttributes.Title }}</a></dt>{{ if .ItemAttributes.Author }}
	<dd>{{ range $i, $v := .ItemAttributes.Author }}{{ if ne $i 0 }}, {{ end }}{{ $v }}{{ end }}</dd>{{ end }}
    <dd>{{ .ItemAttributes.Publisher }}{{ with .ItemAttributes.PublicationDate }} {{ . }}{{ end }}{{ with .ItemAttributes.ReleaseDate }} (Release {{ . }}){{ end }}</dd>
    <dd>{{ .ItemAttributes.ProductGroup }} {{ .ItemAttributes.Binding }}</dd>
    <dd>ASIN: {{ .ASIN }}{{ with .ItemAttributes.EAN }}, EAN: {{ . }}{{ end }}{{ with .ItemAttributes.EISBN }}, EISBN: {{ . }}{{ end }}</dd>{{ if gt $.Rating 0 }}
    <dd>評価<abbr class="rating fa-sm" title="{{ $.Rating }}">{{ range $.Star }}&nbsp;{{ if . }}<i class="fas fa-star"></i>{{ else }}<i class="far fa-star"></i>{{ end }}{{ end }}</abbr></dd>{{ end }}
  </dl>
  <p class="description">{{ $.Description }}</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>mytag-20</a> on <abbr class="dtreviewed" title="{{ $.Date }}">{{ $.Date }}</abbr> (powered by <a href="{{ $.Lookup.AppURL }}" >{{ $.Lookup.AppName }}</a> {{ $.Lookup.AppVersion }})</p>
</div>{{ end }}

$ amazon-item review -d B00I8AT1CM -r 4 -t template/review.html Description
<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/B00I8AT1CM/mytag-20"><img src="https://images-fe.ssl-images-amazon.com/images/I/41vT2D6sERL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/exec/obidos/ASIN/B00I8AT1CM/mytag-20">数学ガール／フェルマーの最終定理</a></dt>
	<dd>結城 浩</dd>
    <dd>SBクリエイティブ 2008-07-29 (Release 2014-03-12)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B00I8AT1CM</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">Description</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>mytag-20</a> on <abbr class="dtreviewed" title="2019-01-27">2019-01-27</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.1)</p>
</div>
```

[amazon-item]: https://github.com/spiegel-im-spiegel/amazon-item "spiegel-im-spiegel/amazon-item: Searching Amazon Items, Powered by PA-API"
