package service

//
//import (
//	"fmt"
//	"github.com/google/uuid"
//	"gitlab.com/dhf0820/ids_release_service/internal/domain"
//	relPB "gitlab.com/dhf0820/ids_release_service/protobufs/relPB"
//	m "github.com/dhf0820./ds_model"
//
//	//"go.mongodb.org/mongo-driver/bson/primitive"
//	"math/rand"
//	"strconv"
//	"time"
//)
//
//func init() {
//	rand.Seed(time.Now().UnixNano())
//}
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
//func RandomZip() string {
//	return fmt.Sprintf("%d", RandomInt(10000, 99999))
//}
//func RandomMRN() string {
//	return RandomStringFromSet(
//		"123456",
//		"539023",
//		"987654",
//		"782145",
//	)
//}
//
//
//
//func SampleRelease() *m.Release {
//
//	return nil
//}
//
//func SamplePbReleaseRequest() *relPB.CreateReleaseRequest {
//	//cr := &relPB.CreateRelease{
//	//	Release:   SamplePbRelease(),
//	//	Recipient: SamplePbRecipient(),
//	//}
//	return &relPB.CreateReleaseRequest{
//		Release: SamplePbRelease(),
//	}
//}
//
//
//
//func SamplePbRelease() *relPB.Release {
//	return &relPB.Release {
//		//Id:       "",
//		//Client:   "Test",
//		//Facility: "Test",
//		//PatMrn:   RandomMRN(),
//		//PatDob:   "1957-10-15",
//		//PatName:  domain.RandomContact(),
//		//SourceId: RandomIntString(10000, 99999),
//		//UserName: "tfrench",
//	}
//}
//
//func SamplePbRecipient() *relPB.Recipient {
//	r := relPB.Recipient {}
//	r.Id			= "5fc184b28b273e53e19b1efa"
//	r.Device		= &relPB.Device{}
//	r.Device.Id		= "5fc163ed8b273e53e19b1ed5"
//	return &r
//}
//
////func SampleEmailDevice(secure bool) *relPB.Device {
////	return &relPB.Device{
////		Method: 	"EMAIL",
////		//Name:		"Vsoft Email",
////		//Address:    "dhf@vertisoft.com",
////	}
////}
////
////func SampleFaxDevice(secure bool) *relPB.Device {
////	return &relPB.Device{
////		Method:  "FAX",
////		//Name:    "VSoft Fax",
////		//Address: "855-998-7638",
////	}
////}
//
//
//func RandomIntString(min, max int) string {
//	return strconv.FormatInt(int64(RandomInt(min, max)), 10)
//}
//
//func NewRelease() *m.Release {
//
//	return &m.Release{
//		Customer: 	testCustomer(),
//		Client: 	testClient(),
//		Patient: 	NewPatient(),
//		Documents:  []*m.Document{},
//
//	}
//}
//
//func NewPatient() *m.Patient {
//	return &m.Patient{
//		PatientName: domain.RandomContact(),
//		MRN:  RandomMRN(),
//		DOB:  "1957-10-15",
//		SSN:  "000-04-5867",
//	}
//}
//
//
//
//func testCustomer() *m.Customer {
//	c := m.Customer {}
//	c.Name 		= "test"
//	c.Code 		= 1
//	c.Facility 	= "main"
//	c.UserId 	= "tfrench"
//	return &c
//}
//
//func testClient() *m.Client {
//	c := m.Client {}
//	c.Source 	= "ca"
//	c.SourceId 	= "12345"
//	return &c
//}
////
////
////func testCustomer() *c.Customer {
////	cus := c.Customer{}
////	cus.Name = "test"
////	cus.Code = 100
////	cus.Facility = "main"
////	cus.UserID = "tfrench"
////	return &cus
////}
//
////func testClient() *c.Client {
////	cl := c.Client{}
////	cl.Source = "test_remote"
////	cl.SourceID = "12345"
////	return &cl
////}
//
//func testDelCustomer() *m.Customer {
//	cus := m.Customer{}
//	cus.Name = "test"
//	cus.Code = 100
//	cus.Facility = "main"
//	cus.UserId = "tfrench"
//	return &cus
//}
//
//func testDelClient() *m.Client {
//	cl := m.Client{}
//	cl.Source = "test_remote"
//	cl.SourceId = "12345"
//	return &cl
//}
//func NewDelivery() *m.Delivery {
//	d := m.Delivery{}
//	d.ReleaseId 	= "5fc3e3b9a652f88181ef5711"
//	d.Documents		= []*m.Document{}
//	d.Customer 		= testDelCustomer()
//	d.Client   		= testDelClient()
//	d.Status 		= &m.Status{}
//	d.RecipientId 	= "5fc172d18b273e53e19b1ee0"
//	d.DeviceId 		= "5fc163ed8b273e53e19b1ed5"
//
//	d.Priority 		= 3
//	d.TimeToLive 	= time.Now().Add(time.Hour * 72)
//	d.NextTryTime 	= time.Now()
//	d.Attempts 	 	= 0
//	//fmt.Printf("\n\n\nNew Test Delivery: %s\n", spew.Sdump(d))
//	return &d
//}
//
//
