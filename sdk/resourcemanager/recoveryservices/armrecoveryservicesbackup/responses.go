//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicesbackup

import "encoding/json"

// BMSPrepareDataMoveOperationResultClientGetResponse contains the response from method BMSPrepareDataMoveOperationResultClient.Get.
type BMSPrepareDataMoveOperationResultClientGetResponse struct {
	// Operation result response for Vault Storage Config
	VaultStorageConfigOperationResultResponseClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type BMSPrepareDataMoveOperationResultClientGetResponse.
func (b *BMSPrepareDataMoveOperationResultClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalVaultStorageConfigOperationResultResponseClassification(data)
	if err != nil {
		return err
	}
	b.VaultStorageConfigOperationResultResponseClassification = res
	return nil
}

// BackupEnginesClientGetResponse contains the response from method BackupEnginesClient.Get.
type BackupEnginesClientGetResponse struct {
	// The base backup engine class. All workload specific backup engines derive from this class.
	BackupEngineBaseResource
}

// BackupEnginesClientListResponse contains the response from method BackupEnginesClient.NewListPager.
type BackupEnginesClientListResponse struct {
	// List of BackupEngineBase resources
	BackupEngineBaseResourceList
}

// BackupJobsClientListResponse contains the response from method BackupJobsClient.NewListPager.
type BackupJobsClientListResponse struct {
	// List of Job resources
	JobResourceList
}

// BackupOperationResultsClientGetResponse contains the response from method BackupOperationResultsClient.Get.
type BackupOperationResultsClientGetResponse struct {
	// placeholder for future response values
}

// BackupOperationStatusesClientGetResponse contains the response from method BackupOperationStatusesClient.Get.
type BackupOperationStatusesClientGetResponse struct {
	// Operation status.
	OperationStatus
}

// BackupPoliciesClientListResponse contains the response from method BackupPoliciesClient.NewListPager.
type BackupPoliciesClientListResponse struct {
	// List of ProtectionPolicy resources
	ProtectionPolicyResourceList
}

// BackupProtectableItemsClientListResponse contains the response from method BackupProtectableItemsClient.NewListPager.
type BackupProtectableItemsClientListResponse struct {
	// List of WorkloadProtectableItem resources
	WorkloadProtectableItemResourceList
}

// BackupProtectedItemsClientListResponse contains the response from method BackupProtectedItemsClient.NewListPager.
type BackupProtectedItemsClientListResponse struct {
	// List of ProtectedItem resources
	ProtectedItemResourceList
}

// BackupProtectionContainersClientListResponse contains the response from method BackupProtectionContainersClient.NewListPager.
type BackupProtectionContainersClientListResponse struct {
	// List of ProtectionContainer resources
	ProtectionContainerResourceList
}

// BackupProtectionIntentClientListResponse contains the response from method BackupProtectionIntentClient.NewListPager.
type BackupProtectionIntentClientListResponse struct {
	// List of ProtectionIntent resources
	ProtectionIntentResourceList
}

// BackupResourceEncryptionConfigsClientGetResponse contains the response from method BackupResourceEncryptionConfigsClient.Get.
type BackupResourceEncryptionConfigsClientGetResponse struct {
	BackupResourceEncryptionConfigExtendedResource
}

// BackupResourceEncryptionConfigsClientUpdateResponse contains the response from method BackupResourceEncryptionConfigsClient.Update.
type BackupResourceEncryptionConfigsClientUpdateResponse struct {
	// placeholder for future response values
}

// BackupResourceStorageConfigsNonCRRClientGetResponse contains the response from method BackupResourceStorageConfigsNonCRRClient.Get.
type BackupResourceStorageConfigsNonCRRClientGetResponse struct {
	// The resource storage details.
	BackupResourceConfigResource
}

// BackupResourceStorageConfigsNonCRRClientPatchResponse contains the response from method BackupResourceStorageConfigsNonCRRClient.Patch.
type BackupResourceStorageConfigsNonCRRClientPatchResponse struct {
	// placeholder for future response values
}

// BackupResourceStorageConfigsNonCRRClientUpdateResponse contains the response from method BackupResourceStorageConfigsNonCRRClient.Update.
type BackupResourceStorageConfigsNonCRRClientUpdateResponse struct {
	// The resource storage details.
	BackupResourceConfigResource
}

// BackupResourceVaultConfigsClientGetResponse contains the response from method BackupResourceVaultConfigsClient.Get.
type BackupResourceVaultConfigsClientGetResponse struct {
	// Backup resource vault config details.
	BackupResourceVaultConfigResource
}

// BackupResourceVaultConfigsClientPutResponse contains the response from method BackupResourceVaultConfigsClient.Put.
type BackupResourceVaultConfigsClientPutResponse struct {
	// Backup resource vault config details.
	BackupResourceVaultConfigResource
}

// BackupResourceVaultConfigsClientUpdateResponse contains the response from method BackupResourceVaultConfigsClient.Update.
type BackupResourceVaultConfigsClientUpdateResponse struct {
	// Backup resource vault config details.
	BackupResourceVaultConfigResource
}

// BackupStatusClientGetResponse contains the response from method BackupStatusClient.Get.
type BackupStatusClientGetResponse struct {
	// BackupStatus response.
	BackupStatusResponse
}

// BackupUsageSummariesClientListResponse contains the response from method BackupUsageSummariesClient.NewListPager.
type BackupUsageSummariesClientListResponse struct {
	// Backup management usage for vault.
	BackupManagementUsageList
}

// BackupWorkloadItemsClientListResponse contains the response from method BackupWorkloadItemsClient.NewListPager.
type BackupWorkloadItemsClientListResponse struct {
	// List of WorkloadItem resources
	WorkloadItemResourceList
}

// BackupsClientTriggerResponse contains the response from method BackupsClient.Trigger.
type BackupsClientTriggerResponse struct {
	// placeholder for future response values
}

// ClientBMSPrepareDataMoveResponse contains the response from method Client.BeginBMSPrepareDataMove.
type ClientBMSPrepareDataMoveResponse struct {
	// placeholder for future response values
}

// ClientBMSTriggerDataMoveResponse contains the response from method Client.BeginBMSTriggerDataMove.
type ClientBMSTriggerDataMoveResponse struct {
	// placeholder for future response values
}

// ClientGetOperationStatusResponse contains the response from method Client.GetOperationStatus.
type ClientGetOperationStatusResponse struct {
	// Operation status.
	OperationStatus
}

// ClientMoveRecoveryPointResponse contains the response from method Client.BeginMoveRecoveryPoint.
type ClientMoveRecoveryPointResponse struct {
	// placeholder for future response values
}

// DeletedProtectionContainersClientListResponse contains the response from method DeletedProtectionContainersClient.NewListPager.
type DeletedProtectionContainersClientListResponse struct {
	// List of ProtectionContainer resources
	ProtectionContainerResourceList
}

// ExportJobsOperationResultsClientGetResponse contains the response from method ExportJobsOperationResultsClient.Get.
type ExportJobsOperationResultsClientGetResponse struct {
	// Base class for operation result info.
	OperationResultInfoBaseResource
}

// FeatureSupportClientValidateResponse contains the response from method FeatureSupportClient.Validate.
type FeatureSupportClientValidateResponse struct {
	// Response for feature support requests for Azure IaasVm
	AzureVMResourceFeatureSupportResponse
}

// FetchTieringCostClientPostResponse contains the response from method FetchTieringCostClient.BeginPost.
type FetchTieringCostClientPostResponse struct {
	// Base class for tiering cost response
	TieringCostInfoClassification
}

// MarshalJSON implements the json.Marshaller interface for type FetchTieringCostClientPostResponse.
func (f FetchTieringCostClientPostResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.TieringCostInfoClassification)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type FetchTieringCostClientPostResponse.
func (f *FetchTieringCostClientPostResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalTieringCostInfoClassification(data)
	if err != nil {
		return err
	}
	f.TieringCostInfoClassification = res
	return nil
}

// GetTieringCostOperationResultClientGetResponse contains the response from method GetTieringCostOperationResultClient.Get.
type GetTieringCostOperationResultClientGetResponse struct {
	// Base class for tiering cost response
	TieringCostInfoClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type GetTieringCostOperationResultClientGetResponse.
func (g *GetTieringCostOperationResultClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalTieringCostInfoClassification(data)
	if err != nil {
		return err
	}
	g.TieringCostInfoClassification = res
	return nil
}

// ItemLevelRecoveryConnectionsClientProvisionResponse contains the response from method ItemLevelRecoveryConnectionsClient.Provision.
type ItemLevelRecoveryConnectionsClientProvisionResponse struct {
	// placeholder for future response values
}

// ItemLevelRecoveryConnectionsClientRevokeResponse contains the response from method ItemLevelRecoveryConnectionsClient.Revoke.
type ItemLevelRecoveryConnectionsClientRevokeResponse struct {
	// placeholder for future response values
}

// JobCancellationsClientTriggerResponse contains the response from method JobCancellationsClient.Trigger.
type JobCancellationsClientTriggerResponse struct {
	// placeholder for future response values
}

// JobDetailsClientGetResponse contains the response from method JobDetailsClient.Get.
type JobDetailsClientGetResponse struct {
	// Defines workload agnostic properties for a job.
	JobResource
}

// JobOperationResultsClientGetResponse contains the response from method JobOperationResultsClient.Get.
type JobOperationResultsClientGetResponse struct {
	// placeholder for future response values
}

// JobsClientExportResponse contains the response from method JobsClient.Export.
type JobsClientExportResponse struct {
	// placeholder for future response values
}

// OperationClientValidateResponse contains the response from method OperationClient.Validate.
type OperationClientValidateResponse struct {
	ValidateOperationsResponse
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// Operations List response which contains list of available APIs.
	ClientDiscoveryResponse
}

// PrivateEndpointClientGetOperationStatusResponse contains the response from method PrivateEndpointClient.GetOperationStatus.
type PrivateEndpointClientGetOperationStatusResponse struct {
	// Operation status.
	OperationStatus
}

// PrivateEndpointConnectionClientDeleteResponse contains the response from method PrivateEndpointConnectionClient.BeginDelete.
type PrivateEndpointConnectionClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionClientGetResponse contains the response from method PrivateEndpointConnectionClient.Get.
type PrivateEndpointConnectionClientGetResponse struct {
	// Private Endpoint Connection Response Properties
	PrivateEndpointConnectionResource
}

// PrivateEndpointConnectionClientPutResponse contains the response from method PrivateEndpointConnectionClient.BeginPut.
type PrivateEndpointConnectionClientPutResponse struct {
	// Private Endpoint Connection Response Properties
	PrivateEndpointConnectionResource
}

// ProtectableContainersClientListResponse contains the response from method ProtectableContainersClient.NewListPager.
type ProtectableContainersClientListResponse struct {
	// List of ProtectableContainer resources
	ProtectableContainerResourceList
}

// ProtectedItemOperationResultsClientGetResponse contains the response from method ProtectedItemOperationResultsClient.Get.
type ProtectedItemOperationResultsClientGetResponse struct {
	// Base class for backup items.
	ProtectedItemResource
}

// ProtectedItemOperationStatusesClientGetResponse contains the response from method ProtectedItemOperationStatusesClient.Get.
type ProtectedItemOperationStatusesClientGetResponse struct {
	// Operation status.
	OperationStatus
}

// ProtectedItemsClientCreateOrUpdateResponse contains the response from method ProtectedItemsClient.CreateOrUpdate.
type ProtectedItemsClientCreateOrUpdateResponse struct {
	// Base class for backup items.
	ProtectedItemResource
}

// ProtectedItemsClientDeleteResponse contains the response from method ProtectedItemsClient.Delete.
type ProtectedItemsClientDeleteResponse struct {
	// placeholder for future response values
}

// ProtectedItemsClientGetResponse contains the response from method ProtectedItemsClient.Get.
type ProtectedItemsClientGetResponse struct {
	// Base class for backup items.
	ProtectedItemResource
}

// ProtectionContainerOperationResultsClientGetResponse contains the response from method ProtectionContainerOperationResultsClient.Get.
type ProtectionContainerOperationResultsClientGetResponse struct {
	// Base class for container with backup items. Containers with specific workloads are derived from this class.
	ProtectionContainerResource
}

// ProtectionContainerRefreshOperationResultsClientGetResponse contains the response from method ProtectionContainerRefreshOperationResultsClient.Get.
type ProtectionContainerRefreshOperationResultsClientGetResponse struct {
	// placeholder for future response values
}

// ProtectionContainersClientGetResponse contains the response from method ProtectionContainersClient.Get.
type ProtectionContainersClientGetResponse struct {
	// Base class for container with backup items. Containers with specific workloads are derived from this class.
	ProtectionContainerResource
}

// ProtectionContainersClientInquireResponse contains the response from method ProtectionContainersClient.Inquire.
type ProtectionContainersClientInquireResponse struct {
	// placeholder for future response values
}

// ProtectionContainersClientRefreshResponse contains the response from method ProtectionContainersClient.Refresh.
type ProtectionContainersClientRefreshResponse struct {
	// placeholder for future response values
}

// ProtectionContainersClientRegisterResponse contains the response from method ProtectionContainersClient.BeginRegister.
type ProtectionContainersClientRegisterResponse struct {
	// Base class for container with backup items. Containers with specific workloads are derived from this class.
	ProtectionContainerResource
}

// ProtectionContainersClientUnregisterResponse contains the response from method ProtectionContainersClient.Unregister.
type ProtectionContainersClientUnregisterResponse struct {
	// placeholder for future response values
}

// ProtectionIntentClientCreateOrUpdateResponse contains the response from method ProtectionIntentClient.CreateOrUpdate.
type ProtectionIntentClientCreateOrUpdateResponse struct {
	// Base class for backup ProtectionIntent.
	ProtectionIntentResource
}

// ProtectionIntentClientDeleteResponse contains the response from method ProtectionIntentClient.Delete.
type ProtectionIntentClientDeleteResponse struct {
	// placeholder for future response values
}

// ProtectionIntentClientGetResponse contains the response from method ProtectionIntentClient.Get.
type ProtectionIntentClientGetResponse struct {
	// Base class for backup ProtectionIntent.
	ProtectionIntentResource
}

// ProtectionIntentClientValidateResponse contains the response from method ProtectionIntentClient.Validate.
type ProtectionIntentClientValidateResponse struct {
	// Response contract for enable backup validation request
	PreValidateEnableBackupResponse
}

// ProtectionPoliciesClientCreateOrUpdateResponse contains the response from method ProtectionPoliciesClient.CreateOrUpdate.
type ProtectionPoliciesClientCreateOrUpdateResponse struct {
	// Base class for backup policy. Workload-specific backup policies are derived from this class.
	ProtectionPolicyResource
}

// ProtectionPoliciesClientDeleteResponse contains the response from method ProtectionPoliciesClient.BeginDelete.
type ProtectionPoliciesClientDeleteResponse struct {
	// placeholder for future response values
}

// ProtectionPoliciesClientGetResponse contains the response from method ProtectionPoliciesClient.Get.
type ProtectionPoliciesClientGetResponse struct {
	// Base class for backup policy. Workload-specific backup policies are derived from this class.
	ProtectionPolicyResource
}

// ProtectionPolicyOperationResultsClientGetResponse contains the response from method ProtectionPolicyOperationResultsClient.Get.
type ProtectionPolicyOperationResultsClientGetResponse struct {
	// Base class for backup policy. Workload-specific backup policies are derived from this class.
	ProtectionPolicyResource
}

// ProtectionPolicyOperationStatusesClientGetResponse contains the response from method ProtectionPolicyOperationStatusesClient.Get.
type ProtectionPolicyOperationStatusesClientGetResponse struct {
	// Operation status.
	OperationStatus
}

// RecoveryPointsClientGetResponse contains the response from method RecoveryPointsClient.Get.
type RecoveryPointsClientGetResponse struct {
	// Base class for backup copies. Workload-specific backup copies are derived from this class.
	RecoveryPointResource
}

// RecoveryPointsClientListResponse contains the response from method RecoveryPointsClient.NewListPager.
type RecoveryPointsClientListResponse struct {
	// List of RecoveryPoint resources
	RecoveryPointResourceList
}

// RecoveryPointsRecommendedForMoveClientListResponse contains the response from method RecoveryPointsRecommendedForMoveClient.NewListPager.
type RecoveryPointsRecommendedForMoveClientListResponse struct {
	// List of RecoveryPoint resources
	RecoveryPointResourceList
}

// ResourceGuardProxiesClientGetResponse contains the response from method ResourceGuardProxiesClient.NewGetPager.
type ResourceGuardProxiesClientGetResponse struct {
	// List of ResourceGuardProxyBase resources
	ResourceGuardProxyBaseResourceList
}

// ResourceGuardProxyClientDeleteResponse contains the response from method ResourceGuardProxyClient.Delete.
type ResourceGuardProxyClientDeleteResponse struct {
	// placeholder for future response values
}

// ResourceGuardProxyClientGetResponse contains the response from method ResourceGuardProxyClient.Get.
type ResourceGuardProxyClientGetResponse struct {
	ResourceGuardProxyBaseResource
}

// ResourceGuardProxyClientPutResponse contains the response from method ResourceGuardProxyClient.Put.
type ResourceGuardProxyClientPutResponse struct {
	ResourceGuardProxyBaseResource
}

// ResourceGuardProxyClientUnlockDeleteResponse contains the response from method ResourceGuardProxyClient.UnlockDelete.
type ResourceGuardProxyClientUnlockDeleteResponse struct {
	// Response of Unlock Delete API.
	UnlockDeleteResponse
}

// RestoresClientTriggerResponse contains the response from method RestoresClient.BeginTrigger.
type RestoresClientTriggerResponse struct {
	// placeholder for future response values
}

// SecurityPINsClientGetResponse contains the response from method SecurityPINsClient.Get.
type SecurityPINsClientGetResponse struct {
	// The token information details.
	TokenInformation
}

// TieringCostOperationStatusClientGetResponse contains the response from method TieringCostOperationStatusClient.Get.
type TieringCostOperationStatusClientGetResponse struct {
	// Operation status.
	OperationStatus
}

// ValidateOperationClientTriggerResponse contains the response from method ValidateOperationClient.BeginTrigger.
type ValidateOperationClientTriggerResponse struct {
	// placeholder for future response values
}

// ValidateOperationResultsClientGetResponse contains the response from method ValidateOperationResultsClient.Get.
type ValidateOperationResultsClientGetResponse struct {
	ValidateOperationsResponse
}

// ValidateOperationStatusesClientGetResponse contains the response from method ValidateOperationStatusesClient.Get.
type ValidateOperationStatusesClientGetResponse struct {
	// Operation status.
	OperationStatus
}