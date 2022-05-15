package common

//
//type Client struct {
//	Source		string 	`json:"source"`
//	SourceID	string	`json:"source_id"`
//}
//
//type Customer struct {
//	Name 		string	`json:"name"`
//	Code 		int32 	`json:"code"`
//	Facility 	string 	`json:"facility"`
//	UserID 		string 	`json:"user_id"`
//}
//
//
//type Fields struct {
//	Name 		string 			`json:"name"`
//	Label 		string 			`json:"label"`
//	Required 	string			`json:"required"`
//	Value 		string 			`json:"value"`
//	UserVisible string 			`json:"user_visible"`
//}
//
//
//type Delivery struct {
//	ID			    primitive.ObjectID 	`json:"_id,omitempty" bson:"_id,omitempty"`
//	ReleaseId      	primitive.ObjectID  `json:"release_id"`
//	Client         	Client            `json:"client,omitempty" bson:"client"`
//	Customer 		Customer			`json::"customer"`
//	Recipient 		RecipientSummary 	`json:"recipient"`
//	Priority 		int32				`json:"priority"`
//	Attempts 		int32 				`json:"attempts"`
//	NextTryTime 	time.Time 			`json:"next_try_time"`
//	TimeToLive 		time.Time 			`json:"time_to_live"`
//	Status 			Status 				`json:"status"` // Latest tatus
//	History 		[]Status 			`json:"history"` // History of all statuses
//	Documents 		[]Document			`json:"documents"`
//	Device			Device 				`json:"device" bson:"device"`
//	Patient 		Patient 			`json:"patient" bson:"patient"`
//	CreatedAt 		time.Time 			`json:"created_at,omitempty" bson:"created_at"`
//	UpdatedAt 		time.Time 			`json:"updated_at,omitempty" bson:"updated_at"`
//}
//
//type Status struct {
//	Status     	string     	`json:"status"` // "submitted", "pending", "queued", "inprocess", "delivered", "error", "failed"
//	StatusTime 	time.Time 	`json:"status_time"`
//	Comment		string     	`json:"comment"`
//}
//
//type RecipientSummary struct {
//	ID 				primitive.ObjectID 	`json:"id"  bson:"_id,omitempty" `
//	Device 			DeviceSummary		`json:"device"`
//}
//
//type DeviceSummary struct {
//	ID     primitive.ObjectID `json:"id" bson:"id"`
//	Method string             `json:"method"`
//	Name   string             `json:"name"`
//}
//type Device struct {
//	ID     primitive.ObjectID `json:"id" bson:"id"`
//	Method string             `json:"method"`
//	Name   string             `json:"name"`
//}
//
//
//type Document struct {
//	ID 			primitive.ObjectID	`json:"id" bson:"_id"`	// Original Document ID
//	ImageRepository string 			`json:"image_repository" bson:"image_repository"`
//	ImageType 		string 			`json:"image_type" bson:"image_type"`
//	ImageURL 			string 			`json:"image_url" bson:"image_url"`
//	DocumentURL 		string 			`json:"document_url" bson:"document_url"`
//	DocClass			string 			`json:"doc_class" bson:"doc_class"`
//	Description 	string 			`json:"description"`
//	Version 		string 			`json:"version"`
//}
//
//type Patient struct {
//	Source		string 	`json:"source"`
//	SourceID 	string `json:"source_id"`
//	PatientName string `json:"patient_name"`
//	MRN      	string `json:"mrn"`
//	DOB      	string `json:"dob"`
//	SSN      	string `json:"ssn"`
//}