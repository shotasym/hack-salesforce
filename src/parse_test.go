package main

import (
	"reflect"
	"testing"
)

func Test_ParseJson(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    DailyWorks
		wantErr bool
	}{
		{
			name:    "OK",
			args:    args{b: []byte(`[{"date": "2019-01-01", "start": "09:00", "end": "19:00", "projects": [{"name": "test", "duration": "08:00"}]}]`)},
			wantErr: false,
			want: DailyWorks{
				DailyWorks: []DailyWork{
					{
						Date:  "2019-01-01",
						Start: "09:00",
						End:   "19:00",
						Projects: []Project{
							{
								Name:     "test",
								Duration: "08:00",
							},
						},
					},
				},
			},
		},
		{
			name:    "NotJson",
			args:    args{b: []byte(``)},
			wantErr: true,
			want:    DailyWorks{},
		},
	}
	for _, tt := range tests {
		var dw DailyWorks
		err := dw.ParseJson(tt.args.b)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. ParseJson() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(dw, tt.want) {
			t.Errorf("%q. ParseJson() = %v, want %v", tt.name, dw, tt.want)
		}
	}
}
