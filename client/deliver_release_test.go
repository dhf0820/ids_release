package client

//import (
//	"context"
//	"fmt"
//	"github.com/stretchr/testify/require"
//	mod "github.com/dhf0820./ds_model"
//	"os"
//	"testing"
//)
//
//func TestDeliverRelease(t *testing.T) {
//	os.Setenv("ENV_ROI", "/Users/dhf/work/roi/services/release_service/client")
//	fmt.Println("\n\n\nTestSubmitRelease")
//	cust := mod.Customer{}
//	cust.UserId = "tfrench"
//	cust.Facility = "demo"
//	cust.Name = "test"
//	cust.Code = 0
//	clnt := mod.Client{}
//	clnt.Source = "ca"
//	clnt.SourceId = "353"
//	relID := "5f397b9d3bd90a0b2f093593"
//	recipID := "5fc184b28b273e53e19b1efa"
//	deviceName := "5fc163ed8b273e53e19b1ed5"
//
//	r, err := DeliverRelease(context.Background(), relID, recipID, deviceName, &cust, &clnt)
//	require.Nil(t, err)
//	require.NotEmpty(t, r)
//
//}
