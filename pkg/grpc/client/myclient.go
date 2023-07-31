package client

import (
	"context"
	"crypto/tls"
	cx509 "crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/types/known/emptypb"
	"io/ioutil"
	"log"
	mygrpc "mycaserver/pkg/grpc"
	"time"
)

const (
	port int32 = 8112
)

func createTLSCredentials(clientId string) (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	rootCAFile, err := ioutil.ReadFile("cert/rootCA/root.crt")
	if err != nil {
		return nil, err
	}

	certPool := cx509.NewCertPool()
	if !certPool.AppendCertsFromPEM(rootCAFile) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Load client's certificate and private key
	clientCert, err := tls.LoadX509KeyPair("cert/clientCert/"+clientId+".crt", "cert/clientCert/"+clientId+".key")
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		ServerName:   "localhost", //server 必须是对应的这个值
		RootCAs:      certPool,
	}

	return credentials.NewTLS(config), nil
}

func Run(clientId string) {
	tlsCredentials, err := createTLSCredentials(clientId)
	if err != nil {
		log.Print("cannot load TLS credentials: ", err)
		return
	}
	dial, err := grpc.Dial(fmt.Sprintf(":%v", port), grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		log.Print("cannot dial server: ", err)
		return
	}
	defer dial.Close()

	rpcClient := mygrpc.NewCertificateServiceClient(dial)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	template, err := rpcClient.CsrTemplate(ctx, &emptypb.Empty{})
	if err != nil {
		log.Print("error happen when call gRPC client:" + err.Error())
		return
	}

	fmt.Print(template)
}
