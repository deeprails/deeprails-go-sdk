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

	"github.com/stainless-sdks/deeprails-go/internal/apijson"
	"github.com/stainless-sdks/deeprails-go/internal/apiquery"
	"github.com/stainless-sdks/deeprails-go/internal/requestconfig"
	"github.com/stainless-sdks/deeprails-go/option"
	"github.com/stainless-sdks/deeprails-go/packages/param"
	"github.com/stainless-sdks/deeprails-go/packages/respjson"
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
func NewMonitorService(opts ...option.RequestOption) (r MonitorService) {
	r = MonitorService{}
	r.Options = opts
	return
}

// Create a new monitor to evaluate model inputs and outputs using guardrails.
func (r *MonitorService) New(ctx context.Context, body MonitorNewParams, opts ...option.RequestOption) (res *APIResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "monitor"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve the details and evaluations associated with a specific monitor.
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

// Update the name, description, or status of an existing monitor.
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

// Submit a model input and output pair to a monitor for evaluation.
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
	Success bool `json:"success,required"`
	// Response payload for creating or updating a monitor.
	Data APIResponseData `json:"data"`
	// The accompanying message for the request. Includes error details when
	// applicable.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		Data        respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIResponse) RawJSON() string { return r.JSON.raw }
func (r *APIResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response payload for creating or updating a monitor.
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
	//
	// Any of "active", "inactive".
	MonitorStatus string `json:"monitor_status"`
	// The most recent time the monitor was modified in UTC.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// User ID of the user who created the monitor.
	UserID string `json:"user_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MonitorID     respjson.Field
		Name          respjson.Field
		CreatedAt     respjson.Field
		Description   respjson.Field
		MonitorStatus respjson.Field
		UpdatedAt     respjson.Field
		UserID        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIResponseData) RawJSON() string { return r.JSON.raw }
func (r *APIResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response wrapper for operations returning a MonitorDetailResponse.
type MonitorGetResponse struct {
	// Represents whether the request was completed successfully.
	Success bool `json:"success,required"`
	// Detailed response payload for retrieving a monitor and its evaluations.
	Data MonitorGetResponseData `json:"data"`
	// The accompanying message for the request. Includes error details when
	// applicable.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		Data        respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detailed response payload for retrieving a monitor and its evaluations.
type MonitorGetResponseData struct {
	// A unique monitor ID.
	MonitorID string `json:"monitor_id,required"`
	// Status of the monitor. Can be `active` or `inactive`. Inactive monitors no
	// longer record and evaluate events.
	//
	// Any of "active", "inactive".
	MonitorStatus string `json:"monitor_status,required"`
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
	UserID string `json:"user_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MonitorID     respjson.Field
		MonitorStatus respjson.Field
		Name          respjson.Field
		CreatedAt     respjson.Field
		Description   respjson.Field
		Evaluations   respjson.Field
		Stats         respjson.Field
		UpdatedAt     respjson.Field
		UserID        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	TotalEvaluations int64 `json:"total_evaluations"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedEvaluations  respjson.Field
		FailedEvaluations     respjson.Field
		InProgressEvaluations respjson.Field
		QueuedEvaluations     respjson.Field
		TotalEvaluations      respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorGetResponseDataStats) RawJSON() string { return r.JSON.raw }
func (r *MonitorGetResponseDataStats) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response wrapper for operations returning a MonitorEventResponse.
type MonitorSubmitEventResponse struct {
	// Represents whether the request was completed successfully.
	Success bool `json:"success,required"`
	// Response payload for monitor event operations.
	Data MonitorSubmitEventResponseData `json:"data"`
	// The accompanying message for the request. Includes error details when
	// applicable.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		Data        respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorSubmitEventResponse) RawJSON() string { return r.JSON.raw }
func (r *MonitorSubmitEventResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response payload for monitor event operations.
type MonitorSubmitEventResponseData struct {
	// A unique evaluation ID associated with this event.
	EvaluationID string `json:"evaluation_id,required"`
	// A unique monitor event ID.
	EventID string `json:"event_id,required"`
	// Monitor ID associated with this event.
	MonitorID string `json:"monitor_id,required"`
	// The time the monitor event was created in UTC.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EvaluationID respjson.Field
		EventID      respjson.Field
		MonitorID    respjson.Field
		CreatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorSubmitEventResponseData) RawJSON() string { return r.JSON.raw }
func (r *MonitorSubmitEventResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorNewParams struct {
	// Name of the new monitor.
	Name string `json:"name,required"`
	// Description of the new monitor.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r MonitorNewParams) MarshalJSON() (data []byte, err error) {
	type shadow MonitorNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorGetParams struct {
	// Limit the returned events associated with this monitor. Defaults to 10.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MonitorGetParams]'s query parameters as `url.Values`.
func (r MonitorGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MonitorUpdateParams struct {
	// Description of the monitor.
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the monitor.
	Name param.Opt[string] `json:"name,omitzero"`
	// Status of the monitor. Can be `active` or `inactive`. Inactive monitors no
	// longer record and evaluate events.
	//
	// Any of "active", "inactive".
	MonitorStatus MonitorUpdateParamsMonitorStatus `json:"monitor_status,omitzero"`
	paramObj
}

func (r MonitorUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow MonitorUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the monitor. Can be `active` or `inactive`. Inactive monitors no
// longer record and evaluate events.
type MonitorUpdateParamsMonitorStatus string

const (
	MonitorUpdateParamsMonitorStatusActive   MonitorUpdateParamsMonitorStatus = "active"
	MonitorUpdateParamsMonitorStatusInactive MonitorUpdateParamsMonitorStatus = "inactive"
)

type MonitorSubmitEventParams struct {
	// An array of guardrail metrics that the model input and output pair will be
	// evaluated on. For non-enterprise users, these will be limited to `correctness`,
	// `completeness`, `instruction_adherence`, `context_adherence`,
	// `ground_truth_adherence`, and/or `comprehensive_safety`.
	//
	// Any of "correctness", "completeness", "instruction_adherence",
	// "context_adherence", "ground_truth_adherence", "comprehensive_safety".
	GuardrailMetrics []string `json:"guardrail_metrics,omitzero,required"`
	// A dictionary of inputs sent to the LLM to generate output. This must contain a
	// `user_prompt` field and an optional `context` field. Additional properties are
	// allowed.
	ModelInput MonitorSubmitEventParamsModelInput `json:"model_input,omitzero,required"`
	// Output generated by the LLM to be evaluated.
	ModelOutput string `json:"model_output,required"`
	// Model ID used to generate the output, like `gpt-4o` or `o3`.
	ModelUsed param.Opt[string] `json:"model_used,omitzero"`
	// An optional, user-defined tag for the event.
	Nametag param.Opt[string] `json:"nametag,omitzero"`
	// Run mode for the monitor event. The run mode allows the user to optimize for
	// speed, accuracy, and cost by determining which models are used to evaluate the
	// event. Available run modes include `precision_plus`, `precision`, `smart`, and
	// `economy`. Defaults to `smart`.
	//
	// Any of "precision_plus", "precision", "smart", "economy".
	RunMode MonitorSubmitEventParamsRunMode `json:"run_mode,omitzero"`
	paramObj
}

func (r MonitorSubmitEventParams) MarshalJSON() (data []byte, err error) {
	type shadow MonitorSubmitEventParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorSubmitEventParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A dictionary of inputs sent to the LLM to generate output. This must contain a
// `user_prompt` field and an optional `context` field. Additional properties are
// allowed.
//
// The property UserPrompt is required.
type MonitorSubmitEventParamsModelInput struct {
	UserPrompt  string            `json:"user_prompt,required"`
	Context     param.Opt[string] `json:"context,omitzero"`
	ExtraFields map[string]any    `json:"-"`
	paramObj
}

func (r MonitorSubmitEventParamsModelInput) MarshalJSON() (data []byte, err error) {
	type shadow MonitorSubmitEventParamsModelInput
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *MonitorSubmitEventParamsModelInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
