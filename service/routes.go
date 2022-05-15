package service

import (
	//"fmt"
	//"time"
	"net/http"
	//"encoding/json"
	//"github.com/gorilla/mux"
	//log "github.com/sirupsen/logrus"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{

	Route{
		"HealthCheck",
		"GET",
		"/api/rest/v1/healthcheck",
		HealthCheck,
	},
	Route{
		"GetDocument",
		"GET",
		"/api/rest/v1/image/{release_id}",
		getReleaseImage,
	},
	// Route{
	// 	"TestForm",
	// 	"POST",
	// 	"/api/rest/v1/test",
	// 	testForm,
	// },

	//Route{
	//	"AddRecipient",
	//	"POST",
	//	"/api/rest/v1/release/{{releaseId}}/recipient",
	//	addRecipient,
	//},
	// Route{
	// 	"GetConfig",
	// 	"GET",
	// 	"/api/rest/v1/config",
	// 	getConfig,
	// },
	// Route{
	// 	"GetRecipient",
	// 	"GET",
	// 	"/api/rest/v1/recipient/{recipient_id}",
	// 	getRecipient,
	// },
	// Route{
	// 	"FindRecipient",
	// 	"GET",
	// 	"/api/rest/v1/recipient",
	// 	findRecipients,
	// },
	// Route{
	// 	"FindRecipientByForeign",
	// 	"GET",
	// 	"/api/rest/v1/recipient_foreign_id",
	// 	getRecipientByForeignId,
	// },
	// Route{
	// 	"Admin",
	// 	"PUT",
	// 	"/api/rest/v1/admin",
	// 	UpdateEnv,
	// },
	Route{
		"CreateRelease",
		"POST",
		"/api/rest/v1/release",
		createRelease,
	},
	Route{
		"GetReleaseSet",
		"GET",
		"/api/rest/v1/release/set/{release_id}",
		getReleaseSet,
	},
	//Route{
	//	"UpdateRelease",
	//	"PUT",
	//	"/api/rest/v1/release/{release_id}",
	//	updateRelease,
	//},
	Route{
		"GetRelease",
		"GET",
		"/api/rest/v1/release/{release_id}",
		getRelease,
	},
	Route{
		"GetReleaseCombined",
		"GET",
		"/api/rest/v1/release/combined/{release_id}",
		getReleaseCombined,
	},
	Route{
		"UpdateRecipients",
		"PUT",
		"/api/rest/v1/release/{release_id}/recipients",
		updateRecipients,
	},
	Route{
		"UpdateRecipients",
		"PUT",
		"/api/rest/v1/release/{release_id}/recipient",
		updateRecipient,
	},
	//Route{
	//	"UpdateDevice",
	//	"PUT",
	//	"/api/rest/v1/release/{release_id}/recipient",
	//	updateDevice,
	//},
	Route{
		"AddDocument",
		"POST",
		"/api/rest/v1/release/{release_id}/document",
		addDocument,
	},
	Route{
		"GetDocuments",
		"GET",
		"/api/rest/v1/release/{release_id}/documents}",
		getDocuments,
	},
	Route{
		"GetTestRelease",
		"GET",
		"/api/rest/v1/test_release",
		getTestRelease,
	},
	Route{
		"DeleteTestRelease",
		"DELETE",
		"/api/rest/v1/test_release/{release_id}",
		deleteRelease,
	},
	Route{
		"SubmitDelivery",
		"POST",
		"/api/rest/v1/release/deliver",
		submitRelease,
	},
	// Route{
	// 	"CreateDevice",
	// 	"POST",
	// 	"/api/rest/v1/device/{recipient_id}",
	// 	createDevice,
	// },
	// Route{
	// 	"GetDevice",
	// 	"GET",
	// 	"/api/rest/v1/device",
	// 	getDevice,
	// },
	// Route{
	// 	"GetConnectorConfig",
	// 	"GET",
	// 	"/api/rest/v1/connector_config",
	// 	getConnectorConf,
	// },

	//Route{
	//	"SetLogLevel",
	//	"GET",
	//	"/api/rest/v1/log_level/{level}",
	//	GetLogLevel,
	//},
	//Route{
	//	"SetLogLevel",
	//	"PUT",
	//	"/api/rest/v1/log_level/{level}",
	//	SetLogLevel,
	//},
}
