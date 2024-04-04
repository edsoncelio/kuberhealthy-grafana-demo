# kuberhealthy-grafana-demo
Configs to setup kuberhealthy with Grafana and Prometheus.

## Prerequisites
* Docker
* kind (or other tool to run k8s locally e.g. K3s)
* Helm (version 3.xx or higher)

## How to Install
1. Add the helm repos
```bash
helm repo add kuberhealthy https://kuberhealthy.github.io/kuberhealthy/helm-repos
helm repo add grafana https://grafana.github.io/helm-charts
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
```

2. Install Kuberhealthy
```bash
helm upgrade --install kuberhealthy --create-namespace -f kuberhealthy/values.yaml kuberhealthy/kuberhealthy -n kuberhealthy
``` 

By default there are only one check enabled (deployment check).

3. Install Prometheus
```bash
helm upgrade --install prometheus --create-namespace -f prometheus/values.yaml prometheus-community/prometheus -n prometheus
``` 

To check if prometheus is working and is getting the kuberhealthy metrics, make a port-forward:   
`kubectl port-forward svc/prometheus-server 8000:80 -n prometheus`   

If everything is working, you will see the target in `http://localhost:8000/targets?search=kuberhealthy`.

4. Install Grafana
```bash
helm upgrade --install grafana --create-namespace -f grafana/values.yaml grafana/grafana -n grafana
``` 

## Accessing the Grafana Instance 
To access the Grafana instance:
```bash
kubectl port-forward svc/grafana 3000:80 -n grafana
```

The user is `admin` and the password is `$ecret`.
You can see the kuberhealthy dashboard at http://localhost:3000/d/kuberhealthy/kuberhealthy. 



