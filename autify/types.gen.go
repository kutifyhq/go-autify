// Package autify provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package autify

import (
	"time"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Defines values for TestCaseResultStatus.
const (
	TestCaseResultStatusFailed TestCaseResultStatus = "failed"

	TestCaseResultStatusPassed TestCaseResultStatus = "passed"

	TestCaseResultStatusRunning TestCaseResultStatus = "running"

	TestCaseResultStatusSkipped TestCaseResultStatus = "skipped"

	TestCaseResultStatusWaiting TestCaseResultStatus = "waiting"
)

// Defines values for TestPlanResultStatus.
const (
	TestPlanResultStatusFailed TestPlanResultStatus = "failed"

	TestPlanResultStatusPassed TestPlanResultStatus = "passed"

	TestPlanResultStatusRunning TestPlanResultStatus = "running"

	TestPlanResultStatusSkipped TestPlanResultStatus = "skipped"

	TestPlanResultStatusWaiting TestPlanResultStatus = "waiting"
)

// Capability defines model for Capability.
type Capability struct {
	Browser        *string    `json:"browser,omitempty"`
	BrowserVersion *string    `json:"browser_version,omitempty"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	Device         *string    `json:"device,omitempty"`
	Id             *int       `json:"id,omitempty"`
	Os             *string    `json:"os,omitempty"`
	OsVersion      *string    `json:"os_version,omitempty"`
	Resolution     *string    `json:"resolution,omitempty"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
}

// Story defines model for Story.
type Story struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Id        *int       `json:"id,omitempty"`
	Name      *string    `json:"name,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// TestCaseResult defines model for TestCaseResult.
type TestCaseResult struct {
	CreatedAt  *time.Time            `json:"created_at,omitempty"`
	Duration   *int                  `json:"duration,omitempty"`
	FinishedAt *time.Time            `json:"finished_at,omitempty"`
	Id         *int                  `json:"id,omitempty"`
	StartedAt  *time.Time            `json:"started_at,omitempty"`
	Status     *TestCaseResultStatus `json:"status,omitempty"`
	UpdatedAt  *time.Time            `json:"updated_at,omitempty"`
}

// TestCaseResultStatus defines model for TestCaseResult.Status.
type TestCaseResultStatus string

// TestPlanResult defines model for TestPlanResult.
type TestPlanResult struct {
	CreatedAt  *time.Time            `json:"created_at,omitempty"`
	Duration   *int                  `json:"duration,omitempty"`
	FinishedAt *time.Time            `json:"finished_at,omitempty"`
	Id         *int                  `json:"id,omitempty"`
	StartedAt  *time.Time            `json:"started_at,omitempty"`
	Status     *TestPlanResultStatus `json:"status,omitempty"`
	TestPlan   *struct {
		CreatedAt *time.Time `json:"created_at,omitempty"`
		Id        *int       `json:"id,omitempty"`
		Name      *string    `json:"name,omitempty"`
		UpdatedAt *time.Time `json:"updated_at,omitempty"`
	} `json:"test_plan,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// TestPlanResultStatus defines model for TestPlanResult.Status.
type TestPlanResultStatus string

// GetProjectsProjectIdResultsParams defines parameters for GetProjectsProjectIdResults.
type GetProjectsProjectIdResultsParams struct {

	// The number of page returns.
	Page *int `json:"page,omitempty"`

	// The number of items returns. Default number is 30 and up to a maximum of 100
	PerPage *int `json:"per_page,omitempty"`

	// Test plan ID used to filter results.
	TestPlanId *int `json:"test_plan_id,omitempty"`
}

// GetProjectsProjectIdScenariosParams defines parameters for GetProjectsProjectIdScenarios.
type GetProjectsProjectIdScenariosParams struct {

	// The number of page returns.
	Page *int `json:"page,omitempty"`
}

