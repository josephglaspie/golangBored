# golangCraigslist

This app will find something fun to do if you're bored. Just enter the number of people you are currently with
`http://127.0.0.1:55020/bored?query=2` 
and return results with something fun to do.

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
In my deployment yaml you'll find I'm using an image I pushed to Dockerhub. If you're using minikube and 
want to use local images only check out this [medium article](https://medium.com/swlh/how-to-run-locally-built-docker-images-in-kubernetes-b28fbc32cc1d) 

