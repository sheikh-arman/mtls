// client/client.go

package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	/*client := http.Client{
		Timeout: time.Minute * 3,
	}*/
	cert, err := ioutil.ReadFile("/home/user/go/src/github.com/sheikh-arman/mtls/cert/ca.crt")
	if err != nil {
		log.Fatalf("could not open certificate file: %v", err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(cert)

	certificate, err := tls.LoadX509KeyPair("/home/user/go/src/github.com/sheikh-arman/mtls/cert/client.crt", "/home/user/go/src/github.com/sheikh-arman/mtls/cert/client.key")
	if err != nil {
		log.Fatalf("could not load certificate: %v", err)
	}

	client := http.Client{
		Timeout: time.Minute * 3,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{certificate},
			},
		},
	}
	/*	resp, err := client.Get("http://localhost:9090")
		if err != nil {
			log.Fatalf("error making get request: %v", err)
		}*/
	// change the address to match the common name of the certificate
	resp, err := client.Get("https://arman:9090")
	if err != nil {
		log.Fatalf("error making get request: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading response: %v", err)
	}
	fmt.Println(string(body))
}
