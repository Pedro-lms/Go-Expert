package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {

	c := http.Client{Timeout: time.Second}
	jsonvar := bytes.NewBuffer([]byte(`{"name":"Pedro"}`))
	resp, err := c.Post("http://google.com", "application/json", jsonvar)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
