package ns1

import (
	"fmt"

	"github.com/mburtless/ns1-go/ns1/model/dns"
)

// ZonesService is an interface for interfacing with the
// Zones endpoint of the NS1 API
type ZonesService interface {
	ListZonesRequest(*ListZonesInput) ListZonesRequest
	GetZoneRequest(*GetZoneInput) GetZoneRequest
}

// ZoneServiceOp handles communication with the Zones related
// methods of the NS1 API
type ZonesServiceOp struct {
	client *Client
}

// ListZonesRequest is a API request type for the ListZones API operation
type ListZonesRequest struct {
	*Request
	Input *ListZonesInput
}

// ListZonesInput is the input for the ListZones API operation
type ListZonesInput struct {
	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response
	// includes only results beyond the token.
	//NextToken *string `type:"string"`
}

// ListZonesOutput
//type ListZonesOutput []struct {
type ListZonesOutput struct {
	// If this value is present, there are additional results to be displayed. To
	// retrieve them, call ListTags again, with NextToken set to this value.
	//NextToken *string
	Zones []dns.Zone
}

//type ListZonesOutput []*dns.Zone

// Send marshals and sends the ListZones API request.
func (r ListZonesRequest) Send() (*ListZonesOutput, error) {
	//SetContext?  Should prob take one as param
	err := r.Request.SendToPointer(&r.Request.Data.(*ListZonesOutput).Zones)
	if err != nil {
		return nil, err
	}
	//return r.Request.Data(*ListZonesOutput)
	// Not yet implemented
	//return nil, nil
	return r.Request.Data.(*ListZonesOutput), nil
}

// ListZonesRequest returns a request value for making API operation for
// NS1 Zones
func (z *ZonesServiceOp) ListZonesRequest(input *ListZonesInput) ListZonesRequest {
	// define Operation
	op := &Operation{
		Name:       "",
		HTTPMethod: "GET",
		HTTPPath:   "zones",
	}

	if input == nil {
		input = &ListZonesInput{}
	}

	// Init empty output struct to hold response later
	output := &ListZonesOutput{}
	// create request obj via call to NewRequest(operation, input, output)
	req := z.client.NewRequest(op, input, output)
	//return ListZonesRequest struct containing request obj and input
	return ListZonesRequest{Request: req, Input: input}
}

// GetZoneRequest is a API request type for the GetZone API operation
type GetZoneRequest struct {
	*Request
	Input *GetZoneInput
}

// GetZoneInput is the input for the GetZone API operation
type GetZoneInput struct {
	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response
	// includes only results beyond the token.
	//NextToken *string `type:"string"`
	ZoneName *string
}

// GetZoneOutput is the output for the GetZone API operation
type GetZoneOutput struct {
	// If this value is present, there are additional results to be displayed. To
	// retrieve them, call ListTags again, with NextToken set to this value.
	//NextToken *string
	*dns.Zone
}

// GetZoneRequest returns a request value for making API operation for
// NS1 Zones
func (z *ZonesServiceOp) GetZoneRequest(input *GetZoneInput) GetZoneRequest {
	if input == nil {
		input = &GetZoneInput{}
	}

	// define Operation
	op := &Operation{
		Name:       "",
		HTTPMethod: "GET",
		HTTPPath:   fmt.Sprintf("zones/%s", *input.ZoneName),
	}

	// Init empty output struct to hold response later
	output := &GetZoneOutput{}
	// create request obj via call to NewRequest(operation, input, output)
	req := z.client.NewRequest(op, input, output)
	//return GetZoneRequest struct containing request obj and input
	return GetZoneRequest{Request: req, Input: input}
}

// Send marshals and sends the GetZone API request.
func (r GetZoneRequest) Send() (*GetZoneOutput, error) {
	//SetContext?  Should prob take one as param
	err := r.Request.SendToData()
	if err != nil {
		return nil, err
	}
	return r.Request.Data.(*GetZoneOutput), nil
}
