package main

import (
	"fmt"
	///"github.com/joho/godotenv"
	"net/http"
	"os"
	"time"

	"github.com/dhf0820/ids_model/common"
	log "github.com/sirupsen/logrus"
	"gitlab.com/dhf0820/ids_release_service/service"
	// "strings"
	//"google.golang.org/grpc/credentials"
)

func main() {
	fmt.Printf("Main for release starting: 220314.0 at %s\n", time.Now().UTC().String())

	service.Start()
}

func Start() {
	//var opts = []grpc.ServerOption{}
	run_version := os.Getenv("RELEASE_VERSION")

	fmt.Printf("Start Called: [%s]\n", run_version)

	_, err := service.Initialize("", "") //Database is opened ini initialize
	if err != nil {
		fmt.Printf("Initialize error: %v\n", err)
		os.Exit(1)
	}
	//fmt.Printf("\n---cfg: %s\n", spew.Sdump(config))
	//service.OpenDB()
	cfg := service.GetConfig()
	//fmt.Printf("\n---cfg: %s\n", spew.Sdump(cfg))
	eps := common.GetMyEndpoints(cfg)
	for _, ep := range eps {
		if ep.Protocol == "grpc" {
			fmt.Printf("GRPC is not implemented\n")
			//baseAddress := ep.Address
			//serverAddress := ""
			//fmt.Printf("RELEASE_ADDRESS - %s\n", baseAddress)
			//port := ep.Port
			//
			//fmt.Printf("RELEASE_PORT - %s\n", port)
			//deployMode := ep.DeployMode
			////deployMode := os.Getenv("RELEASE_DEPLOYMODE")
			////deployMode := strings.ToUpper(ep.DeployMode)
			//fmt.Printf("\nStarting [%s] connection\n\n", deployMode)
			//
			//fmt.Printf("Port: %s\n", port)
			//
			//switch deployMode {
			//case "DOCKER":
			//	serverAddress = fmt.Sprintf("0.0.0.0:%s", port)
			//case "LOCAL":
			//	serverAddress := fmt.Sprintf("localhost:%s", port)
			//	log.Infof("Using LOCAL mode to address [%s]\n", serverAddress)
			//case "K8S": // This may change to just the port.
			//	serverAddress = fmt.Sprintf(":%s", port)
			//}
			//
			//fmt.Printf("Deploying to %s server  - %s\n", deployMode, serverAddress)
			//lis, err := net.Listen("tcp", serverAddress)
			//if err != nil {
			//	log.Fatalf("Listen Failed: %s\n", err)
			//}
			//s := grpc.NewServer(opts...)
			//
			//fmt.Printf("Using ReleaseServiceServer [%s]\n", serverAddress)
			//releaseServiceServer := service.NewReleaseServiceServer()
			//relPB.RegisterReleaseServiceServer(s, releaseServiceServer) //&releaseServiceServer{})
			//reflection.Register(s)
			//go s.Serve(lis) // Run as goroutine so main can start the http handler`
		} else if ep.Protocol == "http" {
			//fmt.Printf("Restful EndPoint: %s\n", spew.Sdump(ep))
			restAddress := fmt.Sprintf("%s:%s", ep.Address, ep.Port)
			router := service.NewRouter()
			log.Infof("listening for restful requests at %s", restAddress)
			mainErr := http.ListenAndServe(restAddress, router)
			if mainErr != nil {
				log.Errorf("Rest Startup error: %v", mainErr)
			}
		}
	}
}
