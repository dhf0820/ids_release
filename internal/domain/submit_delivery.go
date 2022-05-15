package domain

//import (
//	//"fmt"
//	//"github.com/davecgh/go-spew/spew"
//	//"go.mongodb.org/mongo-driver/bson/primitive"
//	//"time"
//
//	//"github.com/dhf0820/ids_release_service/src/settings"
//	//"go.mongodb.org/mongo-driver/bson/primitive"
//)

//func (dr *DeliveryRequest)Submit() (string, error) {
//	//relID, err := primitive.ObjectIDFromHex(dr.ReleaseID)
//	//if err != nil {
//	//	err = fmt.Errorf("RequestID [%s] is invalid: %s", dr.ReleaseID, err.Error())
//	//	return err
//	//}
//
//	//fmt.Printf("\n\n\nServer/Submit started\n\n")
//	//rel, err := GetReleaseById(dr.ReleaseID)
//	//if err != nil {
//	//	err = fmt.Errorf("RequestID [%s] was not found: %s", dr.ReleaseID, err.Error())
//	//	return "", err
//	//}
//	return "", nil
//	//time_now := time.Now()
//	//
//	//his := DeliveryHistory{
//	//	ID: 			primitive.NewObjectID(),
//	//	Recipient: 		dr.Recipient,
//	//	Device: 		dr.Device,
//	//	SubmitTime:  	&time_now,
//	//	DeliveryTime:   nil,
//	//	Status:			"PENDING",
//	//	StatusTime:     &time_now,
//	//	Comments:       fmt.Sprintf("Submitted to %s queue", dr.Device.DeviceName),
//	//}
//
//
//
//	//TODO Queue to proper delivery device
//
//	//fmt.Printf("DeliveryHistory: %s\n", spew.Sdump(his))
//	//rel.Recipients = append(rel.Recipients, his)
//	//err = rel.Update()
//	//if err != nil {
//	//	err = fmt.Errorf("Update Release [%s] failed %s", dr.ReleaseID, err.Error())
//	//	return "", err
//	//}
//	////fmt.Printf("Updated Release: %s\n", spew.Sdump(rel))
//	//return his.ID.Hex(), nil
//}

//DeliveryID     primitive.ObjectID `json:"delivery_id"`            // generated to uniquely id this delivery record
//RecipientID    primitive.ObjectID `json:"recipient_id,omitempty"` // IF set, pull rec our recipient
//Recipient      Recipient          `json:"recipient"`              // If set use this instead of using stored recipient
//Device Device     `json:"delivery_device,omitempty"`
//DeliveryStatus string             `json:"delivery_status"` // InProcess, Success, Error:message
//SubmitTime     time.Time
//ClosedTime     time.Time // Time delivered or closed

/*DeliveryID     primitive.ObjectID `json:"delivery_id"`            // generated to uniquely id this delivery record
RecipientID    primitive.ObjectID `json:"recipient_id,omitempty"` // IF set, pull rec our recipient
Recipient      Recipient          `json:"recipient"`              // If set use this instead of using stored recipient
Device Device     `json:"delivery_device,omitempty"`
DeliveryStatus string             `json:"delivery_status"` // InProcess, Success, Error:message
SubmitTime     time.Time
ClosedTime     time.Time // Time delivered or closed*/
