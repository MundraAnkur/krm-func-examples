package main

import (
	"os"
	"fmt"
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"

)

//remove an Item from struct
func removeItem(ss * fn.KubeObject, ssSlice [] *fn.KubeObject) [] *fn.KubeObject {
    for idx, v := range ssSlice {
        if v == ss {
            return append(ssSlice[0:idx], ssSlice[idx+1:]...)
        }
    }
    return ssSlice
}


// This is the main logic to find resources with ownerReferences and remove them
func Run(rl *fn.ResourceList) (bool,error) {


	for _, kubeObject := range rl.Items {
		
		
		_,found,err := kubeObject.NestedSlice("metadata","ownerReferences")

		
		if(err != nil){
			fmt.Print("Encountered an error")
			return false,err
		}

		if(found){
			rl.Items = removeItem(kubeObject, rl.Items)
		}
		
	}
	
	//This result message will be displayed in the function evaluation time.
	rl.Results = append(rl.Results, fn.GeneralResult("Removed owner references", fn.Info))
	return true,nil

}

func main() {
	// `AsMain` accepts a `ResourceListProcessor` interface.
	if err := fn.AsMain(fn.ResourceListProcessorFunc(Run)); err != nil {
		os.Exit(1)
	}
}
