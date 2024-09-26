package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessAudience struct {
	// The ID of the application.
	ApplicationId nullable.Type[string] `json:"applicationId,omitempty"`

	// Indicates the reasons this audience was included for a sign-in request.
	AudienceReasons *ConditionalAccessAudienceReason `json:"audienceReasons,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}