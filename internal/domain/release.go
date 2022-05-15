package domain

//
//import (
//	"context"
//	"fmt"
//	log "github.com/sirupsen/logrus"
//	db "gitlab.com/dhf0820/ids_release_service/internal/db/vsmongo"
//	m "github.com/dhf0820./ds_model"
//	"google.golang.org/grpc/codes"
//	"google.golang.org/grpc/status"
//	"time"
//	//db "gitlab.com/dhf0820/ids_release_service/internal/db/vsmongo"
//	//"gitlab.com/dhf0820/ids_release_service/internal/settings"
//	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/bson/primitive"
//)
//
//
////type Release struct {
////	ID       		primitive.ObjectID 	`json:"id,omitempty" bson:"_id,omitempty"`
////	Client   		dc.Client            `json:"client,omitempty" bson:"client"`
////	Customer 		dc.Customer 			`json:"customer"`
////	Patient  		Patient            	`json:"patient,omitempty"`
////	Documents 		[]dc.Document  		`json:"documents"`
////	History			[]DeliveryHistory 	`json:"history,omitempty" bson:"delivery_history"`
////	CreatedAt       time.Time          	`json:"created_at,omitempty"`
////	UpdatedAt       time.Time          	`json:"updated_at,omitempty"`
////}
////
////
//////TODO: Is this needed or just use the ID and possible friendly name
////
//////DeviceDetail identifies the method of delivery and details. It is kept in the Recipient history once it is delivered.
////type DeviceDetail struct {
////	Id              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
////	RecipientId     primitive.ObjectID `json:"recipient_id" bson:"recipient_id"`
////	Method          string             `json:"method"`
////	Name            string             `json:"name"`
////	Validated 		string 				`json:"validated"`
////	ValidatedAt 	*time.Time			`json:"validated_at"`
////	Fields 			[]Fields 			`json:"fields"`
////	UpdatedAt       time.Time          `json:"updated_at"`
////}
////
////
////type Fields struct {
////	Name 		string 			`json:"name"`
////	Label 		string 			`json:"label"`
////	Required 	string			`json:"required"`
////	Value 		string 			`json:"value"`
////	UserVisible bool 			`json:"user_visible"`
////}
////
//////TODO: Is this needed or just maintain the Recipient ID
//////Recipient is who the
//////type Recipient struct {
//////	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
//////	Name        string             `json:"name"`
//////	CompanyName string             `json:"company_name"`
//////	Address1    string             `json:"address1"`
//////	Address2    string             `json:"address2"`
//////	City        string             `json:"city"`
//////	State       string             `json:"state"`
//////	Zip         string             `json:"zip"`
//////	Phone       string             `json:"phone"`
//////	Fax         string             `json:"fax"`
//////	Email       string             `json:"email"`
//////	Type        string             `json:"type"` //MedicalPractice, Patient, Relative, Service
//////	Devices     []DeviceSummary    `json:"devices"`
//////	CreatedAt   time.Time          `json:"created_at,omitempty"`
//////	UpdatedAt   time.Time          `json:"updated_at,omitempty"`
//////	Country 	string 				`json:"country"`
//////}
////
////type DeviceSummary struct {
////	Id     primitive.ObjectID `json:"_id"`
////	Method string             `json:"method"`
////	Name   string             `json:"name"`
////}
////
//////Device identifies the method of delivery and details. It is kept in the Recipient history once it is delivered.
////type Device struct {
////	Id              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
////	Method          string             `json:"method"`
////}
////
//////type Document struct {
//////	ID 			primitive.ObjectID	`json:"id" bson:"_id"`	// Original Document ID
//////	ImageRepository string 			`json:"image_repositry"`
//////	ImageType 		string 			`json:"image_type"`
//////	URL 			string 			`json:"url"`
//////	DocClass			string 			`json:"doc_class"`
//////	Description 	string 			`json:"descripition"`
//////	Version 		string 			`json:"version"`
//////}
////
////
////type StatusHistory struct {
////	Status     string     `json:"status"`
////	Comment    string     `json:"comment"`
////	StatusTime *time.Time `json:"status_time"`
////}
////
//////Patient contains the patient information and is a subdocument in Release
////type Patient struct {
////	SourceID string `json:"source_id"`
////	Source 	 string `json:"source"`
////	PatientName  string `json:"patient_name"`
////	MRN      string `json:"mrn"`
////	DOB      string `json:"dob"`
////	SSN      string `json:"ssn"`
////}
////
/////* 	DeliveryHistory need to think about this for right structure.  A release may be delivered to
////multiple recipients, or one Recipient multiple times. What isthe proper structure */
////type DeliveryHistory struct {
////	ID           primitive.ObjectID `json:"id" bson:"_id"`
////	RecipientID  primitive.ObjectID `json:"recipient_id"`
////	RecipientName string			`json:"recipient_name"`
////	DeviceID 	 primitive.ObjectID `json:"device_id"`
////	DeviceName 	 string 			`json:"device_name"`
////	SubmitTime   *time.Time         `json:"submit_time"`
////	DeliveryTime *time.Time         `json:"delivery_time"`
////	Status 		[]StatusHistory		`json:"status"`
////}
////
//////type DeliveryStatus struct {
//////	Status       string             `json:"status"`			// "STARTED", "PENDING", "QUEUED", "DELIVERED", "ERROR", "FAILED"
//////	StatusTime   *time.Time         `json:"status_time"`
//////	Comments     string             `json:"comments"`
//////}
////
////type DeliveryRequest struct {
////	Customer  string 		`json:"customer"`
////	ReleaseID string           `json:"release_id"`
////	RecipientID string 		 `json:"recipient_id"` // If set use this instead of using stored recipie
////	DeviceID   string		`json:"device_id"`
////}
//
//// Release Methods
//func GetReleaseById(oid primitive.ObjectID) (rel *m.Release, err error) {
//	rel = &m.Release{}
//	fmt.Printf("GetReleaseByID [%s]\n", oid)
//	filter := bson.M{"_id": oid}
//	return getRelease(filter)
//}
//
//func GetReleaseByForeignId(fId, customer string)(rel *m.Release, err error) {
//	rel = &m.Release{}
//
//	//TODO: Refactor client to customer and sourceID to foreignID
//	filter := bson.M{"sourceid": fId, "client": customer}
//	return getRelease(filter)
//}
//
//func GetReleaseByClient(client *m.Client)(rel *m.Release, err error) {
//	rel = &m.Release{}
//
//	//TODO: Refactor client to customer and sourceID to foreignID
//	filter := bson.M{"client": *client}
//	return getRelease(filter)
//}
//
//
//func getRelease(filter bson.M) (rel *m.Release, err error) {
//	rel = &m.Release{}
//	fmt.Printf("Using filter: %v\n", filter)
//
//	collection, err := db.GetCollection("releases")
//	if err != nil {
//		return nil, status.Errorf(
//			codes.Internal,
//			fmt.Sprintf("domain.GetRelease GetCollection failed: %v\n", err),
//		)
//	}
//	fmt.Printf("Calling FindOne: %v\n", filter)
//	err = collection.FindOne(context.Background(), filter).Decode(rel)
//	if err != nil {
//		log.Errorf("domain.GetRelease FindOne [%v] NotFound\n", filter)
//		return nil, status.Errorf(
//			codes.NotFound,
//			fmt.Sprintf("domain.GetRelease FindOne [%v] NotFound\n", filter),
//		)
//	}
//	//fmt.Printf("\n\n$$$$getrelease: %v\n", spew.Sdump(rel))
//	return
//}
//
//
//func CreateRelease(data *m.Release) (*m.Release, error) {
//	var err error
//	//var insertResult *mongo.InsertOneResult
//	//fmt.Printf("\n$$$ domain createRelease Received save: %s\n", spew.Sdump(data))
//
//	fmt.Printf("Before Get Collection\n")
//	collection, err := db.GetCollection("releases")
//	if err != nil {
//		fmt.Printf("GetCollection failed: %v\n", err)
//		return nil, err
//	}
//
//	data.CreatedAt = time.Now()
//	data.UpdatedAt = data.CreatedAt
//
//
//	res, err := collection.InsertOne(context.Background(), data)
//	if err != nil {
//		log.Errorf("InsertOne failed: %v\n", err)
//		return nil, status.Errorf(
//			codes.Internal,
//			fmt.Sprintf("Internal error saving Document: %v", err),
//		)
//	}
//	oid, ok := res.InsertedID.(primitive.ObjectID)
//	fmt.Printf("domain/create Hex oid: %s\n", oid.Hex())
//	//fmt.Printf("\nData after save: %s\n", spew.Sdump(data))
//	if !ok {
//		return nil, status.Errorf(
//			codes.Internal,
//			fmt.Sprintf("Cannot convert to OID"),
//		)
//	}
//	data.ID = oid
//	//fmt.Printf("Domain.Create is returning: %s\n", spew.Sdump(data))
//	return data, err
//}
//
////TODO Should Delete physically remove the release or just market as deleted?
////DeleteRelease Deletes a release base upon the provided ObjectID
//func DeleteRelease(oid primitive.ObjectID) error {
//
//	filter := bson.M{"_id": oid}
//	//fmt.Printf("Using filter: %v\n", filter)
//	col, err := db.GetCollection("releases")
//	if err != nil {
//		err = fmt.Errorf("Collection Failed|%d|%s", codes.Internal, err.Error())
//		return err
//	}
//	ctx := context.Background()
//	resp, err := col.DeleteOne(ctx, filter)  // returns result a struct with one element DeletedCount
//	if err != nil {
//		err = fmt.Errorf("Internal Error|%d|Delete of %s failed err: %v", codes.Internal, oid.Hex(), err.Error())
//		return err
//	}
//	if resp.DeletedCount == 0 {
//		err = fmt.Errorf("Not Found|%d|Release not found", codes.NotFound)
//		return err
//	}
//
//
//
//	return err
//}
//
///*TODO: Release Update Should change oly allowed fields and not replace entire Release  History should be manually appended
//	Documents should be updated alone
//*/
////Update  replace the current Release with the new one.
//func Update(rel *m.Release) error {
//
//	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	//defer cancel()
//	collection, err := db.GetCollection("releases")
//	if err != nil {
//		return status.Errorf(
//			codes.Internal,
//			fmt.Sprintf("domain.GetRelease GetCollection failed: %v\n", err),
//		)
//	}
//
//	rel.UpdatedAt = time.Now()
//	filter := bson.M{"_id": rel.ID}
//	_, err = collection.ReplaceOne(context.Background(), filter, rel)
//	if err != nil {
//		return status.Errorf(
//			codes.Internal,
//			fmt.Sprintf("Cannot update Document in database: %v", err),
//		)
//	}
//	return err
//}
//
////func UpdateDocumentCount(id primitive.ObjectID, val int) (*Release, error) {
////	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
////	defer cancel()
////	collection, err := db.GetCollection(settings.CollectionName())
////	if err != nil {
////		fmt.Printf("Collection failed; %s\n", err)
////		return nil, err
////	}
////	filter := bson.M{"_id": id}
////	updateAt := time.Now()
////	update := bson.M{
////		"$inc": bson.M{"documentcount": val},
////		"$set": bson.M{"updatedat": updateAt},
////	}
////	upsert := true
////	after := options.After
////	opt := options.FindOneAndUpdateOptions{
////		ReturnDocument: &after,
////		Upsert: &upsert,
////	}
////
////	result := collection.FindOneAndUpdate(ctx, filter, update, &opt)
////	if result.Err() != nil {
////		fmt.Printf("Update error: %s\n", result.Err())
////		return nil, result.Err()
////	}
////	doc := bson.M{}
////	rel := Release{}
////	decodeErr := result.Decode(&doc)
////	dErr :=result.Decode(&rel)
////	if dErr != nil {
////		fmt.Printf("Decode to struct failed: %s\n", dErr)
////		return &rel, decodeErr
////	}
////	//fmt.Printf("rel: %s\n", spew.Sdump(rel))
////	return &rel, decodeErr
////
////}
