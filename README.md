# golangCraigslist

This app will query Craigslist for computers in East Texas using of your choosing `http://127.0.0.1:55020/computers?query=macbook` 
and return results within 100 miles from my location in Whitehouse.

#### To run locally

`go run app.go`

#### Push to Docker Hub
```
APP_TAG=josephglaspie/golangcraigslist:v0.0.1
docker build . -t $APP_TAG
docker push $APP_TAG
```

#### To deploy to minikube
```
kubectl apply -f web-deployment.yaml
minikube service web 
```
