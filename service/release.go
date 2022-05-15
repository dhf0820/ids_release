package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/davecgh/go-spew/spew"
	common "github.com/dhf0820/ids_model/common"
	delvMod "github.com/dhf0820/ids_model/delivery"
	devMod "github.com/dhf0820/ids_model/device"
	docMod "github.com/dhf0820/ids_model/document"
	recipMod "github.com/dhf0820/ids_model/recipMod"
	relMod "github.com/dhf0820/ids_model/release"
	log "gitgithub.com/dhf0820.ids_model
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	//"strings"
	"time"
)

var updateFields = bson.D{}
var fieldsInQuery = 0

//################################## Main Functions ################################
func CreateRelease(rel *relMod.Release) (*relMod.Release, error) {
	var err error
	collection, err := GetCollection("releases")
	if err != nil {
		fmt.Printf("GetCollection [releases] failed: %v\n", err)
		return nil, err
	}

	rel.CreatedAt = time.Now()
	rel.UpdatedAt = rel.CreatedAt
	res, err := collection.InsertOne(context.Background(), rel)
	if err != nil {
		log.Errorf("InsertOne failed: %v\n", err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error saving Document: %v", err),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot retrieve InsertedId"),
		)
	}
	rel.ID = oid
	return rel, err
}

func GetRelease(releaseId primitive.ObjectID) (*relMod.Release, error) {
	fmt.Printf("\n   GetRelease:62\n")
	//filter := bson.D{{"_id", releaseId}}
	filter := bson.M{"_id": releaseId}
	fmt.Printf("GetRelease-65 filter: %v\n", filter)
	release, err := getReleaseByFilterM(filter)
	return release, err
}

// func GetReleaseImage(relID primitive.ObjectID) (*[]byte, error) {
// 	doc, err := GetDocument(docId)
// 	if err != nil {
// 		fmt.Printf("GetDocumentImage error: %s\n", err.Error())
// 		return nil, err
// 	}
// 	//fmt.Printf("DocumentImage: %s\n", spew.Sdump(doc))
// 	imageId := doc.ImageID

// 	image, err := getGridFsImage(imageId)
// 	fmt.Printf("Length of image: %d\n", len(*image))
// 	if err != nil {
// 		return nil, fmt.Errorf("GetImage error: %s", err.Error())
// 	}
// 	return image, nil
// }

func GetReleaseImage(rel *relMod.Release) (*[]byte, error) {
	imageId := rel.ImageId

	image, err := getGridFsImage(imageId)
	fmt.Printf("Length of image: %d\n", len(*image))
	if err != nil {
		return nil, fmt.Errorf("GetImage error: %s", err.Error())
	}
	return image, nil
}

// doc, err := GetDocument(docId)
// if err != nil {
// 	fmt.Printf("GetDocumentImage error: %s\n", err.Error())
// 	return nil, err
// }
// //fmt.Printf("DocumentImage: %s\n", spew.Sdump(doc))
// imageId := doc.ImageID

// image, err := getGridFsImage(imageId)
// fmt.Printf("Length of image: %d\n", len(*image))
// if err != nil {
// 	return nil, fmt.Errorf("GetImage error: %s", err.Error())
// }
// return image, nil
// }

func GetDocuments(releaseId primitive.ObjectID) ([]*common.DocumentMeta, error) {
	fmt.Printf("GetDocuments: 71 Getrelease: %s\n", releaseId)
	rel, err := GetRelease(releaseId)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Release [%s] not found:  %v", releaseId, err),
		)
	}
	return rel.Documents, nil
}

func GetReleaseBySrcId(srcId, facility string) (*relMod.Release, error) {
	if srcId == "" {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("GetReleaseBySrcIdFacility invalid params srcId: [%s], facility: [%s]", srcId, facility),
		)
	}
	filter := bson.D{{"Client.SourceId", srcId}, {"Customer.Facility", facility}}
	return getReleaseByFilter(filter)
}

func SubmitToDelivery(release *relMod.Release, recipient *recipMod.Recipient, device *devMod.Device) (*delvMod.Delivery, error) {
	//func SubmitToDelivery(releaseId, recipientId, deviceId primitive.ObjectID) (*delvMod.Delivery, error) {
	//recipientId, err := primitive.ObjectIDFromHex("61b53c8389233aab9bc94e31")
	cfg := GetConfig()

	// // recipient, err := GetRecipient(recipientId)
	// // _, err = GetRecipient(recipientId)

	// //recipient, err := CreateRecipient(SampleRecipient())
	// // if err != nil {
	// // 	return nil, fmt.Errorf("Recipient could not be found: %s", err)
	// // }
	// //fmt.Printf("#### Recipient: %s\n", spew.Sdump(recipient))
	// //deviceId := recipient.Devices[0].ID
	// device, err := GetDevice(deviceId)
	// fmt.Printf("Deliverying to %s\n", spew.Sdump(device))
	// //device, err := CreateDevice(SampleEmail("emailTest"), recipient.ID, "emailTest")
	// if err != nil {
	// 	return nil, fmt.Errorf("Device could not be found: %s", err)
	// }

	// //fmt.Printf("#### Device: %s\n", spew.Sdump(device))
	// release, err := GetRelease(releaseId)
	// if err != nil {
	// 	fmt.Errorf("release id %s is invalid: %s", releaseId, err.Error())
	// }

	//fmt.Printf("#### Release: %s\n", spew.Sdump(release))
	curTime := time.Now()
	delv := delvMod.Delivery{}

	common.GetDataByName(cfg.CallBacks, "status")
	delv.ReleaseID = release.ID
	delv.CallBackStatusUrl = common.GetDataByName(cfg.CallBacks, "status")
	delv.CallBackLogMessageUrl = common.GetDataByName(cfg.CallBacks, "log")
	delv.Meta = []*common.KVData{}
	pname := common.KVData{
		Name:  "patient_name",
		Value: release.Patient.PatientName,
	}
	delv.Meta = append(delv.Meta, &pname)
	pmrn := common.KVData{
		Name:  "mrn",
		Value: release.Patient.MRN,
	}
	delv.Meta = append(delv.Meta, &pmrn)
	pssn := common.KVData{
		Name:  "ssn",
		Value: release.Patient.SSN,
	}
	delv.Meta = append(delv.Meta, &pssn)
	pdob := common.KVData{
		Name:  "dob",
		Value: release.Patient.DOB,
	}
	delv.Meta = append(delv.Meta, &pdob)
	delv.Documents = []*common.DocumentMeta{}
	cfg = GetConfig()

	ep := common.GetServiceEndpoint(cfg.ServiceEndPoints, "document")
	docUrl := ep.Address
	fmt.Printf("\n\n\n######  DocumentUrl: %s\n\n\n\n", docUrl)
	delv.Documents = release.Documents
	delv.Recipient.CompanyName = recipient.CompanyName
	delv.Recipient.ID = recipient.ID
	delv.Recipient.Name = recipient.Name
	delv.Recipient.FirstName = recipient.FirstName
	delv.Recipient.LastName = recipient.LastName
	delv.Device.ID = device.ID
	delv.Device.Method = device.Method
	delv.Device.Name = device.Name
	delv.Device.Active = device.Active
	delv.Device.Validated = device.Validated
	delv.Device.ValidationDate = device.ValidationDate
	delv.Device.Fields = device.Fields
	delv.Device.RetryPlan = device.RetryPlan
	delv.Priority = 3 //TODO: where to get default Priority
	delv.Attempts = 0
	delv.NextTryTime = curTime
	delv.TimeToLive = curTime.UTC().Add(time.Hour * 24)
	delv.Status = &common.Status{}
	delv.Status.State = "queued"
	delv.Status.StatusTime = curTime
	delv.CreatedAt = curTime
	delv.UpdatedAt = curTime

	//TODO: Call delivery to create the deliveryQueue Entry
	fmt.Printf("\n\n\n#####Release SubmitToDelivery174: %s\n", spew.Sdump(delv))
	deliv, err := QueueDelivery(&delv)
	if err != nil {
		fmt.Printf("\n@@@ QueueDelivery Error: %s\n", err.Error())
	}
	return deliv, err
}

func SetReleaseLock(oid primitive.ObjectID, mode string) error {
	_, err := GetRelease(oid)
	if err != nil {
		return status.Errorf(
			codes.NotFound,
			fmt.Sprintf("GetRelease-37 Id: %s was not found", oid),
		)
	}
	updRelease := relMod.Release{}
	updRelease.ID = oid
	switch mode {
	case "lock":
		updRelease.Locked = "true"
	case "unlock":
		updRelease.Locked = "false"
	default:
		return status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Invalid Mode: [%s]", mode),
		)
	}
	_, err = UpdateRelease(&updRelease)
	return err
}

//CreateReleaseArchive: Takes the release ID and retrieves all documents and concatinates pdf images into
// one pdf containing all documents in the release, saving it in GridFS and returning the url of the image

// func CreateReleaseArchive(id primitive.ObjectID) (primitive.ObjectID, error) {

// }

func GetDocument(docMeta *common.DocumentMeta) (*docMod.DeliveryDocument, error) {
	documentResp := docMod.DocumentResponse{}
	docUrl := docMeta.DocumentURL
	fmt.Printf("GetDocument:268 - docMeta: %s\n", spew.Sdump(docMeta))
	fmt.Printf("GetDocument:269 - DocUrl: %s\n", docUrl)

	req, err := http.NewRequest("GET", docUrl, nil)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("Access Authorization: %s\n", endPoint.Access)
	//req.Header.Set("AUTHORIZATION", "37")
	//req.Header.Set("facility", "demo")
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		//TODO: Handle Error from getting document
		log.Fatal("Error reading response. ", err)
	}
	if resp.StatusCode != 200 {
		err = fmt.Errorf("GetDocumentError: %d - %s", resp.StatusCode, resp.Status)
		return nil, err
	}
	var f interface{}
	//	hook := HookData{}
	//	err = json.Unmarshal(results, &hook)
	//	So(err, ShouldBeNil)
	defer resp.Body.Close()

	decode := json.NewDecoder(resp.Body).Decode(&f)
	if decode != nil {
		err := fmt.Errorf("error Decoding GetDocumentWithImage:240 -- document: %s", decode.Error())
		//TODO: Handle decoe document error properly
		log.Fatal(err)
	}
	fmt.Printf("f: %s\n", spew.Sdump(f))
	dd := docMod.DeliveryDocument{}
	dd.Description = docMeta.Description
	dd.Image = documentResp.Image
	dd.DocClass = common.GetDataByName(docMeta.Meta, "doc_class")
	dd.DateOfService = common.GetDataByName(docMeta.Meta, "date_of_service")
	dd.ImageType = documentResp.Document.ImageType
	dd.ImageRepository = documentResp.Document.ImageRepository
	dd.ImageURL = docMeta.DocumentURL //This gets teh combined image and document.
	dd.DocumentURL = docMeta.DocumentURL
	//dd.DocumentURL = "http://localhost:29912/api/v1/document/61b6c89bd8cc9541dde09f1b?image=none"
	//dd.FileName = fmt.Sprintf("doc-%s-%s.dd", docMeta.ID.Hex, dd.DocClass)
	//ioutil.WriteFile(dd.FileName, *image,0444)
	//fmt.Printf("dd: %s\n", spew.Sdump(dd))
	dd.Image = documentResp.Image
	return &dd, nil
}
func GetDocumentWithImage(docMeta *common.DocumentMeta) (*docMod.DeliveryDocument, error) {

	fmt.Printf("GetDocumentWithImage meta: %s\n", spew.Sdump(docMeta))
	documentResp := docMod.DocumentResponse{}
	baseDocURL := os.Getenv("DOCUMENT_HOST")
	if baseDocURL == "" {
		log.Errorf("DocumentBase is not setup using localhost")
		baseDocURL = "http://localhost:19203/api/rest/v1/"
	}
	docUrl := fmt.Sprintf("%s%s?image=include", baseDocURL, docMeta.DocumentURL)
	fmt.Printf("\nGetDocumentWithImage docURL: %s\n\n", docUrl)

	req, err := http.NewRequest("GET", docUrl, nil)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("Access Authorization: %s\n", endPoint.Access)
	//req.Header.Set("AUTHORIZATION", "37")
	//req.Header.Set("facility", "demo")
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		//TODO: Handle Error from getting document
		log.Errorf("Error reading response: %s", err.Error())
		return nil, err
	}
	if resp.StatusCode != 200 {
		err = fmt.Errorf("GetDocument: %s error: %s", docUrl, resp.Status)
		log.Errorf("Get %s failed: %s", docUrl, resp.Status)
		return nil, err
	}
	fmt.Printf("resp: Status: %d, message: %s\n", resp.StatusCode, resp.Status)
	defer resp.Body.Close()

	decode := json.NewDecoder(resp.Body).Decode(&documentResp)
	if decode != nil {
		err := fmt.Errorf("error Decoding GetDocumentWithImage:282 -- document: %s", decode.Error())
		//TODO: Handle decoe document error properly
		log.Fatal(err)
	}

	//fmt.Printf("Setup DeliveryDocument: %s\n\n", spew.Sdump(documentResp))
	fmt.Printf("GetDocumentWithImageUrl: Setup DeliveryDocument: %s\n", docMeta.DocumentURL)
	dd := docMod.DeliveryDocument{}
	dd.Description = docMeta.Description
	dd.Image = documentResp.Image
	dd.DocClass = common.GetDataByName(docMeta.Meta, "doc_class")
	dd.DateOfService = common.GetDataByName(docMeta.Meta, "date_of_service")
	dd.ImageType = documentResp.Document.ImageType
	dd.ImageRepository = documentResp.Document.ImageRepository
	dd.ImageURL = docMeta.DocumentURL //This gets teh combined image and document.
	dd.DocumentURL = docUrl           //"http://localhost:29912/api/rest/v1/document/61b6c89bd8cc9541dde09f1b?image=none"
	//dd.FileName = fmt.Sprintf("doc-%s.dd", docMeta.ID.Hex, dd.DocClass)
	//ioutil.WriteFile(dd.FileName, *image,0444)
	//fmt.Printf("\n\nImage Size: %d  -  dd: %s\n", len(*documentResp.Image), spew.Sdump(dd))
	dd.Image = documentResp.Image
	return &dd, nil
}

//func BuildDocumentSet(job *common.Payload) ([]*docMod.DeliveryDocument, error) {
//func BuildDocumentSet(meta []*common.KVData, docs []*common.DocumentMeta) ([]*docMod.DeliveryDocument, error) {
func BuildDocumentSet(relID primitive.ObjectID) ([]*docMod.DeliveryDocument, error) {
	fmt.Printf("\n\n\n\nBuildDocumentSet for release: %s\n", relID.Hex())
	FileSet := []*docMod.DeliveryDocument{}
	release, err := GetRelease(relID)
	if err != nil {
		return nil, err
	}

	n := 0

	for _, doc := range release.Documents {
		n = n + 1

		fmt.Printf("\n\n\nBuild is working on document %s  # %d in build set\n", doc.ID.Hex(), n)
		delvDoc, err := GetDocumentWithImage(doc)

		if err != nil {
			log.Errorf("GetDocumcentWithImage error: %s", err.Error())
			//msg := fmt.Sprintf("Job: %s could not retrieve document Image:[%s]", job.DelvPayload.ID, doc.ID)
			//logging.LogMessage(job, "logs", "warn", msg, "email_connector")
			//TODO: If document missing on email create a place holder giving the class, description and DOS.
			continue
			//TODO: Log unable to retrieve document image for a release to an error list stored in the Delivery History
		}
		FileSet = append(FileSet, delvDoc)

		//tempFolder := GetTempDir()
		// docId := doc.ID.Hex()
		// // tempFolder := common.GetDataByName(Config.Data, "temp_dir")
		// // tempFolder := common.GetDataByName(job.Config.Data, "temp_dir")
		// if tempFolder == "" {
		// 	tempFolder = "./tmp"
		// }
		// //tempFolder := os.Getenv("TEMP_FOLDER")

		// delvDoc.FileName = fmt.Sprintf("%s/%s.pdf", tempFolder, docId)
		// fmt.Printf("FileName: %s\n", delvDoc.FileName)
		// // file, err := os.OpenFile(delvDoc.FileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		// // if err != nil {
		// // 	panic(err)
		// // }
		// // defer file.Close()
		// if delvDoc.Image == nil {
		// 	fmt.Printf("Image is nil\n")
		// } else {
		// 	fmt.Printf("Image is not nil\n")
		// }
		// fmt.Printf("Size of image: %d\n", len(*delvDoc.Image))
		// ioutil.WriteFile(delvDoc.FileName, *delvDoc.Image, 0444)
		// if err != nil {
		// 	fmt.Printf("Write file %s failed: %s\n", delvDoc.FileName, err.Error())
		// 	//TODO: If document missing on email create a place holder giving the class, description and DOS.
		// 	//TODO: Log unable to retrieve document image for a release to an error list stored in the Delivery History
		// 	// msg := fmt.Sprintf("Save Image to %s failed with %s", delvDoc.FileName, err.Error())
		// 	// logging.LogMessage(job, "logs", "error", msg, "email_connector")
		// } else {
		// 	//Since we have a file containing the image, remove the image from memory
		// 	delvDoc.Image = nil
		// }

		// FileSet = append(FileSet, delvDoc)
	}
	return FileSet, nil

	// combinedName := fmt.Sprintf("%s/rel_%s.pdf", relID.Hex(), tempFolder)
	// inFiles := []string{}
	// for _, doc := range FileSet {
	// 	inFiles = append(inFiles, doc.FileName)
	// }
	// err = MergeToCombined(inFiles, combinedName, false)
	// if err != nil {
	// 	return nil, err
	// }

	// content, err := ioutil.ReadFile(combinedName)
	// if err != nil {
	// 	return nil, err
	// }

	// return fileset, nil
}

func GetReleaseCombined(relID primitive.ObjectID) (*[]byte, error) {
	fmt.Printf("\n\nGetReleaseCombined for %s\n", relID.Hex())
	tempFolder := GetTempDir()
	releaseSet, err := BuildDocumentSet(relID)
	if err != nil {
		return nil, err

	}
	fmt.Printf(" \n\n\nBuildDocumentSet returned %d documents in release\n", len(releaseSet))
	relId := relID.Hex()
	inFiles := []string{}
	n := 0
	for _, delvDoc := range releaseSet {
		fmt.Printf("")
		//docId := doc.ID.Hex()
		// tempFolder := common.GetDataByName(Config.Data, "temp_dir")
		// tempFolder := common.GetDataByName(job.Config.Data, "temp_dir")
		if tempFolder == "" {
			tempFolder = "./tmp"
		}
		//tempFolder := os.Getenv("TEMP_FOLDER")

		delvDoc.FileName = fmt.Sprintf("%s/rel_%s_%d.pdf", tempFolder, relId, n)
		fmt.Printf("GetReleaseCombined: 483 -- FileName: %s\n", delvDoc.FileName)
		// file, err := os.OpenFile(delvDoc.FileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		// if err != nil {
		// 	panic(err)
		// }
		// defer file.Close()
		if delvDoc.Image == nil {
			fmt.Printf("Image is nil\n")
		} else {
			fmt.Printf("Image is not nil\n")
		}
		fmt.Printf("Size of image: %d\n", len(*delvDoc.Image))
		fmt.Printf("writing File: %s\n", delvDoc.FileName)
		ioutil.WriteFile(delvDoc.FileName, *delvDoc.Image, 0444)
		if err != nil {
			fmt.Printf("Write file %s failed: %s\n", delvDoc.FileName, err.Error())
			err = nil
			//TODO: If document missing on email create a place holder giving the class, description and DOS.
			//TODO: Log unable to retrieve document image for a release to an error list stored in the Delivery History
			// msg := fmt.Sprintf("Save Image to %s failed with %s", delvDoc.FileName, err.Error())
			// logging.LogMessage(job, "logs", "error", msg, "email_connector")
		}
		inFiles = append(inFiles, delvDoc.FileName)
		n = n + 1
	}
	combinedName := fmt.Sprintf("%s/rel_%s.pdf", tempFolder, relID.Hex())

	// for _, doc := range FileSet {
	// 	inFiles = append(inFiles, doc.FileName)
	// }
	err = MergeToCombined(inFiles, combinedName, false)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadFile(combinedName)
	if err != nil {
		return nil, err
	}
	return &content, nil
}

// Update documents in a release to add a new document

//TODO: Update a change in a document in a release. Add, delete, document Information changed
//func UpdateDocument(releaseId primitive.ObjectID, documentId primitive.ObjectID, document *common.DocumentMeta) (*relMod.Release, error) {
//	//Get find/pull the document
//	filter := bson.M{"_id": releaseId, "documents.id": documentId} //returns the release if it has that document
//
//	release, err := getReleaseByFilterM(filter)
//	if err != nil {
//		return nil, err
//	}
//
//	var existsDocument *common.DocumentMeta
//	if release.Documents != nil {
//		for _, doc := range release.Documents {
//			if doc.ID == documentId {
//				existsDocument = doc
//				break
//			}
//		}
//	}
//
//	hasUpdate := false
//
//	if strings.Trim(document.DocClass, " ") != "" && strings.Trim(document.DocClass, " ") != existsDocument.DocClass {
//		// TODO: Build change audit
//		existsDocument.DocClass = document.DocClass
//		hasUpdate = true
//	}
//	if hasUpdate == false {
//		return nil, errors.New("No changes found.")
//	}
//	updateAt := time.Now()
//	existsDocument.UpdatedAt = updateAt
//	updateQuery := bson.M{"$set": bson.M{"documents.$": existsDocument, "updated_at": updateAt}} //updates the db with the existsDocument
//
//	//update_query := bson.M{"$set": bson.M{"documents.$.description": document.Description,  //updates the description in the found document
//	//	"documents.$.updated_at": updateAt,  // updates the updated_at also
//	//	"updated_at": updateAt}}  // updates the parent updated_at
//	col, _ := GetCollection("releases")
//	_, err = col.UpdateOne(context.Background(), filter, updateQuery)
//	if err != nil {
//		return nil, err
//	}
//	release, err = getReleaseByFilterM(filter)
//	if err != nil {
//		return nil, err
//	}
//
//	return release, nil
//}

//TODO: RemoveDocument is null
func RemoveDocument(releaseId primitive.ObjectID, documentMetadataId primitive.ObjectID) {}

//func AddRecipient(releaseId primitive.ObjectID, recipient *pkg.Recipients)  error{
//	col, _  := GetCollection("releases")
//	rel, err := GetRelease(releaseId)
//	if err != nil {
//		return fmt.Errorf("Add AddRecipient-168 Release Id: %s was not found", releaseId)
//	}
//	//rel.Documents = append(rel.Documents, documentMetadata)
//	//res, err :=col.UpdateOne(context.Background(), bson.M{"_id": releaseID},
//	//	{$set
//	//})
//	//if err != nil {
//	//	fmt.Printf("Update Metadata error: %v\n", err)
//	//	return err
//	//}
//	//count := res.ModifiedCount
//	//upsertcount := res.UpsertedCount
//	//fmt.Printf("inserted: %d - upserted: %d\n", count, upsertcount)
//
//
//	//data := bson.D{{"$set", {"documents": documentMetadata}}}
//
//	////fmt.Printf("Recipient Data to update: %v\n", data)
//	////fmt.Printf("Filter: %v\n", filter)
//	//_, err = collection.UpdateOne(
//	//	ctx,
//	//	filter,
//	//	data)
//	//"updated_at": time.Now(),
//	filter := bson.M{"_id": releaseId}
//
//	updateQry := bson.M{"$addToSet":bson.M{"recipients": recipient }}
//	//fmt.Printf("Filter: %s\n", spew.Sdump(filter))
//	//fmt.Printf("update: %s\n", spew.Sdump(updateQry))
//	_, err = col.UpdateOne(context.Background(), filter,updateQry )
//	if err != nil {
//		fmt.Printf("Update Documents error: %v\n", err)
//		return err
//	}
//	//count := res.ModifiedCount
//	//upsertCount := res.UpsertedCount
//	//update := bson.M{$addToSet: bson.M{"documents": documentMetadata},}
//	//_, err := col.UpdateOne(context.Background(), bson.M{"_id": releaseID}, update)
//
//
//	//count := 1 //res.ModifiedCount
//	//count := res.UpsertedCount
//	//fmt.Printf("inserted: %d - upserted: %d \n", count,upsertCount)
//	return err
//}

func UpdateRecipients(releaseId primitive.ObjectID, recipient *recipMod.Recipient) (*relMod.Release, error) {
	summary := recipMod.SummaryFromRecipient(recipient)
	collection, _ := GetCollection("releases")
	filter := bson.M{"_id": releaseId, "recipients.RecipientId": recipient.ID} //returns the release if it has that document
	updateAt := time.Now()
	var updateQuery bson.M
	_, err := getReleaseByFilterM(filter)
	if err == nil {
		//Found Release and Recipient
		fmt.Printf("Recipient Found, updating it\n")
		updateQuery = bson.M{"$set": bson.M{"recipients.$": summary, "updated_at": updateAt}} //updates the db with the existsDocument
		_, err = collection.UpdateOne(context.Background(), filter, updateQuery)
		if err != nil {
			fmt.Printf("Update Recipients error-227: %v\n", err)
			return nil, err
		}
	} else {
		fmt.Printf("Recipient not found adding\n")
		filter = bson.M{"_id": releaseId}
		//fmt.Println("GetRelease by id only")
		_, err := GetRelease(releaseId)
		if err != nil {
			return nil, fmt.Errorf("finding release: %s - %s", releaseId, err.Error())
		}
		//fmt.Printf("Updating release: %s\n", spew.Sdump(rel))
		//updateQry := bson.M{"$addToSet":bson.M{"documents": document }}
		//TODO: Update the full record LastUpdate time.&
		updateQuery = bson.M{"$addToSet": bson.M{"recipients": summary}}
		//fmt.Printf("Finding one by recipient only: %s\n", releaseId)
		_, err = collection.UpdateOne(context.Background(), filter, updateQuery)
		if err != nil {
			fmt.Printf("Update Recipients error-238: %v\n", err)
			return nil, err
		}
	}
	//_, err = collection.UpdateOne(context.Background(), filter, updateQuery)
	//if err != nil {
	//	fmt.Printf("Update Recipients error-244: %v\n", err)
	//	return nil, err
	//}
	fmt.Printf("Looking for updated release\n")
	rel, err := GetRelease(releaseId)
	if err != nil {
		return nil, fmt.Errorf("Cound not refind release: %s\n", releaseId)
	}
	return rel, nil
}

//fmt.Printf("Found Release: %s\n", spew.Sdump(release))
//var existsRecipient *pkg.Recipients
//if release.Recipients != nil {
//	for _, recip := range release.Recipients {
//		if recip.RecipientId == recip.RecipientId {
//			existsRecipient = recip
//			break
//		}
//	}
//}
//
////hasUpdate := false
//
////updateAt := time.Now()
////var updateQuery bson.M
//for _, recip := range release.Recipients {
//	if recip.RecipientId == recipient.RecipientId {
//		//Exists update it
//		updateQuery = bson.M{"$set":bson.M{"recipients.$" : existsRecipient,"updated_at": updateAt}} //updates the db with the existsDocument
//	} else {
//		//Does not exist add it
//		updateQuery = bson.M{"$addToSet": bson.M{"recipients": recipient}}
//		filter = bson.M{"_id": release.ID}
//	}
//	_, err := collection.UpdateOne(context.Background(), filter,updateQuery )
//	if err != nil {
//		fmt.Printf("Update Recipients error: %v\n", err)
//		return nil, err
//	}
//}
//rel, err := service.GetRelease(releaseId)
//if err != nil {
//	return nil, fmt.Errorf("Cound not refind release: %s\n", releaseId)
//}
//return rel, nil
//}

func AddDocument(releaseID primitive.ObjectID, document *common.DocumentMetadata) error {
	col, _ := GetCollection("releases")
	//rel, err := GetRelease(releaseID)
	//if err != nil {
	//	return status.Errorf(
	//		codes.NotFound,
	//		fmt.Sprintf("Add Document-92 Release Id: %s was not found", releaseID),
	//	)
	//}
	//rel.Documents = append(rel.Documents, documentMetadata)
	//res, err :=col.UpdateOne(context.Background(), bson.M{"_id": releaseID},
	//	{$set
	//})
	//if err != nil {
	//	fmt.Printf("Update Metadata error: %v\n", err)
	//	return err
	//}
	//count := res.ModifiedCount
	//upsertcount := res.UpsertedCount
	//fmt.Printf("inserted: %d - upserted: %d\n", count, upsertcount)

	//data := bson.D{{"$set", {"documents": documentMetadata}}}

	////fmt.Printf("Recipient Data to update: %v\n", data)
	////fmt.Printf("Filter: %v\n", filter)
	//_, err = collection.UpdateOne(
	//	ctx,
	//	filter,
	//	data)
	//"updated_at": time.Now(),
	filter := bson.M{"_id": releaseID}

	updateQry := bson.M{"$addToSet": bson.M{"documents": document}}
	//fmt.Printf("Filter: %s\n", spew.Sdump(filter))
	//fmt.Printf("update: %s\n", spew.Sdump(updateQry))
	_, err := col.UpdateOne(context.Background(), filter, updateQry)
	if err != nil {
		fmt.Printf("Update Release: %s Documents error: %v\n", releaseID, err)
		return err
	}
	//count := res.ModifiedCount
	//upsertCount := res.UpsertedCount
	//update := bson.M{$addToSet: bson.M{"documents": documentMetadata},}
	//_, err := col.UpdateOne(context.Background(), bson.M{"_id": releaseID}, update)

	//count := 1 //res.ModifiedCount
	//count := res.UpsertedCount
	//fmt.Printf("inserted: %d - upserted: %d \n", count,upsertCount)
	return err
}

//ToDo: Update ReleaseStatus

func UpdateRelease(r *relMod.Release) (*relMod.Release, error) {

	collection, err := GetCollection("releases")
	//timeNow := time.Now().Fopkg.t(Fopkg.tDate)
	r.UpdatedAt = time.Now()
	//u := bson.D{}
	//update = bson.D{}
	filter := bson.M{"_id": r.ID}
	//fmt.Printf("RecipientUpdate: %s\n", r.ID.Hex())
	if r.Locked != "" {
		buildUpdate("locked", r.Locked)
	}
	data := bson.D{{"$set", updateFields}}
	//fmt.Printf("Recipient Data to update: %v\n", data)
	//fmt.Printf("Filter: %v\n", filter)
	_, err = collection.UpdateOne(
		context.Background(),
		filter,
		data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error update Release using filter: %s:  %v\n", filter, err),
		)
	}
	//fmt.Printf("Updated %v Documents\n", result.ModifiedCount)
	rel, err := GetRelease(r.ID)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Could not verify the Release with ID: %s\n", r.ID.Hex()),
		)
	}
	r = rel
	return rel, nil
}

func buildUpdate(key, value string) {
	//value = SetDashToBlank(value)
	d := bson.E{key, value}
	updateFields = append(updateFields, d)
}

//func GetRelease(releaseId primitive.ObjectID) (*pkg.Release, error) {
//	fmt.Printf("\n   GetRelease\n")
//	filter := bson.D{{"_id", releaseId}}
//	return getReleaseByFilter(filter)
//}

func GetReleaseByForeignId(srcId string) (rel *relMod.Release, err error) {
	rel = &relMod.Release{}
	filter := bson.D{{"src_id", srcId}}
	//fmt.Printf("\nGetReleaseByForeignId\n")
	return getReleaseByFilter(filter)
}

//TODO Should Remove physically remove the release or just market as deleted?
//RemoveRelease Removes a release base upon the provided ObjectID
func RemoveRelease(oid primitive.ObjectID) error {
	filter := bson.M{"_id": oid}
	//fmt.Printf("Using filter: %v\n", filter)
	col, err := GetCollection("releases")
	if err != nil {
		err = fmt.Errorf("Collection Failed|%d|%s", codes.Internal, err.Error())
		return err
	}
	resp, err := col.DeleteOne(context.Background(), filter) // returns result a struct with one element DeletedCount
	if err != nil {
		err = fmt.Errorf("Internal Error|%d|Delete of %s failed err: %v", codes.Internal, oid.Hex(), err.Error())
		return err
	}
	if resp.DeletedCount == 0 {
		err = errors.New("Release not found, nothing removed")
		return err
	}
	return err
}

//########################################### Helper Functions ###########################
func getReleaseByFilterM(filter bson.M) (rel *relMod.Release, err error) {
	rel = &relMod.Release{}
	//fmt.Printf("getReleaseByFilter - Using filter: %v\n", filter)

	collection, err := GetCollection("releases")
	if err != nil {
		return nil, fmt.Errorf("GetRelease-553 GetCollection failed: %v\n", err)
	}
	fmt.Printf("Calling FindOne:537 -- %v\n", filter)
	err = collection.FindOne(context.Background(), filter).Decode(rel)
	if err != nil {
		return nil, fmt.Errorf("GetRelease-557 FindOne [%v] error: %v\n", filter, err)
	}
	fmt.Printf("\n\n$$$$getReleaseByFilter:542 -- %v\n", spew.Sdump(rel))
	return rel, nil
}

func getReleaseByFilter(filter bson.D) (rel *relMod.Release, err error) {
	rel = &relMod.Release{}
	//fmt.Printf("getReleaseByFilter - Using filter: %v\n", filter)
	//spew.Dump(filter)
	collection, err := GetCollection("releases")
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("GetRelease-568 GetCollection failed: %v\n", err),
		)
	}
	//fmt.Printf("Calling FindOne: %v\n", filter)
	err = collection.FindOne(context.Background(), filter).Decode(rel)
	if err != nil {
		//log.Errorf("GetRelease-302 FindOne [%v]   err: %v\n", filter, err)
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("GetRelease-576 FindOne [%v] NotFound\n", filter),
		)
	}
	//fmt.Printf("\n\n$$$$getrelease: %v\n", spew.Sdump(rel))
	return
}

func updateOne(collection mongo.Collection, filter bson.M, data bson.D) {
	_, err := collection.UpdateOne(
		context.TODO(),
		filter,
		data)
	if err != nil {

		//return nil, status.Errorf(
		//	codes.Internal,
		//	fmt.Sprintf("Error update Release using filter: %s:  %v\n", filter, err),
		//)
	}
	return
}

func CreateSampleRelease() (*relMod.Release, error) {
	var err error
	rel := SampleRelease()

	collection, err := GetCollection("releases")
	if err != nil {
		fmt.Printf("GetCollection [releases] failed: %v\n", err)
		return nil, err
	}

	rel.CreatedAt = time.Now()
	rel.UpdatedAt = rel.CreatedAt
	res, err := collection.InsertOne(context.Background(), rel)
	if err != nil {
		log.Errorf("InsertOne failed: %v\n", err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error saving Document: %v", err),
		)
	}

	relId, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot retrieve InsertedId"),
		)
	}
	rel.ID = relId
	err = AddDocument(rel.ID, MakeSampleMetaData())
	if err != nil {
		return rel, fmt.Errorf("unable to add document: %s", err.Error())
	}
	rel, _ = GetRelease(rel.ID)
	return rel, err
}

//func UpdateDocumentCount(id primitive.ObjectID, val int) (*Release, error) {
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//	collection, err := GetCollection(settings.CollectionName())
//	if err != nil {
//		fmt.Printf("Collection failed; %s\n", err)
//		return nil, err
//	}
//	filter := bson.M{"_id": id}
//	updateAt := time.Now()
//	update := bson.M{
//		"$inc": bson.M{"documentcount": val},
//		"$set": bson.M{"updatedat": updateAt},
//	}
//	upsert := true
//	after := options.After
//	opt := options.FindOneAndUpdateOptions{
//		ReturnDocument: &after,
//		Upsert: &upsert,
//	}
//
//	result := collection.FindOneAndUpdate(ctx, filter, update, &opt)
//	if result.Err() != nil {
//		fmt.Printf("Update error: %s\n", result.Err())
//		return nil, result.Err()
//	}
//	doc := bson.M{}
//	rel := Release{}
//	decodeErr := result.Decode(&doc)
//	dErr :=result.Decode(&rel)
//	if dErr != nil {
//		fmt.Printf("Decode to struct failed: %s\n", dErr)
//		return &rel, decodeErr
//	}
//	//fmt.Printf("rel: %s\n", spew.Sdump(rel))
//	return &rel, decodeErr
//
//}

//func GetReleaseBySrcIdFacility(ctx context.Context, srcId, facility string) (*m.Release, error) {
//	if srcId == "" || facility == "" {
//		return nil, status.Errorf(
//			codes.InvalidArgument,
//			fmt.Sprintf("GetReleaseBySrcIdFacility invalid params srcId: [%s], facility: [%s]",srcId, facility),
//		)
//	}
//	filter := bson.D{{"Client.SourceId", srcId}, {"Customer.Facility", facility}}
//	return getReleaseByFilter(filter)
//}

///
/*TODO: Release Update Should change oly allowed fields and not replace entire Release  History should be manually appended
Documents should be updated alone
*/
//Update  replace the current Release with the new one.
//func Update(rel *pkg.Release) error {
//	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	//defer cancel()
//	collection, err := GetCollection("releases")
//	if err != nil {
//		return status.Errorf(
//			codes.Internal,
//			fmt.Sprintf("domain.GetRelease GetCollection failed: %v\n", err),
//		)
//	}
//
//	rel.UpdatedAt = time.Now()
//	filter := bson.M{"_id": rel.ID}
//	_, err = collection.ReplaceOne(context.Background(), filter, rel)
//	if err != nil {
//		return status.Errorf(
//			codes.Internal,
//			fmt.Sprintf("Cannot update Document in database: %v", err),
//		)
//	}
//	return err
//}
//func GetReleaseById(relId primitive.ObjectID) (rel *pkg.Release, err error) {
//	rel = &pkg.Release{}
//	fmt.Printf("\nGetReleaseById\n")
//	filter := bson.D{{"_id", relId}}
//	return getReleaseByFilter(filter)
//}
