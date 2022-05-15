package service

import (
	"encoding/json"
	//"encoding/json"
	"fmt"

	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
	//mod "github.com/dhf0820./ds_model"
	"github.com/dhf0820/ids_release_service/messaging"
	//mod "github.com/dhf0820./ds_model"
	"os"

	common "github.com/dhf0820./ds_model/common"
	//"strconv"
	//"time"
	//"os"
	//"io/ioutil"
	//mod "github.com/dhf0820./ds_model"
	//"github.com/dhf0820/ca_link_service/messaging"
)

/* DeliveryConfig is the basic configuration for this all services available to the current service
what it can talk to and who can talk to it*/
//type ServiceConfig struct {
//	Name string `json:"name"`
//	//Messaging  		Messaging					// move to endpoints
//	DataConnector    *common.DataConnector  `json:"dataconnector"`
//	Services         []*corepkg.ServiceScope `json:"services" bson:"services"`
//	MyEndPoints      []*common.EndPoint     `json:"myendpoints"`
//	ServiceEndPoints []*common.EndPoint     `json:"serviceendpoints" bson:"endpoints"`
//	ConnectInfo      []*common.ConnectInfo  `json:"connect_info" bson:"connect_info"`
//}
//
//type ServiceScope struct {
//	Name  string `json:"name" bson:"name"`
//	Scope string `json:"scope" bson:"scope"` // min, norm, max
//}
//
//type EndPoint struct { // Replaces BaseUrl
//	Name        string `json:"name" bson:"name"`
//	Label       string `json:"label" bson:"name"`
//	Scope       string `json:"scope, omitempty" bson:"scope"`
//	Protocol    string `json:"protocol" bson:"protocol"` // grpc or amqp
//	Address     string `json:"address" bson:"address"`
//	Port        string `json:"port" bson:"port"`
//	Credentials string `json:"credentials" bson:"credentials"`
//	CertName    string `json:"cert_name" bson:"cert_name"`
//	TLSMode     string `json:"tls_mode" bson:"tls_mode"`
//	DeployMode  string `json:"deploy_mode" bson:"deploy_mode""`
//}
//
//type ConnectInfo struct {
//	//ID    string `json:"id" bson:"id,omitempty"`
//	Name  string `json:"name"`
//	Value string `json:"value"`
//}

//type DataConnector struct {
//	Server     string     `json:"server"`
//	User       string     `json:"user"`
//	Password   string     `json:"passwword"`
//	Database   string     `json:"database"`
//	Collection string     `json:"collection"`
//	Fields     []*KVField `json:"fields"`
//}

//type KVField struct {
//	Name  string `json:"name"`
//	Value string `json:"value"`
//}

type Messaging struct {
	Name   string
	VsAMQP string
	Client *messaging.MessagingClient
}

var (
	Conf *common.ServiceConfig //*config
	//MsgClients []*Messaging
	//*messaging.MessagingClient
	//GWConfig *m.DeliveryConfig
)

func GetServiceEndPoint(value string) *common.EndPoint {
	//config := GetConfig()
	//fmt.Printf("GetConfig : %s\n", spew.Sdump(config))
	endPoints := GetConfig().ServiceEndPoints
	for _, ep := range endPoints {
		//fmt.Printf("Looking at %s for %s\n", ep.Name, value)
		if ep.Name == value {
			//fmt.Printf("Found EndPoint: %s\n", ep.Name)
			return ep
		}
	}
	log.Errorf("endPoint: %s was not found", value)
	return nil
}

//
//func GetMyEndPoint(value string) *common.EndPoint {
//	endPoints := GetConfig().MyEndPoints
//	for _, ep := range endPoints {
//		//fmt.Printf("Looking at %s for %s\n", ep.Name, value)
//		if ep.Name == value {
//			//fmt.Printf("Found EndPoint: %s\n", ep.Name)
//			return ep
//		}
//	}
//	return nil
//}

//func GetMsgClient(name string) *messaging.MessagingClient {
//	for _, msg := range MsgClients {
//		if msg.Name == name {
//			return msg.Client
//		}
//	}
//	return nil
//}

//func InitMessaging(amqp *EndPoint) *messaging.MessagingClient {
//	//fmt.Printf("\n---Initializing %s\n\n", spew.Sdump(amqp))
//	msg := Messaging{}
//	msg.Name = amqp.Name
//	msg.VsAMQP = amqp.Address
//	msg.Client = &messaging.MessagingClient{}
//	msg.Client.ConnectToBroker(msg.VsAMQP)
//	MsgClients = append(MsgClients, &msg)
//	return msg.Client
//}
//TODO: Replace GetConfig in core with calls to core to get the centeralized information

func Initialize(serviceVersion, company string) (*common.ServiceConfig, error) {
	var err error
	//var delay time.Duration
	//var maxAttempts int
	fmt.Printf("\n\n\n -------Initiallizing ReleaseService -----\n\n")
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceVersion == "" {
		serviceVersion = os.Getenv("SERVICE_VERSION")
	}
	companyName := os.Getenv("COMPANY_NAME")
	if companyName == "" {
		companyName = "demo"
	}
	if serviceName == "" {
		serviceName = "release"
		os.Setenv("SERVICE_NAME", "release")
	}
	if serviceVersion == "" {
		serviceVersion = "local_test"
		os.Setenv("SERVICE_VERSION", serviceVersion)
	}
	if company == "" {
		company = os.Getenv("COMPANY")
		os.Setenv("COMPANY", company)
	}

	api := os.Getenv("API")
	if api == "" {
		api = "RESTFUL"
	}
	//configAddress := os.Getenv("CONFIG_ADDRESS")
	//fmt.Printf("Core address: %s\n", configAddress)
	//fmt.Printf("waiting to retrieve the confguration from core\n")
	Conf, err = GetServiceConfig(serviceName, serviceVersion, company, api)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("\n\n---GetServiceConfig returned\n")
	//fmt.Printf("\n\n----config: %s]\n", spew.Sdump(Conf))
	OpenDB()
	return Conf, err
}

type ConfigResp struct {
	Status  int                  `json:"status"`
	Message string               `json:"message"`
	Config  common.ServiceConfig `json:"config"`
}

// Get the service configfrom core. Default is Restful, Optional is GRPC
func GetServiceConfig(name, version, company, mode string) (*common.ServiceConfig, error) {
	var cfg ConfigResp
	var err error
	//var unmarshalErr *json.UnmarshalTypeError
	var bdy []byte
	cfg = ConfigResp{}
	fmt.Printf("GetServiceConfig-198 name: %s, version: %s, company: %s, mode: %s\n",
		name, version, company, mode)
	//coreName := strings.ReplaceAll(os.Getenv("CORE_NAME_PORT"), " ","")
	api := os.Getenv("API")
	configAddr := os.Getenv("CONFIG_ADDRESS")
	if api == "" || api == "RESTFUL" {

		url := fmt.Sprintf("%s/config?name=%s&version=%s&company=%s", configAddr, name, version, company)
		//url := fmt.Sprintf("http://localhost:19900/api/v1/config?name=%s&version=%s&company=%s", name, version, company)
		fmt.Printf("\n#####Config url: %s\n", url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Get config returned error: %v\n", err)
			return nil, err
		}
		defer resp.Body.Close()
		//cfg = mod.ServiceConfig{}
		bdy, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		//fmt.Printf("raw string: %s\n",string(bdy))
		err = json.Unmarshal(bdy, &cfg)
		if err != nil {
			return nil, err
		}
		//fmt.Printf("Config JSON: %s\n", spew.Sdump(cfg))

	} else {
		return nil, fmt.Errorf("%s is not supported yet", mode)
	}
	Conf = &cfg.Config
	//fmt.Printf("CFG: = %s\n", spew.Sdump(cfg.Config))
	//fmt.Printf("\n\n###Conf: = %s\n", spew.Sdump(Conf))

	return &cfg.Config, err
}
func GetConfig() *common.ServiceConfig {

	//fmt.Printf("##### GetConf: %s\n", spew.Sdump(Conf))
	return Conf
}

//func getConf() *ServiceConfig {
//	V = viper.New()
//	os.Getenv("ROI_PREFIX")
//	V.SetEnvPrefix("roi")
//	V.AutomaticEnv()
//
//	V.SetConfigName("config.json")
//	V.AddConfigPath(os.Getenv("ENV_ROI")) // specifies where the location of the config file
//	//V.AddConfigPath("../config")
//	//V.AddConfigPath("./config")
//	//V.AddConfigPath("/etc/config")
//
//	V.SetConfigType("json")
//
//	err := V.ReadInConfig()
//
//	if err != nil {
//		fmt.Printf("%v", err)
//	}
//
//	Conf = &ServiceConfig{}
//	err = V.Unmarshal(Conf)
//	if err != nil {
//		fmt.Printf("unable to decode into config struct, %v", err)
//	}
//	fmt.Printf("\n\n--Core Conf: %s\n", spew.Sdump(Conf))
//
//	return Conf
//}
//
//func GetConfig() *ServiceConfig{
//	fmt.Printf("\n\n--GetConf: %s\n", spew.Sdump(Conf))
//	return Conf
//}

//func DeliveryConfigure() *DeliveryConfig {
//	//Config = fillConfig()
//	fmt.Printf("CaLinkConfigure called\n")
//	if Conf == nil {
//		fmt.Printf("\nConfiguring CaLink Service\n")
//		Conf = getConf()
//		Conf.Messaging.Client = &messaging.MessagingClient{}
//		//GWConfig.MessagingClient = &messaging.MessagingClient{}
//		if Conf.Messaging.Client == nil {
//			fmt.Printf("Messaging.Client is not configured\n")
//		}
//		fmt.Printf("AMQP: [%s]\n", Conf.Messaging.VsAMQP)
//		Conf.Messaging.Client.ConnectToBroker(Conf.Messaging.VsAMQP)
//	} else {
//		fmt.Printf("ChartArchive Link Service Is Currently configured: %s\n", spew.Sdump(GetDelivery()))
//	}
//	return Conf
//}
//
///*GetDelivery returns the active Delivery configation. THis includes CA Database, Messaging and list of services
//available.
//*/
//func GetDelivery() *DeliveryConfig {
//	return Conf
//}
//
//// //Config is the easy methog to access the configurations
//// func GetConfig() *DeliveryConfig {
//// 	return Conf
//// }

//var (
//	V *viper.Viper
//)

// Read the config file from the current directory and marshal
// into the conf config struct.
