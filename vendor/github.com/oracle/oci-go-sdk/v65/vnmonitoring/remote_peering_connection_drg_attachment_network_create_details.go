// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RemotePeeringConnectionDrgAttachmentNetworkCreateDetails The representation of RemotePeeringConnectionDrgAttachmentNetworkCreateDetails
type RemotePeeringConnectionDrgAttachmentNetworkCreateDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of compartment that contains the remote peering connection.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The BGP ASN to use for the remote peering connection's route target.
	RegionalOciAsn *string `mandatory:"true" json:"regionalOciAsn"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network attached to the DRG.
	Id *string `mandatory:"false" json:"id"`

	// Whether the RPC attachment is for a GFC DRG, indicating the mpls label should be
	// allocated from the VC label range.
	// Example: `true`
	IsGlobalFastConnect *bool `mandatory:"false" json:"isGlobalFastConnect"`
}

//GetId returns Id
func (m RemotePeeringConnectionDrgAttachmentNetworkCreateDetails) GetId() *string {
	return m.Id
}

func (m RemotePeeringConnectionDrgAttachmentNetworkCreateDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RemotePeeringConnectionDrgAttachmentNetworkCreateDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m RemotePeeringConnectionDrgAttachmentNetworkCreateDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRemotePeeringConnectionDrgAttachmentNetworkCreateDetails RemotePeeringConnectionDrgAttachmentNetworkCreateDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeRemotePeeringConnectionDrgAttachmentNetworkCreateDetails
	}{
		"REMOTE_PEERING_CONNECTION",
		(MarshalTypeRemotePeeringConnectionDrgAttachmentNetworkCreateDetails)(m),
	}

	return json.Marshal(&s)
}
