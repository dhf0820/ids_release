package service

//
//import (
//	"encoding/json"
//	"fmt"
//	//mod "github.com/dhf0820./ds_model"
//	"io/ioutil"
//	"os"
//
//	//"github.com/davecgh/go-spew/spew"
//	log "github.com/sirupsen/logrus"
//	//"go.mongodb.org/mongo-driver/bson/primitive"
//	//"gitlab.com/dhf0820/ids_release_service/config"
//	//"context"
//	//"fmt"
//	//"github.com/davecgh/go-spew/spew"
//	//db "gitlab.com/dhf0820/ROIEmail/db"
//	//"go.mongodb.org/mongo-driver/bson"
//	//"go.mongodb.org/mongo-driver/mongo/options"
//	//"time"
//	//log "github.com/sirupsen/logrus"
//	//
//	//"google.golang.org/grpc/codes"
//	//"google.golang.org/grpc/status"
//	//"time"
//)
//
//
//
//type ReleaseConfig struct {
//	Name 					string
//	DataConnector DataConnector
//	Services   		[]ServiceConfig
//	EndPoints     []EndPoint         `json:"endpoints" bson:"end_points"`
//	ConnectInfo   []ConnectInfo      `json:"connect_info" bson:"connect_info"`
//}
//
//type ServiceConfig struct {
//	Name    string
//	Port    string
//	Address string
//	Fields  []KVField
//	//scope
//	//protocol
//}
//
//type EndPoint struct { // Replaces BaseUrl move ConnectorURL into this
//	Label  string `json:"label"`
//	Type   string `json:"type" bson:"type"` // grpc or amqp
//	URL    string `json:"url" bson:"url"`
//	Access string `json:"access" bson:"access"`
//}
//
//type KVField struct {
//	Name  string
//	Value string
//}
//
//type ConnectInfo struct {
//	Name  string `json:"name"`
//	Value string `json:"value"`
//}
//
//type  DataConnector struct {
//	Server     string
//	User     string
//	Password string
//	Database string
//}
//
//var currentConfig *ReleaseConfig
//
///*InitializeConfiguration read the configuration file in ./config into cfg
// */
//func InitializeConfig(cfgFile string) (*ReleaseConfig, error) {
//	// read the configuration file in ./config into cfg
//	//if cfgFile == "" {
//	//	os.Getenv()
//	//}
//	//cfgFile := "$home/config/configuration.json"
//	//fmt.Printf("Config: %s\n", cfgFile)
//	if cfgFile == "" {
//		cfgFile = os.Getenv("RELEASE_ENV")
//	}
//	data, err := ioutil.ReadFile(cfgFile)
//	if err != nil {
//		log.Errorf("InitializeConfig Error reading [%s]: %v", cfgFile, err)
//	}
//	//fmt.Printf("Read ConfigFile: %s\n", data)
//	res := &ReleaseConfig{}
//	err = json.Unmarshal(data, res)
//	if err != nil {
//		log.Errorf("Initialize unmarshal failed: %v\n", err)
//	}
//	currentConfig = res
//	//fmt.Printf("parsed config:\n %s\n", spew.Sdump(currentConfig))
//	return res, err
//}
//
//func GetConfig() *ReleaseConfig {
//	fmt.Printf("Getting ReleaseConfig\n")
//	if currentConfig == nil {
//		log.Errorf("GetConfig: Configuration has not been initialized, doing so using %s\n", os.Getenv("CA_LINK"))
//		_, err := InitializeConfig(os.Getenv("CA_LINK"))
//		if err != nil {
//			err = fmt.Errorf("Configuration is not initialized")
//			return nil
//		}
//	}
//	//fmt.Printf("Returning config:\n%s\n", spew.Sdump(currentConfig))
//
//	return currentConfig
//}
//
////TODO: Replace settings calls in core with calls to core to get the centeralized information
//func GetService(key string) (value *ServiceConfig, err error) {
//	var srvVal ServiceConfig
//
//	//LogMessage(&payload, "Detailed", "Info", "Checking config value for field: "+field, payload.Config.Core_log_url)
//	svs := GetConfig().Services
//	for _, sv := range svs {
//		switch {
//		case sv.Name == key:
//			srvVal = sv
//		}
//	}
//	return &srvVal, nil
//}
//
//func (cfg *ReleaseConfig) GetEndPointInfo(key string) (value EndPoint, err error) {
//	var epVal EndPoint
//	//LogMessage(&payload, "Detailed", "Info", "Checking config value for field: "+field, payload.Config.Core_log_url)
//	eps := cfg.EndPoints
//	//fmt.Printf("Looking in CFG:\n%s\n", spew.Sdump(cfg))
//	//fmt.Printf("num endpoints: %d\n", len(eps))
//	for _, ep := range eps {
//		switch {
//		case ep.Label == key:
//			epVal = ep
//		}
//	}
//	return epVal, nil

//}
//
//func (cfg *ReleaseConfig) GetConnectInfo(key string) (value ConnectInfo, err error) {
//	var ciVal ConnectInfo
//	//LogMessage(&payload, "Detailed", "Info", "Checking config value for field: "+field, payload.Config.Core_log_url)
//	cis := cfg.ConnectInfo
//	for _, ci := range cis {
//		switch {
//		case ci.Name == key:
//			ciVal = ci
//		}
//	}
//	return ciVal, nil
//}
//
////func (cfg *ReleaseConfig) GetConfigField(key string) (value ConnectorField, err error) {
////	var fldVal ConnectorField
////	//LogMessage(&payload, "Detailed", "Info", "Checking config value for field: "+field, payload.Config.Core_log_url)
////	flds := cfg.Fields
////	for _, fld := range flds {
////		fmt.Printf("Looking at Field: %s\n\n", spew.Sdump(fld))
////		switch {
////		case fld.Name == key:
////			fldVal = fld
////		}
////	}
////	return fldVal, nil
////}
//
//
//
//
///* GetField returns the Field of the config based upon the ServiceConfig specified
// */
//func (s *ServiceConfig) GetField(key string) (value *KVField, err error) {
//	var fldVal *KVField
//	//LogMessage(&payload, "Detailed", "Info", "Checking config value for field: "+field, payload.Config.Core_log_url)
//
//	flds := s.Fields
//	for _, fld := range flds {
//		fmt.Printf("looking at [%s]\n", fld)
//		switch {
//		case fld.Name == key:
//			fmt.Printf("  found: %s\n", fld)
//			return &fld, nil
//		}
//	}
//	return fldVal, nil
//}
//
