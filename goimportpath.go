package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const (
	goTagsTemplate = `<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="{{.ImportPath}} git {{.RepoUrl}}">
</head>
<body>
</body>
</html>
`
)

type Canonical struct {
	ImportPath string
	RepoUrl    string
}

var (
	temp *template.Template

	canonicals = map[string]Canonical{ // URL.Path ==> canonical
		"/weakand": {
			ImportPath: "topic.ai/weakand",
			RepoUrl:    "github.com/wangkuiyi/weakand"},
	}
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)

	if canonical, ok := canonicals[r.URL.Path]; ok {
		temp.Execute(w, canonical)
	} else {
		fmt.Fprintf(w, "Hello! Welcome to %s!", r.URL.Host)
	}
}

func main() {
	addr := flag.String("addr", ":443", "Listening address")
	cert := flag.String("cert", ".tls/server.crt", "certificate file")
	key := flag.String("key", ".tls/server.key", "private key file")

	flag.Parse()

	if t, e := template.New("gotags").Parse(goTagsTemplate); e != nil {
		log.Fatal(e)
	} else {
		temp = t
	}

	http.HandleFunc("/", viewHandler)
	http.ListenAndServeTLS(*addr, *cert, *key, nil)
}
