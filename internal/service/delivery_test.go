package service

//import (
//	"fmt"
//	//db "gitlab.com/dhf0820/ids_delivery_service/db/vsmongo"
//	"os"
//	"context"
//	//"gitlab.com/dhf0820/ids_delivery_service/protobufs/delPB"
//	. "github.com/smartystreets/goconvey/convey"
//	dc "gitlab.com/dhf0820/ids_delivery_service/client"
//	//"google.golang.org/grpc"
//	"testing"
//)
//
//func TestGetDelivery(t *testing.T) {
//	fmt.Println("TestGetDeliveryById")
//	os.Setenv("ENV_ROI", "/Users/dhf/work/roi/services/release_service/config")
//	//settings.InitDeliveryService("")
//	Convey("Subject: Get a Delivery", t, func() {
//		delID := "5fd52440b5a67e2d1414468a"
//		Convey("Given: A valid Delivery", func() {
//			resp, err := dc.GetDelivery(context.Background(),delID)
//			So(err, ShouldBeNil)
//			So(resp, ShouldNotBeNil)
//			//delv := resp.GetDelivery()
//			//So(delv.GetId(), ShouldEqual, relID)
//		})
//	})
//}
//

//func TestCreateDelivery(t *testing.T) {
//	fmt.Println("TestCreateDelivery")
//
//	Convey("Subject: Create a Delivery", t, func() {
//		d := NewDelivery()
//		//fmt.Printf("NewDeliveryTest: %s\n", spew.Sdump(d))
//		Convey("Given: A valid Delivery", func() {
//			delv, err := dc.CreateDelivery(context.Background(), d)
//			So(err, ShouldBeNil)
//			So(delv, ShouldNotBeNil)
//			userID := delv.Customer.UserID
//			So(userID, ShouldEqual, "tfrench")
//			//fmt.Printf("\n\nNew Delivery - 35: %s\n", spew.Sdump(delv))
//		})
//	})
//}
