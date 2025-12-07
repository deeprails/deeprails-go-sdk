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
func (r *DefendService) NewWorkflow(ctx context.Context, body DefendNewWorkflowParams, opts ...option.RequestOption) (res *DefendNewWorkflowResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "defend"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Use this endpoint to retrieve a specific event of a guardrail workflow
func (r *DefendService) GetEvent(ctx context.Context, workflowID string, eventID string, opts ...option.RequestOption) (res *DefendGetEventResponse, err error) {
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
func (r *DefendService) SubmitEvent(ctx context.Context, workflowID string, body DefendSubmitEventParams, opts ...option.RequestOption) (res *DefendSubmitEventResponse, err error) {
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
func (r *DefendService) UpdateWorkflow(ctx context.Context, workflowID string, body DefendUpdateWorkflowParams, opts ...option.RequestOption) (res *DefendUpdateWorkflowResponse, err error) {
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
	// Mapping of guardrail metric names to tolerance values. Values can be strings
	// (`low`, `medium`, `high`) for automatic tolerance levels.
	AutomaticHallucinationToleranceLevels map[string]DefendResponseAutomaticHallucinationToleranceLevel `json:"automatic_hallucination_tolerance_levels"`
	// The time the workflow was created in UTC.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Mapping of guardrail metric names to threshold values. Values can be floating
	// point numbers (0.0-1.0) for custom thresholds.
	CustomHallucinationThresholdValues interface{} `json:"custom_hallucination_threshold_values"`
	// A description for the workflow, to help you remember what that workflow means to
	// your organization.
	Description string `json:"description"`
	// The action used to improve outputs that fail one or more guardrail metrics for
	// the workflow events.
	ImprovementAction DefendResponseImprovementAction `json:"improvement_action"`
	// A human-readable name for the workflow that will correspond to it's workflow ID.
	Name  string              `json:"name"`
	Stats DefendResponseStats `json:"stats"`
	// Status of the selected workflow. May be `inactive` or `active`. Inactive
	// workflows will not accept events.
	Status DefendResponseStatus `json:"status"`
	// Type of thresholds used to evaluate the event.
	ThresholdType DefendResponseThresholdType `json:"threshold_type"`
	// The most recent time the workflow was updated in UTC.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// A unique workflow ID used to identify the workflow in other endpoints.
	WorkflowID string             `json:"workflow_id"`
	JSON       defendResponseJSON `json:"-"`
}

// defendResponseJSON contains the JSON metadata for the struct [DefendResponse]
type defendResponseJSON struct {
	AutomaticHallucinationToleranceLevels apijson.Field
	CreatedAt                             apijson.Field
	CustomHallucinationThresholdValues    apijson.Field
	Description                           apijson.Field
	ImprovementAction                     apijson.Field
	Name                                  apijson.Field
	Stats                                 apijson.Field
	Status                                apijson.Field
	ThresholdType                         apijson.Field
	UpdatedAt                             apijson.Field
	WorkflowID                            apijson.Field
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

type DefendNewWorkflowResponse = interface{}

type DefendGetEventResponse = interface{}

type DefendSubmitEventResponse = interface{}

type DefendUpdateWorkflowResponse = interface{}

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
	// Any structured information that directly relates to the model’s input and
	// expected output—e.g., the recent turn-by-turn history between an AI tutor and a
	// student, facts or state passed through an agentic workflow, or other
	// domain-specific signals your system already knows and wants the model to
	// condition on.
	Context param.Field[[]string] `json:"context"`
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
