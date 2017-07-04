package main

import (
	"net/http"
	"strings"

	"github.com/linnv/logx"
)

func main() {
	http.HandleFunc("/short", func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("good"))
		fields := []string{
			// "www.jialinwu.com", "redirecting to targer",
			"{{.url}}", "http://www.jialinwu.com",
			"{{.title}}", "redirecting to targer",
		}
		page := strings.NewReplacer(fields...).Replace(pageTemplate)
		w.Write([]byte(page))
	})
	logx.Debugln("listening on ", 61900)
	http.ListenAndServe(":61900", nil)
	return
}

const pageTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="utf-8" />
  <meta http-equiv="refresh" content="0; url={{.url}}" />
  <title>{{.title}}</title>
</head>
<body>
  <p><a href="{{.url}}">Click here</a> if not redirected automatically.</p>
</body>
</html>
`
