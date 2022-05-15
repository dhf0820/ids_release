package service

//
//import (
//	"fmt"
//	"time"
//	"github.com/davecgh/go-spew/spew"
//	log "github.com/sirupsen/logrus"
//	//"gitlab.com/dhf0820/ids_release_service/internal/domain"
//	relPB "gitlab.com/dhf0820/ids_release_service/protobufs/relPB"
//	delPB "gitlab.com/dhf0820/ids_delivery_service/protobufs/delPB"
//	//c "gitlab.com/dhf0820/ids_release_service/pkg/common"
//	mod "github.com/dhf0820./ds_model"
//	"go.mongodb.org/mongo-driver/bson/primitive"
//	"regexp"
//	//"time"
//)
//
//
//func ToClient(pc *relPB.Client) *mod.Client {
//	c :=mod.Client{}
//	c.SourceId = pc.GetSourceId()
//	c.Source 	= pc.GetSource()
//	return &c
//}
//
//func ToPbClient(cl *mod.Client) *relPB.Client {
//	pc := relPB.Client{}
//	pc.SourceId = cl.SourceId
//	pc.Source 	= cl.Source
//	return &pc
//}
//func ToCustomer(pc *relPB.Customer) *mod.Customer {
//	cu := &mod.Customer{}
//	cu.Name = pc.GetName()
//	cu.UserId = pc.GetUserId()
//	cu.Facility = pc.GetFacility()
//	cu.Code 	= pc.GetCode()
//	return cu
//}
//
//func ToPbCustomer(cu *mod.Customer) *relPB.Customer {
//	pc := relPB.Customer{}
//	pc.Name = cu.Name
//	pc.UserId = cu.UserId
//	pc.Facility = cu.Facility
//	pc.Code 	= cu.Code
//	return &pc
//}
//
//func ToField(f *mod.Field) *relPB.Field {
//	pf := relPB.Field{}
//	pf.Name 		= f.Name
//	pf.Label		= f.Label
//	pf.Value		= f.Value
//	pf.UserVisible  = f.UserVisible
//	pf.Required		= f.Required
//	pf.Sensitive	= f.Sensitive
//	return &pf
//
//}
//
//
//func ToPbField(pf *relPB.Field) *mod.Field {
//	f := mod.Field{}
//	f.Name 			= pf.GetName()
//	f.Label			= pf.GetLabel()
//	f.Value			= pf.GetValue()
//	f.Required		= pf.GetRequired()
//	f.UserVisible  	= pf.GetUserVisible()
//	f.Sensitive 	= pf.GetSensitive()
//	return &f
//}
//
//
////UnMarshal takes a pb.Release and unmarshals it to domain.Release
//func ToRelease(pbRel *relPB.Release) *mod.Release {
//
//	//pbRel := frm.GetRelease()
//	rel := mod.Release{}
//	rel.Client = ToClient(pbRel.GetClient())
//	rel.Customer = ToCustomer(pbRel.GetCustomer())
//
//	fmt.Printf("Patient: %v\n", pbRel.GetPatient().GetMrn())
//	rel.Patient = &mod.Patient{}
//	rel.Patient.MRN = pbRel.GetPatient().GetMrn()
//	rel.Patient.PatientName = pbRel.GetPatient().GetPatientName()
//	rel.Patient.DOB = pbRel.GetPatient().GetDob()
//	rel.Patient.SSN = pbRel.GetPatient().GetSsn()
//	rel.Documents 	= RelPbToDocuments(pbRel.GetDocuments())
//	rel.ID, _ = primitive.ObjectIDFromHex(pbRel.GetRelId())
//	//fmt.Printf("FromPB: %s\n", spew.Sdump(rel))
//
//	return &rel
//}
//
////Marshal takes a domain release andmmod.rshals it to a pb.Release
//func ToPbRelease(rel *mod.Release) *relPB.Release {
//	fmt.Printf("ToPbReleaserelease: %s\n", spew.Sdump(rel))
//	fmt.Printf("\n\nPatient Name: [%s]\n", rel.Patient.PatientName)
//	pRel := relPB.Release{}
//	pat := &relPB.Patient{}
//	pRel.Client = ToPbClient(rel.Client)
//	pRel.Customer = ToPbCustomer(rel.Customer)
//	pRel.Patient = pat
//	pRel.Patient.PatientName = rel.Patient.PatientName
//	pRel.Patient.Mrn = rel.Patient.MRN
//	pRel.Patient.Dob = rel.Patient.DOB
//	pRel.Patient.Ssn = rel.Patient.SSN
//	pRel.RelId  = rel.ID.Hex()
//	pRel.Documents = ToPbDocuments(rel.Documents)
//	pRel.CreatedAt = rel.CreatedAt.Unix()
//	pRel.UpdatedAt = rel.UpdatedAt.Unix()
//
//	return &pRel
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
////funcmmod.rshalRecipient(r *Recipient)  (*relPB.Recipient, error) {
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
////ToDelivery accepts the PBDelivery format and converts it to domain so the domain can us it.
//func ToDelivery(pbDev *delPB.Delivery) *mod.Delivery {
//
//	delv := mod.Delivery{}
//	delv.Client	 			= &mod.Client{}
//	delv.Customer 			= &mod.Customer{}
//	delv.ID, _ 				= primitive.ObjectIDFromHex(  pbDev.GetId())
//	delv.Client.SourceId 	= pbDev.GetClient().GetSourceId()
//	delv.Client.Source	 	= pbDev.GetClient().GetSource()
//	delv.Customer.UserId 	= pbDev.GetCustomer().GetUserId()
//	delv.Customer.Name 	 	= pbDev.GetCustomer().GetName()
//	delv.Customer.Code   	= pbDev.GetCustomer().GetCode()
//	delv.Customer.Facility	= pbDev.GetCustomer().GetFacility()
//	delv.ReleaseId 			= pbDev.GetReleaseId()
//	delv.RecipientId		= pbDev.RecipientId
//	delv.Status				= ToStatus(pbDev.GetStatus())
//	delv.Documents 			= DelPbToDocuments(pbDev.GetDocuments())
//	delv.Attempts			= pbDev.GetAttempts()
//	delv.NextTryTime 		= time.Unix(pbDev.GetNextTryTime(), 0)
//	delv.TimeToLive			= time.Unix(pbDev.GetTimeToLive(), 0)
//	delv.Priority			= pbDev.GetPriority()
//	fmt.Printf("ToDelivery: %s\n", spew.Sdump(delv))
//	return &delv
//}
//
////ToPbDelivery accepts a domain delivery struct and puts it into the PB.delivery format
//func ToDelPbDelivery(d *mod.Delivery) *delPB.Delivery {
//	if d == nil {
//		fmt.Printf("Delivery marshal is nil\n")
//	}
//	fmt.Printf("ToPBDelivery - 55: %s\n", spew.Sdump(d))
//	pd := delPB.Delivery{}
//	pd.Client 					= &delPB.Client{}
//	pd.Customer					= &delPB.Customer{}
//	pd.Id 						= d.ID.Hex()
//	pd.Client.Source			= d.Client.Source
//	pd.Client.SourceId			= d.Client.SourceId
//	pd.Customer.UserId			= d.Customer.UserId
//	pd.Customer.Name			= d.Customer.Name
//	pd.Customer.Code			= d.Customer.Code
//	pd.Customer.Facility 		= d.Customer.Facility
//	pd.ReleaseId 				= d.ReleaseId
//	pd.Status 					= ToDelPbStatus(d.Status)
//	pd.Documents 				= ToDelPbDocuments(d.Documents)
//	pd.Attempts					= d.Attempts
//	pd.NextTryTime				= d.NextTryTime.Unix()
//	pd.TimeToLive				= d.TimeToLive.Unix()
//	pd.Priority					= d.Priority
//	fmt.Printf("ToPbDelivery - 80: %v\n", pd)
//	return &pd
//}
//
//func ToStatus(ps *delPB.Status) *mod.Status {
//	s := mod.Status{}
//	s.Status 	= ps.GetStatus()
//	s.StatusTime 	= time.Unix(ps.GetStatusTime(), 0)
//	s.Comment		= ps.GetComment()
//	return &s
//}
//
////func ToHistory(ph []*delPB.Status) []*mod.Status {
////	sh := []*mod.Status{}
////	for _, ps := range ph {
////		sh = append(sh, ToStatus(ps))
////	}
////	return sh
////}
//
//func DelPbToStatus(ps *delPB.Status) *mod.Status {
//	s := mod.Status{}
//	s.Status 	= ps.GetStatus()
//	s.StatusTime 	= time.Unix(ps.GetStatusTime(), 0)
//	s.Comment		= ps.GetComment()
//	return &s
//}
//
////func DelPbToHistory(ph []*delPB.Status) []*mod.Status {
////	sh := []*mod.Status{}
////	for _, ps := range ph {
////		sh = append(sh,DelPbToStatus(ps))
////	}
////	return sh
////}
//
//func ToDelPbStatus(s *mod.Status) *delPB.Status {
//	ps := delPB.Status{}
//	ps.Status 		= s.Status
//	ps.StatusTime 	= s.StatusTime.Unix()
//	ps.Comment		= s.Comment
//	return &ps
//}
//
//func ToDelPbHistory(sh []*mod.Status) []*delPB.Status {
//	psh := []*delPB.Status{}
//	for _, ps := range sh {
//		psh = append(psh, ToDelPbStatus(ps))
//	}
//	return psh
//}
//
//
//
////func RelPbToDelivery(pbDev *relPB.Delivery) *m.Delivery {
////
////	delv :=mmod.Delivery{}
////	delv.Client	 			= &m.Client{}
////	delv.Customer 			= &m.Customer{}
////	delv.ID, _ 				= primitive.ObjectIDFromHex(  pbDev.GetId())
////	delv.Client.SourceId 	= pbDev.GetClient().GetSourceId()
////	delv.Client.Source	 	= pbDev.GetClient().GetSource()
////	delv.Customer.UserId 	= pbDev.GetCustomer().GetUserId()
////	delv.Customer.Name 	 	= pbDev.GetCustomer().GetName()
////	delv.Customer.Code   	= pbDev.GetCustomer().GetCode()
////	delv.Customer.Facility	= pbDev.GetCustomer().GetFacility()
////	delv.ReleaseId  		= pbDev.GetReleaseId()
////	delv.DeviceId			= pbDev.GetId()
////	delv.Status				= RelPbToStatus(pbDev.GetStatus())
////	delv.Documents 			= RelPbToDocuments(pbDev.GetDocuments())
////	delv.Attempts			= pbDev.GetAttempts()
////	delv.NextTryTime 		= time.Unix(pbDev.GetNextTryTime(), 0)
////	delv.TimeToLive			= time.Unix(pbDev.GetTimeToLive(), 0)
////	delv.Priority			= pbDev.GetPriority()
////	fmt.Printf("ToDelivery: %s\n", spew.Sdump(delv))
////	return &delv
////}
//
////ToPbDelivery accepts a domain delivery struct and puts it into the PB.delivery format
////func ToRelPbDelivery(d *m.Delivery) *relPB.Delivery {
////	if d == nil {
////		fmt.Printf("Delivery inmmod.rshal is nil\n")
////	}
////	fmt.Printf("ToPBDelivery - 55: %s\n", spew.Sdump(d))
////	pd := relPB.Delivery{}
////	pd.Client 					= &relPB.Client{}
////	pd.Customer					= &relPB.Customer{}
////	pd.Id 						= d.ID.Hex()
////	pd.Client.Source			= d.Client.Source
////	pd.Client.SourceId			= d.Client.SourceId
////	pd.Customer.UserId			= d.Customer.UserId
////	pd.Customer.Name			= d.Customer.Name
////	pd.Customer.Code			= d.Customer.Code
////	pd.Customer.Facility 		= d.Customer.Facility
////	pd.ReleaseId 				= d.ReleaseId
////	pd.DeviceId					= d.DeviceId
////	pd.RecipientId 				= d.RecipientId
////	pd.Status 					= ToPbStatus(d.Status)
////
////	pd.Documents 				= ToPbDocuments(d.Documents)
////	pd.Attempts					= d.Attempts
////	pd.NextTryTime				= d.NextTryTime.Unix()
////	pd.TimeToLive				= d.TimeToLive.Unix()
////	pd.Priority					= d.Priority
////	fmt.Printf("ToPbDelivery - 80: %v\n", pd)
////	return &pd
////}
//
//
//
//func RelPbToStatus(ps *relPB.Status) *mod.Status {
//	s := mod.Status{}
//	s.Status 	= ps.GetStatus()
//	s.StatusTime 	= time.Unix(ps.GetStatusTime(), 0)
//	s.Comment		= ps.GetComment()
//	return &s
//}
//
////func RelPbToHistory(ph []*relPB.Status) []*mod.Status {
////	sh := []*mod.Status{}
////	for _, ps := range ph {
////		sh = append(sh, RelPbToStatus(ps))
////	}
////	return sh
////}
//
//func ToPbStatus(s *mod.Status) *relPB.Status {
//	ps := relPB.Status{}
//	ps.Status 		= s.Status
//	ps.StatusTime 	= s.StatusTime.Unix()
//	ps.Comment		= s.Comment
//	return &ps
//}
//
//func ToPbHistory(sh []*mod.Status) []*relPB.Status {
//	psh := []*relPB.Status{}
//	for _, ps := range sh {
//		psh = append(psh, ToPbStatus(ps))
//	}
//	return psh
//}
//
//func RelPbToDocuments (pds []*relPB.Document) []*mod.Document {
//	da := []*mod.Document{}
//	for _, d := range pds {
//		da = append(da, RelPbToDocument(d))
//	}
//	return da
//}
//
//func RelPbToDocument(pd *relPB.Document) *mod.Document {
//	d := &mod.Document{}
//	//d.Id  			= pd.GetId()
//	d.Description 	= pd.GetDescription()
//	d.DocClass		= pd.GetDocClass()
//	d.ImageType 	= pd.GetImageType()
//	d.ImageRepository	= pd.GetImageRepository()
//	d.ImageURL 			= pd.GetImageUrl()
//	d.DocumentURL 		= pd.GetDocumentUrl()
//	d.Version		= pd.GetVersion()
//	return d
//}
//
//func ToPbDocuments (pds []*mod.Document) []*relPB.Document {
//	da := []*relPB.Document{}
//	for _, d := range pds {
//		da = append(da, ToPbDocument(d))
//	}
//	return da
//}
//
//func ToPbDocument(pd *mod.Document) *relPB.Document {
//	d := &relPB.Document{}
//	//d.Id				= pd.Id
//	d.Description 		= pd.Description
//	d.DocClass			= pd.DocClass
//	d.ImageType 		= pd.ImageType
//	d.ImageRepository	= pd.ImageRepository
//	d.ImageUrl 				= pd.ImageURL
//	d.DocumentUrl		= pd.DocumentURL
//	d.Version			= pd.Version
//	return d
//}
//
//
//func DelPbToDocuments (pds []*delPB.Document) []*mod.Document {
//	da := []*mod.Document{}
//	for _, d := range pds {
//		da = append(da, DelPbToDocument(d))
//	}
//	return da
//}
//
//func DelPbToDocument(pd *delPB.Document) *mod.Document {
//	d := &mod.Document{}
//	//d.Id 			= pd.GetId()
//	d.Description 	= pd.GetDescription()
//	d.DocClass		= pd.GetDocClass()
//	d.ImageType 	= pd.GetImageType()
//	d.ImageRepository	= pd.GetImageRepository()
//	d.ImageURL 		= pd.GetImageUrl()
//	d.DocumentURL 	= pd.GetDocumentUrl()
//	d.Version		= pd.GetVersion()
//	return d
//}
//
//func ToDelPbDocuments (pds []*mod.Document) []*delPB.Document {
//	da := []*delPB.Document{}
//	for _, d := range pds {
//		da = append(da, ToDelPbDocument(d))
//	}
//	return da
//}
//
//func ToDelPbDocument(pd *mod.Document) *delPB.Document {
//	d := &delPB.Document{}
//	//d.Id				= pd.Id
//	d.Description 		= pd.Description
//	d.DocClass			= pd.DocClass
//	d.ImageType 		= pd.ImageType
//	d.ImageRepository	= pd.ImageRepository
//	d.ImageUrl 				= pd.ImageURL
//	d.DocumentUrl 		= pd.DocumentURL
//	d.Version			= pd.Version
//	return d
//}
//
//
//
////MarshalDocument converts a *document.Document to a *relPB.Document
////func ToPbDocument(doc *c.Document) *relPB.Document {
////	ds := relPB.Document{}
////	ds.Id 				= doc.ID.Hex()
////	ds.Url 				= doc.URL
////	ds.ImageRepository 	= doc.ImageRepository
////	ds.ImageType 		= doc.ImageType
////	ds.Description		= doc.Description
////	ds.DocClass			= doc.DocClass
////
////	return &ds
////}
//
////UnMarshalDocument converts a *relPB.Document to a *domain.Document
////func ToDocument(doc *relPB.Document) *domain.Document {
////	ds := domain.Document{}
////	ds.ImageType 		= 	doc.GetImageType()
////	ds.ImageRepository 	= doc.GetImageRepository()
////	ds.URL 				= doc.GetUrl()
////	ds.DocClass			= doc.GetDocClass()
////	ds.Description		= doc.GetDescription()
////	ds.ID, _ = primitive.ObjectIDFromHex(doc.GetId())
////	return &ds
////}
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
////funcmmod.rshalDevice(d *Device) *relPB.Device {
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
////
//////MarshalDocuments converts from the client's Go Struct version of a Document to Protobuf Documentfor over the wire.
////func ToPbDocuments(ds []domain.Document) []*relPB.Document {
////	var pds []*relPB.Document
////	for _, d := range ds {
////		pd := relPB.Document{}
////		pd.Id				= d.ID.Hex()
////		pd.ImageRepository  = d.ImageRepository
////		pd.ImageType		= d.ImageType
////		pd.Url 				= d.URL
////		pd.Description		= d.Description
////		pd.DocClass			= d.DocClass
////		pds = append(pds, &pd)
////	}
////	return pds
////}
////
//////MarshalDocuments conversts from the client's Go Struct version of a Document to Protobuf Documentfor over the wire.
////func ToDocuments(pds []*relPB.Document) []domain.Document {
////	var dsa []domain.Document
////	for _, d := range pds {
////		pd := domain.Document{}
////		pd.URL 				= d.GetUrl()
////		pd.ImageType		= d.GetImageType()
////		pd.ImageRepository	= d.GetImageRepository()
////		pd.DocClass			= d.GetDocClass()
////		pd.Description		= d.GetDescription()
////		pd.ID, _ 			= primitive.ObjectIDFromHex(d.GetId())
////		dsa = append(dsa, pd)
////	}
////	return dsa
////}
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
