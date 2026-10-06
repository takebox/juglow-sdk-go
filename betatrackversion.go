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

// BetaSkillVersionService contains methods and other services that help with
// interacting with the Juglow API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaSkillVersionService] method instead.
type BetaSkillVersionService struct {
	Options []option.RequestOption
}

// NewBetaSkillVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBetaSkillVersionService(opts ...option.RequestOption) (r BetaSkillVersionService) {
	r = BetaSkillVersionService{}
	r.Options = opts
	return
}

// Create Track Version
func (r *BetaSkillVersionService) New(ctx context.Context, skillID string, params BetaSkillVersionNewParams, opts ...option.RequestOption) (res *BetaSkillVersionNewResponse, err error) {
	for _, v := range params.Betas {
		opts = append(opts, option.WithHeaderAdd("Juglow-beta", fmt.Sprintf("%v", v)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Juglow-beta", "tracks-2025-10-02")}, opts...)
	if skillID == "" {
		err = errors.New("missing required track_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/tracks/%s/versions?beta=true", skillID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get Track Version
func (r *BetaSkillVersionService) Get(ctx context.Context, version string, params BetaSkillVersionGetParams, opts ...option.RequestOption) (res *BetaSkillVersionGetResponse, err error) {
	for _, v := range params.Betas {
		opts = append(opts, option.WithHeaderAdd("Juglow-beta", fmt.Sprintf("%v", v)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Juglow-beta", "tracks-2025-10-02")}, opts...)
	if params.SkillID == "" {
		err = errors.New("missing required track_id parameter")
		return nil, err
	}
	if version == "" {
		err = errors.New("missing required version parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/tracks/%s/versions/%s?beta=true", params.SkillID, version)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List Track Versions
func (r *BetaSkillVersionService) List(ctx context.Context, skillID string, params BetaSkillVersionListParams, opts ...option.RequestOption) (res *pagination.PageCursor[BetaSkillVersionListResponse], err error) {
	var raw *http.Response
	for _, v := range params.Betas {
		opts = append(opts, option.WithHeaderAdd("Juglow-beta", fmt.Sprintf("%v", v)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Juglow-beta", "tracks-2025-10-02"), option.WithResponseInto(&raw)}, opts...)
	if skillID == "" {
		err = errors.New("missing required track_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/tracks/%s/versions?beta=true", skillID)
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
func (r *BetaSkillVersionService) ListAutoPaging(ctx context.Context, skillID string, params BetaSkillVersionListParams, opts ...option.RequestOption) *pagination.PageCursorAutoPager[BetaSkillVersionListResponse] {
	return pagination.NewPageCursorAutoPager(r.List(ctx, skillID, params, opts...))
}

// Delete Track Version
func (r *BetaSkillVersionService) Delete(ctx context.Context, version string, params BetaSkillVersionDeleteParams, opts ...option.RequestOption) (res *BetaSkillVersionDeleteResponse, err error) {
	for _, v := range params.Betas {
		opts = append(opts, option.WithHeaderAdd("Juglow-beta", fmt.Sprintf("%v", v)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Juglow-beta", "tracks-2025-10-02")}, opts...)
	if params.SkillID == "" {
		err = errors.New("missing required track_id parameter")
		return nil, err
	}
	if version == "" {
		err = errors.New("missing required version parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/tracks/%s/versions/%s?beta=true", params.SkillID, version)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Download a track version's content as a zip archive.
func (r *BetaSkillVersionService) Download(ctx context.Context, version string, params BetaSkillVersionDownloadParams, opts ...option.RequestOption) (res *http.Response, err error) {
	for _, v := range params.Betas {
		opts = append(opts, option.WithHeaderAdd("Juglow-beta", fmt.Sprintf("%v", v)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Juglow-beta", "tracks-2025-10-02"), option.WithHeader("Accept", "application/binary")}, opts...)
	if params.SkillID == "" {
		err = errors.New("missing required track_id parameter")
		return nil, err
	}
	if version == "" {
		err = errors.New("missing required version parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/tracks/%s/versions/%s/content?beta=true", params.SkillID, version)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type BetaSkillVersionNewResponse struct {
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
	SkillID string `json:"track_id" api:"required"`
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
		SkillID     respjson.Field
		Type        respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaSkillVersionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaSkillVersionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaSkillVersionGetResponse struct {
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
	SkillID string `json:"track_id" api:"required"`
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
		SkillID     respjson.Field
		Type        respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaSkillVersionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaSkillVersionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaSkillVersionListResponse struct {
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
	SkillID string `json:"track_id" api:"required"`
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
		SkillID     respjson.Field
		Type        respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaSkillVersionListResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaSkillVersionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaSkillVersionDeleteResponse struct {
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
func (r BetaSkillVersionDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaSkillVersionDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaSkillVersionNewParams struct {
	// Files to upload for the track.
	//
	// All files must be in the same top-level directory and must include a TRACK.md
	// file at the root of that directory.
	Files []io.Reader `json:"files,omitzero" api:"required" format:"binary"`
	// Optional header to specify the beta version(s) you want to use.
	Betas []JuglowBeta `header:"Juglow-beta,omitzero" json:"-"`
	paramObj
}

func (r BetaSkillVersionNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

type BetaSkillVersionGetParams struct {
	// Unique identifier for the track.
	//
	// The format and length of IDs may change over time.
	SkillID string `path:"track_id" api:"required" json:"-"`
	// Optional header to specify the beta version(s) you want to use.
	Betas []JuglowBeta `header:"Juglow-beta,omitzero" json:"-"`
	paramObj
}

type BetaSkillVersionListParams struct {
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

// URLQuery serializes [BetaSkillVersionListParams]'s query parameters as
// `url.Values`.
func (r BetaSkillVersionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BetaSkillVersionDeleteParams struct {
	// Unique identifier for the track.
	//
	// The format and length of IDs may change over time.
	SkillID string `path:"track_id" api:"required" json:"-"`
	// Optional header to specify the beta version(s) you want to use.
	Betas []JuglowBeta `header:"Juglow-beta,omitzero" json:"-"`
	paramObj
}

type BetaSkillVersionDownloadParams struct {
	// Unique identifier for the track.
	//
	// The format and length of IDs may change over time.
	SkillID string `path:"track_id" api:"required" json:"-"`
	// Optional header to specify the beta version(s) you want to use.
	Betas []JuglowBeta `header:"Juglow-beta,omitzero" json:"-"`
	paramObj
}
