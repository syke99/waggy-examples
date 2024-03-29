package examples

import (
	"fmt"
	wagi "github.com/syke99/waggy/v2"
	"net/http"
)

var flg waggy.FullCGI

func defRespHandler(w http.ResponseWriter, r *http.Request) {
	params := wagi.Vars(r)

	greetingType := params["type"]

	switch greetingType {
	case "hello":
		fmt.Fprintln(w, "Hello World!!")
	case "goodbye":
		fmt.Fprintln(w, "Goodbye for now!!!")
	case "whatup":
		wagi.WriteDefaultResponse(w, r)
	}
}

func ExampleDefaultResponse() {
	greetingHandler := wagi.NewHandlerWithRoute("/greeting/{type}", &flg).
		MethodHandler(http.MethodGet, defRespHandler).
		WithDefaultResponse([]byte("So what's good?"))

	_ = wagi.Serve(greetingHandler)
}
