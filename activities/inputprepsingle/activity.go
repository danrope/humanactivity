package inputprepsingle

import (
	"strconv"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/danrope/humanactivity/util"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	// Get the activity data from the context
	data := context.GetInput("data").([][]float64)
	//m := stats(data)
	m := util.MakeLags(data, "", "")

	outStr := ""

	//Place all output in required format for input to scoring
	i := 0
	for k, v := range m {
		if i > 0 {
			outStr += ","
		}
		outStr += k + ":" + strconv.FormatFloat(v, 'f', -1, 64)
		i++
	}
	context.SetOutput("predictors", outStr)
	// for k, v := range m {
	// 	fmt.Println(k)

	// 	context.SetOutput(k, v)
	// }

	return true, nil
}
