package firewallrules

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServerGroupsv2Id{}

// ServerGroupsv2Id is a struct representing the Resource ID for a Server Groupsv 2
type ServerGroupsv2Id struct {
	SubscriptionId     string
	ResourceGroupName  string
	ServerGroupsv2Name string
}

// NewServerGroupsv2ID returns a new ServerGroupsv2Id struct
func NewServerGroupsv2ID(subscriptionId string, resourceGroupName string, serverGroupsv2Name string) ServerGroupsv2Id {
	return ServerGroupsv2Id{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		ServerGroupsv2Name: serverGroupsv2Name,
	}
}

// ParseServerGroupsv2ID parses 'input' into a ServerGroupsv2Id
func ParseServerGroupsv2ID(input string) (*ServerGroupsv2Id, error) {
	parser := resourceids.NewParserFromResourceIdType(ServerGroupsv2Id{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServerGroupsv2Id{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServerGroupsv2Name, ok = parsed.Parsed["serverGroupsv2Name"]; !ok {
		return nil, fmt.Errorf("the segment 'serverGroupsv2Name' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseServerGroupsv2IDInsensitively parses 'input' case-insensitively into a ServerGroupsv2Id
// note: this method should only be used for API response data and not user input
func ParseServerGroupsv2IDInsensitively(input string) (*ServerGroupsv2Id, error) {
	parser := resourceids.NewParserFromResourceIdType(ServerGroupsv2Id{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServerGroupsv2Id{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServerGroupsv2Name, ok = parsed.Parsed["serverGroupsv2Name"]; !ok {
		return nil, fmt.Errorf("the segment 'serverGroupsv2Name' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateServerGroupsv2ID checks that 'input' can be parsed as a Server Groupsv 2 ID
func ValidateServerGroupsv2ID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServerGroupsv2ID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Server Groupsv 2 ID
func (id ServerGroupsv2Id) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServerGroupsv2Name)
}

// Segments returns a slice of Resource ID Segments which comprise this Server Groupsv 2 ID
func (id ServerGroupsv2Id) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDBforPostgreSQL", "Microsoft.DBforPostgreSQL", "Microsoft.DBforPostgreSQL"),
		resourceids.StaticSegment("staticServerGroupsv2", "serverGroupsv2", "serverGroupsv2"),
		resourceids.UserSpecifiedSegment("serverGroupsv2Name", "serverGroupsv2Value"),
	}
}

// String returns a human-readable description of this Server Groupsv 2 ID
func (id ServerGroupsv2Id) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Server Groupsv 2 Name: %q", id.ServerGroupsv2Name),
	}
	return fmt.Sprintf("Server Groupsv 2 (%s)", strings.Join(components, "\n"))
}
