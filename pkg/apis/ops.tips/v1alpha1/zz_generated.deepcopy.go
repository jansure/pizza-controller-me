// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PizzaCustomer) DeepCopyInto(out *PizzaCustomer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PizzaCustomer.
func (in *PizzaCustomer) DeepCopy() *PizzaCustomer {
	if in == nil {
		return nil
	}
	out := new(PizzaCustomer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PizzaCustomer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PizzaCustomerList) DeepCopyInto(out *PizzaCustomerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PizzaCustomer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PizzaCustomerList.
func (in *PizzaCustomerList) DeepCopy() *PizzaCustomerList {
	if in == nil {
		return nil
	}
	out := new(PizzaCustomerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PizzaCustomerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PizzaCustomerSpec) DeepCopyInto(out *PizzaCustomerSpec) {
	*out = *in
	out.CreditCardSecretRef = in.CreditCardSecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PizzaCustomerSpec.
func (in *PizzaCustomerSpec) DeepCopy() *PizzaCustomerSpec {
	if in == nil {
		return nil
	}
	out := new(PizzaCustomerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PizzaCustomerStatus) DeepCopyInto(out *PizzaCustomerStatus) {
	*out = *in
	out.ClosestStoreRef = in.ClosestStoreRef
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PizzaCustomerStatus.
func (in *PizzaCustomerStatus) DeepCopy() *PizzaCustomerStatus {
	if in == nil {
		return nil
	}
	out := new(PizzaCustomerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PizzaOrder) DeepCopyInto(out *PizzaOrder) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PizzaOrder.
func (in *PizzaOrder) DeepCopy() *PizzaOrder {
	if in == nil {
		return nil
	}
	out := new(PizzaOrder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PizzaOrder) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PizzaOrderList) DeepCopyInto(out *PizzaOrderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PizzaOrder, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PizzaOrderList.
func (in *PizzaOrderList) DeepCopy() *PizzaOrderList {
	if in == nil {
		return nil
	}
	out := new(PizzaOrderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PizzaOrderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PizzaOrderProduct) DeepCopyInto(out *PizzaOrderProduct) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PizzaOrderProduct.
func (in *PizzaOrderProduct) DeepCopy() *PizzaOrderProduct {
	if in == nil {
		return nil
	}
	out := new(PizzaOrderProduct)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PizzaOrderSpec) DeepCopyInto(out *PizzaOrderSpec) {
	*out = *in
	out.StoreRef = in.StoreRef
	out.CustomerRef = in.CustomerRef
	if in.Products != nil {
		in, out := &in.Products, &out.Products
		*out = make([]PizzaOrderProduct, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PizzaOrderSpec.
func (in *PizzaOrderSpec) DeepCopy() *PizzaOrderSpec {
	if in == nil {
		return nil
	}
	out := new(PizzaOrderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PizzaOrderStatus) DeepCopyInto(out *PizzaOrderStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PizzaOrderStatus.
func (in *PizzaOrderStatus) DeepCopy() *PizzaOrderStatus {
	if in == nil {
		return nil
	}
	out := new(PizzaOrderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PizzaStore) DeepCopyInto(out *PizzaStore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PizzaStore.
func (in *PizzaStore) DeepCopy() *PizzaStore {
	if in == nil {
		return nil
	}
	out := new(PizzaStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PizzaStore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PizzaStoreList) DeepCopyInto(out *PizzaStoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PizzaStore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PizzaStoreList.
func (in *PizzaStoreList) DeepCopy() *PizzaStoreList {
	if in == nil {
		return nil
	}
	out := new(PizzaStoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PizzaStoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PizzaStoreProduct) DeepCopyInto(out *PizzaStoreProduct) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PizzaStoreProduct.
func (in *PizzaStoreProduct) DeepCopy() *PizzaStoreProduct {
	if in == nil {
		return nil
	}
	out := new(PizzaStoreProduct)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PizzaStoreSpec) DeepCopyInto(out *PizzaStoreSpec) {
	*out = *in
	if in.Products != nil {
		in, out := &in.Products, &out.Products
		*out = make([]PizzaStoreProduct, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PizzaStoreSpec.
func (in *PizzaStoreSpec) DeepCopy() *PizzaStoreSpec {
	if in == nil {
		return nil
	}
	out := new(PizzaStoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PizzaStoreStatus) DeepCopyInto(out *PizzaStoreStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PizzaStoreStatus.
func (in *PizzaStoreStatus) DeepCopy() *PizzaStoreStatus {
	if in == nil {
		return nil
	}
	out := new(PizzaStoreStatus)
	in.DeepCopyInto(out)
	return out
}
