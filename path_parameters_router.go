package examples

import (
	"fmt"
	wagi "github.com/syke99/waggy/v2"
	"net/http"
)

var flg waggy.FullCGI

func routerPathParamsHandler(w http.ResponseWriter, r *http.Request) {
	params := wagi.Vars(r)

	greetingType := params["type"]

	switch greetingType {
	case "hello":
		fmt.Fprintln(w, "Hello World!!")
	case "goodbye":
		fmt.Fprintln(w, "Goodbye for now!!!")
	}
}

func ExampleRouterPathParams() {
	greetingHandler := wagi.NewHandler(nil).
		MethodHandler(http.MethodGet, routerPathParamsHandler)

	router := wagi.NewRouter(&flg)

	router.Handle("/greeting/{type}", greetingHandler)

	_ = wagi.Serve(router)
}
