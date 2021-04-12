package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	gw "order_client/gw"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	hostname           = "localhost"
	crtFile            = filepath.Join("..", "certs", "client.crt")
	keyFile            = filepath.Join("..", "certs", "client.key")
	caFile             = filepath.Join("..", "certs", "ca.crt")
	grpcServerEndpoint = "localhost:50051"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	certificate, err := tls.LoadX509KeyPair(crtFile, keyFile)
	if err != nil {
		log.Fatalf("could not load client key pair: %s", err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(caFile)

	if err != nil {
		log.Fatalf("could not read ca certificate: %s", err)
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("failed to append ca certs")
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			ServerName:   hostname,
			Certificates: []tls.Certificate{certificate},
			RootCAs:      certPool,
		})),
	}

	err = gw.RegisterProductInfoHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("Fail to registerr gRPC service endpoint: %v", err)
		return
	}

	if err := http.ListenAndServe(":8081", mux); err != nil {
		log.Fatalf("Could not setup HTTP endpoint: %v", err)
	}

}
