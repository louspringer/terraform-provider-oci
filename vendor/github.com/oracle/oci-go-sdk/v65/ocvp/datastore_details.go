// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use the Oracle Cloud VMware API to create SDDCs and manage ESXi hosts and software.
// For more information, see Oracle Cloud VMware Solution (https://docs.cloud.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm).
//

package ocvp

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatastoreDetails Datastore details for a getting an Sddc.
type DatastoreDetails struct {

	// A list of OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)s of Block Storage Volumes.
	BlockVolumeIds []string `mandatory:"true" json:"blockVolumeIds"`

	// Type of the datastore.
	DatastoreType DatastoreTypesEnum `mandatory:"true" json:"datastoreType"`

	// Size of the Block Storage Volume in GB.
	Capacity *float64 `mandatory:"true" json:"capacity"`
}

func (m DatastoreDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatastoreDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatastoreTypesEnum(string(m.DatastoreType)); !ok && m.DatastoreType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatastoreType: %s. Supported values are: %s.", m.DatastoreType, strings.Join(GetDatastoreTypesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
