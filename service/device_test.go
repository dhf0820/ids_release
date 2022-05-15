package service

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestGetDevice(t *testing.T) {
	t.Parallel()
	InitTest()

	fmt.Printf("\n\nTestGetDevice\n")

	Convey("Subject: Get a Device", t, func() {
		fmt.Printf("Given there is a valid Device\n")
		Convey("Get the Device", func() {
			fmt.Printf("Get a Device\n")
			deviceId, _ := primitive.ObjectIDFromHex("61b551e2c20a4aad4494a758")

			device, err := GetDevice(deviceId)
			So(err, ShouldBeNil)
			So(device, ShouldNotBeNil)
			So(device.Name, ShouldEqual, "email")
			//fmt.Printf("Release: %s\n", spew.Sdump(rel))
		})
	})

}


