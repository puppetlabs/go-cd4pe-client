/*
 * CD4PE API V1
 *
 * API for CD4PE V1
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package cd4pe
// ControlRepo struct for ControlRepo
type ControlRepo struct {
	Created int64 `json:"created,omitempty"`
	Description string `json:"description,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	Domain string `json:"domain,omitempty"`
	LastCommitId string `json:"lastCommitId,omitempty"`
	LastDeploy RollingDeploymentEvent `json:"lastDeploy,omitempty"`
	LastImpactAnalysis PeImpactAnalysisEvent `json:"lastImpactAnalysis,omitempty"`
	LastDeployment DeploymentAppEvent `json:"lastDeployment,omitempty"`
	LastJob VmJobInstanceEvent `json:"lastJob,omitempty"`
	Name string `json:"name,omitempty"`
	SrcRepoDisplayName string `json:"srcRepoDisplayName,omitempty"`
	SrcRepoDisplayOwner string `json:"srcRepoDisplayOwner,omitempty"`
	SrcRepoId string `json:"srcRepoId,omitempty"`
	SrcRepoName string `json:"srcRepoName,omitempty"`
	SrcRepoOwner string `json:"srcRepoOwner,omitempty"`
	SrcRepoProvider string `json:"srcRepoProvider,omitempty"`
	Repository Repository `json:"repository,omitempty"`
}
