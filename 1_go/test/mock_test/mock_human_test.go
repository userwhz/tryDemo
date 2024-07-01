package main

import (
	"github.com/golang/mock/gomock"
	"testing"
	"tryDemo/1_go/test/mock"
)

func Test_mock_human(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockHuman := mock.NewMockHuman(ctrl)
	mockHuman.EXPECT().Speak().DoAndReturn(func() string {
		return "hhh"
	})
	mockHuman.EXPECT().Speak().Return("hhh1").AnyTimes()
	output := mockHuman.Speak()
	output1 := mockHuman.Speak()

	t.Errorf("output: %s", output)
	t.Errorf("output: %s", output1)
}
