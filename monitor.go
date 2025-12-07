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
func (r *MonitorService) GetEvent(ctx context.Context, monitorID string, eventID string, opts ...option.RequestOption) (res *MonitorEventDetailResponse, err error) {
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
	// An array of extended AI capabilities associated with this monitor. Can be
	// `web_search`, `file_search`, and/or `context_awareness`.
	Capabilities []MonitorDetailResponseCapability `json:"capabilities,required"`
	// The time the monitor was created in UTC.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// An array of all evaluations performed by this monitor. Each one corresponds to a
	// separate monitor event.
	Evaluations []MonitorDetailResponseEvaluation `json:"evaluations,required"`
	// An array of files associated with this monitor.
	Files []MonitorDetailResponseFile `json:"files,required"`
	// A unique monitor ID.
	MonitorID string `json:"monitor_id,required"`
	// Name of this monitor.
	Name string `json:"name,required"`
	// Contains five fields used for stats of this monitor: total evaluations,
	// completed evaluations, failed evaluations, queued evaluations, and in progress
	// evaluations.
	Stats MonitorDetailResponseStats `json:"stats,required"`
	// Status of the monitor. Can be `active` or `inactive`. Inactive monitors no
	// longer record and evaluate events.
	Status MonitorDetailResponseStatus `json:"status,required"`
	// The most recent time the monitor was modified in UTC.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Description of this monitor.
	Description string                    `json:"description"`
	JSON        monitorDetailResponseJSON `json:"-"`
}

// monitorDetailResponseJSON contains the JSON metadata for the struct
// [MonitorDetailResponse]
type monitorDetailResponseJSON struct {
	Capabilities apijson.Field
	CreatedAt    apijson.Field
	Evaluations  apijson.Field
	Files        apijson.Field
	MonitorID    apijson.Field
	Name         apijson.Field
	Stats        apijson.Field
	Status       apijson.Field
	UpdatedAt    apijson.Field
	Description  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MonitorDetailResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorDetailResponseJSON) RawJSON() string {
	return r.raw
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
	// An array of guardrail metrics that the input and output pair will be evaluated
	// on.
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
	// Any structured information that directly relates to the model’s input and
	// expected output—e.g., the recent turn-by-turn history between an AI tutor and a
	// student, facts or state passed through an agentic workflow, or other
	// domain-specific signals your system already knows and wants the model to
	// condition on.
	Context []string `json:"context"`
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
	Context      apijson.Field
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

type MonitorEventDetailResponse struct {
	// The extended AI capabilities associated with the monitor event. Can be
	// `web_search`, `file_search`, and/or `context_awareness`.
	Capabilities []MonitorEventDetailResponseCapability `json:"capabilities"`
	// The time spent on the evaluation in seconds.
	EvalTime string `json:"eval_time"`
	// The result of the evaluation of the monitor event.
	EvaluationResult map[string]interface{} `json:"evaluation_result"`
	// A unique monitor event ID.
	EventID string `json:"event_id"`
	// The files associated with the monitor event.
	Files []MonitorEventDetailResponseFile `json:"files"`
	// The guardrail metrics evaluated by the monitor event.
	GuardrailMetrics []string `json:"guardrail_metrics"`
	// The model input used to create the monitor event.
	ModelInput map[string]interface{} `json:"model_input"`
	// The output evaluated by the monitor event.
	ModelOutput string `json:"model_output"`
	// Monitor ID associated with this event.
	MonitorID string `json:"monitor_id"`
	// A human-readable tag for the monitor event.
	Nametag string `json:"nametag"`
	// The run mode used to evaluate the monitor event.
	RunMode MonitorEventDetailResponseRunMode `json:"run_mode"`
	// Status of the monitor event's evaluation.
	Status MonitorEventDetailResponseStatus `json:"status"`
	// The time the monitor event was created in UTC.
	Timestamp time.Time                      `json:"timestamp" format:"date-time"`
	JSON      monitorEventDetailResponseJSON `json:"-"`
}

// monitorEventDetailResponseJSON contains the JSON metadata for the struct
// [MonitorEventDetailResponse]
type monitorEventDetailResponseJSON struct {
	Capabilities     apijson.Field
	EvalTime         apijson.Field
	EvaluationResult apijson.Field
	EventID          apijson.Field
	Files            apijson.Field
	GuardrailMetrics apijson.Field
	ModelInput       apijson.Field
	ModelOutput      apijson.Field
	MonitorID        apijson.Field
	Nametag          apijson.Field
	RunMode          apijson.Field
	Status           apijson.Field
	Timestamp        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MonitorEventDetailResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorEventDetailResponseJSON) RawJSON() string {
	return r.raw
}

type MonitorEventDetailResponseCapability struct {
	// The type of capability.
	Capability string                                   `json:"capability"`
	JSON       monitorEventDetailResponseCapabilityJSON `json:"-"`
}

// monitorEventDetailResponseCapabilityJSON contains the JSON metadata for the
// struct [MonitorEventDetailResponseCapability]
type monitorEventDetailResponseCapabilityJSON struct {
	Capability  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorEventDetailResponseCapability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorEventDetailResponseCapabilityJSON) RawJSON() string {
	return r.raw
}

type MonitorEventDetailResponseFile struct {
	// The ID of the file.
	FileID string `json:"file_id"`
	// The name of the file.
	FileName string `json:"file_name"`
	// The size of the file in bytes.
	FileSize int64                              `json:"file_size"`
	JSON     monitorEventDetailResponseFileJSON `json:"-"`
}

// monitorEventDetailResponseFileJSON contains the JSON metadata for the struct
// [MonitorEventDetailResponseFile]
type monitorEventDetailResponseFileJSON struct {
	FileID      apijson.Field
	FileName    apijson.Field
	FileSize    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorEventDetailResponseFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorEventDetailResponseFileJSON) RawJSON() string {
	return r.raw
}

// The run mode used to evaluate the monitor event.
type MonitorEventDetailResponseRunMode string

const (
	MonitorEventDetailResponseRunModePrecisionPlus MonitorEventDetailResponseRunMode = "precision_plus"
	MonitorEventDetailResponseRunModePrecision     MonitorEventDetailResponseRunMode = "precision"
	MonitorEventDetailResponseRunModeSmart         MonitorEventDetailResponseRunMode = "smart"
	MonitorEventDetailResponseRunModeEconomy       MonitorEventDetailResponseRunMode = "economy"
)

func (r MonitorEventDetailResponseRunMode) IsKnown() bool {
	switch r {
	case MonitorEventDetailResponseRunModePrecisionPlus, MonitorEventDetailResponseRunModePrecision, MonitorEventDetailResponseRunModeSmart, MonitorEventDetailResponseRunModeEconomy:
		return true
	}
	return false
}

// Status of the monitor event's evaluation.
type MonitorEventDetailResponseStatus string

const (
	MonitorEventDetailResponseStatusInProgress MonitorEventDetailResponseStatus = "in_progress"
	MonitorEventDetailResponseStatusCompleted  MonitorEventDetailResponseStatus = "completed"
	MonitorEventDetailResponseStatusCanceled   MonitorEventDetailResponseStatus = "canceled"
	MonitorEventDetailResponseStatusQueued     MonitorEventDetailResponseStatus = "queued"
	MonitorEventDetailResponseStatusFailed     MonitorEventDetailResponseStatus = "failed"
)

func (r MonitorEventDetailResponseStatus) IsKnown() bool {
	switch r {
	case MonitorEventDetailResponseStatusInProgress, MonitorEventDetailResponseStatusCompleted, MonitorEventDetailResponseStatusCanceled, MonitorEventDetailResponseStatusQueued, MonitorEventDetailResponseStatusFailed:
		return true
	}
	return false
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
