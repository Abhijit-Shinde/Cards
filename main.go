package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func (l logWriter) Write(b []byte) (i int, e error) {
	fmt.Println(string(b))
	return len(b), nil
}

func main() {
	// resp, _ := http.Get("https://dummyjson.com/products/1")

	// var result map[string]interface{}

	// body, _ := io.ReadAll(resp.Body)
	// json.Unmarshal(body, &result)

	// fmt.Println(result["products"])

	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	io.Copy(logWriter{}, resp.Body)
}
