package main

import (
	"errors"
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
	// fmt.Println(string(out[:]))
    return out, err
}

func main() {

	filename := os.Getenv("HOME")+"/.k8sgpt.yaml"
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


	path :=  os.Getenv("HOME")+"/.kube"
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
	k8sfilename := os.Getenv("HOME")+"/.kube/config"
	_,err = os.Create(os.Getenv("HOME")+"/.kube/")
	err = os.WriteFile(k8sfilename, []byte(os.Getenv("KUBE_CONFIG")), 0644)
    if err != nil {
        log.Fatal(err)
    }

	cmd := exec.Command("/app/k8sgpt" , "analyze" , "--explain",  "--namespace=default" , "--filter=Pod", "--output=json")
	// output, _ := cmd.CombinedOutput()
	// fmt.Println(fmt.Sprintf(`myOutput=%s`, output))
	os.Setenv("GITHUB_OUTPUT", "myOutput=1"  )
    if err := cmd.Run(); err != nil{
       fmt.Println(err)
    }
}

