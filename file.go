// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package deeprails

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"slices"
	"time"

	"github.com/deeprails/deeprails-go-sdk/internal/apiform"
	"github.com/deeprails/deeprails-go-sdk/internal/apijson"
	"github.com/deeprails/deeprails-go-sdk/internal/param"
	"github.com/deeprails/deeprails-go-sdk/internal/requestconfig"
	"github.com/deeprails/deeprails-go-sdk/option"
)

// FileService contains methods and other services that help with interacting with
// the deeprails API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFileService] method instead.
type FileService struct {
	Options []option.RequestOption
}

// NewFileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFileService(opts ...option.RequestOption) (r *FileService) {
	r = &FileService{}
	r.Options = opts
	return
}

// Use this endpoint to upload a file to the DeepRails API
func (r *FileService) Upload(ctx context.Context, body FileUploadParams, opts ...option.RequestOption) (res *FileUploadResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "files/upload"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type FileUploadResponse struct {
	// The time the file was created in UTC.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// A unique file ID.
	FileID string `json:"file_id"`
	// Name of the file.
	FileName string `json:"file_name"`
	// Path to the s3 bucket where the file is stored.
	FilePath string `json:"file_path"`
	// The most recent time the file was modified in UTC.
	UpdatedAt time.Time              `json:"updated_at" format:"date-time"`
	JSON      fileUploadResponseJSON `json:"-"`
}

// fileUploadResponseJSON contains the JSON metadata for the struct
// [FileUploadResponse]
type fileUploadResponseJSON struct {
	CreatedAt   apijson.Field
	FileID      apijson.Field
	FileName    apijson.Field
	FilePath    apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FileUploadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileUploadResponseJSON) RawJSON() string {
	return r.raw
}

type FileUploadParams struct {
	// The contents of the file to upload.
	File param.Field[io.Reader] `json:"file,required" format:"binary"`
}

func (r FileUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}
