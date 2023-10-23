package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

// templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP request.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates",
			t.filename)))
	})

	// compile the template and write the output to the response writter.
	err := t.templ.Execute(w, r)
	if err != nil {
		log.Fatal("Error compiling template", err)
	}
}

func main() {

	var addr = flag.String("addr", ":8080", "The addr of the  application.")
	flag.Parse() // parse the flags

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte(`<html>
	//       <head>
	//         <title>Chat</title>
	//       </head>
	//       <body>
	//         Let's chat!
	//       </body>
	//     </html>
	// 		`))
	// })

	// templateHandler struct implements a method with signature func(w http.ResponseWriter, r *http.Request)
	// which is all the http.Handle to handle an http request. Therefore we can pass an instance of templateHandler instead of
	// a function.
	http.Handle("/", &templateHandler{
		filename: "chat.html",
	})

	// room implements a serveHTTP method.
	r := newRoom()
	http.Handle("/room", r)
	go r.run()

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("Could not start server", err)
	}
}
