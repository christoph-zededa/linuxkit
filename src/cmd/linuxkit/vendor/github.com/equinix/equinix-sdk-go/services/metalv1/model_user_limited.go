/*
Metal API

# Introduction Equinix Metal provides a RESTful HTTP API which can be reached at <https://api.equinix.com/metal/v1>. This document describes the API and how to use it.  The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account. Every feature of the Equinix Metal web interface is accessible through the API.  The API docs are generated from the Equinix Metal OpenAPI specification and are officially hosted at <https://metal.equinix.com/developers/api>.  # Common Parameters  The Equinix Metal API uses a few methods to minimize network traffic and improve throughput. These parameters are not used in all API calls, but are used often enough to warrant their own section. Look for these parameters in the documentation for the API calls that support them.  ## Pagination  Pagination is used to limit the number of results returned in a single request. The API will return a maximum of 100 results per page. To retrieve additional results, you can use the `page` and `per_page` query parameters.  The `page` parameter is used to specify the page number. The first page is `1`. The `per_page` parameter is used to specify the number of results per page. The maximum number of results differs by resource type.  ## Sorting  Where offered, the API allows you to sort results by a specific field. To sort results use the `sort_by` query parameter with the root level field name as the value. The `sort_direction` parameter is used to specify the sort direction, either either `asc` (ascending) or `desc` (descending).  ## Filtering  Filtering is used to limit the results returned in a single request. The API supports filtering by certain fields in the response. To filter results, you can use the field as a query parameter.  For example, to filter the IP list to only return public IPv4 addresses, you can filter by the `type` field, as in the following request:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/projects/id/ips?type=public_ipv4 ```  Only IP addresses with the `type` field set to `public_ipv4` will be returned.  ## Searching  Searching is used to find matching resources using multiple field comparissons. The API supports searching in resources that define this behavior. Currently the search parameter is only available on devices, ssh_keys, api_keys and memberships endpoints.  To search resources you can use the `search` query parameter.  ## Include and Exclude  For resources that contain references to other resources, sucha as a Device that refers to the Project it resides in, the Equinix Metal API will returns `href` values (API links) to the associated resource.  ```json {   ...   \"project\": {     \"href\": \"/metal/v1/projects/f3f131c8-f302-49ef-8c44-9405022dc6dd\"   } } ```  If you're going need the project details, you can avoid a second API request.  Specify the contained `href` resources and collections that you'd like to have included in the response using the `include` query parameter.  For example:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=projects ```  The `include` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests where `href` resources are presented.  To have multiple resources include, use a comma-separated list (e.g. `?include=emails,projects,memberships`).  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=emails,projects,memberships ```  You may also include nested associations up to three levels deep using dot notation (`?include=memberships.projects`):  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=memberships.projects ```  To exclude resources, and optimize response delivery, use the `exclude` query parameter. The `exclude` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests for fields with nested object responses. When excluded, these fields will be replaced with an object that contains only an `href` field.

API version: 1.0.0
Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
	"fmt"
)

// checks if the UserLimited type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserLimited{}

// UserLimited struct for UserLimited
type UserLimited struct {
	// Avatar thumbnail URL of the User
	AvatarThumbUrl *string `json:"avatar_thumb_url,omitempty"`
	// Avatar URL of the User
	AvatarUrl *string `json:"avatar_url,omitempty"`
	// Full name of the User
	FullName *string `json:"full_name,omitempty"`
	// API URL uniquely representing the User
	Href *string `json:"href,omitempty"`
	// ID of the User
	Id                   string `json:"id"`
	AdditionalProperties map[string]interface{}
}

type _UserLimited UserLimited

// NewUserLimited instantiates a new UserLimited object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserLimited(id string) *UserLimited {
	this := UserLimited{}
	this.Id = id
	return &this
}

// NewUserLimitedWithDefaults instantiates a new UserLimited object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserLimitedWithDefaults() *UserLimited {
	this := UserLimited{}
	return &this
}

// GetAvatarThumbUrl returns the AvatarThumbUrl field value if set, zero value otherwise.
func (o *UserLimited) GetAvatarThumbUrl() string {
	if o == nil || IsNil(o.AvatarThumbUrl) {
		var ret string
		return ret
	}
	return *o.AvatarThumbUrl
}

// GetAvatarThumbUrlOk returns a tuple with the AvatarThumbUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLimited) GetAvatarThumbUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AvatarThumbUrl) {
		return nil, false
	}
	return o.AvatarThumbUrl, true
}

// HasAvatarThumbUrl returns a boolean if a field has been set.
func (o *UserLimited) HasAvatarThumbUrl() bool {
	if o != nil && !IsNil(o.AvatarThumbUrl) {
		return true
	}

	return false
}

// SetAvatarThumbUrl gets a reference to the given string and assigns it to the AvatarThumbUrl field.
func (o *UserLimited) SetAvatarThumbUrl(v string) {
	o.AvatarThumbUrl = &v
}

// GetAvatarUrl returns the AvatarUrl field value if set, zero value otherwise.
func (o *UserLimited) GetAvatarUrl() string {
	if o == nil || IsNil(o.AvatarUrl) {
		var ret string
		return ret
	}
	return *o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLimited) GetAvatarUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AvatarUrl) {
		return nil, false
	}
	return o.AvatarUrl, true
}

// HasAvatarUrl returns a boolean if a field has been set.
func (o *UserLimited) HasAvatarUrl() bool {
	if o != nil && !IsNil(o.AvatarUrl) {
		return true
	}

	return false
}

// SetAvatarUrl gets a reference to the given string and assigns it to the AvatarUrl field.
func (o *UserLimited) SetAvatarUrl(v string) {
	o.AvatarUrl = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *UserLimited) GetFullName() string {
	if o == nil || IsNil(o.FullName) {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLimited) GetFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.FullName) {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *UserLimited) HasFullName() bool {
	if o != nil && !IsNil(o.FullName) {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *UserLimited) SetFullName(v string) {
	o.FullName = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *UserLimited) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLimited) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *UserLimited) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *UserLimited) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value
func (o *UserLimited) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserLimited) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserLimited) SetId(v string) {
	o.Id = v
}

func (o UserLimited) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserLimited) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AvatarThumbUrl) {
		toSerialize["avatar_thumb_url"] = o.AvatarThumbUrl
	}
	if !IsNil(o.AvatarUrl) {
		toSerialize["avatar_url"] = o.AvatarUrl
	}
	if !IsNil(o.FullName) {
		toSerialize["full_name"] = o.FullName
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	toSerialize["id"] = o.Id

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserLimited) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUserLimited := _UserLimited{}

	err = json.Unmarshal(data, &varUserLimited)

	if err != nil {
		return err
	}

	*o = UserLimited(varUserLimited)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "avatar_thumb_url")
		delete(additionalProperties, "avatar_url")
		delete(additionalProperties, "full_name")
		delete(additionalProperties, "href")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserLimited struct {
	value *UserLimited
	isSet bool
}

func (v NullableUserLimited) Get() *UserLimited {
	return v.value
}

func (v *NullableUserLimited) Set(val *UserLimited) {
	v.value = val
	v.isSet = true
}

func (v NullableUserLimited) IsSet() bool {
	return v.isSet
}

func (v *NullableUserLimited) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserLimited(val *UserLimited) *NullableUserLimited {
	return &NullableUserLimited{value: val, isSet: true}
}

func (v NullableUserLimited) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserLimited) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}