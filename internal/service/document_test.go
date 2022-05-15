package service

//import (
//	"github.com/davecgh/go-spew/spew"
//	"github.com/dhf0820/ids_release_service/internal/domain"
//	c "github.com/dhf0820/ids_release_service/pkg/common"
//	relPB "github.com/dhf0820/ids_release_service/protobufs/relPB"
//	mod "github.com/dhf0820./ds_model"
//	"go.mongodb.org/mongo-driver/bson/primitive"
//	"testing"
//	"context"
//	"fmt"
//	"os"
//	"time"
//	"github.com/dhf0820/ids_release_service/internal/settings"
//	db "github.com/dhf0820/ids_release_service/internal/db/vsmongo"
//	"google.golang.org/grpc"
//	"github.com/stretchr/testify/require"
//)
//
//func TestAddDocument( t *testing.T) {
//	fmt.Println("\n\n\nTestAddRemoteDocumentForID")
//	os.Setenv("ENV_ROI", "/Users/dhf/work/roi/services/release_service/config")
//	settings.Initialize("")
//	settings.SetDbName("test_release_service")
//	db.Open()
//	ctx := context.Background()
//	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
//	if err != nil {
//		t.Fatalf("Failed to dial bufnet: %v", err)
//	}
//	defer conn.Close()
//	//fmt.Printf("Creating Client\n")
//	client := relPB.NewReleaseServiceClient(conn)
//	require.NotNil(t, client)
//	relId, _ := primitive.ObjectIDFromHex("5fc3e3b9a652f88181ef5711")
//
//	doc := &mod.Document{}
//
//	doc.Id =primitive.NewObjectID().Hex()
//	doc.ImageType = "pdf"
//	doc.ImageRepository = "ca"
//	doc.ImageURL = "/pdf/175"
//	doc.DocumentURL = "/175"
//	doc.Description = "Mamo"
//	doc.DocClass = "Radiology"
//	req := &relPB.AddDocumentRequest{
//		Doc:   	ToRelPbDocument(*doc),
//		RelId:  "5fc3e3b9a652f88181ef5711",
//	}
//
//	startTime := time.Now()
//
//	resp, err := client.AddDocument(ctx, req)
//	fmt.Printf("\n\nTime call GRPC to AddDocument: %f seconds\n", time.Since(startTime).Seconds())
//
//	require.Nil(t, err)
//	fmt.Printf("resp: %s\n", spew.Sdump(resp))
//
//	rel, err := domain.GetReleaseById(relId)
//	require.Nil(t, err)
//	require.NotEmpty(t, rel.Documents, "Document was not added" )
//	require.NotEqual(t, 0, len(rel.Documents))
//	fmt.Printf("Document list: %v\n", rel.Documents)
//
//}
//
//func NewDocument() *domain.Document {
//	oid, _ := primitive.ObjectIDFromHex("5ee6998c0fd93d42fe68393c")
//	return &domain.Document{
//		ID:				oid,
//		ClientID:		"CADIG",
//		Source:			"Cerner",
//		SourceID:		"sc-7601",
//		Facility:		"Test",
//		DocClass:		"Radiology",
//		Description:	"Left Knee",
//		MRN:			"987654"	,
//		DateOfService:   "2019-10-15",
//	}
//}
//
//
//func NewPbDocument() *relPB.Document {
//	//oid, _ := primitive.ObjectIDFromHex("5ee6998c0fd93d42fe68393c")
//	return &relPB.Document{
//		Id:				"5ee6998c0fd93d42fe68393c",
//		ClientId:		"CADIG",
//		Source:			"Cerner",
//		SourceId:		"sc-7601",
//		Facility:		"Test",
//		DocClass:		"Radiology",
//		Description:	"Left Knee",
//		Mrn:			"987654"	,
//		DateOfService:   "2019-10-15",
//	}
//}
