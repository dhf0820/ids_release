package client

//
//import (
//	"context"
//	"fmt"
//	"github.com/davecgh/go-spew/spew"
//	//"gitlab.com/dhf0820/ids_delivery_service/protobufs/delPB"
//	"go.mongodb.org/mongo-driver/bson/primitive"
//	//"time"
//	m "github.com/dhf0820./ds_model"
//	//"google.golang.org/grpc/codes"
//	//"google.golang.org/grpc/status"
//	//"github.com/davecgh/go-spew/spew"
//	relPB "gitlab.com/dhf0820/ids_release_service/protobufs/relPB"
//)
//
////type Client struct {
////	Source		string 	`json:"source"`
////	SourceID	string	`json:"source_id"`
////}
////
////type Customer struct {
////	Name 		string	`json:"name"`
////	Code 		int32 	`json:"code"`
////	Facility 	string 	`json:"facility"`
////	UserID 		string 	`json:"user_id"`
////}
////
////type Release struct {
////	Name string // Used by grpc http/1
////	//ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
////	ID         string
////	Client     *Client            `json:"client,omitempty" bson:"client"`
////	Customer   *Customer 		 `json:"customer", bson:"customer"`
////	Facility   string            `json:"facility,omitempty" bson:"facility"`
////	SourceID   string            `json:"source_id,omitempty"`
////	Patient    Patient           `json:"patient,omitempty"`
////	UserID     string            `json:"user_id,omitempty"`
////	Documents  []Document        `json:"documents"`
////	Recipients []DeliveryHistory `json:"history,omitempty" bson:"deliveryhistory"`
////	CreatedAt  *time.Time        `json:"created_at,omitempty"`
////	UpdatedAt  *time.Time        `json:"updated_at,omitempty"`
////}
////
//////Patient contains the patient information and is a subdocument in Release
////type Patient struct {
////	SourceID 	string `json:"source_id" bson:"source_id"`
////	Source 		string `json:"source" bson:"source"`
////	Name     	string `json:"name" bson:"patient_name"`
////	MRN      	string `json:"mrn" bson:"mrn"`
////	DOB      	string `json:"dob" bson: "dob"`
////	SSN      	string `json:"ssn" bson:"ssn"`
////}
////
////type DeliveryHistory struct {
////	ID           primitive.ObjectID `json:"id" bson:"_id"`
////	RecipientID  string          	`json:"recipient_id" bson:"recipient_id"`
////	DeviceID     string     		`json:"device_id" bson:"device_id"`
////	SubmitTime   *time.Time         `json:"submit_time" bson:"submit_time"`
////	DeliveryTime *time.Time         `json:"delivery_time" bson:"delivery_time"`
////	Status       string             `json:"status" bson: "status"` // "STARTED", "PENDING", "QUEUED", "DELIVERED", "ERROR", "FAILED"
////	StatusTime   *time.Time         `json:"status_time" bson:"status_time"`
////	Comments     string             `json:"comments" bson:"comments"`
////}
////
////
////type DeliveryDevice struct {
////	ID     primitive.ObjectID	`json:"id" bson:"_id"`
////	Method string             	`json:"method" bson:"method"`
////	Name   string             	`json:"name" bson: "name"`
////	Validated *time.Time  		`json:"validated", bson:"validated"`
////}
//
/////////////////////////// Service Calls //////////////////////////////
//
//func GetRelease(ctx context.Context, id string) (*m.Release, error) {
//	fmt.Printf("Client GetRelease called\n")
//	if RelClientPtr == nil {
//		fmt.Printf("Document Get is connecting to server")
//		ReleaseConnect()
//	}
//	relReq := relPB.GetReleaseRequest{RelId: id}
//	//fmt.Printf("Calling RelClient.GetClient\n")
//
//	relResp, err := RelClient.GetRelease(ctx, &relReq)
//	if err != nil {
//		err = fmt.Errorf("GetRelease: [%s] failed: err: %v", id, err)
//		return nil, err
//	}
//	pRel := relResp.GetRelease()
//	return ToRelease(pRel), nil
//}
//
////// func CreateReleasePB(ctx context.Context, relReq *relPB.CreateReleaseRequest) (*relPB.CreateReleaseResponse, error) {
////// 	fmt.Printf("Client CreateRelease called\n")
////// 	if RelClientPtr == nil {
////// 		fmt.Printf("Document Get is connecting to server")
////// 		ReleaseConnect()
////// 	}
////
////// 	//relReq := relPB.CreateReleaseRequest{Release: rel}
////// 	relResp, err := RelClient.CreateRelease(ctx, relReq)
////// 	if err != nil {
////// 		return nil, err
////// 	}
////// 	//rel :=relResp.GetRelease()
////// 	//fmt.Printf("relResp: %v\n", spew.Sdump(rel))
////// 	//rel.Id = relId
////// 	return relResp, nil
////// }
////
//func CreateRelease(ctx context.Context, rel *m.Release) (*m.Release, error) {
//	fmt.Printf("Client CreateRelease called\n")
//	if RelClientPtr == nil {
//		fmt.Printf("CreateRelease is connecting to server")
//		ReleaseConnect()
//	}
//	relReq := relPB.CreateReleaseRequest{}
//	relReq.Release = ToPbRelease(rel)
//	//ssn := relReq.Release.GetPatient().GetSsn()
//	//fmt.Printf("\n$$$ relReq.Release.Patient.SSN: %s\n", ssn )
//
//	//relReq := relPB.CreateReleaseRequest{Release: rel}
//	relResp, err := RelClient.CreateRelease(ctx, &relReq)
//	if err != nil {
//		return nil, err
//	}
//	pbRel := relResp.GetRelease()
//	rel = ToRelease(pbRel)
//	//fmt.Printf("relResp: %v\n", spew.Sdump(rel))
//	//rel.Id = relId
//	return rel, nil
//}
//
//
//
//
//
//
//
//func  AddDocument(ctx context.Context, rel *m.Release,doc *m.Document) (*m.Document, error) {
//	if RelClientPtr == nil {
//		fmt.Printf("Document Get is connecting to server")
//		ReleaseConnect()
//	}
//	fmt.Printf("Document ReleaseID: %s\n", rel.ID)
//	//doc.Id = primitive.NewObjectID().Hex()
//	pbDoc := ToPbDocument(doc)
//
//	//fmt.Printf("Client pbDoc: %v\n",spew.Sdump(pbDoc))
//	relReq := relPB.AddDocumentRequest{
//		RelId:  rel.ID.Hex(),
//		Doc: pbDoc,
//	}
//	relResp, err := RelClient.AddDocument(ctx, &relReq)
//	//relResp.GetDocument()
//	//relResp.GetNumDocs()
//	if err != nil {
//		return nil, err
//	}
//	newDoc := ToDocument(relResp.GetDocument())
//	//fmt.Printf("New Document: %s\n", spew.Sdump(newDoc))
//	fmt.Printf("AddDocument response: %s\n",  spew.Sdump(relResp.GetDocument()))
//	return  newDoc,  nil
//}
//
//
/////////////////////////////////////////// Helpers  //////////////////////////////////////
//func ToRelease(c *relPB.Release) *m.Release {
//	r := &m.Release{}
//	r.ID,_ = primitive.ObjectIDFromHex( c.GetRelId())
//	r.Client = m.Client{}
//	r.Client.Source = c.GetClient().GetSource()
//	r.Client.SourceId 	= c.GetClient().GetSourceId()
//	r.Customer = m.Customer{}
//	r.Customer.Name 	= c.GetCustomer().GetName()
//	r.Customer.Facility = c.GetCustomer().GetFacility()
//	r.Customer.Code 	= c.GetCustomer().GetCode()
//	r.Customer.UserId	= c.GetCustomer().GetUserId()
//	r.Patient 			= m.Patient{}
//	r.Patient.PatientName = c.GetPatient().GetPatientName()
//	r.Patient.MRN = c.GetPatient().GetMrn()
//	r.Patient.DOB = c.GetPatient().GetDob()
//	r.Patient.SSN = c.GetPatient().GetSsn()
//	r.Documents = ToDocuments(c.GetDocuments())
//	fmt.Printf("ToRelease: %s\n", spew.Sdump(r))
//	return r
//}
//
//func ToPbRelease(r *m.Release) *relPB.Release {
//	c := &relPB.Release{}
//	c.RelId					= c.RelId
//	clnt 					:= relPB.Client{}
//	clnt.Source				= r.Client.Source
//	clnt.SourceId			= r.Client.SourceId
//	//c.Client 				= &clnt
//	//c.Client.SourceId 		=	r.Client.SourceID
//	//c.Client.Source			= 	r.Client.Source
//	c.Client = &relPB.Client{}
//	c.Client.SourceId		= r.Client.SourceId
//	c.Client.Source			= r.Client.Source
//	c.Customer				= &relPB.Customer{}
//	c.Customer.UserId 		= r.Customer.UserId
//	c.Customer.Code			= r.Customer.Code
//	c.Customer.Facility		= r.Customer.Facility
//	c.Customer.Name			= r.Customer.Name
//	c.Patient 				= &relPB.Patient{}
//	c.Patient.Source		= r.Patient.Source
//	c.Patient.SourceId		= r.Patient.SourceId
//	c.Patient.PatientName	= r.Patient.PatientName
//	c.Patient.Mrn			= r.Patient.MRN
//	c.Patient.Ssn			= r.Patient.SSN
//	c.Patient.Dob			= r.Patient.DOB
//
//	c.Documents = ToPBDocuments(r.Documents)
//	fmt.Printf("\n $$$$  ToPbRelease mrn = %s  source MRN %s\n", c.GetPatient().GetMrn(), r.Patient.MRN)
//	return c
//}
//
//
//func ToDocuments(pds []*relPB.Document) []*m.Document {
//	ds := []*m.Document {}
//	for _, pd := range pds {
//		d := ToDocument(pd)
//		ds = append(ds, d)
//	}
//	return ds
//}
//
//func ToDocument(d *relPB.Document) *m.Document {
//	doc := m.Document{}
//	//doc.Id				= d.GetId()
//	doc.DocumentURL		= d.GetDocumentUrl()
//	doc.ImageURL		= d.GetImageUrl()
//	doc.Description		= d.GetDescription()
//	//doc.DateOfService	= d.GetDateOfService()
//	doc.DocClass		= d.GetDocClass()
//	doc.ImageURL		= d.GetImageUrl()
//	doc.ImageRepository	= d.GetImageRepository()
//	doc.ImageType		= d.GetImageType()
//
//	return &doc
//}
//
//func ToPBDocuments(ds []*m.Document) []*relPB.Document {
//	pds := []*relPB.Document {}
//	for _, rd := range ds {
//		pd := ToPbDocument(rd)
//		pds = append(pds, pd)
//	}
//	return pds
//}
//func ToPbDocument(ds *m.Document) *relPB.Document {
//	pd := relPB.Document{}
//	//pd.DateOfService = ds.DateOfService
//	pd.ImageType		= ds.ImageType
//	pd.ImageRepository	= ds.ImageRepository
//	pd.DocClass			= ds.DocClass
//	pd.ImageUrl			= ds.ImageURL
//	pd.Description 		= ds.Description
//	pd.DocumentUrl		= ds.DocumentURL
//	//pd.Id				= ds.Id
//	pd.Version			= ds.Version
//
//	return &pd
//}
