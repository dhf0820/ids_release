package domain

//import (
//	"fmt"
//	"github.com/stretchr/testify/require"
//	"github.com/dhf0820/ids_release_service/internal/settings"
//	//"go.mongodb.org/mongo-driver/bson"
//	"testing"
//	"time"
//)
//
//func TestUpdateDocumentCount(t *testing.T) {
//	t.Parallel()
//	fmt.Printf("TestUpdateRelease\n")
//	settings.SetDbName("test_release_service")
//	rel, err := CreateRelease(NewRelease("EMAIL"))
//	fmt.Printf("AddedID: %s\n", rel.ID.Hex())
//	require.Nil(t, err)
//	require.NotNil(t, rel)
//
//	// Update the number of documents
//	val, err := UpdateDocumentCount(rel.ID, 1)
//	time.Sleep(5 * time.Second)
//	val, err = UpdateDocumentCount(rel.ID, 10)
//
//	fmt.Printf("PatientName: %s\n", val.Patient)  //val["patient"])
//	require.Nil(t, err)
//
//
//	//require.Equal(t, 1, val)
//
//	//DeleteDocument(doc.ID)
//}
