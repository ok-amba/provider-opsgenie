/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AlertPolicyObservation struct {

	// The ID of the Opsgenie Alert Policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AlertPolicyParameters struct {

	// Actions to add to the alerts original actions value as a list of strings. If ignore_original_actions field is set to true, this will replace the original actions.
	// +kubebuilder:validation:Optional
	Actions []*string `json:"actions,omitempty" tf:"actions,omitempty"`

	// Description of the alert. You can use {{description}} to refer to the original alert description. Default: {{description}}
	// +kubebuilder:validation:Optional
	AlertDescription *string `json:"alertDescription,omitempty" tf:"alert_description,omitempty"`

	// Alias of the alert. You can use {{alias}} to refer to the original alias. Default: {{alias}}
	// +kubebuilder:validation:Optional
	Alias *string `json:"alias,omitempty" tf:"alias,omitempty"`

	// It will trigger other modify policies if set to true. Default: false
	// +kubebuilder:validation:Optional
	ContinuePolicy *bool `json:"continuePolicy,omitempty" tf:"continue_policy,omitempty"`

	// If policy should be enabled. Default: true
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Entity field of the alert. You can use {{entity}} to refer to the original entity. Default: {{entity}}
	// +kubebuilder:validation:Optional
	Entity *string `json:"entity,omitempty" tf:"entity,omitempty"`

	// A alert filter which will be applied. This filter can be empty: filter {} - this means match-all. This is a block, structure is documented below.
	// +kubebuilder:validation:Optional
	Filter []FilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// If set to true, policy will ignore the original actions of the alert. Default: false
	// +kubebuilder:validation:Optional
	IgnoreOriginalActions *bool `json:"ignoreOriginalActions,omitempty" tf:"ignore_original_actions,omitempty"`

	// If set to true, policy will ignore the original details of the alert. Default: false
	// +kubebuilder:validation:Optional
	IgnoreOriginalDetails *bool `json:"ignoreOriginalDetails,omitempty" tf:"ignore_original_details,omitempty"`

	// If set to true, policy will ignore the original responders of the alert. Default: false
	// +kubebuilder:validation:Optional
	IgnoreOriginalResponders *bool `json:"ignoreOriginalResponders,omitempty" tf:"ignore_original_responders,omitempty"`

	// If set to true, policy will ignore the original tags of the alert. Default: false
	// +kubebuilder:validation:Optional
	IgnoreOriginalTags *bool `json:"ignoreOriginalTags,omitempty" tf:"ignore_original_tags,omitempty"`

	// Message of the alerts
	// +kubebuilder:validation:Required
	Message *string `json:"message" tf:"message,omitempty"`

	// Name of the alert policy
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Description of the policy. This can be max 512 characters.
	// +kubebuilder:validation:Optional
	PolicyDescription *string `json:"policyDescription,omitempty" tf:"policy_description,omitempty"`

	// Priority of the alert. Should be one of P1, P2, P3, P4, or P5
	// +kubebuilder:validation:Optional
	Priority *string `json:"priority,omitempty" tf:"priority,omitempty"`

	// Responders to add to the alerts original responders value as a list of teams, users or the reserved word none or all. If ignore_original_responders field is set to true, this will replace the original responders. The possible values for responders are: user, team. This is a block, structure is documented below.
	// +kubebuilder:validation:Optional
	Responders []RespondersParameters `json:"responders,omitempty" tf:"responders,omitempty"`

	// Source field of the alert. You can use {{source}} to refer to the original source. Default: {{source}}
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// Tags to add to the alerts original tags value as a list of strings. If ignore_original_responders field is set to true, this will replace the original responders.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Id of team that this policy belongs to.
	// +crossplane:generate:reference:type=Team
	// +kubebuilder:validation:Optional
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`

	// Reference to a Team to populate teamId.
	// +kubebuilder:validation:Optional
	TeamIDRef *v1.Reference `json:"teamIdRef,omitempty" tf:"-"`

	// Selector for a Team to populate teamId.
	// +kubebuilder:validation:Optional
	TeamIDSelector *v1.Selector `json:"teamIdSelector,omitempty" tf:"-"`

	// Time restrictions specified in this field must be met for this policy to work. This is a block, structure is documented below.
	// +kubebuilder:validation:Optional
	TimeRestriction []TimeRestrictionParameters `json:"timeRestriction,omitempty" tf:"time_restriction,omitempty"`
}

type ConditionsObservation struct {
}

type ConditionsParameters struct {

	// User defined value that will be compared with alert field according to the operation. Default: empty string
	// User defined value that will be compared with alert field according to the operation. Default value is empty string
	// +kubebuilder:validation:Optional
	ExpectedValue *string `json:"expectedValue,omitempty" tf:"expected_value,omitempty"`

	// Specifies which alert field will be used in condition. Possible values are message, alias, description, source, entity, tags, actions, details, extra-properties, responders, teams, priority
	// +kubebuilder:validation:Required
	Field *string `json:"field" tf:"field,omitempty"`

	// If field is set as extra-properties, key could be used for key-value pair
	// If 'field' is set as 'extra-properties', key could be used for key-value pair
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Indicates behaviour of the given operation. Default: false
	// Indicates behaviour of the given operation. Default value is false
	// +kubebuilder:validation:Optional
	Not *bool `json:"not,omitempty" tf:"not,omitempty"`

	// It is the operation that will be executed for the given field and key. Possible operations are matches, contains, starts-with, ends-with, equals, contains-key, contains-value, greater-than, less-than, is-empty, equals-ignore-whitespace.
	// +kubebuilder:validation:Required
	Operation *string `json:"operation" tf:"operation,omitempty"`

	// Order of the condition in conditions list
	// Order of the condition in conditions list
	// +kubebuilder:validation:Optional
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`
}

type FilterObservation struct {
}

type FilterParameters struct {

	// Conditions applied to filter. This is a block, structure is documented below.
	// +kubebuilder:validation:Optional
	Conditions []ConditionsParameters `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// A filter type, supported types are: match-all, match-any-condition, match-all-conditions. Default: match-all
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RespondersObservation struct {
}

type RespondersParameters struct {

	// ID of the responder
	// +crossplane:generate:reference:type=Team
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Reference to a Team to populate id.
	// +kubebuilder:validation:Optional
	IDRef *v1.Reference `json:"idRef,omitempty" tf:"-"`

	// Selector for a Team to populate id.
	// +kubebuilder:validation:Optional
	IDSelector *v1.Selector `json:"idSelector,omitempty" tf:"-"`

	// Name of the responder
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Type of responder. Acceptable values are: user or team
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// Username of the responder
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type RestrictionObservation struct {
}

type RestrictionParameters struct {

	// Ending hour of restriction on defined end_day
	// +kubebuilder:validation:Required
	EndHour *float64 `json:"endHour" tf:"end_hour,omitempty"`

	// Ending minute of restriction on defined end_hour
	// +kubebuilder:validation:Required
	EndMin *float64 `json:"endMin" tf:"end_min,omitempty"`

	// Starting hour of restriction on defined start_day
	// +kubebuilder:validation:Required
	StartHour *float64 `json:"startHour" tf:"start_hour,omitempty"`

	// Staring minute of restriction on defined start_hour
	// +kubebuilder:validation:Required
	StartMin *float64 `json:"startMin" tf:"start_min,omitempty"`
}

type RestrictionsObservation struct {
}

type RestrictionsParameters struct {

	// Ending day of restriction (eg. wednesday)
	// +kubebuilder:validation:Required
	EndDay *string `json:"endDay" tf:"end_day,omitempty"`

	// Ending hour of restriction on defined end_day
	// +kubebuilder:validation:Required
	EndHour *float64 `json:"endHour" tf:"end_hour,omitempty"`

	// Ending minute of restriction on defined end_hour
	// +kubebuilder:validation:Required
	EndMin *float64 `json:"endMin" tf:"end_min,omitempty"`

	// Starting day of restriction (eg. monday)
	// +kubebuilder:validation:Required
	StartDay *string `json:"startDay" tf:"start_day,omitempty"`

	// Starting hour of restriction on defined start_day
	// +kubebuilder:validation:Required
	StartHour *float64 `json:"startHour" tf:"start_hour,omitempty"`

	// Staring minute of restriction on defined start_hour
	// +kubebuilder:validation:Required
	StartMin *float64 `json:"startMin" tf:"start_min,omitempty"`
}

type TimeRestrictionObservation struct {
}

type TimeRestrictionParameters struct {

	// A definition of hourly definition applied daily, this has to be used with combination: type = time-of-day. This is a block, structure is documented below.
	// +kubebuilder:validation:Optional
	Restriction []RestrictionParameters `json:"restriction,omitempty" tf:"restriction,omitempty"`

	// List of days and hours definitions for field type = weekday-and-time-of-day. This is a block, structure is documented below.
	// +kubebuilder:validation:Optional
	Restrictions []RestrictionsParameters `json:"restrictions,omitempty" tf:"restrictions,omitempty"`

	// Defines if restriction should apply daily on given hours or on certain days and hours. Possible values are: time-of-day, weekday-and-time-of-day
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// AlertPolicySpec defines the desired state of AlertPolicy
type AlertPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AlertPolicyParameters `json:"forProvider"`
}

// AlertPolicyStatus defines the observed state of AlertPolicy.
type AlertPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AlertPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AlertPolicy is the Schema for the AlertPolicys API. Manages a Alert Policy within Opsgenie.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opsgenie}
type AlertPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AlertPolicySpec   `json:"spec"`
	Status            AlertPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlertPolicyList contains a list of AlertPolicys
type AlertPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlertPolicy `json:"items"`
}

// Repository type metadata.
var (
	AlertPolicy_Kind             = "AlertPolicy"
	AlertPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AlertPolicy_Kind}.String()
	AlertPolicy_KindAPIVersion   = AlertPolicy_Kind + "." + CRDGroupVersion.String()
	AlertPolicy_GroupVersionKind = CRDGroupVersion.WithKind(AlertPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&AlertPolicy{}, &AlertPolicyList{})
}