package inputprepmulti

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/danrope/humanactivity/util"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	// defer func() {
	// 	if r := recover(); r != nil {
	// 		t.Failed()
	// 		t.Errorf("panic during execution: %v", r)
	// 	}
	// }()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	data := util.ReadCSV()
	tc.SetInput("data", data[0:11])
	act.Eval(tc)

	z10 := tc.GetOutput("_z10__")
	assert.Equal(t, z10, 8.067265, "Output should contain expected z10 result")
}
