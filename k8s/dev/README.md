# K8s Dev Env

Steps:

Start the local minikube instance and let that do it's thing.
```sh
minikube start
```

Run the startup script located in the `k8s` directory. This will run the `dev` setup script too.
```sh
./k8s/setup.sh
```

Forward the ingress controller port to access the web app.
```sh
kubectl port-forward -n ingress-nginx service/ingress-nginx-controller 8080:80
```

