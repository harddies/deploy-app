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
	applyCoreV1 "k8s.io/client-go/applyconfigurations/core/v1"
)

func (s *Service) ApplyDeployment(d string) {
	var err error

	data := new(model.Data)
	if err = json.Unmarshal([]byte(d), data); err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	buf := bytes.Buffer{}
	if err = s.template.DeployApp.Execute(&buf, data); err != nil {
		log.Fatal(err)
	}
	dac := new(applyAppsV1.DeploymentApplyConfiguration)
	if err = yaml.Unmarshal(buf.Bytes(), dac); err != nil {
		log.Fatalf("yaml.Unmarshal deployment %+v", err)
	}
	if _, err = s.k8sClient.AppsV1().Deployments(model.K8sNamespaceApp).
		Apply(ctx, dac, metaV1.ApplyOptions{FieldManager: model.K8sFieldManager}); err != nil {
		log.Fatalf("k8s apply deployment %+v", err)
	}

	buf = bytes.Buffer{}
	if err = s.template.ServiceApp.Execute(&buf, data); err != nil {
		log.Fatal(err)
	}
	svc := new(applyCoreV1.ServiceApplyConfiguration)
	if err = yaml.Unmarshal(buf.Bytes(), svc); err != nil {
		log.Fatalf("yaml.Unmarshal svc %+v", err)
	}
	if _, err = s.k8sClient.CoreV1().Services(model.K8sNamespaceApp).
		Apply(ctx, svc, metaV1.ApplyOptions{FieldManager: model.K8sFieldManager}); err != nil {
		log.Fatalf("k8s apply svc %+v", err)
	}

	log.Println("apply deployment successfully")
}
