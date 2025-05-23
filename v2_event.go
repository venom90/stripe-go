//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Event reason type.
type V2EventReasonType string

// List of values that V2EventReasonType can take
const (
	V2EventReasonTypeRequest V2EventReasonType = "request"
)

// Information on the API request that instigated the event.
type V2EventReasonRequest struct {
	// ID of the API request that caused the event.
	ID string `json:"id"`
	// The idempotency key transmitted during the request.
	IdempotencyKey string `json:"idempotency_key"`
}

// Reason for the event.
type V2EventReason struct {
	// Information on the API request that instigated the event.
	Request *V2EventReasonRequest `json:"request"`
	// Event reason type.
	Type V2EventReasonType `json:"type"`
}

// Events are generated to keep you informed of activity in your business account. APIs in the /v2 namespace generate [thin events](https://docs.stripe.com/event-destinations#benefits-of-thin-events) which have small, unversioned payloads that include a reference to the ID of the object that has changed. The Events v2 API returns these new thin events. [Retrieve the event object](https://docs.stripe.com/event-destinations#fetch-data) for additional data about the event. Use the related object ID in the event payload to [fetch the API resource](https://docs.stripe.com/event-destinations#retrieve-the-object-associated-with-thin-events) of the object associated with the event. Comparatively, events generated by most API v1 include a versioned snapshot of an API object in their payload.
type V2BaseEvent struct {
	APIResource
	// Authentication context needed to fetch the event or related object.
	Context string `json:"context"`
	// Time at which the object was created.
	Created time.Time `json:"created"`
	// Unique identifier for the event.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Reason for the event.
	Reason *V2EventReason `json:"reason"`
	// The type of the event.
	Type string `json:"type"`
}

func (e *V2BaseEvent) getBaseEvent() *V2BaseEvent {
	return e
}
