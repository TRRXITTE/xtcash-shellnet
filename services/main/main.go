package main

import (
    "log"
    "net/http"
    "time"

    "github.com/julienschmidt/httprouter"
)

func main() {
    router := httprouter.New()

    srv := &http.Server{
        Addr:         ":443", // HTTPS port
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  10 * time.Second,
        Handler:      router,
    }

    InitHandlers(router)

    log.Println("Info: Starting Service on HTTPS")

    // Provide the paths to your SSL certificate and key files
    certFile := "/etc/letsencrypt/live/wallet.traaitt.com/fullchain.pem"
    keyFile := "/etc/letsencrypt/live/wallet.traaitt.com/privkey.pem"
    log.Fatal(srv.ListenAndServeTLS(certFile, keyFile))
}
