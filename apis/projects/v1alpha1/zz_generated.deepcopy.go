//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023 The Crossplane Authors.

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

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationDestination) DeepCopyInto(out *ApplicationDestination) {
	*out = *in
	if in.Server != nil {
		in, out := &in.Server, &out.Server
		*out = new(string)
		**out = **in
	}
	if in.ServerRef != nil {
		in, out := &in.ServerRef, &out.ServerRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ServerSelector != nil {
		in, out := &in.ServerSelector, &out.ServerSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationDestination.
func (in *ApplicationDestination) DeepCopy() *ApplicationDestination {
	if in == nil {
		return nil
	}
	out := new(ApplicationDestination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JWTToken) DeepCopyInto(out *JWTToken) {
	*out = *in
	if in.ExpiresAt != nil {
		in, out := &in.ExpiresAt, &out.ExpiresAt
		*out = new(int64)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JWTToken.
func (in *JWTToken) DeepCopy() *JWTToken {
	if in == nil {
		return nil
	}
	out := new(JWTToken)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JWTTokens) DeepCopyInto(out *JWTTokens) {
	*out = *in
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]JWTToken, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JWTTokens.
func (in *JWTTokens) DeepCopy() *JWTTokens {
	if in == nil {
		return nil
	}
	out := new(JWTTokens)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrphanedResourceKey) DeepCopyInto(out *OrphanedResourceKey) {
	*out = *in
	if in.Group != nil {
		in, out := &in.Group, &out.Group
		*out = new(string)
		**out = **in
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrphanedResourceKey.
func (in *OrphanedResourceKey) DeepCopy() *OrphanedResourceKey {
	if in == nil {
		return nil
	}
	out := new(OrphanedResourceKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrphanedResourcesMonitorSettings) DeepCopyInto(out *OrphanedResourcesMonitorSettings) {
	*out = *in
	if in.Warn != nil {
		in, out := &in.Warn, &out.Warn
		*out = new(bool)
		**out = **in
	}
	if in.Ignore != nil {
		in, out := &in.Ignore, &out.Ignore
		*out = make([]OrphanedResourceKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrphanedResourcesMonitorSettings.
func (in *OrphanedResourcesMonitorSettings) DeepCopy() *OrphanedResourcesMonitorSettings {
	if in == nil {
		return nil
	}
	out := new(OrphanedResourcesMonitorSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Project) DeepCopyInto(out *Project) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Project.
func (in *Project) DeepCopy() *Project {
	if in == nil {
		return nil
	}
	out := new(Project)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Project) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectList) DeepCopyInto(out *ProjectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Project, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectList.
func (in *ProjectList) DeepCopy() *ProjectList {
	if in == nil {
		return nil
	}
	out := new(ProjectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectObservation) DeepCopyInto(out *ProjectObservation) {
	*out = *in
	if in.JWTTokensByRole != nil {
		in, out := &in.JWTTokensByRole, &out.JWTTokensByRole
		*out = make(map[string]JWTTokens, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectObservation.
func (in *ProjectObservation) DeepCopy() *ProjectObservation {
	if in == nil {
		return nil
	}
	out := new(ProjectObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectParameters) DeepCopyInto(out *ProjectParameters) {
	*out = *in
	if in.SourceRepos != nil {
		in, out := &in.SourceRepos, &out.SourceRepos
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SourceReposRefs != nil {
		in, out := &in.SourceReposRefs, &out.SourceReposRefs
		*out = make([]v1.Reference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SourceReposSelector != nil {
		in, out := &in.SourceReposSelector, &out.SourceReposSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Destinations != nil {
		in, out := &in.Destinations, &out.Destinations
		*out = make([]ApplicationDestination, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]ProjectRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ClusterResourceWhitelist != nil {
		in, out := &in.ClusterResourceWhitelist, &out.ClusterResourceWhitelist
		*out = make([]metav1.GroupKind, len(*in))
		copy(*out, *in)
	}
	if in.NamespaceResourceBlacklist != nil {
		in, out := &in.NamespaceResourceBlacklist, &out.NamespaceResourceBlacklist
		*out = make([]metav1.GroupKind, len(*in))
		copy(*out, *in)
	}
	if in.OrphanedResources != nil {
		in, out := &in.OrphanedResources, &out.OrphanedResources
		*out = new(OrphanedResourcesMonitorSettings)
		(*in).DeepCopyInto(*out)
	}
	if in.SyncWindows != nil {
		in, out := &in.SyncWindows, &out.SyncWindows
		*out = make(SyncWindows, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NamespaceResourceWhitelist != nil {
		in, out := &in.NamespaceResourceWhitelist, &out.NamespaceResourceWhitelist
		*out = make([]metav1.GroupKind, len(*in))
		copy(*out, *in)
	}
	if in.SignatureKeys != nil {
		in, out := &in.SignatureKeys, &out.SignatureKeys
		*out = make([]SignatureKey, len(*in))
		copy(*out, *in)
	}
	if in.ClusterResourceBlacklist != nil {
		in, out := &in.ClusterResourceBlacklist, &out.ClusterResourceBlacklist
		*out = make([]metav1.GroupKind, len(*in))
		copy(*out, *in)
	}
	if in.ProjectLabels != nil {
		in, out := &in.ProjectLabels, &out.ProjectLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectParameters.
func (in *ProjectParameters) DeepCopy() *ProjectParameters {
	if in == nil {
		return nil
	}
	out := new(ProjectParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectRole) DeepCopyInto(out *ProjectRole) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.JWTTokens != nil {
		in, out := &in.JWTTokens, &out.JWTTokens
		*out = make([]JWTToken, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectRole.
func (in *ProjectRole) DeepCopy() *ProjectRole {
	if in == nil {
		return nil
	}
	out := new(ProjectRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectSpec) DeepCopyInto(out *ProjectSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectSpec.
func (in *ProjectSpec) DeepCopy() *ProjectSpec {
	if in == nil {
		return nil
	}
	out := new(ProjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectStatus) DeepCopyInto(out *ProjectStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectStatus.
func (in *ProjectStatus) DeepCopy() *ProjectStatus {
	if in == nil {
		return nil
	}
	out := new(ProjectStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SignatureKey) DeepCopyInto(out *SignatureKey) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SignatureKey.
func (in *SignatureKey) DeepCopy() *SignatureKey {
	if in == nil {
		return nil
	}
	out := new(SignatureKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SyncWindow) DeepCopyInto(out *SyncWindow) {
	*out = *in
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(string)
		**out = **in
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = new(string)
		**out = **in
	}
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(string)
		**out = **in
	}
	if in.Applications != nil {
		in, out := &in.Applications, &out.Applications
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ManualSync != nil {
		in, out := &in.ManualSync, &out.ManualSync
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SyncWindow.
func (in *SyncWindow) DeepCopy() *SyncWindow {
	if in == nil {
		return nil
	}
	out := new(SyncWindow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in SyncWindows) DeepCopyInto(out *SyncWindows) {
	{
		in := &in
		*out = make(SyncWindows, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SyncWindows.
func (in SyncWindows) DeepCopy() SyncWindows {
	if in == nil {
		return nil
	}
	out := new(SyncWindows)
	in.DeepCopyInto(out)
	return *out
}
