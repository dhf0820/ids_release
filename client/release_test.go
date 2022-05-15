package client

// import (
// 	//"context"
// 	"fmt"
// 	//"github.com/davecgh/go-spew/spew"

// 	//"github.com/davecgh/go-spew/spew"
// 	//"github.com/dhf0820/ids_release_service/protobufs/relPB"

// 	"github.com/dhf0820/ids_release_service/service"
// 	"os"
// 	"testing"
// 	. "github.com/smartystreets/goconvey/convey"
// 	//"github.com/stretchr/testify/require"
// )

// func TestGetRelease(t *testing.T) {
// 	Convey("Subject: Get a Release", t, func() {
// 		fmt.Printf("\nTestGet a release\n")
// 		os.Setenv("ENV_ROI", "/Users/dhf/work/roi/services/release_service/config")
// 		os.Setenv("CA_LINK", "/Users/dhf/work/roi/services/ca_link_service/config/configuration.test.json")

// 		service.InitTest()
// 		service.Initialize("local")
// 		//OpenDB

// 		Convey("Get the release for  5fc3e3b9a652f88181ef5711", func() {
// 		// 	ctx := context.Background()
// 		// 	relId := "5fc3e3b9a652f88181ef5711"
// 			Convey("OpenDB ", func() {
// 				mongo := service.OpenDB()
// 				So(mongo, ShouldNotBeNil)
// 			})
// 			// Convey("As Undeliverable ", func() {
// 			// 	rel, err := GetRelease(ctx, relId)
// 			// 	So(err, ShouldBeNil)
// 			// 	So(rel, ShouldNotBeNil)
// 			// 	//So(delv.DeliveryStatus, ShouldEqual, "undeliverable")
// 			// })
// 		})
// 	})
// }

// /* func TestCreateRelease(t *testing.T) {
// 	Convey("Subject: Create a Release", t, func() {
// 		fmt.Printf("\nTesCreateRelease\n")
// 		os.Setenv("ENV_ROI", "/Users/dhf/work/roi/services/release_service/config")
// 		os.Setenv("CA_LINK", "/Users/dhf/work/roi/services/ca_link_service/config/configuration.test.json")
// 		db.Open()
// 		Convey("Get the release for  5fc3e3b9a652f88181ef5711", func() {
// 			ctx := context.Background()
// 			rel := SampleRelease()

// 			Convey("As Undeliverable ", func() {
// 				release, err := CreateRelease(ctx, rel)
// 				//rel, err := GetRelease(ctx, relId)
// 				So(err, ShouldBeNil)
// 				So(release, ShouldNotBeNil)
// 				//So(delv.DeliveryStatus, ShouldEqual, "undeliverable")
// 			})
// 		})
// 	})
// }

// func TestAddDocument(t *testing.T) {
// 	Convey("Subject: Add a docucment to release", t, func(){
// 		os.Setenv("ENV_ROI", "/Users/dhf/work/roi/services/release_service/config")
// 		db.Open()
// 		ReleaseConnect()
// 		Convey("Create New Release", func() {
// 			ctx :=context.Background()
// 			rel := SampleRelease()
// 			newRel, err := CreateRelease(context.Background(), rel)
// 			So(err, ShouldBeNil)
// 			So(newRel, ShouldNotBeNil)
// 			//relId := newRel.ID
// 			doc := SampleDocument()
// 			doc, err = newRel.AddDocument(ctx, doc)
// 			So(err, ShouldBeNil)
// 			So(doc, ShouldNotBeNil)
// 			fmt.Printf("NewDocument: %s\n", spew.Sdump(doc))
// 		})
// 	})
// //TODO: Remove fake document
// }
//  */
