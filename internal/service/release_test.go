package service

//import (
//	"context"
//	"fmt"
//	db "gitlab.com/dhf0820/ids_release_service/internal/db/vsmongo"
//	"gitlab.com/dhf0820/ids_release_service/internal/domain"
//	"gitlab.com/dhf0820/ids_release_service/internal/settings"
//	"go.mongodb.org/mongo-driver/bson/primitive"
//	"os"
//
//	//log "github.com/sirupsen/logrus"
//	"testing"
//	relPB "gitlab.com/dhf0820/ids_release_service/protobufs/relPB"
//	"github.com/stretchr/testify/require"
//
//	"google.golang.org/grpc"
//)
//
//func TestGetReleaseForId(t *testing.T) {
//	fmt.Println("\n\n\nTestGetReleaseForId")
//	os.Setenv("ENV_ROI", "/Users/dhf/work/roi/services/release_service/config")
//	settings.Initialize("")
//	settings.SetDbName("test_release_service")
//	db.Open()
//
//	ctx := context.Background()
//	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
//	if err != nil {
//		t.Fatalf("Failed to dial bufnet: %v", err)
//	}
//	defer conn.Close()
//	//fmt.Printf("Creating Client\n")
//	client := relPB.NewReleaseServiceClient(conn)
//	require.NotNil(t, client)
//
//
//
//	//liveid := "5ee50a5fe8420bcf0d9f68fc"
//	//testID := "5ee52bce3a072fefccad781e"
//	testID := "5fc3e3b9a652f88181ef5711"
//	req := relPB.GetReleaseRequest{
//		RelId:    testID,
//	}
//	resp, err := client.GetRelease(ctx, &req )
//	fmt.Printf("GetRelease error: %v\n", err)
//	require.Nil(t, err)
//	require.NotNil(t,resp.GetRelease())
//	require.Nil(t, err)
//	rel := resp.GetRelease()
//	fmt.Printf("Release: %s\n", rel)
//	require.Equal(t, req.RelId, rel.RelId)
//}
//
//func TestCreateRelease(t *testing.T) {
//	fmt.Println("\n\n\nTestCreateRelease")
//	os.Setenv("ENV_ROI", "/Users/dhf/work/roi/services/release_service/config")
//	settings.Initialize("")
//	settings.SetDbName("test_release_service")
//
//	db.Open()
//	ctx := context.Background()
//	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
//	if err != nil {
//		t.Fatalf("Failed to dial bufnet: %v", err)
//	}
//	defer conn.Close()
//
//	//fmt.Printf("Creating Client\n")
//	client := relPB.NewReleaseServiceClient(conn)
//	require.NotNil(t, client)
//	rel := NewRelease()
//	req := relPB.CreateReleaseRequest{}
//	req.Release = ToPbRelease(rel)
//	// Setup and create a new Release
//	resp, err :=client.CreateRelease(ctx, &req )
//	require.Nil(t, err)
//	require.NotEmpty(t,resp.Release)
//	require.NotEmpty(t, resp.Release.GetRelId())
//	//Validate it was added
//	rID, _ := primitive.ObjectIDFromHex(resp.Release.GetRelId())
//	_, err = domain.GetReleaseById(rID)
//	require.Nil(t, err)
//
//	//// Delete the test Release
//	//oid, _ := primitive.ObjectIDFromHex(resp.Release.GetRelId())
//	//err = domain.DeleteRelease(oid)
//	////fmt.Printf("Delete Document failed: %s\n", err)
//	//require.Nil(t, err)
//}

//func TestGetReleaseByForeignId(t *testing.T) {
//	fmt.Println("TestGetReleaseByForeignId")
//	db.Open()
//	//settings.SetDbName("test_documents")
//
//	ctx := context.Background()
//	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
//	if err != nil {
//		t.Fatalf("Failed to dial bufnet: %v", err)
//	}
//	defer conn.Close()
//	//fmt.Printf("Creating Client\n")
//	client := relPB.NewReleaseServiceClient(conn)
//	require.NotNil(t, client)
//	//liveid := "5ee50a5fe8420bcf0d9f68fc"
//	//testID := "5ee52bce3a072fefccad781e"
//	sourceId := "47578"
//	clientId := "Test"
//	req := relPB.GetReleaseRequest{
//		SourceId:    sourceId,
//		ClientId: 	 clientId,
//	}
//	resp, err :=client.GetRelease(ctx, &req )
//	require.Nil(t, err)
//	require.NotNil(t,resp.GetRelease())
//	require.Nil(t, err)
//	rel := resp.GetRelease()
//	require.Equal(t, req.SourceId, rel.SourceId)
//}
