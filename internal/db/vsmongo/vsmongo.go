package vsmongo

//
//import (
//	"context"
//"fmt"
//"os"
//"strconv"
//"time"
//
//log "github.com/sirupsen/logrus"
//	"gitlab.com/dhf0820/ids_release_service/internal/settings"
//"go.mongodb.org/mongo-driver/mongo"
//"go.mongodb.org/mongo-driver/mongo/options"
//)
//
//type MongoDB struct {
//	Client       *mongo.Client
//	DatabaseName string
//	URL          string
//	Database     *mongo.Database
//	Session      mongo.Session
//	Collection   *mongo.Collection
//}
//
//var DB MongoDB
//var mongoClient *mongo.Client
////var databaseName = settings.DbName()
////var dbURL = settings.DbURL()
//
//var insertResult *mongo.InsertOneResult
//
//func Open() *MongoDB {
//	var err error
//	if GetConfig() == nil {
//
//	}
//	//settings.Init("atlas")
//	dbURL := settings.DbURL()
//
//	fmt.Printf("Open DB Type: %s\n", dbURL)
//	uri := settings.DbURL()
//	fmt.Printf("Opening database: %s\n", uri)
//	//uri := dbURL + databaseName
//	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
//
//	defer cancel()
//
//	// client, error := vsmongo.NewClient(options.Client().ApplyURI("ur_Database_uri"))
//	// error = client.Connect(ctx)
//
//	// //Checking the connection
//	// error = client.Ping(context.TODO(), nil)
//	// fmt.Println("Database connected")
//
//	opts := options.Client()
//	opts.ApplyURI(uri)
//	opts.SetMaxPoolSize(5)
//	if mongoClient, err = mongo.Connect(ctx, opts); err != nil {
//		fmt.Println(err.Error())
//	}
//	err = mongoClient.Ping(context.TODO(), nil)
//	if err != nil {
//		fmt.Printf("Database did not connect: %v\n", err)
//		log.Printf("Database did not connect: %v", err)
//		return nil
//	}
//
//	DB.Client = mongoClient
//	DB.DatabaseName = databaseName
//	DB.Database = mongoClient.Database(DB.DatabaseName)
//	DB.URL = dbURL
//
//	fmt.Println("Database connected")
//	//fmt.Printf("Client: %s\n", spew.Sdump(client))
//	DB.Collection = DB.Client.Database(DB.DatabaseName).Collection("documents")
//	return &DB
//}
//
//// ConnectToDB starts a new database connection and returns a reference to it
//func ConnectToDB(url string) (*MongoDB, error) {
//	databaseName := settings.Conf.ReleaseDatabase
//	//databaseName := os.Getenv("DATABASE")
//	fmt.Printf("Using DB: %s\n", databaseName)
//	if url == "" {
//		url = settings.DbURL()
//	}
//	fmt.Printf("Mongo URL: %s\n", url)
//	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
//	defer cancel()
//	options := options.Client().ApplyURI(url)
//	options.SetMaxPoolSize(DbPoolSize())
//	client, err := mongo.Connect(ctx, options)
//	if err != nil {
//		return nil, err
//	}
//	DB.Client = client
//	DB.DatabaseName = databaseName
//	DB.Database = client.Database(DB.DatabaseName)
//	DB.URL = dbURL
//	return &DB, nil
//}
//
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
//func Current() (*MongoDB, error) {
//	if DB.Client != nil {
//		return &DB, nil
//	}
//	_, err := ConnectToDB("")
//	//client, err := Open("")
//	return &DB, err
//}
//
//func (db *MongoDB) Close() error {
//	err := db.Client.Disconnect(context.TODO())
//	return err
//}
//
//func GetCollection(collection string) (*mongo.Collection, error) {
//	if collection == "" {
//		collection = CollectionName()
//	}
//	fmt.Printf("GetCollection: %s\n", collection)
//	db, err := Current() //"mongodb://admin:Sacj0nhat1@cat.vertisoft.com:27017")
//	if err != nil {
//		fmt.Printf("Current returned error: %s\n", err)
//		log.Fatal(err)
//		//return nil, err
//	}
//	fmt.Printf("GetCollection DBName: %s\n", DB.DatabaseName)
//	client := db.Client
//	coll := client.Database(DB.DatabaseName).Collection(collection)
//	return coll, nil
//}
//
//func CollectionName() string {
//	return "releases"
//}
