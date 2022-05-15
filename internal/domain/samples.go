package domain

//import (
//	"fmt"
//	"github.com/google/uuid"
//	"math/rand"
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
//		)
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
//		)
//}
//
//func RandomContact() string{
//	return fmt.Sprintf("%s, %s", RandomLastName(), RandomFirstName())
//}
//
//func RandomStreet() string {
//	return RandomStringFromSet(
//		"Misty Sage St.",
//			"Mission Crest Ave",
//			"White Waterfall Ave",
//			"Old Tree Trace",
//			"Avon Dr.",
//			"Lemon Ave",
//			"Calavo Dr.",
//			"Dale Ave",
//			"Hale Kanani Rd.",
//			"Denbury Dr",
//			"W. 11 St",
//			"S. Kihei Rd.",
//		)
//}
//
//func RandomCity() string{
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
//		)
//}
//
//func RandomAddress() string {
//	return fmt.Sprintf("%d %s", RandomInt(100,99999), RandomStreet() )
//}
//
//func RandomFax() string {
//	return RandomStringFromSet(
//		"855-998-7638",
//		"855-810-0810",
//
//		)/**/
//}
//
//func RandomZip() string {
//	return fmt.Sprintf("%d", RandomInt(10000, 99999))
//}

//func FaxDelivery() DeliveryFax {
//	return DeliveryFax{
//		FaxNumber: "855-998-7638",
//		ConfirmDate: "2011-08-10",
//	}
//}

//func EmailDevice(secure bool) *Device {
//	return &Device{
//		Method: "EMAIL",
//		Name: "Vsoft Email",
//		//Address: "dhf@vertioft.com",
//		//Combined: "true",
//	}
//}
//
//func AllScriptsDevice(mrn string) *Device {
//	if mrn == "" {
//		mrn = "test77777"
//	}
//	return &Device{
//		Method: "ALLSCRIPTS",
//		Name: "AllScripts",
//		//Address: "dhf@vertioft.com",
//		//Combined:  "false",
//	}
//}

//func FaxDevice(secure bool) *Device {
//	return &Device{
//		Method: "FAX",
//		Name:	"VSoft Fax",
//		//Address: "855-998-7638",
//		//Combined: "true",
//	}
//}

//func SampleRecipient() *Recipient {
//
//	return &Recipient {
//		ID:           primitive.NewObjectID(),
//		Name:      RandomContact(),
//		CompanyName:      RandomCompany(),
//		Address1:     RandomAddress(),
//		City:         RandomCity(),
//		State:        "NV",
//		Zip:       	  RandomZip(),
//	}
//}
//
//func SampleAllScriptsRecipient() *Recipient {
//
//	return &Recipient {
//		ID:           primitive.NewObjectID(),
//		Name:      RandomContact(),
//		CompanyName:      RandomCompany(),
//		Address1:     RandomAddress(),
//		City:         RandomCity(),
//		State:        "NV",
//		Zip:       	  RandomZip(),
//	}
//}