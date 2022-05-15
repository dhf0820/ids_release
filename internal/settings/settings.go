package settings

//import (
//	"fmt"
//	"os"
//	"strconv"
//
//	"github.com/spf13/viper"
//
//	log "github.com/sirupsen/logrus"
//
//	//"github.com/prometheus/common/log"
//	"go.mongodb.org/mongo-driver/mongo"
//)
//
//var atlasURL = "mongodb//dhf:Sacj0nhat1@cluster0-rcwui.mongodb.net/?retryWrites=true&w=majority"
//var localURL = "mongodb://dhf:Sacj0nhat1@docker1.vertisoft.com:27017/?serverSelectionTimeoutMS=5000&connectTimeoutMS=10000&authSource=admin&authMechanism=SCRAM-SHA-256"
//
//var currentDB *mongo.Database
//var currentCollection = "releases"
//var urlToUse = "local"
//var dbName = ""
//var baseURL = localURL
//var dbEnv = ""
//var mongoURL string
//var exists bool
//
//type config struct {
//	DeliveryPort        string
//	DeliveryAddress     string
//	DeliveryTLSMode     string
//	DeliveryCertName    string
//	DeliveryDeployMode  string
//	DeliveryDatabase    string
//
//	ReleasePort        string
//	ReleaseAddress     string
//	ReleaseTLSMode     string
//	ReleaseCertName    string
//	ReleaseDeployMode  string
//	ReleaseDatabase    string
//
//	DocumentPort       string
//	DocumentAddress    string
//	DocumentTLSMode    string
//	DocumentCertName   string
//	DocumentDeployMode string
//	DocumentDatabase   string
//
//	CorePort           string
//	CoreAddress        string
//	CoreTLSMode        string
//	CoreCertName       string
//	CoreDeployType     string
//
//	Atlas              string
//	MongoDB            string
//	Dbenv              string
//	Testing            string
//}
//
//// Create a new config instance.
//var (
//	Conf *config
//)
//
//// Init creates teh configuration information using Viper
//func Initialize(dbEnv string) {
//	Conf = getConf()
//	if dbEnv == "" {
//		//dbEnv = os.Getenv("DBENV")
//		dbEnv = Conf.Dbenv
//	}
//	if dbEnv == "" {
//		fmt.Printf("No DBENV environment set using local\n")
//		dbEnv = "local"
//	}
//}
//
//var (
//	V *viper.Viper
//)
//
//// Read the config file from the current directory and marshal
//// into the conf config struct.
//func getConf() *config {
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
//	conf := &config{}
//	err = V.Unmarshal(conf)
//	if err != nil {
//		fmt.Printf("unable to decode into config struct, %v", err)
//	}
//	//fmt.Printf("ReleaseConf: %s\n", spew.Sdump(conf))
//	test := V.Get("testing")
//	fmt.Printf("Test: %s\n", test)
//	return conf
//}
//
////DbURL returns the setting for the Mongo url or the default: docker1.vertisoft.com:27017/releases
//func DbURL() string {
//
//	if baseURL == "" {
//		return baseURL
//	}
//	//dbEnvToUse := Conf.Dbenv
//	////dbEnvToUse := os.Getenv("DBENV") // local, atlas, DBURL environment name
//	//if dbEnvToUse == "" {
//	//	fmt.Printf("No DB environment set using local\n")
//	//	dbEnvToUse = "local"
//	//}
//
//	switch dbEnv {
//	case "local":
//		fmt.Printf("Using local\n")
//		baseURL = localURL
//	case "atlas":
//		fmt.Printf("Using atlas\n")
//		baseURL = atlasURL
//	case "mongodb":
//		fmt.Printf("Looking up ENV MONGODB\n")
//		baseURL = Conf.MongoDB
///*
//		baseURL, exists = os.LookupEnv("MONGODB")
//		if exists != true {
//			fmt.Printf("MONGODB is not found\n")
//			baseURL = localURL
//			dbEnvToUse = "local"
//		}*/
//	//default:
//	//	baseURL, exists = os.LookupEnv(dbEnvToUse)
//	//	if !exists { // We have no MONGODB URL so use default local one
//	//		fmt.Printf("No mongourl provided using local server\n")
//	//		baseURL = localURL
//	//		dbEnvToUse = "local"
//	//	}
//	}
//
//	mongoURL = baseURL
//	return mongoURL
//}
//
////DbName returns the setting for the name of the db or the default: roi_dev
//func DbName() string {
//	var exists bool
//	//if dbName == "" {
//	log.Debugf("\n\nGetting MongoDB Database Name from Environment")
//	dbName, exists = os.LookupEnv("DATABASE")
//	if exists == false {
//		dbName = "test_release_service"
//		log.Warnf("Using default Database: [%s]", dbName)
//		//dbName = defaultName
//	}
//	fmt.Printf("DbName returning: %s\n", dbName)
//	//}
//	return dbName
//}
//
//func SetDbName(name string) {
//	log.Warnf("Setting Database Name: %s", name)
//	os.Setenv("DATABASE", name)
//	dbName = name
//
//}
//
//// DbPoolSize returns the setting for PoolSize or the default of 100
//func DbPoolSize() uint64 {
//	var poolSize uint64
//	poolSizeString, exists := os.LookupEnv("POOLSIZE")
//	if exists == false {
//		poolSizeString = "100"
//	}
//	poolSizeInt64, err := strconv.ParseInt(poolSizeString, 10, 64)
//	if err == nil {
//		poolSize = uint64(poolSizeInt64)
//	} else {
//		poolSize = 100
//	}
//	return poolSize
//}
//
//func SetCollectionName(name string) {
//	currentCollection = name
//}
//
////CollectionName is the name of the mongo collection used by this microservice.
//func CollectionName() string {
//	return currentCollection
//}
//
//func SetCurrentDB(conn *mongo.Database) {
//	currentDB = conn
//}
//
//func CurrentDB() *mongo.Database {
//	return currentDB
//}
//
////TODO: put certificates in a secure K/V store as actual key and Certificate
//func GetTLSCert() string {
//	var crt = ""
//	cert, exists := os.LookupEnv("CERTNAME")
//	if exists {
//		crt = fmt.Sprintf("/etc/ssl/vscerts/%s-selfsigned.crt", cert)
//	} else {
//		crt = "/etc/ssl/vscerts/localhost-selfsigned.crt"
//	}
//	//crt, exists := os.LookupEnv("CERT.CRT")
//	//if !exists {
//	//	crt = "/etc/ssl/vscerts/localhost-selfsigned.crt"
//	//}
//	fmt.Printf("Using Certificate: %s\n", crt)
//	return crt
//}
//
//func GetTLSKey() string {
//	var key = ""
//	cert, exists := os.LookupEnv("CERTNAME")
//	if exists {
//		key = fmt.Sprintf("/etc/ssl/vsprivate/%s-selfsigned.key", cert)
//	} else {
//		key = "/etc/ssl/vsprivate/localhost-selfsigned.key"
//	}
//
//	//key, exists := os.LookupEnv("CERT.KEY")
//	//if !exists {
//	//	key = "/etc/ssl/vsprivate/localhost-selfsigned.key"
//	//}
//	fmt.Printf("Using Certificate: %s\n", cert)
//	fmt.Printf("Using CertificateKey: %s\n", key)
//	return key
//
//}
