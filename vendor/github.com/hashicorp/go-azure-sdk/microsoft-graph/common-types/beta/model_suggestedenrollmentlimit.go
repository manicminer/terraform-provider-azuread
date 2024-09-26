package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SuggestedEnrollmentLimit struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The suggested enrollment limit within a day
	SuggestedDailyLimit *int64 `json:"suggestedDailyLimit,omitempty"`
}