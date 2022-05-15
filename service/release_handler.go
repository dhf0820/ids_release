package service

import (
	"encoding/json"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/dhf0820./ds_model/common"
	delvMod "github.com/dhf0820./ds_model/delivery"
	docMod "github.com/dhf0820./ds_model/document"
	recipMod "github.com/dhf0820./ds_model/recipMod"
	relMod "github.com/dhf0820./ds_model/release"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"

	//recipMod "github.com/dhf0820./ds_model/recipient"
	"time"

	//dm "gitlab.com/dhf0820/ids_document_service/pkg"
	"io/ioutil"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//type ReleaseResponse struct {
//	Status  int             `json:"status"`
//	Message string          `json:"message"`
//	Release *relMod.Release `json:"release"` // ID of release needed to add documents and submit
//}

//type DocumentResponse struct {
//	Status   int    `json:"status"`
//	Message  string `json:"message"`
//	Document *docMod.Document
//}
//
//type DocumentsResponse struct {
//	Status    int    `json:"status"`
//	Message   string `json:"message"`
//	Documents []*relMod.DocumentMetadata
//}

func addDocument(w http.ResponseWriter, r *http.Request) {
	log.Infof("addDocument Handler is called")
	params := mux.Vars(r)
	fmt.Printf("params: %v\n", params)
	releaseId := params["release_id"]
	fmt.Printf("Received release_id: %s\n", releaseId)
	relId, err := primitive.ObjectIDFromHex(releaseId)
	if err != nil {
		fmt.Printf("Error on fromHex: %vn", err)
		HandleError(w, "addDocument", err)
		return
	}
	//fmt.Printf("\nUpdate ReleaseID: %s\n", relId)
	doc := docMod.Document{}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = fmt.Errorf("400|addDocument read body: %s", err.Error())
		log.Errorf("%s", err.Error())
		HandleError(w, "addDocument", err)
		return
	}
	//fmt.Printf("addDocument Body: %s\n", string(b))

	if err := json.Unmarshal(b, &doc); err != nil {
		err = fmt.Errorf("400|Error marshaling JSON: %s", err.Error())
		log.Errorf("addDocument error: %s", err)
		HandleError(w, "addDocument", err)
		return
	}
	fmt.Printf("\n### Document: %s being added\n\n", doc.ID)
	//fmt.Printf("document: %s\n", spew.Sdump(doc))
	docMeta := common.DocumentMetadata{}
	docMeta.ID = doc.ID
	docMeta.DocClass = common.GetDataByName(doc.Meta, "doc_class")
	docMeta.Description = common.GetDataByName(doc.Meta, "description")
	//docMeta.Status = doc.Status
	//TODO: Lookup the source instead of assuming it is element 0, though has to be.
	srcCor := doc.CorrelationIDs[0]
	docMeta.SrcDocURL = srcCor.SystemDocumentURL
	docMeta.ImageType = doc.ImageType
	docMeta.DateOfService = common.GetDataByName(doc.Meta, "date_of_service")
	docMeta.Meta = doc.Meta

	docEP := common.GetServiceEndpoint(GetConfig().ServiceEndPoints, "document")

	docMeta.DocumentURL = fmt.Sprintf("%s/document/%s", docEP.Address, doc.ID.Hex())

	status := "no_image"
	if doc.ImageID != primitive.NilObjectID {
		status = "has_image"
	}

	stat := common.Status{
		State:      status,
		Comment:    "Image Status",
		StatusTime: time.Now(),
	}
	docMeta.Status = &stat
	docMeta.SrcImageURL = doc.CorrelationIDs[0].SystemImageURL

	resp := relMod.ReleaseResponse{}
	err = AddDocument(relId, &docMeta)
	updRel, err := GetRelease(relId)
	if err != nil {
		resp.Status = 400
		resp.Message = err.Error()
		resp.Release = nil
	} else {
		resp.Status = 201
		resp.Message = "added document to release"
		resp.Release = updRel
	}
	fmt.Printf("\n\n###write response for add document: %s\n\n", doc.ID)
	WriteReleaseResponse(w, resp.Status, &resp)

}

func createRelease(w http.ResponseWriter, r *http.Request) {
	log.Infof("createRelease Handler is called")
	newRel := relMod.Release{}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = fmt.Errorf("400|createRelease read body: %s", err.Error())
		log.Errorf("%s", err.Error())
		HandleError(w, "createRelease", err)
		return
	}
	//fmt.Printf("createRelease Body: %s\n", string(b))
	if err := json.Unmarshal(b, &newRel); err != nil {
		err = fmt.Errorf("400|Error marshaling JSON: %s", err.Error())
		log.Errorf("createRelease error: %s", err)
		HandleError(w, "createRelease", err)
		return
	}
	nRel, err := CreateRelease(&newRel)
	if err != nil || nRel == nil {
		fmt.Printf("Error creatingRelease: %v\n", err)
		HandleError(w, "createRelease-65", err)
		return
	}

	resp := relMod.ReleaseResponse{}
	resp.Status = 201
	resp.Message = "Added Release"
	resp.Release = nRel
	fmt.Printf("resp: %s\n", spew.Sdump(resp))
	WriteReleaseResponse(w, 201, &resp)
}

func deleteRelease(w http.ResponseWriter, r *http.Request) {
	log.Infof("deleteRelease Handler is called")
	params := mux.Vars(r)
	fmt.Printf("params: %v\n", params)
	id := params["release_id"]
	fmt.Printf("Received release_id: %s\n", id)
	relId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("Error on fromHex: %vn", err)
		HandleError(w, "getRelease", err)
		return
	}
	fmt.Printf("Received relId: %s\n", relId)
	err = RemoveRelease(relId)
	if err != nil {
		err = fmt.Errorf("400|getRelease read body: %s", err.Error())
		log.Errorf("%s", err.Error())
		HandleError(w, "getRelease", err)
		return
	}

	resp := common.StatusResponse{
		Status:  200,
		Message: "Ok",
	}
	WriteStatusResponse(w, 200, &resp)

}

func getDocuments(w http.ResponseWriter, r *http.Request) {
	log.Infof("getDocuments Handler is called")
	params := mux.Vars(r)
	fmt.Printf("params: %v\n", params)
	id := params["release_id"]
	fmt.Printf("Received release_id: %s\n", id)
	relId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("Error on fromHex: %vn", err)
		HandleError(w, "getRelease", err)
		return
	}
	documents, err := GetDocuments(relId)
	if err != nil {
		err = fmt.Errorf("400|getDocuments read body: %s", err.Error())
		log.Errorf("%s", err.Error())
		HandleError(w, "getDocuments", err)
		return
	}
	resp := docMod.DocumentMetaResponse{
		Documents: documents,
		Status:    200,
		Message:   "Ok",
	}
	WriteDocumentsResponse(w, 200, &resp)

}

func getRelease(w http.ResponseWriter, r *http.Request) {
	log.Infof("getRelease Handler is called")
	params := mux.Vars(r)
	fmt.Printf("params: %v\n", params)
	id := params["release_id"]
	fmt.Printf("Received release_id: %s\n", id)
	relId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("Error on fromHex: %vn", err)
		HandleError(w, "getRelease", err)
		return
	}
	fmt.Printf("Received relId: %s\n", relId)
	release, err := GetRelease(relId)
	if err != nil {
		err = fmt.Errorf("400|getRelease read body: %s", err.Error())
		log.Errorf("%s", err.Error())
		HandleError(w, "getRelease", err)
		return
	}

	resp := relMod.ReleaseResponse{
		Release: release,
		Status:  200,
		Message: "Ok",
	}
	fmt.Printf("Release returning: %s\n", spew.Sdump(release))
	WriteReleaseResponse(w, 200, &resp)
}

func getReleaseSet(w http.ResponseWriter, r *http.Request) {
	log.Infof("getReleaseSet Handler is called")
	params := mux.Vars(r)
	fmt.Printf("params: %v\n", params)
	id := params["release_id"]
	fmt.Printf("Received release_id: %s\n", id)
	relId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("Error on fromHex: %vn", err)
		HandleError(w, "getReleaseSet", err)
		return
	}
	fmt.Printf("Received relId: %s\n", relId)
	documentSet, err := BuildDocumentSet(relId)
	if err != nil {
		err = fmt.Errorf("400|getReleaseSet read body: %s", err.Error())
		log.Errorf("%s", err.Error())
		HandleError(w, "getReleaseSet", err)
		return
	}

	resp := relMod.ReleaseResponse{
		DocumentSet: documentSet,
		Status:      200,
		Message:     "Ok",
	}
	WriteReleaseResponse(w, 200, &resp)
}

func getReleaseCombined(w http.ResponseWriter, r *http.Request) {
	log.Infof("getReleaseCombined Handler is called")
	params := mux.Vars(r)
	fmt.Printf("params: %v\n", params)
	id := params["release_id"]
	fmt.Printf("Received release_id: %s\n", id)
	relId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("Error on fromHex: %vn", err)
		HandleError(w, "getReleaseCombined", err)
		return
	}
	fmt.Printf("Received relId: %s\n", relId)
	image, err := GetReleaseCombined(relId)
	if err != nil {
		err = fmt.Errorf("400|getReleaseCombined read body: %s", err.Error())
		log.Errorf("%s", err.Error())
		HandleError(w, "getReleaseCombined", err)
		return
	}

	resp := relMod.ReleaseResponse{
		ReleaseCombined: *image,
		Status:          200,
		Message:         "Ok",
	}
	//fmt.Printf("Release returning: %s\n", spew.Sdump(release))
	WriteReleaseResponse(w, 200, &resp)
}

func getReleaseImage(w http.ResponseWriter, r *http.Request) {
	log.Infof("getRelease Handler is called")
	params := mux.Vars(r)
	fmt.Printf("params: %v\n", params)
	id := params["release_id"]
	fmt.Printf("Received Image Request for release_id: %s\n", id)
	relId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("Error on fromHex: %vn", err)
		HandleError(w, "getReleaseImage", err)
		return
	}
	//fmt.Printf("Received relId: %s\n", relId)
	release, err := GetRelease(relId)
	if err != nil {
		err = fmt.Errorf("400|getRelease read body: %s", err.Error())
		log.Errorf("%s", err.Error())
		HandleError(w, "getRelease", err)
		return
	}

	// image, err := GetReleaseImage(release)
	resp := relMod.ReleaseResponse{
		Release: release,
		Status:  200,
		Message: "Ok",
	}
	fmt.Printf("Release returning: %s\n", spew.Sdump(release))
	WriteReleaseResponse(w, 200, &resp)
}

func getTestRelease(w http.ResponseWriter, r *http.Request) {
	log.Infof("getTestRelease Handler is called")
	release, err := CreateSampleRelease()
	if err != nil {
		log.Errorf("getTestRelease error: %s", err)
		HandleError(w, "getTestRelease", err)
	}
	resp := relMod.ReleaseResponse{
		Release: release,
		Status:  200,
		Message: "Ok",
	}
	WriteReleaseResponse(w, 200, &resp)
}

var decoder = schema.NewDecoder()

type requestParams struct {
	ReleaseId   string `schema:"release, required"`
	RecipientId string `schema:"recipient, required"`
	DeviceId    string `schema:"device, required"`
}

func submitRelease(w http.ResponseWriter, r *http.Request) {
	resp := delvMod.DeliveryResponse{}
	fmt.Printf("\n##########submitRelease:265 -- called\n")
	log.Infof("submitRelease Handler is called")
	if err := r.ParseForm(); err != nil {
		msg := fmt.Sprintf("Error parsing Query: %s", err.Error())
		log.Println(msg)
		resp.Status = 400
		resp.Message = msg
		WriteDeliveryResponse(w, 400, msg, nil)
		err := WriteGenericResponse(w, 400, "Error parsing query: "+err.Error())
		if err != nil {
			log.Println("Error writing response: ", err.Error())
		}
		return
	}
	var decoder = schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	filter := new(requestParams)

	if err := decoder.Decode(filter, r.Form); err != nil {
		log.Println("Error decoding query: ", err.Error())
		err := WriteGenericResponse(w, 400, "Error decoding query: "+err.Error())
		if err != nil {
			log.Println("Error writing response: ", err.Error())
			return
		}
		return
	}
	fmt.Printf("filter: %s\n", spew.Sdump(filter))
	releaseID, err := primitive.ObjectIDFromHex(filter.ReleaseId)
	if err != nil {
		fmt.Printf("Error on release fromHex %s: %v\n", filter.ReleaseId, err)
		HandleError(w, "submitRelease", err)
		return
	}
	recipientID, err := primitive.ObjectIDFromHex(filter.RecipientId)
	if err != nil {
		fmt.Printf("Error on recipient fromHex %s: %v\n", filter.RecipientId, err)
		HandleError(w, "submitRelease", err)
		return
	}
	deviceID, err := primitive.ObjectIDFromHex(filter.DeviceId)
	if err != nil {
		fmt.Printf("Error on device fromHex %s: %v\n", filter.DeviceId, err)
		HandleError(w, "submitRelease", err)
		return
	}

	// Validate parameters
	release, err := GetRelease(releaseID)
	if err != nil {
		msg := fmt.Sprintf("Submit GetRelease:305 failed: %s\n", err.Error())
		WriteDeliveryResponse(w, 400, msg, nil)
		return
	}

	recipient, err := GetRecipient(recipientID)
	if err != nil {
		msg := fmt.Sprintf("Submit GetRecipient:316 failed: %s\n", err.Error())
		WriteDeliveryResponse(w, 400, msg, nil)
		return
	}

	device, err := GetDevice(deviceID)
	if err != nil {
		msg := fmt.Sprintf("Submit GetRecipient:316 failed: %s\n", err.Error())
		WriteDeliveryResponse(w, 400, msg, nil)
		return
	}

	//callBack := "http://localhost29911//api/rest/v1/log_status"

	//delv := delvMod.Delivery{}

	delv, err := SubmitToDelivery(release, recipient, device) //, callBack)
	if err != nil {
		fmt.Printf("Submit failed with error: %s\n", err.Error())
		return
	}
	// resp = delvMod.DeliveryResponse{
	// 	Delivery: delv,
	// 	Status:   200,
	// 	Message:  "queued",
	// }
	WriteDeliveryResponse(w, 200, "queued", delv)
	//params := mux.Vars(r)
	//fmt.Printf("params: %s\n", spew.Sdump(params))
	//err := r.ParseForm()
	//if err != nil {
	//	fmt.Printf("Error on parameters: %vn", err)
	//	HandleError(w, "submitRelease", err)
	//	return
	//}
	//var req requestParams
	//err = decoder.Decode(&req, r.PostForm)
	//if err != nil {
	//	fmt.Printf("Error on decoding parameters: %vn", err)
	//	HandleError(w, "submitRelease", err)
	//	return
	//}
	//fmt.Printf("SubmitParams: %sn", spew.Sdump(req))
	//releaseID, err := primitive.ObjectIDFromHex(req.ReleaseId)
	//if err != nil {
	//	fmt.Printf("Error on release fromHex %s: %v\n", req.ReleaseId, err)
	//	HandleError(w, "submitRelease", err)
	//	return
	//}
	//recipientID, err := primitive.ObjectIDFromHex(req.RecipientId)
	//if err != nil {
	//	fmt.Printf("Error on recipient fromHex %s: %v\n", req.RecipientId, err)
	//	HandleError(w, "submitRelease", err)
	//	return
	//}
	//deviceID, err := primitive.ObjectIDFromHex(req.DeviceId)
	//if err != nil {
	//	fmt.Printf("Error on device fromHex %s: %v\n", req.DeviceId, err)
	//	HandleError(w, "submitRelease", err)
	//	return
	//}
	////params := mux.Vars(r)
	////fmt.Printf("params: %v\n", params)
	////releaseId := params["release_id"]
	////fmt.Printf("Received release_id: %s\n", releaseId)
	////releaseID, err := primitive.ObjectIDFromHex(releaseId)
	////if err != nil {
	////	fmt.Printf("Error on fromHex: %vn", err)
	////	HandleError(w, "submitRelease", err)
	////	return
	////}
	////
	//fmt.Printf("\nSubmit ReleaseID: %s\n", relId)\

	//req := common.SubmitDeliveryRequest{}
	//b, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	err = fmt.Errorf("400|submitRelease read body: %s", err.Error())
	//	log.Errorf("%s", err.Error())
	//	HandleError(w, "submitRelease", err)
	//	return
	//}
	//fmt.Printf("submitRelease Body: %s\n", string(b))
	//
	//if err := json.Unmarshal(b, &req); err != nil {
	//	err = fmt.Errorf("400|Error marshaling JSON: %s", err.Error())
	//	log.Errorf("submitRelease error: %s", err)
	//	HandleError(w, "submitRelease", err)
	//	return
	//}
	////callBack := "http://localhost29911//api/rest/v1/log_status"
	//delv, err := SubmitToDelivery(req.ReleaseId, req.RecipientId, req.DeviceId, req.CallBack)
	//resp := delvMod.DeliveryResponse{
	//	Delivery: delv,
	//	Status:   201,
	//	Message:  "queued",
	//}
	//WriteDeliveryResponse(w, 201, &resp)

}

func addRecipient(w http.ResponseWriter, r *http.Request) {
	log.Infof("addRecipient Handler is called")
	params := mux.Vars(r)
	fmt.Printf("params: %v\n", params)
	releaseString := params["release_id"]
	recipientString := params["recipient_id"]
	releaseId, err := primitive.ObjectIDFromHex(releaseString)
	if err != nil {
		err = fmt.Errorf("400|releiase_id invalid %s", err.Error())
		log.Errorf("%s", err.Error())
		HandleError(w, "addRecipient", err)
		return
	}
	recipientId, err := primitive.ObjectIDFromHex(recipientString)
	if err != nil {
		err = fmt.Errorf("400|recipient_id invalid %s", err.Error())
		log.Errorf("%s", err.Error())
		HandleError(w, "addRecipient", err)
		return
	}
	fmt.Printf("Received release_id: %s and recipient_id: %s\n", releaseId, recipientId)
	release, err := GetRelease(releaseId)
	if err != nil {
		err = fmt.Errorf("400|release was not found %s", err.Error())
		log.Errorf("%s", err.Error())
		HandleError(w, "addRecipient", err)
		return
	}
	if err != nil {
		err = fmt.Errorf("400|recipient was not found %s", err.Error())
		log.Errorf("%s", err.Error())
		HandleError(w, "addRecipient", err)
		return
	}
	recipient, err := GetRecipient(recipientId)
	recipSum := recipMod.SummaryFromRecipient(recipient)
	//recipSum.ID = recipient.ID
	//recipSum.LastName = recipient.LastName
	//recipSum.FirstName = recipient.FirstName
	//recipSum.Name = recipient.Name
	//recipSum.CompanyName = recipient.CompanyName

	//TODO: HAndle url query to indicate to update. If already one error unless update set
	release.Recipient = recipSum
	release.Recipients = append(release.Recipients, recipSum)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = fmt.Errorf("400|updateRecipients read body: %s", err.Error())
		log.Errorf("%s", err.Error())
		HandleError(w, "updateRecipients", err)
		return
	}
	fmt.Printf("updateRecipients Body: %s\n", string(b))
	if err := json.Unmarshal(b, &recipient); err != nil {
		err = fmt.Errorf("400|Error marshaling JSON: %s", err.Error())
		log.Errorf("updateRecipients error: %s", err)
		HandleError(w, "updateRecipients", err)
		return
	}
	nRel, err := UpdateRecipients(releaseId, recipient)
	if err != nil || nRel == nil {
		fmt.Printf("Error creatingRelease: %v\n", err)
		HandleError(w, "createRelease-65", err)
		return
	}
	resp := relMod.ReleaseResponse{}
	resp.Status = 200
	resp.Message = "Updated Release"
	resp.Release = nRel
	//fmt.Printf("resp: %s\n", spew.Sdump(resp))
	WriteReleaseResponse(w, 200, &resp)
}

func updateRecipients(w http.ResponseWriter, r *http.Request) {
	log.Infof("updateRecipients Handler is called")
	params := mux.Vars(r)
	fmt.Printf("params: %v\n", params)
	releaseString := params["release_id"]
	releaseId, err := primitive.ObjectIDFromHex(releaseString)
	fmt.Printf("Received release_id: %s\n", releaseId)
	recipient := recipMod.Recipient{}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = fmt.Errorf("400|updateRecipients read body: %s", err.Error())
		log.Errorf("%s", err.Error())
		HandleError(w, "updateRecipients", err)
		return
	}
	fmt.Printf("updateRecipients Body: %s\n", string(b))
	if err := json.Unmarshal(b, &recipient); err != nil {
		err = fmt.Errorf("400|Error marshaling JSON: %s", err.Error())
		log.Errorf("updateRecipients error: %s", err)
		HandleError(w, "updateRecipients", err)
		return
	}
	nRel, err := UpdateRecipients(releaseId, &recipient)
	if err != nil || nRel == nil {
		fmt.Printf("Error creatingRelease: %v\n", err)
		HandleError(w, "createRelease-65", err)
		return
	}
	resp := relMod.ReleaseResponse{}
	resp.Status = 200
	resp.Message = "Updated Release"
	resp.Release = nRel
	//fmt.Printf("resp: %s\n", spew.Sdump(resp))
	WriteReleaseResponse(w, 200, &resp)
}

func updateRecipient(w http.ResponseWriter, r *http.Request) {
	log.Infof("updateRecipient Handler is called")
	params := mux.Vars(r)
	fmt.Printf("params: %v\n", params)
	releaseString := params["release_id"]
	releaseId, err := primitive.ObjectIDFromHex(releaseString)
	fmt.Printf("Received release_id: %s\n", releaseId)
	recipient := recipMod.Recipient{}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = fmt.Errorf("400|updateRecipients read body: %s", err.Error())
		log.Errorf("%s", err.Error())
		HandleError(w, "updateRecipients", err)
		return
	}
	fmt.Printf("updateRecipients Body: %s\n", string(b))
	if err := json.Unmarshal(b, &recipient); err != nil {
		err = fmt.Errorf("400|Error marshaling JSON: %s", err.Error())
		log.Errorf("updateRecipients error: %s", err)
		HandleError(w, "updateRecipients", err)
		return
	}
	nRel, err := UpdateRecipients(releaseId, &recipient)
	if err != nil || nRel == nil {
		fmt.Printf("Error creatingRelease: %v\n", err)
		HandleError(w, "createRelease-65", err)
		return
	}
	resp := relMod.ReleaseResponse{}
	resp.Status = 200
	resp.Message = "Updated Release"
	resp.Release = nRel
	//fmt.Printf("resp: %s\n", spew.Sdump(resp))
	WriteReleaseResponse(w, 200, &resp)
}

//#################################### Response Writers ###########################

func WriteDeliveryResponse(w http.ResponseWriter, status int, msg string, delivery *delvMod.Delivery) error {
	w.Header().Set("Content-Type", "application/json")
	resp := delvMod.DeliveryResponse{}
	resp.Status = 400
	resp.Message = msg
	resp.Delivery = delivery
	switch status {
	case 200:
		w.WriteHeader(http.StatusOK)
	case 400:
		w.WriteHeader(http.StatusBadRequest)
	case 401:
		w.WriteHeader(http.StatusUnauthorized)
	case 403:
		w.WriteHeader(http.StatusForbidden)
	}
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return err
	}
	return nil
}

func WriteDocumentsResponse(w http.ResponseWriter, status int, resp *docMod.DocumentMetaResponse) error {
	w.Header().Set("Content-Type", "application/json")
	switch status {
	case 200:
		w.WriteHeader(http.StatusOK)
	case 400:
		w.WriteHeader(http.StatusBadRequest)
	case 401:
		w.WriteHeader(http.StatusUnauthorized)
	case 403:
		w.WriteHeader(http.StatusForbidden)
	}
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return err
	}
	return nil
}

func WriteDocumentResponse(w http.ResponseWriter, status int, resp *docMod.DocumentResponse) error {
	w.Header().Set("Content-Type", "application/json")

	switch status {
	case 200:
		w.WriteHeader(http.StatusOK)
	case 400:
		w.WriteHeader(http.StatusBadRequest)
	case 401:
		w.WriteHeader(http.StatusUnauthorized)
	case 403:
		w.WriteHeader(http.StatusForbidden)
	}
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return err
	}
	return nil
}

func WriteReleaseResponse(w http.ResponseWriter, status int, resp *relMod.ReleaseResponse) error {
	w.Header().Set("Content-Type", "application/json")

	switch status {
	case 200:
		w.WriteHeader(http.StatusOK)
	case 400:
		w.WriteHeader(http.StatusBadRequest)
	case 401:
		w.WriteHeader(http.StatusUnauthorized)
	case 403:
		w.WriteHeader(http.StatusForbidden)
	}
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return err
	}
	return nil
}

func WriteStatusResponse(w http.ResponseWriter, status int, resp *common.StatusResponse) error {
	w.Header().Set("Content-Type", "application/json")

	switch status {
	case 200:
		w.WriteHeader(http.StatusOK)
	case 400:
		w.WriteHeader(http.StatusBadRequest)
	case 401:
		w.WriteHeader(http.StatusUnauthorized)
	case 403:
		w.WriteHeader(http.StatusForbidden)
	}
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return err

	}
	return nil
}

//fmt.Printf("image: %v\n",spew.Sdump(newUpload) )
// token := r.Header.Get("AUTHORIZATION")
// log.Debugf("Handler token: %s", token)
//fmt.Printf("\n\nCreating new recipient: %s\n", spew.Sdump(newRecip))

//_, err = http.Post("http://localhost:29900/api/v1/release/27/document", "application/json",
//	bytes.NewBuffer(body))
// err = usr.Create(token)
// if err != nil {
// 	log.Errorf("Create User failed: %s", err.Error())
// 	HandleError(w, "Create", err)
// 	return
// }

//Check if the recipient already exists. There should only be one.
// if there were more than one found return error and array of what was found. Otherwise return the exissting one.
// filter := m.SearchRecipientsValues{}
// filter.Source = newRecip.Client.Source
// filter.Customer = newRecip.Customer.Code
// filter.ForeignId = newRecip.Client.SourceId
// filter.Facility = newRecip.Customer.Facility
// ctx := context.Background() //TODO: need to set timeout on create
// recips, err := service.FindRecipients(ctx, &filter)
// if err == nil {
// 	if len(recips) > 0 {
// 		resp := SummaryResponse{}
// 		resp.Recipient = recips[0]
// 		resp.Status = 409
// 		resp.Message = "Already exists"
// 		WriteRecipientSummaryResponse(w, 200, &resp)
// 		return
// 	}
// }

// recip, err := service.CreateRecipient(ctx, &newRecip)
// resp := SummaryResponse{}
// if err != nil {
// 	if err.Error() != "Recipient Exists" {
// 		fmt.Printf("Create error: %v\n", err)
// 		respondWithError(w, 500, "err.Error()")
// 		return
// 	} else {
// 		resp.Message = "Recipient Exist "
// 	}
// } else {
// 	resp.Message = "Created"
// }
// resp.Recipient = recip
// WriteRecipientSummaryResponse(w, 200, &resp)
