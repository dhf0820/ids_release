package service

//import (
//	//"context"
//	//"fmt"
//	//
//	////"github.com/davecgh/go-spew/spew"
//	//dc "gitlab.com/dhf0820/ids_release_service/internal/clients/delivery_client"
//	//
//	////"fmt"
//	//"gitlab.com/dhf0820/ids_release_service/internal/domain"
//)

//sendToDeliveryService accepts the release, recipient, and delivery device returning the id of the delivery and error
//func SendToDeliveryService(ctx context.Context, dr *domain.Release, dRecip *domain.Recipient, device *domain.Device) (string, error) {
//var err error
//delv := delPB.Delivery{}
//delv.ReleaseId = dr.ID.Hex()
//delv.ClientId = dr.Client
//delv.SourceId = dr.SourceID
//delv.UserId = dr.UserID
//delv.Patient = &delPB.Patient{
//	Mrn:    dr.Patient.MRN,
//	Ssn:    dr.Patient.SSN,
//	Dob:    dr.Patient.DOB,
//	Name:   dr.Patient.PatName,
//}
//delv.Documents, err =  dc.FillPbDeliveryDocuments(dr.Documents)
//if err != nil {
//	return "", err
//}
//
// recipient := delPB.Recipient{
// Id: dRecip.ID.Hex(),
// SourceId: 		dRecip.RemoteID,
// Contact:        dRecip.Contact,
// Company:        dRecip.Company,
// Address1:       dRecip.Address1,
// Address2:       dRecip.Address2,
// City:	        dRecip.City,
// State:          dRecip.State,
// Postal:         dRecip.Zip,
// ElectronicId:   dRecip.ElectronicID,
// Facility:      	dRecip.Facility,
// ClientId:       dRecip.ClientID,
// }

// pDevice := delPB.Device{
// Id: 			device.ID,
// Method: 		device.Method,
// DestinationMrn: device.DestinationMRN,
// DeviceName:     device.DeviceName,
// Address:        device.Address,
// Command:  		device.Command,
// Secure:  		device.Secure,

//}
//delvReq := delPB.CreateDeliveryRequest{
//	 Delivery: &delv,
//	 Recipient: &recipient,
//	 Device: &pDevice,
//}

//dc.CreateDelivery(ctx, &delv, &recipient, &pDevice)
//dc.CreateDelivery(ctx, )
//resp, err := dc.CreateDelivery(ctx, dr, dRecip, device)
//
//fmt.Printf("%%%%SendToDelivery returned: %v\n %s\n", err, resp)
//return resp, err
//}
