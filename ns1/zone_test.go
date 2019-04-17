package ns1

import (
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

	expectedReq, _ := http.NewRequest("GET", fmt.Sprintf("%szones", client.config.endpoint), nil)
	expectedReq.Header.Add(headerAuth, client.config.apiKey)
	expectedReq.Header.Add("User-Agent", client.config.userAgent)

	req := client.Zones.ListZonesRequest(&ListZonesInput{})
	assert.Equal(t, expectedReq, req.HTTPRequest)
	assert.Equal(t, req.Config, client.config)
	assert.Equal(t, req.HTTPClient, client.httpClient)
}

func TestListZonesRequest_Send(t *testing.T) {
	setup()
	defer teardown()

	req := client.Zones.ListZonesRequest(&ListZonesInput{})
	expected := &ListZonesOutput{
		Zones: []dns.Zone{
			{Zone: "foo.com", ID: "12345678910111213141516a"},
			{Zone: "bar.com", ID: "12345678910111213141516b"},
		},
	}
	//*expected = append(*expected, struct{ *dns.Zone }{&dns.Zone{Zone: "foo.com", ID: "12345678910111213141516a"}})
	//*expected = append(*expected, struct{ *dns.Zone }{&dns.Zone{Zone: "bar.com", ID: "12345678910111213141516b"}})
	//&struct{Zone: "bar.com", ID: "12345678910111213141516b"},

	mux.HandleFunc("/v1/zones", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{"zone": "foo.com", "id": "12345678910111213141516a"},
		{"zone": "bar.com", "id": "12345678910111213141516b"}]`)
	})
	out, err := req.Send()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	assert.Equal(t, expected, out)
	t.Logf("Out: %v\nExpected:: %v", out, expected)
}
