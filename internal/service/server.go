package service

//
//var (
//	serverAddress string
//	baseAddress string
//	port string
//	tlsMode string
//	deployMode string
//	//err error
//)
//
//
//var lis net.Listener
//type ReleaseServiceServer struct {
//}
//
//func NewReleaseServiceServer() *ReleaseServiceServer {
//	return &ReleaseServiceServer{}
//}
//
//func Start() {
//
//	var opts = []grpc.ServerOption{}
//	fmt.Printf("Start Called\n")
//	settings.Initialize("")
//	settings.SetCollectionName("releases")
//	//settings.SetDbName(os.Getenv("DATABASE"))
//	db.Open()
//
//	baseAddress = settings.Conf.ReleaseAddress
//	fmt.Printf("RELEASE_ADDRESS - %s\n", baseAddress)
//	port := settings.Conf.ReleasePort
//	fmt.Printf("RELEASE_PORT - %s\n", port)
//	deployMode := strings.ToUpper(settings.Conf.ReleaseDeployMode)
//	fmt.Printf("\nStarting [%s] connection\n\n", deployMode)
//
//	fmt.Printf("Port: %s\n", port)
//
//	switch deployMode {
//	case "DOCKER":
//		serverAddress = fmt.Sprintf("0.0.0.0:%s", port)
//	case "LOCAL":
//		serverAddress = fmt.Sprintf("localhost:%s", port)
//		log.Infof("Using LOCAL mode to address [%s]\n", serverAddress)
//	case "K8S": // This may change to just the port.
//		serverAddress = fmt.Sprintf(":%s", port)
//	}
//
//	fmt.Printf("Deploying to %s server  - %s\n", deployMode, serverAddress)
//	lis, err = net.Listen("tcp", serverAddress)
//	if err != nil {
//		log.Fatalf("Listen Failed: %s\n", err)
//	}
//	s := grpc.NewServer(opts...)
//
//	fmt.Printf("Using releaseServiceServer [%s]\n", serverAddress)
//	ReleaseServiceServer := NewReleaseServiceServer()
//	relPB.RegisterReleaseServiceServer(s, ReleaseServiceServer) //&releaseServiceServer{})
//	reflection.Register(s)
//	//fmt.Printf("Starting Server port: %s\n", serverAddress)
//	s.Serve(lis)
//}
//
////func noTLS() grpc.DialOption {
////	return grpc.WithInsecure()
////}
//
////func serverTLSOnly() grpc.DialOption {
////	tlsCredentials, err := loadTLSCredentials("server")
////	if err != nil {
////		log.Fatal("cannot load TLS credentials")
////	}
////	return grpc.WithTransportCredentials(tlsCredentials)
////}
//
//// Not sure this is correct.
////func mutualTLS() grpc.DialOption {
////	tlsCredentials, err := loadTLSCredentials("mutual")
////	if err != nil {
////		log.Fatal("cannot load TLS credentials")
////	}
////	return grpc.WithTransportCredentials(tlsCredentials)
////}
//
////func loadTLSCredentials(mode string) (credentials.TransportCredentials, error) {
////	var config *tls.Config
////	// load certificate of the CA who signed client's certificate
////	pemClientCA , err := ioutil.ReadFile(clientCACertFile)
////	if err != nil {
////		return nil, err
////	}
////	certPool := x509.NewCertPool()
////	if !certPool.AppendCertsFromPEM(pemClientCA) {
////		return nil, fmt.Errorf("failed to add seerver CA's certificate")
////	}
////
////	// Load server's certificate and private key
////	serverCert, err := tls.LoadX509KeyPair(serverCertFile, serverKeyFile)
////	if err != nil {
////		return nil, err
////	}
////
////	switch mode {
////	case "server":
////		if mode == "server" {
////			config = &tls.Config{
////				Certificates: []tls.Certificate{serverCert},
////				ClientAuth:   tls.NoClientCert, // not mutual tls
////				//ClientAuth:     tls.RequireAndVerifyClientCert,   // Mutual TLS
////				//ClientCAs:      certPool,							//Mutual TLS
////			}
////		}
////	case "mutual":
////		config = &tls.Config{
////			Certificates: []tls.Certificate{serverCert},
////			ClientAuth:   tls.RequireAndVerifyClientCert,
////			ClientCAs:      certPool,							//Mutual TLS
////		}
////	}
////	return credentials.NewTLS(config), nil
////}
//
//
//
//
//
//
