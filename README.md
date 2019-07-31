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

### If Mac as work station

If you are using a mac, you can use the following commands to get the prerequisites in place quickly.

* Install home brew: https://brew.sh `/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"`
* Install terraform via brew `brew install terraform` or upgrade if already installed `brew upgrade terraform`
* Install git via brew `brew install git` or upgrade if already installed `brew upgrade git`
* Install gcloud via brew `brew cask install google-cloud-sdk` or upgrade if already installed `brew cask upgrade google-cloud-sdk`

### If setting up gcloud for the first time

If you already have setup GCP via gcloud, you can skip this section

* `gcloud init`
* Answer Y
* This will launch a browser and finish up the setup process
* Back to the cli, select "Create New Project", and call it "automate-all-the-things"

## Setup

* `git clone https://github.com/Artemmkin/terraform-kubernetes.git`
* Enable k8 in gcp - https://console.developers.google.com/apis/library/container.googleapis.com?project=automate-all-the-things&pli=1
* Login to GCP and go to IAM > Service Accounts and create a terraform user > Owner
* Click Create key at the end of this step and move the it to terraform/account.json file
* `cd terraform`
* `terraform init`
* `terraform plan`
* `terraform apply` - this took 6 and half minutes
* TODO explain how to setup golang
* `cd ../code/`
* `go build`
* TODO docker build
* TODO k8 deploy


## Tear Down

## Resources used

* http://artemstar.com/2018/01/15/cicd-with-kubernetes-and-gitlab/
* https://www.terraform.io/docs/providers/google/r/container_cluster.html
* https://cloud.google.com/iam/docs/creating-managing-service-account-keys#creating_service_account_keys
* https://tutorialedge.net/golang/creating-simple-web-server-with-golang/


## Common problems

```
Error: googleapi: Error 403: Kubernetes Engine API has not been used in project 184897089599 before or it is disabled. Enable it by visiting https://console.developers.google.com/apis/api/container.googleapis.com/overview?project=184897089599 then retry. If you enabled this API recently, wait a few minutes for the action to propagate to our systems and retry., accessNotConfigured
```
* Enable billing for the account, and wait a few minutes for it to become active and make the terraform call again

* Ensure you made a service account and not an IAM account for terraform

## TODO

* Write idempotent bash start script to wrap all this and skip gcp setup of terraform if user has an existing k8 setup
* Variablize out terraform
* Find lower level for the service account then owner to perform terraform actions
