package service

import (
	// "bytes"
	// //pm "github.com/dhf0820./ds_model/patient"
	// "io/ioutil"
	// "strconv"

	// //"context"
	// "encoding/json"
	"fmt"
	// "math/rand"
	// "net/http"
	//"context"
	//pb "gitlab.com/dhf0820/ids_core_service/protobufs/corePB"
	//grpc "google.golang.org/grpc"
	//"github.com/joho/godotenv"
	//"os"
	//"github.com/davecgh/go-spew/spew"
	. "github.com/smartystreets/goconvey/convey"

	//docMod "github.com/dhf0820./ds_model/document"

	//log "github.com/sirupsen/logrus"
	"testing"
)

type Upload struct {
	ID             string `json:"ID"`
	Size           int    `json:"Size"`
	SizeIsDeferred bool   `json:"SizeIsDeferred"`
	Offset         int
	MetaData       map[string]interface{}
	IsFinal        bool
	IsPartial      bool
	PartialUploads bool
}

// type HookData struct {
// 	Upload 	Upload 	`json:"Upload"`
// }
func TestProcessJson(t *testing.T) {
	//t.Parallel()
	//fmt.Printf("\n\nTestPocessJson\n")
	//Convey("unmarshal a json string", t, func() {
	//	u := docMod.Upload{}
	//	u.ID = "32f46e9f3052486a022c00548e3b27d1"
	//	u.Size = 27181
	//	u.SizeIsDeferred = true
	//	hd := HookData{}
	//	hd.Upload = u
	//	hb, err := json.Marshal(hd)
	//	ud := HookData{}
	//	err = json.Unmarshal(hb, &ud)
	//	fmt.Printf("ud := %s\n", hb)
	//	//fmt.Printf("um: %s\n", spew.Sdump(ud))
	//
	//	b := []byte("{\"Upload\":{\"ID\":\"32f46e9f3052486a022c00548e3b27d1\",\"Size\":27181,\"SizeIsDeferred\":true}}") //,"Offset":27181,"MetaData":{"Client":"test","Description":"Xray Left Knee"}}}"
	//
	//	fmt.Printf("sample: %s\n", b)
	//	results := []byte("{\"Upload\":{\"ID\":\"32f46e9f3052486a022c00548e3b27d1\",\"Size\":27181,\"SizeIsDeferred\":false,\"Offset\":27181,\"MetaData\":{\"Client\":\"test\",\"Description\":\"Xray Left Knee\"}}}")
	//
	//	//os.Setenv("ENV_CORE_TEST", "/Users/dhf/work/roi/services/core_service/config/core_test.json")
	//	fmt.Printf("\n\n--- Convey\n")
	//	hook := HookData{}
	//	err = json.Unmarshal(results, &hook)
	//	So(err, ShouldBeNil)
	//	//fmt.Printf("hook: %s\n", spew.Sdump(hook))
	//})
}

//var r  {"Upload":{"ID":"32f46e9f3052486a022c00548e3b27d1","Size":27181,"SizeIsDeferred":false,"Offset":27181,"MetaData":{"Client":"test","Description":"Xray Left Knee"}}}

func TestSubmitRelease(t *testing.T) {
	//t.Parallel()
	fmt.Printf("\n\nTestSubmitrelease\n")
	Convey("Upload a document", t, func() {
	})

}

//func executeRequest(req *http.Request) *httptest.ResponseRecorder {
//	rr := httptest.NewRecorder()
//	http.Post
//	a.Router.ServeHTTP(rr, req)
//
//	return rr
//}

/*func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}*/
