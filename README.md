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


## Resources used

* http://artemstar.com/2018/01/15/cicd-with-kubernetes-and-gitlab/
