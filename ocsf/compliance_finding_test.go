package ocsf

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComplianceFindingUnmarshal(t *testing.T) {
	tests := []struct {
		name          string
		sampleData    string
		expectedError bool
	}{
		{"basic sample", "compliance_finding.json", false},
		{"security control profile", "compliance_finding_with_security_control_profile.json", false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			data, err := os.ReadFile(fmt.Sprintf("test_data/%s", test.sampleData))
			if err != nil {
				t.Fatal(err)
			}
			var complianceFinding ComplianceFinding
			err = json.Unmarshal(data, &complianceFinding)
			if !test.expectedError {
				assert.NoError(t, err)
			}
		})
	}
}
