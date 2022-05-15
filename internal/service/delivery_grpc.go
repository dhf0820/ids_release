package service

//import (
//	"context"
//	"crypto/tls"
//	"crypto/x509"
//	"fmt"
//	mod "github.com/dhf0820./ds_model"
//
//	// "github.com/dhf0820/ca_link_service/config"
//
//	//rec "github.com/dhf0820/ca_link_service/clients/recipient"
//	//m "github.com/dhf0820/ca_link_service/models"
//	//rc "github.com/dhf0820/roi-recipient/client"
//
//	"github.com/davecgh/go-spew/spew"
//	log "github.com/sirupsen/logrus"
//	delPB "github.com/dhf0820/ids_delivery/protobufs/delPB"
//
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/credentials"
//
//	//"google.golang.org/grpc/credentials"
//	"io/ioutil"
//	//mod "github.com/dhf0820./ds_model"
//	// "github.com/davecgh/go-spew/spew"
//	// "google.golang.org/grpc/codes"
//	// "google.golang.org/grpc/status"
//	// "github.com/davecgh/go-spew/spew"
//
//)
//
////var baseAddress string
////var address string
////var port int
//
//const (
//	caCertFile             = "cert/ca-cert.pem"
//	deliveryClientCertFile = "cert/client-cert.pem"
//	deliveryClientKeyFile  = "cert/client-key.pem"
//)
//
//var (
//	delvClient      delPB.DeliveryServiceClient
//	delvClientPtr   *delPB.DeliveryServiceClient
//	err             error
//	deliveryPort    string
//	deliveryAddress string
//	deliveryTlsMode string
//)
//
//func GetDelivery(ctx context.Context, delvId string) (*mod.Delivery, error) {
//	client, err := deliveryConnect()
//	if err != nil {
//		return nil, err
//	}
//	req := delPB.GetDeliveryRequest{}
//	req.Id = delvId
//	resp, err := client.GetDelivery(ctx, &req)
//	if err != nil {
//		return nil, err
//	}
//	delv := resp.GetDelivery()
//	return ToDelivery(delv), nil
//}
//
//
//func deliveryConnect() (delPB.DeliveryServiceClient, error) {
//	fmt.Printf("IndeliveryConnect\n\n")
//	if delvClientPtr != nil {
//		fmt.Printf("DelvClientPtr exists, use it\n")
//		return delvClient, nil
//	}
//
//	delivery, err := GetService("delivery_service")
//	fmt.Printf("DeliveryService: %s\n", spew.Sdump(delivery))
//	tlsFld, _ := delivery.GetField("tls_mode")
//	fmt.Printf("tlsMode: %s\n", spew.Sdump(tlsFld))
//
//	//delivery, err := GetService("delivery_service")
//	deliveryPort = delivery.Port
//	deliveryAddress = delivery.Address
//	deliveryTlsMode = tlsFld.Value
//	fmt.Printf("\n\nConnect to DeliveryService: %s, %s, TLSMode: %s\n\n", deliveryAddress, deliveryPort, deliveryTlsMode)
//	delvClientPtr = connect()
//	if err != nil {
//		err = fmt.Errorf("Could not Connect to Delivery Service: %v\n", err)
//		return nil, err
//	}
//	delvClient = *delvClientPtr
//	return delvClient, nil
//}
//
//func deliveryCurrentConnection() (delPB.DeliveryServiceClient, error) {
//	fmt.Printf("\ndeliverCurrentConnection\n")
//	if delvClientPtr == nil {
//		fmt.Printf("DeliveryConnection does not exist, creating!\n")
//		return deliveryConnect()
//	}
//	fmt.Printf("DeliveryConnection exists\n")
//	return delvClient, nil
//}
//
//func deliveryGetConnection() (delPB.DeliveryServiceClient, error) {
//	fmt.Printf("\nIn deliveryGetConnectionn\n")
//	if delvClientPtr == nil {
//		fmt.Printf("DeliveryConnection does not exist, creating!\n")
//		return deliveryConnect()
//	}
//	fmt.Printf("DeliveryConnection exists\n")
//	return delvClient, nil
//}
//
////SetReleaseConnection sets the connection information that will be used on a connect or reconnect
////func SetReleaseConnectMode(mode string) (portNum, baseAddress, tls string){
////	if strings.ToLower(mode) == "local" {
////		portNum = "9902"
////		baseAddress = "localhost"
////		tls = "none"
////	} else {
////		cfg := Config()
////		portNum = cfg.Port
////		baseAddress = cfg.Address
////		tls = cfg.TLSMode
////	}
////	return
////}
//
////SetConnection sets the connection information that will be used on a connect or reconnect
//func deliverySetReleaseConnection(portNum, baseAddress, tls string) {
//	deliveryPort = portNum
//	deliveryAddress = baseAddress
//	deliveryTlsMode = tls
//	fmt.Printf("-----Set connection fields: address: %s, port: %s, tls: %s\n",
//		deliveryAddress, deliveryPort, deliveryTlsMode)
//}
//
//func deliveryGetConnectionDetail() (portNum string, addr string, tls string) {
//	portNum = deliveryPort
//	addr = deliveryAddress
//	tls = deliveryTlsMode
//	return
//}
//
////func SetupClient(addr, portNum, tls string) {
////	address = addr
////	port = portNum
////	tlsMode = tls
////	fmt.Printf("SetupClient: addr: %s, port: %s, tlsMode: %s\n", address, port, tlsMode)
////
////}
//
//// ResetClient returns the current connection information
//func deliveryResetClient() {
//	delvClientPtr = nil
//	//DelvClient = nil
//	deliveryConnect()
//	return
//}
//
//func connect() *delPB.DeliveryServiceClient {
//	fmt.Printf("\nIn connect\n")
//	//var opts = []grpc.DialOption{}
//	//deliveryTlsMode := "none"
//	fmt.Printf("connect to DeliveryServer: %s, %s, %s\n", deliveryAddress, deliveryPort, deliveryTlsMode)
//	var transportOptions grpc.DialOption
//	connectAddress := fmt.Sprintf("%s:%s", deliveryAddress, deliveryPort)
//	//creds, err := credentials.NewClientTLSFromFile("../../server.crt", "")
//	//useSecure, exists := os.LookupEnv("USESECURE")
//
//	switch deliveryTlsMode {
//	case "none":
//		transportOptions = noTLS()
//	case "server":
//		transportOptions = serverTLSOnly()
//	case "mutual":
//		transportOptions = mutualTLS()
//
//	}
//
//	//Not supporting login/jwt at this time
//	fmt.Printf("DeliveryConnect address: %s  --  %s\n", connectAddress, spew.Sdump(transportOptions))
//	conn, err := grpc.Dial(connectAddress, transportOptions)
//	//fmt.Printf("Dial return error: %s\n", err)
//	if err != nil {
//		log.Fatalf("Could not connect Dial error: %v\n", err)
//	}
//	//defer conn.Close()
//	fmt.Printf("Creating Client\n")
//
//	clnt := delPB.NewDeliveryServiceClient(conn)
//	return &clnt
//}
//
//func noTLS() grpc.DialOption {
//	fmt.Printf("\n\nnoTLS: %s\n\n", spew.Sdump(grpc.WithInsecure()))
//	return grpc.WithInsecure()
//}
//
//func serverTLSOnly() grpc.DialOption {
//	fmt.Println("In serverTLSOnly\n)")
//	pemServerCA, err := ioutil.ReadFile(caCertFile)
//	if err != nil {
//		log.Fatal()
//	}
//	certPool := x509.NewCertPool()
//	if !certPool.AppendCertsFromPEM(pemServerCA) {
//		log.Fatal("failed to add server CA'a certificte")
//	}
//
//	// Create the Credentials and return it
//
//	config := &tls.Config{
//		RootCAs: certPool,
//	}
//	//fmt.Printf("Config: %s\n", spew.Sdump(config))
//	//credentials.NewTLS(config)
//	// Create the Credentials and return it
//	return grpc.WithTransportCredentials(credentials.NewTLS(config))
//}
//
//func mutualTLS() grpc.DialOption {
//
//	// load certificate of the CA who signed servers's certificate
//	pemServerCA, err := ioutil.ReadFile(caCertFile)
//	if err != nil {
//		log.Fatal("CACertFile ws not found: ", err)
//	}
//	certPool := x509.NewCertPool()
//	if !certPool.AppendCertsFromPEM(pemServerCA) {
//		log.Fatal("failed to add server CA's certificate")
//	}
//	// Load client's certificate and private key
//	clientCert, err := tls.LoadX509KeyPair(deliveryClientCertFile, deliveryClientKeyFile)
//	if err != nil {
//		log.Fatal("Load X509 failed: ", err)
//	}
//
//	config := &tls.Config{
//		Certificates: []tls.Certificate{clientCert},
//		RootCAs:      certPool,
//	}
//	//credentials.NewTLS(config)
//	// Create the Credentials and return it
//	return grpc.WithTransportCredentials(credentials.NewTLS(config))
//}
//
//
////func DeliverRelease(ctx context.Context, dt *m.DeliverTask) (string,error) {
////	if dt.Patient.RemoteID == "" {
////		dt.Patient.RemoteID = dt.Patient.SSN
////	}
////	//fmt.Printf("DeliverRelease device: %s\n", spew.Sdump(dt.Device))
////	resp, err := dc.CreateDelivery(ctx, dt.ReleaseID, dt.Recipient.ID,
////		dt.Device.Id.Hex(), dt.Patient.RemoteID )
////	if err != nil {
////		return "", err
////	}
////	return resp.DeliveryID, nil
////}
//
//
