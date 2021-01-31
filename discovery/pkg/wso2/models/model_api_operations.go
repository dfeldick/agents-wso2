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

// APIOperations struct for APIOperations
type APIOperations struct {
	Id *string `json:"id,omitempty"`
	Target *string `json:"target,omitempty"`
	Verb *string `json:"verb,omitempty"`
	AuthType *string `json:"authType,omitempty"`
	ThrottlingPolicy *string `json:"throttlingPolicy,omitempty"`
	Scopes *[]string `json:"scopes,omitempty"`
	UsedProductIds *[]string `json:"usedProductIds,omitempty"`
	AmznResourceName *string `json:"amznResourceName,omitempty"`
	AmznResourceTimeout *int32 `json:"amznResourceTimeout,omitempty"`
}

// NewAPIOperations instantiates a new APIOperations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIOperations() *APIOperations {
	this := APIOperations{}
	var authType string = "Any"
	this.AuthType = &authType
	return &this
}

// NewAPIOperationsWithDefaults instantiates a new APIOperations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIOperationsWithDefaults() *APIOperations {
	this := APIOperations{}
	var authType string = "Any"
	this.AuthType = &authType
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *APIOperations) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIOperations) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *APIOperations) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *APIOperations) SetId(v string) {
	o.Id = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *APIOperations) GetTarget() string {
	if o == nil || o.Target == nil {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIOperations) GetTargetOk() (*string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *APIOperations) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *APIOperations) SetTarget(v string) {
	o.Target = &v
}

// GetVerb returns the Verb field value if set, zero value otherwise.
func (o *APIOperations) GetVerb() string {
	if o == nil || o.Verb == nil {
		var ret string
		return ret
	}
	return *o.Verb
}

// GetVerbOk returns a tuple with the Verb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIOperations) GetVerbOk() (*string, bool) {
	if o == nil || o.Verb == nil {
		return nil, false
	}
	return o.Verb, true
}

// HasVerb returns a boolean if a field has been set.
func (o *APIOperations) HasVerb() bool {
	if o != nil && o.Verb != nil {
		return true
	}

	return false
}

// SetVerb gets a reference to the given string and assigns it to the Verb field.
func (o *APIOperations) SetVerb(v string) {
	o.Verb = &v
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *APIOperations) GetAuthType() string {
	if o == nil || o.AuthType == nil {
		var ret string
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIOperations) GetAuthTypeOk() (*string, bool) {
	if o == nil || o.AuthType == nil {
		return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *APIOperations) HasAuthType() bool {
	if o != nil && o.AuthType != nil {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given string and assigns it to the AuthType field.
func (o *APIOperations) SetAuthType(v string) {
	o.AuthType = &v
}

// GetThrottlingPolicy returns the ThrottlingPolicy field value if set, zero value otherwise.
func (o *APIOperations) GetThrottlingPolicy() string {
	if o == nil || o.ThrottlingPolicy == nil {
		var ret string
		return ret
	}
	return *o.ThrottlingPolicy
}

// GetThrottlingPolicyOk returns a tuple with the ThrottlingPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIOperations) GetThrottlingPolicyOk() (*string, bool) {
	if o == nil || o.ThrottlingPolicy == nil {
		return nil, false
	}
	return o.ThrottlingPolicy, true
}

// HasThrottlingPolicy returns a boolean if a field has been set.
func (o *APIOperations) HasThrottlingPolicy() bool {
	if o != nil && o.ThrottlingPolicy != nil {
		return true
	}

	return false
}

// SetThrottlingPolicy gets a reference to the given string and assigns it to the ThrottlingPolicy field.
func (o *APIOperations) SetThrottlingPolicy(v string) {
	o.ThrottlingPolicy = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *APIOperations) GetScopes() []string {
	if o == nil || o.Scopes == nil {
		var ret []string
		return ret
	}
	return *o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIOperations) GetScopesOk() (*[]string, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *APIOperations) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *APIOperations) SetScopes(v []string) {
	o.Scopes = &v
}

// GetUsedProductIds returns the UsedProductIds field value if set, zero value otherwise.
func (o *APIOperations) GetUsedProductIds() []string {
	if o == nil || o.UsedProductIds == nil {
		var ret []string
		return ret
	}
	return *o.UsedProductIds
}

// GetUsedProductIdsOk returns a tuple with the UsedProductIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIOperations) GetUsedProductIdsOk() (*[]string, bool) {
	if o == nil || o.UsedProductIds == nil {
		return nil, false
	}
	return o.UsedProductIds, true
}

// HasUsedProductIds returns a boolean if a field has been set.
func (o *APIOperations) HasUsedProductIds() bool {
	if o != nil && o.UsedProductIds != nil {
		return true
	}

	return false
}

// SetUsedProductIds gets a reference to the given []string and assigns it to the UsedProductIds field.
func (o *APIOperations) SetUsedProductIds(v []string) {
	o.UsedProductIds = &v
}

// GetAmznResourceName returns the AmznResourceName field value if set, zero value otherwise.
func (o *APIOperations) GetAmznResourceName() string {
	if o == nil || o.AmznResourceName == nil {
		var ret string
		return ret
	}
	return *o.AmznResourceName
}

// GetAmznResourceNameOk returns a tuple with the AmznResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIOperations) GetAmznResourceNameOk() (*string, bool) {
	if o == nil || o.AmznResourceName == nil {
		return nil, false
	}
	return o.AmznResourceName, true
}

// HasAmznResourceName returns a boolean if a field has been set.
func (o *APIOperations) HasAmznResourceName() bool {
	if o != nil && o.AmznResourceName != nil {
		return true
	}

	return false
}

// SetAmznResourceName gets a reference to the given string and assigns it to the AmznResourceName field.
func (o *APIOperations) SetAmznResourceName(v string) {
	o.AmznResourceName = &v
}

// GetAmznResourceTimeout returns the AmznResourceTimeout field value if set, zero value otherwise.
func (o *APIOperations) GetAmznResourceTimeout() int32 {
	if o == nil || o.AmznResourceTimeout == nil {
		var ret int32
		return ret
	}
	return *o.AmznResourceTimeout
}

// GetAmznResourceTimeoutOk returns a tuple with the AmznResourceTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIOperations) GetAmznResourceTimeoutOk() (*int32, bool) {
	if o == nil || o.AmznResourceTimeout == nil {
		return nil, false
	}
	return o.AmznResourceTimeout, true
}

// HasAmznResourceTimeout returns a boolean if a field has been set.
func (o *APIOperations) HasAmznResourceTimeout() bool {
	if o != nil && o.AmznResourceTimeout != nil {
		return true
	}

	return false
}

// SetAmznResourceTimeout gets a reference to the given int32 and assigns it to the AmznResourceTimeout field.
func (o *APIOperations) SetAmznResourceTimeout(v int32) {
	o.AmznResourceTimeout = &v
}

func (o APIOperations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.Verb != nil {
		toSerialize["verb"] = o.Verb
	}
	if o.AuthType != nil {
		toSerialize["authType"] = o.AuthType
	}
	if o.ThrottlingPolicy != nil {
		toSerialize["throttlingPolicy"] = o.ThrottlingPolicy
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	if o.UsedProductIds != nil {
		toSerialize["usedProductIds"] = o.UsedProductIds
	}
	if o.AmznResourceName != nil {
		toSerialize["amznResourceName"] = o.AmznResourceName
	}
	if o.AmznResourceTimeout != nil {
		toSerialize["amznResourceTimeout"] = o.AmznResourceTimeout
	}
	return json.Marshal(toSerialize)
}

type NullableAPIOperations struct {
	value *APIOperations
	isSet bool
}

func (v NullableAPIOperations) Get() *APIOperations {
	return v.value
}

func (v *NullableAPIOperations) Set(val *APIOperations) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIOperations) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIOperations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIOperations(val *APIOperations) *NullableAPIOperations {
	return &NullableAPIOperations{value: val, isSet: true}
}

func (v NullableAPIOperations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIOperations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


