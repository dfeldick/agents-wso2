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

// SharedScopeUsedAPIInfo struct for SharedScopeUsedAPIInfo
type SharedScopeUsedAPIInfo struct {
	Name string `json:"name"`
	Context string `json:"context"`
	Version string `json:"version"`
	// If the provider value is not given user invoking the api will be used as the provider. 
	Provider *string `json:"provider,omitempty"`
	// Resource list which have used the shared scope within this API 
	UsedResourceList *[]SharedScopeUsedAPIResourceInfo `json:"usedResourceList,omitempty"`
}

// NewSharedScopeUsedAPIInfo instantiates a new SharedScopeUsedAPIInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharedScopeUsedAPIInfo(name string, context string, version string, ) *SharedScopeUsedAPIInfo {
	this := SharedScopeUsedAPIInfo{}
	this.Name = name
	this.Context = context
	this.Version = version
	return &this
}

// NewSharedScopeUsedAPIInfoWithDefaults instantiates a new SharedScopeUsedAPIInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharedScopeUsedAPIInfoWithDefaults() *SharedScopeUsedAPIInfo {
	this := SharedScopeUsedAPIInfo{}
	return &this
}

// GetName returns the Name field value
func (o *SharedScopeUsedAPIInfo) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SharedScopeUsedAPIInfo) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SharedScopeUsedAPIInfo) SetName(v string) {
	o.Name = v
}

// GetContext returns the Context field value
func (o *SharedScopeUsedAPIInfo) GetContext() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *SharedScopeUsedAPIInfo) GetContextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *SharedScopeUsedAPIInfo) SetContext(v string) {
	o.Context = v
}

// GetVersion returns the Version field value
func (o *SharedScopeUsedAPIInfo) GetVersion() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *SharedScopeUsedAPIInfo) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *SharedScopeUsedAPIInfo) SetVersion(v string) {
	o.Version = v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *SharedScopeUsedAPIInfo) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedScopeUsedAPIInfo) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *SharedScopeUsedAPIInfo) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *SharedScopeUsedAPIInfo) SetProvider(v string) {
	o.Provider = &v
}

// GetUsedResourceList returns the UsedResourceList field value if set, zero value otherwise.
func (o *SharedScopeUsedAPIInfo) GetUsedResourceList() []SharedScopeUsedAPIResourceInfo {
	if o == nil || o.UsedResourceList == nil {
		var ret []SharedScopeUsedAPIResourceInfo
		return ret
	}
	return *o.UsedResourceList
}

// GetUsedResourceListOk returns a tuple with the UsedResourceList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedScopeUsedAPIInfo) GetUsedResourceListOk() (*[]SharedScopeUsedAPIResourceInfo, bool) {
	if o == nil || o.UsedResourceList == nil {
		return nil, false
	}
	return o.UsedResourceList, true
}

// HasUsedResourceList returns a boolean if a field has been set.
func (o *SharedScopeUsedAPIInfo) HasUsedResourceList() bool {
	if o != nil && o.UsedResourceList != nil {
		return true
	}

	return false
}

// SetUsedResourceList gets a reference to the given []SharedScopeUsedAPIResourceInfo and assigns it to the UsedResourceList field.
func (o *SharedScopeUsedAPIInfo) SetUsedResourceList(v []SharedScopeUsedAPIResourceInfo) {
	o.UsedResourceList = &v
}

func (o SharedScopeUsedAPIInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["context"] = o.Context
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.UsedResourceList != nil {
		toSerialize["usedResourceList"] = o.UsedResourceList
	}
	return json.Marshal(toSerialize)
}

type NullableSharedScopeUsedAPIInfo struct {
	value *SharedScopeUsedAPIInfo
	isSet bool
}

func (v NullableSharedScopeUsedAPIInfo) Get() *SharedScopeUsedAPIInfo {
	return v.value
}

func (v *NullableSharedScopeUsedAPIInfo) Set(val *SharedScopeUsedAPIInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSharedScopeUsedAPIInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSharedScopeUsedAPIInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharedScopeUsedAPIInfo(val *SharedScopeUsedAPIInfo) *NullableSharedScopeUsedAPIInfo {
	return &NullableSharedScopeUsedAPIInfo{value: val, isSet: true}
}

func (v NullableSharedScopeUsedAPIInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharedScopeUsedAPIInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


