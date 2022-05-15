package service

// import (
// 	"context"
// 	"fmt"
// 	. "github.com/smartystreets/goconvey/convey"
// 	"github.com/stretchr/testify/require"
// 	relPB "gitlab.com/dhf0820/ids_release_service/protobufs/relPB"
// 	"google.golang.org/grpc"
// 	"os"

// 	//log "github.com/sirupsen/logrus"
// 	"testing"
// )

// func TestGRPCGetDocuments(t *testing.T) {
// 	t.Parallel()
// 	InitTest()
// 	fmt.Printf("\n\nTestGRPCGetDocuments\n")
// 	//os.Setenv("ENV_RELEASE_TEST", "/Users/dhf/work/roi/services/release_service/config/config.json")
// 	os.Setenv("CONFIG_ADDRESS", "localhost:19800")
// 	Initialize("local")
// 	var err error
// 	ctx := context.Background()
// 	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
// 	if err != nil {
// 		t.Fatalf("Failed to dial bufnet: %v", err)
// 	}
// 	defer conn.Close()
// 	//fmt.Printf("Creating Client\n")
// 	client := relPB.NewReleaseServiceClient(conn)
// 	require.NotNil(t, client)

// 	testID := "5fc3e3b9a652f88181ef5711"

// 	Convey("Subject: GetDocuments for a Release", t, func() {

// 		Convey("Given: A Valid release", func() {
// 			req := relPB.GetDocumentsRequest{
// 				RelId: testID,
// 			}
// 			resp, err := client.GetDocuments(ctx, &req)
// 			So(err, ShouldBeNil)
// 			So(resp, ShouldNotBeNil)
// 			docs := ToDocuments(resp.GetDocuments())
// 			So(len(docs), ShouldEqual, 1)
// 		})
// 	})
// }

// func TestGRPCGetRelease(t *testing.T) {
// 	fmt.Println("\n\n\nTestGetGetRelease")
// 	InitTest()

// 	//os.Setenv("ENV_RELEASE_TEST", "/Users/dhf/work/roi/services/release_service/config/config.json")
// 	Convey("GetDocuments for a Release", t, func() {
// 		os.Setenv("CONFIG_ADDRESS", "localhost:19800")
// 		Initialize("local")

// 		ctx := context.Background()
// 		conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
// 		if err != nil {
// 			t.Fatalf("Failed to dial bufnet: %v", err)
// 		}
// 		defer conn.Close()
// 		//fmt.Printf("Creating Client\n")
// 		client := relPB.NewReleaseServiceClient(conn)
// 		So(client, ShouldNotBeNil)
// 		Convey("Given: A Valid release", func() {
// 			testID := "5fc3e3b9a652f88181ef5711"
// 			req := relPB.GetReleaseRequest{
// 				RelId: testID,
// 			}
// 			resp, err := client.GetRelease(ctx, &req)
// 			So(err, ShouldBeNil)
// 			So(resp, ShouldNotBeNil)

// 			rel := ToRelease(resp.GetRelease())
// 			So(rel, ShouldNotBeNil)
// 			So(len(rel.Documents), ShouldEqual, 1)
// 		})
// 		Convey("Get a release using foreignId", func() {
// 			custCode := "demo"
// 			foreignId := "12345"
// 			req := relPB.GetReleaseRequest{
// 				ForeignId:    foreignId,
// 				CustomerCode: custCode,
// 			}
// 			resp, err := client.GetRelease(ctx, &req)
// 			So(err, ShouldBeNil)
// 			So(resp, ShouldNotBeNil)
// 			rel := ToRelease(resp.GetRelease())
// 			So(rel, ShouldNotBeNil)
// 			So(rel.Client.SourceId, ShouldEqual, foreignId)
// 			docReq := relPB.GetDocumentsRequest{}
// 			docReq.RelId = rel.ID.Hex()
// 			docResp, err := client.GetDocuments(ctx, &docReq)
// 			So(err, ShouldBeNil)
// 			So(docResp, ShouldNotBeNil)
// 			docs := ToDocuments(docResp.GetDocuments())
// 			So(len(docs), ShouldEqual, 1)
// 		})
// 	})

// }

// //func TestGRPCGetReleaseForForeignId(t *testing.T) {
// //	fmt.Println("\n\n\nTestGRPCGetReleaseForForeignId")
// //	//os.Setenv("ENV_RELEASE_TEST", "/Users/dhf/work/roi/services/release_service/config/config.json")
// //	InitTest()
// //	Initialize()
// //	ctx := context.Background()
// //	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
// //	if err != nil {
// //		t.Fatalf("Failed to dial bufnet: %v", err)
// //	}
// //	defer conn.Close()
// //	//fmt.Printf("Creating Client\n")
// //	client := relPB.NewReleaseServiceClient(conn)
// //	require.NotNil(t, client)
// //
// //	custCode := int32(101)
// //	foreignId := "12345"
// //
// //	//testID := "60049f4a0e7104875cb6bd29"
// //	req := relPB.GetReleaseRequest{
// //		ForeignId:    foreignId,
// //		CustomerCode:	custCode,
// //	}
// //	resp, err := client.GetRelease(ctx, &req )
// //	fmt.Printf("GetRelease error: %v\n", err)
// //	require.Nil(t, err)
// //	require.NotNil(t,resp.GetRelease())
// //	require.Nil(t, err)
// //	rel := resp.GetRelease()
// //	fmt.Printf("Release: %s\n", rel)
// //	require.Equal(t, req.RelId, rel.RelId)
// //}

// func TestGRPCCreateRelease(t *testing.T) {
// 	fmt.Println("\n\n\nTestCreateRelease")
// 	InitTest()
// 	Initialize("local")

// 	ctx := context.Background()
// 	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
// 	if err != nil {
// 		t.Fatalf("Failed to dial bufnet: %v", err)
// 	}
// 	defer conn.Close()

// 	//fmt.Printf("Creating Client\n")
// 	client := relPB.NewReleaseServiceClient(conn)
// 	require.NotNil(t, client)
// 	rel := NewRelease()
// 	req := relPB.CreateReleaseRequest{}
// 	req.Release = ToPbRelease(rel)
// 	// Setup and create a new Release
// 	resp, err := client.CreateRelease(ctx, &req)
// 	require.Nil(t, err)
// 	require.NotEmpty(t, resp.Release)
// 	require.NotEmpty(t, resp.Release.GetRelId())
// 	//Validate it was added
// 	//rID, _ := primitive.ObjectIDFromHex(resp.Release.GetRelId())
// 	_, err = GetReleaseById(resp.Release.GetRelId())
// 	require.Nil(t, err)

// 	//// Delete the test Release
// 	//oid, _ := primitive.ObjectIDFromHex(resp.Release.GetRelId())
// 	//err = domain.DeleteRelease(oid)
// 	////fmt.Printf("Delete Document failed: %s\n", err)
// 	//require.Nil(t, err)
// }

// //func TestGetReleasesForPatient(t *testing.T){}
// //func TestGetReleasesForForRecipient(t *testing.T){}
// //func TestGetReleasesForDevice(t *testing.T){}

// //func TestGetReleaseByForeignId(t *testing.T) {
// //	fmt.Println("TestGetReleaseByForeignId")
// //	db.Open()
// //	//settings.SetDbName("test_documents")
// //
// //	ctx := context.Background()
// //	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
// //	if err != nil {
// //		t.Fatalf("Failed to dial bufnet: %v", err)
// //	}
// //	defer conn.Close()
// //	//fmt.Printf("Creating Client\n")
// //	client := relPB.NewReleaseServiceClient(conn)
// //	require.NotNil(t, client)
// //	//liveid := "5ee50a5fe8420bcf0d9f68fc"
// //	//testID := "5ee52bce3a072fefccad781e"
// //	sourceId := "47578"
// //	clientId := "Test"
// //	req := relPB.GetReleaseRequest{
// //		SourceId:    sourceId,
// //		ClientId: 	 clientId,
// //	}
// //	resp, err :=client.GetRelease(ctx, &req )
// //	require.Nil(t, err)
// //	require.NotNil(t,resp.GetRelease())
// //	require.Nil(t, err)
// //	rel := resp.GetRelease()
// //	require.Equal(t, req.SourceId, rel.SourceId)
// //}
