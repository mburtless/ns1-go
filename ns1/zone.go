package ns1

import "model/dns"

type ZonesService interface {
	ListZonesRequest(*ListZonesInput) ListZonesRequest
}

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

type ListZonesOutput struct {
	// If this value is present, there are additional results to be displayed. To
	// retrieve them, call ListTags again, with NextToken set to this value.
	//NextToken *string
	Zones []dns.Zone
}

// Send marshals and sends the ListZones API request.
func (r ListZonesRequest) Send() (*ListZonesOutput, error) {
	// Not yet implemented
	return nil, nil
}

// ListZonesRequest returns a request value for making API operation for
// NS1 Zones
func (z *ZonesServiceOp) ListZonesRequest(input *ListZonesInput) ListZonesRequest {
	// Not yet implemented
	return nil
}
