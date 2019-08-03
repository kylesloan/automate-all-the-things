// +build mage

package main

import (
        "io/ioutil"
	"github.com/magefile/mage/sh" 
)

// Default target to run when none is specified
// If not set, running mage will list available targets

// Setup the GCP instance, deploy the k8 configs, and run the tests. Requires that you have setup the GCP account, enabled k8 api, and made the service account.
func Setup() {
  sh.RunV("ansible-playbook", "ansible/setup-playbook.yml")
}

// Destroy artifacts made and the GCK kube instance
func Destroy() {
  sh.RunV("ansible-playbook", "ansible/destroy-playbook.yml")
}

// Run tests, this requires setup to be run already which also runs the test at the end
func Test() {
    file, _ := ioutil.ReadFile("goss/container_name.txt")
    s := string(file)
    sh.RunV("docker", "run", s)
}
