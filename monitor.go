// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package deeprails

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/deeprails/deeprails-go-sdk/internal/apijson"
	"github.com/deeprails/deeprails-go-sdk/internal/apiquery"
	"github.com/deeprails/deeprails-go-sdk/internal/param"
	"github.com/deeprails/deeprails-go-sdk/internal/requestconfig"
	"github.com/deeprails/deeprails-go-sdk/option"
)

// MonitorService contains methods and other services that help with interacting
// with the deep rails API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMonitorService] method instead.
type MonitorService struct {
	Options []option.RequestOption
}

// NewMonitorService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMonitorService(opts ...option.RequestOption) (r *MonitorService) {
	r = &MonitorService{}
	r.Options = opts
	return
}

// Use this endpoint to create a new monitor to evaluate model inputs and outputs
// using guardrails
func (r *MonitorService) New(ctx context.Context, body MonitorNewParams, opts ...option.RequestOption) (res *MonitorNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "monitor"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Use this endpoint to retrieve the details and evaluations associated with a
// specific monitor
func (r *MonitorService) Get(ctx context.Context, monitorID string, query MonitorGetParams, opts ...option.RequestOption) (res *MonitorGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if monitorID == "" {
		err = errors.New("missing required monitor_id parameter")
		return
	}
	path := fmt.Sprintf("monitor/%s", monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Use this endpoint to update the name, status, and/or other details of an
// existing monitor.
func (r *MonitorService) Update(ctx context.Context, monitorID string, body MonitorUpdateParams, opts ...option.RequestOption) (res *MonitorUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if monitorID == "" {
		err = errors.New("missing required monitor_id parameter")
		return
	}
	path := fmt.Sprintf("monitor/%s", monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Use this endpoint to retrieve the details of a specific monitor event
func (r *MonitorService) GetEvent(ctx context.Context, monitorID string, eventID string, opts ...option.RequestOption) (res *MonitorGetEventResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if monitorID == "" {
		err = errors.New("missing required monitor_id parameter")
		return
	}
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	path := fmt.Sprintf("monitor/%s/events/%s", monitorID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Use this endpoint to submit a model input and output pair to a monitor for
// evaluation
func (r *MonitorService) SubmitEvent(ctx context.Context, monitorID string, body MonitorSubmitEventParams, opts ...option.RequestOption) (res *MonitorSubmitEventResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if monitorID == "" {
		err = errors.New("missing required monitor_id parameter")
		return
	}
	path := fmt.Sprintf("monitor/%s/events", monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type MonitorNewResponse = interface{}

type MonitorGetResponse = interface{}

type MonitorUpdateResponse = interface{}

type MonitorGetEventResponse = interface{}

type MonitorSubmitEventResponse = interface{}

type MonitorNewParams struct {
	// An array of guardrail metrics that the model input and output pair will be
	// evaluated on. For non-enterprise users, these will be limited to `correctness`,
	// `completeness`, `instruction_adherence`, `context_adherence`,
	// `ground_truth_adherence`, and/or `comprehensive_safety`.
	GuardrailMetrics param.Field[[]MonitorNewParamsGuardrailMetric] `json:"guardrail_metrics,required"`
	// Name of the new monitor.
	Name param.Field[string] `json:"name,required"`
	// Context includes any structured information that directly relates to the model’s
	// input and expected output—e.g., the recent turn-by-turn history between an AI
	// tutor and a student, facts or state passed through an agentic workflow, or other
	// domain-specific signals your system already knows and wants the model to
	// condition on. This field determines whether to enable context awareness for this
	// monitor's evaluations. Defaults to false.
	ContextAwareness param.Field[bool] `json:"context_awareness"`
	// Description of the new monitor.
	Description param.Field[string] `json:"description"`
	// An array of file IDs to search in the monitor's evaluations. Files must be
	// uploaded via the DeepRails API first.
	FileSearch param.Field[[]string] `json:"file_search"`
	// Whether to enable web search for this monitor's evaluations. Defaults to false.
	WebSearch param.Field[bool] `json:"web_search"`
}

func (r MonitorNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MonitorNewParamsGuardrailMetric string

const (
	MonitorNewParamsGuardrailMetricCorrectness          MonitorNewParamsGuardrailMetric = "correctness"
	MonitorNewParamsGuardrailMetricCompleteness         MonitorNewParamsGuardrailMetric = "completeness"
	MonitorNewParamsGuardrailMetricInstructionAdherence MonitorNewParamsGuardrailMetric = "instruction_adherence"
	MonitorNewParamsGuardrailMetricContextAdherence     MonitorNewParamsGuardrailMetric = "context_adherence"
	MonitorNewParamsGuardrailMetricGroundTruthAdherence MonitorNewParamsGuardrailMetric = "ground_truth_adherence"
	MonitorNewParamsGuardrailMetricComprehensiveSafety  MonitorNewParamsGuardrailMetric = "comprehensive_safety"
)

func (r MonitorNewParamsGuardrailMetric) IsKnown() bool {
	switch r {
	case MonitorNewParamsGuardrailMetricCorrectness, MonitorNewParamsGuardrailMetricCompleteness, MonitorNewParamsGuardrailMetricInstructionAdherence, MonitorNewParamsGuardrailMetricContextAdherence, MonitorNewParamsGuardrailMetricGroundTruthAdherence, MonitorNewParamsGuardrailMetricComprehensiveSafety:
		return true
	}
	return false
}

type MonitorGetParams struct {
	// Limit the number of returned evaluations associated with this monitor. Defaults
	// to 10.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [MonitorGetParams]'s query parameters as `url.Values`.
func (r MonitorGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MonitorUpdateParams struct {
	// New description of the monitor.
	Description param.Field[string] `json:"description"`
	// An array of file IDs to search in the monitor's evaluations. Files must be
	// uploaded via the DeepRails API first.
	FileSearch param.Field[[]string] `json:"file_search"`
	// An array of the new guardrail metrics that model input and output pairs will be
	// evaluated on.
	GuardrailMetrics param.Field[[]MonitorUpdateParamsGuardrailMetric] `json:"guardrail_metrics"`
	// New name of the monitor.
	Name param.Field[string] `json:"name"`
	// Status of the monitor. Can be `active` or `inactive`. Inactive monitors no
	// longer record and evaluate events.
	Status param.Field[MonitorUpdateParamsStatus] `json:"status"`
	// Whether to enable web search for this monitor's evaluations.
	WebSearch param.Field[bool] `json:"web_search"`
}

func (r MonitorUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MonitorUpdateParamsGuardrailMetric string

const (
	MonitorUpdateParamsGuardrailMetricCorrectness          MonitorUpdateParamsGuardrailMetric = "correctness"
	MonitorUpdateParamsGuardrailMetricCompleteness         MonitorUpdateParamsGuardrailMetric = "completeness"
	MonitorUpdateParamsGuardrailMetricInstructionAdherence MonitorUpdateParamsGuardrailMetric = "instruction_adherence"
	MonitorUpdateParamsGuardrailMetricContextAdherence     MonitorUpdateParamsGuardrailMetric = "context_adherence"
	MonitorUpdateParamsGuardrailMetricGroundTruthAdherence MonitorUpdateParamsGuardrailMetric = "ground_truth_adherence"
	MonitorUpdateParamsGuardrailMetricComprehensiveSafety  MonitorUpdateParamsGuardrailMetric = "comprehensive_safety"
)

func (r MonitorUpdateParamsGuardrailMetric) IsKnown() bool {
	switch r {
	case MonitorUpdateParamsGuardrailMetricCorrectness, MonitorUpdateParamsGuardrailMetricCompleteness, MonitorUpdateParamsGuardrailMetricInstructionAdherence, MonitorUpdateParamsGuardrailMetricContextAdherence, MonitorUpdateParamsGuardrailMetricGroundTruthAdherence, MonitorUpdateParamsGuardrailMetricComprehensiveSafety:
		return true
	}
	return false
}

// Status of the monitor. Can be `active` or `inactive`. Inactive monitors no
// longer record and evaluate events.
type MonitorUpdateParamsStatus string

const (
	MonitorUpdateParamsStatusActive   MonitorUpdateParamsStatus = "active"
	MonitorUpdateParamsStatusInactive MonitorUpdateParamsStatus = "inactive"
)

func (r MonitorUpdateParamsStatus) IsKnown() bool {
	switch r {
	case MonitorUpdateParamsStatusActive, MonitorUpdateParamsStatusInactive:
		return true
	}
	return false
}

type MonitorSubmitEventParams struct {
	// A dictionary of inputs sent to the LLM to generate output. The dictionary must
	// contain at least a `user_prompt` field or a `system_prompt` field. For
	// ground_truth_adherence guardrail metric, `ground_truth` should be provided.
	ModelInput param.Field[MonitorSubmitEventParamsModelInput] `json:"model_input,required"`
	// Output generated by the LLM to be evaluated.
	ModelOutput param.Field[string] `json:"model_output,required"`
	// An optional, user-defined tag for the event.
	Nametag param.Field[string] `json:"nametag"`
	// Run mode for the monitor event. The run mode allows the user to optimize for
	// speed, accuracy, and cost by determining which models are used to evaluate the
	// event. Available run modes include `precision_plus`, `precision`, `smart`, and
	// `economy`. Defaults to `smart`.
	RunMode param.Field[MonitorSubmitEventParamsRunMode] `json:"run_mode"`
}

func (r MonitorSubmitEventParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A dictionary of inputs sent to the LLM to generate output. The dictionary must
// contain at least a `user_prompt` field or a `system_prompt` field. For
// ground_truth_adherence guardrail metric, `ground_truth` should be provided.
type MonitorSubmitEventParamsModelInput struct {
	// Any structured information that directly relates to the model’s input and
	// expected output—e.g., the recent turn-by-turn history between an AI tutor and a
	// student, facts or state passed through an agentic workflow, or other
	// domain-specific signals your system already knows and wants the model to
	// condition on.
	Context param.Field[[]string] `json:"context"`
	// The ground truth for evaluating Ground Truth Adherence guardrail.
	GroundTruth param.Field[string] `json:"ground_truth"`
	// The system prompt used to generate the output.
	SystemPrompt param.Field[string] `json:"system_prompt"`
	// The user prompt used to generate the output.
	UserPrompt param.Field[string] `json:"user_prompt"`
}

func (r MonitorSubmitEventParamsModelInput) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Run mode for the monitor event. The run mode allows the user to optimize for
// speed, accuracy, and cost by determining which models are used to evaluate the
// event. Available run modes include `precision_plus`, `precision`, `smart`, and
// `economy`. Defaults to `smart`.
type MonitorSubmitEventParamsRunMode string

const (
	MonitorSubmitEventParamsRunModePrecisionPlus MonitorSubmitEventParamsRunMode = "precision_plus"
	MonitorSubmitEventParamsRunModePrecision     MonitorSubmitEventParamsRunMode = "precision"
	MonitorSubmitEventParamsRunModeSmart         MonitorSubmitEventParamsRunMode = "smart"
	MonitorSubmitEventParamsRunModeEconomy       MonitorSubmitEventParamsRunMode = "economy"
)

func (r MonitorSubmitEventParamsRunMode) IsKnown() bool {
	switch r {
	case MonitorSubmitEventParamsRunModePrecisionPlus, MonitorSubmitEventParamsRunModePrecision, MonitorSubmitEventParamsRunModeSmart, MonitorSubmitEventParamsRunModeEconomy:
		return true
	}
	return false
}
