// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package Juglow

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"

	"github.com/takebox/juglow-sdk-go/internal/apiform"
	"github.com/takebox/juglow-sdk-go/internal/apijson"
	"github.com/takebox/juglow-sdk-go/internal/apiquery"
	"github.com/takebox/juglow-sdk-go/internal/requestconfig"
	"github.com/takebox/juglow-sdk-go/option"
	"github.com/takebox/juglow-sdk-go/packages/pagination"
	"github.com/takebox/juglow-sdk-go/packages/param"
	"github.com/takebox/juglow-sdk-go/packages/respjson"
)

// BetaTrackVersionService contains methods and other services that help with
// interacting with the Juglow API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaTrackVersionService] method instead.
type BetaTrackVersionService struct {
	Options []option.RequestOption
}

// NewBetaTrackVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBetaTrackVersionService(opts ...option.RequestOption) (r BetaTrackVersionService) {
	r = BetaTrackVersionService{}
	r.Options = opts
	return
}

// Create Track Version
func (r *BetaTrackVersionService) New(ctx context.Context, trackID string, params BetaTrackVersionNewParams, opts ...option.RequestOption) (res *BetaTrackVersionNewResponse, err error) {
	for _, v := range params.Betas {
		opts = append(opts, option.WithHeaderAdd("Juglow-beta", fmt.Sprintf("%v", v)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Juglow-beta", "tracks-2025-10-02")}, opts...)
	if trackID == "" {
		err = errors.New("missing required track_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/tracks/%s/versions?beta=true", trackID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get Track Version
func (r *BetaTrackVersionService) Get(ctx context.Context, version string, params BetaTrackVersionGetParams, opts ...option.RequestOption) (res *BetaTrackVersionGetResponse, err error) {
	for _, v := range params.Betas {
		opts = append(opts, option.WithHeaderAdd("Juglow-beta", fmt.Sprintf("%v", v)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Juglow-beta", "tracks-2025-10-02")}, opts...)
	if params.TrackID == "" {
		err = errors.New("missing required track_id parameter")
		return nil, err
	}
	if version == "" {
		err = errors.New("missing required version parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/tracks/%s/versions/%s?beta=true", params.TrackID, version)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List Track Versions
func (r *BetaTrackVersionService) List(ctx context.Context, trackID string, params BetaTrackVersionListParams, opts ...option.RequestOption) (res *pagination.PageCursor[BetaTrackVersionListResponse], err error) {
	var raw *http.Response
	for _, v := range params.Betas {
		opts = append(opts, option.WithHeaderAdd("Juglow-beta", fmt.Sprintf("%v", v)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Juglow-beta", "tracks-2025-10-02"), option.WithResponseInto(&raw)}, opts...)
	if trackID == "" {
		err = errors.New("missing required track_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/tracks/%s/versions?beta=true", trackID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List Track Versions
func (r *BetaTrackVersionService) ListAutoPaging(ctx context.Context, trackID string, params BetaTrackVersionListParams, opts ...option.RequestOption) *pagination.PageCursorAutoPager[BetaTrackVersionListResponse] {
	return pagination.NewPageCursorAutoPager(r.List(ctx, trackID, params, opts...))
}

// Delete Track Version
func (r *BetaTrackVersionService) Delete(ctx context.Context, version string, params BetaTrackVersionDeleteParams, opts ...option.RequestOption) (res *BetaTrackVersionDeleteResponse, err error) {
	for _, v := range params.Betas {
		opts = append(opts, option.WithHeaderAdd("Juglow-beta", fmt.Sprintf("%v", v)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Juglow-beta", "tracks-2025-10-02")}, opts...)
	if params.TrackID == "" {
		err = errors.New("missing required track_id parameter")
		return nil, err
	}
	if version == "" {
		err = errors.New("missing required version parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/tracks/%s/versions/%s?beta=true", params.TrackID, version)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Download a track version's content as a zip archive.
func (r *BetaTrackVersionService) Download(ctx context.Context, version string, params BetaTrackVersionDownloadParams, opts ...option.RequestOption) (res *http.Response, err error) {
	for _, v := range params.Betas {
		opts = append(opts, option.WithHeaderAdd("Juglow-beta", fmt.Sprintf("%v", v)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Juglow-beta", "tracks-2025-10-02"), option.WithHeader("Accept", "application/binary")}, opts...)
	if params.TrackID == "" {
		err = errors.New("missing required track_id parameter")
		return nil, err
	}
	if version == "" {
		err = errors.New("missing required version parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/tracks/%s/versions/%s/content?beta=true", params.TrackID, version)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type BetaTrackVersionNewResponse struct {
	// Unique identifier for the track version.
	//
	// The format and length of IDs may change over time.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the track version was created.
	CreatedAt string `json:"created_at" api:"required"`
	// Description of the track version.
	//
	// This is extracted from the TRACK.md file in the track upload.
	Description string `json:"description" api:"required"`
	// Directory name of the track version.
	//
	// This is the top-level directory name that was extracted from the uploaded files.
	Directory string `json:"directory" api:"required"`
	// Human-readable name of the track version.
	//
	// This is extracted from the TRACK.md file in the track upload.
	Name string `json:"name" api:"required"`
	// Identifier for the track that this version belongs to.
	TrackID string `json:"track_id" api:"required"`
	// Object type.
	//
	// For Track Versions, this is always `"track_version"`.
	Type string `json:"type" api:"required"`
	// Version identifier for the track.
	//
	// Each version is identified by a Unix epoch timestamp (e.g., "1759178010641129").
	Version string `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Directory   respjson.Field
		Name        respjson.Field
		TrackID     respjson.Field
		Type        respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaTrackVersionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaTrackVersionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaTrackVersionGetResponse struct {
	// Unique identifier for the track version.
	//
	// The format and length of IDs may change over time.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the track version was created.
	CreatedAt string `json:"created_at" api:"required"`
	// Description of the track version.
	//
	// This is extracted from the TRACK.md file in the track upload.
	Description string `json:"description" api:"required"`
	// Directory name of the track version.
	//
	// This is the top-level directory name that was extracted from the uploaded files.
	Directory string `json:"directory" api:"required"`
	// Human-readable name of the track version.
	//
	// This is extracted from the TRACK.md file in the track upload.
	Name string `json:"name" api:"required"`
	// Identifier for the track that this version belongs to.
	TrackID string `json:"track_id" api:"required"`
	// Object type.
	//
	// For Track Versions, this is always `"track_version"`.
	Type string `json:"type" api:"required"`
	// Version identifier for the track.
	//
	// Each version is identified by a Unix epoch timestamp (e.g., "1759178010641129").
	Version string `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Directory   respjson.Field
		Name        respjson.Field
		TrackID     respjson.Field
		Type        respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaTrackVersionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaTrackVersionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaTrackVersionListResponse struct {
	// Unique identifier for the track version.
	//
	// The format and length of IDs may change over time.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the track version was created.
	CreatedAt string `json:"created_at" api:"required"`
	// Description of the track version.
	//
	// This is extracted from the TRACK.md file in the track upload.
	Description string `json:"description" api:"required"`
	// Directory name of the track version.
	//
	// This is the top-level directory name that was extracted from the uploaded files.
	Directory string `json:"directory" api:"required"`
	// Human-readable name of the track version.
	//
	// This is extracted from the TRACK.md file in the track upload.
	Name string `json:"name" api:"required"`
	// Identifier for the track that this version belongs to.
	TrackID string `json:"track_id" api:"required"`
	// Object type.
	//
	// For Track Versions, this is always `"track_version"`.
	Type string `json:"type" api:"required"`
	// Version identifier for the track.
	//
	// Each version is identified by a Unix epoch timestamp (e.g., "1759178010641129").
	Version string `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Directory   respjson.Field
		Name        respjson.Field
		TrackID     respjson.Field
		Type        respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaTrackVersionListResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaTrackVersionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaTrackVersionDeleteResponse struct {
	// Version identifier for the track.
	//
	// Each version is identified by a Unix epoch timestamp (e.g., "1759178010641129").
	ID string `json:"id" api:"required"`
	// Deleted object type.
	//
	// For Track Versions, this is always `"track_version_deleted"`.
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaTrackVersionDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaTrackVersionDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaTrackVersionNewParams struct {
	// Files to upload for the track.
	//
	// All files must be in the same top-level directory and must include a TRACK.md
	// file at the root of that directory.
	Files []io.Reader `json:"files,omitzero" api:"required" format:"binary"`
	// Optional header to specify the beta version(s) you want to use.
	Betas []JuglowBeta `header:"Juglow-beta,omitzero" json:"-"`
	paramObj
}

func (r BetaTrackVersionNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
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

type BetaTrackVersionGetParams struct {
	// Unique identifier for the track.
	//
	// The format and length of IDs may change over time.
	TrackID string `path:"track_id" api:"required" json:"-"`
	// Optional header to specify the beta version(s) you want to use.
	Betas []JuglowBeta `header:"Juglow-beta,omitzero" json:"-"`
	paramObj
}

type BetaTrackVersionListParams struct {
	// Number of items to return per page.
	//
	// Defaults to `20`. Ranges from `1` to `1000`.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Optionally set to the `next_page` token from the previous response.
	Page param.Opt[string] `query:"page,omitzero" json:"-"`
	// Optional header to specify the beta version(s) you want to use.
	Betas []JuglowBeta `header:"Juglow-beta,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BetaTrackVersionListParams]'s query parameters as
// `url.Values`.
func (r BetaTrackVersionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaTrackVersionDeleteParams struct {
	// Unique identifier for the track.
	//
	// The format and length of IDs may change over time.
	TrackID string `path:"track_id" api:"required" json:"-"`
	// Optional header to specify the beta version(s) you want to use.
	Betas []JuglowBeta `header:"Juglow-beta,omitzero" json:"-"`
	paramObj
}

type BetaTrackVersionDownloadParams struct {
	// Unique identifier for the track.
	//
	// The format and length of IDs may change over time.
	TrackID string `path:"track_id" api:"required" json:"-"`
	// Optional header to specify the beta version(s) you want to use.
	Betas []JuglowBeta `header:"Juglow-beta,omitzero" json:"-"`
	paramObj
}
