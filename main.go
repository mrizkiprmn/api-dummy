package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	filepath   = os.Getenv("CONFIG_PATH")
	serverport = os.Getenv("SERVER_PORT")
)

func main() {
	mp := make(map[string]interface{})

	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	err = dec.Decode(&mp)
	if err != nil {
		log.Println(err)
	}

	handler := http.NewServeMux()
	for k := range mp {
		fmt.Printf("load: %30v\n", k)
		handler.HandleFunc(k, func(w http.ResponseWriter, r *http.Request) {
			a := mp[r.URL.Path]
			res, _ := json.Marshal(&a)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, string(res))
		})
	}
	http.ListenAndServe(serverport, handler)
}
