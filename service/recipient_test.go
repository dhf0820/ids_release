package service

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestGetRecipient(t *testing.T) {
	t.Parallel()
	InitTest()

	fmt.Printf("\n\nTestAddDocument\n")

	Convey("Subject: Get a Recipient", t, func() {
		fmt.Printf("Given there is a valid recipient\n")
		Convey("Get the recipient", func() {
			fmt.Printf("Get a Recipient\n")
			recipientId, _ := primitive.ObjectIDFromHex("61b53c8389233aab9bc94e31")
			recipient, err := GetRecipient(recipientId)
			So(err, ShouldBeNil)
			So(recipient, ShouldNotBeNil)
			So(recipient.Name, ShouldEqual, "HARMAN, DEBBIE")
			//fmt.Printf("Release: %s\n", spew.Sdump(rel))
		})
	})

}

