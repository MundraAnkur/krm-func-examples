package main

import (
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
)

// clean resource manifest
func cleanResourceManifest(kubeObject *fn.KubeObject) error {
	var err error
	err = cleanMetaData(kubeObject)
	err = cleanStatus(kubeObject)
	err = removeNodeName(kubeObject)
	return err
}

// Neat metadata
func cleanMetaData(kubeObject *fn.KubeObject) error {
	// Remove last-applied-configuration annotation
	if _, err := kubeObject.RemoveNestedField("metadata", "annotations", "kubectl.kubernetes.io/last-applied-configuration"); err != nil {
		return err
	}

	// Get metadata fields
	metadata := make(map[string]interface{})
	metadata["name"] = kubeObject.GetName()
	metadata["namespace"] = kubeObject.GetNamespace()
	metadata["labels"] = kubeObject.GetLabels()
	metadata["annotations"] = kubeObject.GetAnnotations()

	// Remove metadata
	if _, err := kubeObject.RemoveNestedField("metadata"); err != nil {
		return err
	}
	// Add metadata
	kubeObject.SetName(metadata["name"].(string))
	if metadata["namespace"].(string) != "" {
		kubeObject.SetNamespace(metadata["namespace"].(string))
	}
	for k, v := range metadata["labels"].(map[string]string) {
		kubeObject.SetLabel(k, v)
	}
	for k, v := range metadata["annotations"].(map[string]string) {
		kubeObject.SetAnnotation(k, v)
	}
	return nil
}

// clean resource status
func cleanStatus(kubeObject *fn.KubeObject) error {
	if _, err := kubeObject.RemoveNestedField("status"); err != nil {
		return err
	}
	return nil
}

// remove node name
func removeNodeName(kubeObject *fn.KubeObject) error {
	if _, err := kubeObject.RemoveNestedField("spec", "nodeName"); err != nil {
		return err
	}
	return nil
}
