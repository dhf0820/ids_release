package service

// import (
// 	"fmt"

// 	//"github.com/davecgh/go-spew/spew"
// 	. "github.com/smartystreets/goconvey/convey"
// 	//log "github.com/sirupsen/logrus"
// 	"testing"
// )

// func TestOpenDB(t *testing.T) {
// 	//t.Parallel()
// 	fmt.Printf("\n\nTestOpenDB\n")
// 	InitTest()

// 	Convey("Subject: Open the mongo DB", t, func() {
// 		fmt.Printf("\n\n--- Convey")

// 		conf, err := Initialize("local")

// 		So(err, ShouldBeNil)
// 		So(conf,ShouldNotBeNil )
// 		mongo := OpenDB()
// 		So(mongo, ShouldNotBeNil)
// 		c, err := GetCollection("releases")
// 		So(err, ShouldBeNil)
// 		So(c, ShouldNotBeNil)
// 	})
// }
