package service

// import (
// 	"fmt"
// 	"github.com/davecgh/go-spew/spew"
// 	log "github.com/sirupsen/logrus"
// 	"time"
// 	//"github.com/dhf0820/ids_release_service/internal/domain"
// 	//delPB "github.com/dhf0820/ids_delivery/protobufs/delPB"
// 	relPB "github.com/dhf0820/ids_release_service/protobufs/relPB"
// 	//c "github.com/dhf0820/ids_release_service/pkg/common"
// 	mod "github.com/dhf0820./ds_model"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"regexp"
// 	//"time"
// )

// func ToClient(pc *relPB.Client) *mod.Client {
// 	c := mod.Client{}
// 	c.SourceId = pc.GetSourceId()
// 	c.Source = pc.GetSource()
// 	return &c
// }

// func ToPbClient(cl *mod.Client) *relPB.Client {
// 	pc := relPB.Client{}
// 	pc.SourceId = cl.SourceId
// 	pc.Source = cl.Source
// 	return &pc
// }
// func ToCustomer(pc *relPB.Customer) *mod.Customer {
// 	cu := &mod.Customer{}
// 	cu.Name = pc.GetName()

// 	cu.Facility = pc.GetFacility()
// 	cu.Code = pc.GetCode()
// 	return cu
// }

// func ToPbCustomer(cu *mod.Customer) *relPB.Customer {
// 	pc := relPB.Customer{}
// 	pc.Name = cu.Name
// 	pc.Facility = cu.Facility
// 	pc.Code = cu.Code
// 	return &pc
// }

// func ToField(f *mod.Field) *relPB.Field {
// 	pf := relPB.Field{}
// 	pf.Name = f.Name
// 	pf.Label = f.Label
// 	pf.Value = f.Value
// 	pf.UserVisible = f.UserVisible
// 	pf.Required = f.Required
// 	pf.Sensitive = f.Sensitive
// 	return &pf

// }

// func ToPbField(pf *relPB.Field) *mod.Field {
// 	f := mod.Field{}
// 	f.Name = pf.GetName()
// 	f.Label = pf.GetLabel()
// 	f.Value = pf.GetValue()
// 	f.Required = pf.GetRequired()
// 	f.UserVisible = pf.GetUserVisible()
// 	f.Sensitive = pf.GetSensitive()
// 	return &f
// }

// //UnMarshal takes a pb.Release and unmarshals it to domain.Release
// func ToRelease(pbRel *relPB.Release) *mod.Release {

// 	//pbRel := frm.GetRelease()
// 	rel := mod.Release{}
// 	rel.Client = ToClient(pbRel.GetClient())
// 	rel.Customer = ToCustomer(pbRel.GetCustomer())

// 	fmt.Printf("PatientMRN-82: %v\n", pbRel.GetPatient().GetMrn())
// 	rel.Patient = &mod.Patient{}
// 	rel.Patient.MRN = pbRel.GetPatient().GetMrn()
// 	rel.Patient.PatientName = pbRel.GetPatient().GetPatientName()
// 	rel.Patient.DOB = pbRel.GetPatient().GetDob()
// 	rel.Patient.SSN = pbRel.GetPatient().GetSsn()
// 	rel.Documents = RelPbToDocuments(pbRel.GetDocuments())
// 	rel.ID, _ = primitive.ObjectIDFromHex(pbRel.GetRelId())
// 	//fmt.Printf("FromPB: %s\n", spew.Sdump(rel))

// 	return &rel
// }

// //Marshal takes a domain release andmmod.rshals it to a pb.Release
// func ToPbRelease(rel *mod.Release) *relPB.Release {
// 	fmt.Printf("ToPbReleaserelease: %s\n", spew.Sdump(rel))
// 	fmt.Printf("\n\nPatient Name-98: [%s]\n", rel.Patient.PatientName)
// 	pRel := relPB.Release{}
// 	pat := &relPB.Patient{}
// 	pRel.Client = ToPbClient(rel.Client)
// 	pRel.Customer = ToPbCustomer(rel.Customer)
// 	pRel.Patient = pat
// 	pRel.Patient.PatientName = rel.Patient.PatientName
// 	pRel.Patient.Mrn = rel.Patient.MRN
// 	pRel.Patient.Dob = rel.Patient.DOB
// 	pRel.Patient.Ssn = rel.Patient.SSN
// 	pRel.RelId = rel.ID.Hex()
// 	pRel.Documents = ToPbDocuments(rel.Documents)
// 	pRel.CreatedAt = rel.CreatedAt.Unix()
// 	pRel.UpdatedAt = rel.UpdatedAt.Unix()

// 	return &pRel
// }

// func ToDocuments(pds []*relPB.Document) []*mod.Document {
// 	var ds []*mod.Document
// 	for _, pd := range pds {
// 		ds = append(ds, ToDocument(pd))
// 	}
// 	return ds
// }

// func ToDocument(doc *relPB.Document) *mod.Document {
// 	ds := mod.Document{}
// 	ds.ImageType = doc.GetImageType()
// 	ds.ImageRepository = doc.GetImageRepository()
// 	ds.ImageURL = doc.GetImageUrl()
// 	ds.DocumentURL = doc.GetDocumentUrl()
// 	ds.Description = doc.GetDescription()
// 	ds.DocClass = doc.GetDocClass()
// 	ds.Version = doc.Version
// 	//ds.Id 			= doc.GetId()
// 	return &ds
// }

// //ToDelivery accepts the PBDelivery format and converts it to domain so the domain can us it.
// // func ToDelivery(pbDev *delPB.Delivery) *mod.Delivery {

// // 	delv := mod.Delivery{}
// // 	//delv.Client	 			= &mod.Client{}
// // 	delv.Customer = &mod.Customer{}
// // 	delv.ID, _ = primitive.ObjectIDFromHex(pbDev.GetId())
// // 	//delv.Client.SourceId 	= pbDev.GetClient().GetSourceId()
// // 	//delv.Client.Source	 	= pbDev.GetClient().GetSource()
// // 	//delv.Customer.UserId 	= pbDev.GetCustomer().GetUserId()
// // 	delv.Customer.Name = pbDev.GetCustomer().GetName()
// // 	delv.Customer.Code = pbDev.GetCustomer().GetCode()
// // 	delv.Customer.Facility = pbDev.GetCustomer().GetFacility()
// // 	delv.ReleaseId = pbDev.GetReleaseId()
// // 	delv.RecipientId = pbDev.RecipientId
// // 	delv.Status = ToStatus(pbDev.GetStatus())
// // 	delv.Documents = DelPbToDocuments(pbDev.GetDocuments())
// // 	delv.Attempts = pbDev.GetAttempts()
// // 	delv.NextTryTime = time.Unix(pbDev.GetNextTryTime(), 0)
// // 	delv.TimeToLive = time.Unix(pbDev.GetTimeToLive(), 0)
// // 	delv.Priority = pbDev.GetPriority()
// // 	fmt.Printf("ToDelivery: %s\n", spew.Sdump(delv))
// // 	return &delv
// // }

// //ToPbDelivery accepts a domain delivery struct and puts it into the PB.delivery format
// // func ToDelPbDelivery(d *mod.Delivery) *delPB.Delivery {
// // 	if d == nil {
// // 		fmt.Printf("Delivery marshal is nil\n")
// // 	}
// // 	fmt.Printf("ToPBDelivery - 55: %s\n", spew.Sdump(d))
// // 	pd := delPB.Delivery{}
// // 	//pd.Client 					= &delPB.Client{}
// // 	pd.Customer = &delPB.Customer{}
// // 	pd.Id = d.ID.Hex()
// // 	//pd.Client.Source			= d.Client.Source
// // 	//pd.Client.SourceId			= d.Client.SourceId
// // 	//pd.Customer.UserId			= d.Customer.UserId
// // 	pd.Customer.Name = d.Customer.Name
// // 	pd.Customer.Code = d.Customer.Code
// // 	pd.Customer.Facility = d.Customer.Facility
// // 	pd.ReleaseId = d.ReleaseId
// // 	pd.Status = ToDelPbStatus(d.Status)
// // 	pd.Documents = ToDelPbDocuments(d.Documents)
// // 	pd.Attempts = d.Attempts
// // 	pd.NextTryTime = d.NextTryTime.Unix()
// // 	pd.TimeToLive = d.TimeToLive.Unix()
// // 	pd.Priority = d.Priority
// // 	fmt.Printf("ToPbDelivery - 80: %v\n", pd)
// // 	return &pd
// // }

// /* func ToStatus(ps *delPB.Status) *mod.Status {
// 	s := mod.Status{}
// 	s.Status = ps.GetStatus()
// 	s.StatusTime = time.Unix(ps.GetStatusTime(), 0)
// 	s.Comment = ps.GetComment()
// 	return &s
// }

// func DelPbToStatus(ps *delPB.Status) *mod.Status {
// 	s := mod.Status{}
// 	s.Status = ps.GetStatus()
// 	s.StatusTime = time.Unix(ps.GetStatusTime(), 0)
// 	s.Comment = ps.GetComment()
// 	return &s
// }

// func ToDelPbStatus(s *mod.Status) *delPB.Status {
// 	ps := delPB.Status{}
// 	ps.Status = s.Status
// 	ps.StatusTime = s.StatusTime.Unix()
// 	ps.Comment = s.Comment
// 	return &ps
// }

// func ToDelPbHistory(sh []*mod.Status) []*delPB.Status {
// 	psh := []*delPB.Status{}
// 	for _, ps := range sh {
// 		psh = append(psh, ToDelPbStatus(ps))
// 	}
// 	return psh
// } */

// func RelPbToStatus(ps *relPB.Status) *mod.Status {
// 	s := mod.Status{}
// 	s.Status = ps.GetStatus()
// 	s.StatusTime = time.Unix(ps.GetStatusTime(), 0)
// 	s.Comment = ps.GetComment()
// 	return &s
// }

// func ToPbStatus(s *mod.Status) *relPB.Status {
// 	ps := relPB.Status{}
// 	ps.Status = s.Status
// 	ps.StatusTime = s.StatusTime.Unix()
// 	ps.Comment = s.Comment
// 	return &ps
// }

// func ToPbHistory(sh []*mod.Status) []*relPB.Status {
// 	psh := []*relPB.Status{}
// 	for _, ps := range sh {
// 		psh = append(psh, ToPbStatus(ps))
// 	}
// 	return psh
// }

// func RelPbToDocuments(pds []*relPB.Document) []*mod.Document {
// 	da := []*mod.Document{}
// 	for _, d := range pds {
// 		da = append(da, RelPbToDocument(d))
// 	}
// 	return da
// }

// func RelPbToDocument(pd *relPB.Document) *mod.Document {
// 	d := &mod.Document{}
// 	d.Id = pd.GetId()
// 	d.Description = pd.GetDescription()
// 	d.DocClass = pd.GetDocClass()
// 	d.ImageType = pd.GetImageType()
// 	d.ImageRepository = pd.GetImageRepository()
// 	d.ImageURL = pd.GetImageUrl()
// 	d.DocumentURL = pd.GetDocumentUrl()
// 	d.Version = pd.GetVersion()
// 	d.SourceDocumentURL = pd.GetSourceDocumentUrl()
// 	d.SourceImageURL = pd.GetSourceImageUrl()
// 	return d
// }

// func ToPbDocuments(pds []*mod.Document) []*relPB.Document {
// 	da := []*relPB.Document{}
// 	for _, d := range pds {
// 		da = append(da, ToPbDocument(d))
// 	}
// 	return da
// }

// func ToPbDocument(pd *mod.Document) *relPB.Document {
// 	d := &relPB.Document{}
// 	d.Id = pd.Id
// 	d.Description = pd.Description
// 	d.DocClass = pd.DocClass
// 	d.ImageType = pd.ImageType
// 	d.ImageRepository = pd.ImageRepository
// 	d.ImageUrl = pd.ImageURL
// 	d.DocumentUrl = pd.DocumentURL
// 	d.Version = pd.Version
// 	d.SourceImageUrl = pd.SourceImageURL
// 	d.SourceDocumentUrl = pd.SourceDocumentURL
// 	return d
// }

// /* func DelPbToDocuments(pds []*delPB.Document) []*mod.Document {
// 	da := []*mod.Document{}
// 	for _, d := range pds {
// 		da = append(da, DelPbToDocument(d))
// 	}
// 	return da
// }

// func DelPbToDocument(pd *delPB.Document) *mod.Document {
// 	d := &mod.Document{}
// 	//d.Id 			= pd.GetId()
// 	d.Description = pd.GetDescription()
// 	d.DocClass = pd.GetDocClass()
// 	d.ImageType = pd.GetImageType()
// 	d.ImageRepository = pd.GetImageRepository()
// 	d.ImageURL = pd.GetImageUrl()
// 	d.DocumentURL = pd.GetDocumentUrl()
// 	d.Version = pd.GetVersion()
// 	return d
// }

// func ToDelPbDocuments(pds []*mod.Document) []*delPB.Document {
// 	da := []*delPB.Document{}
// 	for _, d := range pds {
// 		da = append(da, ToDelPbDocument(d))
// 	}
// 	return da
// }

// func ToDelPbDocument(pd *mod.Document) *delPB.Document {
// 	d := &delPB.Document{}
// 	//d.Id				= pd.Id
// 	d.Description = pd.Description
// 	d.DocClass = pd.DocClass
// 	d.ImageType = pd.ImageType
// 	d.ImageRepository = pd.ImageRepository
// 	d.ImageUrl = pd.ImageURL
// 	d.DocumentUrl = pd.DocumentURL
// 	d.Version = pd.Version
// 	return d
// } */

// func NormalizePhone(phone string) string {
// 	fmt.Printf("Normalizing : %s\n", phone)
// 	reg, err := regexp.Compile("[^0-9]*")
// 	if err != nil {
// 		log.Errorf("regexp compile failed: %v", err)
// 	}
// 	return reg.ReplaceAllString(phone, "")
// }

// func ValidateEmail(email string) bool {
// 	fmt.Printf("Validating : %s\n", email)
// 	str := "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
// 	re := regexp.MustCompile(str)
// 	return re.MatchString(email)
// }
