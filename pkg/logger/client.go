package logger

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/mehrdadjalili/facegram_common/pkg/logger/pb/pd_logger_client"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"io/ioutil"
	"log"
)

func connectToLoggerServer(url, token, certificate, domain string) (*grpc.ClientConn, pd_logger_client.LoggerServiceClient, error) {
	rpcCreds := oauth.NewOauthAccess(&oauth2.Token{AccessToken: token})
	caCert, err := ioutil.ReadFile(certificate)
	if err != nil {
		log.Fatalln(err)
	}
	rootCAs := x509.NewCertPool()
	rootCAs.AppendCertsFromPEM(caCert)

	tlsConf := &tls.Config{
		RootCAs:            rootCAs,
		InsecureSkipVerify: false,
		MinVersion:         tls.VersionTLS12,
		ServerName:         domain,
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(tlsConf)),
		grpc.WithPerRPCCredentials(rpcCreds),
	}

	connection, err := grpc.Dial(url, opts...)
	if err != nil {
		return nil, nil, err
	}
	return connection, pd_logger_client.NewLoggerServiceClient(connection), nil
}
