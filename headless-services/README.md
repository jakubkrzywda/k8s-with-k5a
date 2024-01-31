# Headless Services

## With Selectors

Create a deployment consisting of two repliacas of Nginx web server:

```bash
kubectl create deployment app --image nginx:1.25-alpine --replicas=2

kubectl get pod -o wide
```

Expose the deployment with a headless service:

```bash
kubectl expose deployment app --name='headless-with-selectors' --cluster-ip='None'

kubectl get service

kubectl get ep
```

Check DNS records and test the connection (replace the IP below with the one of your Pod):

```bash
kubectl run --image busybox:1.35 -it --rm -- bash

nslookup headless-with-selectors

wget -O - 10.244.0.4

wget -O - headless-with-selectors.default.svc.cluster.local

exit
```

## Without Selectors

Create a headless service without selectors:

```bash
kubectl create -f headless-service-without-selectors.yaml

kubectl get service

kubectl get ep
```

Check DNS records and test the connection:

```bash
kubectl run --image busybox:1.35 -it --rm -- bash

nslookup headless-without-selectors

wget -O - headless-without-selectors.default.svc.cluster.local

exit
```

Create an EndpointSlice (remeber to set an IP of your Pod's):

```bash
kubectl create -f endpointslice.yaml

kubectl get endpointslices
```

Check DNS records and test the connection:

```bash
kubectl run --image busybox:1.35 -it --rm -- bash

nslookup headless-without-selectors

wget -O - headless-without-selectors.default.svc.cluster.local

nslookup 10-244-0-3.headless-without-selectors.default.svc.cluster.local

wget -O - 10-244-0-3.headless-without-selectors.default.svc.cluster.local

exit
```

## External Name

Create an ExternalName type of service:

```bash
kubectl create -f headless-service-without-selectors-external-name.yaml

kubectl get service
```

Check DNS records and test the connection:

```bash
kubectl run --image busybox:1.35 -it --rm -- bash

nslookup headless-without-selectors-external-name

ping headless-without-selectors-external-name.default.svc.cluster.local

exit
```
