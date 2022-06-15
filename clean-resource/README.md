
# Clean Resource Metadata

## Overview

This function cleans the resource remove all the non essential fileds from the manifest.
- Clean Metadata - Only keeps name, namespace, labels, and annotations in metadata.
- Clean Status - remove status field
- Remove nodeName from spec 

## Testing the function locally:

```
kpt fn source data/deployment.yaml | go run *.go
```
