package model

type Data struct {
	ProjectName   string `json:"project_name"`
	ServiceName   string `json:"service_name"`
	GitVersionTag string `json:"git_version_tag""`
	PodCount      int    `json:"pod_count"`
}
