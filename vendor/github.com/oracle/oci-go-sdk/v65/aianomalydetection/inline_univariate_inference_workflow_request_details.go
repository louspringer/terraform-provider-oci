// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Anomaly Detection API
//
// OCI AI Service solutions can help Enterprise customers integrate AI into their products immediately by using our proven,
// pre-trained/custom models or containers, and without a need to set up in house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI/ML operations, shortening the time to market.
//

package aianomalydetection

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InlineUnivariateInferenceWorkflowRequestDetails The inline list signal of signal data.
type InlineUnivariateInferenceWorkflowRequestDetails struct {

	// Choose whether you would like the service to return all data points or just anomalies.
	AreAllDataPointsRequired *bool `mandatory:"false" json:"areAllDataPointsRequired"`

	TrainingRequestDetails *UnivariateModelTrainingRequestDetails `mandatory:"false" json:"trainingRequestDetails"`

	// Tune between precision and recall.
	Sensitivity *float32 `mandatory:"false" json:"sensitivity"`

	// The list of all signals and their values.
	SignalData []UnivariateInlineSignalRequestData `mandatory:"false" json:"signalData"`
}

//GetAreAllDataPointsRequired returns AreAllDataPointsRequired
func (m InlineUnivariateInferenceWorkflowRequestDetails) GetAreAllDataPointsRequired() *bool {
	return m.AreAllDataPointsRequired
}

//GetTrainingRequestDetails returns TrainingRequestDetails
func (m InlineUnivariateInferenceWorkflowRequestDetails) GetTrainingRequestDetails() *UnivariateModelTrainingRequestDetails {
	return m.TrainingRequestDetails
}

//GetSensitivity returns Sensitivity
func (m InlineUnivariateInferenceWorkflowRequestDetails) GetSensitivity() *float32 {
	return m.Sensitivity
}

func (m InlineUnivariateInferenceWorkflowRequestDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InlineUnivariateInferenceWorkflowRequestDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m InlineUnivariateInferenceWorkflowRequestDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInlineUnivariateInferenceWorkflowRequestDetails InlineUnivariateInferenceWorkflowRequestDetails
	s := struct {
		DiscriminatorParam string `json:"requestType"`
		MarshalTypeInlineUnivariateInferenceWorkflowRequestDetails
	}{
		"INLINE",
		(MarshalTypeInlineUnivariateInferenceWorkflowRequestDetails)(m),
	}

	return json.Marshal(&s)
}
