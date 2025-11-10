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

// MonitorService contains methods and other services that help with interacting
// with the deeprails API.
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
func (r *MonitorService) New(ctx context.Context, body MonitorNewParams, opts ...option.RequestOption) (res *MonitorCreateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "monitor"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Use this endpoint to retrieve the details and evaluations associated with a
// specific monitor
func (r *MonitorService) Get(ctx context.Context, monitorID string, query MonitorGetParams, opts ...option.RequestOption) (res *MonitorDetailResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if monitorID == "" {
		err = errors.New("missing required monitor_id parameter")
		return
	}
	path := fmt.Sprintf("monitor/%s", monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Use this endpoint to update the name, description, or status of an existing
// monitor
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

// Use this endpoint to submit a model input and output pair to a monitor for
// evaluation
func (r *MonitorService) SubmitEvent(ctx context.Context, monitorID string, body MonitorSubmitEventParams, opts ...option.RequestOption) (res *MonitorEventResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if monitorID == "" {
		err = errors.New("missing required monitor_id parameter")
		return
	}
	path := fmt.Sprintf("monitor/%s/events", monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type MonitorCreateResponse struct {
	// The time the monitor was created in UTC.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A unique monitor ID.
	MonitorID string `json:"monitor_id,required"`
	// Status of the monitor. Can be `active` or `inactive`. Inactive monitors no
	// longer record and evaluate events.
	Status MonitorCreateResponseStatus `json:"status,required"`
	JSON   monitorCreateResponseJSON   `json:"-"`
}

// monitorCreateResponseJSON contains the JSON metadata for the struct
// [MonitorCreateResponse]
type monitorCreateResponseJSON struct {
	CreatedAt   apijson.Field
	MonitorID   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorCreateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorCreateResponseJSON) RawJSON() string {
	return r.raw
}

// Status of the monitor. Can be `active` or `inactive`. Inactive monitors no
// longer record and evaluate events.
type MonitorCreateResponseStatus string

const (
	MonitorCreateResponseStatusActive   MonitorCreateResponseStatus = "active"
	MonitorCreateResponseStatusInactive MonitorCreateResponseStatus = "inactive"
)

func (r MonitorCreateResponseStatus) IsKnown() bool {
	switch r {
	case MonitorCreateResponseStatusActive, MonitorCreateResponseStatusInactive:
		return true
	}
	return false
}

type MonitorDetailResponse struct {
	// A unique monitor ID.
	MonitorID string `json:"monitor_id,required"`
	// Name of this monitor.
	Name string `json:"name,required"`
	// Status of the monitor. Can be `active` or `inactive`. Inactive monitors no
	// longer record and evaluate events.
	Status MonitorDetailResponseStatus `json:"status,required"`
	// An array of capabilities associated with this monitor.
	Capabilities []MonitorDetailResponseCapability `json:"capabilities"`
	// The time the monitor was created in UTC.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Description of this monitor.
	Description string `json:"description"`
	// An array of all evaluations performed by this monitor. Each one corresponds to a
	// separate monitor event.
	Evaluations []MonitorDetailResponseEvaluation `json:"evaluations"`
	// An array of files associated with this monitor.
	Files []MonitorDetailResponseFile `json:"files"`
	// Contains five fields used for stats of this monitor: total evaluations,
	// completed evaluations, failed evaluations, queued evaluations, and in progress
	// evaluations.
	Stats MonitorDetailResponseStats `json:"stats"`
	// The most recent time the monitor was modified in UTC.
	UpdatedAt time.Time                 `json:"updated_at" format:"date-time"`
	JSON      monitorDetailResponseJSON `json:"-"`
}

// monitorDetailResponseJSON contains the JSON metadata for the struct
// [MonitorDetailResponse]
type monitorDetailResponseJSON struct {
	MonitorID    apijson.Field
	Name         apijson.Field
	Status       apijson.Field
	Capabilities apijson.Field
	CreatedAt    apijson.Field
	Description  apijson.Field
	Evaluations  apijson.Field
	Files        apijson.Field
	Stats        apijson.Field
	UpdatedAt    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MonitorDetailResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorDetailResponseJSON) RawJSON() string {
	return r.raw
}

// Status of the monitor. Can be `active` or `inactive`. Inactive monitors no
// longer record and evaluate events.
type MonitorDetailResponseStatus string

const (
	MonitorDetailResponseStatusActive   MonitorDetailResponseStatus = "active"
	MonitorDetailResponseStatusInactive MonitorDetailResponseStatus = "inactive"
)

func (r MonitorDetailResponseStatus) IsKnown() bool {
	switch r {
	case MonitorDetailResponseStatusActive, MonitorDetailResponseStatusInactive:
		return true
	}
	return false
}

type MonitorDetailResponseCapability struct {
	// The type of capability.
	Capability string                              `json:"capability"`
	JSON       monitorDetailResponseCapabilityJSON `json:"-"`
}

// monitorDetailResponseCapabilityJSON contains the JSON metadata for the struct
// [MonitorDetailResponseCapability]
type monitorDetailResponseCapabilityJSON struct {
	Capability  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorDetailResponseCapability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorDetailResponseCapabilityJSON) RawJSON() string {
	return r.raw
}

type MonitorDetailResponseEvaluation struct {
	// Status of the evaluation.
	EvaluationStatus MonitorDetailResponseEvaluationsEvaluationStatus `json:"evaluation_status,required"`
	// A dictionary of inputs sent to the LLM to generate output. The dictionary must
	// contain at least a `user_prompt` field or a `system_prompt` field. For
	// ground_truth_adherence guardrail metric, `ground_truth` should be provided.
	ModelInput MonitorDetailResponseEvaluationsModelInput `json:"model_input,required"`
	// Output generated by the LLM to be evaluated.
	ModelOutput string `json:"model_output,required"`
	// Run mode for the evaluation. The run mode allows the user to optimize for speed,
	// accuracy, and cost by determining which models are used to evaluate the event.
	RunMode MonitorDetailResponseEvaluationsRunMode `json:"run_mode,required"`
	// The time the evaluation was created in UTC.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Error message if the evaluation failed.
	ErrorMessage string `json:"error_message"`
	// Evaluation result consisting of average scores and rationales for each of the
	// evaluated guardrail metrics.
	EvaluationResult map[string]interface{} `json:"evaluation_result"`
	// Total cost of the evaluation.
	EvaluationTotalCost float64 `json:"evaluation_total_cost"`
	// An array of guardrail metrics that the model input and output pair will be
	// evaluated on.
	GuardrailMetrics []MonitorDetailResponseEvaluationsGuardrailMetric `json:"guardrail_metrics"`
	// An optional, user-defined tag for the evaluation.
	Nametag string `json:"nametag"`
	// Evaluation progress. Values range between 0 and 100; 100 corresponds to a
	// completed `evaluation_status`.
	Progress int64                               `json:"progress"`
	JSON     monitorDetailResponseEvaluationJSON `json:"-"`
}

// monitorDetailResponseEvaluationJSON contains the JSON metadata for the struct
// [MonitorDetailResponseEvaluation]
type monitorDetailResponseEvaluationJSON struct {
	EvaluationStatus    apijson.Field
	ModelInput          apijson.Field
	ModelOutput         apijson.Field
	RunMode             apijson.Field
	CreatedAt           apijson.Field
	ErrorMessage        apijson.Field
	EvaluationResult    apijson.Field
	EvaluationTotalCost apijson.Field
	GuardrailMetrics    apijson.Field
	Nametag             apijson.Field
	Progress            apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *MonitorDetailResponseEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorDetailResponseEvaluationJSON) RawJSON() string {
	return r.raw
}

// Status of the evaluation.
type MonitorDetailResponseEvaluationsEvaluationStatus string

const (
	MonitorDetailResponseEvaluationsEvaluationStatusInProgress MonitorDetailResponseEvaluationsEvaluationStatus = "in_progress"
	MonitorDetailResponseEvaluationsEvaluationStatusCompleted  MonitorDetailResponseEvaluationsEvaluationStatus = "completed"
	MonitorDetailResponseEvaluationsEvaluationStatusCanceled   MonitorDetailResponseEvaluationsEvaluationStatus = "canceled"
	MonitorDetailResponseEvaluationsEvaluationStatusQueued     MonitorDetailResponseEvaluationsEvaluationStatus = "queued"
	MonitorDetailResponseEvaluationsEvaluationStatusFailed     MonitorDetailResponseEvaluationsEvaluationStatus = "failed"
)

func (r MonitorDetailResponseEvaluationsEvaluationStatus) IsKnown() bool {
	switch r {
	case MonitorDetailResponseEvaluationsEvaluationStatusInProgress, MonitorDetailResponseEvaluationsEvaluationStatusCompleted, MonitorDetailResponseEvaluationsEvaluationStatusCanceled, MonitorDetailResponseEvaluationsEvaluationStatusQueued, MonitorDetailResponseEvaluationsEvaluationStatusFailed:
		return true
	}
	return false
}

// A dictionary of inputs sent to the LLM to generate output. The dictionary must
// contain at least a `user_prompt` field or a `system_prompt` field. For
// ground_truth_adherence guardrail metric, `ground_truth` should be provided.
type MonitorDetailResponseEvaluationsModelInput struct {
	// The ground truth for evaluating Ground Truth Adherence guardrail.
	GroundTruth string `json:"ground_truth"`
	// The system prompt used to generate the output.
	SystemPrompt string `json:"system_prompt"`
	// The user prompt used to generate the output.
	UserPrompt string                                         `json:"user_prompt"`
	JSON       monitorDetailResponseEvaluationsModelInputJSON `json:"-"`
}

// monitorDetailResponseEvaluationsModelInputJSON contains the JSON metadata for
// the struct [MonitorDetailResponseEvaluationsModelInput]
type monitorDetailResponseEvaluationsModelInputJSON struct {
	GroundTruth  apijson.Field
	SystemPrompt apijson.Field
	UserPrompt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MonitorDetailResponseEvaluationsModelInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorDetailResponseEvaluationsModelInputJSON) RawJSON() string {
	return r.raw
}

// Run mode for the evaluation. The run mode allows the user to optimize for speed,
// accuracy, and cost by determining which models are used to evaluate the event.
type MonitorDetailResponseEvaluationsRunMode string

const (
	MonitorDetailResponseEvaluationsRunModePrecisionPlus MonitorDetailResponseEvaluationsRunMode = "precision_plus"
	MonitorDetailResponseEvaluationsRunModePrecision     MonitorDetailResponseEvaluationsRunMode = "precision"
	MonitorDetailResponseEvaluationsRunModeSmart         MonitorDetailResponseEvaluationsRunMode = "smart"
	MonitorDetailResponseEvaluationsRunModeEconomy       MonitorDetailResponseEvaluationsRunMode = "economy"
)

func (r MonitorDetailResponseEvaluationsRunMode) IsKnown() bool {
	switch r {
	case MonitorDetailResponseEvaluationsRunModePrecisionPlus, MonitorDetailResponseEvaluationsRunModePrecision, MonitorDetailResponseEvaluationsRunModeSmart, MonitorDetailResponseEvaluationsRunModeEconomy:
		return true
	}
	return false
}

type MonitorDetailResponseEvaluationsGuardrailMetric string

const (
	MonitorDetailResponseEvaluationsGuardrailMetricCorrectness          MonitorDetailResponseEvaluationsGuardrailMetric = "correctness"
	MonitorDetailResponseEvaluationsGuardrailMetricCompleteness         MonitorDetailResponseEvaluationsGuardrailMetric = "completeness"
	MonitorDetailResponseEvaluationsGuardrailMetricInstructionAdherence MonitorDetailResponseEvaluationsGuardrailMetric = "instruction_adherence"
	MonitorDetailResponseEvaluationsGuardrailMetricContextAdherence     MonitorDetailResponseEvaluationsGuardrailMetric = "context_adherence"
	MonitorDetailResponseEvaluationsGuardrailMetricGroundTruthAdherence MonitorDetailResponseEvaluationsGuardrailMetric = "ground_truth_adherence"
	MonitorDetailResponseEvaluationsGuardrailMetricComprehensiveSafety  MonitorDetailResponseEvaluationsGuardrailMetric = "comprehensive_safety"
)

func (r MonitorDetailResponseEvaluationsGuardrailMetric) IsKnown() bool {
	switch r {
	case MonitorDetailResponseEvaluationsGuardrailMetricCorrectness, MonitorDetailResponseEvaluationsGuardrailMetricCompleteness, MonitorDetailResponseEvaluationsGuardrailMetricInstructionAdherence, MonitorDetailResponseEvaluationsGuardrailMetricContextAdherence, MonitorDetailResponseEvaluationsGuardrailMetricGroundTruthAdherence, MonitorDetailResponseEvaluationsGuardrailMetricComprehensiveSafety:
		return true
	}
	return false
}

type MonitorDetailResponseFile struct {
	// The ID of the file.
	FileID string `json:"file_id"`
	// The name of the file.
	FileName string `json:"file_name"`
	// The size of the file in bytes.
	FileSize int64                         `json:"file_size"`
	JSON     monitorDetailResponseFileJSON `json:"-"`
}

// monitorDetailResponseFileJSON contains the JSON metadata for the struct
// [MonitorDetailResponseFile]
type monitorDetailResponseFileJSON struct {
	FileID      apijson.Field
	FileName    apijson.Field
	FileSize    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorDetailResponseFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorDetailResponseFileJSON) RawJSON() string {
	return r.raw
}

// Contains five fields used for stats of this monitor: total evaluations,
// completed evaluations, failed evaluations, queued evaluations, and in progress
// evaluations.
type MonitorDetailResponseStats struct {
	// Number of evaluations that completed successfully.
	CompletedEvaluations int64 `json:"completed_evaluations"`
	// Number of evaluations that failed.
	FailedEvaluations int64 `json:"failed_evaluations"`
	// Number of evaluations currently in progress.
	InProgressEvaluations int64 `json:"in_progress_evaluations"`
	// Number of evaluations currently queued.
	QueuedEvaluations int64 `json:"queued_evaluations"`
	// Total number of evaluations performed by this monitor.
	TotalEvaluations int64                          `json:"total_evaluations"`
	JSON             monitorDetailResponseStatsJSON `json:"-"`
}

// monitorDetailResponseStatsJSON contains the JSON metadata for the struct
// [MonitorDetailResponseStats]
type monitorDetailResponseStatsJSON struct {
	CompletedEvaluations  apijson.Field
	FailedEvaluations     apijson.Field
	InProgressEvaluations apijson.Field
	QueuedEvaluations     apijson.Field
	TotalEvaluations      apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *MonitorDetailResponseStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorDetailResponseStatsJSON) RawJSON() string {
	return r.raw
}

type MonitorEventResponse struct {
	// A unique monitor event ID.
	EventID string `json:"event_id,required"`
	// Monitor ID associated with this event.
	MonitorID string `json:"monitor_id,required"`
	// The time the monitor event was created in UTC.
	CreatedAt time.Time                `json:"created_at" format:"date-time"`
	JSON      monitorEventResponseJSON `json:"-"`
}

// monitorEventResponseJSON contains the JSON metadata for the struct
// [MonitorEventResponse]
type monitorEventResponseJSON struct {
	EventID     apijson.Field
	MonitorID   apijson.Field
	CreatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorEventResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorEventResponseJSON) RawJSON() string {
	return r.raw
}

type MonitorUpdateResponse struct {
	// The time the monitor was last modified in UTC.
	ModifiedAt time.Time `json:"modified_at,required" format:"date-time"`
	// A unique monitor ID.
	MonitorID string `json:"monitor_id,required"`
	// Status of the monitor. Can be `active` or `inactive`. Inactive monitors no
	// longer record and evaluate events.
	Status MonitorUpdateResponseStatus `json:"status,required"`
	JSON   monitorUpdateResponseJSON   `json:"-"`
}

// monitorUpdateResponseJSON contains the JSON metadata for the struct
// [MonitorUpdateResponse]
type monitorUpdateResponseJSON struct {
	ModifiedAt  apijson.Field
	MonitorID   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Status of the monitor. Can be `active` or `inactive`. Inactive monitors no
// longer record and evaluate events.
type MonitorUpdateResponseStatus string

const (
	MonitorUpdateResponseStatusActive   MonitorUpdateResponseStatus = "active"
	MonitorUpdateResponseStatusInactive MonitorUpdateResponseStatus = "inactive"
)

func (r MonitorUpdateResponseStatus) IsKnown() bool {
	switch r {
	case MonitorUpdateResponseStatusActive, MonitorUpdateResponseStatusInactive:
		return true
	}
	return false
}

type MonitorNewParams struct {
	// An array of guardrail metrics that the model input and output pair will be
	// evaluated on. For non-enterprise users, these will be limited to `correctness`,
	// `completeness`, `instruction_adherence`, `context_adherence`,
	// `ground_truth_adherence`, and/or `comprehensive_safety`.
	GuardrailMetrics param.Field[[]MonitorNewParamsGuardrailMetric] `json:"guardrail_metrics,required"`
	// Name of the new monitor.
	Name param.Field[string] `json:"name,required"`
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
	// Description of the monitor.
	Description param.Field[string] `json:"description"`
	// Name of the monitor.
	Name param.Field[string] `json:"name"`
	// Status of the monitor. Can be `active` or `inactive`. Inactive monitors no
	// longer record and evaluate events.
	Status param.Field[MonitorUpdateParamsStatus] `json:"status"`
}

func (r MonitorUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
