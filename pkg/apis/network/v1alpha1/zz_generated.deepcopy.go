// +build !ignore_autogenerated

/*
Copyright 2019 The KubeSphere authors.

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

// autogenerated by controller-gen object, do not modify manually

package v1alpha1

import (
	"k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"kubesphere.io/kubesphere/pkg/apis/network/v1alpha1/numorstring"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EntityRule) DeepCopyInto(out *EntityRule) {
	*out = *in
	if in.Nets != nil {
		in, out := &in.Nets, &out.Nets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]numorstring.Port, len(*in))
		copy(*out, *in)
	}
	if in.NotNets != nil {
		in, out := &in.NotNets, &out.NotNets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NotPorts != nil {
		in, out := &in.NotPorts, &out.NotPorts
		*out = make([]numorstring.Port, len(*in))
		copy(*out, *in)
	}
	if in.ServiceAccounts != nil {
		in, out := &in.ServiceAccounts, &out.ServiceAccounts
		*out = new(ServiceAccountMatch)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EntityRule.
func (in *EntityRule) DeepCopy() *EntityRule {
	if in == nil {
		return nil
	}
	out := new(EntityRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPMatch) DeepCopyInto(out *HTTPMatch) {
	*out = *in
	if in.Methods != nil {
		in, out := &in.Methods, &out.Methods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = make([]HTTPPath, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPMatch.
func (in *HTTPMatch) DeepCopy() *HTTPMatch {
	if in == nil {
		return nil
	}
	out := new(HTTPMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPPath) DeepCopyInto(out *HTTPPath) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPPath.
func (in *HTTPPath) DeepCopy() *HTTPPath {
	if in == nil {
		return nil
	}
	out := new(HTTPPath)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ICMPFields) DeepCopyInto(out *ICMPFields) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(int)
		**out = **in
	}
	if in.Code != nil {
		in, out := &in.Code, &out.Code
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ICMPFields.
func (in *ICMPFields) DeepCopy() *ICMPFields {
	if in == nil {
		return nil
	}
	out := new(ICMPFields)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceNetworkPolicy) DeepCopyInto(out *NamespaceNetworkPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceNetworkPolicy.
func (in *NamespaceNetworkPolicy) DeepCopy() *NamespaceNetworkPolicy {
	if in == nil {
		return nil
	}
	out := new(NamespaceNetworkPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespaceNetworkPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceNetworkPolicyList) DeepCopyInto(out *NamespaceNetworkPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NamespaceNetworkPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceNetworkPolicyList.
func (in *NamespaceNetworkPolicyList) DeepCopy() *NamespaceNetworkPolicyList {
	if in == nil {
		return nil
	}
	out := new(NamespaceNetworkPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespaceNetworkPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceNetworkPolicySpec) DeepCopyInto(out *NamespaceNetworkPolicySpec) {
	*out = *in
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = new(int)
		**out = **in
	}
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = make([]Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Egress != nil {
		in, out := &in.Egress, &out.Egress
		*out = make([]Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Types != nil {
		in, out := &in.Types, &out.Types
		*out = make([]PolicyType, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceNetworkPolicySpec.
func (in *NamespaceNetworkPolicySpec) DeepCopy() *NamespaceNetworkPolicySpec {
	if in == nil {
		return nil
	}
	out := new(NamespaceNetworkPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rule) DeepCopyInto(out *Rule) {
	*out = *in
	if in.IPVersion != nil {
		in, out := &in.IPVersion, &out.IPVersion
		*out = new(int)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(v1.Protocol)
		**out = **in
	}
	if in.ICMP != nil {
		in, out := &in.ICMP, &out.ICMP
		*out = new(ICMPFields)
		(*in).DeepCopyInto(*out)
	}
	if in.NotProtocol != nil {
		in, out := &in.NotProtocol, &out.NotProtocol
		*out = new(v1.Protocol)
		**out = **in
	}
	if in.NotICMP != nil {
		in, out := &in.NotICMP, &out.NotICMP
		*out = new(ICMPFields)
		(*in).DeepCopyInto(*out)
	}
	in.Source.DeepCopyInto(&out.Source)
	in.Destination.DeepCopyInto(&out.Destination)
	if in.HTTP != nil {
		in, out := &in.HTTP, &out.HTTP
		*out = new(HTTPMatch)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule.
func (in *Rule) DeepCopy() *Rule {
	if in == nil {
		return nil
	}
	out := new(Rule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountMatch) DeepCopyInto(out *ServiceAccountMatch) {
	*out = *in
	if in.Names != nil {
		in, out := &in.Names, &out.Names
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountMatch.
func (in *ServiceAccountMatch) DeepCopy() *ServiceAccountMatch {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceNetworkPolicy) DeepCopyInto(out *WorkspaceNetworkPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceNetworkPolicy.
func (in *WorkspaceNetworkPolicy) DeepCopy() *WorkspaceNetworkPolicy {
	if in == nil {
		return nil
	}
	out := new(WorkspaceNetworkPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkspaceNetworkPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceNetworkPolicyEgressRule) DeepCopyInto(out *WorkspaceNetworkPolicyEgressRule) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]networkingv1.NetworkPolicyPort, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.To != nil {
		in, out := &in.To, &out.To
		*out = make([]WorkspaceNetworkPolicyPeer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceNetworkPolicyEgressRule.
func (in *WorkspaceNetworkPolicyEgressRule) DeepCopy() *WorkspaceNetworkPolicyEgressRule {
	if in == nil {
		return nil
	}
	out := new(WorkspaceNetworkPolicyEgressRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceNetworkPolicyIngressRule) DeepCopyInto(out *WorkspaceNetworkPolicyIngressRule) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]networkingv1.NetworkPolicyPort, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = make([]WorkspaceNetworkPolicyPeer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceNetworkPolicyIngressRule.
func (in *WorkspaceNetworkPolicyIngressRule) DeepCopy() *WorkspaceNetworkPolicyIngressRule {
	if in == nil {
		return nil
	}
	out := new(WorkspaceNetworkPolicyIngressRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceNetworkPolicyList) DeepCopyInto(out *WorkspaceNetworkPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkspaceNetworkPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceNetworkPolicyList.
func (in *WorkspaceNetworkPolicyList) DeepCopy() *WorkspaceNetworkPolicyList {
	if in == nil {
		return nil
	}
	out := new(WorkspaceNetworkPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkspaceNetworkPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceNetworkPolicyPeer) DeepCopyInto(out *WorkspaceNetworkPolicyPeer) {
	*out = *in
	in.NetworkPolicyPeer.DeepCopyInto(&out.NetworkPolicyPeer)
	if in.WorkspaceSelector != nil {
		in, out := &in.WorkspaceSelector, &out.WorkspaceSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceNetworkPolicyPeer.
func (in *WorkspaceNetworkPolicyPeer) DeepCopy() *WorkspaceNetworkPolicyPeer {
	if in == nil {
		return nil
	}
	out := new(WorkspaceNetworkPolicyPeer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceNetworkPolicySpec) DeepCopyInto(out *WorkspaceNetworkPolicySpec) {
	*out = *in
	if in.PolicyTypes != nil {
		in, out := &in.PolicyTypes, &out.PolicyTypes
		*out = make([]networkingv1.PolicyType, len(*in))
		copy(*out, *in)
	}
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = make([]WorkspaceNetworkPolicyIngressRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Egress != nil {
		in, out := &in.Egress, &out.Egress
		*out = make([]WorkspaceNetworkPolicyEgressRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceNetworkPolicySpec.
func (in *WorkspaceNetworkPolicySpec) DeepCopy() *WorkspaceNetworkPolicySpec {
	if in == nil {
		return nil
	}
	out := new(WorkspaceNetworkPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceNetworkPolicyStatus) DeepCopyInto(out *WorkspaceNetworkPolicyStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceNetworkPolicyStatus.
func (in *WorkspaceNetworkPolicyStatus) DeepCopy() *WorkspaceNetworkPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(WorkspaceNetworkPolicyStatus)
	in.DeepCopyInto(out)
	return out
}
