package service

//
//import (
//	"context"
//	"fmt"
//	"github.com/davecgh/go-spew/spew"
//	"github.com/dhf0820/ids_release_service/internal/domain"
//	m "github.com/dhf0820./ds_model"
//
//	//"github.com/dhf0820/ids_release_service/pkg/common"
//	dc "github.com/dhf0820/ids_delivery/client"
//	"go.mongodb.org/mongo-driver/bson/primitive"
//	"google.golang.org/grpc/codes"
//	"google.golang.org/grpc/status"
//
//	//"github.com/davecgh/go-spew/spew"
//
//	//delvClient "github.com/dhf0820/ids_release_service/internal/clients/delivery_client"
//	"github.com/dhf0820/ids_release_service/protobufs/relPB"
//)
//
//func (s *ReleaseServiceServer) SubmitDelivery(ctx context.Context,
//				req *relPB.SubmitDeliveryRequest) (*relPB.DeliveryResponse, error) {
//	fmt.Printf("\nSubmitting a release for delivery\n")
//
//	relId, _ := primitive.ObjectIDFromHex( req.GetReleaseId())
//	release, err := domain.GetReleaseById(relId)
//
//	//deviceId, _ := primitive.ObjectIDFromHex(req.GetDeviceId())
//	_, err := GetDeviceDetail(ctx, req.GetDeviceId())
//	if err != nil {
//		return nil, status.Errorf(
//			codes.NotFound,
//			fmt.Sprintf("Release: %s not found", relId),
//		)
//	}
//
//	delv := m.Delivery{}
//	delv.ReleaseId = release.ID.Hex()
//	delv.Documents = fillDocuments(release)
//	delv.ReleaseId = req.GetReleaseId()
//	delv.Customer = &m.Customer{}
//	delv.Customer.Code = req.GetCustomer().GetCode()
//	delv.Customer.Facility = req.GetCustomer().GetFacility()
//	delv.Customer.UserId = req.GetCustomer().GetUserId()
//	delv.Customer.Name = req.GetCustomer().GetName()
//	delv.Client = &m.Client{}
//	delv.Client.Source = req.GetClient().GetSource()
//	delv.Client.SourceId = req.GetClient().GetSourceId()
//	delv.RecipientId = req.GetRecipientId()
//	delv.DeviceId 	= req.GetDeviceId()
//
//
//
//	//TODO: Validate the release, recipient and device are valid before sending
//	fmt.Printf("Submitting Delivery: %s\n", spew.Sdump(delv))
//	delivery, err := dc.SubmitDelivery(ctx, &delv)
//	if err != nil {
//		fmt.Printf("Submit Error: %v\n", err)
//		return nil, status.Errorf(
//			codes.Internal,
//			fmt.Sprintf("Submit Delivery error: %v", err),
//		)
//	}
//
//	resp := &relPB.DeliveryResponse{}
//	resp.Delivery = ToRelPbDelivery(delivery)
//	return resp, nil
//}
//
//func fillDocuments(rel *m.Release) []*m.Document{
//	return rel.Documents
//}
//
//func fillDocument(doc relPB.Document) m.Document {
//	d := m.Document{}
//	d.ImageType 		= doc.GetImageType()
//	d.ImageRepository	= doc.GetImageRepository()
//	d.ImageURL 			= doc.GetImageUrl()
//	d.DocClass			= doc.GetDocClass()
//	d.Description		= doc.GetDescription()
//	d.Version			= doc.GetVersion()
//	//d.Id 				= doc.GetId()
//	return d
//
//}
////	//TODO: CallDomain with the info for a deliveryHistory
////
////	fmt.Printf("Getting Release\n")
////	rel, err := domain.GetReleaseById(relID)
////	if err != nil {
////		rel, err = domain.GetReleaseByForeignId(customer, relID)
////		if err != nil {
////			return nil, status.Errorf(
////				codes.NotFound,
////				fmt.Sprintf("domain.GetRelese provided ID: %s is invalid", relID),
////			)
////		}
////	}
////
////	fmt.Printf("Getting Recipient\n")
////	recipID := req.GetRecipientId()
////	recip, err := recipClient.GetRecipient(ctx, recipID)
////	if err != nil {
////		if err != nil {
////			return nil, status.Errorf(
////				codes.NotFound,
////				fmt.Sprintf("domain.GetRecipient provided ID: %s is invalid", relID),
////			)
////		}
////	}
////
////	fmt.Printf("Recipient: %s\n", spew.Sdump(recip))
////	fmt.Printf("Getting DeviceName: %d\n", len(recip.Devices))
////	deviceName := req.GetDeviceName()
////	dName := ""
////	for _, d := range recip.Devices {
////		fmt.Printf("Looking for device: %s  have %s\n", deviceName, d.Name)
////		if d.Name == deviceName {
////			dName = deviceName
////			break
////		} else {
////			continue
////		}
////	}
////	if dName == "" {
////		return nil, status.Errorf(
////			codes.NotFound,
////			fmt.Sprintf("domain.GetRecipient device: %s is invalid", deviceName),
////		)
////	}
////	fmt.Printf("%s is valid\n", deviceName)
////
////
////	clientMRN := req.GetClientMrn()
////
////
////
////
////
////
/////* when we receive the delivery id back from the DeliveryService, create a new DeliveryHistory with the id received,
////the submit time, submitted status in history in the release. Also set the indicator in the release that no mr documents
////can be added to this release.  If there needs to be additional documents added, create a new release or do not allow
////as error
////*/
////
////
////	//fmt.Printf("Checking there are documents to delivery\n")
////	//if len(rel.Documents) == 0 {
////	//	err := fmt.Errorf("Release: [%s] has no documents and can not be submitted.", relID)
////	//	return nil, err
////	//}
////
////
////	//delID, err := delvClient.CreateDelivery(ctx, relID, recipID, deviceName, clientMRN)
////	//fmt.Printf("Create Delivery err: %v\n", err)
////	//fmt.Printf("DelID: [%s]\n", delID)
////	//resp := &relPB.SubmitDeliveryResponse{
////	//	DeliveryId: delID,
////	//}
////	//
////	////TODO: CallDomain with the info for a deliveryHistory
////	//return resp, nil
////	// Create a Delivery history and add to Release showing deliveryid, Recip, device id, Submitted status and time
////
////	// Not sure what is updated
////	//fmt.Printf("Update Release")
////	//err = rel.Update()
////	//if err != nil {
////	//	err = fmt.Errorf("update Release [%s] for number of documents failed: %v", relID, err)
////	//	return nil, err
////	//}
////
////
////	//fmt.Printf("Calling From pbDevice\n")
////
////	//dDev, _ := domain.FromPBDevice(req.GetDevice())
////	//fmt.Printf("Calling FromPBRecipient\n")
////	//recp, _ := domain.FromPBRecipient(req.GetRecipient())
////	//ddr := domain.DeliveryRequest{
////	//	ReleaseID: 		relID,
////	//	Recipient: 		*recp,
////	//	Device: *dDev,
////	//}
////	//fmt.Printf("Request: %s\n", spew.Sdump(ddr))
////	//devID, err := ddr.Submit()
////	//if err != nil {
////	//	fmt.Printf("Submit Error: %s\n", err)
////	//	return nil, err
////	//}
////
//////}
