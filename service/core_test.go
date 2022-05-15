package service

// import (
// 	"fmt"

// 	"github.com/davecgh/go-spew/spew"

// 	//"github.com/davecgh/go-spew/spew"
// 	"os"

// 	. "github.com/smartystreets/goconvey/convey"

// 	//log "github.com/sirupsen/logrus"
// 	"testing"

// 	core "gitlab.com/dhf0820/ids_core_service/connect_core"
// )

// func TestInitialize(t *testing.T) {
// 	//t.Parallel()
// 	fmt.Printf("\n\nTestInititialize\n")
// 	Convey("Initializes the Release system getting configuration from Core", t, func() {
// 		fmt.Printf("\n\n--- Convey\n")
// 		os.Setenv("CONFIG_ADDRESS", "localhost:19800")
// 		os.Setenv("SERVICE_NAME", "release")
// 		os.Setenv("SERVICE_VERSION", "local")
// 		conf, err := Initialize("local")
// 		So(err, ShouldBeNil)
// 		So(conf, ShouldNotBeNil)
// 		//So(conf.Name, ShouldEqual, "delivery")
// 		fmt.Printf("\ncfg: %s\n\n", spew.Sdump(conf))
// 		for i, me := range conf.MyEndPoints {
// 			fmt.Printf("MyConnection: %d - %s\n", i, spew.Sdump(me))
// 		}
// 		fmt.Printf("DataConnector: %s\n", spew.Sdump(conf.DataConnector))

// 		relId := "5fc3e3b9a652f88181ef5711"
// 		rel, err := GetRelease(relId)
// 		So(err, ShouldBeNil)
// 		So(rel, ShouldNotBeNil)
// 		fmt.Printf("Release: %s\n", spew.Sdump(rel))
// 	})
// }

// //func TestGetConfig(t *testing.T) {
// //	//t.Parallel()
// //	fmt.Printf("\n\nTestGetConfig\n")
// //	Convey("Subject: Client Finds possibly multiple Recipients", t, func() {
// //		os.Setenv("ENV_RELEASE_TEST", "/Users/dhf/work/roi/services/release_service/config/config.json")
// //		fmt.Printf("\n\n--- Convey")
// //		conf, err := Initialize("ENV_RELEASE_TEST")
// //		So(err, ShouldBeNil)
// //		So(conf,ShouldNotBeNil )
// //		So(conf.Name, ShouldEqual, "release")
// //		fmt.Printf("\ncfg: %s\n\n", spew.Sdump(conf))
// //		for i, me := range conf.MyEndPoints {
// //			fmt.Printf("MyConnection: %d - %s\n", i, spew.Sdump(me))
// //		}
// //	})
// //}

// func TestGetServiceConfig(t *testing.T) {
// 	//t.Parallel()
// 	InitTest()
// 	Initialize("local")
// 	fmt.Printf("\n\nTestGetConfig\n")
// 	Convey("Retrieves the Release configuration from core", t, func() {
// 		//os.Setenv("ENV_DELIVERY_TEST", "/Users/dhf/work/roi/services/delivery_service/config/config.json")
// 		fmt.Printf("\n\n--- Convey")
// 		conf, err := core.GetServiceConfig("release", "test")
// 		So(err, ShouldBeNil)
// 		So(conf, ShouldNotBeNil)
// 		//So(conf.Name, ShouldEqual, "delivery")
// 		fmt.Printf("\ncfg: %s\n\n", spew.Sdump(conf))
// 		for i, me := range conf.MyEndPoints {
// 			fmt.Printf("MyConnection: %d - %s\n", i, spew.Sdump(me))
// 		}
// 	})
// }

// func TestConnectToCore(t *testing.T) {
// 	fmt.Printf("\n\nTestGetConfig\n")
// 	InitTest()
// 	Initialize("local")
// 	Convey("Connects to core via GRPC", t, func() {
// 		//os.Setenv("CONFIG_ADDRESS", "localhost:19800")
// 		client, err := core.ConnectToCore()
// 		//client, err := ConnectToCore()
// 		So(err, ShouldBeNil)
// 		So(client, ShouldNotBeNil)
// 		//req := pb.ConfigRequest{
// 		//	Name: "delivery",
// 		//	Version: "test",
// 		//}
// 		conf, err := core.GetServiceConfig("release", "test")
// 		//resp, err := client.GetServiceConfig(context.Background(), &req)
// 		So(err, ShouldBeNil)
// 		//So(resp, ShouldNotBeNil)
// 		//conf := toServiceConfig(resp.GetServiceConfig())
// 		fmt.Printf("Received: %s\n", spew.Sdump(conf))
// 	})

// }
