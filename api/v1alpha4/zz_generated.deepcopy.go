//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha4

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	apiv1alpha4 "sigs.k8s.io/cluster-api/api/v1alpha4"
	"sigs.k8s.io/cluster-api/errors"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddressRecord) DeepCopyInto(out *AddressRecord) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddressRecord.
func (in *AddressRecord) DeepCopy() *AddressRecord {
	if in == nil {
		return nil
	}
	out := new(AddressRecord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowedNamespaces) DeepCopyInto(out *AllowedNamespaces) {
	*out = *in
	if in.NamespaceList != nil {
		in, out := &in.NamespaceList, &out.NamespaceList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowedNamespaces.
func (in *AllowedNamespaces) DeepCopy() *AllowedNamespaces {
	if in == nil {
		return nil
	}
	out := new(AllowedNamespaces)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureBastion) DeepCopyInto(out *AzureBastion) {
	*out = *in
	in.Subnet.DeepCopyInto(&out.Subnet)
	out.PublicIP = in.PublicIP
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureBastion.
func (in *AzureBastion) DeepCopy() *AzureBastion {
	if in == nil {
		return nil
	}
	out := new(AzureBastion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureCluster) DeepCopyInto(out *AzureCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureCluster.
func (in *AzureCluster) DeepCopy() *AzureCluster {
	if in == nil {
		return nil
	}
	out := new(AzureCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureClusterIdentity) DeepCopyInto(out *AzureClusterIdentity) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureClusterIdentity.
func (in *AzureClusterIdentity) DeepCopy() *AzureClusterIdentity {
	if in == nil {
		return nil
	}
	out := new(AzureClusterIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureClusterIdentity) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureClusterIdentityList) DeepCopyInto(out *AzureClusterIdentityList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AzureClusterIdentity, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureClusterIdentityList.
func (in *AzureClusterIdentityList) DeepCopy() *AzureClusterIdentityList {
	if in == nil {
		return nil
	}
	out := new(AzureClusterIdentityList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureClusterIdentityList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureClusterIdentitySpec) DeepCopyInto(out *AzureClusterIdentitySpec) {
	*out = *in
	out.ClientSecret = in.ClientSecret
	if in.AllowedNamespaces != nil {
		in, out := &in.AllowedNamespaces, &out.AllowedNamespaces
		*out = new(AllowedNamespaces)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureClusterIdentitySpec.
func (in *AzureClusterIdentitySpec) DeepCopy() *AzureClusterIdentitySpec {
	if in == nil {
		return nil
	}
	out := new(AzureClusterIdentitySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureClusterIdentityStatus) DeepCopyInto(out *AzureClusterIdentityStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(apiv1alpha4.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureClusterIdentityStatus.
func (in *AzureClusterIdentityStatus) DeepCopy() *AzureClusterIdentityStatus {
	if in == nil {
		return nil
	}
	out := new(AzureClusterIdentityStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureClusterList) DeepCopyInto(out *AzureClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AzureCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureClusterList.
func (in *AzureClusterList) DeepCopy() *AzureClusterList {
	if in == nil {
		return nil
	}
	out := new(AzureClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureClusterSpec) DeepCopyInto(out *AzureClusterSpec) {
	*out = *in
	in.NetworkSpec.DeepCopyInto(&out.NetworkSpec)
	out.ControlPlaneEndpoint = in.ControlPlaneEndpoint
	if in.AdditionalTags != nil {
		in, out := &in.AdditionalTags, &out.AdditionalTags
		*out = make(Tags, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.IdentityRef != nil {
		in, out := &in.IdentityRef, &out.IdentityRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	in.BastionSpec.DeepCopyInto(&out.BastionSpec)
	if in.CloudProviderConfigOverrides != nil {
		in, out := &in.CloudProviderConfigOverrides, &out.CloudProviderConfigOverrides
		*out = new(CloudProviderConfigOverrides)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureClusterSpec.
func (in *AzureClusterSpec) DeepCopy() *AzureClusterSpec {
	if in == nil {
		return nil
	}
	out := new(AzureClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureClusterStatus) DeepCopyInto(out *AzureClusterStatus) {
	*out = *in
	if in.FailureDomains != nil {
		in, out := &in.FailureDomains, &out.FailureDomains
		*out = make(apiv1alpha4.FailureDomains, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(apiv1alpha4.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LongRunningOperationStates != nil {
		in, out := &in.LongRunningOperationStates, &out.LongRunningOperationStates
		*out = make(Futures, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureClusterStatus.
func (in *AzureClusterStatus) DeepCopy() *AzureClusterStatus {
	if in == nil {
		return nil
	}
	out := new(AzureClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMachine) DeepCopyInto(out *AzureMachine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMachine.
func (in *AzureMachine) DeepCopy() *AzureMachine {
	if in == nil {
		return nil
	}
	out := new(AzureMachine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureMachine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMachineList) DeepCopyInto(out *AzureMachineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AzureMachine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMachineList.
func (in *AzureMachineList) DeepCopy() *AzureMachineList {
	if in == nil {
		return nil
	}
	out := new(AzureMachineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureMachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMachineSpec) DeepCopyInto(out *AzureMachineSpec) {
	*out = *in
	if in.ProviderID != nil {
		in, out := &in.ProviderID, &out.ProviderID
		*out = new(string)
		**out = **in
	}
	if in.FailureDomain != nil {
		in, out := &in.FailureDomain, &out.FailureDomain
		*out = new(string)
		**out = **in
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(Image)
		(*in).DeepCopyInto(*out)
	}
	if in.UserAssignedIdentities != nil {
		in, out := &in.UserAssignedIdentities, &out.UserAssignedIdentities
		*out = make([]UserAssignedIdentity, len(*in))
		copy(*out, *in)
	}
	in.OSDisk.DeepCopyInto(&out.OSDisk)
	if in.DataDisks != nil {
		in, out := &in.DataDisks, &out.DataDisks
		*out = make([]DataDisk, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AdditionalTags != nil {
		in, out := &in.AdditionalTags, &out.AdditionalTags
		*out = make(Tags, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AcceleratedNetworking != nil {
		in, out := &in.AcceleratedNetworking, &out.AcceleratedNetworking
		*out = new(bool)
		**out = **in
	}
	if in.SpotVMOptions != nil {
		in, out := &in.SpotVMOptions, &out.SpotVMOptions
		*out = new(SpotVMOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.SecurityProfile != nil {
		in, out := &in.SecurityProfile, &out.SecurityProfile
		*out = new(SecurityProfile)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMachineSpec.
func (in *AzureMachineSpec) DeepCopy() *AzureMachineSpec {
	if in == nil {
		return nil
	}
	out := new(AzureMachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMachineStatus) DeepCopyInto(out *AzureMachineStatus) {
	*out = *in
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]v1.NodeAddress, len(*in))
		copy(*out, *in)
	}
	if in.VMState != nil {
		in, out := &in.VMState, &out.VMState
		*out = new(ProvisioningState)
		**out = **in
	}
	if in.FailureReason != nil {
		in, out := &in.FailureReason, &out.FailureReason
		*out = new(errors.MachineStatusError)
		**out = **in
	}
	if in.FailureMessage != nil {
		in, out := &in.FailureMessage, &out.FailureMessage
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(apiv1alpha4.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LongRunningOperationStates != nil {
		in, out := &in.LongRunningOperationStates, &out.LongRunningOperationStates
		*out = make(Futures, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMachineStatus.
func (in *AzureMachineStatus) DeepCopy() *AzureMachineStatus {
	if in == nil {
		return nil
	}
	out := new(AzureMachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMachineTemplate) DeepCopyInto(out *AzureMachineTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMachineTemplate.
func (in *AzureMachineTemplate) DeepCopy() *AzureMachineTemplate {
	if in == nil {
		return nil
	}
	out := new(AzureMachineTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureMachineTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMachineTemplateList) DeepCopyInto(out *AzureMachineTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AzureMachineTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMachineTemplateList.
func (in *AzureMachineTemplateList) DeepCopy() *AzureMachineTemplateList {
	if in == nil {
		return nil
	}
	out := new(AzureMachineTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureMachineTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMachineTemplateResource) DeepCopyInto(out *AzureMachineTemplateResource) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMachineTemplateResource.
func (in *AzureMachineTemplateResource) DeepCopy() *AzureMachineTemplateResource {
	if in == nil {
		return nil
	}
	out := new(AzureMachineTemplateResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMachineTemplateSpec) DeepCopyInto(out *AzureMachineTemplateSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMachineTemplateSpec.
func (in *AzureMachineTemplateSpec) DeepCopy() *AzureMachineTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(AzureMachineTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMarketplaceImage) DeepCopyInto(out *AzureMarketplaceImage) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMarketplaceImage.
func (in *AzureMarketplaceImage) DeepCopy() *AzureMarketplaceImage {
	if in == nil {
		return nil
	}
	out := new(AzureMarketplaceImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureSharedGalleryImage) DeepCopyInto(out *AzureSharedGalleryImage) {
	*out = *in
	if in.Publisher != nil {
		in, out := &in.Publisher, &out.Publisher
		*out = new(string)
		**out = **in
	}
	if in.Offer != nil {
		in, out := &in.Offer, &out.Offer
		*out = new(string)
		**out = **in
	}
	if in.SKU != nil {
		in, out := &in.SKU, &out.SKU
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureSharedGalleryImage.
func (in *AzureSharedGalleryImage) DeepCopy() *AzureSharedGalleryImage {
	if in == nil {
		return nil
	}
	out := new(AzureSharedGalleryImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackOffConfig) DeepCopyInto(out *BackOffConfig) {
	*out = *in
	if in.CloudProviderBackoffExponent != nil {
		in, out := &in.CloudProviderBackoffExponent, &out.CloudProviderBackoffExponent
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.CloudProviderBackoffJitter != nil {
		in, out := &in.CloudProviderBackoffJitter, &out.CloudProviderBackoffJitter
		x := (*in).DeepCopy()
		*out = &x
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackOffConfig.
func (in *BackOffConfig) DeepCopy() *BackOffConfig {
	if in == nil {
		return nil
	}
	out := new(BackOffConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BastionSpec) DeepCopyInto(out *BastionSpec) {
	*out = *in
	if in.AzureBastion != nil {
		in, out := &in.AzureBastion, &out.AzureBastion
		*out = new(AzureBastion)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BastionSpec.
func (in *BastionSpec) DeepCopy() *BastionSpec {
	if in == nil {
		return nil
	}
	out := new(BastionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildParams) DeepCopyInto(out *BuildParams) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
	if in.Additional != nil {
		in, out := &in.Additional, &out.Additional
		*out = make(Tags, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildParams.
func (in *BuildParams) DeepCopy() *BuildParams {
	if in == nil {
		return nil
	}
	out := new(BuildParams)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudProviderConfigOverrides) DeepCopyInto(out *CloudProviderConfigOverrides) {
	*out = *in
	if in.RateLimits != nil {
		in, out := &in.RateLimits, &out.RateLimits
		*out = make([]RateLimitSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.BackOffs.DeepCopyInto(&out.BackOffs)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudProviderConfigOverrides.
func (in *CloudProviderConfigOverrides) DeepCopy() *CloudProviderConfigOverrides {
	if in == nil {
		return nil
	}
	out := new(CloudProviderConfigOverrides)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataDisk) DeepCopyInto(out *DataDisk) {
	*out = *in
	if in.ManagedDisk != nil {
		in, out := &in.ManagedDisk, &out.ManagedDisk
		*out = new(ManagedDiskParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Lun != nil {
		in, out := &in.Lun, &out.Lun
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataDisk.
func (in *DataDisk) DeepCopy() *DataDisk {
	if in == nil {
		return nil
	}
	out := new(DataDisk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiffDiskSettings) DeepCopyInto(out *DiffDiskSettings) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiffDiskSettings.
func (in *DiffDiskSettings) DeepCopy() *DiffDiskSettings {
	if in == nil {
		return nil
	}
	out := new(DiffDiskSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskEncryptionSetParameters) DeepCopyInto(out *DiskEncryptionSetParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskEncryptionSetParameters.
func (in *DiskEncryptionSetParameters) DeepCopy() *DiskEncryptionSetParameters {
	if in == nil {
		return nil
	}
	out := new(DiskEncryptionSetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FrontendIP) DeepCopyInto(out *FrontendIP) {
	*out = *in
	if in.PublicIP != nil {
		in, out := &in.PublicIP, &out.PublicIP
		*out = new(PublicIPSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FrontendIP.
func (in *FrontendIP) DeepCopy() *FrontendIP {
	if in == nil {
		return nil
	}
	out := new(FrontendIP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Future) DeepCopyInto(out *Future) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Future.
func (in *Future) DeepCopy() *Future {
	if in == nil {
		return nil
	}
	out := new(Future)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Futures) DeepCopyInto(out *Futures) {
	{
		in := &in
		*out = make(Futures, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Futures.
func (in Futures) DeepCopy() Futures {
	if in == nil {
		return nil
	}
	out := new(Futures)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image) DeepCopyInto(out *Image) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.SharedGallery != nil {
		in, out := &in.SharedGallery, &out.SharedGallery
		*out = new(AzureSharedGalleryImage)
		(*in).DeepCopyInto(*out)
	}
	if in.Marketplace != nil {
		in, out := &in.Marketplace, &out.Marketplace
		*out = new(AzureMarketplaceImage)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image.
func (in *Image) DeepCopy() *Image {
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerSpec) DeepCopyInto(out *LoadBalancerSpec) {
	*out = *in
	if in.FrontendIPs != nil {
		in, out := &in.FrontendIPs, &out.FrontendIPs
		*out = make([]FrontendIP, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FrontendIPsCount != nil {
		in, out := &in.FrontendIPsCount, &out.FrontendIPsCount
		*out = new(int32)
		**out = **in
	}
	if in.IdleTimeoutInMinutes != nil {
		in, out := &in.IdleTimeoutInMinutes, &out.IdleTimeoutInMinutes
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSpec.
func (in *LoadBalancerSpec) DeepCopy() *LoadBalancerSpec {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedDiskParameters) DeepCopyInto(out *ManagedDiskParameters) {
	*out = *in
	if in.DiskEncryptionSet != nil {
		in, out := &in.DiskEncryptionSet, &out.DiskEncryptionSet
		*out = new(DiskEncryptionSetParameters)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedDiskParameters.
func (in *ManagedDiskParameters) DeepCopy() *ManagedDiskParameters {
	if in == nil {
		return nil
	}
	out := new(ManagedDiskParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NatGateway) DeepCopyInto(out *NatGateway) {
	*out = *in
	out.NatGatewayIP = in.NatGatewayIP
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NatGateway.
func (in *NatGateway) DeepCopy() *NatGateway {
	if in == nil {
		return nil
	}
	out := new(NatGateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkSpec) DeepCopyInto(out *NetworkSpec) {
	*out = *in
	in.Vnet.DeepCopyInto(&out.Vnet)
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make(Subnets, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.APIServerLB.DeepCopyInto(&out.APIServerLB)
	if in.NodeOutboundLB != nil {
		in, out := &in.NodeOutboundLB, &out.NodeOutboundLB
		*out = new(LoadBalancerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ControlPlaneOutboundLB != nil {
		in, out := &in.ControlPlaneOutboundLB, &out.ControlPlaneOutboundLB
		*out = new(LoadBalancerSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkSpec.
func (in *NetworkSpec) DeepCopy() *NetworkSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OSDisk) DeepCopyInto(out *OSDisk) {
	*out = *in
	if in.DiskSizeGB != nil {
		in, out := &in.DiskSizeGB, &out.DiskSizeGB
		*out = new(int32)
		**out = **in
	}
	if in.ManagedDisk != nil {
		in, out := &in.ManagedDisk, &out.ManagedDisk
		*out = new(ManagedDiskParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.DiffDiskSettings != nil {
		in, out := &in.DiffDiskSettings, &out.DiffDiskSettings
		*out = new(DiffDiskSettings)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OSDisk.
func (in *OSDisk) DeepCopy() *OSDisk {
	if in == nil {
		return nil
	}
	out := new(OSDisk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicIPSpec) DeepCopyInto(out *PublicIPSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicIPSpec.
func (in *PublicIPSpec) DeepCopy() *PublicIPSpec {
	if in == nil {
		return nil
	}
	out := new(PublicIPSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitConfig) DeepCopyInto(out *RateLimitConfig) {
	*out = *in
	if in.CloudProviderRateLimitQPS != nil {
		in, out := &in.CloudProviderRateLimitQPS, &out.CloudProviderRateLimitQPS
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.CloudProviderRateLimitQPSWrite != nil {
		in, out := &in.CloudProviderRateLimitQPSWrite, &out.CloudProviderRateLimitQPSWrite
		x := (*in).DeepCopy()
		*out = &x
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitConfig.
func (in *RateLimitConfig) DeepCopy() *RateLimitConfig {
	if in == nil {
		return nil
	}
	out := new(RateLimitConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitSpec) DeepCopyInto(out *RateLimitSpec) {
	*out = *in
	in.Config.DeepCopyInto(&out.Config)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitSpec.
func (in *RateLimitSpec) DeepCopy() *RateLimitSpec {
	if in == nil {
		return nil
	}
	out := new(RateLimitSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteTable) DeepCopyInto(out *RouteTable) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteTable.
func (in *RouteTable) DeepCopy() *RouteTable {
	if in == nil {
		return nil
	}
	out := new(RouteTable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroup) DeepCopyInto(out *SecurityGroup) {
	*out = *in
	if in.SecurityRules != nil {
		in, out := &in.SecurityRules, &out.SecurityRules
		*out = make(SecurityRules, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(Tags, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroup.
func (in *SecurityGroup) DeepCopy() *SecurityGroup {
	if in == nil {
		return nil
	}
	out := new(SecurityGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityProfile) DeepCopyInto(out *SecurityProfile) {
	*out = *in
	if in.EncryptionAtHost != nil {
		in, out := &in.EncryptionAtHost, &out.EncryptionAtHost
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityProfile.
func (in *SecurityProfile) DeepCopy() *SecurityProfile {
	if in == nil {
		return nil
	}
	out := new(SecurityProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityRule) DeepCopyInto(out *SecurityRule) {
	*out = *in
	if in.SourcePorts != nil {
		in, out := &in.SourcePorts, &out.SourcePorts
		*out = new(string)
		**out = **in
	}
	if in.DestinationPorts != nil {
		in, out := &in.DestinationPorts, &out.DestinationPorts
		*out = new(string)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
	if in.Destination != nil {
		in, out := &in.Destination, &out.Destination
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityRule.
func (in *SecurityRule) DeepCopy() *SecurityRule {
	if in == nil {
		return nil
	}
	out := new(SecurityRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in SecurityRules) DeepCopyInto(out *SecurityRules) {
	{
		in := &in
		*out = make(SecurityRules, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityRules.
func (in SecurityRules) DeepCopy() SecurityRules {
	if in == nil {
		return nil
	}
	out := new(SecurityRules)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpotVMOptions) DeepCopyInto(out *SpotVMOptions) {
	*out = *in
	if in.MaxPrice != nil {
		in, out := &in.MaxPrice, &out.MaxPrice
		x := (*in).DeepCopy()
		*out = &x
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpotVMOptions.
func (in *SpotVMOptions) DeepCopy() *SpotVMOptions {
	if in == nil {
		return nil
	}
	out := new(SpotVMOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetSpec) DeepCopyInto(out *SubnetSpec) {
	*out = *in
	if in.CIDRBlocks != nil {
		in, out := &in.CIDRBlocks, &out.CIDRBlocks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.SecurityGroup.DeepCopyInto(&out.SecurityGroup)
	out.RouteTable = in.RouteTable
	out.NatGateway = in.NatGateway
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetSpec.
func (in *SubnetSpec) DeepCopy() *SubnetSpec {
	if in == nil {
		return nil
	}
	out := new(SubnetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Subnets) DeepCopyInto(out *Subnets) {
	{
		in := &in
		*out = make(Subnets, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subnets.
func (in Subnets) DeepCopy() Subnets {
	if in == nil {
		return nil
	}
	out := new(Subnets)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Tags) DeepCopyInto(out *Tags) {
	{
		in := &in
		*out = make(Tags, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tags.
func (in Tags) DeepCopy() Tags {
	if in == nil {
		return nil
	}
	out := new(Tags)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserAssignedIdentity) DeepCopyInto(out *UserAssignedIdentity) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserAssignedIdentity.
func (in *UserAssignedIdentity) DeepCopy() *UserAssignedIdentity {
	if in == nil {
		return nil
	}
	out := new(UserAssignedIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VM) DeepCopyInto(out *VM) {
	*out = *in
	in.Image.DeepCopyInto(&out.Image)
	in.OSDisk.DeepCopyInto(&out.OSDisk)
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(Tags, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]v1.NodeAddress, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VM.
func (in *VM) DeepCopy() *VM {
	if in == nil {
		return nil
	}
	out := new(VM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VnetSpec) DeepCopyInto(out *VnetSpec) {
	*out = *in
	if in.CIDRBlocks != nil {
		in, out := &in.CIDRBlocks, &out.CIDRBlocks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(Tags, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VnetSpec.
func (in *VnetSpec) DeepCopy() *VnetSpec {
	if in == nil {
		return nil
	}
	out := new(VnetSpec)
	in.DeepCopyInto(out)
	return out
}
