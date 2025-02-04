//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/intel/nri-resmgr/pkg/apis/resmgr"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Adjustment) DeepCopyInto(out *Adjustment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Adjustment.
func (in *Adjustment) DeepCopy() *Adjustment {
	if in == nil {
		return nil
	}
	out := new(Adjustment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Adjustment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdjustmentList) DeepCopyInto(out *AdjustmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Adjustment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdjustmentList.
func (in *AdjustmentList) DeepCopy() *AdjustmentList {
	if in == nil {
		return nil
	}
	out := new(AdjustmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AdjustmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdjustmentNodeStatus) DeepCopyInto(out *AdjustmentNodeStatus) {
	*out = *in
	if in.Errors != nil {
		in, out := &in.Errors, &out.Errors
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdjustmentNodeStatus.
func (in *AdjustmentNodeStatus) DeepCopy() *AdjustmentNodeStatus {
	if in == nil {
		return nil
	}
	out := new(AdjustmentNodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdjustmentScope) DeepCopyInto(out *AdjustmentScope) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]*resmgr.Expression, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto((*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdjustmentScope.
func (in *AdjustmentScope) DeepCopy() *AdjustmentScope {
	if in == nil {
		return nil
	}
	out := new(AdjustmentScope)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdjustmentSpec) DeepCopyInto(out *AdjustmentSpec) {
	*out = *in
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = make([]AdjustmentScope, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Classes != nil {
		in, out := &in.Classes, &out.Classes
		*out = new(Classes)
		(*in).DeepCopyInto(*out)
	}
	if in.ToptierLimit != nil {
		in, out := &in.ToptierLimit, &out.ToptierLimit
		x := (*in).DeepCopy()
		*out = &x
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdjustmentSpec.
func (in *AdjustmentSpec) DeepCopy() *AdjustmentSpec {
	if in == nil {
		return nil
	}
	out := new(AdjustmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdjustmentStatus) DeepCopyInto(out *AdjustmentStatus) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make(map[string]AdjustmentNodeStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdjustmentStatus.
func (in *AdjustmentStatus) DeepCopy() *AdjustmentStatus {
	if in == nil {
		return nil
	}
	out := new(AdjustmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Classes) DeepCopyInto(out *Classes) {
	*out = *in
	if in.BlockIO != nil {
		in, out := &in.BlockIO, &out.BlockIO
		*out = new(string)
		**out = **in
	}
	if in.RDT != nil {
		in, out := &in.RDT, &out.RDT
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Classes.
func (in *Classes) DeepCopy() *Classes {
	if in == nil {
		return nil
	}
	out := new(Classes)
	in.DeepCopyInto(out)
	return out
}
