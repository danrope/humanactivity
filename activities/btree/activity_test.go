package btree

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
	countJog := 0
	trials := 100

	for i := 100; i < 100+trials; i++ {
		m := util.MakeLags(data[i:i+11], "_", "__")

		for k, v := range m {
			tc.SetInput(k, v)
		}
		act.Eval(tc)

		result := tc.GetOutput("result")

		if result == "Jogging" {
			countJog++
		}

	}

	assert.True(t, float64(countJog)/float64(trials) >= 0.95, "Not enough Jogging predictions")

}
