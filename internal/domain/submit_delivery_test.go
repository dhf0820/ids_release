package domain

//import (
//	"fmt"
//	"github.com/stretchr/testify/require"
//	"github.com/dhf0820/ids_release_service/internal/settings"
//	"testing"
//)
//
//func TestSubmitEmailDelivery(t *testing.T) {
//	t.Parallel()
//	fmt.Printf("TestSubmitDelivery\n")
//	settings.SetDbName("test_release_service")
//	//relID, err := primitive.ObjectIDFromHex("5ee900a942396a5236640788")
//
//	sdReq := &DeliveryRequest{
//		ReleaseID:      "5ee900a942396a5236640788",
//		Recipient:      NewRecipient(),
//		Device: *AllScriptsDevice("test77777"),
//	}
//	delId, err := sdReq.Submit()
//	require.Nil(t, err)
//	require.NotNil(t, delId)
//	//require.NotEmpty(t, sdReq.ReleaseID)
//	rel, err := GetReleaseById(sdReq.ReleaseID)
//	require.NotEmpty(t, rel.Recipients)
//	//rel.History = nil
//	//err = rel.Update()
//	require.Nil(t, err)
//}
//
//
//func TestSubmitAllScriptsDelivery(t *testing.T) {
//	t.Parallel()
//	fmt.Printf("TestSubmitDelivery\n")
//	settings.SetDbName("test_release_service")
//	//relID, err := primitive.ObjectIDFromHex("5ee900a942396a5236640788")
//
//	sdReq := &DeliveryRequest{
//		ReleaseID:      "5ee510df6f4841277a5029ac",
//		Recipient:      *SampleAllScriptsRecipient(),
//		Device: *AllScriptsDevice("test77777"),
//	}
//	delId, err := sdReq.Submit()
//	require.Nil(t, err)
//	require.NotNil(t, delId)
//	//require.NotEmpty(t, sdReq.ReleaseID)
//	rel, err := GetReleaseById(sdReq.ReleaseID)
//	require.NotEmpty(t, rel.Recipients)
//	//rel.History = nil
//	//err = rel.Update()
//	require.Nil(t, err)
//}
