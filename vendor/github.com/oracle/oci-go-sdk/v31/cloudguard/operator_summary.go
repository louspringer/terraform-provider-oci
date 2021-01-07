// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/v31/common"
)

// OperatorSummary Summary of Operator
type OperatorSummary struct {

	// name of the operand
	Name *string `mandatory:"true" json:"name"`

	// display name of the operand
	DisplayName *string `mandatory:"true" json:"displayName"`

	// data type of operand
	Datatype *string `mandatory:"true" json:"datatype"`

	// operand list type
	ManagedListtype *string `mandatory:"true" json:"managedListtype"`

	// Filter type can be config filter or condition filter
	FilterType ConditionFilterTypeEnum `mandatory:"true" json:"filterType"`

	// List of parameters
	Operators []ConditionOperator `mandatory:"true" json:"operators"`

	// configuration value type list for multilist data type
	MultiListTypes []string `mandatory:"false" json:"multiListTypes"`
}

func (m OperatorSummary) String() string {
	return common.PointerString(m)
}
