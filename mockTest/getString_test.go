package mockTest

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
)

// mock test
func TestMock(t *testing.T) {
	mockCtr := gomock.NewController(t)
	defer mockCtr.Finish()

	mockGetString := NewMockmockDemo(mockCtr)
	mockGetString.EXPECT().GetString().Return("123")
	//mockGetString.EXPECT().Recv().Return("").Times(3)
	goVer := GoMockDemo(mockGetString)

	if goVer != "12345" {
		t.Error("Get wrong version %s", goVer)
	}
}
