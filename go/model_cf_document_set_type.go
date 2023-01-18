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

// CfDocumentSetType - This is the container for a collection of CFDocuments. There must be at least one CFDocument. 
type CfDocumentSetType struct {

	CFDocuments []CfDocumentType `json:"CFDocuments"`
}
