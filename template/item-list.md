| ASIN | Title | Author | Binding | Publisher | PublicationDate | URL |
| ---- | ----- | ------ | ------- | --------- | --------------- | --- |
{{ range .Items }}| {{ .ASIN }} | {{ .ItemAttributes.Title }} | {{ range .ItemAttributes.Author }} {{ . }}{{ end }} | {{ .ItemAttributes.Binding }} | {{ .ItemAttributes.Publisher }} | {{ .ItemAttributes.PublicationDate }} | {{ .URL }} |
{{ end }}
