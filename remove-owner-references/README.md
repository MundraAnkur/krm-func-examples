
# remove-owner-references

## Overview

The KRM function created in code removes resources with owner references.  


## Usage  

To test the function locally:

Convert the KRM resources and FunctionConfig resource to `ResourceList`, and then pipe the ResourceList as stdin to your function

`kpt fn source data | go run main.go`

  
Using Docker
`export FN_CONTAINER_REGISTRY=<Your  GCR  or  docker  hub>`
`export TAG= <Your  KRM  function  tag>`
`docker build . -t ${FN_CONTAINER_REGISTRY}/${FUNCTION_NAME}:${TAG}`


For testing out the function
`kpt fn eval ./data --image ${FN_CONTAINER_REGISTRY}/${FUNCTION_NAME}:${TAG}`

