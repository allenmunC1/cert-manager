// Package face implements the Azure ARM Face service API version 1.0.
//
// An API for face detection, verification, and identification.
//
// Deprecated: Please instead use github.com/Azure/azure-sdk-for-go/services
package face

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
)

const ()

// ManagementClient is the base client for Face.
type ManagementClient struct {
	autorest.Client
	SubscriptionKey string
	AzureRegion     AzureRegions
}

// New creates an instance of the ManagementClient client.
func New(subscriptionKey string, azureRegion AzureRegions) ManagementClient {
	return NewWithoutDefaults(subscriptionKey, azureRegion)
}

// NewWithoutDefaults creates an instance of the ManagementClient client.
func NewWithoutDefaults(subscriptionKey string, azureRegion AzureRegions) ManagementClient {
	return ManagementClient{
		Client:          autorest.NewClientWithUserAgent(UserAgent()),
		SubscriptionKey: subscriptionKey,
		AzureRegion:     azureRegion,
	}
}