package client

//import (
//	"fmt"
//	"github.com/davecgh/go-spew/spew"
//	relPB "gitlab.com/dhf0820/ids_release_service/protobufs/relPB"
//)
//
//func ToRelease(c *relPB.Release) *Release {
//	r := &Release{}
//	r.ID = c.Id
//	r.Client = c.GetClient()
//	r.Facility = c.GetFacility()
//	r.SourceID = c.GetSourceId()
//	r.UserID = c.GetUserName()
//	r.Patient.Name = c.GetPatName()
//	r.Patient.MRN = c.GetPatMrn()
//	r.Patient.DOB = c.GetPatDob()
//	r.Patient.SSN = c.GetPatSsn()
//	r.Documents = ToDocuments(c.GetDocuments())
//	fmt.Printf("ToRelease: %s\n", spew.Sdump(r))
//	return r
//}
//
//func ToPbRelease(r *Release) *relPB.Release {
//	c := &relPB.Release{}
//	c.Id = r.ID
//	c.Client = r.Client
//	c.Facility = r.Facility
//	c.SourceId = r.SourceID
//	c.UserName = r.UserID
//	c.PatName = r.Patient.Name
//	c.PatMrn = r.Patient.MRN
//	c.PatDob = r.Patient.DOB
//	c.PatSsn = r.Patient.SSN
//	//c.Documents = ToPBDocuments(r.Documents)
//	fmt.Printf("\n $$$$  ToPbRelease ssn = %s  source SSN %s\n", c.GetPatSsn(), r.Patient.SSN)
//	return c
//}
//
////UnMarshal converts from a PB.Release to Gostruct Release
////This allows all client to use normal GO Structs and the system converts to proper ProtoBuf format to send over the wire
//// func UnMarshal(r *relPB.Release) *Release {
//// 	rel := Release {}
//// 	rel.ID  		= r.GetId()
//// 	rel.SourceID	= r.GetSourceId()
//// 	rel.Facility	= r.GetFacility()
//// 	rel.Patient.Name = r.GetPatName()
//// 	rel.Patient.MRN  = r.GetPatMrn()
//// 	rel.Patient.DOB  = r.GetPatDob()
//// 	rel.Patient.SSN  = r.GetPatSsn()
//// 	rel.UserID 		 = r.GetUserName()
//// 	rel.Client 		 = r.GetClient()
//
//// 	return &rel
//// }
//
////Marshal converts from Go Struct version of a Release to the protobuf format to send over the wire
//// func Marshal(r * Release) *relPB.Release {
//
//// 	pbRel := relPB.Release{}
//// 	pbRel.Client		= r.Client
//// 	pbRel.UserName 		= r.UserID
//// 	pbRel.Facility		= r.Facility
//// 	pbRel.SourceId		= r.SourceID
//// 	pbRel.Id			= r.ID
//// 	pbRel.PatSsn		= r.Patient.SSN
//// 	pbRel.PatDob		= r.Patient.DOB
//// 	pbRel.PatMrn		= r.Patient.MRN
//// 	pbRel.PatName		= r.Patient.Name
//
//// 	return &pbRel
//// }
//
//func ToDocuments(pds []*relPB.Document) []Document {
//	ds := []Document {}
//	for _, pd := range pds {
//		d := ToDocument(pd)
//		ds = append(ds, *d)
//	}
//	return ds
//}
//
//func ToDocument(d *relPB.Document) *Document {
//	doc := Document{}
//	doc.MRN = d.GetMrn()
//	doc.ID = d.GetId()
//	doc.SourceID = d.GetSourceId()
//	doc.Source = d.GetSource()
//	doc.Facility = d.GetFacility()
//	doc.Description = d.GetDescription()
//	doc.DocClass = d.GetDocClass()
//	doc.ImageType = d.GetImageType()
//	doc.ClientId = d.GetClientId()
//	doc.DateOfService = d.GetDateOfService()
//	return &doc
//}
//
//func ToPBDocuments(ds []Document) []*relPB.Document {
//	pds := []*relPB.Document {}
//	for _, rd := range ds {
//		pd := ToPbDocument(&rd)
//		pds = append(pds, pd)
//	}
//	return pds
//}
//func ToPbDocument(ds *Document) *relPB.Document {
//	pd := relPB.Document{}
//	pd.DateOfService = ds.DateOfService
//	pd.ClientId = ds.ClientId
//	pd.ImageType = ds.ImageType
//	pd.DocClass = ds.DocClass
//	pd.Description = ds.Description
//	pd.Facility = ds.Facility
//	pd.Source = ds.Source
//	pd.SourceId = ds.SourceID
//	pd.ReleaseId = ds.ReleaseId
//	pd.Mrn = ds.MRN
//	pd.Id = ds.ID
//	return &pd
//}
//
////UnMarshalDocumentsummary converts from PB format of Document to the internal GO Struct version for clients
//// func UnMarshalDocument(d *relPB.Document) *Document {
//// 	doc :=Document{}
//// 	doc.MRN			= d.GetMrn()
//// 	doc.ID 			= d.GetId()
//// 	doc.SourceID	= d.GetSourceId()
//// 	doc.Source		= d.GetSource()
//// 	doc.Facility	= d.GetFacility()
//// 	doc.Description = d.GetDescription()
//// 	doc.DocClass	= d.GetDocClass()
//// 	doc.ImageType	= d.GetImageType()
//// 	doc.ClientId    = d.GetClientId()
//// 	doc.DateOfService = d.GetDateOfService()
//
//// 	return &doc
//// }
//
//// //MarshalDocument conversts from the client's Go Struct version of a Document to Protobuf for over the wire.
//// func MarshalDocument(ds *Document) *relPB.Document {
//// 	pd := relPB.Document{}
//// 	pd.DateOfService 	= ds.DateOfService
//// 	pd.ClientId			= ds.ClientId
//// 	pd.ImageType		= ds.ImageType
//// 	pd.DocClass			= ds.DocClass
//// 	pd.Description		= ds.Description
//// 	pd.Facility			= ds.Facility
//// 	pd.Source			= ds.Source
//// 	pd.SourceId			= ds.SourceID
//// 	pd.ReleaseId        = ds.ReleaseId
//// 	pd.Mrn				= ds.MRN
//// 	pd.Id				= ds.ID
//// 	return &pd
//// }
