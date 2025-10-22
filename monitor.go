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
func (r *MonitorService) New(ctx context.Context, body MonitorNewParams, opts ...option.RequestOption) (res *APIResponse, err error) {
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

// Use this endpoint to update the name, description, or status of an existing
// monitor
func (r *MonitorService) Update(ctx context.Context, monitorID string, body MonitorUpdateParams, opts ...option.RequestOption) (res *APIResponse, err error) {
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

// Response wrapper for operations returning a MonitorResponse.
type APIResponse struct {
	// Represents whether the request was completed successfully.
	Success bool            `json:"success,required"`
	Data    APIResponseData `json:"data"`
	// The accompanying message for the request. Includes error details when
	// applicable.
	Message string          `json:"message"`
	JSON    apiResponseJSON `json:"-"`
}

// apiResponseJSON contains the JSON metadata for the struct [APIResponse]
type apiResponseJSON struct {
	Success     apijson.Field
	Data        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseJSON) RawJSON() string {
	return r.raw
}

type APIResponseData struct {
	// A unique monitor ID.
	MonitorID string `json:"monitor_id,required"`
	// Name of the monitor.
	Name string `json:"name,required"`
	// The time the monitor was created in UTC.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Description of the monitor.
	Description string `json:"description"`
	// Status of the monitor. Can be `active` or `inactive`. Inactive monitors no
	// longer record and evaluate events.
	MonitorStatus APIResponseDataMonitorStatus `json:"monitor_status"`
	// The most recent time the monitor was modified in UTC.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// User ID of the user who created the monitor.
	UserID string              `json:"user_id"`
	JSON   apiResponseDataJSON `json:"-"`
}

// apiResponseDataJSON contains the JSON metadata for the struct [APIResponseData]
type apiResponseDataJSON struct {
	MonitorID     apijson.Field
	Name          apijson.Field
	CreatedAt     apijson.Field
	Description   apijson.Field
	MonitorStatus apijson.Field
	UpdatedAt     apijson.Field
	UserID        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *APIResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseDataJSON) RawJSON() string {
	return r.raw
}

// Status of the monitor. Can be `active` or `inactive`. Inactive monitors no
// longer record and evaluate events.
type APIResponseDataMonitorStatus string

const (
	APIResponseDataMonitorStatusActive   APIResponseDataMonitorStatus = "active"
	APIResponseDataMonitorStatusInactive APIResponseDataMonitorStatus = "inactive"
)

func (r APIResponseDataMonitorStatus) IsKnown() bool {
	switch r {
	case APIResponseDataMonitorStatusActive, APIResponseDataMonitorStatusInactive:
		return true
	}
	return false
}

// Response wrapper for operations returning a MonitorDetailResponse.
type MonitorGetResponse struct {
	// Represents whether the request was completed successfully.
	Success bool                   `json:"success,required"`
	Data    MonitorGetResponseData `json:"data"`
	// The accompanying message for the request. Includes error details when
	// applicable.
	Message string                 `json:"message"`
	JSON    monitorGetResponseJSON `json:"-"`
}

// monitorGetResponseJSON contains the JSON metadata for the struct
// [MonitorGetResponse]
type monitorGetResponseJSON struct {
	Success     apijson.Field
	Data        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorGetResponseJSON) RawJSON() string {
	return r.raw
}

type MonitorGetResponseData struct {
	// A unique monitor ID.
	MonitorID string `json:"monitor_id,required"`
	// Status of the monitor. Can be `active` or `inactive`. Inactive monitors no
	// longer record and evaluate events.
	MonitorStatus MonitorGetResponseDataMonitorStatus `json:"monitor_status,required"`
	// Name of this monitor.
	Name string `json:"name,required"`
	// The time the monitor was created in UTC.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Description of this monitor.
	Description string `json:"description"`
	// An array of all evaluations performed by this monitor. Each one corresponds to a
	// separate monitor event.
	Evaluations []Evaluation `json:"evaluations"`
	// Contains five fields used for stats of this monitor: total evaluations,
	// completed evaluations, failed evaluations, queued evaluations, and in progress
	// evaluations.
	Stats MonitorGetResponseDataStats `json:"stats"`
	// The most recent time the monitor was modified in UTC.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// User ID of the user who created the monitor.
	UserID string                     `json:"user_id"`
	JSON   monitorGetResponseDataJSON `json:"-"`
}

// monitorGetResponseDataJSON contains the JSON metadata for the struct
// [MonitorGetResponseData]
type monitorGetResponseDataJSON struct {
	MonitorID     apijson.Field
	MonitorStatus apijson.Field
	Name          apijson.Field
	CreatedAt     apijson.Field
	Description   apijson.Field
	Evaluations   apijson.Field
	Stats         apijson.Field
	UpdatedAt     apijson.Field
	UserID        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *MonitorGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorGetResponseDataJSON) RawJSON() string {
	return r.raw
}

// Status of the monitor. Can be `active` or `inactive`. Inactive monitors no
// longer record and evaluate events.
type MonitorGetResponseDataMonitorStatus string

const (
	MonitorGetResponseDataMonitorStatusActive   MonitorGetResponseDataMonitorStatus = "active"
	MonitorGetResponseDataMonitorStatusInactive MonitorGetResponseDataMonitorStatus = "inactive"
)

func (r MonitorGetResponseDataMonitorStatus) IsKnown() bool {
	switch r {
	case MonitorGetResponseDataMonitorStatusActive, MonitorGetResponseDataMonitorStatusInactive:
		return true
	}
	return false
}

// Contains five fields used for stats of this monitor: total evaluations,
// completed evaluations, failed evaluations, queued evaluations, and in progress
// evaluations.
type MonitorGetResponseDataStats struct {
	// Number of evaluations that completed successfully.
	CompletedEvaluations int64 `json:"completed_evaluations"`
	// Number of evaluations that failed.
	FailedEvaluations int64 `json:"failed_evaluations"`
	// Number of evaluations currently in progress.
	InProgressEvaluations int64 `json:"in_progress_evaluations"`
	// Number of evaluations currently queued.
	QueuedEvaluations int64 `json:"queued_evaluations"`
	// Total number of evaluations performed by this monitor.
	TotalEvaluations int64                           `json:"total_evaluations"`
	JSON             monitorGetResponseDataStatsJSON `json:"-"`
}

// monitorGetResponseDataStatsJSON contains the JSON metadata for the struct
// [MonitorGetResponseDataStats]
type monitorGetResponseDataStatsJSON struct {
	CompletedEvaluations  apijson.Field
	FailedEvaluations     apijson.Field
	InProgressEvaluations apijson.Field
	QueuedEvaluations     apijson.Field
	TotalEvaluations      apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *MonitorGetResponseDataStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorGetResponseDataStatsJSON) RawJSON() string {
	return r.raw
}

// Response wrapper for operations returning a MonitorEventResponse.
type MonitorSubmitEventResponse struct {
	// Represents whether the request was completed successfully.
	Success bool                           `json:"success,required"`
	Data    MonitorSubmitEventResponseData `json:"data"`
	// The accompanying message for the request. Includes error details when
	// applicable.
	Message string                         `json:"message"`
	JSON    monitorSubmitEventResponseJSON `json:"-"`
}

// monitorSubmitEventResponseJSON contains the JSON metadata for the struct
// [MonitorSubmitEventResponse]
type monitorSubmitEventResponseJSON struct {
	Success     apijson.Field
	Data        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorSubmitEventResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorSubmitEventResponseJSON) RawJSON() string {
	return r.raw
}

type MonitorSubmitEventResponseData struct {
	// A unique evaluation ID associated with this event.
	EvaluationID string `json:"evaluation_id,required"`
	// A unique monitor event ID.
	EventID string `json:"event_id,required"`
	// Monitor ID associated with this event.
	MonitorID string `json:"monitor_id,required"`
	// The time the monitor event was created in UTC.
	CreatedAt time.Time                          `json:"created_at" format:"date-time"`
	JSON      monitorSubmitEventResponseDataJSON `json:"-"`
}

// monitorSubmitEventResponseDataJSON contains the JSON metadata for the struct
// [MonitorSubmitEventResponseData]
type monitorSubmitEventResponseDataJSON struct {
	EvaluationID apijson.Field
	EventID      apijson.Field
	MonitorID    apijson.Field
	CreatedAt    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MonitorSubmitEventResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorSubmitEventResponseDataJSON) RawJSON() string {
	return r.raw
}

type MonitorNewParams struct {
	// Name of the new monitor.
	Name param.Field[string] `json:"name,required"`
	// Description of the new monitor.
	Description param.Field[string] `json:"description"`
}

func (r MonitorNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MonitorGetParams struct {
	// Limit the returned events associated with this monitor. Defaults to 10.
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
	// Status of the monitor. Can be `active` or `inactive`. Inactive monitors no
	// longer record and evaluate events.
	MonitorStatus param.Field[MonitorUpdateParamsMonitorStatus] `json:"monitor_status"`
	// Name of the monitor.
	Name param.Field[string] `json:"name"`
}

func (r MonitorUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Status of the monitor. Can be `active` or `inactive`. Inactive monitors no
// longer record and evaluate events.
type MonitorUpdateParamsMonitorStatus string

const (
	MonitorUpdateParamsMonitorStatusActive   MonitorUpdateParamsMonitorStatus = "active"
	MonitorUpdateParamsMonitorStatusInactive MonitorUpdateParamsMonitorStatus = "inactive"
)

func (r MonitorUpdateParamsMonitorStatus) IsKnown() bool {
	switch r {
	case MonitorUpdateParamsMonitorStatusActive, MonitorUpdateParamsMonitorStatusInactive:
		return true
	}
	return false
}

type MonitorSubmitEventParams struct {
	// An array of guardrail metrics that the model input and output pair will be
	// evaluated on. For non-enterprise users, these will be limited to `correctness`,
	// `completeness`, `instruction_adherence`, `context_adherence`,
	// `ground_truth_adherence`, and/or `comprehensive_safety`.
	GuardrailMetrics param.Field[[]MonitorSubmitEventParamsGuardrailMetric] `json:"guardrail_metrics,required"`
	// A dictionary of inputs sent to the LLM to generate output. The dictionary must
	// contain at least a `user_prompt` or `system_prompt` field. For
	// ground_truth_adherence guardrail metric, `ground_truth` should be provided.
	ModelInput param.Field[MonitorSubmitEventParamsModelInput] `json:"model_input,required"`
	// Output generated by the LLM to be evaluated.
	ModelOutput param.Field[string] `json:"model_output,required"`
	// Model ID used to generate the output, like `gpt-4o` or `o3`.
	ModelUsed param.Field[string] `json:"model_used"`
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

type MonitorSubmitEventParamsGuardrailMetric string

const (
	MonitorSubmitEventParamsGuardrailMetricCorrectness          MonitorSubmitEventParamsGuardrailMetric = "correctness"
	MonitorSubmitEventParamsGuardrailMetricCompleteness         MonitorSubmitEventParamsGuardrailMetric = "completeness"
	MonitorSubmitEventParamsGuardrailMetricInstructionAdherence MonitorSubmitEventParamsGuardrailMetric = "instruction_adherence"
	MonitorSubmitEventParamsGuardrailMetricContextAdherence     MonitorSubmitEventParamsGuardrailMetric = "context_adherence"
	MonitorSubmitEventParamsGuardrailMetricGroundTruthAdherence MonitorSubmitEventParamsGuardrailMetric = "ground_truth_adherence"
	MonitorSubmitEventParamsGuardrailMetricComprehensiveSafety  MonitorSubmitEventParamsGuardrailMetric = "comprehensive_safety"
)

func (r MonitorSubmitEventParamsGuardrailMetric) IsKnown() bool {
	switch r {
	case MonitorSubmitEventParamsGuardrailMetricCorrectness, MonitorSubmitEventParamsGuardrailMetricCompleteness, MonitorSubmitEventParamsGuardrailMetricInstructionAdherence, MonitorSubmitEventParamsGuardrailMetricContextAdherence, MonitorSubmitEventParamsGuardrailMetricGroundTruthAdherence, MonitorSubmitEventParamsGuardrailMetricComprehensiveSafety:
		return true
	}
	return false
}

// A dictionary of inputs sent to the LLM to generate output. The dictionary must
// contain at least a `user_prompt` or `system_prompt` field. For
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
