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

type AcknowledgeFilterObservation struct {
}

type AcknowledgeFilterParameters struct {

	// +kubebuilder:validation:Optional
	Conditions []FilterConditionsParameters `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// The responder type - can be escalation, team or user.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type AcknowledgeObservation struct {
}

type AcknowledgeParameters struct {

	// An identifier that is used for alert deduplication. Default: {{alias}}.
	// +kubebuilder:validation:Optional
	Alias *string `json:"alias,omitempty" tf:"alias,omitempty"`

	// Used to specify rules for matching alerts and the filter type. Please note that depending on the integration type the field names in the filter conditions are:
	// +kubebuilder:validation:Optional
	Filter []AcknowledgeFilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// Name of the integration action.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Additional alert action note.
	// +kubebuilder:validation:Optional
	Note *string `json:"note,omitempty" tf:"note,omitempty"`

	// Integer value that defines in which order the action will be performed. Default: 1.
	// +kubebuilder:validation:Optional
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`

	// The responder type - can be escalation, team or user.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Owner of the execution for integration action.
	// +kubebuilder:validation:Optional
	User *string `json:"user,omitempty" tf:"user,omitempty"`
}

type AddNoteFilterConditionsObservation struct {
}

type AddNoteFilterConditionsParameters struct {

	// +kubebuilder:validation:Optional
	ExpectedValue *string `json:"expectedValue,omitempty" tf:"expected_value,omitempty"`

	// +kubebuilder:validation:Required
	Field *string `json:"field" tf:"field,omitempty"`

	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Not *bool `json:"not,omitempty" tf:"not,omitempty"`

	// +kubebuilder:validation:Required
	Operation *string `json:"operation" tf:"operation,omitempty"`

	// Integer value that defines in which order the action will be performed. Default: 1.
	// +kubebuilder:validation:Optional
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`
}

type AddNoteFilterObservation struct {
}

type AddNoteFilterParameters struct {

	// +kubebuilder:validation:Optional
	Conditions []AddNoteFilterConditionsParameters `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// The responder type - can be escalation, team or user.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type AddNoteObservation struct {
}

type AddNoteParameters struct {

	// An identifier that is used for alert deduplication. Default: {{alias}}.
	// +kubebuilder:validation:Optional
	Alias *string `json:"alias,omitempty" tf:"alias,omitempty"`

	// Used to specify rules for matching alerts and the filter type. Please note that depending on the integration type the field names in the filter conditions are:
	// +kubebuilder:validation:Optional
	Filter []AddNoteFilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// Name of the integration action.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Additional alert action note.
	// +kubebuilder:validation:Optional
	Note *string `json:"note,omitempty" tf:"note,omitempty"`

	// Integer value that defines in which order the action will be performed. Default: 1.
	// +kubebuilder:validation:Optional
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`

	// The responder type - can be escalation, team or user.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Owner of the execution for integration action.
	// +kubebuilder:validation:Optional
	User *string `json:"user,omitempty" tf:"user,omitempty"`
}

type CloseFilterConditionsObservation struct {
}

type CloseFilterConditionsParameters struct {

	// +kubebuilder:validation:Optional
	ExpectedValue *string `json:"expectedValue,omitempty" tf:"expected_value,omitempty"`

	// +kubebuilder:validation:Required
	Field *string `json:"field" tf:"field,omitempty"`

	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Not *bool `json:"not,omitempty" tf:"not,omitempty"`

	// +kubebuilder:validation:Required
	Operation *string `json:"operation" tf:"operation,omitempty"`

	// Integer value that defines in which order the action will be performed. Default: 1.
	// +kubebuilder:validation:Optional
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`
}

type CloseFilterObservation struct {
}

type CloseFilterParameters struct {

	// +kubebuilder:validation:Optional
	Conditions []CloseFilterConditionsParameters `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// The responder type - can be escalation, team or user.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type CloseObservation struct {
}

type CloseParameters struct {

	// An identifier that is used for alert deduplication. Default: {{alias}}.
	// +kubebuilder:validation:Optional
	Alias *string `json:"alias,omitempty" tf:"alias,omitempty"`

	// Used to specify rules for matching alerts and the filter type. Please note that depending on the integration type the field names in the filter conditions are:
	// +kubebuilder:validation:Optional
	Filter []CloseFilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// Name of the integration action.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Additional alert action note.
	// +kubebuilder:validation:Optional
	Note *string `json:"note,omitempty" tf:"note,omitempty"`

	// Integer value that defines in which order the action will be performed. Default: 1.
	// +kubebuilder:validation:Optional
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`

	// The responder type - can be escalation, team or user.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Owner of the execution for integration action.
	// +kubebuilder:validation:Optional
	User *string `json:"user,omitempty" tf:"user,omitempty"`
}

type CreateFilterConditionsObservation struct {
}

type CreateFilterConditionsParameters struct {

	// +kubebuilder:validation:Optional
	ExpectedValue *string `json:"expectedValue,omitempty" tf:"expected_value,omitempty"`

	// +kubebuilder:validation:Required
	Field *string `json:"field" tf:"field,omitempty"`

	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Not *bool `json:"not,omitempty" tf:"not,omitempty"`

	// +kubebuilder:validation:Required
	Operation *string `json:"operation" tf:"operation,omitempty"`

	// Integer value that defines in which order the action will be performed. Default: 1.
	// +kubebuilder:validation:Optional
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`
}

type CreateFilterObservation struct {
}

type CreateFilterParameters struct {

	// +kubebuilder:validation:Optional
	Conditions []CreateFilterConditionsParameters `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// The responder type - can be escalation, team or user.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type CreateObservation struct {
}

type CreateParameters struct {

	// , alias, entity, Message, recipients, responders, Subject, tags, teams, eventType, Timestamp, TopicArn.
	// +kubebuilder:validation:Optional
	AlertActions []*string `json:"alertActions,omitempty" tf:"alert_actions,omitempty"`

	// An identifier that is used for alert deduplication. Default: {{alias}}.
	// +kubebuilder:validation:Optional
	Alias *string `json:"alias,omitempty" tf:"alias,omitempty"`

	// +kubebuilder:validation:Optional
	AppendAttachments *bool `json:"appendAttachments,omitempty" tf:"append_attachments,omitempty"`

	// Custom alert priority. e.g. {{message.substring(0,2)}}
	// +kubebuilder:validation:Optional
	CustomPriority *string `json:"customPriority,omitempty" tf:"custom_priority,omitempty"`

	// Detailed description of the alert, anything that may not have fit in the message field.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The entity the alert is related to.
	// +kubebuilder:validation:Optional
	Entity *string `json:"entity,omitempty" tf:"entity,omitempty"`

	// Set of user defined properties specified as a map.
	// +kubebuilder:validation:Optional
	ExtraProperties map[string]*string `json:"extraProperties,omitempty" tf:"extra_properties,omitempty"`

	// Used to specify rules for matching alerts and the filter type. Please note that depending on the integration type the field names in the filter conditions are:
	// +kubebuilder:validation:Optional
	Filter []CreateFilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// +kubebuilder:validation:Optional
	IgnoreAlertActionsFromPayload *bool `json:"ignoreAlertActionsFromPayload,omitempty" tf:"ignore_alert_actions_from_payload,omitempty"`

	// +kubebuilder:validation:Optional
	IgnoreExtraPropertiesFromPayload *bool `json:"ignoreExtraPropertiesFromPayload,omitempty" tf:"ignore_extra_properties_from_payload,omitempty"`

	// If enabled, the integration will ignore responders sent in request payloads.
	// +kubebuilder:validation:Optional
	IgnoreRespondersFromPayload *bool `json:"ignoreRespondersFromPayload,omitempty" tf:"ignore_responders_from_payload,omitempty"`

	// +kubebuilder:validation:Optional
	IgnoreTagsFromPayload *bool `json:"ignoreTagsFromPayload,omitempty" tf:"ignore_tags_from_payload,omitempty"`

	// If enabled, the integration will ignore teams sent in request payloads.
	// +kubebuilder:validation:Optional
	IgnoreTeamsFromPayload *bool `json:"ignoreTeamsFromPayload,omitempty" tf:"ignore_teams_from_payload,omitempty"`

	// properties, recipients, teams, priority, eventType.
	// +kubebuilder:validation:Optional
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// Name of the integration action.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Additional alert action note.
	// +kubebuilder:validation:Optional
	Note *string `json:"note,omitempty" tf:"note,omitempty"`

	// Integer value that defines in which order the action will be performed. Default: 1.
	// +kubebuilder:validation:Optional
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`

	// Alert priority.
	// +kubebuilder:validation:Optional
	Priority *string `json:"priority,omitempty" tf:"priority,omitempty"`

	// User, schedule, teams or escalation names to calculate which users will receive notifications of the alert.
	// +kubebuilder:validation:Optional
	Responders []CreateRespondersParameters `json:"responders,omitempty" tf:"responders,omitempty"`

	// User defined field to specify source of action.
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// Comma separated list of labels to be attached to the alert.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The responder type - can be escalation, team or user.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Owner of the execution for integration action.
	// +kubebuilder:validation:Optional
	User *string `json:"user,omitempty" tf:"user,omitempty"`
}

type CreateRespondersObservation struct {
}

type CreateRespondersParameters struct {

	// The id of the responder.
	// +crossplane:generate:reference:type=Team
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Reference to a Team to populate id.
	// +kubebuilder:validation:Optional
	IDRef *v1.Reference `json:"idRef,omitempty" tf:"-"`

	// Selector for a Team to populate id.
	// +kubebuilder:validation:Optional
	IDSelector *v1.Selector `json:"idSelector,omitempty" tf:"-"`

	// The responder type - can be escalation, team or user.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type FilterConditionsObservation struct {
}

type FilterConditionsParameters struct {

	// +kubebuilder:validation:Optional
	ExpectedValue *string `json:"expectedValue,omitempty" tf:"expected_value,omitempty"`

	// +kubebuilder:validation:Required
	Field *string `json:"field" tf:"field,omitempty"`

	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Not *bool `json:"not,omitempty" tf:"not,omitempty"`

	// +kubebuilder:validation:Required
	Operation *string `json:"operation" tf:"operation,omitempty"`

	// Integer value that defines in which order the action will be performed. Default: 1.
	// +kubebuilder:validation:Optional
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`
}

type IgnoreFilterConditionsObservation struct {
}

type IgnoreFilterConditionsParameters struct {

	// +kubebuilder:validation:Optional
	ExpectedValue *string `json:"expectedValue,omitempty" tf:"expected_value,omitempty"`

	// +kubebuilder:validation:Required
	Field *string `json:"field" tf:"field,omitempty"`

	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Not *bool `json:"not,omitempty" tf:"not,omitempty"`

	// +kubebuilder:validation:Required
	Operation *string `json:"operation" tf:"operation,omitempty"`

	// Integer value that defines in which order the action will be performed. Default: 1.
	// +kubebuilder:validation:Optional
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`
}

type IgnoreFilterObservation struct {
}

type IgnoreFilterParameters struct {

	// +kubebuilder:validation:Optional
	Conditions []IgnoreFilterConditionsParameters `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// The responder type - can be escalation, team or user.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type IgnoreObservation struct {
}

type IgnoreParameters struct {

	// Used to specify rules for matching alerts and the filter type. Please note that depending on the integration type the field names in the filter conditions are:
	// +kubebuilder:validation:Optional
	Filter []IgnoreFilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// Name of the integration action.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Integer value that defines in which order the action will be performed. Default: 1.
	// +kubebuilder:validation:Optional
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`

	// The responder type - can be escalation, team or user.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IntegrationActionObservation struct {

	// The id of the responder.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type IntegrationActionParameters struct {

	// +kubebuilder:validation:Optional
	Acknowledge []AcknowledgeParameters `json:"acknowledge,omitempty" tf:"acknowledge,omitempty"`

	// Additional alert action note.
	// +kubebuilder:validation:Optional
	AddNote []AddNoteParameters `json:"addNote,omitempty" tf:"add_note,omitempty"`

	// +kubebuilder:validation:Optional
	Close []CloseParameters `json:"close,omitempty" tf:"close,omitempty"`

	// +kubebuilder:validation:Optional
	Create []CreateParameters `json:"create,omitempty" tf:"create,omitempty"`

	// +kubebuilder:validation:Optional
	Ignore []IgnoreParameters `json:"ignore,omitempty" tf:"ignore,omitempty"`

	// ID of the parent integration resource to bind to.
	// +crossplane:generate:reference:type=ApiIntegration
	// +kubebuilder:validation:Optional
	IntegrationID *string `json:"integrationId,omitempty" tf:"integration_id,omitempty"`

	// Reference to a ApiIntegration to populate integrationId.
	// +kubebuilder:validation:Optional
	IntegrationIDRef *v1.Reference `json:"integrationIdRef,omitempty" tf:"-"`

	// Selector for a ApiIntegration to populate integrationId.
	// +kubebuilder:validation:Optional
	IntegrationIDSelector *v1.Selector `json:"integrationIdSelector,omitempty" tf:"-"`
}

// IntegrationActionSpec defines the desired state of IntegrationAction
type IntegrationActionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IntegrationActionParameters `json:"forProvider"`
}

// IntegrationActionStatus defines the observed state of IntegrationAction.
type IntegrationActionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IntegrationActionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IntegrationAction is the Schema for the IntegrationActions API. Manages advanced actions for integrations within Opsgenie
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opsgenie}
type IntegrationAction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IntegrationActionSpec   `json:"spec"`
	Status            IntegrationActionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IntegrationActionList contains a list of IntegrationActions
type IntegrationActionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IntegrationAction `json:"items"`
}

// Repository type metadata.
var (
	IntegrationAction_Kind             = "IntegrationAction"
	IntegrationAction_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IntegrationAction_Kind}.String()
	IntegrationAction_KindAPIVersion   = IntegrationAction_Kind + "." + CRDGroupVersion.String()
	IntegrationAction_GroupVersionKind = CRDGroupVersion.WithKind(IntegrationAction_Kind)
)

func init() {
	SchemeBuilder.Register(&IntegrationAction{}, &IntegrationActionList{})
}
