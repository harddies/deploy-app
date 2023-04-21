package service

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"deploy/internal/model"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	applyAppsV1 "k8s.io/client-go/applyconfigurations/apps/v1"
)

func (s *Service) ApplyDeployment(tplDir string, d string) {
	var err error

	data := new(model.Data)
	if err = json.Unmarshal([]byte(d), data); err != nil {
		log.Fatal(err)
	}

	buf := bytes.Buffer{}
	if err = s.template.DeployApp.Execute(&buf, data); err != nil {
		log.Fatal(err)
	}
	b := buf.Bytes()

	dac := new(applyAppsV1.DeploymentApplyConfiguration)
	if err = yaml.Unmarshal(b, dac); err != nil {
		log.Fatalf("yaml.Unmarshal %+v", err)
	}
	if _, err = s.k8sClient.AppsV1().Deployments(model.K8sNamespaceApp).
		Apply(context.Background(), dac, metaV1.ApplyOptions{FieldManager: "kubectl-apply-controller"}); err != nil {
		log.Fatalf("k8s apply %+v", err)
	}
	log.Println("apply deployment successfully")
}
