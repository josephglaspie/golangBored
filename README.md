# golangCraigslist

This app will find something fun to do if you're bored. Just enter the number of people you are currently with
`http://127.0.0.1:55020/bored?query=2` 
and return results with something fun to do.

### Prerequisites
```
#install Go
brew install go

#install Docker
brew cask install docker

#install minikube
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-darwin-amd64
sudo install minikube-darwin-amd64 /usr/local/bin/minikube

#install Kubectl
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/darwin/amd64/kubectl"

```
#### To run locally

`go run app.go`

#### Push to Docker Hub
```
APP_TAG=josephglaspie/bored:v0.0.5
docker build . -t $APP_TAG
docker push $APP_TAG
```

#### To deploy to minikube
```
kubectl apply -f web-deployment.yaml
# The following will return the port
minikube service web 
```
#### Make available online with ngrok
```
ngrok http http://127.0.0.1:56952
```
Take the return from ngrok and add it to the twilio webhook

#### Use local images
In the deployment yaml you'll find  an image pushed to Dockerhub. If you're using minikube and 
want to use local images only, check out this [medium article](https://medium.com/swlh/how-to-run-locally-built-docker-images-in-kubernetes-b28fbc32cc1d) 

#### EKS
- Used eks module: https://registry.terraform.io/modules/terraform-aws-modules/eks/aws/latest
- Create main.tf as seen in this repo and tfvars.tf with the following:

```
provider "aws" {
  region     = "us-east-2"
  access_key = ""
  secret_key = ""
}
```
- You can get your access_key and secret_key from the security credential part of the console 
https://console.aws.amazon.com/iam/home?region=us-east-2#/security_credentials
- Acess Keys > Create new key
```
terraform init
terraform plan
terraform apply #Go make some tea or coffee this will take around 10 minutes to build out
# Connect your kubectl to the cluster
aws eks --region us-east-2 update-kubeconfig --name test-cluster
# Test connection
k get pods
k apply -f web-deployment.yaml
terraform destroy
```
- Now we need to expose the application https://aws.amazon.com/premiumsupport/knowledge-center/eks-kubernetes-services-cluster/