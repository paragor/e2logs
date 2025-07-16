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
### To Deploy on the cluster via a Helm Chart

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

