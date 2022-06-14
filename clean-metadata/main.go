package main

import (
	"os"

	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
)

// Run
// This is the main logic. rl is the input `ResourceList` which has the `FunctionConfig` and `Items` fields.
// You can modify the `Items` and add result information to `rl.Result`.
func Run(rl *fn.ResourceList) (bool, error) {
	for _, kubeObject := range rl.Items {
		if err := neatMetadata(kubeObject); err != nil {
			return false, err
		}
	}

	//	This result message will be displayed in the function evaluation time.
	rl.Results = append(rl.Results, fn.GeneralResult("Remove unnecessary fields from metadata", fn.Info))
	return true, nil
}

func main() {
	if err := fn.AsMain(fn.ResourceListProcessorFunc(Run)); err != nil {
		os.Exit(1)
	}
}
