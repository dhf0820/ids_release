package service

import (
	"fmt"

	"github.com/google/uuid"

	docMod "github.com/dhf0820./ds_model/document"

	common "github.com/dhf0820./ds_model/common"
	delvMod "github.com/dhf0820./ds_model/delivery"
	patMod "github.com/dhf0820./ds_model/patient"
	relMod "github.com/dhf0820./ds_model/release"
	relPB "gitlab.com/dhf0820/ids_release_service/protobufs/relPB"
	"go.mongodb.org/mongo-driver/bson/primitive"

	//"go.mongodb.org/mongo-driver/bson/primitive"
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

func RandomBool() bool {
	return rand.Intn(2) == 1
}

func RandomInt(min, max int) int {
	return min + rand.Int()%(max-min+1)
}

func RandomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func RandomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func RandomID() string {
	return uuid.New().String()
}

func RandomZip() string {
	return fmt.Sprintf("%d", RandomInt(10000, 99999))
}
func RandomMRN() string {
	return RandomStringFromSet(
		"123456",
		"539023",
		"987654",
		"782145",
	)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

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

func RandomFirstName() string {
	return RandomStringFromSet(
		"Theresa",
		"Don",
		"Donald",
		"Preston",
		"Clayton",
		"Lacey",
		"Destiny",
		"Andrea",
		"Steven",
		"Matt",
		"Maria",
		"Foster",
	)
}

func RandomLastName() string {
	return RandomStringFromSet(
		"French",
		"Enyedi",
		"Novomeszky",
		"Kiss",
		"Parnell",
		"Trump",
		"Gray",
		"Schaufler",
		"Friedman",
		"Powell",
	)
}

func RandomContact() string {
	return fmt.Sprintf("%s, %s", RandomLastName(), RandomFirstName())
}

func RandomStreet() string {
	return RandomStringFromSet(
		"Misty Sage St.",
		"Mission Crest Ave",
		"White Waterfall Ave",
		"Old Tree Trace",
		"Avon Dr.",
		"Lemon Ave",
		"Calavo Dr.",
		"Dale Ave",
		"Hale Kanani Rd.",
		"Denbury Dr",
		"W. 11 St",
		"S. Kihei Rd.",
	)
}

func RandomCity() string {
	return RandomStringFromSet(
		"Las Vegas",
		"Roswell",
		"La Mesa",
		"Maui",
		"El Cajon",
		"Fort Worth",
		"Albany",
	)
}

func RandomCompany() string {
	return RandomStringFromSet(
		"VertiSoft Corp",
		"Enyedi Enterprises",
		"Enyedi Consulting",
		"Paxamor LLC",
		"Parnell Electric",
		"Friedman Consulting",
		"Innovotics",
		"InNEVotics",
		"Innovotix",
		"French Investments",
	)
}

func RandomAddress() string {
	return fmt.Sprintf("%d %s", RandomInt(100, 99999), RandomStreet())
}

func RandomFax() string {
	return RandomStringFromSet(
		"855-998-7638",
		"855-810-0810",
	) /**/
}

func SampleRelease() *relMod.Release {
	return &relMod.Release{
		//Customer:  testCustomer(),
		//Client:    testClient(),
		SrcID:      "1958",
		Patient:    SamplePatient(),
		Documents:  []*common.DocumentMeta{},
		Recipients: []*common.RecipientSummary{},
	}
}

func SamplePbReleaseRequest() *relPB.CreateReleaseRequest {
	//cr := &relPB.CreateRelease{
	//	Release:   SamplePbRelease(),
	//	Recipient: SamplePbRecipient(),
	//}
	return &relPB.CreateReleaseRequest{
		Release: SamplePbRelease(),
	}
}

func SamplePbRelease() *relPB.Release {
	return &relPB.Release{
		//Id:       "",
		//Client:   "Test",
		//Facility: "Test",
		//PatMrn:   RandomMRN(),
		//PatDob:   "1957-10-15",
		//PatName:  domain.RandomContact(),
		//SourceId: RandomIntString(10000, 99999),
		//UserName: "tfrench",
	}
}

func SamplePbRecipient() *relPB.Recipient {
	r := relPB.Recipient{}
	r.Id = "5fc184b28b273e53e19b1efa"
	r.Device = &relPB.Device{}
	r.Device.Id = "5fc163ed8b273e53e19b1ed5"
	return &r
}

//func SampleEmailDevice(secure bool) *relPB.Device {
//	return &relPB.Device{
//		Method: 	"EMAIL",
//		//Name:		"Vsoft Email",
//		//Address:    "dhf@vertisoft.com",
//	}
//}
//
//func SampleFaxDevice(secure bool) *relPB.Device {
//	return &relPB.Device{
//		Method:  "FAX",
//		//Name:    "VSoft Fax",
//		//Address: "855-998-7638",
//	}
//}

func RandomIntString(min, max int) string {
	return strconv.FormatInt(int64(RandomInt(min, max)), 10)
}

//func SampleRelease() *relMod.Release {
//
//	return &relMod.Release{
//		//Customer:  testCustomer(),
//		//Client:    testClient(),
//		SrcID:     "199",
//		Patient:   NewPatient(),
//		Documents: []*relMod.DocumentMetadata{},
//		Recipients: []*common.RecipientSummary{},
//	}
//}

//func NewPatient() *patMod.Patient {
//	return &patMod.Patient{
//		PatientName: RandomContact(),
//		MRN:         RandomMRN(),
//		DOB:         "1958-01-27",
//		SSN:         "000-04-5867",
//	}
//}

func SamplePatient() *patMod.Patient {
	return &patMod.Patient{
		PatientName: RandomContact(),
		MRN:         RandomMRN(),
		DOB:         "1958-01-27",
		SSN:         "000-04-5867",
	}
}

func SampleDocumentMetaData(srcId string, documentId primitive.ObjectID) *common.DocumentMetadata {
	dm := common.DocumentMetadata{}
	dm.ID = primitive.NewObjectID()
	dm.ImageType = "PDF"
	dm.DateOfService = "2021-12-01"
	dm.Description = "Cerner Description"
	dm.DocClass = "Radiology"
	dm.ID = documentId
	return &dm
}

//func newDocumentMetaData(srcId string, documentId primitive.ObjectID) *relMod.DocumentMetadata {
//	dm := relMod.DocumentMetadata{}
//	dm.ID = primitive.NewObjectID()
//	dm.ImageType = "PDF"
//	dm.DateOfService = "2021-12-01"
//	dm.Description = "Cerner Description"
//	dm.DocClass = "Radiology"
//	dm.ID = documentId
//	return &dm
//}

func SampleDocument() *docMod.Document {
	doc := docMod.Document{}
	corsrc := docMod.CorrelationId{}
	correl := docMod.CorrelationId{}
	corsrc.ReferenceType = "source"
	corsrc.ReferenceID = "37710"
	corsrc.SystemDocumentURL = "http://docker1.ihids.com:4567/api/patient_detail/37710"
	corsrc.SystemFacility = "demo"
	corsrc.SystemImageURL = "http://docker1.ihids.com:4567/api/pdf/37710"
	corsrc.Removable = "false"
	correl.ReferenceType = "release"
	correl.ReferenceID = strconv.Itoa(RandomInt(100, 99999))
	correl.SystemFacility = "demo"
	now := time.Now().UTC()
	corsrc.CreateTime = &now
	correl.CreateTime = &now
	corsrc.OriginatingIP = "192.168.1.26"
	correl.OriginatingIP = "192.168.1.26"
	doc.StorageType = "perm"
	doc.CorrelationIDs = append(doc.CorrelationIDs, &corsrc)
	doc.CorrelationIDs = append(doc.CorrelationIDs, &correl)
	return &doc
}

//
//func MakeSampleDocument() *docMod.Document {
//	doc := docMod.Document{}
//	src := docMod.CorrelationId{}
//	rel := docMod.CorrelationId{}
//	src.ReferenceType = "source"
//	src.ReferenceID = "37710"
//	src.SystemDocumentURL = "http://docker1.ihids.com:4567/api/patient_detail/37710"
//	src.SystemImageURL = "http://docker1.ihids.com:4567/api/pdf/37710"
//	src.OriginatingIP = "192.168.1.26"
//	src.SystemFacility = "demo"
//	now := time.Now().UTC()
//	src.CreateTime = &now
//	rel.CreateTime = &now
//	doc.CorrelationIDs = append(doc.CorrelationIDs, &src)
//	rel.ReferenceType = "source"
//	rel.ReferenceID = strconv.Itoa(RandomInt(100, 99999))
//	rel.SystemFacility = "demo"
//	rel.OriginatingIP = "192.168.1.26"
//	doc.ID = primitive.NewObjectID()
//	doc.CorrelationIDs = append(doc.CorrelationIDs, &rel)
//	doc.ImageType = "pdf"
//
//	return &doc
//}
//type DocumentMetadata struct {
//	ID 				primitive.ObjectID	`json:"id" bson:"id"` // actually the DocumentID
//	ImageType       string				`json:"image_type" bson:"image_type"`
//	DocumentURL		string 				`json:"document_url" bson:"document_url"` // our imageurl
//	//ImageID			primitive.ObjectID	`json:"image_id" bson:"image_id"`
//	Options 		[]*common.Option		`json:"options" bson:"options"`
//	Status 			*common.Status			`json:"status" bson:"status"`
//	DocClass		string				`json:"doc_class" bson:"doc_class"`  // in options
//	Description		string				`json:"description" bson:"description"` //options
//	DateOfService	string				`json:"date_of_service" bson:"date_of_service"` //Options
//	SrcDocURL		string				`json:"src_doc_url" bson:"src_doc_url"` // Options
//	SrcImageURL		string				`json:"src_image_url" bson:"src_image_url"`// Options
//	UpdatedAt 		time.Time			`json:"updated_at,omitempty" bson:"updated_at"`
//}
func MakeSampleMetaData() *common.DocumentMetadata {
	md := common.DocumentMetadata{}
	md.ImageType = "pdf"
	md.DateOfService = "2021-12-03"
	md.DocClass = "Radiology"
	md.Description = "X-Ray Left Knee"
	md.SrcDocURL = "http://docker1.ihids.com:4567/api/patient_detail/37710"
	md.ImageType = "PDF"
	md.SrcDocURL = "http://docker1.ihids.com:4567/api/pdf/37710"
	md.ID, _ = primitive.ObjectIDFromHex("61aaf12247b75e4e346e74c5")
	md.Status = &common.Status{
		Comment:    "Complete",
		State:      "OK",
		StatusTime: time.Now().UTC(),
	}

	return &md
}

//func testCustomer() *cm.Customer {
//	c := cm.Customer{}
//	c.Name = "test"
//	c.Code = "1"
//	c.Facility = "main"
//	//c.UserId 	= "tfrench"
//	return &c
//}
//
//func testClient() *cm.Client {
//	c := cm.Client{}
//	c.Source = "ca"
//	c.SourceId = "12345"
//	return &c
//}

//
//
//func testCustomer() *c.Customer {
//	cus := c.Customer{}
//	cus.Name = "test"
//	cus.Code = 100
//	cus.Facility = "main"
//	cus.UserID = "tfrench"
//	return &cus
//}

//func testClient() *c.Client {
//	cl := c.Client{}
//	cl.Source = "test_remote"
//	cl.SourceID = "12345"
//	return &cl
//}

//func testDelCustomer() *cm.Customer {
//	cus := cm.Customer{}
//	cus.Name = "test"
//	cus.Code = "100"
//	cus.Facility = "main"
//	//cus.UserId = "tfrench"
//	return &cus
//}
//
//func testDelClient() *cm.Client {
//	cl := cm.Client{}
//	cl.Source = "test_remote"
//	cl.SourceId = "12345"
//	return &cl
//}

func NewDelivery() *delvMod.Delivery {
	d := delvMod.Delivery{}
	//d.ReleaseId = "5fc3e3b9a652f88181ef5711"
	//d.Documents = []*pkg.DocumentMetadata{}
	//d.Customer = testDelCustomer()
	//d.Client = testDelClient()
	//d.Status = &m.Status{}
	//d.RecipientId = "5fc172d18b273e53e19b1ee0"
	//d.DeviceId = "5fc163ed8b273e53e19b1ed5"
	//
	//d.Priority = 3
	//d.TimeToLive = time.Now().Add(time.Hour * 72)
	//d.NextTryTime = time.Now()
	//d.Attempts = 0
	////fmt.Printf("\n\n\nNew Test Delivery: %s\n", spew.Sdump(d))
	return &d
}
