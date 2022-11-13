package main

import (
	"fmt"
	"lilurl/handler"
	"net/http"
)

const serverPort = ":8080"

func main() {
	// start the server
	mux := defaultMux()
	// build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/renee": "https://github.com/reneemeyer",
		"/chill": "https://www.youtube.com/watch?v=MkuPV_5kbAg",
	}

	mapHandler := handler.MapHandler(pathsToUrls, mux)

	// build the YAMLHandler using the maphandler as the fallback
	yaml := `
- path: /biscuits
  url: https://bakerbettie.com/5-ingredient-basic-drop-biscuits
- path: /megatron
  url: https://static.wikia.nocookie.net/transformers/images/7/7a/G1_Megatron.png/revision/latest?cb=20210722144704
`
	yamlHandler, err := handler.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting server on: ", serverPort)
	http.ListenAndServe(serverPort, yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hey good lookin")
}
