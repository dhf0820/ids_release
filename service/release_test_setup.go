package service

import (
	"context"
	"fmt"
	"google.golang.org/grpc/test/bufconn"
	"net"
	//"testing"
	"os"
)

//https://stackoverflow.com/questions/42102496/testing-a-grpc-service
const bufSize = 1024 * 1024

var bufLis *bufconn.Listener

const chunkSize = uint32(1 << 14)

func InitTest() {
	fmt.Printf("\nInitTest\n\n")
	os.Setenv("CONFIG_ADDRESS", "http://docker1.ihids.com:19200/api/rest/v1")
	os.Setenv("SERVICE_NAME", "release")
	os.Setenv("SERVICE_VERSION", "local_test")
	os.Setenv("COMPANY", "demo")
	os.Setenv("DOCUMENT_HOST", "http://localhost:19203/api/rest/v1/")

	_, err := GetServiceConfig("release", "local_test", "demo", "")
	if err != nil {
		fmt.Printf("Cound not GetConfig: %v\n", err)
		os.Exit(1)
	}

	//TODO: Start api for testing
	//bufLis = bufconn.Listen(bufSize)
	//s := grpc.NewServer()
	//rss := &ReleaseServiceServer{}
	//relPB.RegisterReleaseServiceServer(s, rss)
	//go func() {
	//	if err := s.Serve(bufLis); err != nil {
	//		log.Fatalf("Server exited with error: %v", err)
	//	}
	//}()
	//settings.SetDbName("test_release_service")
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return bufLis.Dial()
}

/*func SetupDomainRelease(t *testing.T, create bool) *domain.Release {
	//settings.SetDbName("test_documents")
	data := sample.NewDomainDocument(1)
	//data.ID = primitive.NilObjectID
	// if create {
	// 	fmt.Printf("Creating new document: %s\n", spew.Sdump(data))
	// 	doc, err = domain.AddDocument(data)
	// 	if err != nil {
	// 		t.Fatalf("Error setupDomainDocument Creating: %v", err)
	// 	}
	// } else {
	// 	doc = data
	// }
	return data
}

func SetupPbCreateRelease(t *testing.T) *pb.CreateRelease {
	//settings.SetDbName("test_documents")
	data := sample.NewDocument(1)
	//imageFile := sample.ImageFileName
	//_, err := data.FromDocumentPB(pbDoc)
	// if err != nil {
	// 	err := fmt.Errorf("FromDocumentPB failed: %v", err)
	// 	t.Fatal(err)
	// }
	return data
}*/
