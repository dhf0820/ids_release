package service

import (
	//"bytes"
	//"context"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	relMod "github.com/dhf0820./ds_model/release"
	"go.mongodb.org/mongo-driver/bson/primitive"

	//"github.com/davecgh/go-spew/spew"
	. "github.com/smartystreets/goconvey/convey"
	//documentModel "github.com/dhf0820/ids_document/pkg"
	//"io/ioutil"

	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestAddDocument(t *testing.T) {
	t.Parallel()
	InitTest()
	//Initialize("local", "demo")

	release, _ := CreateRelease(SampleRelease())
	Convey("Subject: AddDocument to release", t, func() {
		fmt.Printf("\nTest AddDocument\n")
		err := AddDocument(release.ID, MakeSampleMetaData())
		So(err, ShouldBeNil)
		docs, err := GetDocuments(release.ID)
		So(err, ShouldBeNil)
		So(len(docs), ShouldEqual, 1)
	})
	RemoveRelease(release.ID)

}

func TestSubmitToDelivery(t *testing.T) {
	t.Parallel()
	InitTest()
	//Initialize("local", "demo")
	recipientId, _ := primitive.ObjectIDFromHex("61d742e2ce5e7cd79f5a5b3d")
	deviceId, _ := primitive.ObjectIDFromHex("622eb6c1f7ac30519ed693f2")
	releaseId, _ := primitive.ObjectIDFromHex("623268b436ab2ee6f654b517")
	//callBackURL := "http://localhost:29911/api/rest/v1/log_status"
	//release, _ := CreateRelease(SampleRelease())
	Convey("Subject: SubmitToDelivery", t, func() {
		fmt.Printf("\nTest SubmitToDelivery\n")
		release, err := getReleaseByFilterM(bson.M{"_id": releaseId})
		So(err, ShouldBeNil)
		So(release, ShouldNotBeNil)
		recipient, err := GetRecipient(recipientId)
		So(err, ShouldBeNil)
		So(recipient, ShouldNotBeNil)
		device, err := GetDevice(deviceId)
		So(err, ShouldBeNil)
		So(device, ShouldNotBeNil)
		// devSum, err := RecipientDeviceSummaryByName(recipient, "dhfrench@gmail.com")
		// So(err, ShouldBeNil)
		// So(devSum, ShouldNotBeNil)
		// device, err := GetDevice(devSum.ID)
		// So(err, ShouldBeNil)
		// So(device, ShouldNotBeNil)

		deliv, err := SubmitToDelivery(release, recipient, device)
		So(err, ShouldBeNil)
		So(deliv, ShouldNotBeNil)

		So(deliv.Status.State, ShouldEqual, "queued")
		//fmt.Printf("Result Delivery-56: %s\n", spew.Sdump(deliv))
	})
	//RemoveRelease(release.ID)
	//RemoveDelivery(deliv.ID)
}

//func TestUpdateDocument(t *testing.T) {
//	t.Parallel()
//	InitTest()
//	Initialize("local", "demo")
//
//	document := MakeSampleMetaData()
//
//	release, _ := CreateRelease(SampleRelease())
//	AddDocument(release.ID, document)
//	Convey("Subject: TestUpdateDocument", t, func() {
//		fmt.Printf("\nTest UpdateDcoument\n")
//		var doc common.DocumentMetadata
//		doc.DocClass = "XXXXXX"
//		rel, err := UpdateDocument(release.ID, document.ID, &doc)
//		So(err, ShouldBeNil)
//		So(rel, ShouldNotBeNil)
//		urel, err := GetRelease(release.ID)
//		So(err, ShouldBeNil)
//		So(urel.Documents[0].DocClass, ShouldEqual, "XXXXXX")
//		fmt.Printf("Updated Document: %s\n", spew.Sdump(urel))
//	})
//	RemoveRelease(release.ID)
//}

func TestCreateRelease(t *testing.T) {
	t.Parallel()
	InitTest()
	fmt.Printf("\n\nTestCreateRelease\n")
	Initialize("local_test", "demo")

	//doc := NewTestDocument("", "")
	//ctx := context.Background()
	var release *relMod.Release
	//var document *pkg.DocumentMetadata
	var err error
	//var releaseId primitive.ObjectID
	Convey("Subject: Create a new Release", t, func() {
		//Convey("Given: A Valid release", func() {
		rel := SampleRelease()
		Convey("Create the Release", func() {
			release, err = CreateRelease(rel)
			So(err, ShouldBeNil)
			So(rel, ShouldNotBeNil)
			//releaseId = release.ID
		})
		Convey("Get a release", func() {
			fmt.Printf("\n##########################GetRelease: %s\n\n", release.ID)
			rel, err := GetRelease(release.ID)
			So(err, ShouldBeNil)
			So(rel, ShouldNotBeNil)
			//fmt.Printf("GotRelease: %s\n", spew.Sdump(rel))
		})
		Convey("Get a release NotFound", func() {
			fmt.Printf("\n##########################GetRelease NotFound: %s\n\n", "61abe571d3bfe446df8f3cdf")
			oid, err := primitive.ObjectIDFromHex("61abe571d3bfe446df8f3cdf")
			So(err, ShouldBeNil)
			rel, err := GetRelease(oid)
			So(err, ShouldNotBeNil)
			So(rel, ShouldBeNil)
			//fmt.Printf("GotRelease: %s\n", spew.Sdump(rel))
		})
		Convey("Lock the Release", func() {
			fmt.Printf("\nLockRelease\n")
			err := SetReleaseLock(release.ID, "lock")
			So(err, ShouldBeNil)
			rel, err := GetRelease(release.ID)
			So(err, ShouldBeNil)
			So(rel.Locked, ShouldEqual, "true")
			//fmt.Printf("GotRelease: %s\n", spew.Sdump(rel))
		})
		//Convey("Subject: GetReleaseByFilter", t, func() {
		//	fmt.Printf("TestGetReleaseByFilter\n")
		//	filter := bson.D{{"_id", release.ID}}
		//	rel, err := getReleaseByFilter(filter)
		//	So(err, ShouldBeNil)
		//	So(rel, ShouldNotBeNil)
		//})
		//Convey("Subject: GetReleaseByForeignId", t, func() {
		//	filter := bson.D{{"src_id", "1958"}}
		//	fmt.Printf("\nGetReleaseByForeignId\n")
		//	rel, err := getReleaseByFilter(filter)
		//	//rel, err := GetReleaseByForeignId("1958")
		//	So(err, ShouldBeNil)
		//	So(rel, ShouldNotBeNil)
		//	So(rel.SrcID, ShouldEqual, "1958")
		//})
		//Convey("Subject: AddDocument to release", t, func() {
		//	//doc := sampleDocument()
		//	fmt.Printf("\nTest AddDocument\n")
		//	doc := MakeSampleMetaData()
		//	err := AddDocument(release.ID, doc)
		//	//count, err := AddDocument(ctx, release.ID, doc)
		//	So(err, ShouldBeNil)
		//	docs, err := GetDocuments(release.ID)
		//	So(err, ShouldBeNil)
		//	So(len(docs), ShouldEqual, 1)
		//	So(doc.ID, ShouldEqual, docs[1].ID)
		//	//document = docs[1]
		//	//fmt.Printf("Document %s\n", spew.Sdump(document))
		//})
		////Convey("Subject: Update a document", t, func() {
		////	fmt.Printf("\nTest UpdateDcoument\n")
		////	//relId, _ := primitive.ObjectIDFromHex("61aac0e16702437a525bf32b")
		////	//docId =  primitive.ObjectIDFromHex("61aaac4847b75e4e346e74c9")
		////	var doc pkg.DocumentMetadata
		////	doc.DocClass = "XXXXXX"
		////	release, err := UpdateDocument(release.ID, document.ID, &doc)
		////	So(err, ShouldBeNil)
		////	So(release, ShouldNotBeNil)
		////	So(release.Documents[0].DocClass, ShouldEqual, "XXXXXX")
		////})
		Convey("Remove the Release", func() {
			fmt.Printf("\nRemove Release\n")
			err := RemoveRelease(release.ID)
			So(err, ShouldBeNil)
			_, err = GetRelease(release.ID)
			So(err, ShouldNotBeNil)
		})

		//})
	})
}

func TestGetReleaseByFilter(t *testing.T) {
	t.Parallel()
	InitTest()
	fmt.Printf("\n\nTestGetReleaseByFilter\n")
	InitTest()
	Initialize("local", "demo")
	//var release *pkg.Release
	//var err error

	release, _ := CreateRelease(SampleRelease())
	Convey("Subject: Get a releaseByFilter", t, func() {
		filter := bson.D{{"_id", release.ID}}
		rel, err := getReleaseByFilter(filter)
		So(err, ShouldBeNil)
		So(rel, ShouldNotBeNil)

	})
	RemoveRelease(release.ID)
}

func TestGetRelease(t *testing.T) {
	t.Parallel()
	fmt.Printf("\n\nTestGetRelease\n")
	InitTest()
	//Initialize("local", "demo")
	//release, _ := CreateRelease(SampleRelease())
	releaseID, _ := primitive.ObjectIDFromHex("61ce422b2c4df1d789338321")
	Convey("Subject: Create a new Release", t, func() {
		//oid, err := primitive.ObjectIDFromHex("61a7cec75fa8fca189445c76")
		release, err := GetRelease(releaseID)
		So(err, ShouldBeNil)
		So(release, ShouldNotBeNil)
		//So(rel.SrcID, ShouldEqual, "199")
		fmt.Printf("Release: %s\n", spew.Sdump(release))
	})
	//RemoveRelease(release.ID)

}

func TestGetTestRelease(t *testing.T) {
	t.Parallel()
	fmt.Printf("\n\nTestGetTestRelease\n")
	InitTest()
	Initialize("local_test", "demo")
	release, _ := CreateSampleRelease()
	//release, _ := CreateRelease(SampleRelease())
	Convey("Subject: Create a new Release", t, func() {
		//oid, err := primitive.ObjectIDFromHex("61a7cec75fa8fca189445c76")
		rel, err := GetRelease(release.ID)
		So(err, ShouldBeNil)
		So(rel, ShouldNotBeNil)
		So(rel.SrcID, ShouldEqual, "1958")
	})
	RemoveRelease(release.ID)

}
func TestGetReleaseDocuments(t *testing.T) {
	t.Parallel()
	fmt.Printf("\n\nTestGetDocuments\n")
	InitTest()
	Initialize("local_test", "demo")
	doc1 := MakeSampleMetaData()
	doc2 := MakeSampleMetaData()
	doc2Id, _ := primitive.ObjectIDFromHex("61aaf11b2ea223828e2aa3f0")
	doc2.ID = doc2Id
	release, _ := CreateRelease(SampleRelease())
	AddDocument(release.ID, doc1)
	AddDocument(release.ID, doc2)
	Convey("Subject: Create a new Release", t, func() {
		//oid, err := primitive.ObjectIDFromHex("61a7cec75fa8fca189445c76")
		docs, err := GetDocuments(release.ID)
		So(err, ShouldBeNil)
		So(len(docs), ShouldEqual, 2)
		So(docs[1].ID, ShouldEqual, doc2Id)
		d := docs[1]
		//fmt.Printf("document; %s\n", spew.Sdump(d))
		doc, err := GetDocument(d)
		So(err, ShouldBeNil)
		So(doc, ShouldNotBeNil)

		fmt.Printf("Version: [%s]\n", doc.Version)
	})

	RemoveRelease(release.ID)

}

func TestBuildDocumentSet(t *testing.T) {
	t.Parallel()
	fmt.Printf("\n\nTestBuildReleaseSet\n")
	InitTest()
	Initialize("local_test", "demo")
	// doc1 := MakeSampleMetaData()
	// doc2 := MakeSampleMetaData()
	// doc2Id, _ := primitive.ObjectIDFromHex("61aaf11b2ea223828e2aa3f0")
	// doc2.ID = doc2Id
	// release, _ := CreateRelease(SampleRelease())
	// AddDocument(release.ID, doc1)
	// AddDocument(release.ID, doc2)
	Convey("Subject: BuildDocumentSet", t, func() {
		rel_oid, err := primitive.ObjectIDFromHex("62328c7a36ab2ee6f654b529")
		docSet, err := BuildDocumentSet(rel_oid)

		//docs, err := GetDocuments(rel_oid)
		So(err, ShouldBeNil)
		So(len(docSet), ShouldNotEqual, 0)
		fmt.Printf("Number of documents in release : %d\n", len(docSet))
	})

	//RemoveRelease(release.ID)

}

func TestGetReleaseCombined(t *testing.T) {
	t.Parallel()
	fmt.Printf("\n\nTestGetReleaseDCombined\n")
	InitTest()
	Initialize("local_test", "demo")
	// doc1 := MakeSampleMetaData()
	// doc2 := MakeSampleMetaData()
	// doc2Id, _ := primitive.ObjectIDFromHex("61aaf11b2ea223828e2aa3f0")
	// doc2.ID = doc2Id
	// release, _ := CreateRelease(SampleRelease())
	// AddDocument(release.ID, doc1)
	// AddDocument(release.ID, doc2)
	Convey("Subject: Create a new Release", t, func() {
		rel_oid, err := primitive.ObjectIDFromHex("62328c7a36ab2ee6f654b529")
		combinedRelease, err := GetReleaseCombined(rel_oid)
		So(err, ShouldBeNil)
		So(len(*combinedRelease), ShouldNotEqual, 0)
		fmt.Printf("Size of combined: %d\n", len(*combinedRelease))
	})

	//RemoveRelease(release.ID)

}

func TestGetReleaseByForeignId(t *testing.T) {
	t.Parallel()
	fmt.Printf("\n\nTestGetReleaseByFilter\n")
	InitTest()
	Initialize("local", "demo")
	nRelease := SampleRelease()
	nRelease.SrcID = "1958"
	release, _ := CreateRelease(nRelease)
	Convey("Subject: Create a new Release", t, func() {
		rel, err := GetReleaseByForeignId("1958")
		So(err, ShouldBeNil)
		So(rel, ShouldNotBeNil)
		So(rel.SrcID, ShouldEqual, "1958")
	})

	RemoveRelease(release.ID)
}
