//
//
// File generated from our OpenAPI spec
//
//

// Package settings provides the /v1/tax/settings APIs
package settings

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
)

// Client is used to invoke /v1/tax/settings APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves Tax Settings for a merchant.
func Get(params *stripe.TaxSettingsParams) (*stripe.TaxSettings, error) {
	return getC().Get(params)
}

// Retrieves Tax Settings for a merchant.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(params *stripe.TaxSettingsParams) (*stripe.TaxSettings, error) {
	settings := &stripe.TaxSettings{}
	err := c.B.Call(http.MethodGet, "/v1/tax/settings", c.Key, params, settings)
	return settings, err
}

// Updates Tax Settings parameters used in tax calculations. All parameters are editable but none can be removed once set.
func Update(params *stripe.TaxSettingsParams) (*stripe.TaxSettings, error) {
	return getC().Update(params)
}

// Updates Tax Settings parameters used in tax calculations. All parameters are editable but none can be removed once set.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(params *stripe.TaxSettingsParams) (*stripe.TaxSettings, error) {
	settings := &stripe.TaxSettings{}
	err := c.B.Call(http.MethodPost, "/v1/tax/settings", c.Key, params, settings)
	return settings, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
