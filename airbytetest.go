// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package airbytetest

import (
	"airbyte-test/pkg/utils"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	"http://localhost:8000/api",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

// AirbyteTest - Airbyte Configuration API
// [https://airbyte.io](https://airbyte.io).
//
// This API is a collection of HTTP RPC-style methods. While it is not a REST API, those familiar with REST should find the conventions of this API recognizable.
//
// Here are some conventions that this API follows:
// * All endpoints are http POST methods.
// * All endpoints accept data via `application/json` request bodies. The API does not accept any data via query params.
// * The naming convention for endpoints is: localhost:8000/{VERSION}/{METHOD_FAMILY}/{METHOD_NAME} e.g. `localhost:8000/v1/connections/create`.
// * For all `update` methods, the whole object must be passed in, even the fields that did not change.
//
// Authentication (OSS):
// * When authenticating to the Configuration API, you must use Basic Authentication by setting the Authentication Header to Basic and base64 encoding the username and password (which are `airbyte` and `password` by default - so base64 encoding `airbyte:password` results in `YWlyYnl0ZTpwYXNzd29yZA==`). So the full header reads `'Authorization': "Basic YWlyYnl0ZTpwYXNzd29yZA=="`
//
// https://airbyte.io - Find out more about Airbyte
type AirbyteTest struct {
	// Attempt - Interactions with attempt related resources.
	Attempt *attempt
	// Connection - Connection between sources and destinations.
	Connection *connection
	// Destination - Destination related resources.
	Destination *destination
	// DestinationDefinition - DestinationDefinition related resources.
	DestinationDefinition *destinationDefinition
	// DestinationDefinitionSpecification - DestinationDefinitionSpecification related resources.
	DestinationDefinitionSpecification *destinationDefinitionSpecification
	// DestinationOauth - Source OAuth related resources to delegate access from user.
	DestinationOauth *destinationOauth
	// Health - Healthchecks
	Health        *health
	Internal      *internal
	Jobs          *jobs
	Logs          *logs
	Notifications *notifications
	Openapi       *openapi
	Operation     *operation
	Scheduler     *scheduler
	// Source - Source related resources.
	Source *source
	// SourceDefinition - SourceDefinition related resources.
	SourceDefinition *sourceDefinition
	// SourceDefinitionSpecification - SourceDefinition specification related resources.
	SourceDefinitionSpecification *sourceDefinitionSpecification
	// SourceOauth - Source OAuth related resources to delegate access from user.
	SourceOauth *sourceOauth
	// State - Interactions with state related resources.
	State *state
	// WebBackend - Endpoints for the Airbyte web application. Those APIs should not be called outside the web application implementation and are not
	// guaranteeing any backwards compatibility.
	//
	WebBackend *webBackend
	// Workspace - Workspace related resources.
	Workspace *workspace

	// Non-idiomatic field names below are to namespace fields from the fields names above to avoid name conflicts
	_defaultClient  HTTPClient
	_securityClient HTTPClient

	_serverURL  string
	_language   string
	_sdkVersion string
	_genVersion string
}

type SDKOption func(*AirbyteTest)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *AirbyteTest) {
		sdk._serverURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *AirbyteTest) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk._serverURL = serverURL
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *AirbyteTest) {
		sdk._defaultClient = client
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *AirbyteTest {
	sdk := &AirbyteTest{
		_language:   "go",
		_sdkVersion: "1.3.0",
		_genVersion: "2.34.2",
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk._defaultClient == nil {
		sdk._defaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk._securityClient == nil {
		sdk._securityClient = sdk._defaultClient
	}

	if sdk._serverURL == "" {
		sdk._serverURL = ServerList[0]
	}

	sdk.Attempt = newAttempt(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Connection = newConnection(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Destination = newDestination(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.DestinationDefinition = newDestinationDefinition(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.DestinationDefinitionSpecification = newDestinationDefinitionSpecification(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.DestinationOauth = newDestinationOauth(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Health = newHealth(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Internal = newInternal(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Jobs = newJobs(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Logs = newLogs(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Notifications = newNotifications(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Openapi = newOpenapi(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Operation = newOperation(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Scheduler = newScheduler(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Source = newSource(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.SourceDefinition = newSourceDefinition(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.SourceDefinitionSpecification = newSourceDefinitionSpecification(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.SourceOauth = newSourceOauth(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.State = newState(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.WebBackend = newWebBackend(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Workspace = newWorkspace(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	return sdk
}
