
# Clean Resource Metadata

## Overview

This function cleans the resource metadata and remove all the non essential fileds from metadata.
Only keeps name, namespace, labels, and annotations in metadata.  

## Testing the function locally:
Run 

```
kpt fn source data/service.yaml | go run *.go
```

### Example
**Input**

```
apiVersion: v1
kind: Service
metadata:
  annotations:
    kubectl.kubernetes.io/last-applied-configuration: |
      {"apiVersion":"v1","kind":"Service","metadata":{"annotations":{},"labels":{"app":"guestbook","tier":"frontend"},"name":"frontend","namespace":"guestbook"},"spec":{"ports":[{"port":80}],"selector":{"app":"guestbook","tier":"frontend"}}}
  creationTimestamp: "2022-06-14T16:49:17Z"
  labels:
    app: guestbook
    tier: frontend
  name: frontend
  namespace: guestbook
  resourceVersion: "479"
  uid: 0e19ac91-c96d-4e64-b443-c72733bf9734
spec:
  clusterIP: 10.109.22.148
  ...
```

**Output**
```
apiVersion: v1
kind: Service
metadata:
  labels:
    app: guestbook
    tier: frontend
  name: frontend
  namespace: guestbook
spec:
  clusterIP: 10.109.22.148
  ...
```
