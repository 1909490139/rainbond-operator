// +build !ignore_autogenerated

// RAINBOND, Application Management Platform
// Copyright (C) 2014-2017 Goodrain Co., Ltd.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerStatus) DeepCopyInto(out *ControllerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerStatus.
func (in *ControllerStatus) DeepCopy() *ControllerStatus {
	if in == nil {
		return nil
	}
	out := new(ControllerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Database) DeepCopyInto(out *Database) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Database.
func (in *Database) DeepCopy() *Database {
	if in == nil {
		return nil
	}
	out := new(Database)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EtcdConfig) DeepCopyInto(out *EtcdConfig) {
	*out = *in
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EtcdConfig.
func (in *EtcdConfig) DeepCopy() *EtcdConfig {
	if in == nil {
		return nil
	}
	out := new(EtcdConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageHub) DeepCopyInto(out *ImageHub) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageHub.
func (in *ImageHub) DeepCopy() *ImageHub {
	if in == nil {
		return nil
	}
	out := new(ImageHub)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageStatus) DeepCopyInto(out *ImageStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageStatus.
func (in *ImageStatus) DeepCopy() *ImageStatus {
	if in == nil {
		return nil
	}
	out := new(ImageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletConfig) DeepCopyInto(out *KubeletConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletConfig.
func (in *KubeletConfig) DeepCopy() *KubeletConfig {
	if in == nil {
		return nil
	}
	out := new(KubeletConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeAvailPorts) DeepCopyInto(out *NodeAvailPorts) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeAvailPorts.
func (in *NodeAvailPorts) DeepCopy() *NodeAvailPorts {
	if in == nil {
		return nil
	}
	out := new(NodeAvailPorts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RainbondCluster) DeepCopyInto(out *RainbondCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(RainbondClusterStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RainbondCluster.
func (in *RainbondCluster) DeepCopy() *RainbondCluster {
	if in == nil {
		return nil
	}
	out := new(RainbondCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RainbondCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RainbondClusterCondition) DeepCopyInto(out *RainbondClusterCondition) {
	*out = *in
	if in.LastProbeTime != nil {
		in, out := &in.LastProbeTime, &out.LastProbeTime
		*out = (*in).DeepCopy()
	}
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RainbondClusterCondition.
func (in *RainbondClusterCondition) DeepCopy() *RainbondClusterCondition {
	if in == nil {
		return nil
	}
	out := new(RainbondClusterCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RainbondClusterList) DeepCopyInto(out *RainbondClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RainbondCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RainbondClusterList.
func (in *RainbondClusterList) DeepCopy() *RainbondClusterList {
	if in == nil {
		return nil
	}
	out := new(RainbondClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RainbondClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RainbondClusterSpec) DeepCopyInto(out *RainbondClusterSpec) {
	*out = *in
	if in.GatewayIngressIPs != nil {
		in, out := &in.GatewayIngressIPs, &out.GatewayIngressIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.GatewayNodes != nil {
		in, out := &in.GatewayNodes, &out.GatewayNodes
		*out = make([]NodeAvailPorts, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ImageHub != nil {
		in, out := &in.ImageHub, &out.ImageHub
		*out = new(ImageHub)
		**out = **in
	}
	if in.RegionDatabase != nil {
		in, out := &in.RegionDatabase, &out.RegionDatabase
		*out = new(Database)
		**out = **in
	}
	if in.UIDatabase != nil {
		in, out := &in.UIDatabase, &out.UIDatabase
		*out = new(Database)
		**out = **in
	}
	if in.EtcdConfig != nil {
		in, out := &in.EtcdConfig, &out.EtcdConfig
		*out = new(EtcdConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.KubeletConfig != nil {
		in, out := &in.KubeletConfig, &out.KubeletConfig
		*out = new(KubeletConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RainbondClusterSpec.
func (in *RainbondClusterSpec) DeepCopy() *RainbondClusterSpec {
	if in == nil {
		return nil
	}
	out := new(RainbondClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RainbondClusterStatus) DeepCopyInto(out *RainbondClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]RainbondClusterCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeAvailPorts != nil {
		in, out := &in.NodeAvailPorts, &out.NodeAvailPorts
		*out = make([]*NodeAvailPorts, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(NodeAvailPorts)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.StorageClasses != nil {
		in, out := &in.StorageClasses, &out.StorageClasses
		*out = make([]*StorageClass, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(StorageClass)
				**out = **in
			}
		}
	}
	if in.ControllerStatues != nil {
		in, out := &in.ControllerStatues, &out.ControllerStatues
		*out = make([]*ControllerStatus, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ControllerStatus)
				**out = **in
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RainbondClusterStatus.
func (in *RainbondClusterStatus) DeepCopy() *RainbondClusterStatus {
	if in == nil {
		return nil
	}
	out := new(RainbondClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RainbondPackage) DeepCopyInto(out *RainbondPackage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(RainbondPackageStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RainbondPackage.
func (in *RainbondPackage) DeepCopy() *RainbondPackage {
	if in == nil {
		return nil
	}
	out := new(RainbondPackage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RainbondPackage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RainbondPackageList) DeepCopyInto(out *RainbondPackageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RainbondPackage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RainbondPackageList.
func (in *RainbondPackageList) DeepCopy() *RainbondPackageList {
	if in == nil {
		return nil
	}
	out := new(RainbondPackageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RainbondPackageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RainbondPackageSpec) DeepCopyInto(out *RainbondPackageSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RainbondPackageSpec.
func (in *RainbondPackageSpec) DeepCopy() *RainbondPackageSpec {
	if in == nil {
		return nil
	}
	out := new(RainbondPackageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RainbondPackageStatus) DeepCopyInto(out *RainbondPackageStatus) {
	*out = *in
	if in.ImageStatus != nil {
		in, out := &in.ImageStatus, &out.ImageStatus
		*out = make(map[string]ImageStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RainbondPackageStatus.
func (in *RainbondPackageStatus) DeepCopy() *RainbondPackageStatus {
	if in == nil {
		return nil
	}
	out := new(RainbondPackageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RbdComponent) DeepCopyInto(out *RbdComponent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(RbdComponentStatus)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RbdComponent.
func (in *RbdComponent) DeepCopy() *RbdComponent {
	if in == nil {
		return nil
	}
	out := new(RbdComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RbdComponent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RbdComponentList) DeepCopyInto(out *RbdComponentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RbdComponent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RbdComponentList.
func (in *RbdComponentList) DeepCopy() *RbdComponentList {
	if in == nil {
		return nil
	}
	out := new(RbdComponentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RbdComponentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RbdComponentSpec) DeepCopyInto(out *RbdComponentSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RbdComponentSpec.
func (in *RbdComponentSpec) DeepCopy() *RbdComponentSpec {
	if in == nil {
		return nil
	}
	out := new(RbdComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RbdComponentStatus) DeepCopyInto(out *RbdComponentStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RbdComponentStatus.
func (in *RbdComponentStatus) DeepCopy() *RbdComponentStatus {
	if in == nil {
		return nil
	}
	out := new(RbdComponentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageClass) DeepCopyInto(out *StorageClass) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageClass.
func (in *StorageClass) DeepCopy() *StorageClass {
	if in == nil {
		return nil
	}
	out := new(StorageClass)
	in.DeepCopyInto(out)
	return out
}
