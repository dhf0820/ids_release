package service

//import (
//	"context"
//	"fmt"
//	"github.com/davecgh/go-spew/spew"
//	"gitlab.com/dhf0820/ids_delivery_service/protobufs/delPB"
//	mod "github.com/dhf0820./ds_model"
//	"go.mongodb.org/mongo-driver/bson/primitive"
//)
//
////GetDeviceDetail returns the device detail for an ID string provided
//func GetDeviceDetail(ctx context.Context, id string) (*mod.DeviceDetail, error) {
//	deliveryGetConnection()
//
//	req := delPB.GetDeviceDetailRequest{}
//	req.Id  = id
//	resp, err := delvClient.GetDeviceDetail(ctx, &req)
//	if err != nil {
//		return nil, err
//	}
//	dev := toDeviceDetail(resp.GetDevice())
//	return dev, nil
//}
//
//func toPbDevice(d *mod.DeviceDetail) *delPB.Device {
//	pd := delPB.Device{}
//	pd.Id 						= d.ID.Hex()
//	pd.Method					= d.Method
//	pd.Name 					= d.Name
//	pd.RecipientId		= d.RecipientId
//	pd.Validated			= d.Validated  //toBoolString(d.Validated)
//
//	//if d.ValidatedAt != nil {
//	pd.ValidatedAt		= d.ValidatedAt
//	//}
//	pd.CreateMethod		= d.CreateMethod
//
//	//pd.CreatedAt			= d.CreatedAt.Unix()
//	pd.UpdatedAt			= d.UpdatedAt.Unix()
//	pd.Fields					= toPbFields(d.Fields)
//	return &pd
//}
//
//func toDeviceDetail(pd *delPB.Device) *mod.DeviceDetail {
//	d := mod.DeviceDetail{}
//	d.Name			= pd.GetName()
//	d.Method		= pd.GetMethod()
//	d.RecipientId	= pd.GetRecipientId()
//	d.ID, _ 		= primitive.ObjectIDFromHex(pd.GetId())
//	d.Validated		= pd.Validated
//	d.ValidatedAt 	= pd.ValidatedAt
//	d.CreateMethod 	= pd.GetCreateMethod()
//	d.Fields 		= toFields(pd.GetFields())
//	return &d
//}
//
//func toPBDeviceDetail(d *mod.DeviceDetail) *delPB.Device {
//	pd := delPB.Device{}
//	fmt.Printf("toPBDeviceDetai: %s\n",spew.Sdump(d))
//	pd.Validated			= d.Validated
//	pd.ValidatedAt			= d.ValidatedAt
//	pd.Method				= d.Method
//	pd.Name					= d.Name
//	pd.RecipientId			= d.RecipientId
//	pd.Id 					= d.ID.Hex()
//	pd.CreateMethod 		= d.CreateMethod
//	pd.UpdatedAt			= d.UpdatedAt.Unix()
//	//pd.CreatedAt			= d.CreatedAt.Unix()
//	pd.Fields 				= toPbFields(d.Fields)
//
//	return &pd
//}
//
//
//
//func toBool(val string) bool {
//	var bval bool
//	if val == "true" {
//		bval = true
//	}else {
//		bval = false
//	}
//	return bval
//}
//
//func toBoolString(val bool) string {
//	var strVal string
//	if val == true {
//		strVal = "true"
//	} else {
//		strVal = "false"
//	}
//	return strVal
//}
//
//func toPbFields(fa []*mod.Field) []*delPB.Field {
//	pfa := []*delPB.Field{}
//	for _, f := range fa {
//		pf := delPB.Field{}
//		pf.UserVisible		= toBoolString(f.UserVisible)
//		pf.Required				= toBoolString(f.Required)
//		pf.Sensitive			= toBoolString(f.Sensitive)
//		pf.Default				= f.Default
//		pf.Value					= f.Value
//		pf.Name 					= f.Name
//		pf.DisplayValue		= f.DisplayValue
//		pf.IsNameVisible	= toBoolString(f.IsNameValue)
//		pf.Label					= f.Label
//		pfa = append(pfa, &pf)
//	}
//	return pfa
//}
//
//func toFields(pfa []*delPB.Field) []*mod.Field {
//	flds := []*mod.Field{}
//	for _, pf := range pfa {
//		f := mod.Field{}
//		f.Name					= pf.GetName()
//		f.Value					= pf.GetValue()
//		f.Default				= pf.GetDefault()
//		f.DisplayValue 	= pf.GetDisplayValue()
//		f.Sensitive			= toBool(pf.GetSensitive())
//		f.Required			= toBool(pf.GetRequired())
//		f.UserVisible		= toBool(pf.GetUserVisible())
//		f.IsNameValue		= toBool(pf.GetIsNameVisible())
//		flds = append(flds, &f)
//	}
//	return flds
//}
