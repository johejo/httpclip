package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
)

var (
	endpoint string
)

func init() {
	flag.StringVar(&endpoint, "endpoint", "http://localhost:9999", "endpoint")
}

func main() {
	if err := mainE(); err != nil {
		log.Fatal(err)
	}
}

func mainE() error {
	flag.Parse()
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, os.Stdin)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
