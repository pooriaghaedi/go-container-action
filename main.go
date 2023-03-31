package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"gopkg.in/yaml.v3"
)

type myData struct {
        Backend_type          string  `yaml:"backend_type"`
        Kubernetesclient      string
        Openai_key 			  string  `yaml:"openai_key"`
}

func marshalYaml(filename string) ([]byte, error) {
	
	out, err := yaml.Marshal(&myData{Backend_type: os.Getenv("BACKENDTYPE"),Kubernetesclient: "{}",Openai_key: os.Getenv("OAIKEY")})
    if err != nil {
        return nil, fmt.Errorf("in file %q: %w", filename, err)
    }
	fmt.Println(string(out[:]))
    return out, err
}

func main() {

	filename := "/root/.k8sgpt.yaml"
	_,err := os.Create(filename)
	if err != nil {
    	 fmt.Println( err)
    }
	out, err := marshalYaml(filename)
	if err != nil {
        log.Fatal(err)
    }
	err = os.WriteFile(filename, out, 0644)
    if err != nil {
        log.Fatal(err)
    }

	k8sfilename := "/root/.kube/config"
	err = os.WriteFile(k8sfilename, []byte(os.Getenv("K8SCONFIG")), 0644)
    if err != nil {
        log.Fatal(err)
    }

	cmd := exec.Command("./k8sgpt" , "analyze" , "--explain",  "--namespace=default" , "--filter=Pod")
    if err := cmd.Run(); err != nil{
       fmt.Println(err)
    }
}

