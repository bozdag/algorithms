package brackets

import "testing"

// TestIsBalanced_Balanced a small unit test for a few cases
func TestIsBalanced_Balanced(t *testing.T) {
	testData := []string{
		"(())",
		"(()",
		")(())",
		"((",
		"(())(())()",
		"{[()]}",
		"{[(])}",
		"{{[[(())]]}}",
	}

	expectedResults := []bool {
		true,
		false,
		false,
		false,
		true,
		true,
		false,
		true,
	}

	for i := 0; i <len(testData); i++ {
		if actualResult := IsBalanced(testData[i]); actualResult != expectedResults[i] {
			t.Errorf("%v: Expected %v, Actual : %v", testData[i], actualResult, expectedResults[i])
		} 
	}
}


