package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func moo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	_, err := fmt.Fprint(w, "Oh lols hi")
	if err != nil {
		return
	}

}

func main() {
	http.HandleFunc("/", moo)

	caCert, _ := ioutil.ReadFile("../rootCA.pem")
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		ClientCAs:  caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}
	tlsConfig.BuildNameToCertificate()
	server := &http.Server{
		Addr:      ":9443",
		TLSConfig: tlsConfig,
	}

	err := server.ListenAndServeTLS("server.moo.pem", "server.moo-key.pem")

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
