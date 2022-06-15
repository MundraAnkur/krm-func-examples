package main

import (
	"os"
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
)


// This is the main logic. rl is the input `ResourceList` which has the `FunctionConfig` and `Items` fields.
// You can modify the `Items` and add result information to `rl.Result`.
func Run(rl *fn.ResourceList) (bool,error) {
	
	for _, kubeObject := range rl.Items {
		//remove owner references
		kubeObject.RemoveNestedField("metadata", "ownerReferences")	
	}
	
	//This result message will be displayed in the function evaluation time.
	rl.Results = append(rl.Results, fn.GeneralResult("Removed owner references", fn.Info))
	return true,nil

}

func main() {
	// `AsMain` accepts a `ResourceListProcessor` interface.
	// You can explore other `ResourceListProcessor` structs in the SDK or define your own.
	if err := fn.AsMain(fn.ResourceListProcessorFunc(Run)); err != nil {
		os.Exit(1)
	}
}
