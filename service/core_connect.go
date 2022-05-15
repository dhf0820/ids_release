package service

// import (
// 	"crypto/tls"
// 	"crypto/x509"
// 	"fmt"

// 	"github.com/davecgh/go-spew/spew"
// 	pb "github.com/dhf0820/ids_core/protobufs/corePB"
// 	log "github.com/sirupsen/logrus"

// 	//mod "github.com/dhf0820./ds_model"
// 	"io/ioutil"
// 	"os"

// 	common "github.com/dhf0820/ids_model/common"
// 	corePkg "github.com/dhf0820/ids_core/pkg"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials"
// )

// //var baseAddress string
// //var address string
// //var port int

// const (
// 	coreCertFile       = "cert/core-cert.pem"
// 	coreClientCertFile = "cert/core-cert.pem"
// 	coreClientKeyFile  = "cert/core-key.pem"
// )

// var (
// 	coreClient    pb.CoreServiceClient
// 	coreClientPtr *pb.CoreServiceClient
// 	corePort      string
// 	coreAddress   string
// 	coreTlsMode   string
// 	//releasePort 	string
// 	//releaseAddress  string
// 	//releaseTlsMode  string
// )

// //func CurrentReleaseConnection() (relPB.ReleaseServiceClient, error) {
// //	fmt.Printf("\ndeliverCurrentConnection\n")
// //	if relClientPtr == nil {
// //		fmt.Printf("DeliveryConnection does not exist, creating!\n")
// //		return ConnectToRelease()
// //	}
// //	fmt.Printf("DeliveryConnection exists\n")
// //	return relClient, nil
// //}

// //SetReleaseConnection sets the connection information that will be used on a connect or reconnect
// //func SetReleaseConnectMode(mode string) (portNum, baseAddress, tls string){
// //	if strings.ToLower(mode) == "local" {
// //		portNum = "9902"
// //		baseAddress = "localhost"
// //		tls = "none"
// //	} else {
// //		cfg := Config()
// //		portNum = cfg.Port
// //		baseAddress = cfg.Address
// //		tls = cfg.TLSMode
// //	}
// //	return
// //}

// //SetConnection sets the connection information that will be used on a connect or reconnect
// func coreSetConnection(portNum, baseAddress, tls string) {
// 	corePort = portNum
// 	coreAddress = baseAddress
// 	coreTlsMode = tls
// 	fmt.Printf("-----Set connection fields: address: %s, port: %s, tls: %s\n",
// 		coreAddress, corePort, coreTlsMode)
// }

// func CoreGetConnectionDetail() (portNum string, addr string, tls string) {
// 	portNum = corePort
// 	addr = coreAddress
// 	tls = coreTlsMode
// 	return
// }

// //func SetupClient(addr, portNum, tls string) {
// //	address = addr
// //	port = portNum
// //	tlsMode = tls
// //	fmt.Printf("SetupClient: addr: %s, port: %s, tlsMode: %s\n", address, port, tlsMode)
// //
// //}

// // ResetClient returns the current connection information
// func CoreResetClient() {
// 	coreClientPtr = nil
// 	coreClient = nil
// 	ConnectToCore()
// 	return
// }

// func ConnectToCore() (pb.CoreServiceClient, error) {
// 	var transportOptions grpc.DialOption
// 	if coreClientPtr != nil {
// 		fmt.Printf("DeliveryConnection exist, Using It!\n")
// 		return coreClient, nil
// 	}
// 	//release := GetServiceEndPoint("grpc_release")
// 	//fmt.Printf("GrpcReleaseService: %s\n", spew.Sdump(release))
// 	//corePort 	:= os.Getenv("CORE_PORT")
// 	coreAddress := os.Getenv("CONFIG_ADDRESS")
// 	coreTlsMode := os.Getenv("CONFIG_TLSMODE")

// 	//fmt.Printf("\n\nConnect to ReleaseService: address: %s, port: %s, TLSMode: %s\n\n",
// 	//	releaseAddress, releasePort, releaseTlsMode)
// 	//fmt.Printf("connect to ReleaseServer: %s, port: %s, %s\n",
// 	//	releaseAddress, releasePort, releaseTlsMode)
// 	connectAddress := coreAddress //fmt.Sprintf("%s", coreAddress)
// 	//creds, err := credentials.NewClientTLSFromFile("../../server.crt", "")
// 	//useSecure, exists := os.LookupEnv("USESECURE")

// 	switch coreTlsMode {
// 	case "none":
// 		transportOptions = coreNoTLS()
// 	case "server":
// 		transportOptions = coreTLSOnly()
// 	case "mutual":
// 		transportOptions = coreMutualTLS()
// 	default:
// 		transportOptions = coreNoTLS()
// 	}
// 	transportOptions = grpc.WithInsecure()
// 	//Not supporting login/jwt at this time
// 	//fmt.Printf("ReleaseConnect address: %s  --  %s\n", connectAddress, spew.Sdump(transportOptions))
// 	conn, err := grpc.Dial(connectAddress, transportOptions)
// 	//fmt.Printf("Dial return error: %s\n", err)
// 	if err != nil {
// 		return nil, err
// 	}
// 	//defer conn.Close()
// 	//fmt.Printf("Creating Client\n")

// 	coreClient = pb.NewCoreServiceClient(conn)
// 	coreClientPtr = &coreClient
// 	fmt.Printf("Connected to coreservice: %s\n", connectAddress)
// 	return coreClient, nil
// }

// func coreNoTLS() grpc.DialOption {
// 	fmt.Printf("\n\nnoTLS: %s\n\n", spew.Sdump(grpc.WithInsecure()))
// 	return grpc.WithInsecure()
// }

// func coreTLSOnly() grpc.DialOption {
// 	fmt.Println("In serverTLSOnly\n)")
// 	pemServerCA, err := ioutil.ReadFile(coreCertFile)
// 	if err != nil {
// 		log.Fatal()
// 	}
// 	certPool := x509.NewCertPool()
// 	if !certPool.AppendCertsFromPEM(pemServerCA) {
// 		log.Fatal("failed to add server CA'a certificte")
// 	}

// 	// Create the Credentials and return it

// 	config := &tls.Config{
// 		RootCAs: certPool,
// 	}
// 	//fmt.Printf("Config: %s\n", spew.Sdump(config))
// 	//credentials.NewTLS(config)
// 	// Create the Credentials and return it
// 	return grpc.WithTransportCredentials(credentials.NewTLS(config))
// }

// func coreMutualTLS() grpc.DialOption {

// 	// load certificate of the CA who signed servers's certificate
// 	pemServerCA, err := ioutil.ReadFile(coreCertFile)
// 	if err != nil {
// 		log.Fatal("CACertFile ws not found: ", err)
// 	}
// 	certPool := x509.NewCertPool()
// 	if !certPool.AppendCertsFromPEM(pemServerCA) {
// 		log.Fatal("failed to add server CA's certificate")
// 	}
// 	// Load client's certificate and private key
// 	clientCert, err := tls.LoadX509KeyPair(coreClientCertFile, coreClientKeyFile)
// 	if err != nil {
// 		log.Fatal("Load X509 failed: ", err)
// 	}

// 	config := &tls.Config{
// 		Certificates: []tls.Certificate{clientCert},
// 		RootCAs:      certPool,
// 	}
// 	//credentials.NewTLS(config)
// 	// Create the Credentials and return it
// 	return grpc.WithTransportCredentials(credentials.NewTLS(config))
// }

// func toServiceConfig(psc *pb.ServiceConfig) *corePkg.ServiceConfig {
// 	sc := corePkg.ServiceConfig{}
// 	sc.ConnectInfo = toKVData(psc.ConnectInfo)
// 	sc.DataConnector = toDataConnector(psc.DataConnector)
// 	sc.ID, _ = primitive.ObjectIDFromHex(psc.GetId())
// 	sc.MyEndPoints = toEndPoints(psc.MyEndPoints)
// 	sc.Name = psc.Name
// 	sc.Version = psc.Version
// 	sc.ServiceEndPoints = toEndPoints(psc.ServiceEndPoints)
// 	return &sc
// }

// func toKVData(pkvd []*pb.KVData) []*common.KVData {
// 	kvda := []*common.KVData{}
// 	for _, d := range pkvd {
// 		kvd := common.KVData{}
// 		kvd.Name = d.Name
// 		kvd.Value = d.Value
// 		kvda = append(kvda, &kvd)
// 	}
// 	return kvda
// }

// func toDataConnector(pdc *pb.DataConnector) *common.DataConnector {
// 	dc := common.DataConnector{}
// 	dc.Collection = pdc.Collection
// 	dc.Database = pdc.Database
// 	dc.Fields = toKVData(pdc.Fields)
// 	dc.Password = pdc.Password
// 	dc.Server = pdc.Server
// 	return &dc
// }
// func toEndPoints(peps []*pb.EndPoint) []*common.EndPoint {
// 	epa := []*common.EndPoint{}
// 	for _, pep := range peps {
// 		epa = append(epa, toEndPoint(pep))
// 	}
// 	return epa
// }

// func toEndPoint(pep *pb.EndPoint) *common.EndPoint {
// 	ep := common.EndPoint{}
// 	ep.Address = pep.Address
// 	ep.CertName = pep.CertName
// 	ep.Credentials = pep.Credentials
// 	ep.DeployMode = pep.DeployMode
// 	ep.Label = pep.Label
// 	ep.Name = pep.Name
// 	ep.Port = pep.Port
// 	ep.Protocol = pep.Protocol
// 	ep.Scope = pep.Scope
// 	ep.TLSMode = pep.TlsMode
// 	return &ep
// }
