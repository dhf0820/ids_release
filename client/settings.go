package client

//import (
//	"fmt"
//	"github.com/spf13/viper"
//	"os"
//
//	//"strconv"
//
//	//log "github.com/sirupsen/logrus"
//
//	//"github.com/prometheus/common/log"
//	//"go.mongodb.org/mongo-driver/mongo"
//)
//
//
//type rel_client_config struct {
//	ReleasePort		string
//	ReleaseAddress 	string
//	ReleaseTLSMode     string
//	ReleaseCertName 	string
//	Initialized string
//}
//var (
//	ReleaseConf *rel_client_config
//)
//
//
//// Init does nothing right now. It will fill values from a yml file or config server.
//func RelClientSettingsInit() {
//	fmt.Printf("Intializing Release Client setting\n")
//	if ReleaseConf  == nil {
//		ReleaseConf = getConf()
//	}
//
//	//fmt.Printf("CONFIG: %s\n", spew.Sdump(ReleaseConf))
//
//}
//
//
//func Config() *rel_client_config {
//	return  ReleaseConf
//}
//
//var (
//	V *viper.Viper
//)
//// Read the config file from the current directory and marshal
//// into the conf config struct.
//func getConf() *rel_client_config {
//	V = viper.New()
//	V.SetEnvPrefix("roi")
//	V.AutomaticEnv()
//
//	V.AddConfigPath(os.Getenv("ENV_ROI"))
//	V.AddConfigPath("./config")
//	V.AddConfigPath("/config")
//	V.SetConfigName("config")
//	V.SetConfigType("yml")
//	err := V.ReadInConfig()
//
//	if err != nil {
//		fmt.Printf("%v", err)
//	}
//
//	conf := &rel_client_config{}
//	err = V.Unmarshal(conf)
//	if err != nil {
//		fmt.Printf("unable to decode into config struct, %v", err)
//	}
//	conf.Initialized = "ReleaseClient"
//	return conf
//}
//
//
//func GetTLSCert() string {
//	var crt = ""
//	cert := ReleaseConf.ReleaseCertName
//	if cert == "" {
//		crt = fmt.Sprintf("/etc/ssl/vscerts/%s-selfsigned.crt", cert)
//	} else {
//		crt = "/etc/ssl/vscerts/localhost-selfsigned.crt"
//	}
//	fmt.Printf("Using Certificate: %s\n", crt)
//	return crt
//}
//
//func GetTLSKey() string {
//	var key = ""
//	cert := ReleaseConf.ReleaseCertName
//	if cert == "" {
//		key = fmt.Sprintf("/etc/ssl/vsprivate/%s-selfsigned.key", cert)
//	} else {
//		key = "/etc/ssl/vsprivate/localhost-selfsigned.key"
//	}
//
//
//	fmt.Printf("Using Certificate: %s\n", cert)
//	fmt.Printf("Using CertificateKey: %s\n", key)
//	return key
//}
//
//
//func SetPort(p string) {
//	viper.Set("Port", p)
//}
//
//func SetAddress(a string) {
//	viper.Set("Address", a)
//}
