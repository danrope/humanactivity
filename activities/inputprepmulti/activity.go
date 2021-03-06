package inputprepmulti

import (
	"encoding/json"

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
	data := context.GetInput("data")
	var tmpData [][]float64
	switch data.(type) {
	case []string:
		json.Unmarshal([]byte(data.([]string)[0]), &tmpData)
	case string:
		json.Unmarshal([]byte(data.(string)), &tmpData)
	case []interface{}:
		t1 := data.([]interface{})[0]
		var t2 []float64
		for flv := range t1.([]interface{}) {
			t2 = append(t2, float64(flv))
		}
		tmpData = [][]float64{t2}

	default:
		tmpData = data.([][]float64)
	}

	m := util.MakeLags(tmpData, "_", "__")

	for k, v := range m {
		context.SetOutput(k, v)
	}

	return true, nil
}
