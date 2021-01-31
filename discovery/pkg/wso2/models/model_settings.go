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

// Settings struct for Settings
type Settings struct {
	// Store URL
	StoreUrl *string `json:"storeUrl,omitempty"`
	Environment *[]Environment `json:"environment,omitempty"`
	Scopes *[]string `json:"scopes,omitempty"`
	MonetizationAttributes *[]MonetizationAttribute `json:"monetizationAttributes,omitempty"`
	SecurityAuditProperties *map[string]interface{} `json:"securityAuditProperties,omitempty"`
	// Is External Stores configuration enabled 
	ExternalStoresEnabled *bool `json:"externalStoresEnabled,omitempty"`
	// Is Document Visibility configuration enabled 
	DocVisibilityEnabled *bool `json:"docVisibilityEnabled,omitempty"`
	// Is Cross Tenant Subscriptions Enabled 
	CrossTenantSubscriptionEnabled *bool `json:"crossTenantSubscriptionEnabled,omitempty"`
	Deployments *[]Deployments `json:"deployments,omitempty"`
}

// NewSettings instantiates a new Settings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettings() *Settings {
	this := Settings{}
	var crossTenantSubscriptionEnabled bool = false
	this.CrossTenantSubscriptionEnabled = &crossTenantSubscriptionEnabled
	return &this
}

// NewSettingsWithDefaults instantiates a new Settings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingsWithDefaults() *Settings {
	this := Settings{}
	var crossTenantSubscriptionEnabled bool = false
	this.CrossTenantSubscriptionEnabled = &crossTenantSubscriptionEnabled
	return &this
}

// GetStoreUrl returns the StoreUrl field value if set, zero value otherwise.
func (o *Settings) GetStoreUrl() string {
	if o == nil || o.StoreUrl == nil {
		var ret string
		return ret
	}
	return *o.StoreUrl
}

// GetStoreUrlOk returns a tuple with the StoreUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Settings) GetStoreUrlOk() (*string, bool) {
	if o == nil || o.StoreUrl == nil {
		return nil, false
	}
	return o.StoreUrl, true
}

// HasStoreUrl returns a boolean if a field has been set.
func (o *Settings) HasStoreUrl() bool {
	if o != nil && o.StoreUrl != nil {
		return true
	}

	return false
}

// SetStoreUrl gets a reference to the given string and assigns it to the StoreUrl field.
func (o *Settings) SetStoreUrl(v string) {
	o.StoreUrl = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *Settings) GetEnvironment() []Environment {
	if o == nil || o.Environment == nil {
		var ret []Environment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Settings) GetEnvironmentOk() (*[]Environment, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *Settings) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given []Environment and assigns it to the Environment field.
func (o *Settings) SetEnvironment(v []Environment) {
	o.Environment = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *Settings) GetScopes() []string {
	if o == nil || o.Scopes == nil {
		var ret []string
		return ret
	}
	return *o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Settings) GetScopesOk() (*[]string, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *Settings) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *Settings) SetScopes(v []string) {
	o.Scopes = &v
}

// GetMonetizationAttributes returns the MonetizationAttributes field value if set, zero value otherwise.
func (o *Settings) GetMonetizationAttributes() []MonetizationAttribute {
	if o == nil || o.MonetizationAttributes == nil {
		var ret []MonetizationAttribute
		return ret
	}
	return *o.MonetizationAttributes
}

// GetMonetizationAttributesOk returns a tuple with the MonetizationAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Settings) GetMonetizationAttributesOk() (*[]MonetizationAttribute, bool) {
	if o == nil || o.MonetizationAttributes == nil {
		return nil, false
	}
	return o.MonetizationAttributes, true
}

// HasMonetizationAttributes returns a boolean if a field has been set.
func (o *Settings) HasMonetizationAttributes() bool {
	if o != nil && o.MonetizationAttributes != nil {
		return true
	}

	return false
}

// SetMonetizationAttributes gets a reference to the given []MonetizationAttribute and assigns it to the MonetizationAttributes field.
func (o *Settings) SetMonetizationAttributes(v []MonetizationAttribute) {
	o.MonetizationAttributes = &v
}

// GetSecurityAuditProperties returns the SecurityAuditProperties field value if set, zero value otherwise.
func (o *Settings) GetSecurityAuditProperties() map[string]interface{} {
	if o == nil || o.SecurityAuditProperties == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.SecurityAuditProperties
}

// GetSecurityAuditPropertiesOk returns a tuple with the SecurityAuditProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Settings) GetSecurityAuditPropertiesOk() (*map[string]interface{}, bool) {
	if o == nil || o.SecurityAuditProperties == nil {
		return nil, false
	}
	return o.SecurityAuditProperties, true
}

// HasSecurityAuditProperties returns a boolean if a field has been set.
func (o *Settings) HasSecurityAuditProperties() bool {
	if o != nil && o.SecurityAuditProperties != nil {
		return true
	}

	return false
}

// SetSecurityAuditProperties gets a reference to the given map[string]interface{} and assigns it to the SecurityAuditProperties field.
func (o *Settings) SetSecurityAuditProperties(v map[string]interface{}) {
	o.SecurityAuditProperties = &v
}

// GetExternalStoresEnabled returns the ExternalStoresEnabled field value if set, zero value otherwise.
func (o *Settings) GetExternalStoresEnabled() bool {
	if o == nil || o.ExternalStoresEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ExternalStoresEnabled
}

// GetExternalStoresEnabledOk returns a tuple with the ExternalStoresEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Settings) GetExternalStoresEnabledOk() (*bool, bool) {
	if o == nil || o.ExternalStoresEnabled == nil {
		return nil, false
	}
	return o.ExternalStoresEnabled, true
}

// HasExternalStoresEnabled returns a boolean if a field has been set.
func (o *Settings) HasExternalStoresEnabled() bool {
	if o != nil && o.ExternalStoresEnabled != nil {
		return true
	}

	return false
}

// SetExternalStoresEnabled gets a reference to the given bool and assigns it to the ExternalStoresEnabled field.
func (o *Settings) SetExternalStoresEnabled(v bool) {
	o.ExternalStoresEnabled = &v
}

// GetDocVisibilityEnabled returns the DocVisibilityEnabled field value if set, zero value otherwise.
func (o *Settings) GetDocVisibilityEnabled() bool {
	if o == nil || o.DocVisibilityEnabled == nil {
		var ret bool
		return ret
	}
	return *o.DocVisibilityEnabled
}

// GetDocVisibilityEnabledOk returns a tuple with the DocVisibilityEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Settings) GetDocVisibilityEnabledOk() (*bool, bool) {
	if o == nil || o.DocVisibilityEnabled == nil {
		return nil, false
	}
	return o.DocVisibilityEnabled, true
}

// HasDocVisibilityEnabled returns a boolean if a field has been set.
func (o *Settings) HasDocVisibilityEnabled() bool {
	if o != nil && o.DocVisibilityEnabled != nil {
		return true
	}

	return false
}

// SetDocVisibilityEnabled gets a reference to the given bool and assigns it to the DocVisibilityEnabled field.
func (o *Settings) SetDocVisibilityEnabled(v bool) {
	o.DocVisibilityEnabled = &v
}

// GetCrossTenantSubscriptionEnabled returns the CrossTenantSubscriptionEnabled field value if set, zero value otherwise.
func (o *Settings) GetCrossTenantSubscriptionEnabled() bool {
	if o == nil || o.CrossTenantSubscriptionEnabled == nil {
		var ret bool
		return ret
	}
	return *o.CrossTenantSubscriptionEnabled
}

// GetCrossTenantSubscriptionEnabledOk returns a tuple with the CrossTenantSubscriptionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Settings) GetCrossTenantSubscriptionEnabledOk() (*bool, bool) {
	if o == nil || o.CrossTenantSubscriptionEnabled == nil {
		return nil, false
	}
	return o.CrossTenantSubscriptionEnabled, true
}

// HasCrossTenantSubscriptionEnabled returns a boolean if a field has been set.
func (o *Settings) HasCrossTenantSubscriptionEnabled() bool {
	if o != nil && o.CrossTenantSubscriptionEnabled != nil {
		return true
	}

	return false
}

// SetCrossTenantSubscriptionEnabled gets a reference to the given bool and assigns it to the CrossTenantSubscriptionEnabled field.
func (o *Settings) SetCrossTenantSubscriptionEnabled(v bool) {
	o.CrossTenantSubscriptionEnabled = &v
}

// GetDeployments returns the Deployments field value if set, zero value otherwise.
func (o *Settings) GetDeployments() []Deployments {
	if o == nil || o.Deployments == nil {
		var ret []Deployments
		return ret
	}
	return *o.Deployments
}

// GetDeploymentsOk returns a tuple with the Deployments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Settings) GetDeploymentsOk() (*[]Deployments, bool) {
	if o == nil || o.Deployments == nil {
		return nil, false
	}
	return o.Deployments, true
}

// HasDeployments returns a boolean if a field has been set.
func (o *Settings) HasDeployments() bool {
	if o != nil && o.Deployments != nil {
		return true
	}

	return false
}

// SetDeployments gets a reference to the given []Deployments and assigns it to the Deployments field.
func (o *Settings) SetDeployments(v []Deployments) {
	o.Deployments = &v
}

func (o Settings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StoreUrl != nil {
		toSerialize["storeUrl"] = o.StoreUrl
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	if o.MonetizationAttributes != nil {
		toSerialize["monetizationAttributes"] = o.MonetizationAttributes
	}
	if o.SecurityAuditProperties != nil {
		toSerialize["securityAuditProperties"] = o.SecurityAuditProperties
	}
	if o.ExternalStoresEnabled != nil {
		toSerialize["externalStoresEnabled"] = o.ExternalStoresEnabled
	}
	if o.DocVisibilityEnabled != nil {
		toSerialize["docVisibilityEnabled"] = o.DocVisibilityEnabled
	}
	if o.CrossTenantSubscriptionEnabled != nil {
		toSerialize["crossTenantSubscriptionEnabled"] = o.CrossTenantSubscriptionEnabled
	}
	if o.Deployments != nil {
		toSerialize["deployments"] = o.Deployments
	}
	return json.Marshal(toSerialize)
}

type NullableSettings struct {
	value *Settings
	isSet bool
}

func (v NullableSettings) Get() *Settings {
	return v.value
}

func (v *NullableSettings) Set(val *Settings) {
	v.value = val
	v.isSet = true
}

func (v NullableSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettings(val *Settings) *NullableSettings {
	return &NullableSettings{value: val, isSet: true}
}

func (v NullableSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


