/*
 * Competencies and Academic Standards Exchange (CASE) Service OpenAPI (YAML) Definition
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Contact: lmattson@imsglobal.org
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// CfConceptType - The container for the definition of a concept which is addressed by the competency framework. 
type CfConceptType struct {

	// The data-type for establishing a Globally Unique Identifier (GUID). The form of the GUID is a Universally Unique Identifier (UUID) of 16 hexadecimal characters (lower case) in the format 8-4-4-4-12. All permitted versions (1-5) and variants (1-2) are supported. 
	Identifier string `json:"identifier"`

	// Model Primitive Datatype = AnyURI
	Uri string `json:"uri"`

	// Model Primitive Datatype = NormalizedString
	Title string `json:"title"`

	// Model Primitive Datatype = NormalizedString
	Keywords string `json:"keywords,omitempty"`

	// Model Primitive Datatype = NormalizedString
	HierarchyCode string `json:"hierarchyCode"`

	// Model Primitive Datatype = String
	Description string `json:"description,omitempty"`

	// Model Primitive Datatype = DateTime
	LastChangeDateTime time.Time `json:"lastChangeDateTime"`
}
