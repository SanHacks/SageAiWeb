package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {

	//Initiate Router
	router := mux.NewRouter()

	//router.Use(checkLogin)
	//Map the routes to the handlers in the backend api handler and
	//Handle 404 errors and redirect to /platform
	//Open Up routes for the frontend engine to use the backend api handler
	//Such as /backend/api/products (api) and /generatedProducts/ directory
	platformRouter(router)

	//Templates allocation to variables for the frontend engine
	joinPlatform, loginPlatform, platform := templates()

	templateHandler(joinPlatform, loginPlatform, platform)
	http.Handle("/", router)

	port := openPort()

	log.Printf("Listening on port %s", port)
	log.Printf("🚀🚀🚀🚀AIGEN🚀🚀🚀🚀")
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

}

// Opens Up The Port 8080, although it can get changed by the PORT env variable
func openPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Enojy! %s", port)
	}
	return port
}

func templateHandler(login *template.Template, join *template.Template, platform *template.Template) {
	http.HandleFunc("/join", func(w http.ResponseWriter, r *http.Request) {
		//IF THE REQUEST IS NOT A POST
		//authCheckIn(w, r)

		if r.Method != http.MethodPost {
			//Render the Home Page
			err := join.Execute(w, nil)
			if err != nil {
				log.Println("Error in Rendering the Join Page")
			}
			return
		}

	})

	http.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		//authCheckIn(w, r)
		if r.Method != http.MethodPost {
			//Render the Home Page
			err := login.Execute(w, nil)
			if err != nil {
				log.Println("Error in Rendering the Join Page")
			}
			return
		}

	})

	http.HandleFunc("/platform", func(w http.ResponseWriter, r *http.Request) {
		//authCheckIn(w, r)
		if r.Method != http.MethodPost {
			//Render the Home Page
			err := platform.Execute(w, nil)
			if err != nil {
				log.Println("Error in Rendering the Join Page")
			}
			return
		}

	})
}

func templates() (*template.Template, *template.Template, *template.Template) {
	platform := template.Must(template.ParseFiles("Web/index.html"))
	//inventory := template.Must(template.ParseFiles("frontend/inventory.html"))
	//product := template.Must(template.ParseFiles("frontend/product.html"))
	//order := template.Must(template.ParseFiles("frontend/purchase.html"))
	//errorPage := template.Must(template.ParseFiles("frontend/404.html"))
	//orderSuccess := template.Must(template.ParseFiles("frontend/orderSuccess.html"))
	joinPlatform := template.Must(template.ParseFiles("Web/signup.html"))
	loginPlatform := template.Must(template.ParseFiles("Web/login.html"))
	return joinPlatform, loginPlatform, platform
}
