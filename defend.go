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
)

// DefendService contains methods and other services that help with interacting
// with the deeprails API.
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

// Use this endpoint to create a new guardrail workflow with optional guardrail
// thresholds and improvement actions
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

// Use this endpoint to update an existing defend workflow
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
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Status of the selected workflow. May be `inactive` or `active`. Inactive
	// workflows will not accept events.
	Status DefendCreateResponseStatus `json:"status,required"`
	// A unique workflow ID.
	WorkflowID string                   `json:"workflow_id,required"`
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
	// Name of the workflow.
	Name string `json:"name,required"`
	// A unique workflow ID.
	WorkflowID string `json:"workflow_id,required"`
	// Mapping of guardrail metric names to tolerance values. Values can be strings
	// (`low`, `medium`, `high`) for automatic tolerance levels.
	AutomaticHallucinationToleranceLevels map[string]DefendResponseAutomaticHallucinationToleranceLevel `json:"automatic_hallucination_tolerance_levels"`
	// Extended AI capabilities available to the event, if any. Can be `web_search`
	// and/or `file_search`.
	Capabilities []DefendResponseCapability `json:"capabilities"`
	// The time the workflow was created in UTC.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Mapping of guardrail metric names to threshold values. Values can be floating
	// point numbers (0.0-1.0) for custom thresholds.
	CustomHallucinationThresholdValues map[string]float64 `json:"custom_hallucination_threshold_values"`
	// Description for the workflow.
	Description string `json:"description"`
	// An array of events associated with this workflow.
	Events []DefendResponseEvent `json:"events"`
	// List of files associated with the workflow. If this is not empty, models can
	// search these files when performing evaluations or remediations
	Files []DefendResponseFile `json:"files"`
	Stats DefendResponseStats  `json:"stats"`
	// Status of the selected workflow. May be `inactive` or `active`. Inactive
	// workflows will not accept events.
	Status DefendResponseStatus `json:"status"`
	// Type of thresholds used to evaluate the event.
	ThresholdType DefendResponseThresholdType `json:"threshold_type"`
	// The most recent time the workflow was updated in UTC.
	UpdatedAt time.Time          `json:"updated_at" format:"date-time"`
	JSON      defendResponseJSON `json:"-"`
}

// defendResponseJSON contains the JSON metadata for the struct [DefendResponse]
type defendResponseJSON struct {
	Name                                  apijson.Field
	WorkflowID                            apijson.Field
	AutomaticHallucinationToleranceLevels apijson.Field
	Capabilities                          apijson.Field
	CreatedAt                             apijson.Field
	CustomHallucinationThresholdValues    apijson.Field
	Description                           apijson.Field
	Events                                apijson.Field
	Files                                 apijson.Field
	Stats                                 apijson.Field
	Status                                apijson.Field
	ThresholdType                         apijson.Field
	UpdatedAt                             apijson.Field
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
	// An array of evaluations for this event.
	Evaluations []DefendResponseEventsEvaluation `json:"evaluations"`
	// A unique workflow event ID.
	EventID string `json:"event_id"`
	// Improved model output after improvement tool was applied.
	ImprovedModelOutput string `json:"improved_model_output"`
	// Status of the improvement tool used to improve the event.
	ImprovementToolStatus string                  `json:"improvement_tool_status"`
	JSON                  defendResponseEventJSON `json:"-"`
}

// defendResponseEventJSON contains the JSON metadata for the struct
// [DefendResponseEvent]
type defendResponseEventJSON struct {
	Evaluations           apijson.Field
	EventID               apijson.Field
	ImprovedModelOutput   apijson.Field
	ImprovementToolStatus apijson.Field
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
	Attempt             apijson.Field
	CreatedAt           apijson.Field
	ErrorMessage        apijson.Field
	EvaluationResult    apijson.Field
	EvaluationStatus    apijson.Field
	EvaluationTotalCost apijson.Field
	GuardrailMetrics    apijson.Field
	ModelInput          apijson.Field
	ModelOutput         apijson.Field
	ModifiedAt          apijson.Field
	Nametag             apijson.Field
	Progress            apijson.Field
	RunMode             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *DefendResponseEventsEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r defendResponseEventsEvaluationJSON) RawJSON() string {
	return r.raw
}

type DefendResponseFile struct {
	FileID   string                 `json:"file_id"`
	FileName string                 `json:"file_name"`
	FileSize int64                  `json:"file_size"`
	JSON     defendResponseFileJSON `json:"-"`
}

// defendResponseFileJSON contains the JSON metadata for the struct
// [DefendResponseFile]
type defendResponseFileJSON struct {
	FileID      apijson.Field
	FileName    apijson.Field
	FileSize    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DefendResponseFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r defendResponseFileJSON) RawJSON() string {
	return r.raw
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

type DefendUpdateResponse struct {
	// The time the workflow was last modified in UTC.
	ModifiedAt time.Time `json:"modified_at,required" format:"date-time"`
	// Status of the selected workflow. May be `inactive` or `active`. Inactive
	// workflows will not accept events.
	Status DefendUpdateResponseStatus `json:"status,required"`
	// A unique workflow ID.
	WorkflowID string                   `json:"workflow_id,required"`
	JSON       defendUpdateResponseJSON `json:"-"`
}

// defendUpdateResponseJSON contains the JSON metadata for the struct
// [DefendUpdateResponse]
type defendUpdateResponseJSON struct {
	ModifiedAt  apijson.Field
	Status      apijson.Field
	WorkflowID  apijson.Field
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
	// A unique workflow event ID.
	EventID string `json:"event_id,required"`
	// Status of the event.
	EventStatus WorkflowEventDetailResponseEventStatus `json:"event_status,required"`
	// Whether the event was filtered and requires improvement.
	Filtered bool `json:"filtered,required"`
	// Type of improvement tool used to improve the event.
	ImprovementToolType WorkflowEventDetailResponseImprovementToolType `json:"improvement_tool_type,required"`
	// Workflow ID associated with the event.
	WorkflowID string `json:"workflow_id,required"`
	// Mapping of guardrail metric names to tolerance values. Values are strings
	// (`low`, `medium`, `high`) representing automatic tolerance levels.
	AutomaticHallucinationToleranceLevels map[string]WorkflowEventDetailResponseAutomaticHallucinationToleranceLevel `json:"automatic_hallucination_tolerance_levels"`
	// Extended AI capabilities available to the event, if any. Can be `web_search`
	// and/or `file_search`.
	Capabilities []WorkflowEventDetailResponseCapability `json:"capabilities"`
	// Mapping of guardrail metric names to threshold values. Values are floating point
	// numbers (0.0-1.0) representing custom thresholds.
	CustomHallucinationThresholdValues map[string]float64 `json:"custom_hallucination_threshold_values"`
	// History of evaluations for the event.
	EvaluationHistory []WorkflowEventDetailResponseEvaluationHistory `json:"evaluation_history"`
	// Evaluation result consisting of average scores and rationales for each of the
	// evaluated guardrail metrics.
	EvaluationResult map[string]interface{} `json:"evaluation_result"`
	// List of files available to the event, if any. Will only be present if
	// `file_search` is enabled.
	Files []WorkflowEventDetailResponseFile `json:"files"`
	// Improved model output after improvement tool was applied and each metric passed
	// evaluation.
	ImprovedModelOutput string `json:"improved_model_output"`
	// Status of the improvement tool used to improve the event.
	ImprovementToolStatus WorkflowEventDetailResponseImprovementToolStatus `json:"improvement_tool_status,nullable"`
	// Type of thresholds used to evaluate the event.
	ThresholdType WorkflowEventDetailResponseThresholdType `json:"threshold_type"`
	JSON          workflowEventDetailResponseJSON          `json:"-"`
}

// workflowEventDetailResponseJSON contains the JSON metadata for the struct
// [WorkflowEventDetailResponse]
type workflowEventDetailResponseJSON struct {
	EventID                               apijson.Field
	EventStatus                           apijson.Field
	Filtered                              apijson.Field
	ImprovementToolType                   apijson.Field
	WorkflowID                            apijson.Field
	AutomaticHallucinationToleranceLevels apijson.Field
	Capabilities                          apijson.Field
	CustomHallucinationThresholdValues    apijson.Field
	EvaluationHistory                     apijson.Field
	EvaluationResult                      apijson.Field
	Files                                 apijson.Field
	ImprovedModelOutput                   apijson.Field
	ImprovementToolStatus                 apijson.Field
	ThresholdType                         apijson.Field
	raw                                   string
	ExtraFields                           map[string]apijson.Field
}

func (r *WorkflowEventDetailResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowEventDetailResponseJSON) RawJSON() string {
	return r.raw
}

// Status of the event.
type WorkflowEventDetailResponseEventStatus string

const (
	WorkflowEventDetailResponseEventStatusInProgress WorkflowEventDetailResponseEventStatus = "In Progress"
	WorkflowEventDetailResponseEventStatusCompleted  WorkflowEventDetailResponseEventStatus = "Completed"
)

func (r WorkflowEventDetailResponseEventStatus) IsKnown() bool {
	switch r {
	case WorkflowEventDetailResponseEventStatusInProgress, WorkflowEventDetailResponseEventStatusCompleted:
		return true
	}
	return false
}

// Type of improvement tool used to improve the event.
type WorkflowEventDetailResponseImprovementToolType string

const (
	WorkflowEventDetailResponseImprovementToolTypeRegen     WorkflowEventDetailResponseImprovementToolType = "regen"
	WorkflowEventDetailResponseImprovementToolTypeFixit     WorkflowEventDetailResponseImprovementToolType = "fixit"
	WorkflowEventDetailResponseImprovementToolTypeDoNothing WorkflowEventDetailResponseImprovementToolType = "do_nothing"
)

func (r WorkflowEventDetailResponseImprovementToolType) IsKnown() bool {
	switch r {
	case WorkflowEventDetailResponseImprovementToolTypeRegen, WorkflowEventDetailResponseImprovementToolTypeFixit, WorkflowEventDetailResponseImprovementToolTypeDoNothing:
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

type WorkflowEventDetailResponseEvaluationHistory struct {
	Attempt             string                                           `json:"attempt"`
	CreatedAt           time.Time                                        `json:"created_at" format:"date-time"`
	ErrorMessage        string                                           `json:"error_message"`
	EvaluationResult    map[string]interface{}                           `json:"evaluation_result"`
	EvaluationStatus    string                                           `json:"evaluation_status"`
	EvaluationTotalCost float64                                          `json:"evaluation_total_cost"`
	GuardrailMetrics    []string                                         `json:"guardrail_metrics"`
	ModelInput          map[string]interface{}                           `json:"model_input"`
	ModelOutput         string                                           `json:"model_output"`
	ModifiedAt          time.Time                                        `json:"modified_at" format:"date-time"`
	Nametag             string                                           `json:"nametag"`
	Progress            int64                                            `json:"progress"`
	RunMode             string                                           `json:"run_mode"`
	JSON                workflowEventDetailResponseEvaluationHistoryJSON `json:"-"`
}

// workflowEventDetailResponseEvaluationHistoryJSON contains the JSON metadata for
// the struct [WorkflowEventDetailResponseEvaluationHistory]
type workflowEventDetailResponseEvaluationHistoryJSON struct {
	Attempt             apijson.Field
	CreatedAt           apijson.Field
	ErrorMessage        apijson.Field
	EvaluationResult    apijson.Field
	EvaluationStatus    apijson.Field
	EvaluationTotalCost apijson.Field
	GuardrailMetrics    apijson.Field
	ModelInput          apijson.Field
	ModelOutput         apijson.Field
	ModifiedAt          apijson.Field
	Nametag             apijson.Field
	Progress            apijson.Field
	RunMode             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *WorkflowEventDetailResponseEvaluationHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowEventDetailResponseEvaluationHistoryJSON) RawJSON() string {
	return r.raw
}

type WorkflowEventDetailResponseFile struct {
	FileID   string                              `json:"file_id"`
	FileName string                              `json:"file_name"`
	FileSize int64                               `json:"file_size"`
	JSON     workflowEventDetailResponseFileJSON `json:"-"`
}

// workflowEventDetailResponseFileJSON contains the JSON metadata for the struct
// [WorkflowEventDetailResponseFile]
type workflowEventDetailResponseFileJSON struct {
	FileID      apijson.Field
	FileName    apijson.Field
	FileSize    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowEventDetailResponseFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowEventDetailResponseFileJSON) RawJSON() string {
	return r.raw
}

// Status of the improvement tool used to improve the event.
type WorkflowEventDetailResponseImprovementToolStatus string

const (
	WorkflowEventDetailResponseImprovementToolStatusImproved            WorkflowEventDetailResponseImprovementToolStatus = "improved"
	WorkflowEventDetailResponseImprovementToolStatusFailedOnMaxRetries  WorkflowEventDetailResponseImprovementToolStatus = "failed on max retries"
	WorkflowEventDetailResponseImprovementToolStatusImprovementRequired WorkflowEventDetailResponseImprovementToolStatus = "improvement_required"
)

func (r WorkflowEventDetailResponseImprovementToolStatus) IsKnown() bool {
	switch r {
	case WorkflowEventDetailResponseImprovementToolStatusImproved, WorkflowEventDetailResponseImprovementToolStatusFailedOnMaxRetries, WorkflowEventDetailResponseImprovementToolStatusImprovementRequired:
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

type WorkflowEventResponse struct {
	// The time the event was created in UTC.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A unique workflow event ID.
	EventID string `json:"event_id,required"`
	// Status of the event.
	Status WorkflowEventResponseStatus `json:"status,required"`
	// Workflow ID associated with the event.
	WorkflowID string                    `json:"workflow_id,required"`
	JSON       workflowEventResponseJSON `json:"-"`
}

// workflowEventResponseJSON contains the JSON metadata for the struct
// [WorkflowEventResponse]
type workflowEventResponseJSON struct {
	CreatedAt   apijson.Field
	EventID     apijson.Field
	Status      apijson.Field
	WorkflowID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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
	ImprovementAction param.Field[DefendNewWorkflowParamsImprovementAction] `json:"improvement_action,required"`
	// Name of the workflow.
	Name param.Field[string] `json:"name,required"`
	// Type of thresholds to use for the workflow, either `automatic` or `custom`.
	// Automatic thresholds are assigned internally after the user specifies a
	// qualitative tolerance for the metrics, whereas custom metrics allow the user to
	// set the threshold for each metric as a floating point number between 0.0 and
	// 1.0.
	ThresholdType param.Field[DefendNewWorkflowParamsThresholdType] `json:"threshold_type,required"`
	// Mapping of guardrail metrics to hallucination tolerance levels (either `low`,
	// `medium`, or `high`). Possible metrics are `completeness`,
	// `instruction_adherence`, `context_adherence`, `ground_truth_adherence`, or
	// `comprehensive_safety`.
	AutomaticHallucinationToleranceLevels param.Field[map[string]DefendNewWorkflowParamsAutomaticHallucinationToleranceLevels] `json:"automatic_hallucination_tolerance_levels"`
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

type DefendSubmitEventParams struct {
	// A dictionary of inputs sent to the LLM to generate output. The dictionary must
	// contain at least a `user_prompt` field or a `system_prompt` field. For the
	// ground_truth_adherence guardrail metric, `ground_truth` should be provided.
	ModelInput param.Field[DefendSubmitEventParamsModelInput] `json:"model_input,required"`
	// Output generated by the LLM to be evaluated.
	ModelOutput param.Field[string] `json:"model_output,required"`
	// Model ID used to generate the output, like `gpt-4o` or `o3`.
	ModelUsed param.Field[string] `json:"model_used,required"`
	// Run mode for the workflow event. The run mode allows the user to optimize for
	// speed, accuracy, and cost by determining which models are used to evaluate the
	// event. Available run modes include `precision_plus`, `precision`, `smart`, and
	// `economy`. Defaults to `smart`.
	RunMode param.Field[DefendSubmitEventParamsRunMode] `json:"run_mode,required"`
	// An optional, user-defined tag for the event.
	Nametag param.Field[string] `json:"nametag"`
}

func (r DefendSubmitEventParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A dictionary of inputs sent to the LLM to generate output. The dictionary must
// contain at least a `user_prompt` field or a `system_prompt` field. For the
// ground_truth_adherence guardrail metric, `ground_truth` should be provided.
type DefendSubmitEventParamsModelInput struct {
	// The ground truth for evaluating the Ground Truth Adherence guardrail.
	GroundTruth param.Field[string] `json:"ground_truth"`
	// The system prompt used to generate the output.
	SystemPrompt param.Field[string] `json:"system_prompt"`
	// The user prompt used to generate the output.
	UserPrompt param.Field[string] `json:"user_prompt"`
}

func (r DefendSubmitEventParamsModelInput) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Run mode for the workflow event. The run mode allows the user to optimize for
// speed, accuracy, and cost by determining which models are used to evaluate the
// event. Available run modes include `precision_plus`, `precision`, `smart`, and
// `economy`. Defaults to `smart`.
type DefendSubmitEventParamsRunMode string

const (
	DefendSubmitEventParamsRunModePrecisionPlus DefendSubmitEventParamsRunMode = "precision_plus"
	DefendSubmitEventParamsRunModePrecision     DefendSubmitEventParamsRunMode = "precision"
	DefendSubmitEventParamsRunModeSmart         DefendSubmitEventParamsRunMode = "smart"
	DefendSubmitEventParamsRunModeEconomy       DefendSubmitEventParamsRunMode = "economy"
)

func (r DefendSubmitEventParamsRunMode) IsKnown() bool {
	switch r {
	case DefendSubmitEventParamsRunModePrecisionPlus, DefendSubmitEventParamsRunModePrecision, DefendSubmitEventParamsRunModeSmart, DefendSubmitEventParamsRunModeEconomy:
		return true
	}
	return false
}

type DefendUpdateWorkflowParams struct {
	// Description for the workflow.
	Description param.Field[string] `json:"description"`
	// Name of the workflow.
	Name param.Field[string] `json:"name"`
}

func (r DefendUpdateWorkflowParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
