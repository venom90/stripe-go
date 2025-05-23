//
//
// File generated from our OpenAPI spec
//
//

// Package outboundtransfer provides the /v1/treasury/outbound_transfers APIs
package outboundtransfer

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
)

// Client is used to invoke /v1/treasury/outbound_transfers APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Updates a test mode created OutboundTransfer with tracking details. The OutboundTransfer must not be cancelable, and cannot be in the canceled or failed states.
func Update(id string, params *stripe.TestHelpersTreasuryOutboundTransferParams) (*stripe.TreasuryOutboundTransfer, error) {
	return getC().Update(id, params)
}

// Updates a test mode created OutboundTransfer with tracking details. The OutboundTransfer must not be cancelable, and cannot be in the canceled or failed states.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.TestHelpersTreasuryOutboundTransferParams) (*stripe.TreasuryOutboundTransfer, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/treasury/outbound_transfers/%s", id)
	outboundtransfer := &stripe.TreasuryOutboundTransfer{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, outboundtransfer)
	return outboundtransfer, err
}

// Transitions a test mode created OutboundTransfer to the failed status. The OutboundTransfer must already be in the processing state.
func Fail(id string, params *stripe.TestHelpersTreasuryOutboundTransferFailParams) (*stripe.TreasuryOutboundTransfer, error) {
	return getC().Fail(id, params)
}

// Transitions a test mode created OutboundTransfer to the failed status. The OutboundTransfer must already be in the processing state.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Fail(id string, params *stripe.TestHelpersTreasuryOutboundTransferFailParams) (*stripe.TreasuryOutboundTransfer, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/treasury/outbound_transfers/%s/fail", id)
	outboundtransfer := &stripe.TreasuryOutboundTransfer{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, outboundtransfer)
	return outboundtransfer, err
}

// Transitions a test mode created OutboundTransfer to the posted status. The OutboundTransfer must already be in the processing state.
func Post(id string, params *stripe.TestHelpersTreasuryOutboundTransferPostParams) (*stripe.TreasuryOutboundTransfer, error) {
	return getC().Post(id, params)
}

// Transitions a test mode created OutboundTransfer to the posted status. The OutboundTransfer must already be in the processing state.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Post(id string, params *stripe.TestHelpersTreasuryOutboundTransferPostParams) (*stripe.TreasuryOutboundTransfer, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/treasury/outbound_transfers/%s/post", id)
	outboundtransfer := &stripe.TreasuryOutboundTransfer{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, outboundtransfer)
	return outboundtransfer, err
}

// Transitions a test mode created OutboundTransfer to the returned status. The OutboundTransfer must already be in the processing state.
func ReturnOutboundTransfer(id string, params *stripe.TestHelpersTreasuryOutboundTransferReturnOutboundTransferParams) (*stripe.TreasuryOutboundTransfer, error) {
	return getC().ReturnOutboundTransfer(id, params)
}

// Transitions a test mode created OutboundTransfer to the returned status. The OutboundTransfer must already be in the processing state.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ReturnOutboundTransfer(id string, params *stripe.TestHelpersTreasuryOutboundTransferReturnOutboundTransferParams) (*stripe.TreasuryOutboundTransfer, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/treasury/outbound_transfers/%s/return", id)
	outboundtransfer := &stripe.TreasuryOutboundTransfer{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, outboundtransfer)
	return outboundtransfer, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
