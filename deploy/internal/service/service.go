package service

import (
	"log"

	"deploy/internal/model"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Service struct {
	k8sClient *kubernetes.Clientset
	template  model.Template
}

func New(kubeConfig string, t model.Template) *Service {
	cfg, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		log.Fatalln(err)
	}

	s := &Service{
		k8sClient: kubernetes.NewForConfigOrDie(cfg),
		template:  t,
	}
	return s
}
