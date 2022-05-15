package domain

//import ()
//
//func Unimplemented() {
//
//}

//import (
//	"fmt"
//	log "github.com/sirupsen/logrus"
//	relPB "github.com/dhf0820/ids_release_service/protobufs/relPB"
//	m "github.com/dhf0820./ds_model"
//
//	"go.mongodb.org/mongo-driver/bson/primitive"
//	"regexp"
//	//"time"
//)
//
//
//func ToClient(pc *relPB.Client) *m.Client {
//	c := m.Client{}
//	c.SourceId = pc.GetSourceId()
//	c.Source 	= pc.GetSource()
//	return &c
//}
//
//func ToPbClient(cl *m.Client) *relPB.Client {
//	pc := relPB.Client{}
//	pc.SourceId = cl.SourceId
//	pc.Source 	= cl.Source
//	return &pc
//}
//func ToCustomer(pc *relPB.Customer) *m.Customer {
//	cu := m.Customer{}
//	cu.Name = pc.GetName()
//	cu.UserId = pc.GetUserId()
//	cu.Facility = pc.GetFacility()
//	cu.Code 	= pc.GetCode()
//	return &cu
//}
//
//func ToPbCustomer(cu *m.Customer) *relPB.Customer {
//	pc := relPB.Customer{}
//	pc.Name = cu.Name
//	pc.UserId = cu.UserId
//	pc.Facility = cu.Facility
//	pc.Code 	= cu.Code
//	return &pc
//}
//
//func ToField(f *m.Field) *relPB.Field {
//	pf := relPB.Field{}
//	pf.Name 		= f.Name
//	pf.Label		= f.Label
//	pf.Value		= f.Value
//	pf.Required		= f.Required
//	pf.UserVisible  = f.UserVisible
//	pf.Sensitive 	= f.Sensitive
//	return &pf
//}
//
//
//func ToPbField(pf *relPB.Field) *m.Field {
//	f := m.Field{}
//	f.Name 			= pf.GetName()
//	f.Label			= pf.GetLabel()
//	f.Value			= pf.GetValue()
//	f.Required		= pf.GetRequired()
//	f.UserVisible  	= pf.GetUserVisible()
//	return &f
//}
//
//
////UnMarshal takes a pb.Release and unmarshals it to domain.Release
//func UnMarshal(pbRel *relPB.Release) (*m.Release, error) {
//	var err error
//	//pbRel := frm.GetRelease()
//	rel := m.Release{}
//	rel.Client = ToClient(pbRel.GetClient())
//	rel.Customer = ToCustomer(pbRel.GetCustomer())
//
//
//	rel.Patient.MRN = pbRel.GetPatient().GetMrn()
//	rel.Patient.PatientName = pbRel.GetPatient().GetPatientName()
//	rel.Patient.DOB = pbRel.GetPatient().GetDob()
//	rel.Patient.SSN = pbRel.GetPatient().GetSsn()
//	rel.Documents = ToDocuments(pbRel.Documents)
//	rel.ID, _ = primitive.ObjectIDFromHex(pbRel.GetRelId())
//	//fmt.Printf("FromPB: %s\n", spew.Sdump(rel))
//	return &rel, err
//}
//
////Marshal takes a domain releasse and marshals it to a pb.Release
//func Marshal(rel *m.Release) (*relPB.Release, error) {
//	pRel := relPB.Release{}
//	pRel.Client = ToPbClient(rel.Client)
//	pRel.Customer = ToPbCustomer(rel.Customer)
//	pRel.Patient.PatientName = rel.Patient.PatientName
//	pRel.Patient.Mrn = rel.Patient.MRN
//	pRel.Patient.Dob = rel.Patient.DOB
//	pRel.Patient.Ssn = rel.Patient.SSN
//	pRel.RelId  = rel.ID.Hex()
//
//	pRel.Documents = ToPbDocuments(rel.Documents)
//
//	return &pRel, nil
//}
//
//
////func UnMarshalRecipient(p *relPB.Recipient)  (*Recipient, error) {
////	recip := Recipient{
////		//ID:			"",
////		Name: 	p.GetContact(),
////		CompanyName:  	p.GetCompany(),
////		Address1:   p.GetAddress1(),
////		City:  		p.GetCity(),
////		State:      p.GetState(),
////		Zip: 		p.GetPostal(),
////
////	}
////	//fmt.Printf("FromPBRecipient: %s\n", spew.Sdump(recip))
////	return &recip, nil
////}
//
////func MarshalRecipient(r *Recipient)  (*relPB.Recipient, error) {
////	recip := relPB.Recipient{}
////		recip.Id			=	r.ID.Hex()
////		recip.Contact 	=	r.Name
////		recip.Company 	=  		r.CompanyName
////		recip.Address1 	=   	r.Address1
////		recip.City 		=  			r.City
////		recip.State 	=     		r.State
////		recip.Postal 	= 		r.Zip
////		recip
////
////
////
////	}
////	//fmt.Printf("FromPBRecipient: %s\n", spew.Sdump(recip))
////	return &recip, nil
////}
//
////func UnMarshalRecipientFull(p *relPB.Recipient)  (*Recipient, error) {
////	oid, err := primitive.ObjectIDFromHex(p.GetId())
////	if err != nil {
////		return nil, err
////	}
////	recip := Recipient{
////		ID:  		oid,
////		Contact: 	p.GetContact(),
////		Company:  	p.GetCompany(),
////		Address1:   p.GetAddress1(),
////		Address2:   p.GetAddress2(),
////		City:  		p.GetCity(),
////		State:		p.GetState(),
////		Zip:  		p.GetPostal(),
////		ElectronicID: p.GetElectronicId(),
////		RemoteID:  	p.GetRemoteId(),
////	}
////	return &recip, nil
////}
//
//
//
//
////MarshalDocument converts a *document.Document to a *relPB.Document
//func ToPbDocument(doc *m.Document) *relPB.Document {
//	ds := &relPB.Document{}
//	//ds.Id 			= doc.Id
//	ds.ImageUrl 			= doc.ImageURL
//	ds.DocumentUrl 			= doc.DocumentURL
//	ds.ImageRepository = doc.ImageRepository
//	ds.ImageType 	= doc.ImageType
//	ds.Version		= doc.Version
//	ds.DocClass		= doc.DocClass
//	ds.Description	= doc.Description
//	return ds
//}
//
//
//
////
////func UnMarshalDevice(pDev *relPB.Device) *Device {
////
////	dev := c.DeviceDetail{}
////	dev.Id, _ 			= primitive.ObjectIDFromHex(pDev.GetId())
////	dev.Method 			= pDev.GetMethod()
////	dev.UpdatedAt 		= time.Unix(pDev.UpdatedAt, 0)
////	dev.Validated 		= pDev.GetValidated()
////	vTime  	:= time.Unix(pDev.GetUpdatedAt(), 0)
////	dev.ValidatedAt 	= &vTime
////	for _, pf := range pDev.GetFields() {
////		f := Fields{}
////		f.Name 		= pf.GetName()
////		f.UserVisible 	= pf.GetUserVisible()
////		f.Required 		= pf.GetRequired()
////		f.Label 		= pf.GetLabel()
////		f.Value 		= pf.GetValue()
////		dev.Fields = append(dev.Fields, f)
////	}
////	return &dev
////}
//
////func MarshalDevice(d *Device) *relPB.Device {
////	dev := relPB.Device{}
////	dev.Name			= d.Name
////	dev.Id				= d.Id.Hex()
////	dev.RecipientId		= d.RecipientId.Hex()
////	dev.Method			= d.Method
////	dev.UpdatedAt		= d.UpdatedAt.Unix()
////	dev.ValidatedAt 	= d.ValidatedAt.Unix()
////	dev.Validated		= d.Validated
////	for _, f := range d.Fields {
////		pf := relPB.Fields{}
////		pf.Name = f.Name
////		pf.UserVisible = f.UserVisible
////		pf.Required = f.Required
////		pf.Label = f.Label
////		pf.Value = f.Value
////		dev.Fields = append(dev.Fields, &pf)
////	}
////	return &dev
////}
////
//
////func UnMarshalDevice(pdd *relPB.Device) *Device{
////	dd :=Device{}
////	dd.ID			= pdd.Id
////	dd.DeviceName	= pdd.DeviceName
////	dd.Method		= pdd.Method
////	dd.Command 		= pdd.Command
////	dd.Address		= pdd.Address
////	dd.Secure		= pdd.Secure
////	return &dd
////}
//
////MarshalDocuments converts from the client's Go Struct version of a Document to Protobuf Documentfor over the wire.
//func ToPbDocuments(ds []*m.Document) []*relPB.Document {
//	var pds []*relPB.Document
//	for _, d := range ds {
//
//		pds = append(pds, ToPbDocument(d))
//	}
//	return pds
//}
//
//func ToDocument(doc *relPB.Document) *m.Document {
//	ds := m.Document{}
//	ds.ImageType 	= 	doc.GetImageType()
//	ds.ImageRepository = doc.GetImageRepository()
//	ds.ImageURL 			= doc.GetImageUrl()
//	ds.DocumentURL			= doc.GetDocumentUrl()
//	ds.Description  = doc.GetDescription()
//	ds.DocClass		= doc.GetDocClass()
//	ds.Version		= doc.Version
//	//ds.Id 			= doc.GetId()
//	return &ds
//}
////MarshalDocuments conversts from the client's Go Struct version of a Document to Protobuf Documentfor over the wire.
//func ToDocuments(pds []*relPB.Document) []*m.Document {
//	var dsa []*m.Document
//	for _, d := range pds {
//		dsa = append(dsa, ToDocument(d))
//	}
//	return dsa
//}
//
//
//func NormalizePhone(phone string) string {
//	fmt.Printf("Normalizing : %s\n", phone)
//	reg, err := regexp.Compile("[^0-9]*")
//	if err != nil {
//		log.Errorf("regexp compile failed: %v", err)
//	}
//	return reg.ReplaceAllString(phone, "")
//}
//
//func ValidateEmail(email string) bool {
//	fmt.Printf("Validating : %s\n", email)
//	str:="^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
//	re := regexp.MustCompile(str)
//	return re.MatchString(email)
//}
