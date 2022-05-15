package connect


import (
	"fmt"
	//"go.mongodb.org/mongo-driver/bson"

	//"context"
	//"github.com/davecgh/go-spew/spew"
	. "github.com/smartystreets/goconvey/convey"

	"testing"
)

func TestConnect(t *testing.T) {
	fmt.Printf("----TestConnectToRelease\n")
	Convey("Connects to Release via GRPC", t, func() {
		conn, err := Connect("ihids.vertisoft.com", "9210","none")
		So(err, ShouldBeNil)
		So(conn, ShouldNotBeNil)
		So(releaseClientPtr, ShouldNotBeNil)
		So(*releaseClientPtr, ShouldEqual, conn)
	})
}

