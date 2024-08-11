package mocktsdbclient

import "github.com/bluebreezecf/opentsdb-goclient/client"

type MockTsdbClient struct {
	dataPoints [][]client.DataPoint
}

func New() *MockTsdbClient {
	return &MockTsdbClient{}
}

func (c *MockTsdbClient) GetData() [][]client.DataPoint {
	return c.dataPoints
}

func (c *MockTsdbClient) Ping() error {
	return nil
}

func (c *MockTsdbClient) Put(datas []client.DataPoint, queryParam string) (*client.PutResponse, error) {
	c.dataPoints = append(c.dataPoints, datas)
	return nil, nil
}

func (c *MockTsdbClient) Query(param client.QueryParam) (*client.QueryResponse, error) {
	return nil, nil
}

func (c *MockTsdbClient) QueryLast(param client.QueryLastParam) (*client.QueryLastResponse, error) {
	return nil, nil
}

func (c *MockTsdbClient) Aggregators() (*client.AggregatorsResponse, error) {
	return nil, nil
}

func (c *MockTsdbClient) Config() (*client.ConfigResponse, error) {
	return nil, nil
}

func (c *MockTsdbClient) Serializers() (*client.SerialResponse, error) {
	return nil, nil
}

func (c *MockTsdbClient) Stats() (*client.StatsResponse, error) {
	return nil, nil
}

func (c *MockTsdbClient) Suggest(sugParm client.SuggestParam) (*client.SuggestResponse, error) {
	return nil, nil
}

func (c *MockTsdbClient) Version() (*client.VersionResponse, error) {
	return nil, nil
}

func (c *MockTsdbClient) Dropcaches() (*client.DropcachesResponse, error) {
	return nil, nil
}

func (c *MockTsdbClient) QueryAnnotation(queryAnnoParam map[string]interface{}) (*client.AnnotationResponse, error) {
	return nil, nil
}

func (c *MockTsdbClient) UpdateAnnotation(annotation client.Annotation) (*client.AnnotationResponse, error) {
	return nil, nil
}

func (c *MockTsdbClient) DeleteAnnotation(annotation client.Annotation) (*client.AnnotationResponse, error) {
	return nil, nil
}

func (c *MockTsdbClient) BulkUpdateAnnotations(annotations []client.Annotation) (*client.BulkAnnotatResponse, error) {
	return nil, nil
}

func (c *MockTsdbClient) BulkDeleteAnnotations(bulkDelParam client.BulkAnnoDeleteInfo) (*client.BulkAnnotatResponse, error) {
	return nil, nil
}

func (c *MockTsdbClient) QueryUIDMetaData(metaQueryParam map[string]string) (*client.UIDMetaDataResponse, error) {
	return nil, nil
}

func (c *MockTsdbClient) UpdateUIDMetaData(uidMetaData client.UIDMetaData) (*client.UIDMetaDataResponse, error) {
	return nil, nil
}

func (c *MockTsdbClient) DeleteUIDMetaData(uidMetaData client.UIDMetaData) (*client.UIDMetaDataResponse, error) {
	return nil, nil
}

func (c *MockTsdbClient) AssignUID(assignParam client.UIDAssignParam) (*client.UIDAssignResponse, error) {
	return nil, nil
}

func (c *MockTsdbClient) QueryTSMetaData(tsuid string) (*client.TSMetaDataResponse, error) {
	return nil, nil
}

func (c *MockTsdbClient) UpdateTSMetaData(tsMetaData client.TSMetaData) (*client.TSMetaDataResponse, error) {
	return nil, nil
}

func (c *MockTsdbClient) DeleteTSMetaData(tsMetaData client.TSMetaData) (*client.TSMetaDataResponse, error) {
	return nil, nil
}
