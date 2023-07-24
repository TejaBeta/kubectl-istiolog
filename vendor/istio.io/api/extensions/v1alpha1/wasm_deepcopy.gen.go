// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package v1alpha1

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using WasmPlugin within kubernetes types, where deepcopy-gen is used.
func (in *WasmPlugin) DeepCopyInto(out *WasmPlugin) {
	p := proto.Clone(in).(*WasmPlugin)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WasmPlugin. Required by controller-gen.
func (in *WasmPlugin) DeepCopy() *WasmPlugin {
	if in == nil {
		return nil
	}
	out := new(WasmPlugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new WasmPlugin. Required by controller-gen.
func (in *WasmPlugin) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using WasmPlugin_TrafficSelector within kubernetes types, where deepcopy-gen is used.
func (in *WasmPlugin_TrafficSelector) DeepCopyInto(out *WasmPlugin_TrafficSelector) {
	p := proto.Clone(in).(*WasmPlugin_TrafficSelector)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WasmPlugin_TrafficSelector. Required by controller-gen.
func (in *WasmPlugin_TrafficSelector) DeepCopy() *WasmPlugin_TrafficSelector {
	if in == nil {
		return nil
	}
	out := new(WasmPlugin_TrafficSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new WasmPlugin_TrafficSelector. Required by controller-gen.
func (in *WasmPlugin_TrafficSelector) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using VmConfig within kubernetes types, where deepcopy-gen is used.
func (in *VmConfig) DeepCopyInto(out *VmConfig) {
	p := proto.Clone(in).(*VmConfig)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VmConfig. Required by controller-gen.
func (in *VmConfig) DeepCopy() *VmConfig {
	if in == nil {
		return nil
	}
	out := new(VmConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new VmConfig. Required by controller-gen.
func (in *VmConfig) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using EnvVar within kubernetes types, where deepcopy-gen is used.
func (in *EnvVar) DeepCopyInto(out *EnvVar) {
	p := proto.Clone(in).(*EnvVar)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvVar. Required by controller-gen.
func (in *EnvVar) DeepCopy() *EnvVar {
	if in == nil {
		return nil
	}
	out := new(EnvVar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new EnvVar. Required by controller-gen.
func (in *EnvVar) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
