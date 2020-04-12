
# Docker

```
cd source
go build
docker build -t quotes:1.0.1

docker run --rm -p 3000:3000 -e USER_DB=sapo -e PWD_DB=123456 quotes:1.0.1
```



# K8s
### Crear Namespace

```
kubectl create namespace quotes-app-ns
```


### Deployment y Servicio de Mongo

```
kubectl create -f ./k8s/mongo.yml --namespace=quotes-app-ns
kubectl apply -f ./k8s/mongo.yml --namespace=quotes-app-ns

# Obtener los deployments
kubectl get deployment --namespace=quotes-app-ns

# Obtener los pods
kubectl get pods --namespace=quotes-app-ns

# Obtener info del deployment
kubectl describe deployment/quotes-app-deployment --namespace=quotes-app-ns

```




### Deployment

```
# Realizar el deployment

kubectl create -f ./k8s/deployment.yml --namespace=quotes-app-ns
kubectl apply -f ./k8s/deployment.yml --namespace=quotes-app-ns

# Obtener los deployments
kubectl get deployment --namespace=quotes-app-ns

# Obtener los pods
kubectl get pods --namespace=quotes-app-ns

# Obtener info del deployment
kubectl describe deployment/quotes-app-deployment --namespace=quotes-app-ns


kubectl delete deployment/quotes-app-deployment --namespace=quotes-app-ns

```

### Servicio

```
kubectl create -f ./k8s/service.yml --namespace=quotes-app-ns
kubectl apply -f ./k8s/service.yml --namespace=quotes-app-ns

# Obtener los servicios
kubectl get service --namespace=quotes-app-ns

# Obtener la url (minikube)
minikube service quotes-app-service --url --namespace=quotes-app-ns


kubectl delete service/quotes-app-service --namespace=quotes-app-ns
```


### Ingress
```
kubectl create -f ./k8s/ingress.yml --namespace=quotes-app-ns
kubectl apply -f ./k8s/ingress.yml --namespace=quotes-app-ns


kubectl get ingress --namespace=quotes-app-ns

kubectl describe ing quotes-app-ingress --namespace=quotes-app-ns

NAME                 HOSTS             ADDRESS        PORTS   AGE
quotes-app-ingress   quotes-app.info   192.168.64.3   80      4m55s


#Si estamos usando minikube:

minikube ip

Agregar al archivo /etc/hosts

192.168.64.3 quotes-app.info


# Delete
kubectl delete ingress/quotes-app-ingress --namespace=quotes-app-ns
```

### Logs
```
kubectl -n quotes-app-ns logs -f quotes-app-deployment-57f4bdc5f4-6kc7j
```





