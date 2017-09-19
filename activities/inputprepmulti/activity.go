package inputprepmulti

import (
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

	m := util.MakeLags(data, "_", "__")

	for k, v := range m {
		context.SetOutput(k, v)
	}

	return true, nil
}
