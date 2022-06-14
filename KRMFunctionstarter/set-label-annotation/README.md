
# set-label

## Overview

This is a starter code just to understand and setup the workflow for KRM functions.

The KRM function created in code adds a label `tier:mysql` to all the resources of kind Deployment.  

## Testing the function :

## Locally

Convert the KRM resources and FunctionConfig resource to `ResourceList`, and then pipe the ResourceList as stdin to your function

`kpt fn source data | go run main.go`

## Using Docker:

`export FN_CONTAINER_REGISTRY=<Your  GCR  or  docker  hub>`

`export TAG= <Your  KRM  function  tag>`

`docker build . -t ${FN_CONTAINER_REGISTRY}/${FUNCTION_NAME}:${TAG}`


For testing out the function
`kpt fn eval ./data --image ${FN_CONTAINER_REGISTRY}/${FUNCTION_NAME}:${TAG}`

