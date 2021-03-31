package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func main() {

	cert, _ := tls.LoadX509KeyPair("client.moo-client.pem", "client.moo-client-key.pem")
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}
	client := &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				fmt.Println("address original =", addr)
				if addr == "server.moo:9443" {
					addr = "127.0.0.1:9443"
					fmt.Println("address modified =", addr)
				}
				return dialer.DialContext(ctx, network, addr)
			},
			TLSClientConfig: &tls.Config{
				Certificates:       []tls.Certificate{cert},
				Renegotiation:      tls.RenegotiateOnceAsClient,
				InsecureSkipVerify: true,
			},
		},
	}

	r, err := client.Get("https://server.moo:9443")
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

}
