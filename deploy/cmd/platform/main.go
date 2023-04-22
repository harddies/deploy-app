package main

import (
	"flag"
	"log"
	"path/filepath"
	"text/template"

	"deploy/internal/model"
	"deploy/internal/service"

	"k8s.io/client-go/util/homedir"
)

func main() {
	var d, tplDir, kubeConfig string
	flag.StringVar(&d, "data", "", "please specify deployed arg for deploying app, it's json")
	flag.StringVar(&tplDir, "tpl_dir", "", "please specify yaml tpl dir")
	flag.StringVar(&kubeConfig, "kubeconfig", "", "(optional) absolute path to the kubeconfig file")
	flag.Parse()

	if kubeConfig == "" {
		home := homedir.HomeDir()
		if home == "" {
			log.Fatalln("kubeconnfig is empty and home dir is empty")
		}
		kubeConfig = filepath.Join(home, ".kube", "config")
	}
	if tplDir == "" {
		tplDir = "../../tpl"
	}
	if d == "" {
		log.Fatalln("data is empty")
	}

	t, err := getTpl(tplDir)
	if err != nil {
		log.Fatalln(err)
	}

	s := service.New(kubeConfig, t)
	s.ApplyDeployment(d)
}

func getTpl(tplDir string) (t model.Template, err error) {
	t = model.Template{}
	t.DeployApp, err = template.ParseFiles(filepath.Join(tplDir, "deploy-app.yaml"))
	if err != nil {
		return
	}
	return
}
