package service

//import (
//	"fmt"
//	"os"
//
//	"github.com/davecgh/go-spew/spew"
//	log "github.com/sirupsen/logrus"
//	"github.com/spf13/viper"
//)
//
///* GatewayConfig is the basic configuration for this all services available to the current service
//what it can talk to and who can talk to it*/
//type ReleaseConfig struct {
//	Name 					string
//	//Messaging  		Messaging					//possible move to endpoints
//	//DataConnector DataConnector
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
//type ConnectInfo struct {
//	Name  string `json:"name"`
//	Value string `json:"value"`
//}
//
//////ConnectField will be build on the fly based upon the Requestor's connector
////type ConnectorField struct {
////	Name         string `json:"name"`
////	Label        string `json:"label"`
////	Default      string `json:"default"`
////	Value        string `json:"value"`
////	DisplayValue string `json:"display_value" bson:"display_value"`
////	Required     bool   `json:"required"`
////	UserVisible  bool   `json:"user_visible" bson:"user_visible"`
////	IsNameValue  bool   `json:"is_name_value" bson:"is_name_value"`
////	Sensitive    bool   `json:"sensitive"`
////}
//
//
////type Messaging struct {
////	VsAMQP string
////	Client *messaging.MessagingClient
////}
//
//type  DataConnector struct {
//	Server     string
//	User     string
//	Password string
//	Database string
//}
//
//type KVField struct {
//	Name  string
//	Value string
//}
//
//var (
//	Conf *ReleaseConfig //*config
//	//GWConfig *m.GatewayConfig
//)
//
////TODO: Replace settings calls in core with calls to core to get the centeralized information
//
//func GatewayConfigure() *ReleaseConfig {
//	//Config = fillConfig()
//	fmt.Printf("CaLinkConfigure called\n")
//	if Conf == nil {
//		fmt.Printf("\nConfiguring CaLink Service\n")
//		Conf = getConf()
//		//Conf.Messaging.Client = &messaging.MessagingClient{}
//		//GWConfig.MessagingClient = &messaging.MessagingClient{}
//		//if Conf.Messaging.Client == nil {
//		//	fmt.Printf("Messaging.Client is not configured\n")
//		//}
//		//fmt.Printf("AMQP: [%s]\n", Conf.Messaging.VsAMQP)
//		//Conf.Messaging.Client.ConnectToBroker(Conf.Messaging.VsAMQP)
//	} else {
//		fmt.Printf("ChartArchive Link Service Is Currently configured: %s\n", spew.Sdump(GetReleaseConfig()))
//	}
//	return Conf
//}
//
///*GetGateway returns the active Gateway configation. THis includes CA Database, Messaging and list of services
//available.
//*/
//func GetReleaseConfig() *ReleaseConfig {
//	return Conf
//}
//
//// //Config is the easy methog to access the configurations
//// func GetConfig() *GatewayConfig {
//// 	return Conf
//// }
//
//var (
//	V *viper.Viper
//)
//
//// Read the config file from the current directory and marshal
//// into the conf config struct.
//func getConf() *ReleaseConfig {
//	if Conf != nil {
//		log.Errorf("Already configured settings returning current")
//		return Conf
//	}
//	V = viper.New()
//	V.SetEnvPrefix("roi")
//	V.AutomaticEnv()
//
//	V.AddConfigPath(os.Getenv("ENV_ROI"))
//	V.AddConfigPath("./config")
//	V.AddConfigPath("../config")
//	V.AddConfigPath("/etc/config")
//	V.SetConfigName("config.json")
//	V.SetConfigType("json")
//	fmt.Printf("Starting Gateway read config\n")
//	err := V.ReadInConfig()
//	fmt.Printf("End Gateway read config\n")
//	if err != nil {
//		fmt.Printf("Settings 170 ReadConfig: %v\n", err)
//	}
//
//	//conf := &config{}
//	conf := &ReleaseConfig{}
//	err = V.Unmarshal(conf)
//	if err != nil {
//		fmt.Printf("unable to decode into config struct, %v", err)
//	}
//
//	//test := V.Get("testing")
//	//fmt.Printf("Test: %s\n", test)
//	//conf.Initialized = "ChartArchiveLink"
//	return conf
//}
//
////TODO: Replace settings calls in core with calls to core to get the centeralized information
//// func GetService(key string) (value *ServiceConfig, err error) {
//// 	var srvVal ServiceConfig
//// 	//LogMessage(&payload, "Detailed", "Info", "Checking config value for field: "+field, payload.Config.Core_log_url)
//// 	svs := Config().Services
//// 	for _, sv := range svs {
//// 		switch {
//// 		case sv.Name == key:
//// 			srvVal = sv
//// 		}
//// 	}
//// 	return &srvVal, nil
//// }
//
//// func (s *ServiceConfig) GetField(key string) (value *Field, err error) {
//// 	var fldVal *Field
//// 	//LogMessage(&payload, "Detailed", "Info", "Checking config value for field: "+field, payload.Config.Core_log_url)
//
//// 	flds := s.Fields
//// 	for _, fld := range flds {
//// 		switch {
//// 		case fld.Name == key:
//// 			fldVal = &fld
//// 		}
//// 	}
//// 	return fldVal, nil
//// }
//
////TODO: put certificates in a secure K/V store as actual key and Certificate
////func GetTLSCert() string {
////	var crt = ""
////	cert := GetConfig().TLS.CertName
////	if cert == "" {
////		crt = fmt.Sprintf("/etc/ssl/vsprivate/%s-selfsigned.crt", cert)
////	} else {
////		crt = "/etc/ssl/vsprivate/localhost-selfsigned.crt"
////	}
////	fmt.Printf("Using Certificate: %s\n", crt)
////	return crt
////}
////
////
////func GetTLSKey() string {
////	var key = ""
////	cert := GetConfig().TLS.GatewayCertName
////	if cert == "" {
////		key = fmt.Sprintf("/etc/ssl/vsprivate/%s-selfsigned.key", cert)
////	} else {
////		key = "/etc/ssl/vsprivate/localhost-selfsigned.key"
////	}
////
////	fmt.Printf("Using Certificate: %s\n", cert)
////	fmt.Printf("Using CertificateKey: %s\n", key)
////	return key
////
////}
//
