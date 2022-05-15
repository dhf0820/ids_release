package service

//import (
//	"context"
//	"fmt"
//	"github.com/davecgh/go-spew/spew"
//	"github.com/stretchr/testify/require"
//	db "github.com/dhf0820/ids_release_service/internal/db/vsmongo"
//	"github.com/dhf0820/ids_release_service/internal/settings"
//	"github.com/dhf0820/ids_release_service/protobufs/relPB"
//	"google.golang.org/grpc"
//	"os"
//	"testing"
//)
//
//func TestSubmitDelivery(t *testing.T) {
//	fmt.Println("\n\n\nTestServerSubmitDelivery")
//	os.Setenv("ENV_ROI", "/Users/dhf/work/roi/services/release_service/config")
//	settings.SetDbName("test_release_service")
//	db.Open()
//
//	ctx := context.Background()
//	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
//	if err != nil {
//		t.Fatalf("Failed to dial bufnet: %v", err)
//	}
//	defer conn.Close()
//	//fmt.Printf("Creating Client\n")
//	client := relPB.NewReleaseServiceClient(conn)
//	require.NotNil(t, client)
//
//	//ctx := context.Background()
//	//conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
//	//if err != nil {
//	//	t.Fatalf("Failed to dial bufnet: %v", err)
//	//}
//	//defer conn.Close()
//	////fmt.Printf("Creating Client\n")
//	//client := pb.NewReleaseServiceClient(conn)
//	//require.NotNil(t, client)
//
//	subReq := relPB.SubmitDeliveryRequest{}
//	subReq.RecipientId = "5fc184b28b273e53e19b1efa"
//	subReq.ReleaseId = "5fd2f1e8a477bc58f440dc87"
//	clnt := &relPB.Client{}
//	clnt.SourceId ="353"
//	clnt.Source = "ca"
//	subReq.Client = clnt
//	cust := &relPB.Customer{}
//	cust.Name = "test"
//	cust.UserId = "tenyedi"
//	cust.Facility = "demo"
//	cust.Code = 1
//	subReq.Customer = cust
//	subReq.DeviceId = "5fc163ed8b273e53e19b1ed5"
//
//	pbResp, err := client.SubmitDelivery(ctx, &subReq)
//	require.Nil(t,err)
//	delv := RelPbToDelivery(pbResp.GetDelivery())
//	fmt.Printf("Delivery: %s\n", spew.Sdump(delv))
//
//
//delvReq := relPB.DeliveryRequest {}
//delvReq.Delivery = ToRelPbDelivery(d)
//
//relDelResp, err := client.SubmitDelivery(ctx, &delvReq)
//respDelv := relDelResp.GetDelivery()
//fmt.Printf("Submit Delivery error: %v\n", err)
//require.Nil(t, err)
//require.NotNil(t, respDelv)
//delv := relDelResp.GetDelivery()
//relPbDelv := RelPbToDelivery(delv)
////TODO: Accept Delivery information
////delID := delResp.GetDeliveryId()
////require.NotEqual(t, "", delID)
//
//fmt.Printf("Resp: %s\n", spew.Sdump(relPbDelv))

//}
