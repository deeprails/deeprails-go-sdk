// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package deeprails

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/deeprails/deeprails-go-sdk/internal/apijson"
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
func (r *DefendService) NewWorkflow(ctx context.Context, body DefendNewWorkflowParams, opts ...option.RequestOption) (res *DefendResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "defend"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Use this endpoint to retrieve a specific event of a guardrail workflow
func (r *DefendService) GetEvent(ctx context.Context, workflowID string, eventID string, opts ...option.RequestOption) (res *WorkflowEventResponse, err error) {
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

// Use this endpoint to update an existing guardrail workflow
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
	// the workflow events. May be `regen`, `fixit`, or `do_nothing`. ReGen runs the
	// user's input prompt with minor induced variance. FixIt attempts to directly
	// address the shortcomings of the output using the guardrail failure rationale. Do
	// Nothing does not attempt any improvement.
	ImprovementAction DefendResponseImprovementAction `json:"improvement_action"`
	// Max. number of improvement action retries until a given event passes the
	// guardrails.
	MaxImprovementAttempts int64 `json:"max_improvement_attempts"`
	// The most recent time the workflow was modified in UTC.
	ModifiedAt time.Time `json:"modified_at" format:"date-time"`
	// Status of the selected workflow. May be `inactive` or `active`. Inactive
	// workflows will not accept events.
	Status DefendResponseStatus `json:"status"`
	// Rate of events associated with this workflow that passed evaluation.
	SuccessRate float64            `json:"success_rate"`
	JSON        defendResponseJSON `json:"-"`
}

// defendResponseJSON contains the JSON metadata for the struct [DefendResponse]
type defendResponseJSON struct {
	Name                   apijson.Field
	WorkflowID             apijson.Field
	CreatedAt              apijson.Field
	Description            apijson.Field
	ImprovementAction      apijson.Field
	MaxImprovementAttempts apijson.Field
	ModifiedAt             apijson.Field
	Status                 apijson.Field
	SuccessRate            apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *DefendResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r defendResponseJSON) RawJSON() string {
	return r.raw
}

// The action used to improve outputs that fail one or more guardrail metrics for
// the workflow events. May be `regen`, `fixit`, or `do_nothing`. ReGen runs the
// user's input prompt with minor induced variance. FixIt attempts to directly
// address the shortcomings of the output using the guardrail failure rationale. Do
// Nothing does not attempt any improvement.
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
	Filtered bool                      `json:"filtered"`
	JSON     workflowEventResponseJSON `json:"-"`
}

// workflowEventResponseJSON contains the JSON metadata for the struct
// [WorkflowEventResponse]
type workflowEventResponseJSON struct {
	EventID       apijson.Field
	WorkflowID    apijson.Field
	AttemptNumber apijson.Field
	EvaluationID  apijson.Field
	Filtered      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkflowEventResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowEventResponseJSON) RawJSON() string {
	return r.raw
}

type DefendNewWorkflowParams struct {
	// The action used to improve outputs that fail one or guardrail metrics for the
	// workflow events. May be `regen`, `fixit`, or `do_nothing`. ReGen runs the user's
	// input prompt with minor induced variance. FixIt attempts to directly address the
	// shortcomings of the output using the guardrail failure rationale. Do Nothing
	// does not attempt any improvement.
	ImprovementAction param.Field[DefendNewWorkflowParamsImprovementAction] `json:"improvement_action,required"`
	// Name of the workflow.
	Name param.Field[string] `json:"name,required"`
	// Type of thresholds to use for the workflow, either `automatic` or `custom`.
	// Automatic thresholds are assigned internally after the user specifies a
	// qualitative tolerance for the metrics, whereas custom metrics allow the user to
	// set the threshold for each metric as a floating point number between 0.0 and
	// 1.0.
	Type param.Field[DefendNewWorkflowParamsType] `json:"type,required"`
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
	// Max. number of improvement action retries until a given event passes the
	// guardrails. Defaults to 10.
	MaxImprovementAttempts param.Field[int64] `json:"max_improvement_attempts"`
	// Whether to enable web search for this workflow's evaluations. Defaults to false.
	WebSearch param.Field[bool] `json:"web_search"`
}

func (r DefendNewWorkflowParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action used to improve outputs that fail one or guardrail metrics for the
// workflow events. May be `regen`, `fixit`, or `do_nothing`. ReGen runs the user's
// input prompt with minor induced variance. FixIt attempts to directly address the
// shortcomings of the output using the guardrail failure rationale. Do Nothing
// does not attempt any improvement.
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
type DefendNewWorkflowParamsType string

const (
	DefendNewWorkflowParamsTypeAutomatic DefendNewWorkflowParamsType = "automatic"
	DefendNewWorkflowParamsTypeCustom    DefendNewWorkflowParamsType = "custom"
)

func (r DefendNewWorkflowParamsType) IsKnown() bool {
	switch r {
	case DefendNewWorkflowParamsTypeAutomatic, DefendNewWorkflowParamsTypeCustom:
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
