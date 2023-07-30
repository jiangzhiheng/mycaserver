package server

import (
	"crypto/tls"
	cx509 "crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	mygrpc "mycaserver/pkg/grpc"
	"net"
)

const (
	port int32 = 8112
)

var server *certificateServiceServer = &certificateServiceServer{}

/*
Run start gRPC server to accept certificate related request
*/
func Run(enableMTLS bool) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var s *grpc.Server
	if enableMTLS {
		tlsCre, err := createTLSCridentials()
		if err != nil {
			return
		}
		s = grpc.NewServer(grpc.Creds(tlsCre))
	} else {
		s = grpc.NewServer()
	}
	mygrpc.RegisterCertificateServiceServer(s, server)
	log.Printf("server listening at %v, gRpc", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func createTLSCridentials() (credentials.TransportCredentials, error) {
	// 加载根证书
	caPEMFile, err := ioutil.ReadFile("cert/rootCA/root.crt") //assume both grpc server and client's certificate are signed by same CA
	if err != nil {
		return nil, err
	}
	// 加入到证书池
	caPool := cx509.NewCertPool()
	if !caPool.AppendCertsFromPEM(caPEMFile) {
		return nil, &ServerError{msg: "load local cert fail"}
	}
	// 加载server端证书
	localCert, err := tls.LoadX509KeyPair("cert/localCert/local.crt", "cert/localCert/local.private.key")
	if err != nil {
		log.Print("load local certificate and key file fail")
		return nil, err
	}
	//返回配置
	config := &tls.Config{
		Certificates: []tls.Certificate{localCert},
		ClientAuth:   tls.RequireAndVerifyClientCert, //means mTLS, will check client's certificate
		ClientCAs:    caPool,
	}
	return credentials.NewTLS(config), nil
}
