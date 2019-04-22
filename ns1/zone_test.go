package ns1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/mburtless/ns1-go/ns1/model/dns"
	"github.com/stretchr/testify/assert"
)

/*
type MockZonesServiceOp struct {
	mock.Mock
}

func (m *MockZonesServiceOp) ListZonesRequest(input *ListZonesInput) ListZonesRequest {
	return ListZonesRequest{}
}

func (m *MockZonesServiceOp) GetZoneRequest(input *GetZoneInput) GetZoneRequest {
	return GetZoneRequest{}
}
*/
func TestListZonesRequest(t *testing.T) {
	setup()
	defer teardown()

	expectedReq, _ := http.NewRequest("GET", fmt.Sprintf("%szones", client.config.endpoint), new(bytes.Buffer))
	addRequestHeaders(expectedReq)

	req := client.Zones.ListZonesRequest(&ListZonesInput{})

	// GetBody is a func, can't compare equality on this so must unset
	expectedReq.GetBody, req.HTTPRequest.GetBody = nil, nil
	assert.Equal(t, expectedReq, req.HTTPRequest)
	assert.Equal(t, req.Config, client.config)
	assert.Equal(t, req.HTTPClient, client.httpClient)
}

func TestListZonesRequest_Send(t *testing.T) {
	setup()
	defer teardown()

	req := client.Zones.ListZonesRequest(&ListZonesInput{})
	//req := client.Zones.GetZoneRequest(&GetZoneInput{"foo.com"})
	expectedList := &ListZonesOutput{
		Zones: []dns.Zone{
			{Zone: "foo.com", ID: "12345678910111213141516a"},
			{Zone: "bar.com", ID: "12345678910111213141516b"},
		},
	}

	mux.HandleFunc("/v1/zones", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{"zone": "foo.com", "id": "12345678910111213141516a"},
		{"zone": "bar.com", "id": "12345678910111213141516b"}]`)
	})
	actualList, err := req.Send()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	assert.Equal(t, expectedList, actualList)
	//t.Logf("Out: %v\nExpected:: %v", out, expected)
}

func TestGetZoneRequest(t *testing.T) {
	setup()
	defer teardown()

	expectedReq, _ := http.NewRequest("GET", fmt.Sprintf("%szones/foo.com", client.config.endpoint), new(bytes.Buffer))
	expectedReq.Header.Add(headerAuth, client.config.apiKey)
	expectedReq.Header.Add("User-Agent", client.config.userAgent)

	req := client.Zones.GetZoneRequest(&GetZoneInput{Zone: "foo.com"})
	// GetBody is a func, can't compare equality on this so must unset
	expectedReq.GetBody, req.HTTPRequest.GetBody = nil, nil
	assert.Equal(t, expectedReq, req.HTTPRequest)
	assert.Equal(t, req.Config, client.config)
	assert.Equal(t, req.HTTPClient, client.httpClient)
}

func TestGetZoneRequest_Send(t *testing.T) {
	setup()
	defer teardown()

	req := client.Zones.GetZoneRequest(&GetZoneInput{Zone: "foo.com"})
	expectedZone := &GetZoneOutput{
		&dns.Zone{
			Zone: "foo.com",
			ID:   "12345678910111213141516a",
		},
	}

	mux.HandleFunc("/v1/zones/foo.com", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"zone": "foo.com", "id": "12345678910111213141516a"}`)
	})
	actualZone, err := req.Send()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	assert.Equal(t, expectedZone, actualZone)
	//t.Logf("Out: %+v\nExpected: %+v", actualZone.ID, expectedZone.ID)
}

func TestCreateZoneRequest(t *testing.T) {
	setup()
	defer teardown()

	input := &CreateZoneInput{
		Zone:    "foo.com",
		TTL:     4000,
		Refresh: 40000,
	}

	expectedBody := new(bytes.Buffer)
	_ = json.NewEncoder(expectedBody).Encode(input)
	expectedReq, _ := http.NewRequest("PUT", fmt.Sprintf("%szones/foo.com", client.config.endpoint), expectedBody)
	addRequestHeaders(expectedReq)

	req := client.Zones.CreateZoneRequest(input)
	// GetBody is a func, can't compare equality on this so must unset
	expectedReq.GetBody, req.HTTPRequest.GetBody = nil, nil
	assert.Equal(t, expectedReq, req.HTTPRequest)
	assert.Equal(t, req.Config, client.config)
	assert.Equal(t, req.HTTPClient, client.httpClient)
}

func TestCreateZoneRequest_Send(t *testing.T) {
	setup()
	defer teardown()

	req := client.Zones.CreateZoneRequest(&CreateZoneInput{Zone: "foo.com"})
	expectedZone := &CreateZoneOutput{
		&dns.Zone{
			Zone: "foo.com",
			ID:   "12345678910111213141516a",
		},
	}

	mux.HandleFunc("/v1/zones/foo.com", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprint(w, `{"zone": "foo.com", "id": "12345678910111213141516a"}`)
	})
	actualZone, err := req.Send()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	assert.Equal(t, expectedZone, actualZone)
}

func TestUpdateZoneRequest(t *testing.T) {
	setup()
	defer teardown()

	input := &UpdateZoneInput{
		Zone:    "foo.com",
		Refresh: 40000,
	}

	expectedBody := new(bytes.Buffer)
	_ = json.NewEncoder(expectedBody).Encode(input)
	expectedReq, _ := http.NewRequest("POST", fmt.Sprintf("%szones/foo.com", client.config.endpoint), expectedBody)
	addRequestHeaders(expectedReq)

	req := client.Zones.UpdateZoneRequest(input)
	// GetBody is a func, can't compare equality on this so must unset
	expectedReq.GetBody, req.HTTPRequest.GetBody = nil, nil
	assert.Equal(t, expectedReq, req.HTTPRequest)
	assert.Equal(t, req.Config, client.config)
	assert.Equal(t, req.HTTPClient, client.httpClient)
}

func TestDeleteZoneRequest(t *testing.T) {
	setup()
	defer teardown()

	input := &DeleteZoneInput{
		Zone: "foo.com",
	}

	expectedBody := new(bytes.Buffer)
	expectedReq, _ := http.NewRequest("DELETE", fmt.Sprintf("%szones/foo.com", client.config.endpoint), expectedBody)
	addRequestHeaders(expectedReq)

	req := client.Zones.DeleteZoneRequest(input)
	// GetBody is a func, can't compare equality on this so must unset
	expectedReq.GetBody, req.HTTPRequest.GetBody = nil, nil
	assert.Equal(t, expectedReq, req.HTTPRequest)
	assert.Equal(t, req.Config, client.config)
	assert.Equal(t, req.HTTPClient, client.httpClient)
}
