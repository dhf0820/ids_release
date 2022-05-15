package service

// import (
// 	//"fmt"
// 	"context"
// 	"fmt"
// 	"github.com/davecgh/go-spew/spew"
// 	log "github.com/sirupsen/logrus"
// 	m "github.com/dhf0820./ds_model"
// 	relPB "gitlab.com/dhf0820/ids_release_service/protobufs/relPB"
// 	//"go.mongodb.org/mongo-driver/bson/primitive"
// )

// func (s ReleaseServiceServer) GetReleaseForForeignId(ctx context.Context, req *relPB.GetReleaseRequest) (*relPB.ReleaseResponse, error) {
// 	var domRel *m.Release
// 	//var rel *relPB.GetReleaseResponse
// 	var err error
// 	fmt.Printf("\nserver getrelease req: %s\n", spew.Sdump(req))
// 	if req.ForeignId != "" {
// 		domRel, err = GetReleaseByForeignId(req.GetForeignId(), req.GetCustomerCode())
// 	}
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Marshal to ProtoBuf Release
// 	pRel := ToPbRelease(domRel)
// 	if err != nil {
// 		err = fmt.Errorf("GetReleaseForForeignId Marshal failed: %v", err)
// 		return nil, err
// 	}
// 	resp := relPB.ReleaseResponse{
// 		Release: pRel,
// 	}
// 	return &resp, nil
// }

// func (s ReleaseServiceServer) GetRelease(ctx context.Context, req *relPB.GetReleaseRequest) (*relPB.ReleaseResponse, error) {
// 	var domRel *m.Release
// 	//var rel *relPB.GetReleaseResponse
// 	var err error
// 	fmt.Printf("\nserver getrelease req: %s\n", spew.Sdump(req))
// 	if req.GetRelId() != "" {
// 		domRel, err = GetRelease(req.GetRelId())
// 	} else if req.GetForeignId() != "" {
// 		domRel, err = GetReleaseByForeignId(req.GetForeignId(), req.GetCustomerCode())
// 	} else {
// 		err = fmt.Errorf("GetRelease requires either internal releaseId or combined CustomerCode and ForeignId")
// 	}
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Marshal to ProtoBuf Release
// 	pRel := ToPbRelease(domRel)
// 	if err != nil {
// 		err = fmt.Errorf("GetRelease Marshal failed: %v", err)
// 		return nil, err
// 	}
// 	resp := relPB.ReleaseResponse{
// 		Release: pRel,
// 	}
// 	return &resp, nil
// }

// func (s *ReleaseServiceServer) CreateRelease(ctx context.Context, req *relPB.CreateReleaseRequest) (*relPB.CreateReleaseResponse, error) {
// 	var err error
// 	//fmt.Printf("\nCreateReleaseRequest: %s\n", spew.Sdump(req))
// 	pbRel := req.GetRelease()
// 	fmt.Printf("\nServer CreateRelease Before UnMarshal: %v\n", spew.Sdump(pbRel))
// 	data := ToRelease(pbRel)

// 	// Do what is necessary with the Go Release Structure

// 	data, err = CreateRelease(data)
// 	if err != nil {
// 		log.Errorf("\ndomain.CreateRelease Error: %s\n", err)
// 		return nil, err
// 	}
// 	fmt.Printf("\nServer CreateRelease returns: %v\n", spew.Sdump(data))
// 	//At this point everything is done and Marshal back to a PB Release
// 	pRel := ToPbRelease(data)

// 	//pbRel.Id = data.ID.Hex()
// 	resp := relPB.CreateReleaseResponse{Release: pRel}
// 	fmt.Printf("Server/CreateRelease resp: %s\n", spew.Sdump(resp))
// 	return &resp, nil
// }

// func (s *ReleaseServiceServer) LockRelease(ctx context.Context, req *relPB.LockReleaseRequest) (*relPB.CallStatus, error) {
// 	var err error
// 	//fmt.Printf("\nCreateReleaseRequest: %s\n", spew.Sdump(req))
// 	relId := req.GetReleaseId()
// 	mode := req.GetMode()
// 	err = LockRelease(ctx, relId, mode)
// 	resp := relPB.CallStatus{}
// 	if err != nil {
// 		resp.Ok = false
// 		resp.Comment = err.Error()
// 	} else {
// 		resp.Ok = true
// 	}
// 	return &resp, err
// }

// func (s *ReleaseServiceServer) GetDocuments(ctx context.Context, req *relPB.GetDocumentsRequest) (*relPB.DocumentsResponse, error) {
// 	var err error
// 	//fmt.Printf("\nCreateReleaseRequest: %s\n", spew.Sdump(req))
// 	relId := req.GetRelId()
// 	docs, err := GetDocuments(ctx, relId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	resp := relPB.DocumentsResponse{}
// 	resp.Documents = ToPbDocuments(docs)
// 	return &resp, err
// }
