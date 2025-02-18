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

// Code generated by deepcopy-gen. DO NOT EDIT.

package resource

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/kubernetes/pkg/apis/core"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocationResult) DeepCopyInto(out *AllocationResult) {
	*out = *in
	if in.ResourceHandles != nil {
		in, out := &in.ResourceHandles, &out.ResourceHandles
		*out = make([]ResourceHandle, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AvailableOnNodes != nil {
		in, out := &in.AvailableOnNodes, &out.AvailableOnNodes
		*out = new(core.NodeSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocationResult.
func (in *AllocationResult) DeepCopy() *AllocationResult {
	if in == nil {
		return nil
	}
	out := new(AllocationResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocationResultModel) DeepCopyInto(out *AllocationResultModel) {
	*out = *in
	if in.NamedResources != nil {
		in, out := &in.NamedResources, &out.NamedResources
		*out = new(NamedResourcesAllocationResult)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocationResultModel.
func (in *AllocationResultModel) DeepCopy() *AllocationResultModel {
	if in == nil {
		return nil
	}
	out := new(AllocationResultModel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DriverAllocationResult) DeepCopyInto(out *DriverAllocationResult) {
	*out = *in
	if in.VendorRequestParameters != nil {
		out.VendorRequestParameters = in.VendorRequestParameters.DeepCopyObject()
	}
	in.AllocationResultModel.DeepCopyInto(&out.AllocationResultModel)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DriverAllocationResult.
func (in *DriverAllocationResult) DeepCopy() *DriverAllocationResult {
	if in == nil {
		return nil
	}
	out := new(DriverAllocationResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DriverRequests) DeepCopyInto(out *DriverRequests) {
	*out = *in
	if in.VendorParameters != nil {
		out.VendorParameters = in.VendorParameters.DeepCopyObject()
	}
	if in.Requests != nil {
		in, out := &in.Requests, &out.Requests
		*out = make([]ResourceRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DriverRequests.
func (in *DriverRequests) DeepCopy() *DriverRequests {
	if in == nil {
		return nil
	}
	out := new(DriverRequests)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedResourcesAllocationResult) DeepCopyInto(out *NamedResourcesAllocationResult) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedResourcesAllocationResult.
func (in *NamedResourcesAllocationResult) DeepCopy() *NamedResourcesAllocationResult {
	if in == nil {
		return nil
	}
	out := new(NamedResourcesAllocationResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedResourcesAttribute) DeepCopyInto(out *NamedResourcesAttribute) {
	*out = *in
	in.NamedResourcesAttributeValue.DeepCopyInto(&out.NamedResourcesAttributeValue)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedResourcesAttribute.
func (in *NamedResourcesAttribute) DeepCopy() *NamedResourcesAttribute {
	if in == nil {
		return nil
	}
	out := new(NamedResourcesAttribute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedResourcesAttributeValue) DeepCopyInto(out *NamedResourcesAttributeValue) {
	*out = *in
	if in.QuantityValue != nil {
		in, out := &in.QuantityValue, &out.QuantityValue
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.BoolValue != nil {
		in, out := &in.BoolValue, &out.BoolValue
		*out = new(bool)
		**out = **in
	}
	if in.IntValue != nil {
		in, out := &in.IntValue, &out.IntValue
		*out = new(int64)
		**out = **in
	}
	if in.IntSliceValue != nil {
		in, out := &in.IntSliceValue, &out.IntSliceValue
		*out = new(NamedResourcesIntSlice)
		(*in).DeepCopyInto(*out)
	}
	if in.StringValue != nil {
		in, out := &in.StringValue, &out.StringValue
		*out = new(string)
		**out = **in
	}
	if in.StringSliceValue != nil {
		in, out := &in.StringSliceValue, &out.StringSliceValue
		*out = new(NamedResourcesStringSlice)
		(*in).DeepCopyInto(*out)
	}
	if in.VersionValue != nil {
		in, out := &in.VersionValue, &out.VersionValue
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedResourcesAttributeValue.
func (in *NamedResourcesAttributeValue) DeepCopy() *NamedResourcesAttributeValue {
	if in == nil {
		return nil
	}
	out := new(NamedResourcesAttributeValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedResourcesFilter) DeepCopyInto(out *NamedResourcesFilter) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedResourcesFilter.
func (in *NamedResourcesFilter) DeepCopy() *NamedResourcesFilter {
	if in == nil {
		return nil
	}
	out := new(NamedResourcesFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedResourcesInstance) DeepCopyInto(out *NamedResourcesInstance) {
	*out = *in
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = make([]NamedResourcesAttribute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedResourcesInstance.
func (in *NamedResourcesInstance) DeepCopy() *NamedResourcesInstance {
	if in == nil {
		return nil
	}
	out := new(NamedResourcesInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedResourcesIntSlice) DeepCopyInto(out *NamedResourcesIntSlice) {
	*out = *in
	if in.Ints != nil {
		in, out := &in.Ints, &out.Ints
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedResourcesIntSlice.
func (in *NamedResourcesIntSlice) DeepCopy() *NamedResourcesIntSlice {
	if in == nil {
		return nil
	}
	out := new(NamedResourcesIntSlice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedResourcesRequest) DeepCopyInto(out *NamedResourcesRequest) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedResourcesRequest.
func (in *NamedResourcesRequest) DeepCopy() *NamedResourcesRequest {
	if in == nil {
		return nil
	}
	out := new(NamedResourcesRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedResourcesResources) DeepCopyInto(out *NamedResourcesResources) {
	*out = *in
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = make([]NamedResourcesInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedResourcesResources.
func (in *NamedResourcesResources) DeepCopy() *NamedResourcesResources {
	if in == nil {
		return nil
	}
	out := new(NamedResourcesResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedResourcesStringSlice) DeepCopyInto(out *NamedResourcesStringSlice) {
	*out = *in
	if in.Strings != nil {
		in, out := &in.Strings, &out.Strings
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedResourcesStringSlice.
func (in *NamedResourcesStringSlice) DeepCopy() *NamedResourcesStringSlice {
	if in == nil {
		return nil
	}
	out := new(NamedResourcesStringSlice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeResourceModel) DeepCopyInto(out *NodeResourceModel) {
	*out = *in
	if in.NamedResources != nil {
		in, out := &in.NamedResources, &out.NamedResources
		*out = new(NamedResourcesResources)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeResourceModel.
func (in *NodeResourceModel) DeepCopy() *NodeResourceModel {
	if in == nil {
		return nil
	}
	out := new(NodeResourceModel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSchedulingContext) DeepCopyInto(out *PodSchedulingContext) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSchedulingContext.
func (in *PodSchedulingContext) DeepCopy() *PodSchedulingContext {
	if in == nil {
		return nil
	}
	out := new(PodSchedulingContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodSchedulingContext) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSchedulingContextList) DeepCopyInto(out *PodSchedulingContextList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PodSchedulingContext, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSchedulingContextList.
func (in *PodSchedulingContextList) DeepCopy() *PodSchedulingContextList {
	if in == nil {
		return nil
	}
	out := new(PodSchedulingContextList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodSchedulingContextList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSchedulingContextSpec) DeepCopyInto(out *PodSchedulingContextSpec) {
	*out = *in
	if in.PotentialNodes != nil {
		in, out := &in.PotentialNodes, &out.PotentialNodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSchedulingContextSpec.
func (in *PodSchedulingContextSpec) DeepCopy() *PodSchedulingContextSpec {
	if in == nil {
		return nil
	}
	out := new(PodSchedulingContextSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSchedulingContextStatus) DeepCopyInto(out *PodSchedulingContextStatus) {
	*out = *in
	if in.ResourceClaims != nil {
		in, out := &in.ResourceClaims, &out.ResourceClaims
		*out = make([]ResourceClaimSchedulingStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSchedulingContextStatus.
func (in *PodSchedulingContextStatus) DeepCopy() *PodSchedulingContextStatus {
	if in == nil {
		return nil
	}
	out := new(PodSchedulingContextStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceClaim) DeepCopyInto(out *ResourceClaim) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceClaim.
func (in *ResourceClaim) DeepCopy() *ResourceClaim {
	if in == nil {
		return nil
	}
	out := new(ResourceClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceClaim) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceClaimConsumerReference) DeepCopyInto(out *ResourceClaimConsumerReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceClaimConsumerReference.
func (in *ResourceClaimConsumerReference) DeepCopy() *ResourceClaimConsumerReference {
	if in == nil {
		return nil
	}
	out := new(ResourceClaimConsumerReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceClaimList) DeepCopyInto(out *ResourceClaimList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceClaimList.
func (in *ResourceClaimList) DeepCopy() *ResourceClaimList {
	if in == nil {
		return nil
	}
	out := new(ResourceClaimList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceClaimList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceClaimParameters) DeepCopyInto(out *ResourceClaimParameters) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.GeneratedFrom != nil {
		in, out := &in.GeneratedFrom, &out.GeneratedFrom
		*out = new(ResourceClaimParametersReference)
		**out = **in
	}
	if in.DriverRequests != nil {
		in, out := &in.DriverRequests, &out.DriverRequests
		*out = make([]DriverRequests, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceClaimParameters.
func (in *ResourceClaimParameters) DeepCopy() *ResourceClaimParameters {
	if in == nil {
		return nil
	}
	out := new(ResourceClaimParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceClaimParameters) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceClaimParametersList) DeepCopyInto(out *ResourceClaimParametersList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceClaimParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceClaimParametersList.
func (in *ResourceClaimParametersList) DeepCopy() *ResourceClaimParametersList {
	if in == nil {
		return nil
	}
	out := new(ResourceClaimParametersList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceClaimParametersList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceClaimParametersReference) DeepCopyInto(out *ResourceClaimParametersReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceClaimParametersReference.
func (in *ResourceClaimParametersReference) DeepCopy() *ResourceClaimParametersReference {
	if in == nil {
		return nil
	}
	out := new(ResourceClaimParametersReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceClaimSchedulingStatus) DeepCopyInto(out *ResourceClaimSchedulingStatus) {
	*out = *in
	if in.UnsuitableNodes != nil {
		in, out := &in.UnsuitableNodes, &out.UnsuitableNodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceClaimSchedulingStatus.
func (in *ResourceClaimSchedulingStatus) DeepCopy() *ResourceClaimSchedulingStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceClaimSchedulingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceClaimSpec) DeepCopyInto(out *ResourceClaimSpec) {
	*out = *in
	if in.ParametersRef != nil {
		in, out := &in.ParametersRef, &out.ParametersRef
		*out = new(ResourceClaimParametersReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceClaimSpec.
func (in *ResourceClaimSpec) DeepCopy() *ResourceClaimSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceClaimSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceClaimStatus) DeepCopyInto(out *ResourceClaimStatus) {
	*out = *in
	if in.Allocation != nil {
		in, out := &in.Allocation, &out.Allocation
		*out = new(AllocationResult)
		(*in).DeepCopyInto(*out)
	}
	if in.ReservedFor != nil {
		in, out := &in.ReservedFor, &out.ReservedFor
		*out = make([]ResourceClaimConsumerReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceClaimStatus.
func (in *ResourceClaimStatus) DeepCopy() *ResourceClaimStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceClaimStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceClaimTemplate) DeepCopyInto(out *ResourceClaimTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceClaimTemplate.
func (in *ResourceClaimTemplate) DeepCopy() *ResourceClaimTemplate {
	if in == nil {
		return nil
	}
	out := new(ResourceClaimTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceClaimTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceClaimTemplateList) DeepCopyInto(out *ResourceClaimTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceClaimTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceClaimTemplateList.
func (in *ResourceClaimTemplateList) DeepCopy() *ResourceClaimTemplateList {
	if in == nil {
		return nil
	}
	out := new(ResourceClaimTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceClaimTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceClaimTemplateSpec) DeepCopyInto(out *ResourceClaimTemplateSpec) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceClaimTemplateSpec.
func (in *ResourceClaimTemplateSpec) DeepCopy() *ResourceClaimTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceClaimTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceClass) DeepCopyInto(out *ResourceClass) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.ParametersRef != nil {
		in, out := &in.ParametersRef, &out.ParametersRef
		*out = new(ResourceClassParametersReference)
		**out = **in
	}
	if in.SuitableNodes != nil {
		in, out := &in.SuitableNodes, &out.SuitableNodes
		*out = new(core.NodeSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.StructuredParameters != nil {
		in, out := &in.StructuredParameters, &out.StructuredParameters
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceClass.
func (in *ResourceClass) DeepCopy() *ResourceClass {
	if in == nil {
		return nil
	}
	out := new(ResourceClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceClass) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceClassList) DeepCopyInto(out *ResourceClassList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceClassList.
func (in *ResourceClassList) DeepCopy() *ResourceClassList {
	if in == nil {
		return nil
	}
	out := new(ResourceClassList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceClassList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceClassParameters) DeepCopyInto(out *ResourceClassParameters) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.GeneratedFrom != nil {
		in, out := &in.GeneratedFrom, &out.GeneratedFrom
		*out = new(ResourceClassParametersReference)
		**out = **in
	}
	if in.VendorParameters != nil {
		in, out := &in.VendorParameters, &out.VendorParameters
		*out = make([]VendorParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Filters != nil {
		in, out := &in.Filters, &out.Filters
		*out = make([]ResourceFilter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceClassParameters.
func (in *ResourceClassParameters) DeepCopy() *ResourceClassParameters {
	if in == nil {
		return nil
	}
	out := new(ResourceClassParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceClassParameters) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceClassParametersList) DeepCopyInto(out *ResourceClassParametersList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceClassParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceClassParametersList.
func (in *ResourceClassParametersList) DeepCopy() *ResourceClassParametersList {
	if in == nil {
		return nil
	}
	out := new(ResourceClassParametersList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceClassParametersList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceClassParametersReference) DeepCopyInto(out *ResourceClassParametersReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceClassParametersReference.
func (in *ResourceClassParametersReference) DeepCopy() *ResourceClassParametersReference {
	if in == nil {
		return nil
	}
	out := new(ResourceClassParametersReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceFilter) DeepCopyInto(out *ResourceFilter) {
	*out = *in
	in.ResourceFilterModel.DeepCopyInto(&out.ResourceFilterModel)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceFilter.
func (in *ResourceFilter) DeepCopy() *ResourceFilter {
	if in == nil {
		return nil
	}
	out := new(ResourceFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceFilterModel) DeepCopyInto(out *ResourceFilterModel) {
	*out = *in
	if in.NamedResources != nil {
		in, out := &in.NamedResources, &out.NamedResources
		*out = new(NamedResourcesFilter)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceFilterModel.
func (in *ResourceFilterModel) DeepCopy() *ResourceFilterModel {
	if in == nil {
		return nil
	}
	out := new(ResourceFilterModel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceHandle) DeepCopyInto(out *ResourceHandle) {
	*out = *in
	if in.StructuredData != nil {
		in, out := &in.StructuredData, &out.StructuredData
		*out = new(StructuredResourceHandle)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceHandle.
func (in *ResourceHandle) DeepCopy() *ResourceHandle {
	if in == nil {
		return nil
	}
	out := new(ResourceHandle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRequest) DeepCopyInto(out *ResourceRequest) {
	*out = *in
	if in.VendorParameters != nil {
		out.VendorParameters = in.VendorParameters.DeepCopyObject()
	}
	in.ResourceRequestModel.DeepCopyInto(&out.ResourceRequestModel)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRequest.
func (in *ResourceRequest) DeepCopy() *ResourceRequest {
	if in == nil {
		return nil
	}
	out := new(ResourceRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRequestModel) DeepCopyInto(out *ResourceRequestModel) {
	*out = *in
	if in.NamedResources != nil {
		in, out := &in.NamedResources, &out.NamedResources
		*out = new(NamedResourcesRequest)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRequestModel.
func (in *ResourceRequestModel) DeepCopy() *ResourceRequestModel {
	if in == nil {
		return nil
	}
	out := new(ResourceRequestModel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSlice) DeepCopyInto(out *ResourceSlice) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.NodeResourceModel.DeepCopyInto(&out.NodeResourceModel)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSlice.
func (in *ResourceSlice) DeepCopy() *ResourceSlice {
	if in == nil {
		return nil
	}
	out := new(ResourceSlice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceSlice) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSliceList) DeepCopyInto(out *ResourceSliceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceSlice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSliceList.
func (in *ResourceSliceList) DeepCopy() *ResourceSliceList {
	if in == nil {
		return nil
	}
	out := new(ResourceSliceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceSliceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StructuredResourceHandle) DeepCopyInto(out *StructuredResourceHandle) {
	*out = *in
	if in.VendorClassParameters != nil {
		out.VendorClassParameters = in.VendorClassParameters.DeepCopyObject()
	}
	if in.VendorClaimParameters != nil {
		out.VendorClaimParameters = in.VendorClaimParameters.DeepCopyObject()
	}
	if in.Results != nil {
		in, out := &in.Results, &out.Results
		*out = make([]DriverAllocationResult, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StructuredResourceHandle.
func (in *StructuredResourceHandle) DeepCopy() *StructuredResourceHandle {
	if in == nil {
		return nil
	}
	out := new(StructuredResourceHandle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VendorParameters) DeepCopyInto(out *VendorParameters) {
	*out = *in
	if in.Parameters != nil {
		out.Parameters = in.Parameters.DeepCopyObject()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VendorParameters.
func (in *VendorParameters) DeepCopy() *VendorParameters {
	if in == nil {
		return nil
	}
	out := new(VendorParameters)
	in.DeepCopyInto(out)
	return out
}
