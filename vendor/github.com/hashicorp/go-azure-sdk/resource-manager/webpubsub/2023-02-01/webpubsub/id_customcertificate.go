package webpubsub

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = CustomCertificateId{}

// CustomCertificateId is a struct representing the Resource ID for a Custom Certificate
type CustomCertificateId struct {
	SubscriptionId        string
	ResourceGroupName     string
	WebPubSubName         string
	CustomCertificateName string
}

// NewCustomCertificateID returns a new CustomCertificateId struct
func NewCustomCertificateID(subscriptionId string, resourceGroupName string, webPubSubName string, customCertificateName string) CustomCertificateId {
	return CustomCertificateId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		WebPubSubName:         webPubSubName,
		CustomCertificateName: customCertificateName,
	}
}

// ParseCustomCertificateID parses 'input' into a CustomCertificateId
func ParseCustomCertificateID(input string) (*CustomCertificateId, error) {
	parser := resourceids.NewParserFromResourceIdType(CustomCertificateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CustomCertificateId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WebPubSubName, ok = parsed.Parsed["webPubSubName"]; !ok {
		return nil, fmt.Errorf("the segment 'webPubSubName' was not found in the resource id %q", input)
	}

	if id.CustomCertificateName, ok = parsed.Parsed["customCertificateName"]; !ok {
		return nil, fmt.Errorf("the segment 'customCertificateName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseCustomCertificateIDInsensitively parses 'input' case-insensitively into a CustomCertificateId
// note: this method should only be used for API response data and not user input
func ParseCustomCertificateIDInsensitively(input string) (*CustomCertificateId, error) {
	parser := resourceids.NewParserFromResourceIdType(CustomCertificateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CustomCertificateId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WebPubSubName, ok = parsed.Parsed["webPubSubName"]; !ok {
		return nil, fmt.Errorf("the segment 'webPubSubName' was not found in the resource id %q", input)
	}

	if id.CustomCertificateName, ok = parsed.Parsed["customCertificateName"]; !ok {
		return nil, fmt.Errorf("the segment 'customCertificateName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateCustomCertificateID checks that 'input' can be parsed as a Custom Certificate ID
func ValidateCustomCertificateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseCustomCertificateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Custom Certificate ID
func (id CustomCertificateId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.SignalRService/webPubSub/%s/customCertificates/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WebPubSubName, id.CustomCertificateName)
}

// Segments returns a slice of Resource ID Segments which comprise this Custom Certificate ID
func (id CustomCertificateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSignalRService", "Microsoft.SignalRService", "Microsoft.SignalRService"),
		resourceids.StaticSegment("staticWebPubSub", "webPubSub", "webPubSub"),
		resourceids.UserSpecifiedSegment("webPubSubName", "webPubSubValue"),
		resourceids.StaticSegment("staticCustomCertificates", "customCertificates", "customCertificates"),
		resourceids.UserSpecifiedSegment("customCertificateName", "customCertificateValue"),
	}
}

// String returns a human-readable description of this Custom Certificate ID
func (id CustomCertificateId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Web Pub Sub Name: %q", id.WebPubSubName),
		fmt.Sprintf("Custom Certificate Name: %q", id.CustomCertificateName),
	}
	return fmt.Sprintf("Custom Certificate (%s)", strings.Join(components, "\n"))
}
