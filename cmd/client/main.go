package main

import (
	//"context"
	//"flag"
	//"fmt"
	//"github.com/davecgh/go-spew/spew"
	//relSrv "gitlab.com/dhf0820/ids_release_service/connect"
	relPB "gitlab.com/dhf0820/ids_release_service/protobufs/relPB"
	//"os"

	//"google.golang.org/grpc"
	//"time"

	//"github.com/davecgh/go-spew/spew"
	//log "github.com/sirupsen/logrus"
	//"gitlab.com/dhf0820/ids_release_service/internal/service"
)

var baseAddr *string

//var address string
var portNum *string

func main() {
	//var opts = []grpc.DialOption{}
	//portNum = flag.String("port", "9210", "the server port")
	//baseAddr = flag.String("address", "192.168.1.26", "the server address")
	//fmt.Printf("BaseAddress: %s\n", *baseAddr)
	//option := flag.Int("o", 3, "Command to run")
	//tlsMode := flag.String("secure", "none", "Use TLS(none, server, mutual) (default none")
	//flag.Parse()
	//
	//fmt.Printf("Option: %d\n", *option)
	///*
	//
	//	address := fmt.Sprintf("%s:%s", *baseAddr, *portNum)
	//	fmt.Printf("Create Client: %s\n", address)
	//	transportOptions := grpc.WithInsecure()
	//	conn, err := grpc.Dial(address, transportOptions)
	//	//fmt.Printf("Dial return error: %s\n", err)
	//	if err != nil {
	//		log.Fatalf("Dial could not connect to address: %s - %s", address, err)
	//	}
	//	//defer conn.Close()
	//	fmt.Printf("Creating Client\n")
	//	relClient := pb.NewReleaseServiceClient(conn)*/
	//
	//conn, err := relSrv.Connect(*baseAddr, *portNum, *tlsMode)
	//if err != nil {
	//	err = fmt.Errorf("relSvr.Connect failed: %v\n", err)
	//	os.Exit(1)
	//}
	////client.SetReleaseConnection(*portNum, *baseAddr, *tlsMode)
	////
	////client.ReleaseConnect()
	//
	//fmt.Printf("Client created\n")
	//fmt.Printf("Option: %v\n", *option)
	//switch *option {
	///*	case 1:
	//		fmt.Printf("Executing AddDocument\n")
	//		AddDocument(client)
	//	case 2:
	//		fmt.Printf("Executing Update Image ID")
	//		UpdateImageID(client)*/
	//case 3:
	//	fmt.Printf("Executing Get Release\n")
	//	startTime := time.Now()
	//	id := "5ee52bce3a072fefccad781e"
	//	//liveid := "5ee50a5fe8420bcf0d9f68fc"
	//	//testID := "5ee52bce3a072fefccad781e"
	//
	//	req := relPB.GetReleaseRequest{}
	//	req.RelId = id
	//	resp, err := conn.GetRelease(context.Background(), &req)
	//	if err != nil {
	//		log.Errorf("GetRelease returned error: %v", err)
	//	} else {
	//		//fmt.Printf("Release: %s\n", spew.Sdump(resp))
	//	}
	//	fmt.Printf("Elapsed time: %f\n", time.Since(startTime).Seconds())
	//
	//	//time.Sleep(5 * time.Second)
	//
	//	//startTime = time.Now()
	//	//GetRelease(relClient, id)
	//	//GetRelease(relClient, id)
	//	//GetRelease(relClient, id)
	//	//GetRelease(relClient, id)
	//	//GetRelease(relClient, id)
	//	//GetRelease(relClient, id)
	//	//fmt.Printf("Elapsed time: %f\n", time.Since(startTime).Seconds())
	//	//time.Sleep(10 * time.Second)
	//	//startTime = time.Now()
	//	//GetRelease(relClient, id)
	//	//fmt.Printf("Elapsed time: %f\n", time.Since(startTime).Seconds())
	//case 4:
	//	fmt.Printf("Executing Create Release\n")
	//	//pbReq := pb.CreateReleaseRequest{
	//	//	Release:  server.SamplePbRelease(),
	//	//}
	//	startTime := time.Now()
	//	//CreateRelease(relClient, server.SamplePbRelease())
	//	fmt.Printf("Elapsed time: %f\n", time.Since(startTime).Seconds())
	//}
	////imageid: "5e9d23b6881fd0e043111169"
	////docid:
	////AddDocument(client)

}

//func AddDocument(client pb.ReleaseServiceClient) {
//	document := pb.Release{}
//	document.Description = "Test Document"
//	document.ReleaseId = "r1234"
//
//	resp, err := client.CreateDocument(context.Background(),
//		&pb.CreateDocumentRequest{Document: &document})
//	if err != nil {
//		log.Fatalf("CreateDocument error: %v\n", err)
//	}
//	//fmt.Printf("Response: %s\n", spew.Sdump(resp))
//}

//func UpdateImageID(client pb.DocumentServiceClient) {
//	req := pb.UpdateDocumentImageRequest{}
//	req.ImageId = "5e9d23b6881fd0e043111169"
//	req.DocId = "5ea0d398122a7f0e8581f2d3"
//	_, err := client.UpdateDocumentImageID(context.Background(), &req)
//	if err != nil {
//		log.Fatal("UpdateDocumentImageID failed: %s\n", err)
//	}
//}

//func GetRelease(client relPB.ReleaseServiceClient, id string) {
//	req := relPB.GetReleaseRequest{
//		RelID: id,
//	}
//	fmt.Printf("Calling ReleaseService\n")
//	dRel, err := client.GetRelease(context.Background(), &req)
//	if err != nil {
//		log.Fatalf("GetRelease failed: %v\n", err)
//	}
//	fmt.Printf("Release: %s\n", spew.Sdump(dRel))
//}

func CreateRelease(client relPB.ReleaseServiceClient, rel *relPB.Release) {
	//req := relPB.CreateReleaseRequest{
	//	Release: rel,
	//}
	//fmt.Printf("Calling ReleaseService\n")
	//resp, err := client.CreateRelease(context.Background(), &req)
	//if err != nil {
	//	log.Fatalf("CreateRelease failed: %v\n", err)
	//}
	//fmt.Printf("Release Response: %s\n", spew.Sdump(resp))
}
