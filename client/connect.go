package client

/* import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	relPB "gitlab.com/dhf0820/ids_release_service/protobufs/relPB"
	"strings"
	//cc "gitlab.com/dhf0820/roi-core/client"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
)


//var baseAddress string
//var address string
//var port int

const (
	CACertFile		= "cert/ca-cert.pem"
	clientCertFile 	= "cert/client-cert.pem"
	clientKeyFile  	= "cert/client-key.pem"

)
var (
	RelClient    relPB.ReleaseServiceClient
	RelClientPtr *relPB.ReleaseServiceClient
	err          error
	port         string
	address      string
	tlsMode      string
)

//func init() {

//err = Connect()
//if err != nil {
//	fmt.Printf("settings init failed to connect: %v\n", err)
//}
//}

func ReleaseConnect() (*relPB.ReleaseServiceClient, error) {
	if RelClientPtr != nil {
		return RelClientPtr, nil
	}

	//var transportOptions grpc.DialOption
	if ReleaseConf == nil{
		fmt.Printf("\nViper is not configured\n")
		RelClientSettingsInit()
	}
	port = ReleaseConf.ReleasePort
	address = ReleaseConf.ReleaseAddress
	tlsMode = ReleaseConf.ReleaseTLSMode
	fmt.Printf("\n\nConnect to release: %s, %s, TLSMode: %s\n\n", address, port, tlsMode)
	RelClientPtr = connect(port, address, tlsMode)   //Connect(port, address, tlsMode)
	if err != nil {
		fmt.Printf("Cound not Connect to release service: %v\n", err)
		return nil, err
	}
	RelClient = *RelClientPtr
	return RelClientPtr, nil
}

func CurrentConnection() (*relPB.ReleaseServiceClient, error) {
	if RelClientPtr == nil {
		return ReleaseConnect()
	}
	return RelClientPtr, nil
}

//func junk() {
//	cc.Connect()
//}


//func ReleaseConnect() (*relPB.ReleaseServiceClient, error) {
//	if RelClientPtr != nil {
//		return RelClientPtr, nil
//	}
//	fmt.Printf("\n\nConnect to release: %s, %s, TLSMode: %s\n\n", address, port, tlsMode)
//	RelClientPtr = connect(port, address, tlsMode)   //Connect(port, address, tlsMode)
//	if err != nil {
//		fmt.Printf("Cound not Connect to release service: %v\n", err)
//		return nil, err
//	}
//
//	RelClient = *RelClientPtr
//
//
//	//fmt.Printf("relClient: %s\n", spew.Sdump(relClient))
//	return RelClientPtr, nil
//}
//
//func CurrentConnection() (*relPB.ReleaseServiceClient, error) {
//	if RelClientPtr == nil {
//		return ReleaseConnect()
//	}
//	return RelClientPtr, nil
//}

//SetReleaseConnection sets the connection information that will be used on a connect or reconnect
func SetReleaseConnectMode(mode string) (portNum, baseAddress, tls string){
	if strings.ToLower(mode) == "local" {
		portNum = "9901"
		baseAddress = "localhost"
		tls = "none"
	} else {
		cfg := Config()
		portNum = cfg.ReleasePort
		baseAddress = cfg.ReleaseAddress
		tls = cfg.ReleaseTLSMode
	}
	return
}


//SetConnection sets the connection information that will be used on a connect or reconnect
func SetReleaseConnection(portNum, baseAddress, tls string) {
	fmt.Printf("SetConnection received: port: %s, baseAddr: %s, tls: %s\n", portNum, baseAddress, tls)
	port = portNum
	address = baseAddress
	tlsMode = tls
	fmt.Printf("Set connection fields: address: %s, port: %s, tls: %s\n", address, port,  tlsMode)
}

func GetConnection() (portNum string, addr string, tls string) {

	portNum = port
	addr = address
	tls = tlsMode
	return
}

// ResetClient returns the current connection information
func ResetClient() () {
	RelClientPtr = nil
	RelClient = nil
	ReleaseConnect()
	return
}
func connect(port, baseAddress, tls string) *relPB.ReleaseServiceClient {
	//var opts = []grpc.DialOption{}
	fmt.Printf("connect to ReleaseServer: %s, %s, %s\n", baseAddress, port, tls)
	var transportOptions  grpc.DialOption
	connectAddress := fmt.Sprintf("%s:%s", baseAddress, port)
	//creds, err := credentials.NewClientTLSFromFile("../../server.crt", "")
	//useSecure, exists := os.LookupEnv("USESECURE")

	switch tls {
	case "none":
		transportOptions = noTLS()
	case "server":
		transportOptions = serverTLSOnly()
	case "mutual":
		transportOptions = mutualTLS()
	default:
		transportOptions = noTLS()
	}

	//Not supporting login/jwt at this time
	conn, err := grpc.Dial(connectAddress, transportOptions)
	//fmt.Printf("Dial return error: %s\n", err)
	if err != nil {
		log.Fatalf("Could not connect Dial error: %v\n", err)
	}
	//defer conn.Close()
	fmt.Printf("Creating ReleaseClient\n")

	clnt :=  relPB.NewReleaseServiceClient(conn)
	fmt.Printf("Returning new connection client\n")
	return &clnt
}


func noTLS() grpc.DialOption {
	return grpc.WithInsecure()
}


func serverTLSOnly() grpc.DialOption {
	fmt.Println("In serverTLSOnly\n)")
	pemServerCA , err := ioutil.ReadFile(CACertFile)
	if err != nil {
		log.Fatal()
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		log.Fatal("failed to add server CA'a certificte")
	}

	// Create the Credentials and return it

	config:= &tls.Config{
		RootCAs: certPool,
	}
	//fmt.Printf("Config: %s\n", spew.Sdump(config))
	//credentials.NewTLS(config)
	// Create the Credentials and return it
	return grpc.WithTransportCredentials(credentials.NewTLS(config))
}


func mutualTLS() grpc.DialOption {
	// load certificate of the CA who signed servers's certificate
	pemServerCA , err := ioutil.ReadFile(CACertFile)
	if err != nil {
		log.Fatal("CACertFile ws not found: ", err)
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		log.Fatal("failed to add server CA's certificate")
	}
	// Load client's certificate and private key
	clientCert, err := tls.LoadX509KeyPair(clientCertFile, clientKeyFile)
	if err != nil {
		log.Fatal("Load X509 failed: ", err)
	}

	config := &tls.Config{
		Certificates: 	[]tls.Certificate{clientCert},
		RootCAs:			certPool,
	}
	//credentials.NewTLS(config)
	// Create the Credentials and return it
	return grpc.WithTransportCredentials(credentials.NewTLS(config))
}
 */

