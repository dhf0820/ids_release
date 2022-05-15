package pkg

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//"status" : {
//"status" : "",
//"status_time" : "0001-01-01T00:00:00Z",
//"comment" : ""
//},
//type Release struct {
//	ID         primitive.ObjectID     `json:"id,omitempty" bson:"_id,omitempty"`
//	Patient    *patMod.Patient        `json:"patient,omitempty"bson:"patient"`
//	Documents  []*common.DocumentMeta `json:"documents" bson:"documents"`
//	Recipients []*Recipients          `json:"recipients" bson:"recipients"`
//	Status     *common.Status         `json:"status" bson:"status"`
//	Locked     string                 `json:"locked" bson:"locked"`   // truefalse
//	SrcID      string                 `json:"src_id" bson:"src_id"`   //Source release ID
//	SrcURL     string                 `json:"src_url" bson:"src_url"` // Source Release url
//	CreatedAt  time.Time              `json:"created_at,omitempty" bson:"created_at"`
//	UpdatedAt  time.Time              `json:"updated_at,omitempty" bson:"updated_at"`
//}

type Recipients struct {
	RecipientId   primitive.ObjectID `json:"recipient_id" bson:"recipient_id"`
	RecipientName string             `json:"recipient_name" bson:"recipient_name"`
	DeviceId      primitive.ObjectID `json:"device_id" bson:"device_id"`
}

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

//type SubmitDeliveryRequest struct {
//	ReleaseId 	primitive.ObjectID      `json:"release_id"`
//	RecipientId primitive.ObjectID 		`json:"recipient_id"` // If set use this instead of using stored recip
//	DeviceId   	primitive.ObjectID		`json:"device_id"`
//	CallBack 	string 					`json:"call_back"`
//}
//
//type SubmitDeliveryResponse struct {
//	Status 		string 			`json:"status"`
//	Message 	string 			`json:"message"`
//	Delivery 	*delvPkg.Delivery		`json:"delivery"`
//
//	//DeliveryId    primitive.ObjectID	`json:"delivery_id"`
//	//RecipientId   primitive.ObjectID	`json:"recipient_id"`
//	//RecipientName string 				`json:"recipient_name"`
//	//DeviceId      primitive.ObjectID 	`json:"device_id"`
//	//DeviceName    string 				`json:"device_name"`
//}
