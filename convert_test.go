package kata

import "testing"

func Test_ConvertToDigitalNumber_By_Number_1_Should_Get_1(t *testing.T) {
	expectedResult := 1
	actualResult := convertToDigitalNumber(1)

	if expectedResult != actualResult {
		t.Errorf("Expect %v but got %v", expectedResult, actualResult)
	}
}
