package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// request, error
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8000", nil) // (context,metodo,url,mux)

	if err != nil {
		panic(err)
	}

	// response, error
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	io.Copy(os.Stdout, res.Body)
}
