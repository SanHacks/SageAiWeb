package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func platformRouter(router *mux.Router) {
	router.NotFoundHandler = routeNotFoundError()
	//Backend API Routes Definition (API) (ie: open route and map to handler)
	//router.HandleFunc("/backend/api/products", ProductsHandler).Methods("GET")

	////Submission Handlers (ie: open route and map to handler)
	//router.HandleFunc("/checkout", CheckoutHandler).Methods("POST")
	//router.HandleFunc("/postsignup", SignupHandler).Methods("POST")
	//router.HandleFunc("/postlogin", LoginHandler).Methods("POST")
	//router.HandleFunc("/postlogout", LogoutHandler).Methods("POST")

	//File Server for Platform CDN and Frontend Engine Resources (CSS, JS, Images, Fonts and stuff)
	router.PathPrefix("/generatedProducts/").Handler(http.StripPrefix("/generatedProducts/", http.FileServer(http.Dir("./generatedProducts/"))))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./Web/assets/css/"))))
	router.PathPrefix("/assets/css/").Handler(http.StripPrefix("/assets/css/", http.FileServer(http.Dir("./Web/assets/css/"))))
	router.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("./Web/assets/img/"))))
	router.PathPrefix("/js/").Handler(http.StripPrefix("/Web/assets/js/", http.FileServer(http.Dir("./Web/assets/js/"))))
	router.PathPrefix("/images/").Handler(http.StripPrefix("/images /", http.FileServer(http.Dir("./Web/assets/img/"))))
	router.PathPrefix("/Web/assets/fonts/").Handler(http.StripPrefix("/Web/assets/fonts/", http.FileServer(http.Dir("./Web/assets/fonts/"))))
	router.PathPrefix("/Web/assets/").Handler(http.StripPrefix("/Web/assets/", http.FileServer(http.Dir("./Web/assets/"))))
	router.PathPrefix("/Web/").Handler(http.StripPrefix("/Web/", http.FileServer(http.Dir("./Web/"))))
}

func routeNotFoundError() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/platform", http.StatusFound)
	})
}
