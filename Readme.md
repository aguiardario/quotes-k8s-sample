
### Construir la imagen Docker a utilizar

```
cd source
go build
docker build -t quotes:1.0.1 .
```


# K8s

### Creamos un Namespace

```
kubectl create namespace quotes-app-ns
```


### Deployment y Servicio de Mongo

```
kubectl create -f ./k8s/mongo.yml --namespace=quotes-app-ns
kubectl apply -f ./k8s/mongo.yml --namespace=quotes-app-ns

### Obtener los deployments
kubectl get deployment --namespace=quotes-app-ns

### Obtener los pods
kubectl get pods --namespace=quotes-app-ns

### Obtener info del deployment
kubectl describe deployment/quotes-app-deployment --namespace=quotes-app-ns

```

### Opción 1: Port-Forward

```
kubectl port-forward <NOMBRE-POD> TARGET-PORT:CONTAINER-PORT
```

por ejemplo:

```
kubectl port-forward quotes-app-deployment-f9c45fc9b-mz2tf 8080:3000 --namespace=quotes-app-ns
```

luego podemos ejecutar

```
curl localhost:8080
```


### Opción 2: Bastión

```
kubectl run -i --tty --rm debug --image=curlimages/curl --restart=Never -- sh --namespace=quotes-app-ns
```

ya metidos dentro del pod ejecutamos:

```
curl quotes-app-service/
```

donde curl **quotes-app-service** es el nombre del servicio definido en el archivo **service.yml**:

```
apiVersion: v1
kind: Service
metadata:
  name: quotes-app-service
  labels:
    app: quotes-app
    tier: backend
    environment: dev
[...]
```



### Opcion 3: NodePort

Una opción es realizar los cambios en el archivo service.yml y aplicar los cambios sobre el servicio.

Otra opción es crear un servicio paralelo así (aplicado al deployment):

```
kubectl expose deploy quotes-app-deployment --name=svc-test --type=NodePort --port=3000 --target-port=3000 --namespace=quotes-app-ns
```

También podríamos exponer directamente el pod:

```
kubectl expose pod quotes-app-deployment-f9c45fc9b-mz2tf --name=svc-test  --type=NodePort --port=3000 --target-port=3000 --namespace=quotes-app-ns
```

Luego obtenemos la ip del cluster:

```
minikube service quotes-app-service --url -n --namespace=quotes-app-ns
```


### Opción Ingress
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





