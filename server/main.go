// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// [START gae_go111_app]

// Sample helloworld is an App Engine app.
package main

// [START import]
import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	// "github.com/go-chi/chi"
	//"google.golang.org/appengine"
)

// [END import]
// [START main_func]

type Homepage struct {
	Title string
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/vue/", vueHandler)

	http.Handle("/dist/", http.FileServer(http.Dir(".")))

	// [START setting_port]
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
	// [END setting_port]

	//r := app.SetupEngine(".")
	// http.Handle("/", r)
	// appengine.Main()
}

/* func SetupEngine(basePath string) *chi.Mux {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Get("/dist")

	})
} */

// [END main_func]

// [START indexHandler]

// indexHandler responds to requests with our greeting.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, World!")
}

// [END indexHandler]
// [END gae_go111_app]

func vueHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/vue/" {
		http.NotFound(w, r)
		return
	}
	tmpl, err := template.ParseFiles("./templates/home.html")

	if err != nil {
		fmt.Fprint(w, err)
	}

	data := Homepage{
		Title: "VueJS Page",
	}

	tmpl.Execute(w, data)
}
