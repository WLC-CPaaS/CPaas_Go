/*
White Label Communications CPaas API Documentation

A CPaaS platform API

API version: 1.1
Contact: support@whitelabelcomm.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"os"
)


// MediaAPIService MediaAPI service
type MediaAPIService service

type ApiV1AccountAccountIDMediaMediaIDFileGetRequest struct {
	ctx context.Context
	ApiService *MediaAPIService
	accountID string
	mediaID string
}

func (r ApiV1AccountAccountIDMediaMediaIDFileGetRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.V1AccountAccountIDMediaMediaIDFileGetExecute(r)
}

/*
V1AccountAccountIDMediaMediaIDFileGet Get Media File

Gather data about the media objects in an account.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountID Account ID, 32 alpha numeric
 @param mediaID Media ID, 32 alpha numeric
 @return ApiV1AccountAccountIDMediaMediaIDFileGetRequest
*/
func (a *MediaAPIService) V1AccountAccountIDMediaMediaIDFileGet(ctx context.Context, accountID string, mediaID string) ApiV1AccountAccountIDMediaMediaIDFileGetRequest {
	return ApiV1AccountAccountIDMediaMediaIDFileGetRequest{
		ApiService: a,
		ctx: ctx,
		accountID: accountID,
		mediaID: mediaID,
	}
}

// Execute executes the request
//  @return *os.File
func (a *MediaAPIService) V1AccountAccountIDMediaMediaIDFileGetExecute(r ApiV1AccountAccountIDMediaMediaIDFileGetRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MediaAPIService.V1AccountAccountIDMediaMediaIDFileGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/account/{accountID}/media/{mediaID}/file"
	localVarPath = strings.Replace(localVarPath, "{"+"accountID"+"}", url.PathEscape(parameterValueToString(r.accountID, "accountID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mediaID"+"}", url.PathEscape(parameterValueToString(r.mediaID, "mediaID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "audio/mp3", "audio/mpeg", "audio/mpeg3", "audio/x-wav", "audio/wav", "audio/ogg", "video/x-flv", "video/h264", "video/mpeg", "video/quicktime", "video/mp4", "video/webm"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["BearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v UtilCPAASError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiV1AccountAccountIDMediaMediaIDFilePostRequest struct {
	ctx context.Context
	ApiService *MediaAPIService
	accountID string
	mediaID string
	file *os.File
}

// Media file
func (r ApiV1AccountAccountIDMediaMediaIDFilePostRequest) File(file *os.File) ApiV1AccountAccountIDMediaMediaIDFilePostRequest {
	r.file = file
	return r
}

func (r ApiV1AccountAccountIDMediaMediaIDFilePostRequest) Execute() (*ServiceDocsMediaGetSingle, *http.Response, error) {
	return r.ApiService.V1AccountAccountIDMediaMediaIDFilePostExecute(r)
}

/*
V1AccountAccountIDMediaMediaIDFilePost Add Media File

Include a media file that is connected to a media object in an account.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountID Account ID, 32 alpha numeric
 @param mediaID Media ID, 32 alpha numeric
 @return ApiV1AccountAccountIDMediaMediaIDFilePostRequest
*/
func (a *MediaAPIService) V1AccountAccountIDMediaMediaIDFilePost(ctx context.Context, accountID string, mediaID string) ApiV1AccountAccountIDMediaMediaIDFilePostRequest {
	return ApiV1AccountAccountIDMediaMediaIDFilePostRequest{
		ApiService: a,
		ctx: ctx,
		accountID: accountID,
		mediaID: mediaID,
	}
}

// Execute executes the request
//  @return ServiceDocsMediaGetSingle
func (a *MediaAPIService) V1AccountAccountIDMediaMediaIDFilePostExecute(r ApiV1AccountAccountIDMediaMediaIDFilePostRequest) (*ServiceDocsMediaGetSingle, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServiceDocsMediaGetSingle
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MediaAPIService.V1AccountAccountIDMediaMediaIDFilePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/account/{accountID}/media/{mediaID}/file"
	localVarPath = strings.Replace(localVarPath, "{"+"accountID"+"}", url.PathEscape(parameterValueToString(r.accountID, "accountID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mediaID"+"}", url.PathEscape(parameterValueToString(r.mediaID, "mediaID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.file == nil {
		return localVarReturnValue, nil, reportError("file is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"
	fileLocalVarFile := r.file

	if fileLocalVarFile != nil {
		fbs, _ := io.ReadAll(fileLocalVarFile)

		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["BearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v UtilCPAASError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiV1AccountAccountidMediaGetRequest struct {
	ctx context.Context
	ApiService *MediaAPIService
	accountid string
	startKey *string
	pageSize *int32
}

// start_key for pagination that was returned as next_start_key from your previous call
func (r ApiV1AccountAccountidMediaGetRequest) StartKey(startKey string) ApiV1AccountAccountidMediaGetRequest {
	r.startKey = &startKey
	return r
}

// number of records to return, range 1 to 50
func (r ApiV1AccountAccountidMediaGetRequest) PageSize(pageSize int32) ApiV1AccountAccountidMediaGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiV1AccountAccountidMediaGetRequest) Execute() (*ServiceDocsMediaGetAll, *http.Response, error) {
	return r.ApiService.V1AccountAccountidMediaGetExecute(r)
}

/*
V1AccountAccountidMediaGet Get Media List

View all media files for an account in your organization.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountid Account ID, 32 alpha numeric
 @return ApiV1AccountAccountidMediaGetRequest
*/
func (a *MediaAPIService) V1AccountAccountidMediaGet(ctx context.Context, accountid string) ApiV1AccountAccountidMediaGetRequest {
	return ApiV1AccountAccountidMediaGetRequest{
		ApiService: a,
		ctx: ctx,
		accountid: accountid,
	}
}

// Execute executes the request
//  @return ServiceDocsMediaGetAll
func (a *MediaAPIService) V1AccountAccountidMediaGetExecute(r ApiV1AccountAccountidMediaGetRequest) (*ServiceDocsMediaGetAll, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServiceDocsMediaGetAll
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MediaAPIService.V1AccountAccountidMediaGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/account/{accountid}/media"
	localVarPath = strings.Replace(localVarPath, "{"+"accountid"+"}", url.PathEscape(parameterValueToString(r.accountid, "accountid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_key", r.startKey, "", "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize, "", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["BearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v UtilCPAASError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiV1AccountAccountidMediaMediaidDeleteRequest struct {
	ctx context.Context
	ApiService *MediaAPIService
	accountid string
	mediaid string
}

func (r ApiV1AccountAccountidMediaMediaidDeleteRequest) Execute() (*ServiceDocsMediaGetSingle, *http.Response, error) {
	return r.ApiService.V1AccountAccountidMediaMediaidDeleteExecute(r)
}

/*
V1AccountAccountidMediaMediaidDelete Delete Media

Remove a media file that is no longer in use from an account.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountid Account ID, 32 alpha numeric
 @param mediaid Device ID, 32 alpha numeric
 @return ApiV1AccountAccountidMediaMediaidDeleteRequest
*/
func (a *MediaAPIService) V1AccountAccountidMediaMediaidDelete(ctx context.Context, accountid string, mediaid string) ApiV1AccountAccountidMediaMediaidDeleteRequest {
	return ApiV1AccountAccountidMediaMediaidDeleteRequest{
		ApiService: a,
		ctx: ctx,
		accountid: accountid,
		mediaid: mediaid,
	}
}

// Execute executes the request
//  @return ServiceDocsMediaGetSingle
func (a *MediaAPIService) V1AccountAccountidMediaMediaidDeleteExecute(r ApiV1AccountAccountidMediaMediaidDeleteRequest) (*ServiceDocsMediaGetSingle, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServiceDocsMediaGetSingle
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MediaAPIService.V1AccountAccountidMediaMediaidDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/account/{accountid}/media/{mediaid}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountid"+"}", url.PathEscape(parameterValueToString(r.accountid, "accountid")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mediaid"+"}", url.PathEscape(parameterValueToString(r.mediaid, "mediaid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["BearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v UtilCPAASError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiV1AccountAccountidMediaMediaidGetRequest struct {
	ctx context.Context
	ApiService *MediaAPIService
	accountid string
	mediaid string
}

func (r ApiV1AccountAccountidMediaMediaidGetRequest) Execute() (*ServiceDocsMediaGetSingle, *http.Response, error) {
	return r.ApiService.V1AccountAccountidMediaMediaidGetExecute(r)
}

/*
V1AccountAccountidMediaMediaidGet Get Media Details

Permit users to view an account's specific media information.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountid Account ID, 32 alpha numeric
 @param mediaid Media ID, 32 alpha numeric
 @return ApiV1AccountAccountidMediaMediaidGetRequest
*/
func (a *MediaAPIService) V1AccountAccountidMediaMediaidGet(ctx context.Context, accountid string, mediaid string) ApiV1AccountAccountidMediaMediaidGetRequest {
	return ApiV1AccountAccountidMediaMediaidGetRequest{
		ApiService: a,
		ctx: ctx,
		accountid: accountid,
		mediaid: mediaid,
	}
}

// Execute executes the request
//  @return ServiceDocsMediaGetSingle
func (a *MediaAPIService) V1AccountAccountidMediaMediaidGetExecute(r ApiV1AccountAccountidMediaMediaidGetRequest) (*ServiceDocsMediaGetSingle, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServiceDocsMediaGetSingle
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MediaAPIService.V1AccountAccountidMediaMediaidGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/account/{accountid}/media/{mediaid}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountid"+"}", url.PathEscape(parameterValueToString(r.accountid, "accountid")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mediaid"+"}", url.PathEscape(parameterValueToString(r.mediaid, "mediaid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["BearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v UtilCPAASError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiV1AccountAccountidMediaPostRequest struct {
	ctx context.Context
	ApiService *MediaAPIService
	accountid string
	media *ServiceVOIPMediaAddEditData
}

// Media creation or update payload
func (r ApiV1AccountAccountidMediaPostRequest) Media(media ServiceVOIPMediaAddEditData) ApiV1AccountAccountidMediaPostRequest {
	r.media = &media
	return r
}

func (r ApiV1AccountAccountidMediaPostRequest) Execute() (*ServiceDocsMediaGetSingle, *http.Response, error) {
	return r.ApiService.V1AccountAccountidMediaPostExecute(r)
}

/*
V1AccountAccountidMediaPost Create Media

Generate a media object to allow users to upload a media file in an account.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountid Account ID, 32 alpha numeric
 @return ApiV1AccountAccountidMediaPostRequest
*/
func (a *MediaAPIService) V1AccountAccountidMediaPost(ctx context.Context, accountid string) ApiV1AccountAccountidMediaPostRequest {
	return ApiV1AccountAccountidMediaPostRequest{
		ApiService: a,
		ctx: ctx,
		accountid: accountid,
	}
}

// Execute executes the request
//  @return ServiceDocsMediaGetSingle
func (a *MediaAPIService) V1AccountAccountidMediaPostExecute(r ApiV1AccountAccountidMediaPostRequest) (*ServiceDocsMediaGetSingle, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServiceDocsMediaGetSingle
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MediaAPIService.V1AccountAccountidMediaPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/account/{accountid}/media"
	localVarPath = strings.Replace(localVarPath, "{"+"accountid"+"}", url.PathEscape(parameterValueToString(r.accountid, "accountid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.media == nil {
		return localVarReturnValue, nil, reportError("media is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.media
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["BearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v UtilCPAASError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
