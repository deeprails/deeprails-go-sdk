// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package deeprails

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/deeprails/deeprails-go-sdk/internal/apijson"
	"github.com/deeprails/deeprails-go-sdk/internal/apiquery"
	"github.com/deeprails/deeprails-go-sdk/internal/param"
	"github.com/deeprails/deeprails-go-sdk/internal/requestconfig"
	"github.com/deeprails/deeprails-go-sdk/option"
	"github.com/deeprails/deeprails-go-sdk/packages/ssestream"
)

// DefendService contains methods and other services that help with interacting
// with the deep rails API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDefendService] method instead.
type DefendService struct {
	Options []option.RequestOption
}

// NewDefendService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDefendService(opts ...option.RequestOption) (r *DefendService) {
	r = &DefendService{}
	r.Options = opts
	return
}

// Use this endpoint to create a new guardrail workflow by specifying guardrail
// thresholds, an improvement action, and optional extended capabilities.
func (r *DefendService) NewWorkflow(ctx context.Context, body DefendNewWorkflowParams, opts ...option.RequestOption) (res *DefendCreateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "defend"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Use this endpoint to retrieve a specific event of a guardrail workflow
func (r *DefendService) GetEvent(ctx context.Context, workflowID string, eventID string, opts ...option.RequestOption) (res *WorkflowEventDetailResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if workflowID == "" {
		err = errors.New("missing required workflow_id parameter")
		return
	}
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	path := fmt.Sprintf("defend/%s/events/%s", workflowID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Use this endpoint to retrieve the details for a specific defend workflow
func (r *DefendService) GetWorkflow(ctx context.Context, workflowID string, query DefendGetWorkflowParams, opts ...option.RequestOption) (res *DefendResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if workflowID == "" {
		err = errors.New("missing required workflow_id parameter")
		return
	}
	path := fmt.Sprintf("defend/%s", workflowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Use this endpoint to submit a model input and output pair to a workflow for
// evaluation with streaming responses.
func (r *DefendService) SubmitAndStreamEventStreaming(ctx context.Context, workflowID string, params DefendSubmitAndStreamEventParams, opts ...option.RequestOption) (stream *ssestream.Stream[string]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	if workflowID == "" {
		err = errors.New("missing required workflow_id parameter")
		return
	}
	path := fmt.Sprintf("defend/%s/events?stream=true", workflowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &raw, opts...)
	return ssestream.NewStream[string](ssestream.NewDecoder(raw), err)
}

// Use this endpoint to submit a model input and output pair to a workflow for
// evaluation
func (r *DefendService) SubmitEvent(ctx context.Context, workflowID string, body DefendSubmitEventParams, opts ...option.RequestOption) (res *WorkflowEventResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if workflowID == "" {
		err = errors.New("missing required workflow_id parameter")
		return
	}
	path := fmt.Sprintf("defend/%s/events", workflowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Use this endpoint to update an existing defend workflow if its details change.
func (r *DefendService) UpdateWorkflow(ctx context.Context, workflowID string, body DefendUpdateWorkflowParams, opts ...option.RequestOption) (res *DefendUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if workflowID == "" {
		err = errors.New("missing required workflow_id parameter")
		return
	}
	path := fmt.Sprintf("defend/%s", workflowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type DefendCreateResponse struct {
	// The time the workflow was created in UTC.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Status of the selected workflow. May be `inactive` or `active`. Inactive
	// workflows will not accept events.
	Status DefendCreateResponseStatus `json:"status" api:"required"`
	// A unique workflow ID.
	WorkflowID string                   `json:"workflow_id" api:"required"`
	JSON       defendCreateResponseJSON `json:"-"`
}

// defendCreateResponseJSON contains the JSON metadata for the struct
// [DefendCreateResponse]
type defendCreateResponseJSON struct {
	CreatedAt   apijson.Field
	Status      apijson.Field
	WorkflowID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DefendCreateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r defendCreateResponseJSON) RawJSON() string {
	return r.raw
}

// Status of the selected workflow. May be `inactive` or `active`. Inactive
// workflows will not accept events.
type DefendCreateResponseStatus string

const (
	DefendCreateResponseStatusInactive DefendCreateResponseStatus = "inactive"
	DefendCreateResponseStatusActive   DefendCreateResponseStatus = "active"
)

func (r DefendCreateResponseStatus) IsKnown() bool {
	switch r {
	case DefendCreateResponseStatusInactive, DefendCreateResponseStatusActive:
		return true
	}
	return false
}

type DefendResponse struct {
	// Mapping of guardrail metric names to tolerance values. Values can be strings
	// (`low`, `medium`, `high`) for automatic tolerance levels.
	AutomaticHallucinationToleranceLevels map[string]DefendResponseAutomaticHallucinationToleranceLevel `json:"automatic_hallucination_tolerance_levels" api:"required"`
	// Extended AI capabilities available to the event, if any. Can be `web_search`,
	// `file_search`, and/or `context_awareness`.
	Capabilities []DefendResponseCapability `json:"capabilities" api:"required"`
	// The time the workflow was created in UTC.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Mapping of guardrail metric names to threshold values. Values can be floating
	// point numbers (0.0-1.0) for custom thresholds.
	CustomHallucinationThresholdValues map[string]float64 `json:"custom_hallucination_threshold_values" api:"required"`
	// A description for the workflow, to help you remember what that workflow means to
	// your organization.
	Description string `json:"description" api:"required"`
	// An array of events associated with this workflow.
	Events []DefendResponseEvent `json:"events" api:"required"`
	// List of files associated with the workflow. If this is not empty, models can
	// search these files when performing evaluations or remediations
	Files []DefendResponseFile `json:"files" api:"required"`
	// A human-readable name for the workflow that will correspond to it's workflow ID.
	Name string `json:"name" api:"required"`
	// Status of the selected workflow. May be `inactive` or `active`. Inactive
	// workflows will not accept events.
	Status DefendResponseStatus `json:"status" api:"required"`
	// Type of thresholds used to evaluate the event.
	ThresholdType DefendResponseThresholdType `json:"threshold_type" api:"required"`
	// The most recent time the workflow was updated in UTC.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// A unique workflow ID used to identify the workflow in other endpoints.
	WorkflowID string `json:"workflow_id" api:"required"`
	// The action used to improve outputs that fail one or more guardrail metrics for
	// the workflow events.
	ImprovementAction DefendResponseImprovementAction `json:"improvement_action"`
	Stats             DefendResponseStats             `json:"stats"`
	JSON              defendResponseJSON              `json:"-"`
}

// defendResponseJSON contains the JSON metadata for the struct [DefendResponse]
type defendResponseJSON struct {
	AutomaticHallucinationToleranceLevels apijson.Field
	Capabilities                          apijson.Field
	CreatedAt                             apijson.Field
	CustomHallucinationThresholdValues    apijson.Field
	Description                           apijson.Field
	Events                                apijson.Field
	Files                                 apijson.Field
	Name                                  apijson.Field
	Status                                apijson.Field
	ThresholdType                         apijson.Field
	UpdatedAt                             apijson.Field
	WorkflowID                            apijson.Field
	ImprovementAction                     apijson.Field
	Stats                                 apijson.Field
	raw                                   string
	ExtraFields                           map[string]apijson.Field
}

func (r *DefendResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r defendResponseJSON) RawJSON() string {
	return r.raw
}

type DefendResponseAutomaticHallucinationToleranceLevel string

const (
	DefendResponseAutomaticHallucinationToleranceLevelLow    DefendResponseAutomaticHallucinationToleranceLevel = "low"
	DefendResponseAutomaticHallucinationToleranceLevelMedium DefendResponseAutomaticHallucinationToleranceLevel = "medium"
	DefendResponseAutomaticHallucinationToleranceLevelHigh   DefendResponseAutomaticHallucinationToleranceLevel = "high"
)

func (r DefendResponseAutomaticHallucinationToleranceLevel) IsKnown() bool {
	switch r {
	case DefendResponseAutomaticHallucinationToleranceLevelLow, DefendResponseAutomaticHallucinationToleranceLevelMedium, DefendResponseAutomaticHallucinationToleranceLevelHigh:
		return true
	}
	return false
}

type DefendResponseCapability struct {
	Capability string                       `json:"capability"`
	JSON       defendResponseCapabilityJSON `json:"-"`
}

// defendResponseCapabilityJSON contains the JSON metadata for the struct
// [DefendResponseCapability]
type defendResponseCapabilityJSON struct {
	Capability  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DefendResponseCapability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r defendResponseCapabilityJSON) RawJSON() string {
	return r.raw
}

type DefendResponseEvent struct {
	// The ID of the billing request for the event.
	BillingRequestID string `json:"billing_request_id"`
	// An array of evaluations for this event.
	Evaluations []DefendResponseEventsEvaluation `json:"evaluations"`
	// A unique workflow event ID.
	EventID string `json:"event_id"`
	// Improved model output after improvement tool was applied.
	ImprovedModelOutput string `json:"improved_model_output"`
	// Status of the improvement tool used to improve the event. `improvement_required`
	// indicates that the evaluation is complete and the improvement action is needed
	// but is not taking place. `improved` and `improvement_failed` indicate when the
	// improvement action concludes, successfully and unsuccessfully, respectively.
	// `no_improvement_required` means that the first evaluation passed all its
	// metrics!
	ImprovementToolStatus DefendResponseEventsImprovementToolStatus `json:"improvement_tool_status"`
	// Status of the event.
	Status DefendResponseEventsStatus `json:"status"`
	JSON   defendResponseEventJSON    `json:"-"`
}

// defendResponseEventJSON contains the JSON metadata for the struct
// [DefendResponseEvent]
type defendResponseEventJSON struct {
	BillingRequestID      apijson.Field
	Evaluations           apijson.Field
	EventID               apijson.Field
	ImprovedModelOutput   apijson.Field
	ImprovementToolStatus apijson.Field
	Status                apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *DefendResponseEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r defendResponseEventJSON) RawJSON() string {
	return r.raw
}

type DefendResponseEventsEvaluation struct {
	// Analysis of the failures of the model_output according to the guardrail metrics
	// evaluated.
	AnalysisOfFailures string `json:"analysis_of_failures"`
	// The attempt number or identifier for this evaluation.
	Attempt string `json:"attempt"`
	// The time the evaluation was created in UTC.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Error message if the evaluation failed.
	ErrorMessage string `json:"error_message"`
	// The result of the evaluation.
	EvaluationResult map[string]interface{} `json:"evaluation_result"`
	// Status of the evaluation.
	EvaluationStatus string `json:"evaluation_status"`
	// Total cost of the evaluation.
	EvaluationTotalCost float64 `json:"evaluation_total_cost"`
	// An array of guardrail metrics evaluated.
	GuardrailMetrics []string `json:"guardrail_metrics"`
	// Status of the improvement tool used to improve the event. `improvement_required`
	// indicates that the evaluation is complete and the improvement action is needed
	// but is not taking place. `improved` and `improvement_failed` indicate when the
	// improvement action concludes, successfully and unsuccessfully, respectively.
	// `no_improvement_required` means that the first evaluation passed all its
	// metrics!
	ImprovementToolStatus DefendResponseEventsEvaluationsImprovementToolStatus `json:"improvement_tool_status"`
	// A list of key improvements made to the model_output to address the failures.
	KeyImprovements []string `json:"key_improvements"`
	// The model input used for the evaluation.
	ModelInput map[string]interface{} `json:"model_input"`
	// The model output that was evaluated.
	ModelOutput string `json:"model_output"`
	// The time the evaluation was last modified in UTC.
	ModifiedAt time.Time `json:"modified_at" format:"date-time"`
	// An optional tag for the evaluation.
	Nametag string `json:"nametag"`
	// Evaluation progress (0-100).
	Progress int64 `json:"progress"`
	// Run mode used for the evaluation.
	RunMode string                             `json:"run_mode"`
	JSON    defendResponseEventsEvaluationJSON `json:"-"`
}

// defendResponseEventsEvaluationJSON contains the JSON metadata for the struct
// [DefendResponseEventsEvaluation]
type defendResponseEventsEvaluationJSON struct {
	AnalysisOfFailures    apijson.Field
	Attempt               apijson.Field
	CreatedAt             apijson.Field
	ErrorMessage          apijson.Field
	EvaluationResult      apijson.Field
	EvaluationStatus      apijson.Field
	EvaluationTotalCost   apijson.Field
	GuardrailMetrics      apijson.Field
	ImprovementToolStatus apijson.Field
	KeyImprovements       apijson.Field
	ModelInput            apijson.Field
	ModelOutput           apijson.Field
	ModifiedAt            apijson.Field
	Nametag               apijson.Field
	Progress              apijson.Field
	RunMode               apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *DefendResponseEventsEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r defendResponseEventsEvaluationJSON) RawJSON() string {
	return r.raw
}

// Status of the improvement tool used to improve the event. `improvement_required`
// indicates that the evaluation is complete and the improvement action is needed
// but is not taking place. `improved` and `improvement_failed` indicate when the
// improvement action concludes, successfully and unsuccessfully, respectively.
// `no_improvement_required` means that the first evaluation passed all its
// metrics!
type DefendResponseEventsEvaluationsImprovementToolStatus string

const (
	DefendResponseEventsEvaluationsImprovementToolStatusImproved              DefendResponseEventsEvaluationsImprovementToolStatus = "improved"
	DefendResponseEventsEvaluationsImprovementToolStatusImprovementFailed     DefendResponseEventsEvaluationsImprovementToolStatus = "improvement_failed"
	DefendResponseEventsEvaluationsImprovementToolStatusNoImprovementRequired DefendResponseEventsEvaluationsImprovementToolStatus = "no_improvement_required"
	DefendResponseEventsEvaluationsImprovementToolStatusImprovementRequired   DefendResponseEventsEvaluationsImprovementToolStatus = "improvement_required"
)

func (r DefendResponseEventsEvaluationsImprovementToolStatus) IsKnown() bool {
	switch r {
	case DefendResponseEventsEvaluationsImprovementToolStatusImproved, DefendResponseEventsEvaluationsImprovementToolStatusImprovementFailed, DefendResponseEventsEvaluationsImprovementToolStatusNoImprovementRequired, DefendResponseEventsEvaluationsImprovementToolStatusImprovementRequired:
		return true
	}
	return false
}

// Status of the improvement tool used to improve the event. `improvement_required`
// indicates that the evaluation is complete and the improvement action is needed
// but is not taking place. `improved` and `improvement_failed` indicate when the
// improvement action concludes, successfully and unsuccessfully, respectively.
// `no_improvement_required` means that the first evaluation passed all its
// metrics!
type DefendResponseEventsImprovementToolStatus string

const (
	DefendResponseEventsImprovementToolStatusImproved              DefendResponseEventsImprovementToolStatus = "improved"
	DefendResponseEventsImprovementToolStatusImprovementFailed     DefendResponseEventsImprovementToolStatus = "improvement_failed"
	DefendResponseEventsImprovementToolStatusNoImprovementRequired DefendResponseEventsImprovementToolStatus = "no_improvement_required"
	DefendResponseEventsImprovementToolStatusImprovementRequired   DefendResponseEventsImprovementToolStatus = "improvement_required"
)

func (r DefendResponseEventsImprovementToolStatus) IsKnown() bool {
	switch r {
	case DefendResponseEventsImprovementToolStatusImproved, DefendResponseEventsImprovementToolStatusImprovementFailed, DefendResponseEventsImprovementToolStatusNoImprovementRequired, DefendResponseEventsImprovementToolStatusImprovementRequired:
		return true
	}
	return false
}

// Status of the event.
type DefendResponseEventsStatus string

const (
	DefendResponseEventsStatusCompleted  DefendResponseEventsStatus = "completed"
	DefendResponseEventsStatusFailed     DefendResponseEventsStatus = "failed"
	DefendResponseEventsStatusInProgress DefendResponseEventsStatus = "in_progress"
)

func (r DefendResponseEventsStatus) IsKnown() bool {
	switch r {
	case DefendResponseEventsStatusCompleted, DefendResponseEventsStatusFailed, DefendResponseEventsStatusInProgress:
		return true
	}
	return false
}

type DefendResponseFile struct {
	FileID                string                 `json:"file_id"`
	FileName              string                 `json:"file_name"`
	FileSize              int64                  `json:"file_size"`
	PresignedURL          string                 `json:"presigned_url"`
	PresignedURLExpiresAt time.Time              `json:"presigned_url_expires_at" format:"date-time"`
	JSON                  defendResponseFileJSON `json:"-"`
}

// defendResponseFileJSON contains the JSON metadata for the struct
// [DefendResponseFile]
type defendResponseFileJSON struct {
	FileID                apijson.Field
	FileName              apijson.Field
	FileSize              apijson.Field
	PresignedURL          apijson.Field
	PresignedURLExpiresAt apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *DefendResponseFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r defendResponseFileJSON) RawJSON() string {
	return r.raw
}

// Status of the selected workflow. May be `inactive` or `active`. Inactive
// workflows will not accept events.
type DefendResponseStatus string

const (
	DefendResponseStatusInactive DefendResponseStatus = "inactive"
	DefendResponseStatusActive   DefendResponseStatus = "active"
)

func (r DefendResponseStatus) IsKnown() bool {
	switch r {
	case DefendResponseStatusInactive, DefendResponseStatusActive:
		return true
	}
	return false
}

// Type of thresholds used to evaluate the event.
type DefendResponseThresholdType string

const (
	DefendResponseThresholdTypeCustom    DefendResponseThresholdType = "custom"
	DefendResponseThresholdTypeAutomatic DefendResponseThresholdType = "automatic"
)

func (r DefendResponseThresholdType) IsKnown() bool {
	switch r {
	case DefendResponseThresholdTypeCustom, DefendResponseThresholdTypeAutomatic:
		return true
	}
	return false
}

// The action used to improve outputs that fail one or more guardrail metrics for
// the workflow events.
type DefendResponseImprovementAction string

const (
	DefendResponseImprovementActionRegen     DefendResponseImprovementAction = "regen"
	DefendResponseImprovementActionFixit     DefendResponseImprovementAction = "fixit"
	DefendResponseImprovementActionDoNothing DefendResponseImprovementAction = "do_nothing"
)

func (r DefendResponseImprovementAction) IsKnown() bool {
	switch r {
	case DefendResponseImprovementActionRegen, DefendResponseImprovementActionFixit, DefendResponseImprovementActionDoNothing:
		return true
	}
	return false
}

type DefendResponseStats struct {
	// Number of AI outputs that failed the guardrails.
	OutputsBelowThreshold int64 `json:"outputs_below_threshold"`
	// Number of AI outputs that were improved.
	OutputsImproved int64 `json:"outputs_improved"`
	// Total number of AI outputs processed by the workflow.
	OutputsProcessed int64                   `json:"outputs_processed"`
	JSON             defendResponseStatsJSON `json:"-"`
}

// defendResponseStatsJSON contains the JSON metadata for the struct
// [DefendResponseStats]
type defendResponseStatsJSON struct {
	OutputsBelowThreshold apijson.Field
	OutputsImproved       apijson.Field
	OutputsProcessed      apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *DefendResponseStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r defendResponseStatsJSON) RawJSON() string {
	return r.raw
}

type DefendUpdateResponse struct {
	// The time the workflow was last modified in UTC.
	ModifiedAt time.Time `json:"modified_at" api:"required" format:"date-time"`
	// Status of the selected workflow. May be `inactive` or `active`. Inactive
	// workflows will not accept events.
	Status DefendUpdateResponseStatus `json:"status" api:"required"`
	// A unique workflow ID.
	WorkflowID string `json:"workflow_id" api:"required"`
	// The name of the workflow.
	Name string                   `json:"name"`
	JSON defendUpdateResponseJSON `json:"-"`
}

// defendUpdateResponseJSON contains the JSON metadata for the struct
// [DefendUpdateResponse]
type defendUpdateResponseJSON struct {
	ModifiedAt  apijson.Field
	Status      apijson.Field
	WorkflowID  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DefendUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r defendUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Status of the selected workflow. May be `inactive` or `active`. Inactive
// workflows will not accept events.
type DefendUpdateResponseStatus string

const (
	DefendUpdateResponseStatusInactive DefendUpdateResponseStatus = "inactive"
	DefendUpdateResponseStatusActive   DefendUpdateResponseStatus = "active"
)

func (r DefendUpdateResponseStatus) IsKnown() bool {
	switch r {
	case DefendUpdateResponseStatusInactive, DefendUpdateResponseStatusActive:
		return true
	}
	return false
}

type WorkflowEventDetailResponse struct {
	AnalysisOfFailures []string `json:"analysis_of_failures" api:"required"`
	// History of evaluations for the event.
	EvaluationHistory []WorkflowEventDetailResponseEvaluationHistory `json:"evaluation_history" api:"required"`
	// Evaluation result consisting of average scores and rationales for each of the
	// evaluated guardrail metrics.
	EvaluationResult map[string]interface{} `json:"evaluation_result" api:"required"`
	// A unique workflow event ID.
	EventID string `json:"event_id" api:"required"`
	// Whether the event was filtered and requires improvement.
	Filtered bool `json:"filtered" api:"required"`
	// Improved model output after improvement tool was applied and each metric passed
	// evaluation.
	ImprovedModelOutput string `json:"improved_model_output" api:"required"`
	// Type of improvement action used to improve the event.
	ImprovementAction WorkflowEventDetailResponseImprovementAction `json:"improvement_action" api:"required"`
	// Status of the improvement tool used to improve the event. `improvement_required`
	// indicates that the evaluation is complete and the improvement action is needed
	// but is not taking place. `improved` and `improvement_failed` indicate when the
	// improvement action concludes, successfully and unsuccessfully, respectively.
	// `no_improvement_required` means that the first evaluation passed all its
	// metrics!
	ImprovementToolStatus WorkflowEventDetailResponseImprovementToolStatus `json:"improvement_tool_status" api:"required,nullable"`
	KeyImprovements       []interface{}                                    `json:"key_improvements" api:"required"`
	// Status of the event.
	Status WorkflowEventDetailResponseStatus `json:"status" api:"required"`
	// Type of thresholds used to evaluate the event.
	ThresholdType WorkflowEventDetailResponseThresholdType `json:"threshold_type" api:"required"`
	// Workflow ID associated with the event.
	WorkflowID string `json:"workflow_id" api:"required"`
	// Mapping of guardrail metric names to tolerance values. Values are strings
	// (`low`, `medium`, `high`) representing automatic tolerance levels.
	AutomaticHallucinationToleranceLevels map[string]WorkflowEventDetailResponseAutomaticHallucinationToleranceLevel `json:"automatic_hallucination_tolerance_levels"`
	// Extended AI capabilities available to the event, if any. Can be `web_search`,
	// `file_search`, and/or `context_awareness`.
	Capabilities []WorkflowEventDetailResponseCapability `json:"capabilities"`
	// Mapping of guardrail metric names to threshold values. Values are floating point
	// numbers (0.0-1.0) representing custom thresholds.
	CustomHallucinationThresholdValues map[string]float64 `json:"custom_hallucination_threshold_values"`
	// List of files available to the event, if any. Will only be present if
	// `file_search` is enabled.
	Files []WorkflowEventDetailResponseFile `json:"files"`
	// The maximum number of improvement attempts to be applied to one event before it
	// is considered failed.
	MaxImprovementAttempts int64                           `json:"max_improvement_attempts"`
	JSON                   workflowEventDetailResponseJSON `json:"-"`
}

// workflowEventDetailResponseJSON contains the JSON metadata for the struct
// [WorkflowEventDetailResponse]
type workflowEventDetailResponseJSON struct {
	AnalysisOfFailures                    apijson.Field
	EvaluationHistory                     apijson.Field
	EvaluationResult                      apijson.Field
	EventID                               apijson.Field
	Filtered                              apijson.Field
	ImprovedModelOutput                   apijson.Field
	ImprovementAction                     apijson.Field
	ImprovementToolStatus                 apijson.Field
	KeyImprovements                       apijson.Field
	Status                                apijson.Field
	ThresholdType                         apijson.Field
	WorkflowID                            apijson.Field
	AutomaticHallucinationToleranceLevels apijson.Field
	Capabilities                          apijson.Field
	CustomHallucinationThresholdValues    apijson.Field
	Files                                 apijson.Field
	MaxImprovementAttempts                apijson.Field
	raw                                   string
	ExtraFields                           map[string]apijson.Field
}

func (r *WorkflowEventDetailResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowEventDetailResponseJSON) RawJSON() string {
	return r.raw
}

type WorkflowEventDetailResponseEvaluationHistory struct {
	AnalysisOfFailures    string                                                            `json:"analysis_of_failures"`
	Attempt               string                                                            `json:"attempt"`
	CreatedAt             time.Time                                                         `json:"created_at" format:"date-time"`
	ErrorMessage          string                                                            `json:"error_message"`
	EvaluationResult      map[string]interface{}                                            `json:"evaluation_result"`
	EvaluationStatus      string                                                            `json:"evaluation_status"`
	EvaluationTotalCost   float64                                                           `json:"evaluation_total_cost"`
	GuardrailMetrics      []string                                                          `json:"guardrail_metrics"`
	ImprovementToolStatus WorkflowEventDetailResponseEvaluationHistoryImprovementToolStatus `json:"improvement_tool_status"`
	KeyImprovements       []string                                                          `json:"key_improvements"`
	ModelInput            map[string]interface{}                                            `json:"model_input"`
	ModelOutput           string                                                            `json:"model_output"`
	Nametag               string                                                            `json:"nametag"`
	Progress              int64                                                             `json:"progress"`
	RunMode               string                                                            `json:"run_mode"`
	JSON                  workflowEventDetailResponseEvaluationHistoryJSON                  `json:"-"`
}

// workflowEventDetailResponseEvaluationHistoryJSON contains the JSON metadata for
// the struct [WorkflowEventDetailResponseEvaluationHistory]
type workflowEventDetailResponseEvaluationHistoryJSON struct {
	AnalysisOfFailures    apijson.Field
	Attempt               apijson.Field
	CreatedAt             apijson.Field
	ErrorMessage          apijson.Field
	EvaluationResult      apijson.Field
	EvaluationStatus      apijson.Field
	EvaluationTotalCost   apijson.Field
	GuardrailMetrics      apijson.Field
	ImprovementToolStatus apijson.Field
	KeyImprovements       apijson.Field
	ModelInput            apijson.Field
	ModelOutput           apijson.Field
	Nametag               apijson.Field
	Progress              apijson.Field
	RunMode               apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *WorkflowEventDetailResponseEvaluationHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowEventDetailResponseEvaluationHistoryJSON) RawJSON() string {
	return r.raw
}

type WorkflowEventDetailResponseEvaluationHistoryImprovementToolStatus string

const (
	WorkflowEventDetailResponseEvaluationHistoryImprovementToolStatusImproved              WorkflowEventDetailResponseEvaluationHistoryImprovementToolStatus = "improved"
	WorkflowEventDetailResponseEvaluationHistoryImprovementToolStatusImprovementFailed     WorkflowEventDetailResponseEvaluationHistoryImprovementToolStatus = "improvement_failed"
	WorkflowEventDetailResponseEvaluationHistoryImprovementToolStatusNoImprovementRequired WorkflowEventDetailResponseEvaluationHistoryImprovementToolStatus = "no_improvement_required"
	WorkflowEventDetailResponseEvaluationHistoryImprovementToolStatusImprovementRequired   WorkflowEventDetailResponseEvaluationHistoryImprovementToolStatus = "improvement_required"
)

func (r WorkflowEventDetailResponseEvaluationHistoryImprovementToolStatus) IsKnown() bool {
	switch r {
	case WorkflowEventDetailResponseEvaluationHistoryImprovementToolStatusImproved, WorkflowEventDetailResponseEvaluationHistoryImprovementToolStatusImprovementFailed, WorkflowEventDetailResponseEvaluationHistoryImprovementToolStatusNoImprovementRequired, WorkflowEventDetailResponseEvaluationHistoryImprovementToolStatusImprovementRequired:
		return true
	}
	return false
}

// Type of improvement action used to improve the event.
type WorkflowEventDetailResponseImprovementAction string

const (
	WorkflowEventDetailResponseImprovementActionRegen     WorkflowEventDetailResponseImprovementAction = "regen"
	WorkflowEventDetailResponseImprovementActionFixit     WorkflowEventDetailResponseImprovementAction = "fixit"
	WorkflowEventDetailResponseImprovementActionDoNothing WorkflowEventDetailResponseImprovementAction = "do_nothing"
)

func (r WorkflowEventDetailResponseImprovementAction) IsKnown() bool {
	switch r {
	case WorkflowEventDetailResponseImprovementActionRegen, WorkflowEventDetailResponseImprovementActionFixit, WorkflowEventDetailResponseImprovementActionDoNothing:
		return true
	}
	return false
}

// Status of the improvement tool used to improve the event. `improvement_required`
// indicates that the evaluation is complete and the improvement action is needed
// but is not taking place. `improved` and `improvement_failed` indicate when the
// improvement action concludes, successfully and unsuccessfully, respectively.
// `no_improvement_required` means that the first evaluation passed all its
// metrics!
type WorkflowEventDetailResponseImprovementToolStatus string

const (
	WorkflowEventDetailResponseImprovementToolStatusImproved              WorkflowEventDetailResponseImprovementToolStatus = "improved"
	WorkflowEventDetailResponseImprovementToolStatusImprovementFailed     WorkflowEventDetailResponseImprovementToolStatus = "improvement_failed"
	WorkflowEventDetailResponseImprovementToolStatusNoImprovementRequired WorkflowEventDetailResponseImprovementToolStatus = "no_improvement_required"
	WorkflowEventDetailResponseImprovementToolStatusImprovementRequired   WorkflowEventDetailResponseImprovementToolStatus = "improvement_required"
)

func (r WorkflowEventDetailResponseImprovementToolStatus) IsKnown() bool {
	switch r {
	case WorkflowEventDetailResponseImprovementToolStatusImproved, WorkflowEventDetailResponseImprovementToolStatusImprovementFailed, WorkflowEventDetailResponseImprovementToolStatusNoImprovementRequired, WorkflowEventDetailResponseImprovementToolStatusImprovementRequired:
		return true
	}
	return false
}

// Status of the event.
type WorkflowEventDetailResponseStatus string

const (
	WorkflowEventDetailResponseStatusInProgress WorkflowEventDetailResponseStatus = "In Progress"
	WorkflowEventDetailResponseStatusCompleted  WorkflowEventDetailResponseStatus = "Completed"
)

func (r WorkflowEventDetailResponseStatus) IsKnown() bool {
	switch r {
	case WorkflowEventDetailResponseStatusInProgress, WorkflowEventDetailResponseStatusCompleted:
		return true
	}
	return false
}

// Type of thresholds used to evaluate the event.
type WorkflowEventDetailResponseThresholdType string

const (
	WorkflowEventDetailResponseThresholdTypeCustom    WorkflowEventDetailResponseThresholdType = "custom"
	WorkflowEventDetailResponseThresholdTypeAutomatic WorkflowEventDetailResponseThresholdType = "automatic"
)

func (r WorkflowEventDetailResponseThresholdType) IsKnown() bool {
	switch r {
	case WorkflowEventDetailResponseThresholdTypeCustom, WorkflowEventDetailResponseThresholdTypeAutomatic:
		return true
	}
	return false
}

type WorkflowEventDetailResponseAutomaticHallucinationToleranceLevel string

const (
	WorkflowEventDetailResponseAutomaticHallucinationToleranceLevelLow    WorkflowEventDetailResponseAutomaticHallucinationToleranceLevel = "low"
	WorkflowEventDetailResponseAutomaticHallucinationToleranceLevelMedium WorkflowEventDetailResponseAutomaticHallucinationToleranceLevel = "medium"
	WorkflowEventDetailResponseAutomaticHallucinationToleranceLevelHigh   WorkflowEventDetailResponseAutomaticHallucinationToleranceLevel = "high"
)

func (r WorkflowEventDetailResponseAutomaticHallucinationToleranceLevel) IsKnown() bool {
	switch r {
	case WorkflowEventDetailResponseAutomaticHallucinationToleranceLevelLow, WorkflowEventDetailResponseAutomaticHallucinationToleranceLevelMedium, WorkflowEventDetailResponseAutomaticHallucinationToleranceLevelHigh:
		return true
	}
	return false
}

type WorkflowEventDetailResponseCapability struct {
	Capability string                                    `json:"capability"`
	JSON       workflowEventDetailResponseCapabilityJSON `json:"-"`
}

// workflowEventDetailResponseCapabilityJSON contains the JSON metadata for the
// struct [WorkflowEventDetailResponseCapability]
type workflowEventDetailResponseCapabilityJSON struct {
	Capability  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowEventDetailResponseCapability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowEventDetailResponseCapabilityJSON) RawJSON() string {
	return r.raw
}

type WorkflowEventDetailResponseFile struct {
	FileID                string                              `json:"file_id"`
	FileName              string                              `json:"file_name"`
	FileSize              int64                               `json:"file_size"`
	PresignedURL          string                              `json:"presigned_url"`
	PresignedURLExpiresAt time.Time                           `json:"presigned_url_expires_at" format:"date-time"`
	JSON                  workflowEventDetailResponseFileJSON `json:"-"`
}

// workflowEventDetailResponseFileJSON contains the JSON metadata for the struct
// [WorkflowEventDetailResponseFile]
type workflowEventDetailResponseFileJSON struct {
	FileID                apijson.Field
	FileName              apijson.Field
	FileSize              apijson.Field
	PresignedURL          apijson.Field
	PresignedURLExpiresAt apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *WorkflowEventDetailResponseFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowEventDetailResponseFileJSON) RawJSON() string {
	return r.raw
}

type WorkflowEventResponse struct {
	// The ID of the billing request for the event.
	BillingRequestID string `json:"billing_request_id" api:"required"`
	// The time the event was created in UTC.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A unique workflow event ID.
	EventID string `json:"event_id" api:"required"`
	// Status of the event.
	Status WorkflowEventResponseStatus `json:"status" api:"required"`
	// Workflow ID associated with the event.
	WorkflowID string                    `json:"workflow_id" api:"required"`
	JSON       workflowEventResponseJSON `json:"-"`
}

// workflowEventResponseJSON contains the JSON metadata for the struct
// [WorkflowEventResponse]
type workflowEventResponseJSON struct {
	BillingRequestID apijson.Field
	CreatedAt        apijson.Field
	EventID          apijson.Field
	Status           apijson.Field
	WorkflowID       apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *WorkflowEventResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowEventResponseJSON) RawJSON() string {
	return r.raw
}

// Status of the event.
type WorkflowEventResponseStatus string

const (
	WorkflowEventResponseStatusInProgress WorkflowEventResponseStatus = "In Progress"
	WorkflowEventResponseStatusCompleted  WorkflowEventResponseStatus = "Completed"
)

func (r WorkflowEventResponseStatus) IsKnown() bool {
	switch r {
	case WorkflowEventResponseStatusInProgress, WorkflowEventResponseStatusCompleted:
		return true
	}
	return false
}

type DefendNewWorkflowParams struct {
	// The action used to improve outputs that fail one or more guardrail metrics for
	// the workflow events. May be `regen`, `fixit`, or `do_nothing`. ReGen runs the
	// user's input prompt with minor induced variance. FixIt attempts to directly
	// address the shortcomings of the output using the guardrail failure rationale. Do
	// Nothing does not attempt any improvement.
	ImprovementAction param.Field[DefendNewWorkflowParamsImprovementAction] `json:"improvement_action" api:"required"`
	// Name of the workflow.
	Name param.Field[string] `json:"name" api:"required"`
	// Type of thresholds to use for the workflow, either `automatic` or `custom`.
	// Automatic thresholds are assigned internally after the user specifies a
	// qualitative tolerance for the metrics, whereas custom metrics allow the user to
	// set the threshold for each metric as a floating point number between 0.0 and
	// 1.0.
	ThresholdType param.Field[DefendNewWorkflowParamsThresholdType] `json:"threshold_type" api:"required"`
	// Mapping of guardrail metrics to hallucination tolerance levels (either `low`,
	// `medium`, or `high`). Possible metrics are `completeness`,
	// `instruction_adherence`, `context_adherence`, `ground_truth_adherence`, or
	// `comprehensive_safety`.
	AutomaticHallucinationToleranceLevels param.Field[map[string]DefendNewWorkflowParamsAutomaticHallucinationToleranceLevels] `json:"automatic_hallucination_tolerance_levels"`
	// Context includes any structured information that directly relates to the model’s
	// input and expected output—e.g., the recent turn-by-turn history between an AI
	// tutor and a student, facts or state passed through an agentic workflow, or other
	// domain-specific signals your system already knows and wants the model to
	// condition on. This field determines whether to enable context awareness for this
	// workflow's evaluations. Defaults to false.
	ContextAwareness param.Field[bool] `json:"context_awareness"`
	// Mapping of guardrail metrics to floating point threshold values. Possible
	// metrics are `correctness`, `completeness`, `instruction_adherence`,
	// `context_adherence`, `ground_truth_adherence`, or `comprehensive_safety`.
	CustomHallucinationThresholdValues param.Field[map[string]float64] `json:"custom_hallucination_threshold_values"`
	// Description for the workflow.
	Description param.Field[string] `json:"description"`
	// An array of file IDs to search in the workflow's evaluations. Files must be
	// uploaded via the DeepRails API first.
	FileSearch param.Field[[]string] `json:"file_search"`
	// Max. number of improvement action attempts until a given event passes the
	// guardrails. Defaults to 10.
	MaxImprovementAttempts param.Field[int64] `json:"max_improvement_attempts"`
	// Whether to enable web search for this workflow's evaluations. Defaults to false.
	WebSearch param.Field[bool] `json:"web_search"`
}

func (r DefendNewWorkflowParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action used to improve outputs that fail one or more guardrail metrics for
// the workflow events. May be `regen`, `fixit`, or `do_nothing`. ReGen runs the
// user's input prompt with minor induced variance. FixIt attempts to directly
// address the shortcomings of the output using the guardrail failure rationale. Do
// Nothing does not attempt any improvement.
type DefendNewWorkflowParamsImprovementAction string

const (
	DefendNewWorkflowParamsImprovementActionRegen     DefendNewWorkflowParamsImprovementAction = "regen"
	DefendNewWorkflowParamsImprovementActionFixit     DefendNewWorkflowParamsImprovementAction = "fixit"
	DefendNewWorkflowParamsImprovementActionDoNothing DefendNewWorkflowParamsImprovementAction = "do_nothing"
)

func (r DefendNewWorkflowParamsImprovementAction) IsKnown() bool {
	switch r {
	case DefendNewWorkflowParamsImprovementActionRegen, DefendNewWorkflowParamsImprovementActionFixit, DefendNewWorkflowParamsImprovementActionDoNothing:
		return true
	}
	return false
}

// Type of thresholds to use for the workflow, either `automatic` or `custom`.
// Automatic thresholds are assigned internally after the user specifies a
// qualitative tolerance for the metrics, whereas custom metrics allow the user to
// set the threshold for each metric as a floating point number between 0.0 and
// 1.0.
type DefendNewWorkflowParamsThresholdType string

const (
	DefendNewWorkflowParamsThresholdTypeAutomatic DefendNewWorkflowParamsThresholdType = "automatic"
	DefendNewWorkflowParamsThresholdTypeCustom    DefendNewWorkflowParamsThresholdType = "custom"
)

func (r DefendNewWorkflowParamsThresholdType) IsKnown() bool {
	switch r {
	case DefendNewWorkflowParamsThresholdTypeAutomatic, DefendNewWorkflowParamsThresholdTypeCustom:
		return true
	}
	return false
}

type DefendNewWorkflowParamsAutomaticHallucinationToleranceLevels string

const (
	DefendNewWorkflowParamsAutomaticHallucinationToleranceLevelsLow    DefendNewWorkflowParamsAutomaticHallucinationToleranceLevels = "low"
	DefendNewWorkflowParamsAutomaticHallucinationToleranceLevelsMedium DefendNewWorkflowParamsAutomaticHallucinationToleranceLevels = "medium"
	DefendNewWorkflowParamsAutomaticHallucinationToleranceLevelsHigh   DefendNewWorkflowParamsAutomaticHallucinationToleranceLevels = "high"
)

func (r DefendNewWorkflowParamsAutomaticHallucinationToleranceLevels) IsKnown() bool {
	switch r {
	case DefendNewWorkflowParamsAutomaticHallucinationToleranceLevelsLow, DefendNewWorkflowParamsAutomaticHallucinationToleranceLevelsMedium, DefendNewWorkflowParamsAutomaticHallucinationToleranceLevelsHigh:
		return true
	}
	return false
}

type DefendGetWorkflowParams struct {
	// Limit the number of returned events associated with this workflow. Defaults
	// to 10.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [DefendGetWorkflowParams]'s query parameters as
// `url.Values`.
func (r DefendGetWorkflowParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DefendSubmitAndStreamEventParams struct {
	// The input provided to the model (e.g., prompt, messages).
	ModelInput param.Field[map[string]interface{}] `json:"model_input" api:"required"`
	// The output generated by the model to be evaluated.
	ModelOutput param.Field[string] `json:"model_output" api:"required"`
	// The model that generated the output (e.g., "gpt-4", "claude-3").
	ModelUsed param.Field[string] `json:"model_used" api:"required"`
	// The evaluation run mode. Streaming is supported on all run modes except
	// precision_max and precision_max_codex. Note: super_fast does not support Web
	// Search or File Search — if your workflow has these enabled, use a different run
	// mode or disable the capability on the workflow.
	RunMode param.Field[DefendSubmitAndStreamEventParamsRunMode] `json:"run_mode" api:"required"`
	// Enable SSE streaming for real-time token feedback. Supported on all run modes
	// except precision_max and precision_max_codex.
	Stream param.Field[bool] `query:"stream"`
	// Optional tag to identify this event.
	Nametag param.Field[string] `json:"nametag"`
}

func (r DefendSubmitAndStreamEventParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [DefendSubmitAndStreamEventParams]'s query parameters as
// `url.Values`.
func (r DefendSubmitAndStreamEventParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The evaluation run mode. Streaming is supported on all run modes except
// precision_max and precision_max_codex. Note: super_fast does not support Web
// Search or File Search — if your workflow has these enabled, use a different run
// mode or disable the capability on the workflow.
type DefendSubmitAndStreamEventParamsRunMode string

const (
	DefendSubmitAndStreamEventParamsRunModeSuperFast      DefendSubmitAndStreamEventParamsRunMode = "super_fast"
	DefendSubmitAndStreamEventParamsRunModeFast           DefendSubmitAndStreamEventParamsRunMode = "fast"
	DefendSubmitAndStreamEventParamsRunModePrecision      DefendSubmitAndStreamEventParamsRunMode = "precision"
	DefendSubmitAndStreamEventParamsRunModePrecisionCodex DefendSubmitAndStreamEventParamsRunMode = "precision_codex"
)

func (r DefendSubmitAndStreamEventParamsRunMode) IsKnown() bool {
	switch r {
	case DefendSubmitAndStreamEventParamsRunModeSuperFast, DefendSubmitAndStreamEventParamsRunModeFast, DefendSubmitAndStreamEventParamsRunModePrecision, DefendSubmitAndStreamEventParamsRunModePrecisionCodex:
		return true
	}
	return false
}

type DefendSubmitEventParams struct {
	// A dictionary of inputs sent to the LLM to generate output. The dictionary must
	// contain a `user_prompt` field. For the ground_truth_adherence guardrail metric,
	// `ground_truth` should be provided.
	ModelInput param.Field[DefendSubmitEventParamsModelInput] `json:"model_input" api:"required"`
	// Output generated by the LLM to be evaluated.
	ModelOutput param.Field[string] `json:"model_output" api:"required"`
	// Model ID used to generate the output, like `gpt-4o` or `o3`.
	ModelUsed param.Field[string] `json:"model_used" api:"required"`
	// Run mode for the workflow event. The run mode allows the user to optimize for
	// speed, accuracy, and cost by determining which models are used to evaluate the
	// event. Available run modes (fastest to most thorough): `super_fast`, `fast`,
	// `precision`, `precision_codex`, `precision_max`, and `precision_max_codex`.
	// Defaults to `fast`. Note: `super_fast` does not support Web Search or File
	// Search — if your workflow has these capabilities enabled, use a different run
	// mode or edit the workflow to disable them.
	RunMode param.Field[DefendSubmitEventParamsRunMode] `json:"run_mode" api:"required"`
	// An optional, user-defined tag for the event.
	Nametag param.Field[string] `json:"nametag"`
}

func (r DefendSubmitEventParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A dictionary of inputs sent to the LLM to generate output. The dictionary must
// contain a `user_prompt` field. For the ground_truth_adherence guardrail metric,
// `ground_truth` should be provided.
type DefendSubmitEventParamsModelInput struct {
	// The user prompt used to generate the output.
	UserPrompt param.Field[string] `json:"user_prompt" api:"required"`
	// Any structured information that directly relates to the model’s input and
	// expected output—e.g., the recent turn-by-turn history between an AI tutor and a
	// student, facts or state passed through an agentic workflow, or other
	// domain-specific signals your system already knows and wants the model to
	// condition on.
	Context param.Field[[]DefendSubmitEventParamsModelInputContext] `json:"context"`
	// The ground truth for evaluating the Ground Truth Adherence guardrail.
	GroundTruth param.Field[string] `json:"ground_truth"`
	// The system prompt used to generate the output.
	SystemPrompt param.Field[string] `json:"system_prompt"`
}

func (r DefendSubmitEventParamsModelInput) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DefendSubmitEventParamsModelInputContext struct {
	// The content of the message.
	Content param.Field[string] `json:"content"`
	// The role of the speaker.
	Role param.Field[string] `json:"role"`
}

func (r DefendSubmitEventParamsModelInputContext) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Run mode for the workflow event. The run mode allows the user to optimize for
// speed, accuracy, and cost by determining which models are used to evaluate the
// event. Available run modes (fastest to most thorough): `super_fast`, `fast`,
// `precision`, `precision_codex`, `precision_max`, and `precision_max_codex`.
// Defaults to `fast`. Note: `super_fast` does not support Web Search or File
// Search — if your workflow has these capabilities enabled, use a different run
// mode or edit the workflow to disable them.
type DefendSubmitEventParamsRunMode string

const (
	DefendSubmitEventParamsRunModeSuperFast         DefendSubmitEventParamsRunMode = "super_fast"
	DefendSubmitEventParamsRunModeFast              DefendSubmitEventParamsRunMode = "fast"
	DefendSubmitEventParamsRunModePrecision         DefendSubmitEventParamsRunMode = "precision"
	DefendSubmitEventParamsRunModePrecisionCodex    DefendSubmitEventParamsRunMode = "precision_codex"
	DefendSubmitEventParamsRunModePrecisionMax      DefendSubmitEventParamsRunMode = "precision_max"
	DefendSubmitEventParamsRunModePrecisionMaxCodex DefendSubmitEventParamsRunMode = "precision_max_codex"
)

func (r DefendSubmitEventParamsRunMode) IsKnown() bool {
	switch r {
	case DefendSubmitEventParamsRunModeSuperFast, DefendSubmitEventParamsRunModeFast, DefendSubmitEventParamsRunModePrecision, DefendSubmitEventParamsRunModePrecisionCodex, DefendSubmitEventParamsRunModePrecisionMax, DefendSubmitEventParamsRunModePrecisionMaxCodex:
		return true
	}
	return false
}

type DefendUpdateWorkflowParams struct {
	// New mapping of guardrail metrics to hallucination tolerance levels (either
	// `low`, `medium`, or `high`) to be used when `threshold_type` is set to
	// `automatic`. Possible metrics are `completeness`, `instruction_adherence`,
	// `context_adherence`, `ground_truth_adherence`, or `comprehensive_safety`.
	AutomaticHallucinationToleranceLevels param.Field[map[string]DefendUpdateWorkflowParamsAutomaticHallucinationToleranceLevels] `json:"automatic_hallucination_tolerance_levels"`
	// Whether to enable context awareness for this workflow's evaluations.
	ContextAwareness param.Field[bool] `json:"context_awareness"`
	// New mapping of guardrail metrics to floating point threshold values to be used
	// when `threshold_type` is set to `custom`. Possible metrics are `correctness`,
	// `completeness`, `instruction_adherence`, `context_adherence`,
	// `ground_truth_adherence`, or `comprehensive_safety`.
	CustomHallucinationThresholdValues param.Field[map[string]float64] `json:"custom_hallucination_threshold_values"`
	// New description for the workflow.
	Description param.Field[string] `json:"description"`
	// An array of file IDs to search in the workflow's evaluations. Files must be
	// uploaded via the DeepRails API first.
	FileSearch param.Field[[]string] `json:"file_search"`
	// The new action used to improve outputs that fail one or more guardrail metrics
	// for the workflow events. May be `regen`, `fixit`, or `do_nothing`. ReGen runs
	// the user's input prompt with minor induced variance. FixIt attempts to directly
	// address the shortcomings of the output using the guardrail failure rationale. Do
	// Nothing does not attempt any improvement.
	ImprovementAction param.Field[DefendUpdateWorkflowParamsImprovementAction] `json:"improvement_action"`
	// Max. number of improvement action attempts until a given event passes the
	// guardrails. Defaults to 10.
	MaxImprovementAttempts param.Field[int64] `json:"max_improvement_attempts"`
	// New name for the workflow.
	Name param.Field[string] `json:"name"`
	// New type of thresholds to use for the workflow, either `automatic` or `custom`.
	// Automatic thresholds are assigned internally after the user specifies a
	// qualitative tolerance for the metrics, whereas custom metrics allow the user to
	// set the threshold for each metric as a floating point number between 0.0 and
	// 1.0.
	ThresholdType param.Field[DefendUpdateWorkflowParamsThresholdType] `json:"threshold_type"`
	// Whether to enable web search for this workflow's evaluations.
	WebSearch param.Field[bool] `json:"web_search"`
}

func (r DefendUpdateWorkflowParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DefendUpdateWorkflowParamsAutomaticHallucinationToleranceLevels string

const (
	DefendUpdateWorkflowParamsAutomaticHallucinationToleranceLevelsLow    DefendUpdateWorkflowParamsAutomaticHallucinationToleranceLevels = "low"
	DefendUpdateWorkflowParamsAutomaticHallucinationToleranceLevelsMedium DefendUpdateWorkflowParamsAutomaticHallucinationToleranceLevels = "medium"
	DefendUpdateWorkflowParamsAutomaticHallucinationToleranceLevelsHigh   DefendUpdateWorkflowParamsAutomaticHallucinationToleranceLevels = "high"
)

func (r DefendUpdateWorkflowParamsAutomaticHallucinationToleranceLevels) IsKnown() bool {
	switch r {
	case DefendUpdateWorkflowParamsAutomaticHallucinationToleranceLevelsLow, DefendUpdateWorkflowParamsAutomaticHallucinationToleranceLevelsMedium, DefendUpdateWorkflowParamsAutomaticHallucinationToleranceLevelsHigh:
		return true
	}
	return false
}

// The new action used to improve outputs that fail one or more guardrail metrics
// for the workflow events. May be `regen`, `fixit`, or `do_nothing`. ReGen runs
// the user's input prompt with minor induced variance. FixIt attempts to directly
// address the shortcomings of the output using the guardrail failure rationale. Do
// Nothing does not attempt any improvement.
type DefendUpdateWorkflowParamsImprovementAction string

const (
	DefendUpdateWorkflowParamsImprovementActionRegen     DefendUpdateWorkflowParamsImprovementAction = "regen"
	DefendUpdateWorkflowParamsImprovementActionFixit     DefendUpdateWorkflowParamsImprovementAction = "fixit"
	DefendUpdateWorkflowParamsImprovementActionDoNothing DefendUpdateWorkflowParamsImprovementAction = "do_nothing"
)

func (r DefendUpdateWorkflowParamsImprovementAction) IsKnown() bool {
	switch r {
	case DefendUpdateWorkflowParamsImprovementActionRegen, DefendUpdateWorkflowParamsImprovementActionFixit, DefendUpdateWorkflowParamsImprovementActionDoNothing:
		return true
	}
	return false
}

// New type of thresholds to use for the workflow, either `automatic` or `custom`.
// Automatic thresholds are assigned internally after the user specifies a
// qualitative tolerance for the metrics, whereas custom metrics allow the user to
// set the threshold for each metric as a floating point number between 0.0 and
// 1.0.
type DefendUpdateWorkflowParamsThresholdType string

const (
	DefendUpdateWorkflowParamsThresholdTypeAutomatic DefendUpdateWorkflowParamsThresholdType = "automatic"
	DefendUpdateWorkflowParamsThresholdTypeCustom    DefendUpdateWorkflowParamsThresholdType = "custom"
)

func (r DefendUpdateWorkflowParamsThresholdType) IsKnown() bool {
	switch r {
	case DefendUpdateWorkflowParamsThresholdTypeAutomatic, DefendUpdateWorkflowParamsThresholdTypeCustom:
		return true
	}
	return false
}
