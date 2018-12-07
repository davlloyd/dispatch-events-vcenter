///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package main

import (
	"github.com/vmware/govmomi/vim25/types"
)

// NO TESTS

func processEventMetadata(e types.BaseEvent) interface{} {
	// TODO: Implement automated way of generating payload based on vSphere API WSDL
	switch concreteEvent := e.(type) {
	case *types.VmBeingCreatedEvent:
		return handleVMBeingCreatedEvent(concreteEvent)
	case *types.VmBeingDeployedEvent:
		return handleVMBeingDeployedEvent(concreteEvent)
	case *types.VmDeployedEvent:
		return handleVMDeployedEvent(concreteEvent)
	case *types.VmPoweredOnEvent:
		return handleVMPoweredOnEvent(concreteEvent)
	case *types.VmPoweredOffEvent:
		return handleVMPoweredOffEvent(concreteEvent)
	case *types.VmEvent:
		return handleVMEvent(concreteEvent)
	default:
		return handleUnknownEvent(concreteEvent)
	}
}

func handleVMBeingCreatedEvent(e *types.VmBeingCreatedEvent) interface{} {
	return struct {
		VMName  string `json:"vm_name"`
		VMID    string `json:"vm_id"`
		NumCPUs int32  `json:"num_cpus"`
		MemMB   int64  `json:"mem_mb"`
	}{
		VMName:  e.ConfigSpec.Name,
		VMID:    e.Vm.Vm.String(),
		NumCPUs: e.ConfigSpec.NumCPUs * e.ConfigSpec.NumCoresPerSocket,
		MemMB:   e.ConfigSpec.MemoryMB,
	}
}

func handleVMBeingDeployedEvent(e *types.VmBeingDeployedEvent) interface{} {
	return struct {
		VMName      string `json:"vm_name"`
		VMID        string `json:"vm_id"`
		SrcTemplate string `json:"src_template"`
	}{
		VMName:      e.Vm.Name,
		VMID:        e.Vm.Vm.String(),
		SrcTemplate: e.SrcTemplate.Name,
	}
}

func handleVMDeployedEvent(e *types.VmDeployedEvent) interface{} {
	return struct {
		VMName      string `json:"vm_name"`
		VMID        string `json:"vm_id"`
		SrcTemplate string `json:"src_template"`
	}{
		VMName:      e.Vm.Name,
		VMID:        e.Vm.Vm.String(),
		SrcTemplate: e.SrcTemplate.Name,
	}
}

func handleVMEvent(e *types.VmEvent) interface{} {
	return struct {
		VMName     string `json:"vm_name"`
		VMID       string `json:"vm_id"`
	}{
		VMName:     e.Vm.Name,
		VMID:       e.Vm.Vm.String(),
	}
}

func handleVMPoweredOnEvent(e *types.VmPoweredOnEvent) interface{} {
	return struct {
		VMName     string `json:"vm_name"`
		VMID       string `json:"vm_id"`
	}{
		VMName:     e.Vm.Name,
		VMID:       e.Vm.Vm.String(),
	}
}

func handleVMPoweredOffEvent(e *types.VmPoweredOffEvent) interface{} {
	return struct {
		VMName     string `json:"vm_name"`
		VMID       string `json:"vm_id"`
	}{
		VMName:     e.Vm.Name,
		VMID:       e.Vm.Vm.String(),
	}
}

func handleUnknownEvent(e *types.VmEvent) interface{} {
	return struct {
		VMName     string `json:"vm_name"`
		VMID       string `json:"vm_id"`
		Comment    string `json:"comment"`
		Dump       string `json:"dump"`
	}{
		VMName:     e.Vm.Name,
		VMID:       e.Vm.Vm.String(),
		Comment:    "Event Type not declared",
		Dump:       fmt.Sprintf("%v", e),
	}
}
