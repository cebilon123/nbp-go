package console

import "testing"

func Test_consoleWriter_WriteString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantN   int
		wantErr bool
	}{
		{name: "Valid length", args: struct{ s string }{s: "test"}, wantN: 4, wantErr: false},
		{name: "Valid length", args: struct{ s string }{s: "tested"}, wantN: 6, wantErr: false},
		{name: "Valid length", args: struct{ s string }{s: "tester"}, wantN: 6, wantErr: false},
		{name: "Zero length", args: struct{ s string }{s: ""}, wantN: 0, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := consoleWriter{}
			gotN, err := c.WriteString(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("WriteString() gotN = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
