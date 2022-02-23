/*
 * CD4PE API V1
 *
 * API for CD4PE V1
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package cd4pe

// PuppetEnterpriseCredentials struct for PuppetEnterpriseCredentials
type PuppetEnterpriseCredentials struct {
	Certificate                         string `json:"certificate,omitempty"`
	CodeManagerEndpoint                 string `json:"codeManagerEndpoint,omitempty"`
	DefaultBranchOverrideSupported      bool   `json:"defaultBranchOverrideSupported,omitempty"`
	Domain                              string `json:"domain,omitempty"`
	Name                                string `json:"name,omitempty"`
	NodeClassifierEndpoint              string `json:"nodeClassifierEndpoint,omitempty"`
	OrchestratorEndpoint                string `json:"orchestratorEndpoint,omitempty"`
	ProtectedEnvironmentsCount          int32  `json:"protectedEnvironmentsCount,omitempty"`
	PuppetDbEndpoint                    string `json:"puppetDbEndpoint,omitempty"`
	PuppetServerCertificate             string `json:"puppetServerCertificate,omitempty"`
	PuppetServerEndpoint                string `json:"puppetServerEndpoint,omitempty"`
	PuppetServerPrivateKey              string `json:"puppetServerPrivateKey,omitempty"`
	Token                               string `json:"token,omitempty"`
	TokenExpiration                     int64  `json:"tokenExpiration,omitempty"`
	UseLegacyCompileEndpoint            bool   `json:"useLegacyCompileEndpoint,omitempty"`
	DefaultBranchOverrideSupportedAsStr string `json:"defaultBranchOverrideSupported_asStr,omitempty"`
	ProtectedEnvironmentsCountAsStr     string `json:"protectedEnvironmentsCount_asStr,omitempty"`
	TokenExpirationAsStr                string `json:"tokenExpiration_asStr,omitempty"`
	UseLegacyCompileEndpointAsStr       string `json:"useLegacyCompileEndpoint_asStr,omitempty"`
}
