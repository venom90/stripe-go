//
//
// File generated from our OpenAPI spec
//
//

// Package receiveddebit provides the /v1/treasury/received_debits APIs
package receiveddebit

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/treasury/received_debits APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the details of an existing ReceivedDebit by passing the unique ReceivedDebit ID from the ReceivedDebit list
func Get(id string, params *stripe.TreasuryReceivedDebitParams) (*stripe.TreasuryReceivedDebit, error) {
	return getC().Get(id, params)
}

// Retrieves the details of an existing ReceivedDebit by passing the unique ReceivedDebit ID from the ReceivedDebit list
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.TreasuryReceivedDebitParams) (*stripe.TreasuryReceivedDebit, error) {
	path := stripe.FormatURLPath("/v1/treasury/received_debits/%s", id)
	receiveddebit := &stripe.TreasuryReceivedDebit{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, receiveddebit)
	return receiveddebit, err
}

// Returns a list of ReceivedDebits.
func List(params *stripe.TreasuryReceivedDebitListParams) *Iter {
	return getC().List(params)
}

// Returns a list of ReceivedDebits.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.TreasuryReceivedDebitListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.TreasuryReceivedDebitList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/treasury/received_debits", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for treasury received debits.
type Iter struct {
	*stripe.Iter
}

// TreasuryReceivedDebit returns the treasury received debit which the iterator is currently pointing to.
func (i *Iter) TreasuryReceivedDebit() *stripe.TreasuryReceivedDebit {
	return i.Current().(*stripe.TreasuryReceivedDebit)
}

// TreasuryReceivedDebitList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TreasuryReceivedDebitList() *stripe.TreasuryReceivedDebitList {
	return i.List().(*stripe.TreasuryReceivedDebitList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
