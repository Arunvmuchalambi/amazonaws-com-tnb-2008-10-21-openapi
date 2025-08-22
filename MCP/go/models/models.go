package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// GetSolVnfInfo represents the GetSolVnfInfo schema from the OpenAPI specification
type GetSolVnfInfo struct {
	Vnfstate interface{} `json:"vnfState,omitempty"`
	Vnfcresourceinfo interface{} `json:"vnfcResourceInfo,omitempty"`
}

// GetSolNetworkPackageContentOutput represents the GetSolNetworkPackageContentOutput schema from the OpenAPI specification
type GetSolNetworkPackageContentOutput struct {
	Nsdcontent interface{} `json:"nsdContent,omitempty"`
}

// DeleteSolNetworkPackageInput represents the DeleteSolNetworkPackageInput schema from the OpenAPI specification
type DeleteSolNetworkPackageInput struct {
}

// GetSolNetworkInstanceInput represents the GetSolNetworkInstanceInput schema from the OpenAPI specification
type GetSolNetworkInstanceInput struct {
}

// ToscaOverride represents the ToscaOverride schema from the OpenAPI specification
type ToscaOverride struct {
	Defaultvalue interface{} `json:"defaultValue,omitempty"`
	Name interface{} `json:"name,omitempty"`
}

// GetSolFunctionPackageInput represents the GetSolFunctionPackageInput schema from the OpenAPI specification
type GetSolFunctionPackageInput struct {
}

// CreateSolNetworkPackageInput represents the CreateSolNetworkPackageInput schema from the OpenAPI specification
type CreateSolNetworkPackageInput struct {
	Tags interface{} `json:"tags,omitempty"`
}

// GetSolFunctionPackageContentOutput represents the GetSolFunctionPackageContentOutput schema from the OpenAPI specification
type GetSolFunctionPackageContentOutput struct {
	Packagecontent interface{} `json:"packageContent,omitempty"`
}

// StringMap represents the StringMap schema from the OpenAPI specification
type StringMap struct {
}

// InstantiateSolNetworkInstanceInput represents the InstantiateSolNetworkInstanceInput schema from the OpenAPI specification
type InstantiateSolNetworkInstanceInput struct {
	Additionalparamsforns interface{} `json:"additionalParamsForNs,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
}

// ListSolFunctionPackageInfo represents the ListSolFunctionPackageInfo schema from the OpenAPI specification
type ListSolFunctionPackageInfo struct {
	Vnfdversion interface{} `json:"vnfdVersion,omitempty"`
	Arn interface{} `json:"arn"`
	Vnfdid interface{} `json:"vnfdId,omitempty"`
	Id interface{} `json:"id"`
	Metadata interface{} `json:"metadata,omitempty"`
	Operationalstate interface{} `json:"operationalState"`
	Usagestate interface{} `json:"usageState"`
	Vnfprovider interface{} `json:"vnfProvider,omitempty"`
	Onboardingstate interface{} `json:"onboardingState"`
	Vnfproductname interface{} `json:"vnfProductName,omitempty"`
}

// GetSolFunctionInstanceOutput represents the GetSolFunctionInstanceOutput schema from the OpenAPI specification
type GetSolFunctionInstanceOutput struct {
	Vnfprovider interface{} `json:"vnfProvider,omitempty"`
	Vnfdid interface{} `json:"vnfdId"`
	Arn interface{} `json:"arn"`
	Instantiationstate interface{} `json:"instantiationState"`
	Tags interface{} `json:"tags,omitempty"`
	Id interface{} `json:"id"`
	Nsinstanceid interface{} `json:"nsInstanceId"`
	Vnfpkgid interface{} `json:"vnfPkgId"`
	Vnfdversion interface{} `json:"vnfdVersion,omitempty"`
	Metadata GetSolFunctionInstanceMetadata `json:"metadata"` // <p>The metadata of a network function instance.</p> <p>A network function instance is a function in a function package .</p>
	Instantiatedvnfinfo GetSolVnfInfo `json:"instantiatedVnfInfo,omitempty"` // <p>Information about the network function.</p> <p>A network function instance is a function in a function package .</p>
	Vnfproductname interface{} `json:"vnfProductName,omitempty"`
}

// GetSolNetworkPackageDescriptorInput represents the GetSolNetworkPackageDescriptorInput schema from the OpenAPI specification
type GetSolNetworkPackageDescriptorInput struct {
}

// UpdateSolNetworkPackageInput represents the UpdateSolNetworkPackageInput schema from the OpenAPI specification
type UpdateSolNetworkPackageInput struct {
	Nsdoperationalstate interface{} `json:"nsdOperationalState"`
}

// PutSolNetworkPackageContentOutput represents the PutSolNetworkPackageContentOutput schema from the OpenAPI specification
type PutSolNetworkPackageContentOutput struct {
	Metadata interface{} `json:"metadata"`
	Nsdid interface{} `json:"nsdId"`
	Nsdname interface{} `json:"nsdName"`
	Nsdversion interface{} `json:"nsdVersion"`
	Vnfpkgids interface{} `json:"vnfPkgIds"`
	Arn interface{} `json:"arn"`
	Id interface{} `json:"id"`
}

// GetSolNetworkOperationMetadata represents the GetSolNetworkOperationMetadata schema from the OpenAPI specification
type GetSolNetworkOperationMetadata struct {
	Createdat interface{} `json:"createdAt"`
	Lastmodified interface{} `json:"lastModified"`
}

// GetSolNetworkInstanceOutput represents the GetSolNetworkInstanceOutput schema from the OpenAPI specification
type GetSolNetworkInstanceOutput struct {
	Arn interface{} `json:"arn"`
	Id interface{} `json:"id"`
	Nsinstancedescription interface{} `json:"nsInstanceDescription"`
	Nsinstancename interface{} `json:"nsInstanceName"`
	Metadata GetSolNetworkInstanceMetadata `json:"metadata"` // <p>The metadata of a network instance.</p> <p>A network instance is a single network created in Amazon Web Services TNB that can be deployed and on which life-cycle operations (like terminate, update, and delete) can be performed.</p>
	Nsstate interface{} `json:"nsState,omitempty"`
	Nsdid interface{} `json:"nsdId"`
	Nsdinfoid interface{} `json:"nsdInfoId"`
	Lcmopinfo LcmOperationInfo `json:"lcmOpInfo,omitempty"` // <p>Lifecycle management operation details on the network instance.</p> <p>Lifecycle management operations are deploy, update, or delete operations.</p>
	Tags interface{} `json:"tags,omitempty"`
}

// GetSolVnfcResourceInfo represents the GetSolVnfcResourceInfo schema from the OpenAPI specification
type GetSolVnfcResourceInfo struct {
	Metadata interface{} `json:"metadata,omitempty"`
}

// NetworkArtifactMeta represents the NetworkArtifactMeta schema from the OpenAPI specification
type NetworkArtifactMeta struct {
	Overrides interface{} `json:"overrides,omitempty"`
}

// CreateSolFunctionPackageInput represents the CreateSolFunctionPackageInput schema from the OpenAPI specification
type CreateSolFunctionPackageInput struct {
	Tags interface{} `json:"tags,omitempty"`
}

// ValidateSolNetworkPackageContentInput represents the ValidateSolNetworkPackageContentInput schema from the OpenAPI specification
type ValidateSolNetworkPackageContentInput struct {
	File interface{} `json:"file"`
}

// GetSolFunctionPackageDescriptorInput represents the GetSolFunctionPackageDescriptorInput schema from the OpenAPI specification
type GetSolFunctionPackageDescriptorInput struct {
}

// GetSolNetworkInstanceMetadata represents the GetSolNetworkInstanceMetadata schema from the OpenAPI specification
type GetSolNetworkInstanceMetadata struct {
	Createdat interface{} `json:"createdAt"`
	Lastmodified interface{} `json:"lastModified"`
}

// ListSolNetworkPackagesInput represents the ListSolNetworkPackagesInput schema from the OpenAPI specification
type ListSolNetworkPackagesInput struct {
}

// ValidateSolFunctionPackageContentMetadata represents the ValidateSolFunctionPackageContentMetadata schema from the OpenAPI specification
type ValidateSolFunctionPackageContentMetadata struct {
	Vnfd FunctionArtifactMeta `json:"vnfd,omitempty"` // <p>Metadata for function package artifacts.</p> <p>Artifacts are the contents of the package descriptor file and the state of the package.</p>
}

// TagResourceInput represents the TagResourceInput schema from the OpenAPI specification
type TagResourceInput struct {
	Tags interface{} `json:"tags"`
}

// ListSolNetworkInstanceInfo represents the ListSolNetworkInstanceInfo schema from the OpenAPI specification
type ListSolNetworkInstanceInfo struct {
	Nsinstancedescription interface{} `json:"nsInstanceDescription"`
	Nsinstancename interface{} `json:"nsInstanceName"`
	Nsstate interface{} `json:"nsState"`
	Nsdid interface{} `json:"nsdId"`
	Nsdinfoid interface{} `json:"nsdInfoId"`
	Arn interface{} `json:"arn"`
	Id interface{} `json:"id"`
	Metadata interface{} `json:"metadata"`
}

// ListTagsForResourceInput represents the ListTagsForResourceInput schema from the OpenAPI specification
type ListTagsForResourceInput struct {
}

// ListSolNetworkOperationsOutput represents the ListSolNetworkOperationsOutput schema from the OpenAPI specification
type ListSolNetworkOperationsOutput struct {
	Networkoperations interface{} `json:"networkOperations,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// ListSolFunctionInstancesOutput represents the ListSolFunctionInstancesOutput schema from the OpenAPI specification
type ListSolFunctionInstancesOutput struct {
	Functioninstances interface{} `json:"functionInstances,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// GetSolInstantiatedVnfInfo represents the GetSolInstantiatedVnfInfo schema from the OpenAPI specification
type GetSolInstantiatedVnfInfo struct {
	Vnfstate interface{} `json:"vnfState,omitempty"`
}

// UpdateSolNetworkInstanceInput represents the UpdateSolNetworkInstanceInput schema from the OpenAPI specification
type UpdateSolNetworkInstanceInput struct {
	Modifyvnfinfodata interface{} `json:"modifyVnfInfoData,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Updatetype interface{} `json:"updateType"`
}

// PutSolFunctionPackageContentMetadata represents the PutSolFunctionPackageContentMetadata schema from the OpenAPI specification
type PutSolFunctionPackageContentMetadata struct {
	Vnfd FunctionArtifactMeta `json:"vnfd,omitempty"` // <p>Metadata for function package artifacts.</p> <p>Artifacts are the contents of the package descriptor file and the state of the package.</p>
}

// TerminateSolNetworkInstanceOutput represents the TerminateSolNetworkInstanceOutput schema from the OpenAPI specification
type TerminateSolNetworkInstanceOutput struct {
	Nslcmopoccid interface{} `json:"nsLcmOpOccId,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
}

// CreateSolNetworkPackageOutput represents the CreateSolNetworkPackageOutput schema from the OpenAPI specification
type CreateSolNetworkPackageOutput struct {
	Arn interface{} `json:"arn"`
	Id interface{} `json:"id"`
	Nsdonboardingstate interface{} `json:"nsdOnboardingState"`
	Nsdoperationalstate interface{} `json:"nsdOperationalState"`
	Nsdusagestate interface{} `json:"nsdUsageState"`
	Tags interface{} `json:"tags,omitempty"`
}

// ListSolNetworkOperationsInfo represents the ListSolNetworkOperationsInfo schema from the OpenAPI specification
type ListSolNetworkOperationsInfo struct {
	Id interface{} `json:"id"`
	Lcmoperationtype interface{} `json:"lcmOperationType"`
	Metadata interface{} `json:"metadata,omitempty"`
	Nsinstanceid interface{} `json:"nsInstanceId"`
	Operationstate interface{} `json:"operationState"`
	Arn interface{} `json:"arn"`
	ErrorField interface{} `json:"error,omitempty"`
}

// UpdateSolNetworkModify represents the UpdateSolNetworkModify schema from the OpenAPI specification
type UpdateSolNetworkModify struct {
	Vnfinstanceid interface{} `json:"vnfInstanceId"`
	Vnfconfigurableproperties interface{} `json:"vnfConfigurableProperties"`
}

// CancelSolNetworkOperationInput represents the CancelSolNetworkOperationInput schema from the OpenAPI specification
type CancelSolNetworkOperationInput struct {
}

// GetSolNetworkPackageContentInput represents the GetSolNetworkPackageContentInput schema from the OpenAPI specification
type GetSolNetworkPackageContentInput struct {
}

// ErrorInfo represents the ErrorInfo schema from the OpenAPI specification
type ErrorInfo struct {
	Details interface{} `json:"details,omitempty"`
	Cause interface{} `json:"cause,omitempty"`
}

// GetSolVnfcResourceInfoMetadata represents the GetSolVnfcResourceInfoMetadata schema from the OpenAPI specification
type GetSolVnfcResourceInfoMetadata struct {
	Helmchart interface{} `json:"helmChart,omitempty"`
	Nodegroup interface{} `json:"nodeGroup,omitempty"`
	Cluster interface{} `json:"cluster,omitempty"`
}

// ListSolNetworkPackagesOutput represents the ListSolNetworkPackagesOutput schema from the OpenAPI specification
type ListSolNetworkPackagesOutput struct {
	Networkpackages interface{} `json:"networkPackages"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// TagMap represents the TagMap schema from the OpenAPI specification
type TagMap struct {
}

// ListSolFunctionInstanceMetadata represents the ListSolFunctionInstanceMetadata schema from the OpenAPI specification
type ListSolFunctionInstanceMetadata struct {
	Lastmodified interface{} `json:"lastModified"`
	Createdat interface{} `json:"createdAt"`
}

// PutSolNetworkPackageContentMetadata represents the PutSolNetworkPackageContentMetadata schema from the OpenAPI specification
type PutSolNetworkPackageContentMetadata struct {
	Nsd NetworkArtifactMeta `json:"nsd,omitempty"` // <p>Metadata for network package artifacts.</p> <p>Artifacts are the contents of the package descriptor file and the state of the package.</p>
}

// GetSolNetworkPackageDescriptorOutput represents the GetSolNetworkPackageDescriptorOutput schema from the OpenAPI specification
type GetSolNetworkPackageDescriptorOutput struct {
	Nsd interface{} `json:"nsd,omitempty"`
}

// UpdateSolNetworkPackageOutput represents the UpdateSolNetworkPackageOutput schema from the OpenAPI specification
type UpdateSolNetworkPackageOutput struct {
	Nsdoperationalstate interface{} `json:"nsdOperationalState"`
}

// InstantiateSolNetworkInstanceOutput represents the InstantiateSolNetworkInstanceOutput schema from the OpenAPI specification
type InstantiateSolNetworkInstanceOutput struct {
	Tags interface{} `json:"tags,omitempty"`
	Nslcmopoccid interface{} `json:"nsLcmOpOccId"`
}

// LcmOperationInfo represents the LcmOperationInfo schema from the OpenAPI specification
type LcmOperationInfo struct {
	Nslcmopoccid interface{} `json:"nsLcmOpOccId"`
}

// ProblemDetails represents the ProblemDetails schema from the OpenAPI specification
type ProblemDetails struct {
	Detail interface{} `json:"detail"`
	Title interface{} `json:"title,omitempty"`
}

// UntagResourceInput represents the UntagResourceInput schema from the OpenAPI specification
type UntagResourceInput struct {
}

// GetSolFunctionInstanceInput represents the GetSolFunctionInstanceInput schema from the OpenAPI specification
type GetSolFunctionInstanceInput struct {
}

// UpdateSolFunctionPackageOutput represents the UpdateSolFunctionPackageOutput schema from the OpenAPI specification
type UpdateSolFunctionPackageOutput struct {
	Operationalstate interface{} `json:"operationalState"`
}

// GetSolNetworkPackageMetadata represents the GetSolNetworkPackageMetadata schema from the OpenAPI specification
type GetSolNetworkPackageMetadata struct {
	Createdat interface{} `json:"createdAt"`
	Lastmodified interface{} `json:"lastModified"`
	Nsd interface{} `json:"nsd,omitempty"`
}

// ListSolFunctionInstanceInfo represents the ListSolFunctionInstanceInfo schema from the OpenAPI specification
type ListSolFunctionInstanceInfo struct {
	Instantiationstate interface{} `json:"instantiationState"`
	Metadata interface{} `json:"metadata"`
	Nsinstanceid interface{} `json:"nsInstanceId"`
	Vnfpkgid interface{} `json:"vnfPkgId"`
	Vnfpkgname interface{} `json:"vnfPkgName,omitempty"`
	Arn interface{} `json:"arn"`
	Id interface{} `json:"id"`
	Instantiatedvnfinfo GetSolInstantiatedVnfInfo `json:"instantiatedVnfInfo,omitempty"` // <p>Information about a network function.</p> <p>A network instance is a single network created in Amazon Web Services TNB that can be deployed and on which life-cycle operations (like terminate, update, and delete) can be performed.</p>
}

// FunctionArtifactMeta represents the FunctionArtifactMeta schema from the OpenAPI specification
type FunctionArtifactMeta struct {
	Overrides interface{} `json:"overrides,omitempty"`
}

// UpdateSolFunctionPackageInput represents the UpdateSolFunctionPackageInput schema from the OpenAPI specification
type UpdateSolFunctionPackageInput struct {
	Operationalstate interface{} `json:"operationalState"`
}

// ListSolFunctionPackageMetadata represents the ListSolFunctionPackageMetadata schema from the OpenAPI specification
type ListSolFunctionPackageMetadata struct {
	Createdat interface{} `json:"createdAt"`
	Lastmodified interface{} `json:"lastModified"`
}

// ListSolNetworkInstancesOutput represents the ListSolNetworkInstancesOutput schema from the OpenAPI specification
type ListSolNetworkInstancesOutput struct {
	Networkinstances interface{} `json:"networkInstances,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// UpdateSolNetworkInstanceOutput represents the UpdateSolNetworkInstanceOutput schema from the OpenAPI specification
type UpdateSolNetworkInstanceOutput struct {
	Nslcmopoccid interface{} `json:"nsLcmOpOccId,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
}

// ListSolFunctionPackagesOutput represents the ListSolFunctionPackagesOutput schema from the OpenAPI specification
type ListSolFunctionPackagesOutput struct {
	Functionpackages interface{} `json:"functionPackages"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// ListSolNetworkOperationsInput represents the ListSolNetworkOperationsInput schema from the OpenAPI specification
type ListSolNetworkOperationsInput struct {
}

// TerminateSolNetworkInstanceInput represents the TerminateSolNetworkInstanceInput schema from the OpenAPI specification
type TerminateSolNetworkInstanceInput struct {
	Tags interface{} `json:"tags,omitempty"`
}

// PutSolFunctionPackageContentInput represents the PutSolFunctionPackageContentInput schema from the OpenAPI specification
type PutSolFunctionPackageContentInput struct {
	File interface{} `json:"file"`
}

// ListTagsForResourceOutput represents the ListTagsForResourceOutput schema from the OpenAPI specification
type ListTagsForResourceOutput struct {
	Tags interface{} `json:"tags"`
}

// GetSolFunctionPackageOutput represents the GetSolFunctionPackageOutput schema from the OpenAPI specification
type GetSolFunctionPackageOutput struct {
	Metadata GetSolFunctionPackageMetadata `json:"metadata,omitempty"` // <p>Metadata related to the function package.</p> <p>A function package is a .zip file in CSAR (Cloud Service Archive) format that contains a network function (an ETSI standard telecommunication application) and function package descriptor that uses the TOSCA standard to describe how the network functions should run on your network.</p>
	Vnfdversion interface{} `json:"vnfdVersion,omitempty"`
	Arn interface{} `json:"arn"`
	Operationalstate interface{} `json:"operationalState"`
	Usagestate interface{} `json:"usageState"`
	Vnfdid interface{} `json:"vnfdId,omitempty"`
	Id interface{} `json:"id"`
	Onboardingstate interface{} `json:"onboardingState"`
	Tags interface{} `json:"tags,omitempty"`
	Vnfproductname interface{} `json:"vnfProductName,omitempty"`
	Vnfprovider interface{} `json:"vnfProvider,omitempty"`
}

// GetSolFunctionPackageMetadata represents the GetSolFunctionPackageMetadata schema from the OpenAPI specification
type GetSolFunctionPackageMetadata struct {
	Vnfd interface{} `json:"vnfd,omitempty"`
	Createdat interface{} `json:"createdAt"`
	Lastmodified interface{} `json:"lastModified"`
}

// ListSolNetworkInstancesInput represents the ListSolNetworkInstancesInput schema from the OpenAPI specification
type ListSolNetworkInstancesInput struct {
}

// ListSolNetworkPackageInfo represents the ListSolNetworkPackageInfo schema from the OpenAPI specification
type ListSolNetworkPackageInfo struct {
	Nsdid interface{} `json:"nsdId,omitempty"`
	Nsdusagestate interface{} `json:"nsdUsageState"`
	Nsdversion interface{} `json:"nsdVersion,omitempty"`
	Vnfpkgids interface{} `json:"vnfPkgIds,omitempty"`
	Nsddesigner interface{} `json:"nsdDesigner,omitempty"`
	Nsdinvariantid interface{} `json:"nsdInvariantId,omitempty"`
	Nsdname interface{} `json:"nsdName,omitempty"`
	Nsdonboardingstate interface{} `json:"nsdOnboardingState"`
	Nsdoperationalstate interface{} `json:"nsdOperationalState"`
	Arn interface{} `json:"arn"`
	Id interface{} `json:"id"`
	Metadata interface{} `json:"metadata"`
}

// Document represents the Document schema from the OpenAPI specification
type Document struct {
}

// ListSolFunctionPackagesInput represents the ListSolFunctionPackagesInput schema from the OpenAPI specification
type ListSolFunctionPackagesInput struct {
}

// DeleteSolFunctionPackageInput represents the DeleteSolFunctionPackageInput schema from the OpenAPI specification
type DeleteSolFunctionPackageInput struct {
}

// CreateSolNetworkInstanceInput represents the CreateSolNetworkInstanceInput schema from the OpenAPI specification
type CreateSolNetworkInstanceInput struct {
	Nsdescription interface{} `json:"nsDescription,omitempty"`
	Nsname interface{} `json:"nsName"`
	Nsdinfoid interface{} `json:"nsdInfoId"`
	Tags interface{} `json:"tags,omitempty"`
}

// TagResourceOutput represents the TagResourceOutput schema from the OpenAPI specification
type TagResourceOutput struct {
}

// ListSolNetworkInstanceMetadata represents the ListSolNetworkInstanceMetadata schema from the OpenAPI specification
type ListSolNetworkInstanceMetadata struct {
	Createdat interface{} `json:"createdAt"`
	Lastmodified interface{} `json:"lastModified"`
}

// UntagResourceOutput represents the UntagResourceOutput schema from the OpenAPI specification
type UntagResourceOutput struct {
}

// GetSolNetworkPackageInput represents the GetSolNetworkPackageInput schema from the OpenAPI specification
type GetSolNetworkPackageInput struct {
}

// PutSolNetworkPackageContentInput represents the PutSolNetworkPackageContentInput schema from the OpenAPI specification
type PutSolNetworkPackageContentInput struct {
	File interface{} `json:"file"`
}

// GetSolFunctionPackageContentInput represents the GetSolFunctionPackageContentInput schema from the OpenAPI specification
type GetSolFunctionPackageContentInput struct {
}

// ValidateSolFunctionPackageContentOutput represents the ValidateSolFunctionPackageContentOutput schema from the OpenAPI specification
type ValidateSolFunctionPackageContentOutput struct {
	Vnfdid interface{} `json:"vnfdId"`
	Vnfdversion interface{} `json:"vnfdVersion"`
	Id interface{} `json:"id"`
	Metadata interface{} `json:"metadata"`
	Vnfproductname interface{} `json:"vnfProductName"`
	Vnfprovider interface{} `json:"vnfProvider"`
}

// ValidateSolNetworkPackageContentMetadata represents the ValidateSolNetworkPackageContentMetadata schema from the OpenAPI specification
type ValidateSolNetworkPackageContentMetadata struct {
	Nsd NetworkArtifactMeta `json:"nsd,omitempty"` // <p>Metadata for network package artifacts.</p> <p>Artifacts are the contents of the package descriptor file and the state of the package.</p>
}

// ValidateSolNetworkPackageContentOutput represents the ValidateSolNetworkPackageContentOutput schema from the OpenAPI specification
type ValidateSolNetworkPackageContentOutput struct {
	Nsdid interface{} `json:"nsdId"`
	Nsdname interface{} `json:"nsdName"`
	Nsdversion interface{} `json:"nsdVersion"`
	Vnfpkgids interface{} `json:"vnfPkgIds"`
	Arn interface{} `json:"arn"`
	Id interface{} `json:"id"`
	Metadata interface{} `json:"metadata"`
}

// CreateSolNetworkInstanceOutput represents the CreateSolNetworkInstanceOutput schema from the OpenAPI specification
type CreateSolNetworkInstanceOutput struct {
	Arn interface{} `json:"arn"`
	Id interface{} `json:"id"`
	Nsinstancename interface{} `json:"nsInstanceName"`
	Nsdinfoid interface{} `json:"nsdInfoId"`
	Tags interface{} `json:"tags,omitempty"`
}

// GetSolNetworkPackageOutput represents the GetSolNetworkPackageOutput schema from the OpenAPI specification
type GetSolNetworkPackageOutput struct {
	Nsdname interface{} `json:"nsdName"`
	Nsdusagestate interface{} `json:"nsdUsageState"`
	Nsdoperationalstate interface{} `json:"nsdOperationalState"`
	Vnfpkgids interface{} `json:"vnfPkgIds"`
	Arn interface{} `json:"arn"`
	Metadata GetSolNetworkPackageMetadata `json:"metadata"` // <p>Metadata associated with a network package.</p> <p>A network package is a .zip file in CSAR (Cloud Service Archive) format defines the function packages you want to deploy and the Amazon Web Services infrastructure you want to deploy them on.</p>
	Nsdonboardingstate interface{} `json:"nsdOnboardingState"`
	Tags interface{} `json:"tags,omitempty"`
	Id interface{} `json:"id"`
	Nsdid interface{} `json:"nsdId"`
	Nsdversion interface{} `json:"nsdVersion"`
}

// GetSolNetworkOperationInput represents the GetSolNetworkOperationInput schema from the OpenAPI specification
type GetSolNetworkOperationInput struct {
}

// ListSolNetworkOperationsMetadata represents the ListSolNetworkOperationsMetadata schema from the OpenAPI specification
type ListSolNetworkOperationsMetadata struct {
	Createdat interface{} `json:"createdAt"`
	Lastmodified interface{} `json:"lastModified"`
}

// ListSolFunctionInstancesInput represents the ListSolFunctionInstancesInput schema from the OpenAPI specification
type ListSolFunctionInstancesInput struct {
}

// ListSolNetworkPackageMetadata represents the ListSolNetworkPackageMetadata schema from the OpenAPI specification
type ListSolNetworkPackageMetadata struct {
	Lastmodified interface{} `json:"lastModified"`
	Createdat interface{} `json:"createdAt"`
}

// GetSolFunctionPackageDescriptorOutput represents the GetSolFunctionPackageDescriptorOutput schema from the OpenAPI specification
type GetSolFunctionPackageDescriptorOutput struct {
	Vnfd interface{} `json:"vnfd,omitempty"`
}

// DeleteSolNetworkInstanceInput represents the DeleteSolNetworkInstanceInput schema from the OpenAPI specification
type DeleteSolNetworkInstanceInput struct {
}

// GetSolNetworkOperationTaskDetails represents the GetSolNetworkOperationTaskDetails schema from the OpenAPI specification
type GetSolNetworkOperationTaskDetails struct {
	Taskcontext interface{} `json:"taskContext,omitempty"`
	Taskendtime interface{} `json:"taskEndTime,omitempty"`
	Taskerrordetails interface{} `json:"taskErrorDetails,omitempty"`
	Taskname interface{} `json:"taskName,omitempty"`
	Taskstarttime interface{} `json:"taskStartTime,omitempty"`
	Taskstatus interface{} `json:"taskStatus,omitempty"`
}

// GetSolFunctionInstanceMetadata represents the GetSolFunctionInstanceMetadata schema from the OpenAPI specification
type GetSolFunctionInstanceMetadata struct {
	Createdat interface{} `json:"createdAt"`
	Lastmodified interface{} `json:"lastModified"`
}

// PutSolFunctionPackageContentOutput represents the PutSolFunctionPackageContentOutput schema from the OpenAPI specification
type PutSolFunctionPackageContentOutput struct {
	Vnfdid interface{} `json:"vnfdId"`
	Vnfdversion interface{} `json:"vnfdVersion"`
	Id interface{} `json:"id"`
	Metadata interface{} `json:"metadata"`
	Vnfproductname interface{} `json:"vnfProductName"`
	Vnfprovider interface{} `json:"vnfProvider"`
}

// GetSolNetworkOperationOutput represents the GetSolNetworkOperationOutput schema from the OpenAPI specification
type GetSolNetworkOperationOutput struct {
	Nsinstanceid interface{} `json:"nsInstanceId,omitempty"`
	Metadata interface{} `json:"metadata,omitempty"`
	Operationstate interface{} `json:"operationState,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Lcmoperationtype interface{} `json:"lcmOperationType,omitempty"`
	Tasks interface{} `json:"tasks,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Arn interface{} `json:"arn"`
	ErrorField interface{} `json:"error,omitempty"`
}

// CreateSolFunctionPackageOutput represents the CreateSolFunctionPackageOutput schema from the OpenAPI specification
type CreateSolFunctionPackageOutput struct {
	Tags interface{} `json:"tags,omitempty"`
	Usagestate interface{} `json:"usageState"`
	Arn interface{} `json:"arn"`
	Id interface{} `json:"id"`
	Onboardingstate interface{} `json:"onboardingState"`
	Operationalstate interface{} `json:"operationalState"`
}

// ValidateSolFunctionPackageContentInput represents the ValidateSolFunctionPackageContentInput schema from the OpenAPI specification
type ValidateSolFunctionPackageContentInput struct {
	File interface{} `json:"file"`
}
