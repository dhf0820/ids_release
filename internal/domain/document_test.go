package domain

//
//import (
//	"context"
//	"fmt"
//	db "github.com/dhf0820/ids_release_service/internal/db/vsmongo"
//	"github.com/dhf0820/ids_release_service/internal/settings"
//	m "github.com/dhf0820./ds_model"
//	"go.mongodb.org/mongo-driver/bson/primitive"
//
//	. "github.com/smartystreets/goconvey/convey"
//	//c "github.com/dhf0820/ids_release_service/pkg/common"
//	//"go.mongodb.org/mongo-driver/bson/primitive"
//	"os"
//	"testing"
//)
//
//func TestAddRemoteDocument(t *testing.T) {
//	fmt.Printf("\n\nTestAddDocument (Remote) to Release\n\n")
//	Convey("Subject: Add a Remote Document to a Release", t, func() {
//		os.Setenv("ENV_ROI", "/Users/dhf/work/roi/services/release_service/config")
//		settings.Initialize("")
//		db.Open()
//		Convey("Given: A valid remote document reference", func() {
//			doc := m.Document{}
//			doc.ID =primitive.NewObjectID().Hex()
//			doc.ImageType = "pdf"
//			doc.ImageRepository = "ca"
//			doc.ImageURL = "/pdf/175"
//			doc.DocumentURL = "/175"
//			doc.Description = "Mamo"
//			doc.DocClass = "Radiology"
//			releaseID,_ 	:= primitive.ObjectIDFromHex("5fc3e3b9a652f88181ef5711")
//			Convey("Add the document", func() {
//				ctx := context.Background()
//				newDoc, err := AddRemoteDocumentForID(ctx, releaseID, &doc)
//				So(err, ShouldBeNil)
//				So(newDoc, ShouldNotBeNil)
//				//fmt.Printf("Device: %s\n", spew.Sdump(newDoc))
//			})
//		})
//	})
//}
//
////func TestAddDocument(t *testing.T) {
//	//	t.Parallel()
//	//	os.Setenv("ENV_ROI", "/Users/dhf/work/roi/delivery_service/client")
//	//	settings.Initialize("")
//	//	settings.SetDbName("test_release_service")
//	//	db.Open()
//	//	fmt.Printf("\n\nTestAddDocument\n")
//	//	doc := NewDocument()
//	//	pbDoc, err := MarshalDocument(doc)
//	//	fmt.Printf("pbDoc: %s\n", spew.Sdump(pbDoc))
//	//	cnt, err := AddDocument(context.Background(), pbDoc)
//	//	require.Nil(t, err)
//	//	//require.NotEmpty(t, rel.Documents)
//	//	require.NotEqual(t,0, cnt)
//	//	fmt.Printf("total documents: %d\n", cnt)
//	//	//rel.Documents = nil
//	//	//rel.Update()
//	//
//	//
//	//	cnt, err = AddDocument(context.Background(), pbDoc)
//	//	require.Nil(t, err)
//	//	require.NotEqual(t, 0, cnt)
//	//	//fmt.Printf("\n\nTime to directly add Document: %f seconds\n\n", time.Since(startTime).Seconds())
//	//
//	//	//rel, err := GetReleaseById(relID)
//	//	//require.Nil(t, err)
//	//	//require.NotEmpty(t,rel.Documents, "Nothing in Documents")
//	//	////fmt.Printf("Updated Release: %s\n", spew.Sdump(rel))
//	//	//rel.Documents = nil
//	//	//rel.Update()
//	//	//require.NotEqual(t, rel.CreatedAt, rel.UpdatedAt)
//	//
//	//
////}
////
////
////func NewDocument() *Document {
////	ds := Document{
////		ID:          primitive.NewObjectID(),
////		ReleaseID:   "5f397b9d3bd90a0b2f093593",
////		Source:      "Cerner",
////		SourceID:    "abcde",
////		Facility:    "Test",
////		DocClass:    "AUTH",
////		Description: "Document description",
////		MRN:         RandomMRN(),
////		DateOfService: "2019-10-15",
////	}
////	//fmt.Printf("NewDoc: %s\n", spew.Sdump(ds))
////	return &ds
////}
//
//
//
