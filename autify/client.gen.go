// Package autify provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package autify

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetProjectsProjectIdResults request
	GetProjectsProjectIdResults(ctx context.Context, projectId int, params *GetProjectsProjectIdResultsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetProjectsProjectIdResultsResultId request
	GetProjectsProjectIdResultsResultId(ctx context.Context, projectId int, resultId int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetProjectsProjectIdScenarios request
	GetProjectsProjectIdScenarios(ctx context.Context, projectId int, params *GetProjectsProjectIdScenariosParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetProjectsProjectIdScenariosScenarioId request
	GetProjectsProjectIdScenariosScenarioId(ctx context.Context, projectId int, scenarioId int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostSchedulesScheduleId request
	PostSchedulesScheduleId(ctx context.Context, scheduleId int, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetProjectsProjectIdResults(ctx context.Context, projectId int, params *GetProjectsProjectIdResultsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetProjectsProjectIdResultsRequest(c.Server, projectId, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetProjectsProjectIdResultsResultId(ctx context.Context, projectId int, resultId int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetProjectsProjectIdResultsResultIdRequest(c.Server, projectId, resultId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetProjectsProjectIdScenarios(ctx context.Context, projectId int, params *GetProjectsProjectIdScenariosParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetProjectsProjectIdScenariosRequest(c.Server, projectId, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetProjectsProjectIdScenariosScenarioId(ctx context.Context, projectId int, scenarioId int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetProjectsProjectIdScenariosScenarioIdRequest(c.Server, projectId, scenarioId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostSchedulesScheduleId(ctx context.Context, scheduleId int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostSchedulesScheduleIdRequest(c.Server, scheduleId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetProjectsProjectIdResultsRequest generates requests for GetProjectsProjectIdResults
func NewGetProjectsProjectIdResultsRequest(server string, projectId int, params *GetProjectsProjectIdResultsParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "project_id", runtime.ParamLocationPath, projectId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/projects/%s/results", pathParam0)
	if operationPath[0] == '/' {
		operationPath = operationPath[1:]
	}
	operationURL := url.URL{
		Path: operationPath,
	}

	queryURL := serverURL.ResolveReference(&operationURL)

	queryValues := queryURL.Query()

	if params.Page != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.PerPage != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "per_page", runtime.ParamLocationQuery, *params.PerPage); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.TestPlanId != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "test_plan_id", runtime.ParamLocationQuery, *params.TestPlanId); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetProjectsProjectIdResultsResultIdRequest generates requests for GetProjectsProjectIdResultsResultId
func NewGetProjectsProjectIdResultsResultIdRequest(server string, projectId int, resultId int) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "project_id", runtime.ParamLocationPath, projectId)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "result_id", runtime.ParamLocationPath, resultId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/projects/%s/results/%s", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = operationPath[1:]
	}
	operationURL := url.URL{
		Path: operationPath,
	}

	queryURL := serverURL.ResolveReference(&operationURL)

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetProjectsProjectIdScenariosRequest generates requests for GetProjectsProjectIdScenarios
func NewGetProjectsProjectIdScenariosRequest(server string, projectId int, params *GetProjectsProjectIdScenariosParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "project_id", runtime.ParamLocationPath, projectId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/projects/%s/scenarios", pathParam0)
	if operationPath[0] == '/' {
		operationPath = operationPath[1:]
	}
	operationURL := url.URL{
		Path: operationPath,
	}

	queryURL := serverURL.ResolveReference(&operationURL)

	queryValues := queryURL.Query()

	if params.Page != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "page", runtime.ParamLocationQuery, *params.Page); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetProjectsProjectIdScenariosScenarioIdRequest generates requests for GetProjectsProjectIdScenariosScenarioId
func NewGetProjectsProjectIdScenariosScenarioIdRequest(server string, projectId int, scenarioId int) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "project_id", runtime.ParamLocationPath, projectId)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "scenario_id", runtime.ParamLocationPath, scenarioId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/projects/%s/scenarios/%s", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = operationPath[1:]
	}
	operationURL := url.URL{
		Path: operationPath,
	}

	queryURL := serverURL.ResolveReference(&operationURL)

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPostSchedulesScheduleIdRequest generates requests for PostSchedulesScheduleId
func NewPostSchedulesScheduleIdRequest(server string, scheduleId int) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "schedule_id", runtime.ParamLocationPath, scheduleId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/schedules/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = operationPath[1:]
	}
	operationURL := url.URL{
		Path: operationPath,
	}

	queryURL := serverURL.ResolveReference(&operationURL)

	req, err := http.NewRequest("POST", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetProjectsProjectIdResults request
	GetProjectsProjectIdResultsWithResponse(ctx context.Context, projectId int, params *GetProjectsProjectIdResultsParams, reqEditors ...RequestEditorFn) (*GetProjectsProjectIdResultsResponse, error)

	// GetProjectsProjectIdResultsResultId request
	GetProjectsProjectIdResultsResultIdWithResponse(ctx context.Context, projectId int, resultId int, reqEditors ...RequestEditorFn) (*GetProjectsProjectIdResultsResultIdResponse, error)

	// GetProjectsProjectIdScenarios request
	GetProjectsProjectIdScenariosWithResponse(ctx context.Context, projectId int, params *GetProjectsProjectIdScenariosParams, reqEditors ...RequestEditorFn) (*GetProjectsProjectIdScenariosResponse, error)

	// GetProjectsProjectIdScenariosScenarioId request
	GetProjectsProjectIdScenariosScenarioIdWithResponse(ctx context.Context, projectId int, scenarioId int, reqEditors ...RequestEditorFn) (*GetProjectsProjectIdScenariosScenarioIdResponse, error)

	// PostSchedulesScheduleId request
	PostSchedulesScheduleIdWithResponse(ctx context.Context, scheduleId int, reqEditors ...RequestEditorFn) (*PostSchedulesScheduleIdResponse, error)
}

type GetProjectsProjectIdResultsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]TestPlanResult
}

// Status returns HTTPResponse.Status
func (r GetProjectsProjectIdResultsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetProjectsProjectIdResultsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetProjectsProjectIdResultsResultIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		CreatedAt  *time.Time  `json:"created_at,omitempty"`
		Duration   *int        `json:"duration,omitempty"`
		FinishedAt *time.Time  `json:"finished_at,omitempty"`
		Id         *int        `json:"id,omitempty"`
		StartedAt  *time.Time  `json:"started_at,omitempty"`
		Status     *N200Status `json:"status,omitempty"`
		TestPlan   *struct {
			CreatedAt *time.Time `json:"created_at,omitempty"`
			Id        *int       `json:"id,omitempty"`
			Name      *string    `json:"name,omitempty"`
			UpdatedAt *time.Time `json:"updated_at,omitempty"`
		} `json:"test_plan,omitempty"`
		TestPlanCapabilityResults *[]struct {
			Capability      *Capability       `json:"capability,omitempty"`
			Id              *int              `json:"id,omitempty"`
			TestCaseResults *[]TestCaseResult `json:"test_case_results,omitempty"`
		} `json:"test_plan_capability_results,omitempty"`
		UpdatedAt *time.Time `json:"updated_at,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r GetProjectsProjectIdResultsResultIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetProjectsProjectIdResultsResultIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetProjectsProjectIdScenariosResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]Story
}

// Status returns HTTPResponse.Status
func (r GetProjectsProjectIdScenariosResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetProjectsProjectIdScenariosResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetProjectsProjectIdScenariosScenarioIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Story
}

// Status returns HTTPResponse.Status
func (r GetProjectsProjectIdScenariosScenarioIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetProjectsProjectIdScenariosScenarioIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostSchedulesScheduleIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]TestPlanResult
}

// Status returns HTTPResponse.Status
func (r PostSchedulesScheduleIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostSchedulesScheduleIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetProjectsProjectIdResultsWithResponse request returning *GetProjectsProjectIdResultsResponse
func (c *ClientWithResponses) GetProjectsProjectIdResultsWithResponse(ctx context.Context, projectId int, params *GetProjectsProjectIdResultsParams, reqEditors ...RequestEditorFn) (*GetProjectsProjectIdResultsResponse, error) {
	rsp, err := c.GetProjectsProjectIdResults(ctx, projectId, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetProjectsProjectIdResultsResponse(rsp)
}

// GetProjectsProjectIdResultsResultIdWithResponse request returning *GetProjectsProjectIdResultsResultIdResponse
func (c *ClientWithResponses) GetProjectsProjectIdResultsResultIdWithResponse(ctx context.Context, projectId int, resultId int, reqEditors ...RequestEditorFn) (*GetProjectsProjectIdResultsResultIdResponse, error) {
	rsp, err := c.GetProjectsProjectIdResultsResultId(ctx, projectId, resultId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetProjectsProjectIdResultsResultIdResponse(rsp)
}

// GetProjectsProjectIdScenariosWithResponse request returning *GetProjectsProjectIdScenariosResponse
func (c *ClientWithResponses) GetProjectsProjectIdScenariosWithResponse(ctx context.Context, projectId int, params *GetProjectsProjectIdScenariosParams, reqEditors ...RequestEditorFn) (*GetProjectsProjectIdScenariosResponse, error) {
	rsp, err := c.GetProjectsProjectIdScenarios(ctx, projectId, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetProjectsProjectIdScenariosResponse(rsp)
}

// GetProjectsProjectIdScenariosScenarioIdWithResponse request returning *GetProjectsProjectIdScenariosScenarioIdResponse
func (c *ClientWithResponses) GetProjectsProjectIdScenariosScenarioIdWithResponse(ctx context.Context, projectId int, scenarioId int, reqEditors ...RequestEditorFn) (*GetProjectsProjectIdScenariosScenarioIdResponse, error) {
	rsp, err := c.GetProjectsProjectIdScenariosScenarioId(ctx, projectId, scenarioId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetProjectsProjectIdScenariosScenarioIdResponse(rsp)
}

// PostSchedulesScheduleIdWithResponse request returning *PostSchedulesScheduleIdResponse
func (c *ClientWithResponses) PostSchedulesScheduleIdWithResponse(ctx context.Context, scheduleId int, reqEditors ...RequestEditorFn) (*PostSchedulesScheduleIdResponse, error) {
	rsp, err := c.PostSchedulesScheduleId(ctx, scheduleId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostSchedulesScheduleIdResponse(rsp)
}

// ParseGetProjectsProjectIdResultsResponse parses an HTTP response from a GetProjectsProjectIdResultsWithResponse call
func ParseGetProjectsProjectIdResultsResponse(rsp *http.Response) (*GetProjectsProjectIdResultsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetProjectsProjectIdResultsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []TestPlanResult
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetProjectsProjectIdResultsResultIdResponse parses an HTTP response from a GetProjectsProjectIdResultsResultIdWithResponse call
func ParseGetProjectsProjectIdResultsResultIdResponse(rsp *http.Response) (*GetProjectsProjectIdResultsResultIdResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetProjectsProjectIdResultsResultIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			CreatedAt  *time.Time  `json:"created_at,omitempty"`
			Duration   *int        `json:"duration,omitempty"`
			FinishedAt *time.Time  `json:"finished_at,omitempty"`
			Id         *int        `json:"id,omitempty"`
			StartedAt  *time.Time  `json:"started_at,omitempty"`
			Status     *N200Status `json:"status,omitempty"`
			TestPlan   *struct {
				CreatedAt *time.Time `json:"created_at,omitempty"`
				Id        *int       `json:"id,omitempty"`
				Name      *string    `json:"name,omitempty"`
				UpdatedAt *time.Time `json:"updated_at,omitempty"`
			} `json:"test_plan,omitempty"`
			TestPlanCapabilityResults *[]struct {
				Capability      *Capability       `json:"capability,omitempty"`
				Id              *int              `json:"id,omitempty"`
				TestCaseResults *[]TestCaseResult `json:"test_case_results,omitempty"`
			} `json:"test_plan_capability_results,omitempty"`
			UpdatedAt *time.Time `json:"updated_at,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetProjectsProjectIdScenariosResponse parses an HTTP response from a GetProjectsProjectIdScenariosWithResponse call
func ParseGetProjectsProjectIdScenariosResponse(rsp *http.Response) (*GetProjectsProjectIdScenariosResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetProjectsProjectIdScenariosResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []Story
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetProjectsProjectIdScenariosScenarioIdResponse parses an HTTP response from a GetProjectsProjectIdScenariosScenarioIdWithResponse call
func ParseGetProjectsProjectIdScenariosScenarioIdResponse(rsp *http.Response) (*GetProjectsProjectIdScenariosScenarioIdResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetProjectsProjectIdScenariosScenarioIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Story
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParsePostSchedulesScheduleIdResponse parses an HTTP response from a PostSchedulesScheduleIdWithResponse call
func ParsePostSchedulesScheduleIdResponse(rsp *http.Response) (*PostSchedulesScheduleIdResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &PostSchedulesScheduleIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []TestPlanResult
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

