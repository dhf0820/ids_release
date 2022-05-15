package connect
//
//import (
//	"crypto/tls"
//	"crypto/x509"
//	"fmt"
//	"os"
//
//	//"github.com/aws/aws-sdk-go/service/mediaconvert"
//
//	log "github.com/sirupsen/logrus"
//	pb "gitlab.com/dhf0820/ids_release_service/protobufs/relPB"
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/credentials"
//	"io/ioutil"
//)
//
////var baseAddress string
////var address string
////var port int
//
//const (
//	releaseCertFile       = "cert/release-cert.pem"
//	releaseClientCertFile = "cert/release-cert.pem"
//	releaseClientKeyFile  = "cert/release-key.pem"
//)
//
//var (
//	releaseClient    pb.ReleaseServiceClient
//	releaseClientPtr *pb.ReleaseServiceClient
//	releasePort      string
//	releaseAddress   string
//	releaseTLSMode   string
//)
//
//// ResetClient returns the current connection information
//func ReleaseResetClient() (pb.ReleaseServiceClient, error) {
//	releaseClientPtr = nil
//	releaseClient = nil
//	return Connect(releaseAddress, releasePort, releaseTLSMode)
//}
//
//func Connect(address string, port, tlsMode string) (pb.ReleaseServiceClient, error) {
//	releaseAddress = address
//	releasePort = port
//	releaseTLSMode = tlsMode
//	connectAddress := fmt.Sprintf("%s:%s", address, port)
//
//	releaseTLSMode = tlsMode
//
//	if releaseClientPtr != nil {
//		fmt.Printf("ReleaseConnection to %s exist, Using It!\n", connectAddress)
//		return releaseClient, nil
//	}
//	var transportOptions grpc.DialOption
//
//	if tlsMode == "" {
//		tlsMode = os.Getenv("RELEASETLS")
//		if tlsMode == "" {
//			tlsMode = "none"
//		}
//	}
//
//	switch releaseTLSMode {
//	case "none":
//		transportOptions = noTLS()
//	case "server":
//		transportOptions = releaseTLSOnly()
//	case "mutual":
//		transportOptions = releaseMutualTLS()
//	default:
//		transportOptions = noTLS()
//	}
//	transportOptions = grpc.WithInsecure()
//	//Not supporting login/jwt at this time
//	fmt.Printf("ReleaseConnect address: [%s]  \n", connectAddress)
//	conn, err := grpc.Dial(connectAddress, transportOptions)
//	//fmt.Printf("Dial return error: %s\n", err)
//	if err != nil {
//		return nil, err
//	}
//	//defer conn.Close()
//	//fmt.Printf("Creating Client\n")
//
//	releaseClient = pb.NewReleaseServiceClient(conn)
//	releaseClientPtr = &releaseClient
//	//fmt.Printf("Connected to ReleaseService: %s\n",connectAddress)
//	return releaseClient, nil
//}
//
//func noTLS() grpc.DialOption {
//	//fmt.Printf("\n\nnoTLS: %s\n\n", spew.Sdump(grpc.WithInsecure()))
//	return grpc.WithInsecure()
//}
//
//func releaseTLSOnly() grpc.DialOption {
//	//fmt.Println("In releaseTLSOnly\n)")
//	pemServerCA, err := ioutil.ReadFile(releaseCertFile)
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
//func releaseMutualTLS() grpc.DialOption {
//
//	// load certificate of the CA who signed servers's certificate
//	pemServerCA, err := ioutil.ReadFile(releaseCertFile)
//	if err != nil {
//		log.Fatal("CACertFile ws not found: ", err)
//	}
//	certPool := x509.NewCertPool()
//	if !certPool.AppendCertsFromPEM(pemServerCA) {
//		log.Fatal("failed to add server CA's certificate")
//	}
//	// Load client's certificate and private key
//	clientCert, err := tls.LoadX509KeyPair(releaseClientCertFile, releaseClientKeyFile)
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
////SetConnection sets the connection information that will be used on a connect or reconnect
//func releaseSetConnection(portNum, baseAddress, tls string) {
//	releasePort = portNum
//	releaseAddress = baseAddress
//	releaseTLSMode = tls
//	// fmt.Printf("-----Set connection fields: address: %s, port: %s, tls: %s\n",
//	// 	releaseAddress, releasePort, releaseTLSMode)
//}
//
//func ReleaseGetConnectionDetail() (portNum string, addr string, tls string) {
//	portNum = releasePort
//	addr = releaseAddress
//	tls = releaseTLSMode
//	return
//}
//
////func toConnectorConfig(pcc *pb.ConnectorConfig) *mod.ConnectorConfig{
////	cc := mod.ConnectorConfig{}
////	cc.ID,_ 				= primitive.ObjectIDFromHex(pcc.GetId())
////	cc.Name 				= pcc.GetName()
////	cc.Version				= pcc.GetVersion()
////	cc.Customer 			= toCustomer(pcc.GetCustomer())
////	cc.ListenAddresses		= toListenAddresses(pcc.GetListenAddresses())
////	cc.EndPoints 			= toEndPoints(pcc.GetEndPoints())
////	cc.Pulls				= pcc.Pulls
////	cc.ConnectAddress		= toConnectAddress(pcc.GetConnectAddress())
////	cc.Fields 				= toFields(pcc.GetFields())
////	cc.Data					= toKVData(pcc.GetData())
////	return &cc
////}
////
////func toPayloadConfig(pcc *pb.PayloadConfig) *mod.PayloadConfig{
////	cc := mod.PayloadConfig{}
////	cc.ID,_ 				= primitive.ObjectIDFromHex(pcc.GetId())
////	cc.Name 				= pcc.GetName()
////	cc.Version				= pcc.GetVersion()
////	cc.Customer 			= toCustomer(pcc.GetCustomer())
////	cc.EndPoints 			= toEndPoints(pcc.GetEndPoints())
////	cc.ConnectAddress 		= toConnectAddress(pcc.GetConnectAddress())
////	cc.Priority				= pcc.GetPriority()
////	cc.Enabled				= pcc.GetEnabled()
////	cc.Fields 				= toFields(pcc.GetFields())
////	cc.Data					= toKVData(pcc.GetData())
////	return &cc
////}
////
////func toConnectAddress(pc *pb.ConnectAddress) *mod.ConnectAddress {
////	ca := mod.ConnectAddress{}
////	ca.Data				= toKVData(pc.GetData())
////	ca.Address			= pc.GetAddress()
////	ca.Authorization	= pc.GetAuthorization()
////	ca.Name				= pc.GetName()
////	ca.Protocol			= pc.GetProtocol()
////	return &ca
////}
////
////func toListenAddresses(pla []*pb.ListenAddress) []*mod.ListenAddress {
////	laa := []*mod.ListenAddress{}
////	for _, pl := range pla {
////		la := mod.ListenAddress{}
////		la.Name 			= pl.GetName()
////		la.Pulls 			= pl.GetPulls()
////		la.Authorization 	= pl.GetAuthorization()
////		la.Address 			= pl.GetAddress()
////		la.Data 			= toKVData(pl.GetData())
////		laa = append(laa, &la)
////	}
////	return laa
////}
////
////func toFields(pfa []*pb.Field) []*mod.Field {
////	flds := []*mod.Field{}
////	for _, pf := range pfa {
////		f := mod.Field{}
////		f.Name				= pf.GetName()
////		f.Value				= pf.GetValue()
////		f.Default			= pf.GetDefault()
////		f.DisplayValue 		= pf.GetDisplayValue()
////		f.Sensitive			= pf.GetSensitive()
////		f.Required			= pf.GetRequired()
////		f.UserVisible		= pf.GetUserVisible()
////		f.IsNameValue		= pf.GetIsNameVisible()
////		flds = append(flds, &f)
////	}
////	return flds
////}
////
////
////func toKVData( pkvd []*pb.KVData) []*mod.KVData {
////	kvda := []*mod.KVData{}
////	for _, d := range pkvd {
////		kvd := mod.KVData{}
////		kvd.Name 	= d.GetName()
////		kvd.Value	= d.GetValue()
////		kvda = append(kvda, &kvd)
////	}
////	return kvda
////}
////
////func toEndPoints(peps []*pb.EndPoint) []*mod.EndPoint {
////	epa := []*mod.EndPoint{}
////	for _, pep := range peps {
////		epa = append(epa, toEndPoint(pep))
////	}
////	return epa
////}
////
////func toEndPoint(pep  *pb.EndPoint) *mod.EndPoint {
////	ep := mod.EndPoint{}
////	ep.Address				= pep.GetAddress()
////	ep.CertName				= pep.GetCertName()
////	ep.Credentials		= pep.GetCredentials()
////	ep.DeployMode			= pep.GetDeployMode()
////	ep.Label					= pep.GetLabel()
////	ep.Name						= pep.GetName()
////	ep.Port						= pep.GetPort()
////	ep.Protocol				= pep.GetProtocol()
////	ep.Scope					= pep.GetScope()
////	ep.QueueName		= pep.GetQueueName()
////	return &ep
////}
////
////func toCustomer(pc *pb.Customer) *mod.Customer {
////	c := mod.Customer{}
////	c.Name			= pc.GetName()
////	c.Code			= pc.GetCode()
////	c.Facility		= pc.GetFacility()
////	return &c
////}
