//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicesdatareplication

import "encoding/json"

func unmarshalDraModelCustomPropertiesClassification(rawMsg json.RawMessage) (DraModelCustomPropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DraModelCustomPropertiesClassification
	switch m["instanceType"] {
	case "VMware":
		b = &VMwareDraModelCustomProperties{}
	default:
		b = &DraModelCustomProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalEventModelCustomPropertiesClassification(rawMsg json.RawMessage) (EventModelCustomPropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b EventModelCustomPropertiesClassification
	switch m["instanceType"] {
	case "HyperVToAzStackHCI":
		b = &HyperVToAzStackHCIEventModelCustomProperties{}
	default:
		b = &EventModelCustomProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalFabricModelCustomPropertiesClassification(rawMsg json.RawMessage) (FabricModelCustomPropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b FabricModelCustomPropertiesClassification
	switch m["instanceType"] {
	case "AzStackHCI":
		b = &AzStackHCIFabricModelCustomProperties{}
	case "HyperVMigrate":
		b = &HyperVMigrateFabricModelCustomProperties{}
	case "VMwareMigrate":
		b = &VMwareMigrateFabricModelCustomProperties{}
	default:
		b = &FabricModelCustomProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalPlannedFailoverModelCustomPropertiesClassification(rawMsg json.RawMessage) (PlannedFailoverModelCustomPropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b PlannedFailoverModelCustomPropertiesClassification
	switch m["instanceType"] {
	case "HyperVToAzStackHCI":
		b = &HyperVToAzStackHCIPlannedFailoverModelCustomProperties{}
	case "VMwareToAzStackHCI":
		b = &VMwareToAzStackHCIPlannedFailoverModelCustomProperties{}
	default:
		b = &PlannedFailoverModelCustomProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalPolicyModelCustomPropertiesClassification(rawMsg json.RawMessage) (PolicyModelCustomPropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b PolicyModelCustomPropertiesClassification
	switch m["instanceType"] {
	case "HyperVToAzStackHCI":
		b = &HyperVToAzStackHCIPolicyModelCustomProperties{}
	case "VMwareToAzStackHCI":
		b = &VMwareToAzStackHCIPolicyModelCustomProperties{}
	default:
		b = &PolicyModelCustomProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalProtectedItemModelCustomPropertiesClassification(rawMsg json.RawMessage) (ProtectedItemModelCustomPropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ProtectedItemModelCustomPropertiesClassification
	switch m["instanceType"] {
	case "HyperVToAzStackHCI":
		b = &HyperVToAzStackHCIProtectedItemModelCustomProperties{}
	case "VMwareToAzStackHCI":
		b = &VMwareToAzStackHCIProtectedItemModelCustomProperties{}
	default:
		b = &ProtectedItemModelCustomProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalRecoveryPointModelCustomPropertiesClassification(rawMsg json.RawMessage) (RecoveryPointModelCustomPropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b RecoveryPointModelCustomPropertiesClassification
	switch m["instanceType"] {
	case "HyperVToAzStackHCI":
		b = &HyperVToAzStackHCIRecoveryPointModelCustomProperties{}
	default:
		b = &RecoveryPointModelCustomProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalReplicationExtensionModelCustomPropertiesClassification(rawMsg json.RawMessage) (ReplicationExtensionModelCustomPropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ReplicationExtensionModelCustomPropertiesClassification
	switch m["instanceType"] {
	case "HyperVToAzStackHCI":
		b = &HyperVToAzStackHCIReplicationExtensionModelCustomProperties{}
	case "VMwareToAzStackHCI":
		b = &VMwareToAzStackHCIReplicationExtensionModelCustomProperties{}
	default:
		b = &ReplicationExtensionModelCustomProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalWorkflowModelCustomPropertiesClassification(rawMsg json.RawMessage) (WorkflowModelCustomPropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b WorkflowModelCustomPropertiesClassification
	switch m["instanceType"] {
	case "FailoverWorkflowDetails":
		b = &FailoverWorkflowModelCustomProperties{}
	case "TestFailoverCleanupWorkflowDetails":
		b = &TestFailoverCleanupWorkflowModelCustomProperties{}
	case "TestFailoverWorkflowDetails":
		b = &TestFailoverWorkflowModelCustomProperties{}
	default:
		b = &WorkflowModelCustomProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}