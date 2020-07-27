package bihua

import (
	"testing"
)

// TestGeNum TestGeNum
func TestGeNum(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		wantNum int
	}{
		{name: "test1",args: struct{ str string }{str: "李"},wantNum: 7},
		{name: "test2",args: struct{ str string }{str: "启"},wantNum: 7},
		{name: "test3",args: struct{ str string }{str: "爱"},wantNum: 10},
		{name: "test4",args: struct{ str string }{str: "李启爱"},wantNum: 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNum := GeNum(tt.args.str); gotNum != tt.wantNum {
				t.Errorf("GeNum() = %v, want %v", gotNum, tt.wantNum)
			}
		})
	}
}