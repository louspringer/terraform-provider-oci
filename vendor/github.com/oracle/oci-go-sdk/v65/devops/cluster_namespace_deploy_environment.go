// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ClusterNamespaceDeployEnvironment Specifies cluster namespace environment.
type ClusterNamespaceDeployEnvironment struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of a project.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID of a compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the Cluster Namespace to deploy to.
	ClusterNamespaceId *string `mandatory:"true" json:"clusterNamespaceId"`

	// Optional description about the deployment environment.
	Description *string `mandatory:"false" json:"description"`

	// Deployment environment display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Time the deployment environment was created. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time the deployment environment was updated. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	NetworkChannel NetworkChannel `mandatory:"false" json:"networkChannel"`

	// The current state of the deployment environment.
	LifecycleState DeployEnvironmentLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

//GetId returns Id
func (m ClusterNamespaceDeployEnvironment) GetId() *string {
	return m.Id
}

//GetDescription returns Description
func (m ClusterNamespaceDeployEnvironment) GetDescription() *string {
	return m.Description
}

//GetDisplayName returns DisplayName
func (m ClusterNamespaceDeployEnvironment) GetDisplayName() *string {
	return m.DisplayName
}

//GetProjectId returns ProjectId
func (m ClusterNamespaceDeployEnvironment) GetProjectId() *string {
	return m.ProjectId
}

//GetCompartmentId returns CompartmentId
func (m ClusterNamespaceDeployEnvironment) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetTimeCreated returns TimeCreated
func (m ClusterNamespaceDeployEnvironment) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m ClusterNamespaceDeployEnvironment) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m ClusterNamespaceDeployEnvironment) GetLifecycleState() DeployEnvironmentLifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecycleDetails returns LifecycleDetails
func (m ClusterNamespaceDeployEnvironment) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

//GetFreeformTags returns FreeformTags
func (m ClusterNamespaceDeployEnvironment) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m ClusterNamespaceDeployEnvironment) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m ClusterNamespaceDeployEnvironment) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m ClusterNamespaceDeployEnvironment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ClusterNamespaceDeployEnvironment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDeployEnvironmentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDeployEnvironmentLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ClusterNamespaceDeployEnvironment) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeClusterNamespaceDeployEnvironment ClusterNamespaceDeployEnvironment
	s := struct {
		DiscriminatorParam string `json:"deployEnvironmentType"`
		MarshalTypeClusterNamespaceDeployEnvironment
	}{
		"CLUSTER_NAMESPACE",
		(MarshalTypeClusterNamespaceDeployEnvironment)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ClusterNamespaceDeployEnvironment) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description        *string                             `json:"description"`
		DisplayName        *string                             `json:"displayName"`
		TimeCreated        *common.SDKTime                     `json:"timeCreated"`
		TimeUpdated        *common.SDKTime                     `json:"timeUpdated"`
		LifecycleState     DeployEnvironmentLifecycleStateEnum `json:"lifecycleState"`
		LifecycleDetails   *string                             `json:"lifecycleDetails"`
		FreeformTags       map[string]string                   `json:"freeformTags"`
		DefinedTags        map[string]map[string]interface{}   `json:"definedTags"`
		SystemTags         map[string]map[string]interface{}   `json:"systemTags"`
		NetworkChannel     networkchannel                      `json:"networkChannel"`
		Id                 *string                             `json:"id"`
		ProjectId          *string                             `json:"projectId"`
		CompartmentId      *string                             `json:"compartmentId"`
		ClusterNamespaceId *string                             `json:"clusterNamespaceId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.DisplayName = model.DisplayName

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	nn, e = model.NetworkChannel.UnmarshalPolymorphicJSON(model.NetworkChannel.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.NetworkChannel = nn.(NetworkChannel)
	} else {
		m.NetworkChannel = nil
	}

	m.Id = model.Id

	m.ProjectId = model.ProjectId

	m.CompartmentId = model.CompartmentId

	m.ClusterNamespaceId = model.ClusterNamespaceId

	return
}
