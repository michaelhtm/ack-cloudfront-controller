// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OriginAccessControlSpec defines the desired state of OriginAccessControl.
//
// A CloudFront origin access control, including its unique identifier.
type OriginAccessControlSpec struct {

// Contains the origin access control.
// +kubebuilder:validation:Required
OriginAccessControlConfig *OriginAccessControlConfig `json:"originAccessControlConfig"`
}

// OriginAccessControlStatus defines the observed state of OriginAccessControl
type OriginAccessControlStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRs managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// +kubebuilder:validation:Optional
	ETag *string `json:"eTag,omitempty"`
	// The unique identifier of the origin access control.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty"`
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty"`
}

// OriginAccessControl is the Schema for the OriginAccessControls API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type OriginAccessControl struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec   OriginAccessControlSpec   `json:"spec,omitempty"`
	Status OriginAccessControlStatus `json:"status,omitempty"`
}

// OriginAccessControlList contains a list of OriginAccessControl
// +kubebuilder:object:root=true
type OriginAccessControlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items []OriginAccessControl `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OriginAccessControl{}, &OriginAccessControlList{})
}