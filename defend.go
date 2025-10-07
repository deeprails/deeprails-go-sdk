// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomdeeprailsdeeprailsgosdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/deeprails/deeprails-go-sdk/internal/apijson"
	"github.com/deeprails/deeprails-go-sdk/internal/requestconfig"
	"github.com/deeprails/deeprails-go-sdk/option"
	"github.com/deeprails/deeprails-go-sdk/packages/param"
	"github.com/deeprails/deeprails-go-sdk/packages/respjson"
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
func NewDefendService(opts ...option.RequestOption) (r DefendService) {
	r = DefendService{}
	r.Options = opts
	return
}

// Create a new guardrail workflow with optional guardrail thresholds and
// improvement actions.
func (r *DefendService) NewWorkflow(ctx context.Context, body DefendNewWorkflowParams, opts ...option.RequestOption) (res *DefendResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "defend"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a specific event of a guardrail workflow.
func (r *DefendService) GetEvent(ctx context.Context, eventID string, query DefendGetEventParams, opts ...option.RequestOption) (res *WorkflowEventResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.WorkflowID == "" {
		err = errors.New("missing required workflow_id parameter")
		return
	}
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	path := fmt.Sprintf("defend/%s/events/%s", query.WorkflowID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve the details for a specific guardrail workflow.
func (r *DefendService) GetWorkflow(ctx context.Context, workflowID string, opts ...option.RequestOption) (res *DefendResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if workflowID == "" {
		err = errors.New("missing required workflow_id parameter")
		return
	}
	path := fmt.Sprintf("defend/%s", workflowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Submit a model input and output pair to a workflow for evaluation.
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

// Update an existing guardrail workflow.
func (r *DefendService) UpdateWorkflow(ctx context.Context, workflowID string, body DefendUpdateWorkflowParams, opts ...option.RequestOption) (res *DefendResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if workflowID == "" {
		err = errors.New("missing required workflow_id parameter")
		return
	}
	path := fmt.Sprintf("defend/%s", workflowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Response payload for guardrail workflow operations.
type DefendResponse struct {
	// Name of the workflow.
	Name string `json:"name,required"`
	// A unique workflow ID.
	WorkflowID string `json:"workflow_id,required"`
	// The time the workflow was created in UTC.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Description for the workflow.
	Description string `json:"description"`
	// The action used to improve outputs that fail one or more guardrail metrics for
	// the workflow events. May be `regenerate`, `fixit`, or null which represents “do
	// nothing”. ReGen runs the user's exact input prompt with minor induced variance.
	// Fixit attempts to directly address the shortcomings of the output using the
	// guardrail failure rationale. Do nothing does not attempt any improvement.
	//
	// Any of "regenerate", "fixit".
	ImprovementAction DefendResponseImprovementAction `json:"improvement_action,nullable"`
	// Max. number of improvement action retries until a given event passes the
	// guardrails.
	MaxRetries int64 `json:"max_retries"`
	// The most recent time the workflow was modified in UTC.
	ModifiedAt time.Time `json:"modified_at" format:"date-time"`
	// Status of the selected workflow. May be `archived` or `active`. Archived
	// workflows will not accept events.
	//
	// Any of "archived", "active".
	Status DefendResponseStatus `json:"status"`
	// Rate of events associated with this workflow that passed evaluation.
	SuccessRate float64 `json:"success_rate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name              respjson.Field
		WorkflowID        respjson.Field
		CreatedAt         respjson.Field
		Description       respjson.Field
		ImprovementAction respjson.Field
		MaxRetries        respjson.Field
		ModifiedAt        respjson.Field
		Status            respjson.Field
		SuccessRate       respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DefendResponse) RawJSON() string { return r.JSON.raw }
func (r *DefendResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The action used to improve outputs that fail one or more guardrail metrics for
// the workflow events. May be `regenerate`, `fixit`, or null which represents “do
// nothing”. ReGen runs the user's exact input prompt with minor induced variance.
// Fixit attempts to directly address the shortcomings of the output using the
// guardrail failure rationale. Do nothing does not attempt any improvement.
type DefendResponseImprovementAction string

const (
	DefendResponseImprovementActionRegenerate DefendResponseImprovementAction = "regenerate"
	DefendResponseImprovementActionFixit      DefendResponseImprovementAction = "fixit"
)

// Status of the selected workflow. May be `archived` or `active`. Archived
// workflows will not accept events.
type DefendResponseStatus string

const (
	DefendResponseStatusArchived DefendResponseStatus = "archived"
	DefendResponseStatusActive   DefendResponseStatus = "active"
)

// Response payload for workflow event operations.
type WorkflowEventResponse struct {
	// A unique workflow event ID.
	EventID string `json:"event_id,required"`
	// Workflow ID associated with the event.
	WorkflowID string `json:"workflow_id,required"`
	// Count of improvement attempts for the event. If greater than one then all
	// previous improvement attempts failed.
	AttemptNumber int64 `json:"attempt_number"`
	// A unique evaluation ID associated with this event. Every event has one or more
	// evaluation attempts.
	EvaluationID string `json:"evaluation_id"`
	// `False` if evaluation passed all of the guardrail metrics, `True` if evaluation
	// failed any of the guardrail metrics.
	Filtered bool `json:"filtered"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventID       respjson.Field
		WorkflowID    respjson.Field
		AttemptNumber respjson.Field
		EvaluationID  respjson.Field
		Filtered      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowEventResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowEventResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DefendNewWorkflowParams struct {
	// The action used to improve outputs that fail one or guardrail metrics for the
	// workflow events. May be `regenerate`, `fixit`, or null which represents “do
	// nothing”. ReGen runs the user's exact input prompt with minor induced variance.
	// Fixit attempts to directly address the shortcomings of the output using the
	// guardrail failure rationale. Do nothing does not attempt any improvement.
	//
	// Any of "regenerate", "fixit".
	ImprovementAction DefendNewWorkflowParamsImprovementAction `json:"improvement_action,omitzero,required"`
	// Mapping of guardrail metrics to floating point threshold values. If the workflow
	// type is automatic, only the metric names are used (`automatic_tolerance`
	// determines thresholds). Possible metrics are `correctness`, `completeness`,
	// `instruction_adherence`, `context_adherence`, `ground_truth_adherence`, or
	// `comprehensive_safety`.
	Metrics map[string]float64 `json:"metrics,omitzero,required"`
	// Name of the workflow.
	Name string `json:"name,required"`
	// Type of thresholds to use for the workflow, either `automatic` or `custom`.
	// Automatic thresholds are assigned internally after the user specifies a
	// qualitative tolerance for the metrics, whereas custom metrics allow the user to
	// set the threshold for each metric as a floating point number between 0.0 and
	// 1.0.
	//
	// Any of "automatic", "custom".
	Type DefendNewWorkflowParamsType `json:"type,omitzero,required"`
	// Description for the workflow.
	Description param.Opt[string] `json:"description,omitzero"`
	// Max. number of improvement action retries until a given event passes the
	// guardrails. Defaults to 10.
	MaxRetries param.Opt[int64] `json:"max_retries,omitzero"`
	// Hallucination tolerance for automatic workflows; may be `low`, `medium`, or
	// `high`. Ignored if `type` is `custom`.
	//
	// Any of "low", "medium", "high".
	AutomaticTolerance DefendNewWorkflowParamsAutomaticTolerance `json:"automatic_tolerance,omitzero"`
	paramObj
}

func (r DefendNewWorkflowParams) MarshalJSON() (data []byte, err error) {
	type shadow DefendNewWorkflowParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DefendNewWorkflowParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The action used to improve outputs that fail one or guardrail metrics for the
// workflow events. May be `regenerate`, `fixit`, or null which represents “do
// nothing”. ReGen runs the user's exact input prompt with minor induced variance.
// Fixit attempts to directly address the shortcomings of the output using the
// guardrail failure rationale. Do nothing does not attempt any improvement.
type DefendNewWorkflowParamsImprovementAction string

const (
	DefendNewWorkflowParamsImprovementActionRegenerate DefendNewWorkflowParamsImprovementAction = "regenerate"
	DefendNewWorkflowParamsImprovementActionFixit      DefendNewWorkflowParamsImprovementAction = "fixit"
)

// Type of thresholds to use for the workflow, either `automatic` or `custom`.
// Automatic thresholds are assigned internally after the user specifies a
// qualitative tolerance for the metrics, whereas custom metrics allow the user to
// set the threshold for each metric as a floating point number between 0.0 and
// 1.0.
type DefendNewWorkflowParamsType string

const (
	DefendNewWorkflowParamsTypeAutomatic DefendNewWorkflowParamsType = "automatic"
	DefendNewWorkflowParamsTypeCustom    DefendNewWorkflowParamsType = "custom"
)

// Hallucination tolerance for automatic workflows; may be `low`, `medium`, or
// `high`. Ignored if `type` is `custom`.
type DefendNewWorkflowParamsAutomaticTolerance string

const (
	DefendNewWorkflowParamsAutomaticToleranceLow    DefendNewWorkflowParamsAutomaticTolerance = "low"
	DefendNewWorkflowParamsAutomaticToleranceMedium DefendNewWorkflowParamsAutomaticTolerance = "medium"
	DefendNewWorkflowParamsAutomaticToleranceHigh   DefendNewWorkflowParamsAutomaticTolerance = "high"
)

type DefendGetEventParams struct {
	WorkflowID string `path:"workflow_id,required" json:"-"`
	paramObj
}

type DefendSubmitEventParams struct {
	// A dictionary of inputs sent to the LLM to generate output. This must contain a
	// `user_prompt` field and an optional `context` field. Additional properties are
	// allowed.
	ModelInput DefendSubmitEventParamsModelInput `json:"model_input,omitzero,required"`
	// Output generated by the LLM to be evaluated.
	ModelOutput string `json:"model_output,required"`
	// Model ID used to generate the output, like `gpt-4o` or `o3`.
	ModelUsed string `json:"model_used,required"`
	// An optional, user-defined tag for the event.
	Nametag string `json:"nametag,required"`
	// Run mode for the workflow event. The run mode allows the user to optimize for
	// speed, accuracy, and cost by determining which models are used to evaluate the
	// event. Available run modes include `precision_plus`, `precision`, `smart`, and
	// `economy`. Defaults to `smart`.
	//
	// Any of "precision_plus", "precision", "smart", "economy".
	RunMode DefendSubmitEventParamsRunMode `json:"run_mode,omitzero,required"`
	paramObj
}

func (r DefendSubmitEventParams) MarshalJSON() (data []byte, err error) {
	type shadow DefendSubmitEventParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DefendSubmitEventParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A dictionary of inputs sent to the LLM to generate output. This must contain a
// `user_prompt` field and an optional `context` field. Additional properties are
// allowed.
//
// The property UserPrompt is required.
type DefendSubmitEventParamsModelInput struct {
	UserPrompt  string            `json:"user_prompt,required"`
	Context     param.Opt[string] `json:"context,omitzero"`
	ExtraFields map[string]any    `json:"-"`
	paramObj
}

func (r DefendSubmitEventParamsModelInput) MarshalJSON() (data []byte, err error) {
	type shadow DefendSubmitEventParamsModelInput
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *DefendSubmitEventParamsModelInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

type DefendUpdateWorkflowParams struct {
	// Description for the workflow.
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the workflow.
	Name param.Opt[string] `json:"name,omitzero"`
	// Type of thresholds to use for the workflow, either `automatic` or `custom`.
	//
	// Any of "automatic", "custom".
	Type DefendUpdateWorkflowParamsType `json:"type,omitzero"`
	paramObj
}

func (r DefendUpdateWorkflowParams) MarshalJSON() (data []byte, err error) {
	type shadow DefendUpdateWorkflowParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DefendUpdateWorkflowParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of thresholds to use for the workflow, either `automatic` or `custom`.
type DefendUpdateWorkflowParamsType string

const (
	DefendUpdateWorkflowParamsTypeAutomatic DefendUpdateWorkflowParamsType = "automatic"
	DefendUpdateWorkflowParamsTypeCustom    DefendUpdateWorkflowParamsType = "custom"
)
