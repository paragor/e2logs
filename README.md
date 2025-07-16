# e2logs
Read and log k8s events in json format.

## Description
K8s have a greate tool for debugging - event system, but there is luck of handling such events in common monitoring system.
Project `e2logs` help to handle events - it just listen events, as `kubectl events -w` does, and put event to stdout log
in json format. 

You can find structure or event at [internal/controller/event_controller.go:56](internal/controller/event_controller.go).
To separate k8s events log from another logs you can filter by field `logger==k8s_events_logger` in your logger collector.

Example of event:
```
{
  "level": "info",
  "ts": "2025-07-16T15:28:00+07:00",
  "logger": "k8s_events_logger",
  "msg": "Error: ErrImagePull",
  "controller": "event",
  "controllerGroup": "",
  "controllerKind": "Event",
  "Event": {
    "name": "debug.1852adf2c2ef5128",
    "namespace": "metrics"
  },
  "name": "debug.1852adf2c2ef5128",
  "reconcileID": "3a033c56-418a-4d18-b347-aa8a9202b5b8",
  "namespace": "metrics",
  "reason": "Failed",
  "source": {
    "component": "kubelet",
    "host": "k8s-control-plane-1"
  },
  "type": "Warning",
  "action": "",
  "reporting_controller": "kubelet",
  "reporting_instance": "hel1-10-10-32-11",
  "event_time": "2025-07-16T15:28:00+07:00",
  "involved_object": {
    "kind": "Pod",
    "namespace": "metrics",
    "name": "debug",
    "uid": "33b5cea4-1e6c-4ef5-93bc-9ebfe6d05387",
    "apiVersion": "v1",
    "resourceVersion": "8489307",
    "fieldPath": "spec.containers{debug}"
  },
  "related": "nil"
}

```
## Getting Started

### Prerequisites
- go version v1.24.0+
- docker version 17.03+.
- kubectl version v1.11.3+.
- Access to a Kubernetes v1.11.3+ cluster.

### To Deploy on the cluster
**Build and push your image to the location specified by `IMG`:**

```sh
make docker-build docker-push IMG=<some-registry>/e2logs:tag
```

**NOTE:** This image ought to be published in the personal registry you specified.
And it is required to have access to pull the image from the working environment.
Make sure you have the proper permission to the registry if the above commands don’t work.

**Install the CRDs into the cluster:**

```sh
make install
```

**Deploy the Manager to the cluster with the image specified by `IMG`:**

```sh
make deploy IMG=<some-registry>/e2logs:tag
```

> **NOTE**: If you encounter RBAC errors, you may need to grant yourself cluster-admin
privileges or be logged in as admin.

**Create instances of your solution**
You can apply the samples (examples) from the config/sample:

```sh
kubectl apply -k config/samples/
```

>**NOTE**: Ensure that the samples has default values to test it out.

### To Uninstall
**Delete the instances (CRs) from the cluster:**

```sh
kubectl delete -k config/samples/
```

**Delete the APIs(CRDs) from the cluster:**

```sh
make uninstall
```

**UnDeploy the controller from the cluster:**

```sh
make undeploy
```

## Project Distribution

Following the options to release and provide this solution to the users.

### By providing a bundle with all YAML files

1. Build the installer for the image built and published in the registry:

```sh
make build-installer IMG=<some-registry>/e2logs:tag
```

**NOTE:** The makefile target mentioned above generates an 'install.yaml'
file in the dist directory. This file contains all the resources built
with Kustomize, which are necessary to install this project without its
dependencies.

2. Using the installer

Users can just run 'kubectl apply -f <URL for YAML BUNDLE>' to install
the project, i.e.:

```sh
kubectl apply -f https://raw.githubusercontent.com/<org>/e2logs/<tag or branch>/dist/install.yaml
```

### By providing a Helm Chart

```
helm repo add e2logs https://paragor.github.io/e2logs
helm repo update
helm install e2logs e2logs/e2logs
```

## License

Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

