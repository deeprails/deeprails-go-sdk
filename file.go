// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package deeprails

import (
	"bytes"
	"context"
	"mime/multipart"
	"net/http"
	"slices"

	"github.com/deeprails/deeprails-go-sdk/internal/apiform"
	"github.com/deeprails/deeprails-go-sdk/internal/apijson"
	"github.com/deeprails/deeprails-go-sdk/internal/param"
	"github.com/deeprails/deeprails-go-sdk/internal/requestconfig"
	"github.com/deeprails/deeprails-go-sdk/option"
)

// FileService contains methods and other services that help with interacting with
// the deep rails API.
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
func (r *FileService) Upload(ctx context.Context, body FileUploadParams, opts ...option.RequestOption) (res *FileResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "files/upload"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type FileResponse struct {
	// A unique file ID.
	FileID string `json:"file_id"`
	// Name of the file.
	FileName string `json:"file_name"`
	// The size of the file in bytes.
	FileSize int64            `json:"file_size"`
	JSON     fileResponseJSON `json:"-"`
}

// fileResponseJSON contains the JSON metadata for the struct [FileResponse]
type fileResponseJSON struct {
	FileID      apijson.Field
	FileName    apijson.Field
	FileSize    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FileResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileResponseJSON) RawJSON() string {
	return r.raw
}

type FileUploadParams struct {
	// The contents of the files to upload.
	Files param.Field[[]string] `json:"files,required"`
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
