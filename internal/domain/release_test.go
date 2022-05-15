package domain

//
//
//import (
//	"fmt"
//	"github.com/stretchr/testify/require"
//	"go.mongodb.org/mongo-driver/bson/primitive"
//
//	//"context"
//	//"github.com/davecgh/go-spew/spew"
//	//"github.com/stretchr/testify/require"
//	db "github.com/dhf0820/ids_release_service/internal/db/vsmongo"
//	"github.com/dhf0820/ids_release_service/internal/settings"
//	m "github.com/dhf0820./ds_model"
//	//"go.mongodb.org/mongo-driver/bson/primitive"
//	"os"
//	"testing"
//	. "github.com/smartystreets/goconvey/convey"
//)
//
//func TestCreateDeliveryQueue(t *testing.T) {
//	fmt.Printf("\n\nTestCreateDeliveryQueue\n\n")
//	//Convey("Subject: Create a DeliveryQueue", t, func() {
//	//	os.Setenv("ENV_ROI", "/Users/dhf/work/roi/delivery_service/config")
//	//	Convey("Given: A Valid Delivery Information", func() {
//	//		deviceID 	:= primitive.ObjectIDFromHex("5fc163ed8b273e53e19b1ed5")
//	//		recipientID := primitive.ObjectIDFromHex("5fc184b28b273e53e19b1efa")
//	//		releaseID 	:= primitive.ObjectIDFromHex("5fc03d058b273e53e19b1eb5")
//	//		Convey("Create the DeliveryQueue", func() {
//	//			ctx := context.Background()
//	//			delv, err := CreateDelivery(ctx, releaseID, recipientID, deviceID)
//	//			So(err, ShouldBeNil)
//	//			So(delv, ShouldNotBeNil)
//	//			fmt.Printf("Device: %s\n", spew.Sdump(delv))
//	//		})
//	//	})
//	//})
//}
//
//func TestCreateRelease(t *testing.T) {
//	t.Parallel()
//	db.Open()
//	fmt.Printf("\n\nTestCreateRelease\n")
//	settings.SetDbName("test_release_service")
//	Convey("Subject: Create a new Release", t, func() {
//		os.Setenv("ENV_ROI", "/Users/dhf/work/roi/services/release_service/config")
//		Convey("Given: A Valid release", func() {
//			release := NewRelease()
//			//fmt.Printf("\n\nNewRelease: %s\n\n", spew.Sdump(release))
//			//deviceID 	:= primitive.ObjectIDFromHex("5fc163ed8b273e53e19b1ed5")
//			//recipientID := primitive.ObjectIDFromHex("5fc184b28b273e53e19b1efa")
//			//releaseID 	:= primitive.ObjectIDFromHex("5fc03d058b273e53e19b1eb5")
//			Convey("Create the Release", func() {
//				//ctx := context.Background()
//				rel, err := CreateRelease(release)
//				So(err, ShouldBeNil)
//				So(rel, ShouldNotBeNil)
//				//fmt.Printf("Release: %s\n", spew.Sdump(rel))
//			})
//		})
//	})
//}
//
//
//
//func TestGetReleaseByID(t *testing.T) {
//	t.Parallel()
//	db.Open()
//	fmt.Printf("\n\nTestGetReleaseByID\n")
//	settings.SetDbName("test_release_service")
//	Convey("Subject: Get an existing Release", t, func() {
//		os.Setenv("ENV_ROI", "/Users/dhf/work/roi/services/release_service/config")
//		Convey("Given: A Valid release", func() {
//			release := NewRelease()
//			//deviceID 	:= primitive.ObjectIDFromHex("5fc163ed8b273e53e19b1ed5")
//			//recipientID := primitive.ObjectIDFromHex("5fc184b28b273e53e19b1efa")
//			//releaseID 	:= primitive.ObjectIDFromHex("5fc03d058b273e53e19b1eb5")
//			Convey("Create the Release", func() {
//				//ctx := context.Background()
//				rel, err := CreateRelease(release)
//				So(err, ShouldBeNil)
//				So(rel, ShouldNotBeNil)
//				//fmt.Printf("Release: %s\n", spew.Sdump(rel))
//			})
//		})
//	})
//}
//
//func TestGetReleaseById(t *testing.T) {
//	t.Parallel()
//	db.Open()
//	fmt.Printf("\n\nTestGetReleaseById\n")
//	settings.SetDbName("test_release_service")
//	relID, _ := primitive.ObjectIDFromHex("5fd266801cdeaa46c542526b")
//	rel, err := GetReleaseById(relID)
//	fmt.Printf("GetRelease Error: %v\n", err)
//	require.Nil(t, err)
//	fmt.Printf("returned ReleaseId = %s   requested: %s\n",rel.ID.Hex(), relID )
//	require.Equal(t, relID.Hex(), rel.ID.Hex())
//
//}
//
////func TestGetReleaseByForeignId(t *testing.T) {
////	t.Parallel()
////	fmt.Printf("Open DB\n")
////	db.Open()
////	fmt.Printf("DB is open\n")
////	sourceId := "253"
////	clientId := "CA"
////	fmt.Printf("\n\nTestGetReleaseBySourceId\n")
////	settings.SetDbName("test_release_service")
////	//relID := "5ee510df6f4841277a5029ac"
////	rel, err := GetReleaseByForeignId(sourceId, clientId)
////	require.Nil(t, err)
////	//require.Equal(t, sourceId, rel.SourceID)
////}
//
//
//
//func NewRelease() *m.Release {
//	return &m.Release{
//		Customer: 	testCustomer(),
//		Client: 	testClient(),
//		Patient: 	NewPatient(),
//		Documents:  []m.Document{},
//		History:    []m.DeliveryHistory{},
//	}
//}
//
//func NewPatient() m.Patient {
//	return m.Patient{
//		PatientName: RandomContact(),
//		MRN:  RandomMRN(),
//		DOB:  "1957-10-15",
//		SSN:  "000-04-5867",
//	}
//}
//
//
//
//func testCustomer() m.Customer {
//	c := m.Customer {}
//	c.Name 		= "test"
//	c.Code 		= 1
//	c.Facility 	= "main"
//	c.UserID 	= "tfrench"
//	return c
//}
//
//func testClient() m.Client {
//	c := m.Client {}
//	c.Source 	= "ca"
//	c.SourceID 	= "12345"
//	return c
//}
//
////func NewRecipient() Recipient {
////	return Recipient{
////		Name:  RandomContact(),
////		CompanyName:  RandomCompany(),
////		Address1: RandomAddress(),
////		City:     RandomCity(),
////		State:    "NV",
////		Zip:      "89100",
////		//Fax:      RandomFax(),
////		//Email:    "dhf@vertisoft.com",
////	}
////}
//
////func NewRecipientSummary() RecipientSummary {
////	return RecipientSummary{
////		Contact:  RandomContact(),
////		Company:  RandomCompany(),
////		Address1: RandomAddress(),
////		City:     RandomCity(),
////		State:    "NV",
////		Zip:      "89100",
////	}
////}
