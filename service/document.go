package service

//import (
//	devMod "github.com/dhf0820/ids_model/device"
//)

// /*AddDocument Service Accepts a AddDocumentRequest containing the releaseID and the new docment.
// It returns the new list of documents after the one added in the AddDocumentResponse
// */

// ///////////////  GRPC services //////////////////
// func (s *ReleaseServiceServer) AddDocument(ctx context.Context,
// 	req *relPB.AddDocumentRequest) (*relPB.AddDocumentResponse, error) {
// 	var err error
// 	doc := RelPbToDocument(req.GetDoc())

// 	//TODO: AddDocumentRequest the release should be a pointer or an ID
// 	//relID, _ := primitive.ObjectIDFromHex(req.GetRelId())
// 	//doc, err := domain.UnMarshalDocument(pdoc)
// 	//fmt.Printf("server doc received: %v\n", spew.Sdump(doc))
// 	//doc := RelPbToDocument(pdoc)
// 	startTime := time.Now()

// 	rel, err := AddDocument(ctx, req.GetRelId(), doc)

// 	if err != nil {
// 		fmt.Printf("server/addDocument failed: %v\n", err)
// 		return nil, err
// 	}
// 	fmt.Printf("\nTime Internal GRPC calling domain to AddDocument: %f seconds\n", time.Since(startTime).Seconds())
// 	pbResp := relPB.AddDocumentResponse{}
// 	pbResp.Release = ToPbRelease(rel)

// 	////fmt.Printf("Server AddDocument Count: %d returning resp: %s\n", count, spew.Sdump(pbResp))
// 	return &pbResp, err
// }

// //AddDocument accepts a *Release  and *Documents and adds the document to the release if it does not already exist.
// //Parameters: *Release, *Document
// //Returns number of documents, error
// func AddDocument(ctx context.Context, relId string, doc *mod.Document) (*mod.Release, error) {
// 	fmt.Printf("AddDocument: %s\n", spew.Sdump(doc))
// 	doc.Id = primitive.NewObjectID().Hex()
// 	rel, err := GetRelease(relId)
// 	if err != nil {
// 		return nil, status.Errorf(
// 			codes.NotFound,
// 			fmt.Sprintf("ReleaseID: [%s] was not found", relId),
// 		)
// 	}
// 	if documentExists(rel, doc) == true {
// 		return nil, status.Errorf(
// 			codes.AlreadyExists,
// 			fmt.Sprintf("ReleaseID: [%s] Already has document: [%s]", rel.ID.Hex(), doc.ImageURL),
// 		)
// 	}
// 	fmt.Printf("SourceImageURL: [%s]\nSourceDocumentURL: [%s]\n", doc.SourceImageURL, doc.SourceDocumentURL)
// 	if doc.SourceImageURL == "" || doc.SourceDocumentURL == "" {
// 		return nil, status.Errorf(
// 			codes.InvalidArgument,
// 			fmt.Sprintf("ReleaseID: [%s] SourceImageURL:[%s] and SourceDocumentURL: [%s] are required", rel.ID.Hex(), doc.SourceImageURL, doc.SourceDocumentURL),
// 		)
// 	}
// 	if doc.Description == "" {
// 		return nil, status.Errorf(
// 			codes.InvalidArgument,
// 			fmt.Sprintf("ReleaseID: [%s]Description is required", rel.ID.Hex()),
// 		)
// 	}
// 	//if doc.DocClass == "" {
// 	//	return nil, status.Errorf(
// 	//		codes.InvalidArgument,
// 	//		fmt.Sprintf("ReleaseID: [%s] DocClass is required", rel.ID.Hex()),
// 	//	)
// 	//}
// 	if doc.ImageURL == "" {
// 		doc.ImageURL = doc.SourceImageURL
// 	}
// 	if doc.DocumentURL == "" {
// 		doc.DocumentURL = doc.SourceDocumentURL
// 	}
// 	//fmt.Printf("\nAdding the document to release documents: %s\n", spew.Sdump(doc))
// 	rel.Documents = append(rel.Documents, doc)
// 	err = Update(rel)
// 	if err != nil {
// 		fmt.Printf("Update Release documents returning err: %v\n", err)
// 		return nil, status.Errorf(
// 			codes.Internal,
// 			fmt.Sprintf("ReleaseID: [%s] Update failed: %v", rel.ID.Hex(), err),
// 		)
// 	}
// 	fmt.Printf("addDocument returning release with %d documents\n", len(rel.Documents))
// 	rel, err = GetRelease(rel.ID.Hex())
// 	if err != nil {
// 		fmt.Printf("GetRelease: [%s] failed with %v\n", rel.ID.Hex(), err)
// 		return nil, status.Errorf(
// 			codes.Internal,
// 			fmt.Sprintf("ReleaseID: [%s] Rereading failed: %v", rel.ID.Hex(), err),
// 		)
// 	}
// 	return rel, nil
// }

// func documentExists(rel *mod.Release, doc *mod.Document) bool {
// 	for _, d := range rel.Documents {
// 		if d.SourceImageURL == doc.SourceImageURL {
// 			return true
// 		}
// 	}
// 	return false
// }
