# automate-all-the-things
Learn some golang and GCP (google cloud compute). https://github.com/kylesloan/automate-all-the-things


## Problem Statement

Build an application in the programming language of your choice that exposes a REST endpoint that returns the following JSON payload with the current timestamp and a static message:
```
{
  "message": "Automate all the things!",
  "timestamp": 1529729125
}
```
The application must be deployed on a Kubernetes cluster running in a public cloud provider of your choice. The provisioning of the cluster as well as the deployment of the application must be done through code.

## Prerequisites and the specific version used while building out this app

* git (2.22.0_1)
* terraform (0.12.5)
* GCP account (new accounts get $300 free credit at this time) - https://console.cloud.google.com/
* gcloud (google-cloud-sdk) cli tool
* golang (1.12.7)
* kubernetes-cli 1.15.1


### If Mac as work station

If you are using a mac, you can use the following commands to get the prerequisites in place quickly.

* Install home brew: https://brew.sh `/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"`
* Install git via brew `brew install git` or upgrade if already installed `brew upgrade git`
* Install terraform via brew `brew install terraform` or upgrade if already installed `brew upgrade terraform`
* Install gcloud via brew `brew cask install google-cloud-sdk` or upgrade if already installed `brew cask upgrade google-cloud-sdk`
* Install golang via brew `brew install golang` or upgrade if already installed `brew upgrade golang`
* Install kubectl via brew `brew install kubernetes-cli` or upgrade if already installed `brew upgrade kubernetes-cli`

### If setting up gcloud for the first time

If you already have setup GCP via gcloud, you can skip this section

* `gcloud init`
* Answer Y
* This will launch a browser and finish up the setup process
* Back to the cli, select "Create New Project", and call it "automate-all-the-things"

## Setup

* `git clone https://github.com/kylesloan/automate-all-the-things.git`
* Enable k8 in gcp - https://console.developers.google.com/apis/library/container.googleapis.com?
* Login to GCP and go to IAM > Service Accounts and create a terraform user with Owner permissions
* Click Create key at the end of this step and move the it to terraform/account.json file
* `cd $PATH_TO_CHECKOUT/terraform/`
* `terraform init`
* `terraform apply` - this took 6 and half minutes
* `cd $PATH_TO_CHECKOUT/code/`
* `env GOOS=linux GOARCH=amd64 go build -o code.bin main.go` - need the linux/amd64 to run properly in gcp
* `docker login` if not already logged into docker hub
* `docker build -t kylesloan/automate-all-the-things:latest .`
* `docker tag automate-all-the-things:latest kylesloan/automate-all-the-things:latest` - TODO determine the user repo that they pushed to
* `docker push kylesloan/automate-all-the-things:latest`
* `cd $PATH_TO_CHECKOUT/k8/`
* `kubectl apply -f deploy.yml` - TODO this needs the proper username from the hub
* `kubectl apply -f service.yml`
* `kubectl apply -f ingress.yml`
* `kubectl get ingress -o wide` - you can run this under watch and wait until you see the IP address appear, this took several minutes even after kubectl said it had assigned the ip to no longer get GCP error page
* `curl -iL $IP_FROM_PREVIOUS_COMMAND`



## Tear Down

* `cd ~/$CHECKOUT_PATH/terraform`
* `terraform destroy` - this took about 7 minutes

## Resources used

* http://artemstar.com/2018/01/15/cicd-with-kubernetes-and-gitlab/
* https://www.terraform.io/docs/providers/google/r/container_cluster.html
* https://cloud.google.com/iam/docs/creating-managing-service-account-keys#creating_service_account_keys
* https://tutorialedge.net/golang/creating-simple-web-server-with-golang/
* https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04#step-4-—-building-executables-for-different-architectures
* https://cloud.google.com/kubernetes-engine/docs/tutorials/http-balancer


## Common problems

```
Error: googleapi: Error 403: Kubernetes Engine API has not been used in project 184897089599 before or it is disabled. Enable it by visiting https://console.developers.google.com/apis/api/container.googleapis.com/overview?project=184897089599 then retry. If you enabled this API recently, wait a few minutes for the action to propagate to our systems and retry., accessNotConfigured
```
* Enable billing for the account, and wait a few minutes for it to become active and make the terraform call again

* Ensure you made a service account and not an IAM account for terraform

```
Failed to import the required Python library (openshift) on imac-6.local's Python /usr/local/Cellar/ansible/2.8.3/libexec/bin/python3.7. Please read module documentation and install in the appropriate location
```
* https://docs.ansible.com/ansible/latest/modules/k8s_module.html#requirements

## gcloud debugging commands

* `gcloud container clusters list`
* `gcloud container clusters get-credentials my-gke-cluster --region us-central1`
* `kubectl config current-context`
* `kubectl run -i --tty ubuntu --image=ubuntu:16.04 --restart=Never -- bash -il` - `apt update && apt install -y curl iputils-ping host`

## TODO

* Write idempotent bash start script to wrap all this and skip gcp setup of terraform if user has an existing k8 setup
* Variablize out terraform
* Find lower level for the service account then owner to perform terraform actions
* monitoring/metrics/graphing
