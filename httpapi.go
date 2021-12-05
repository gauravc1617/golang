package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"./handler"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	lp := handler.NewProduct(l)

	sm := http.NewServeMux()
	sm.Handle("/", lp)

	ser := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		err := ser.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Printf("Receive terminate gracefully", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	ser.Shutdown(tc)

}
