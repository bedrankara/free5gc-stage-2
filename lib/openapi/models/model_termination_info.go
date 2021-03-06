/*
 * Npcf_PolicyAuthorization Service API
 *
 * This is the Policy Authorization Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// indicates the cause for requesting the deletion of the Individual Application Session Context resource
type TerminationInfo struct {
	TermCause TerminationCause `json:"termCause" bson:"termCause"`
	ResUri    string           `json:"resUri" bson:"resUri"`
}
