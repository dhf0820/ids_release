package client

//import (
//	"fmt"
//	"github.com/dhf0820/ids_release_service/internal/domain"
//	m "github.com/dhf0820./ds_model"
//	"go.mongodb.org/mongo-driver/bson/primitive"
//	"strconv"
//
//	//"github.com/dhf0820/ids_release_service/internal/service"
//	//"github.com/dhf0820/ids_release_service/protobufs/relPB"
//
//	//"github.com/dhf0820/ids_release_service/internal/domain"
//	"math/rand"
//	//"strconv"
//	"time"
//
//	"github.com/google/uuid"
//)
//
//func init() {
//	rand.Seed(time.Now().UnixNano())
//}
//
//
//func RandomStringFromSet(a ...string) string {
//	n := len(a)
//	if n == 0 {
//		return ""
//	}
//	return a[rand.Intn(n)]
//}
//
//func RandomBool() bool {
//	return rand.Intn(2) == 1
//}
//
//func RandomInt(min, max int) int {
//	return min + rand.Int()%(max-min+1)
//}
//
//func RandomFloat64(min, max float64) float64 {
//	return min + rand.Float64()*(max-min)
//}
//
//func RandomFloat32(min, max float32) float32 {
//	return min + rand.Float32()*(max-min)
//}
//
//func RandomID() string {
//	return uuid.New().String()
//}
//
//func RandomMRN() string {
//	return RandomStringFromSet(
//		"123456",
//		"539023",
//		"987654",
//		"782145",
//	)
//}
//
//func RandomFirstName() string {
//	return RandomStringFromSet(
//		"Theresa",
//		"Don",
//		"Donald",
//		"Preston",
//		"Clayton",
//		"Lacey",
//		"Destiny",
//		"Andrea",
//		"Steven",
//		"Matt",
//		"Maria",
//		"Foster",
//	)
//}
//
//func RandomLastName() string {
//	return RandomStringFromSet(
//		"French",
//		"Enyedi",
//		"Novomeszky",
//		"Kiss",
//		"Parnell",
//		"Trump",
//		"Gray",
//		"Schaufler",
//		"Friedman",
//		"Powell",
//	)
//}
//
//func RandomContact() string {
//	return fmt.Sprintf("%s, %s", RandomLastName(), RandomFirstName())
//}
//
//func RandomStreet() string {
//	return RandomStringFromSet(
//		"Misty Sage St.",
//		"Mission Crest Ave",
//		"White Waterfall Ave",
//		"Old Tree Trace",
//		"Avon Dr.",
//		"Lemon Ave",
//		"Calavo Dr.",
//		"Dale Ave",
//		"Hale Kanani Rd.",
//		"Denbury Dr",
//		"W. 11 St",
//		"S. Kihei Rd.",
//	)
//}
//
//func RandomCity() string {
//	return RandomStringFromSet(
//		"Las Vegas",
//		"Roswell",
//		"La Mesa",
//		"Maui",
//		"El Cajon",
//		"Fort Worth",
//		"Albany",
//	)
//}
//
//func RandomCompany() string {
//	return RandomStringFromSet(
//		"VertiSoft Corp",
//		"Enyedi Enterprises",
//		"Enyedi Consulting",
//		"Paxamor LLC",
//		"Parnell Electric",
//		"Friedman Consulting",
//		"Innovotics",
//		"InNEVotics",
//		"Innovotix",
//		"French Investments",
//	)
//}
//
//func RandomAddress() string {
//	return fmt.Sprintf("%d %s", RandomInt(100, 99999), RandomStreet())
//}
//
//func RandomFax() string {
//	return RandomStringFromSet(
//		"855-998-7638",
//		"855-810-0810",
//	) /**/
//}
//
//func RandomIntString(min, max int) string {
//	return strconv.FormatInt(int64(RandomInt(min, max)), 10)
//}
//
//
//func RandomZip() string {
//	return fmt.Sprintf("%d", RandomInt(10000, 99999))
//}
//
////
////func FaxDelivery() DeliveryFax {
////	return DeliveryFax{
////		FaxNumber: "855-998-7638",
////		ConfirmDate: "2011-08-10",
////	}
////}
////
////func EmailDevice(secure bool) *Device {
////	return &Device{
////		Method: "EMAIL",
////		Name: "Vsoft Email",
////		Address: "dhf@vertioft.com",
////		Secure: secure,
////	}
////}
////
////func FaxDevice(secure bool) *Device {
////	return &Device{
////		Method: "FAX",
////		Name:	"VSoft Fax",
////		Address: "855-998-7638",
////		Command: "",
////		Secure: secure,
////	}
////}
//
//
////func SampleRelPbReleaseRequest() *relPB.CreateReleaseRequest {
////	//cr := &pb.CreateRelease{
////	//	Release:   SamplePbRelease(),
////	//	Recipient: SamplePbRecipient(),
////	//}
////	return &relPB.CreateReleaseRequest{
////		Release: SampleRelPbRelease(),
////	}
////}
//func SampleRelease() *m.Release {
//	r := &m.Release{}
//	r.ID = primitive.NewObjectID()
//	r.Client = &m.Client{}
//	r.Customer = &m.Customer{}
//	r.Patient  = &m.Patient{}
//	r.Patient.MRN = RandomMRN()
//	r.Patient.DOB = "1957-10-15"
//	r.Patient.PatientName = domain.RandomContact()
//	r.Client.SourceId = RandomIntString(10000, 99999)
//	r.Client.Source   = "ca"
//	r.Customer.Name		= "test"
//	r.Customer.Code 	= 1
//	r.Customer.Facility	= "demo"
//	r.Customer.UserId	= "tfrench"
//	return r
//}
//
//func SampleDocument() *m.Document {
//	doc := &m.Document{}
//	//doc.Id = primitive.NewObjectID().Hex()
//	doc.ImageType = "pdf"
//	doc.ImageRepository = "ca"
//	doc.ImageURL = "/pdf/175"
//	doc.DocumentURL = "/175"
//	doc.Description = "Mamo"
//	doc.DocClass = "Radiology"
//	return doc
//}
//
////
////func SampleRelPbRelease() *relPB.Release {
////	return &relPB.Release{
////		Id:       "",
////		Client:   "Test",
////		Facility: "Test",
////		PatMrn:   RandomMRN(),
////		PatDob:   "1957-10-15",
////		PatName:  domain.RandomContact(),
////		SourceId: RandomIntString(10000, 99999),
////		UserName: "tfrench",
////	}
////}
////
////func SampleRelPbRecipient() *relPB.Recipient {
////
////	return &relPB.Recipient{
////		Id:           "",
////		RemoteId:     "Rec-1",
////		ElectronicId: "1234567",
////		Contact:      domain.RandomContact(),
////		Company:      domain.RandomCompany(),
////		Address1:     domain.RandomAddress(),
////		City:         domain.RandomCity(),
////		State:        "NV",
////		Postal:       RandomZip(),
////	}
////}
////func SampleRelAllScriptsDevice(secure bool) *relPB.Device {
////	return &relPB.Device{
////		Method:         "ALLSCRIPTS",
////		Name:     "AllScripts",
////		Address:        "",
////	}
////}
////
////func SampleRelPbEmailDevice(secure bool) *relPB.Device {
////	return &relPB.Device{
////		Method:     "EMAIL",
////		Name: "Vsoft Email",
////		Address:    "dhf@vertisoft.com",
////
////	}
////}
////
////func SampleRelPbFaxDevice(secure bool) *relPB.Device {
////	return &relPB.Device{
////		Method:     "FAX",
////		Name: "VSoft Fax",
////		Address:    "855-998-7638",
////
////	}
////}
////
////func RandomIntString(min, max int) string {
////	return strconv.FormatInt(int64(RandomInt(min, max)), 10)
////}
