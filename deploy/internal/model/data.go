package model

type Data struct {
	ProjectName string `json:"project_name"`
	ServiceName string `json:"service_name"`
	PodCount    int    `json:"pod_count"`
}
