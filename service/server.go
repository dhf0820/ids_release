package service

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	//corepkg "github.com/dhf0820/ids_core/pkg"
	"net/http"
	"os"
	"sync"

	"github.com/dhf0820/ids_model/common"

	//"github.com/dhf0820/ids_release_service/protobufs/relPB"
	//"google.golang.org/grpc"
	//"google.golang.org/grpc/reflection"
	"net"
	//"os"
	//"strings"
	//"google.golang.org/grpc/credentials"
)

var (
	serverAddress string
	baseAddress   string
	port          string
	tlsMode       string
	deployMode    string
	tempDir       string
	//err error
)

var lis net.Listener

type ReleaseServiceServer struct {
}

func NewReleaseServiceServer() *ReleaseServiceServer {
	return &ReleaseServiceServer{}
}

func Start() {
	run_env := os.Getenv("SERVICE_VERSION")
	company := os.Getenv("COMPANY")
	tempDir = os.Getenv("TEMP_DIR")
	//fmt.Printf("Start Called: [%s]\n", run_env)
	Initialize(run_env, company)
	//service.InitCore(run_env) //TODO: get the env value from flag
	//OpenDB()
	cfg := GetConfig()
	//fmt.Printf("\n---cfg: %s\n", spew.Sdump(cfg))
	//ep := mod.GetMyEndpoint("core")
	eps := common.GetMyEndpoints(cfg)

	var wg sync.WaitGroup
	for _, ep := range eps {
		if ep.Protocol == "http" {
			//fmt.Printf("Setting up restful endpoint %s\n", spew.Sdump(ep))
			defer wg.Done()
			wg.Add(1)

			//fmt.Printf("Starting Restful EndPoint: %s\n", spew.Sdump(ep))
			go restful_worker(&wg, ep)
		}
	}
	fmt.Printf("Wait for restful listeners\n")
	wg.Wait()
	fmt.Printf("Restful Servers stopping\n")
}

func restful_worker(wg *sync.WaitGroup, ep *common.EndPoint) {
	//defer wg.Done()
	//fmt.Printf("Restful EndPoint: %s\n", spew.Sdump(ep))
	//restAddress := restEp.Address
	//restAddress := fmt.Sprintf("%s:%s", "0.0.0.0", ep.Port)
	restAddress := fmt.Sprintf("%s:%s", "0.0.0.0", ep.Port)
	router := NewRouter()
	log.Infof("listening for restful requests at %s", restAddress)
	http.ListenAndServe(restAddress, router)
	// if mainErr != nil {
	// 	log.Errorf("Rest Startup error: %v", mainErr)
	// }
}

func GetTempDir() string {
	return tempDir
}

// func Start() {
// 	var err error
// 	var opts = []grpc.ServerOption{}
// 	var config *mod.ServiceConfig
// 	fmt.Printf("Start Called\n")
// 	config, err = Initialize()
// 	if err != nil {
// 		log.Errorf("Release initialization failed with err: %v", err)
// 		os.Exit(1)
// 	}

// 	fmt.Printf("\n\n\nInitialize Returned: %s\n", spew.Sdump(config))
// 	//settings.SetCollectionName("releases")
// 	//settings.SetDbName(os.Getenv("DATABASE"))
// 	ep := GetMyEndPoint("release")
// 	baseAddress = ep.Address
// 	fmt.Printf("RELEASE_ADDRESS - %s\n", baseAddress)
// 	port := ep.Port

// 	fmt.Printf("RELEASE_PORT - %s\n", port)
// 	deployMode := strings.ToUpper(ep.DeployMode)
// 	fmt.Printf("\nStarting [%s] connection\n\n", deployMode)

// 	fmt.Printf("Port: %s\n", port)

// 	switch deployMode {
// 	case "DOCKER":
// 		serverAddress = fmt.Sprintf("0.0.0.0:%s", port)
// 	case "LOCAL":
// 		serverAddress = fmt.Sprintf("localhost:%s", port)
// 		log.Infof("Using LOCAL mode to address [%s]\n", serverAddress)
// 	case "K8S": // This may change to just the port.
// 		serverAddress = fmt.Sprintf(":%s", port)
// 	}

// 	fmt.Printf("Deploying to %s server  - %s\n", deployMode, serverAddress)
// 	lis, err = net.Listen("tcp", serverAddress)
// 	if err != nil {
// 		log.Fatalf("Listen Failed: %s\n", err)
// 	}
// 	s := grpc.NewServer(opts...)

// 	fmt.Printf("Using releaseServiceServer [%s]\n", serverAddress)
// 	ReleaseServiceServer := NewReleaseServiceServer()
// 	relPB.RegisterReleaseServiceServer(s, ReleaseServiceServer) //&releaseServiceServer{})
// 	reflection.Register(s)
// 	//fmt.Printf("Starting Server port: %s\n", serverAddress)
// 	s.Serve(lis)
// }

//func GetMyEndpoints() []*common.EndPoint {
//	endPoints :=  GetConfig().MyEndPoints
//	fmt.Printf("Release Endpoints: %s", spew.Sdump(endPoints))
//	return endPoints
//}
