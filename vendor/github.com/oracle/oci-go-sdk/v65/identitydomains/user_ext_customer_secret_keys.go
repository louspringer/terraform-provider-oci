// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Domains API
//
// Use the Identity Domains API to manage resources within an identity domain, for example, users, dynamic resource groups, groups, and identity providers. For information about managing resources within identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm). This REST API is SCIM compliant.
// Use the table of contents and search tool to explore the Identity Domains API.
//

package identitydomains

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UserExtCustomerSecretKeys A list of customer secret keys corresponding to user.
// **Added In:** 2102181953
// **SCIM++ Properties:**
//  - idcsCompositeKey: [value]
//  - idcsSearchable: true
//  - multiValued: true
//  - mutability: readOnly
//  - required: false
//  - returned: request
//  - type: complex
//  - uniqueness: none
type UserExtCustomerSecretKeys struct {

	// The user's customer secret key identifier.
	// **Added In:** 2102181953
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: always
	//  - type: string
	//  - uniqueness: none
	Value *string `mandatory:"false" json:"value"`

	// The user's customer secret key OCID.
	// **Added In:** 2102181953
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: always
	//  - type: string
	//  - uniqueness: none
	Ocid *string `mandatory:"false" json:"ocid"`

	// The URI of the corresponding CustomerSecretKey resource to which the user belongs.
	// **Added In:** 2102181953
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: reference
	//  - uniqueness: none
	Ref *string `mandatory:"false" json:"$ref"`
}

func (m UserExtCustomerSecretKeys) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UserExtCustomerSecretKeys) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
