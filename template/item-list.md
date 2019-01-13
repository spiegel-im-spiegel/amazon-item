| ASIN | Title | Author | Binding | EAN | Publisher | PublicationDate | URL |
| ---- | ----- | ------ | ------- | --- | --------- | --------------- | --- |
{{ range .Items }}| {{ .ASIN }} | {{ .ItemAttributes.Title }} | {{ range .ItemAttributes.Author }} {{ . }}{{ end }} | {{ .ItemAttributes.Binding }} | {{ .ItemAttributes.EAN }} | {{ .ItemAttributes.Publisher }} | {{ .ItemAttributes.ReleaseDate }} | {{ .URL }} |
{{ end }}
