/*
 * WSO2 API Manager - Publisher API
 *
 * This document specifies a **RESTful API** for WSO2 **API Manager** - **Publisher**.  # Authentication Our REST APIs are protected using OAuth2 and access control is achieved through scopes. Before you start invoking the the API you need to obtain an access token with the required scopes. This guide will walk you through the steps that you will need to follow to obtain an access token. First you need to obtain the consumer key/secret key pair by calling the dynamic client registration (DCR) endpoint. You can add your preferred grant types in the payload. A Sample payload is shown below. ```   {   \"callbackUrl\":\"www.google.lk\",   \"clientName\":\"rest_api_publisher\",   \"owner\":\"admin\",   \"grantType\":\"client_credentials password refresh_token\",   \"saasApp\":true   } ``` Create a file (payload.json) with the above sample payload, and use the cURL shown bellow to invoke the DCR endpoint. Authorization header of this should contain the base64 encoded admin username and password. **Format of the request** ```   curl -X POST -H \"Authorization: Basic Base64(admin_username:admin_password)\" -H \"Content-Type: application/json\"   \\ -d @payload.json https://<host>:<servlet_port>/client-registration/v0.17/register ``` **Sample request** ```   curl -X POST -H \"Authorization: Basic YWRtaW46YWRtaW4=\" -H \"Content-Type: application/json\"   \\ -d @payload.json https://localhost:9443/client-registration/v0.17/register ``` Following is a sample response after invoking the above curl. ``` { \"clientId\": \"fOCi4vNJ59PpHucC2CAYfYuADdMa\", \"clientName\": \"rest_api_publisher\", \"callBackURL\": \"www.google.lk\", \"clientSecret\": \"a4FwHlq0iCIKVs2MPIIDnepZnYMa\", \"isSaasApplication\": true, \"appOwner\": \"admin\", \"jsonString\": \"{\\\"grant_types\\\":\\\"client_credentials password refresh_token\\\",\\\"redirect_uris\\\":\\\"www.google.lk\\\",\\\"client_name\\\":\\\"rest_api123\\\"}\", \"jsonAppAttribute\": \"{}\", \"tokenType\": null } ``` Next you must use the above client id and secret to obtain the access token. We will be using the password grant type for this, you can use any grant type you desire. You also need to add the proper **scope** when getting the access token. All possible scopes for publisher REST API can be viewed in **OAuth2 Security** section of this document and scope for each resource is given in **authorization** section of resource documentation. Following is the format of the request if you are using the password grant type. ``` curl -k -d \"grant_type=password&username=<admin_username>&password=<admin_passowrd&scope=<scopes seperated by space>\" \\ -H \"Authorization: Basic base64(cliet_id:client_secret)\" \\ https://<host>:<gateway_port>/token ``` **Sample request** ``` curl https://localhost:8243/token -k \\ -H \"Authorization: Basic Zk9DaTR2Tko1OVBwSHVjQzJDQVlmWXVBRGRNYTphNEZ3SGxxMGlDSUtWczJNUElJRG5lcFpuWU1h\" \\ -d \"grant_type=password&username=admin&password=admin&scope=apim:api_view apim:api_create\" ``` Shown below is a sample response to the above request. ``` { \"access_token\": \"e79bda48-3406-3178-acce-f6e4dbdcbb12\", \"refresh_token\": \"a757795d-e69f-38b8-bd85-9aded677a97c\", \"scope\": \"apim:api_create apim:api_view\", \"token_type\": \"Bearer\", \"expires_in\": 3600 } ``` Now you have a valid access token, which you can use to invoke an API. Navigate through the API descriptions to find the required API, obtain an access token as described above and invoke the API with the authentication header. If you use a different authentication mechanism, this process may change.  # Try out in Postman If you want to try-out the embedded postman collection with \"Run in Postman\" option, please follow the guidelines listed below. * All of the OAuth2 secured endpoints have been configured with an Authorization Bearer header with a parameterized access token. Before invoking any REST API resource make sure you run the `Register DCR Application` and `Generate Access Token` requests to fetch an access token with all required scopes. * Make sure you have an API Manager instance up and running. * Update the `basepath` parameter to match the hostname and port of the APIM instance.  [![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/a09044034b5c3c1b01a9) 
 *
 * API version: v1.1
 * Contact: architecture@wso2.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
)

// APISearchResultAllOf struct for APISearchResultAllOf
type APISearchResultAllOf struct {
	// A brief description about the API
	Description *string `json:"description,omitempty"`
	// A string that represents the context of the user's request
	Context *string `json:"context,omitempty"`
	// The version of the API
	Version *string `json:"version,omitempty"`
	// If the provider value is not given, the user invoking the API will be used as the provider. 
	Provider *string `json:"provider,omitempty"`
	// This describes in which status of the lifecycle the API is
	Status *string `json:"status,omitempty"`
	ThumbnailUri *string `json:"thumbnailUri,omitempty"`
}

// NewAPISearchResultAllOf instantiates a new APISearchResultAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPISearchResultAllOf() *APISearchResultAllOf {
	this := APISearchResultAllOf{}
	return &this
}

// NewAPISearchResultAllOfWithDefaults instantiates a new APISearchResultAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPISearchResultAllOfWithDefaults() *APISearchResultAllOf {
	this := APISearchResultAllOf{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *APISearchResultAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APISearchResultAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *APISearchResultAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *APISearchResultAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *APISearchResultAllOf) GetContext() string {
	if o == nil || o.Context == nil {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APISearchResultAllOf) GetContextOk() (*string, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *APISearchResultAllOf) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *APISearchResultAllOf) SetContext(v string) {
	o.Context = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *APISearchResultAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APISearchResultAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *APISearchResultAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *APISearchResultAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *APISearchResultAllOf) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APISearchResultAllOf) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *APISearchResultAllOf) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *APISearchResultAllOf) SetProvider(v string) {
	o.Provider = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *APISearchResultAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APISearchResultAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *APISearchResultAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *APISearchResultAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetThumbnailUri returns the ThumbnailUri field value if set, zero value otherwise.
func (o *APISearchResultAllOf) GetThumbnailUri() string {
	if o == nil || o.ThumbnailUri == nil {
		var ret string
		return ret
	}
	return *o.ThumbnailUri
}

// GetThumbnailUriOk returns a tuple with the ThumbnailUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APISearchResultAllOf) GetThumbnailUriOk() (*string, bool) {
	if o == nil || o.ThumbnailUri == nil {
		return nil, false
	}
	return o.ThumbnailUri, true
}

// HasThumbnailUri returns a boolean if a field has been set.
func (o *APISearchResultAllOf) HasThumbnailUri() bool {
	if o != nil && o.ThumbnailUri != nil {
		return true
	}

	return false
}

// SetThumbnailUri gets a reference to the given string and assigns it to the ThumbnailUri field.
func (o *APISearchResultAllOf) SetThumbnailUri(v string) {
	o.ThumbnailUri = &v
}

func (o APISearchResultAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ThumbnailUri != nil {
		toSerialize["thumbnailUri"] = o.ThumbnailUri
	}
	return json.Marshal(toSerialize)
}

type NullableAPISearchResultAllOf struct {
	value *APISearchResultAllOf
	isSet bool
}

func (v NullableAPISearchResultAllOf) Get() *APISearchResultAllOf {
	return v.value
}

func (v *NullableAPISearchResultAllOf) Set(val *APISearchResultAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAPISearchResultAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAPISearchResultAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPISearchResultAllOf(val *APISearchResultAllOf) *NullableAPISearchResultAllOf {
	return &NullableAPISearchResultAllOf{value: val, isSet: true}
}

func (v NullableAPISearchResultAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPISearchResultAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


