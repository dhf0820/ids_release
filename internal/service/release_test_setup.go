package service

//import (
//	"context"
//	relPB "github.com/dhf0820/ids_release_service/protobufs/relPB"
//	"net"
//	//"testing"
//
//	log "github.com/sirupsen/logrus"
//
//	//sample "github.com/dhf0820/ids_release_service/src/sample"
//	"github.com/dhf0820/ids_release_service/internal/settings"
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/test/bufconn"
//)
//
////https://stackoverflow.com/questions/42102496/testing-a-grpc-service
//const bufSize = 1024 * 1024
//
//var bufLis *bufconn.Listener
//
//const chunkSize = uint32(1 << 14)
//
////
//func init() {
//
//	bufLis = bufconn.Listen(bufSize)
//	s := grpc.NewServer()
//	rss := &ReleaseServiceServer{}
//	relPB.RegisterReleaseServiceServer(s, rss)
//	go func() {
//		if err := s.Serve(bufLis); err != nil {
//			log.Fatalf("Server exited with error: %v", err)
//		}
//	}()
//	settings.SetDbName("test_release_service")
//}
//
//func bufDialer(context.Context, string) (net.Conn, error) {
//	return bufLis.Dial()
//}
//
///*func SetupDomainRelease(t *testing.T, create bool) *domain.Release {
//	//settings.SetDbName("test_documents")
//	data := sample.NewDomainDocument(1)
//	//data.ID = primitive.NilObjectID
//	// if create {
//	// 	fmt.Printf("Creating new document: %s\n", spew.Sdump(data))
//	// 	doc, err = domain.AddDocument(data)
//	// 	if err != nil {
//	// 		t.Fatalf("Error setupDomainDocument Creating: %v", err)
//	// 	}
//	// } else {
//	// 	doc = data
//	// }
//	return data
//}
//
//func SetupPbCreateRelease(t *testing.T) *pb.CreateRelease {
//	//settings.SetDbName("test_documents")
//	data := sample.NewDocument(1)
//	//imageFile := sample.ImageFileName
//	//_, err := data.FromDocumentPB(pbDoc)
//	// if err != nil {
//	// 	err := fmt.Errorf("FromDocumentPB failed: %v", err)
//	// 	t.Fatal(err)
//	// }
//	return data
//}*/
//
