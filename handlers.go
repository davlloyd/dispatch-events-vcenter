///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package main

import (
	"github.com/vmware/govmomi/vim25/types"
	"reflect"
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
	case *types.VmCloneEvent:
		return handleVMCloneEvent(concreteEvent)
	case *types.VmDeployFailedEvent:
		return handleVMDeployFailedEvent(concreteEvent)
	case *types.VmEvent:
		return handleVMEvent(concreteEvent)
	case *types.VmFailedToPowerOffEvent:
		return handleVMFailedToPowerOffEvent(concreteEvent)
	case *types.VmFailedToPowerOnEvent:
		return handleVMFailedToPowerOnEvent(concreteEvent)
	case *types.VmFailedToRebootGuestEvent:
		return handleVMFailedToRebootGuestEvent(concreteEvent)
	case *types.VmFailedToResetEvent:
		return handleVMFailedToResetEvent(concreteEvent)
	case *types.VmFailedToShutdownGuestEvent:
		return handleVMFailedToShutdownGuestEvent(concreteEvent)
	case *types.VmFailedToStandbyGuestEvent:
		return handleVMFailedToStandbyGuestEvent(concreteEvent)
	case *types.VmFailedToSuspendEvent:
		return handleVMFailedToSuspendEvent(concreteEvent)
	case *types.VmFailoverFailed:
		return handleVMEvent(concreteEvent)
	default:
		return handleUnknownEvent(reflect.TypeOf(e).String())
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

func handleVMCloneEvent(e *types.VmEvent) interface{} {
	return struct {
		VMName     string `json:"vm_name"`
		VMID       string `json:"vm_id"`
	}{
		VMName:     e.Vm.Name,
		VMID:       e.Vm.Vm.String(),
	}
}

func handleVMDeployFailedEvent(e *types.VmEvent) interface{} {
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
		VMName     	string `json:"vm_name"`
		VMID		string `json:"vm_id"`
	}{
		VMName:     e.Vm.Name,
		VMID:       e.Vm.Vm.String(),
	}
}

// Failure Events


func handleVMFailedToPowerOffEvent(e *types.VmPoweredOffEvent) interface{} {
	return struct {
		VMName     	string `json:"vm_name"`
		VMID		string `json:"vm_id"`
	}{
		VMName:     e.Vm.Name,
		VMID:       e.Vm.Vm.String(),
	}
}

func handleVMFailedToPowerOnEvent(e *types.VmPoweredOffEvent) interface{} {
	return struct {
		VMName     	string `json:"vm_name"`
		VMID		string `json:"vm_id"`
	}{
		VMName:     e.Vm.Name,
		VMID:       e.Vm.Vm.String(),
	}
}
func handleVMFailedToRebootGuestEvent(e *types.VmPoweredOffEvent) interface{} {
	return struct {
		VMName     	string `json:"vm_name"`
		VMID		string `json:"vm_id"`
	}{
		VMName:     e.Vm.Name,
		VMID:       e.Vm.Vm.String(),
	}
}

func handleVMFailedToResetEvent(e *types.VmPoweredOffEvent) interface{} {
	return struct {
		VMName     	string `json:"vm_name"`
		VMID		string `json:"vm_id"`
	}{
		VMName:     e.Vm.Name,
		VMID:       e.Vm.Vm.String(),
	}
}

func handleVMFailedToSuspendEvent(e *types.VmPoweredOffEvent) interface{} {
	return struct {
		VMName     	string `json:"vm_name"`
		VMID		string `json:"vm_id"`
	}{
		VMName:     e.Vm.Name,
		VMID:       e.Vm.Vm.String(),
	}
}

func handleVMFailedToShutdownGuestEvent(e *types.VmPoweredOffEvent) interface{} {
	return struct {
		VMName     	string `json:"vm_name"`
		VMID		string `json:"vm_id"`
	}{
		VMName:     e.Vm.Name,
		VMID:       e.Vm.Vm.String(),
	}
}

func handleVMFailedToStandbyGuestEvent(e *types.VmPoweredOffEvent) interface{} {
	return struct {
		VMName     	string `json:"vm_name"`
		VMID		string `json:"vm_id"`
	}{
		VMName:     e.Vm.Name,
		VMID:       e.Vm.Vm.String(),
	}
}

// Catch-all event
func handleUnknownEvent(eventtype string) interface{} {
	return struct {
		EventType	string `json:"eventtype"`
		Comment    	string `json:"comment"`
	}{
		EventType:  eventtype,
		Comment:    "Event Type not declared",
	}
}


