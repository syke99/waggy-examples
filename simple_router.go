package examples

import (
	"fmt"
	wagi "github.com/syke99/waggy/v2"
	"net/http"
)

var flg waggy.FullCGI

func routerHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}

func routerGoodbye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Goodbye")
}

func ExampleRouter() {
	greetingHandler := wagi.NewHandler(nil).
		MethodHandler(http.MethodGet, routerHello).
		MethodHandler(http.MethodDelete, routerGoodbye)

	router := wagi.NewRouter(&flg)

	router.Handle("/greeting", greetingHandler)

	_ = wagi.Serve(router)
}
