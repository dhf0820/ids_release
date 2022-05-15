package domain

//
//import (
//	"context"
//	"fmt"
//	mod "github.com/dhf0820./ds_model"
//
//	//c "github.com/dhf0820/ids_release_service/pkg/common"
//	"go.mongodb.org/mongo-driver/bson/primitive"
//)
//
////type server struct {
////}
//
//func AddRemoteDocument(rel *mod.Release, doc *mod.Document) (*mod.Document, error) {
//	//fmt.Printf("\nAdding the document to release documents: %s\n", spew.Sdump(doc))
//	rel.Documents = append(rel.Documents, doc)
//	err := Update(rel)
//	if err != nil {
//		fmt.Printf("Update Release documents returning err: %v\n", err)
//		return nil, err
//	}
//	fmt.Printf("$$$addDocument returning %d documents added\n", len(rel.Documents))
//	return doc, err
//}
//
//func AddRemoteDocumentForID(ctx context.Context,  relID primitive.ObjectID, doc *mod.Document) (*mod.Document, error){
//
//	//fmt.Printf("Retrieving ReleaseID from doument: %s\n", spew.Sdump(doc))
//
//	if relID == primitive.NilObjectID{
//		err := fmt.Errorf("Document has invalid releaseId: [%s]", relID)
//		return nil, err
//	}
//
//	fmt.Printf("\n\nGet Release: %s\n", relID)
//	rel, err := GetReleaseById(relID)
//	if err != nil {
//		err = fmt.Errorf("Release [%s] was NOT FOUND)", relID)
//		return nil, err
//	}
//
//	//fmt.Printf("Calling addDocument with doc: %v\n", spew.Sdump(doc) )
//	doc, err = AddRemoteDocument(rel, doc)
//	return doc, err
//
//}
//
