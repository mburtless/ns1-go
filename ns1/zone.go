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
	CreateZoneRequest(*CreateZoneInput) CreateZoneRequest
	UpdateZoneRequest(*UpdateZoneInput) UpdateZoneRequest
	DeleteZoneRequest(*DeleteZoneInput) DeleteZoneRequest
}

// ZonesServiceOp handles communication with the Zones related
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

// ListZonesOutput is the output for the ListZones API operation
type ListZonesOutput struct {
	// If this value is present, there are additional results to be displayed. To
	// retrieve them, call ListTags again, with NextToken set to this value.
	//NextToken *string
	Zones []dns.Zone
}

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

// ListZonesRequest returns a request value for making an API operation for NS1 Zones
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
	Zone string
}

// GetZoneOutput is the output for the GetZone API operation
type GetZoneOutput struct {
	*dns.Zone
}

// GetZoneRequest returns a request value for making an API operation for NS1 Zones
func (z *ZonesServiceOp) GetZoneRequest(input *GetZoneInput) GetZoneRequest {
	if input == nil {
		input = &GetZoneInput{}
	}

	// define Operation
	op := &Operation{
		Name:       "",
		HTTPMethod: "GET",
		HTTPPath:   fmt.Sprintf("zones/%s", input.Zone),
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

// CreateZoneRequest is an API request type for the CreateZone API operation
type CreateZoneRequest struct {
	*Request
	Input *CreateZoneInput
}

// CreateZoneInput is the input for the CreateZone API operation
type CreateZoneInput struct {
	Zone       string `json:"zone,omitempty"`
	TTL        int    `json:"ttl,omitempty"`
	NxTTL      int    `json:"nx_ttl,omitempty"`
	Retry      int    `json:"retry,omitempty"`
	Serial     int    `json:"serial,omitempty"`
	Refresh    int    `json:"refresh,omitempty"`
	Expiry     int    `json:"expiry,omitempty"`
	Hostmaster string `json:"hostmaster,omitempty"`
}

// CreateZoneOutput is the output for the CreateZone API operation
type CreateZoneOutput struct {
	*dns.Zone
}

// CreateZoneRequest returns a request value for making an API operation for NS1 Zones
func (z *ZonesServiceOp) CreateZoneRequest(input *CreateZoneInput) CreateZoneRequest {
	if input == nil {
		input = &CreateZoneInput{}
	}

	op := &Operation{
		Name:       "",
		HTTPMethod: "PUT",
		HTTPPath:   fmt.Sprintf("zones/%s", input.Zone),
	}

	output := &CreateZoneOutput{&dns.Zone{}}
	req := z.client.NewRequest(op, input, output)

	return CreateZoneRequest{Request: req, Input: input}
}

// Send marshals and sends the CreateZone API request.
func (r CreateZoneRequest) Send() (*CreateZoneOutput, error) {
	//SetContext?  Should prob take one as param
	err := r.Request.SendToData()
	if err != nil {
		return nil, err
	}
	return r.Request.Data.(*CreateZoneOutput), nil
}

// UpdateZoneRequest is an API request type for the UpdateZone API operation
type UpdateZoneRequest struct {
	*Request
	Input *UpdateZoneInput
}

// UpdateZoneInput is the input for the UpdateZone API operation
type UpdateZoneInput struct {
	Zone       string `json:"zone,omitempty"`
	TTL        int    `json:"ttl,omitempty"`
	NxTTL      int    `json:"nx_ttl,omitempty"`
	Retry      int    `json:"retry,omitempty"`
	Serial     int    `json:"serial,omitempty"`
	Refresh    int    `json:"refresh,omitempty"`
	Expiry     int    `json:"expiry,omitempty"`
	Hostmaster string `json:"hostmaster,omitempty"`
}

// UpdateZoneOutput is the output for the UpdateZone API operation
type UpdateZoneOutput struct {
	*dns.Zone
}

// UpdateZoneRequest returns a request value for making an API operation for NS1 Zones
func (z *ZonesServiceOp) UpdateZoneRequest(input *UpdateZoneInput) UpdateZoneRequest {
	if input == nil {
		input = &UpdateZoneInput{}
	}

	op := &Operation{
		Name:       "",
		HTTPMethod: "POST",
		HTTPPath:   fmt.Sprintf("zones/%s", input.Zone),
	}

	output := &UpdateZoneOutput{&dns.Zone{}}
	req := z.client.NewRequest(op, input, output)

	return UpdateZoneRequest{Request: req, Input: input}
}

// Send marshals and sends the UpdateZone API request.
func (r UpdateZoneRequest) Send() (*UpdateZoneOutput, error) {
	//SetContext?  Should prob take one as param
	err := r.Request.SendToData()
	if err != nil {
		return nil, err
	}
	return r.Request.Data.(*UpdateZoneOutput), nil
}

// DeleteZoneRequest is an API request type for the DeleteZone API operation
type DeleteZoneRequest struct {
	*Request
	Input *DeleteZoneInput
}

// DeleteZoneInput is the input for the DeleteZone API operation
type DeleteZoneInput struct {
	Zone string `json:"zone,omitempty"`
}

// DeleteZoneOutput is the output for the DeleteZone API operation
type DeleteZoneOutput struct {
	*dns.Zone
}

// DeleteZoneRequest returns a request value for making an API operation for NS1 Zones
func (z *ZonesServiceOp) DeleteZoneRequest(input *DeleteZoneInput) DeleteZoneRequest {
	if input == nil {
		input = &DeleteZoneInput{}
	}

	op := &Operation{
		Name:       "",
		HTTPMethod: "DELETE",
		HTTPPath:   fmt.Sprintf("zones/%s", input.Zone),
	}

	output := &DeleteZoneOutput{&dns.Zone{}}
	req := z.client.NewRequest(op, input, output)

	return DeleteZoneRequest{Request: req, Input: input}
}

// Send marshals and sends the DeleteZone API request.
func (r DeleteZoneRequest) Send() (*DeleteZoneOutput, error) {
	//SetContext?  Should prob take one as param
	err := r.Request.SendToData()
	if err != nil {
		return nil, err
	}
	return r.Request.Data.(*DeleteZoneOutput), nil
}
