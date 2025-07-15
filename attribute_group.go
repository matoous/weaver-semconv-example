// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated from semantic convention specification. DO NOT EDIT.

package semconv // import "go.opentelemetry.io/otel/semconv/"

import "go.opentelemetry.io/otel/attribute"

// Namespace: acme
const (
	// AcmeProductIDKey is the attribute Key conforming to the "acme.product.id"
	// semantic conventions. It represents the unique identifier of an Acme corp
	// product.
	//
	// Type: string
	// RequirementLevel: Required
	// Stability: Stable
	//
	// Examples: "my_product_id", "1234"
	AcmeProductIDKey = attribute.Key("acme.product.id")
)

// AcmeProductID returns an attribute KeyValue conforming to the
// "acme.product.id" semantic conventions. It represents the unique identifier of
// an Acme corp product.
func AcmeProductID(val string) attribute.KeyValue {
	return AcmeProductIDKey.String(val)
}

// Namespace: oauth2
const (
	// Oauth2ClientIDKey is the attribute Key conforming to the "oauth2.client.id"
	// semantic conventions. It represents the unique identifier of the OAuth 2.0
	// client.
	//
	// Type: string
	// RequirementLevel: Recommended
	// Stability: Stable
	//
	// Examples: "example_client_id"
	Oauth2ClientIDKey = attribute.Key("oauth2.client.id")
)

// Oauth2ClientID returns an attribute KeyValue conforming to the
// "oauth2.client.id" semantic conventions. It represents the unique identifier
// of the OAuth 2.0 client.
func Oauth2ClientID(val string) attribute.KeyValue {
	return Oauth2ClientIDKey.String(val)
}