package client

//
//type RecipientSummary struct {
//	ID          primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
//	Name        string              `json:"name"`
//	FirstName	string 				`json:"first_name"`
//	LastName	string				`json:"last_name"`
//	CompanyName string           	`json:"company_name"`
//	Address1    string           	`json:"address1"`
//	Address2    string           	`json:"address2"`
//	City        string           	`json:"city"`
//	State       string           	`json:"state"`
//	Zip         string           	`json:"zip"`
//
//}


//DeliverRelease starts a delivery using the DeliveryService
//func DeliverRelease(ctx context.Context,
//	relID string,
//	recipID string,
//	deviceName string,
//	cust m.Customer,
//	clnt m.Client) (string, error) {
//
//	// If the deviceName is longer than 10 characters it is a device id otherwise it is the name the recipient uses
//
//	ReleaseConnect()
//	//pbRelease := toDelvPBRelease(dr)
//	//pbRecipient := toRelPBRecipient(recip)
//	//pbDevice := toRelPBDevice(device)
//	delvReq := relPB.SubmitDeliveryRequest{}
//
//	//	SubmitDeliveryRequest{
//	//	RelId:          relID,
//	//	RecipientId:    recipID,
//	//	DeviceName: 	deviceName,
//	//	ClientMrn:      clientMRN,
//	//}
//	resp, err := RelClient.SubmitDelivery(ctx, &delvReq)
//	if err != nil {
//		return "", err
//	}
//
//	return resp.GetDelivery().GetId(), nil
//}
//
//func toDelvPBRelease(d *m.Release) *relPB.Release {
//	r := relPB.Release{
//		//Facility: d.Facility,
//		//Id:       d.ID,
//		//Client:   d.Client,
//		//SourceId: d.SourceID,
//		//PatName:  d.Patient.Name,
//	}
//	return &r
//}
//
//func toRelPBDevice(dd *m.DeliveryDevice) *relPB.Device {
//	dev := relPB.Device{}
//	dev.Id				= dd.ID.Hex()
//	dev.Name			= dd.Name
//	dev.Method			= dd.Method
//	//dev.Validated		= dd.Validated.Unix()
//	//
//	//dev.Address 		= dd.Address
//	//dev.ID				= dd.Id.Hex()
//	//dev.Validated		= dd.Validated.Unix()
//	//dev.Brand 			= dd.Brand
//	//dev.ServicePassword	= dd.ServicePassword
//	//dev.ServiceLogin	= dd.ServiceLogin
//	//dev.EmrLogin		= dd.EmrLogin
//	//dev.EmrPassword		= dd.EmrPassword
//	//dev.EmrType			= dd.EmrType
//	//dev.EmrUrl			= dd.EmrURL
//	//dev.Validated		= dd.Validated.Unix()
//	//dev.RecipientID		= dd.RecipientId.Hex()
//	//dev.Model 			= dd.Model
//	//dev.Country			= dd.Country
//	return &dev
//}

//func toRelPBRecipient(r *RecipientSummary) *relPB.Recipient {
//	return &relPB.Recipient{
//		Id:           r.ID.Hex(),
//		Contact:      r.Name,
//		FirstName:		r.FirstName,
//		LastName:		r.LastName,
//		Company:      r.CompanyName,
//		Address1:     r.Address1,
//		Address2:     r.Address2,
//		City:         r.City,
//		State:        r.State,
//		Postal:       r.Zip,
//	}
//}
