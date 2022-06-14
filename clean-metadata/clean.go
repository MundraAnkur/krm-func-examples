package main

import (
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
)

// Neat metadata
func neatMetadata(kubeObject *fn.KubeObject) error {
	// Remove last-applied-configuration annotation
	annotation := kubeObject.GetAnnotation("kubectl.kubernetes.io/last-applied-configuration")
	if annotation != "" {
		_, err := kubeObject.RemoveNestedField("metadata", "annotations", "kubectl.kubernetes.io/last-applied-configuration")
		if err != nil {
			return err
		}
	}

	// Get metadata fields
	metadata := make(map[string]interface{})
	metadata["name"] = kubeObject.GetName()
	metadata["namespace"] = kubeObject.GetNamespace()
	metadata["labels"] = kubeObject.GetLabels()
	metadata["annotations"] = kubeObject.GetAnnotations()

	// Remove metadata
	kubeObject.RemoveNestedField("metadata")
	// Add metadata
	kubeObject.SetName(metadata["name"].(string))
	kubeObject.SetNamespace(metadata["namespace"].(string))
	for k, v := range metadata["labels"].(map[string]string) {
		kubeObject.SetLabel(k, v)
	}

	for k, v := range metadata["annotations"].(map[string]string) {
		kubeObject.SetAnnotation(k, v)
	}

	return nil
}
