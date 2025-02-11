package sample

import (
	"io"
	btz "bytes"
	"hash"
	"log"
	"net"
	"net/http"
	t1 "html/template"
	t2 "text/template"
	"github.com/golang/mock/gomock"
	"github.com/golang/mock/sample/imp1"
	renamed2 "github.com/golang/mock/sample/imp2"
	. "github.com/golang/mock/sample/imp3"
	imp_four "github.com/golang/mock/sample/imp4"
	"testing"
)

func TestRemember(t *testing.T) {
	// Arranging the mock controller
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// The test scenarios to check
	testCases := []struct {
		name           string
		keys           []string
		values         []interface{}
		expectedError  error
		firstKeyIsA    bool
	}{
		// Scenario 1
		{
			name:           "Normal Operation Test for Remember Function",
			keys:           []string{"key1", "key2"},
			values:         []interface{}{"value1", "value2"},
			expectedError:  nil,
			firstKeyIsA:    false,
		},
		// Scenario 2
		{
			name:           "Error Testing for Remember Function",
			keys:           []string{"key1", "key2"},
			values:         []interface{}{"value1", "value2"},
			expectedError:  net.UnknownNetworkError("error"),
			firstKeyIsA:    false,
		},
		// Scenario 3
		{
			name:          "Ellip Method Test for Remember Function",
			keys:          []string{"a", "key2"},
			values:        []interface{}{"value1", "value2"},
			firstKeyIsA:   true,
		},
		// Scenario 4
		{
			name:          "EllipOnly Method Test for Remember Function",
			keys:          []string{"a", "key2"},
			values:        []interface{}{"value1", "value2"},
			firstKeyIsA:   true,
		},
	}

	// Run the tests cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			// Arrange
			index := NewMockIndex(mockCtrl)

			// Define our expectations
			for i, k := range tc.keys {
				index.EXPECT().Put(k, tc.values[i])
			}
			index.EXPECT().NillableRet().Return(tc.expectedError)

			if tc.firstKeyIsA {
				index.EXPECT().Ellip(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any())
				index.EXPECT().Ellip(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any())
				index.EXPECT().EllipOnly(gomock.Any())
			} 

			// Act
			Remember(index, tc.keys, tc.values)

			// Assert 
			// checking log output in case of error scenarios
			if tc.expectedError != nil {
				buf := btz.Buffer{}
				log.SetOutput(&buf)
				defer func() {
					log.SetOutput(io.Discard)
				}()
				if buf.String() == "" {
					t.Errorf("Error case did not log any message.")
				}
			}

		}) // closing t.Run
	} // closing the loop
}
