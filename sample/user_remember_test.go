package sample

import (
	"testing"
	"bytes"
	"log"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// A mock type to use for testing of the Remember function.
type MockIndex struct {
	mock *gomock.Controller
	Index
}

func (m *MockIndex) Put(key string, value interface{}) {
	m.mock.RecordCall(m, "Put", key, value)
}

func (m *MockIndex) NillableRet() error {
	return m.mock.Call(m, "NillableRet").Error(0)
}

func (m *MockIndex) Ellip(format string, a ...interface{}) {
	m.mock.RecordCall(m, "Ellip", format, a)
}

func (m *MockIndex) EllipOnly(a ...string) {
	m.mock.RecordCall(m, "EllipOnly", a)
}

// TestRemember function to test the Remember function.
func TestRemember(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// TODO: Ensure these are your expected calls. If not, please adjust accordingly.
	index := &MockIndex{
		mock: ctrl,
	}

	tests := []struct {
		name    string
		keys    []string
		values  []interface{}
		// Define any other variables you expect for the test.
		expectNillableRetError bool
		expectEllipCall        bool
	}{
		// TODO: Add Test Scenarios Here
		{
			name:    "Valid Inputs",
			keys:    []string{"a", "b"},
			values:  []interface{}{"value1", "value2"},
			expectNillableRetError: false,
			expectEllipCall:        true,
		},
		{
			name:    "Empty keys and values",
			keys:    []string{},
			values:  []interface{}{},
			expectNillableRetError: false,
			expectEllipCall:        false,
		},
		{
			name:    "Error from NillableRet",
			keys:    []string{"b", "c"},
			values:  []interface{}{"value3", "value4"},
			expectNillableRetError: true,
			expectEllipCall:        false,
		},
		{
			name:    "Check handling the first key as 'a'",
			keys:    []string{"a", "d"},
			values:  []interface{}{"value5", "value6"},
			expectNillableRetError: false,
			expectEllipCall:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			gomock.InOrder(
				index.EXPECT().Put(gomock.Any(), gomock.Any()).AnyTimes(),
				index.EXPECT().NillableRet().Return(nil).Times(1),
			)
			if tt.expectNillableRetError {
				index.EXPECT().NillableRet().Return(fmt.Errorf("some error")).Times(1)
			}
			if tt.expectEllipCall {
				index.EXPECT().Ellip(gomock.Any(), gomock.Any()).AnyTimes()
				index.EXPECT().EllipOnly(gomock.Any()).AnyTimes()
			}

			// Act
			Remember(index, tt.keys, tt.values)

			// Assert Calls Occurred
			index.mock.AssertExpectations(t)
		})
	}
}
