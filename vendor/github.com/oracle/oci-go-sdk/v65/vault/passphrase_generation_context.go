// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vault Secret Management API
//
// Use the Secret Management API to manage secrets and secret versions. For more information, see Managing Secrets (https://docs.cloud.oracle.com/Content/KeyManagement/Tasks/managingsecrets.htm).
//

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PassphraseGenerationContext Generates Passphrase type of secrets.
type PassphraseGenerationContext struct {

	// Name of the predefined secret generation template.
	GenerationTemplate SecretGenerationContextGenerationTemplateEnum `mandatory:"true" json:"generationTemplate"`
}

//GetGenerationTemplate returns GenerationTemplate
func (m PassphraseGenerationContext) GetGenerationTemplate() SecretGenerationContextGenerationTemplateEnum {
	return m.GenerationTemplate
}

func (m PassphraseGenerationContext) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PassphraseGenerationContext) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSecretGenerationContextGenerationTemplateEnum(string(m.GenerationTemplate)); !ok && m.GenerationTemplate != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GenerationTemplate: %s. Supported values are: %s.", m.GenerationTemplate, strings.Join(GetSecretGenerationContextGenerationTemplateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PassphraseGenerationContext) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePassphraseGenerationContext PassphraseGenerationContext
	s := struct {
		DiscriminatorParam string `json:"generationType"`
		MarshalTypePassphraseGenerationContext
	}{
		"PASSPHRASE",
		(MarshalTypePassphraseGenerationContext)(m),
	}

	return json.Marshal(&s)
}
