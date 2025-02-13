// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokeredInfraConfig) DeepCopyInto(out *BrokeredInfraConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokeredInfraConfig.
func (in *BrokeredInfraConfig) DeepCopy() *BrokeredInfraConfig {
	if in == nil {
		return nil
	}
	out := new(BrokeredInfraConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BrokeredInfraConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokeredInfraConfigList) DeepCopyInto(out *BrokeredInfraConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BrokeredInfraConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokeredInfraConfigList.
func (in *BrokeredInfraConfigList) DeepCopy() *BrokeredInfraConfigList {
	if in == nil {
		return nil
	}
	out := new(BrokeredInfraConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BrokeredInfraConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokeredInfraConfigSpec) DeepCopyInto(out *BrokeredInfraConfigSpec) {
	*out = *in
	out.Admin = in.Admin
	out.Broker = in.Broker
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokeredInfraConfigSpec.
func (in *BrokeredInfraConfigSpec) DeepCopy() *BrokeredInfraConfigSpec {
	if in == nil {
		return nil
	}
	out := new(BrokeredInfraConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokeredInfraConfigStatus) DeepCopyInto(out *BrokeredInfraConfigStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokeredInfraConfigStatus.
func (in *BrokeredInfraConfigStatus) DeepCopy() *BrokeredInfraConfigStatus {
	if in == nil {
		return nil
	}
	out := new(BrokeredInfraConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfraConfigAdmin) DeepCopyInto(out *InfraConfigAdmin) {
	*out = *in
	out.Resources = in.Resources
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfraConfigAdmin.
func (in *InfraConfigAdmin) DeepCopy() *InfraConfigAdmin {
	if in == nil {
		return nil
	}
	out := new(InfraConfigAdmin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfraConfigBroker) DeepCopyInto(out *InfraConfigBroker) {
	*out = *in
	out.Resources = in.Resources
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfraConfigBroker.
func (in *InfraConfigBroker) DeepCopy() *InfraConfigBroker {
	if in == nil {
		return nil
	}
	out := new(InfraConfigBroker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfraConfigResources) DeepCopyInto(out *InfraConfigResources) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfraConfigResources.
func (in *InfraConfigResources) DeepCopy() *InfraConfigResources {
	if in == nil {
		return nil
	}
	out := new(InfraConfigResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfraConfigRouter) DeepCopyInto(out *InfraConfigRouter) {
	*out = *in
	out.Resources = in.Resources
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfraConfigRouter.
func (in *InfraConfigRouter) DeepCopy() *InfraConfigRouter {
	if in == nil {
		return nil
	}
	out := new(InfraConfigRouter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StandardInfraConfig) DeepCopyInto(out *StandardInfraConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StandardInfraConfig.
func (in *StandardInfraConfig) DeepCopy() *StandardInfraConfig {
	if in == nil {
		return nil
	}
	out := new(StandardInfraConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StandardInfraConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StandardInfraConfigList) DeepCopyInto(out *StandardInfraConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StandardInfraConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StandardInfraConfigList.
func (in *StandardInfraConfigList) DeepCopy() *StandardInfraConfigList {
	if in == nil {
		return nil
	}
	out := new(StandardInfraConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StandardInfraConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StandardInfraConfigSpec) DeepCopyInto(out *StandardInfraConfigSpec) {
	*out = *in
	out.Admin = in.Admin
	out.Broker = in.Broker
	out.Router = in.Router
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StandardInfraConfigSpec.
func (in *StandardInfraConfigSpec) DeepCopy() *StandardInfraConfigSpec {
	if in == nil {
		return nil
	}
	out := new(StandardInfraConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StandardInfraConfigStatus) DeepCopyInto(out *StandardInfraConfigStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StandardInfraConfigStatus.
func (in *StandardInfraConfigStatus) DeepCopy() *StandardInfraConfigStatus {
	if in == nil {
		return nil
	}
	out := new(StandardInfraConfigStatus)
	in.DeepCopyInto(out)
	return out
}
