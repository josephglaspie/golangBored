# golangBored

This app will find something fun to do if you're bored. Just enter the number of people you are currently with
```
curl --location --request POST 'http://k8s-bored-ingressb-a7cccb6576-1355983709.us-east-2.elb.amazonaws.com/bored' \
--header 'Content-Type: text/plain' \
--data-raw '5'
```

and return results with something fun to do.

### Prerequisites
```
#install Go
brew install go

#install Terraform
brew install terraform

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
kubectl apply -f deployments/mini-deployment.yaml
# The following will return the port
minikube service web 
```
#### Make available online with ngrok
```
ngrok http http://127.0.0.1:56952
```
Take the return from ngrok and add it to the twilio webhook

#### EKS
- Used eks module: https://registry.terraform.io/modules/terraform-aws-modules/eks/aws/latest
- Create main.tf as seen in this repo and tfvars.tf (.gitignore) with the following:

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
cd terraform
terraform init
terraform plan
terraform apply #Go make some tea or coffee this will take around 10 minutes to build out
cd ..

# Connect your kubectl to the cluster
aws configure #paste your keys into the correct places
aws eks --region us-east-2 update-kubeconfig --name mycluster
# Will return something like: Updated context arn:aws:eks:us-east-2:244172242562:cluster/mycluster in /Users/name/.kube/config

```
### Ingress
- [The Best Ingress Doc](https://docs.aws.amazon.com/eks/latest/userguide/alb-ingress.html)
- [AWS Ingress Github](https://github.com/kubernetes-sigs/aws-load-balancer-controller)
- [AWS App Load Balancer]( https://docs.aws.amazon.com/eks/latest/userguide/aws-load-balancer-controller.html)
- [How AWS Load Balancer controller works](https://kubernetes-sigs.github.io/aws-load-balancer-controller/v2.3/how-it-works/)

```
#Create IAM Policy
aws iam create-policy \
--policy-name AWSLoadBalancerControllerIAMPolicy \
--policy-document file://iam_policy.json

#Create service account
myaccount=244172242562
clustername=mycluster
eksctl create iamserviceaccount \
--cluster=$clustername \
--namespace=kube-system \
--name=aws-load-balancer-controller \
--attach-policy-arn=arn:aws:iam::$myaccount:policy/AWSLoadBalancerControllerIAMPolicy \
--override-existing-serviceaccounts \
--approve

eksctl utils associate-iam-oidc-provider --region=us-east-2 --cluster=mycluster --approve

# Install cert-manager to inject certificate configuration into the webhooks.
kubectl apply \
--validate=false \
-f https://github.com/jetstack/cert-manager/releases/download/v1.5.4/cert-manager.yaml

# Download and replace default cluster name with your own
curl -Lo deployments/loadbalancer.yaml https://github.com/kubernetes-sigs/aws-load-balancer-controller/releases/download/v2.3.0/v2_3_0_full.yaml
# Apply ALB with correct cluster name
kubectl apply -f deployments/loadbalancer.yaml

# Check on ALB deployment
kubectl get deployment -n kube-system aws-load-balancer-controller
kubectl logs -n kube-system   deployment.apps/aws-load-balancer-controller

```
1. Log into console and tage subnets with the following:
Key – kubernetes.io/role/elb
Value – 1
2. Also, download the sample app and change out the settings with your own AND add your subnets to the ingress annotation:
```
curl -Lo deployments/web-deployment.yaml https://raw.githubusercontent.com/kubernetes-sigs/aws-load-balancer-controller/v2.3.0/docs/examples/2048/2048_full.yaml
alb.ingress.kubernetes.io/subnets: subnet-0213058ab09c17cca, subnet-003c36f6f9c457619, subnet-00a17755b8471dbd5
```
```
# Test connection
k get pods -n kube-system
# Deploy app with service and ingress
k apply -f deployments/web-deployment.yaml
# Check ingress
kubectl get ingress/ingress-bored -n bored

#get external public ip
k get svc
NAME                       TYPE           CLUSTER-IP       EXTERNAL-IP                                                               PORT(S)          AGE
kubernetes                 ClusterIP      10.100.0.1       <none>                                                                    443/TCP          10m
web                        NodePort       10.100.130.96    <none>                                                                    8080:31317/TCP   10s
web-service-cluster-ip     ClusterIP      10.100.231.227   <none>                                                                    8080/TCP         10s
web-service-loadbalancer   LoadBalancer   10.100.121.221   a93868263dcf34d12b029f8c080c8951-1996062109.us-east-2.elb.amazonaws.com   8080:30061/TCP   10s
web-service-nodeport       NodePort       10.100.143.240   <none>                                                                    80:31549/TCP     10s

curl --location --request POST 'http://k8s-bored-ingressb-a7cccb6576-1355983709.us-east-2.elb.amazonaws.com/bored' \
--header 'Content-Type: text/plain' \
--data-raw '5'

terraform destroy
```
