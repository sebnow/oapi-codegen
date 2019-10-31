// Package petstore provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/sebnow/oapi-codegen DO NOT EDIT.
package petstore

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/sebnow/oapi-codegen/pkg/runtime"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Error defines model for Error.
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// NewPet defines model for NewPet.
type NewPet struct {
	Name string  `json:"name"`
	Tag  *string `json:"tag,omitempty"`
}

// Pet defines model for Pet.
type Pet struct {
	// Embedded struct due to allOf(#/components/schemas/NewPet)
	NewPet
	// Embedded fields due to inline allOf schema
	Id int64 `json:"id"`
}

// FindPetsParams defines parameters for FindPets.
type FindPetsParams struct {

	// tags to filter by
	Tags *[]string `json:"tags,omitempty"`

	// maximum number of results to return
	Limit *int32 `json:"limit,omitempty"`
}

// addPetJSONBody defines parameters for AddPet.
type addPetJSONBody NewPet

// AddPetRequestBody defines body for AddPet for application/json ContentType.
type AddPetJSONRequestBody addPetJSONBody

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(req *http.Request, ctx context.Context) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client Doer

	// A callback for modifying requests which are generated before sending over
	// the network.
	RequestEditor RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// The interface specification for the client above.
type ClientInterface interface {
	// FindPets request
	FindPets(ctx context.Context, params *FindPetsParams) (*http.Response, error)

	// AddPet request  with any body
	AddPetWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error)

	AddPet(ctx context.Context, body AddPetJSONRequestBody) (*http.Response, error)

	// DeletePet request
	DeletePet(ctx context.Context, id int64) (*http.Response, error)

	// FindPetById request
	FindPetById(ctx context.Context, id int64) (*http.Response, error)
}

func (c *Client) FindPets(ctx context.Context, params *FindPetsParams) (*http.Response, error) {
	req, err := NewFindPetsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(req, ctx)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) AddPetWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewAddPetRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(req, ctx)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) AddPet(ctx context.Context, body AddPetJSONRequestBody) (*http.Response, error) {
	req, err := NewAddPetRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(req, ctx)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) DeletePet(ctx context.Context, id int64) (*http.Response, error) {
	req, err := NewDeletePetRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(req, ctx)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) FindPetById(ctx context.Context, id int64) (*http.Response, error) {
	req, err := NewFindPetByIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(req, ctx)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

// NewFindPetsRequest generates requests for FindPets
func NewFindPetsRequest(server string, params *FindPetsParams) (*http.Request, error) {
	var err error

	queryUrl := fmt.Sprintf("%s/pets", server)

	var queryStrings []string

	var queryParam0 string
	if params.Tags != nil {

		queryParam0, err = runtime.StyleParam("form", true, "tags", *params.Tags)
		if err != nil {
			return nil, err
		}

		queryStrings = append(queryStrings, queryParam0)
	}

	var queryParam1 string
	if params.Limit != nil {

		queryParam1, err = runtime.StyleParam("form", true, "limit", *params.Limit)
		if err != nil {
			return nil, err
		}

		queryStrings = append(queryStrings, queryParam1)
	}

	if len(queryStrings) != 0 {
		queryUrl += "?" + strings.Join(queryStrings, "&")
	}

	req, err := http.NewRequest("GET", queryUrl, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewAddPetRequest calls the generic AddPet builder with application/json body
func NewAddPetRequest(server string, body AddPetJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAddPetRequestWithBody(server, "application/json", bodyReader)
}

// NewAddPetRequestWithBody generates requests for AddPet with any type of body
func NewAddPetRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl := fmt.Sprintf("%s/pets", server)

	req, err := http.NewRequest("POST", queryUrl, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// NewDeletePetRequest generates requests for DeletePet
func NewDeletePetRequest(server string, id int64) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "id", id)
	if err != nil {
		return nil, err
	}

	queryUrl := fmt.Sprintf("%s/pets/%s", server, pathParam0)

	req, err := http.NewRequest("DELETE", queryUrl, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewFindPetByIdRequest generates requests for FindPetById
func NewFindPetByIdRequest(server string, id int64) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "id", id)
	if err != nil {
		return nil, err
	}

	queryUrl := fmt.Sprintf("%s/pets/%s", server, pathParam0)

	req, err := http.NewRequest("GET", queryUrl, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClient creates a new Client.
func NewClient(ctx context.Context, opts ...ClientOption) (*ClientWithResponses, error) {
	// create a client with sane default values
	client := Client{
		// must have a slash in order to resolve relative paths correctly.
		Server: "",
	}
	// mutate defaultClient and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}

	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = http.DefaultClient
	}

	return &ClientWithResponses{
		ClientInterface: &client,
	}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		if !strings.HasSuffix(baseURL, "/") {
			baseURL += "/"
		}
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer Doer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditor = fn
		return nil
	}
}

// NewClientWithResponses returns a ClientWithResponses with a default Client:
func NewClientWithResponses(server string) *ClientWithResponses {
	return &ClientWithResponses{
		ClientInterface: &Client{
			Client: &http.Client{},
			Server: server,
		},
	}
}

// NewClientWithResponsesAndRequestEditorFunc takes in a RequestEditorFn callback function and returns a ClientWithResponses with a default Client:
func NewClientWithResponsesAndRequestEditorFunc(server string, reqEditorFn RequestEditorFn) *ClientWithResponses {
	return &ClientWithResponses{
		ClientInterface: &Client{
			Client:        &http.Client{},
			Server:        server,
			RequestEditor: reqEditorFn,
		},
	}
}

type findPetsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]Pet
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r findPetsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r findPetsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type addPetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Pet
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r addPetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r addPetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type deletePetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r deletePetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r deletePetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type findPetByIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Pet
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r findPetByIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r findPetByIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// FindPetsWithResponse request returning *FindPetsResponse
func (c *ClientWithResponses) FindPetsWithResponse(ctx context.Context, params *FindPetsParams) (*findPetsResponse, error) {
	rsp, err := c.FindPets(ctx, params)
	if err != nil {
		return nil, err
	}
	return ParsefindPetsResponse(rsp)
}

// AddPetWithBodyWithResponse request with arbitrary body returning *AddPetResponse
func (c *ClientWithResponses) AddPetWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*addPetResponse, error) {
	rsp, err := c.AddPetWithBody(ctx, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseaddPetResponse(rsp)
}

func (c *ClientWithResponses) AddPetWithResponse(ctx context.Context, body AddPetJSONRequestBody) (*addPetResponse, error) {
	rsp, err := c.AddPet(ctx, body)
	if err != nil {
		return nil, err
	}
	return ParseaddPetResponse(rsp)
}

// DeletePetWithResponse request returning *DeletePetResponse
func (c *ClientWithResponses) DeletePetWithResponse(ctx context.Context, id int64) (*deletePetResponse, error) {
	rsp, err := c.DeletePet(ctx, id)
	if err != nil {
		return nil, err
	}
	return ParsedeletePetResponse(rsp)
}

// FindPetByIdWithResponse request returning *FindPetByIdResponse
func (c *ClientWithResponses) FindPetByIdWithResponse(ctx context.Context, id int64) (*findPetByIdResponse, error) {
	rsp, err := c.FindPetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return ParsefindPetByIdResponse(rsp)
}

// ParsefindPetsResponse parses an HTTP response from a FindPetsWithResponse call
func ParsefindPetsResponse(rsp *http.Response) (*findPetsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &findPetsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		response.JSON200 = &[]Pet{}
		if err := json.Unmarshal(bodyBytes, response.JSON200); err != nil {
			return nil, err
		}

	case strings.Contains(rsp.Header.Get("Content-Type"), "json"):
		response.JSONDefault = &Error{}
		if err := json.Unmarshal(bodyBytes, response.JSONDefault); err != nil {
			return nil, err
		}

	}

	return response, nil
}

// ParseaddPetResponse parses an HTTP response from a AddPetWithResponse call
func ParseaddPetResponse(rsp *http.Response) (*addPetResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &addPetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		response.JSON200 = &Pet{}
		if err := json.Unmarshal(bodyBytes, response.JSON200); err != nil {
			return nil, err
		}

	case strings.Contains(rsp.Header.Get("Content-Type"), "json"):
		response.JSONDefault = &Error{}
		if err := json.Unmarshal(bodyBytes, response.JSONDefault); err != nil {
			return nil, err
		}

	}

	return response, nil
}

// ParsedeletePetResponse parses an HTTP response from a DeletePetWithResponse call
func ParsedeletePetResponse(rsp *http.Response) (*deletePetResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &deletePetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json"):
		response.JSONDefault = &Error{}
		if err := json.Unmarshal(bodyBytes, response.JSONDefault); err != nil {
			return nil, err
		}

	}

	return response, nil
}

// ParsefindPetByIdResponse parses an HTTP response from a FindPetByIdWithResponse call
func ParsefindPetByIdResponse(rsp *http.Response) (*findPetByIdResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &findPetByIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		response.JSON200 = &Pet{}
		if err := json.Unmarshal(bodyBytes, response.JSON200); err != nil {
			return nil, err
		}

	case strings.Contains(rsp.Header.Get("Content-Type"), "json"):
		response.JSONDefault = &Error{}
		if err := json.Unmarshal(bodyBytes, response.JSONDefault); err != nil {
			return nil, err
		}

	}

	return response, nil
}
