package service

//import (
//	"context"
//	"fmt"
//	domain "github.com/dhf0820/ids_release_service/internal/domain"
//	relPB "github.com/dhf0820/ids_release_service/protobufs/relPB"
//	"go.mongodb.org/mongo-driver/bson/primitive"
//	"time"
//)
//
///*AddDocument Service Accepts a AddDocumentRequest contaning teh releaseID and the new docment.
// It returns the new list of documentsafter the one added in the AddDocmentResponse
//*/
//
//func (s *ReleaseServiceServer) AddDocument(ctx context.Context,
//					req *relPB.AddDocumentRequest) (*relPB.AddDocumentResponse, error) {
//	var err error
//	pdoc := req.GetDoc()
//	relID, _ := primitive.ObjectIDFromHex(req.GetRelId())
//	//doc, err := domain.UnMarshalDocument(pdoc)
//	//fmt.Printf("server doc received: %v\n", spew.Sdump(doc))
//	doc := RelPbToDocument(pdoc)
//	startTime := time.Now()
//
//	nDoc, err := domain.AddRemoteDocumentForID(ctx, relID, doc)
//	if err != nil {
//		fmt.Printf("server/addDocument failed: %v\n", err)
//		return nil, err
//	}
//	fmt.Printf("\nTime Internal GRPC calling domain to AddDocument: %f seconds\n", time.Since(startTime).Seconds())
//	pbResp := relPB.AddDocumentResponse{
//		Document: ToPbDocument(nDoc),
//	}
//	//fmt.Printf("Server AddDocument Count: %d returning resp: %s\n", count, spew.Sdump(pbResp))
//	return  &pbResp, err
//
//}
//
