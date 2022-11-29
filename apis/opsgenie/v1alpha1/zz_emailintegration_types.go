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

type EmailIntegrationObservation struct {

	// The id of the responder.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type EmailIntegrationParameters struct {

	// The username part of the email address. It must be unique for each integration.
	// +kubebuilder:validation:Required
	EmailUsername *string `json:"emailUsername" tf:"email_username,omitempty"`

	// A Member block as documented below.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// If enabled, the integration will ignore recipients sent in request payloads. Default: false.
	// +kubebuilder:validation:Optional
	IgnoreRespondersFromPayload *bool `json:"ignoreRespondersFromPayload,omitempty" tf:"ignore_responders_from_payload,omitempty"`

	// Name of the integration. Name must be unique for each integration.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Owner team id of the integration.
	// +crossplane:generate:reference:type=Team
	// +kubebuilder:validation:Optional
	OwnerTeamID *string `json:"ownerTeamId,omitempty" tf:"owner_team_id,omitempty"`

	// Reference to a Team to populate ownerTeamId.
	// +kubebuilder:validation:Optional
	OwnerTeamIDRef *v1.Reference `json:"ownerTeamIdRef,omitempty" tf:"-"`

	// Selector for a Team to populate ownerTeamId.
	// +kubebuilder:validation:Optional
	OwnerTeamIDSelector *v1.Selector `json:"ownerTeamIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Responders []EmailIntegrationRespondersParameters `json:"responders,omitempty" tf:"responders,omitempty"`

	// If enabled, notifications that come from alerts will be suppressed. Default: false.
	// +kubebuilder:validation:Optional
	SuppressNotifications *bool `json:"suppressNotifications,omitempty" tf:"suppress_notifications,omitempty"`
}

type EmailIntegrationRespondersObservation struct {
}

type EmailIntegrationRespondersParameters struct {

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

	// The responder type.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// EmailIntegrationSpec defines the desired state of EmailIntegration
type EmailIntegrationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EmailIntegrationParameters `json:"forProvider"`
}

// EmailIntegrationStatus defines the observed state of EmailIntegration.
type EmailIntegrationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EmailIntegrationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EmailIntegration is the Schema for the EmailIntegrations API. Manages an Email Integration within Opsgenie.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opsgenie}
type EmailIntegration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EmailIntegrationSpec   `json:"spec"`
	Status            EmailIntegrationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EmailIntegrationList contains a list of EmailIntegrations
type EmailIntegrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EmailIntegration `json:"items"`
}

// Repository type metadata.
var (
	EmailIntegration_Kind             = "EmailIntegration"
	EmailIntegration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EmailIntegration_Kind}.String()
	EmailIntegration_KindAPIVersion   = EmailIntegration_Kind + "." + CRDGroupVersion.String()
	EmailIntegration_GroupVersionKind = CRDGroupVersion.WithKind(EmailIntegration_Kind)
)

func init() {
	SchemeBuilder.Register(&EmailIntegration{}, &EmailIntegrationList{})
}