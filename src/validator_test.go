package main

import (
	"testing"
)

func Test_DailyWorksValidate(t *testing.T) {
	type args struct {
		DailyWorks DailyWorks
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "OK",
			args: args{DailyWorks: DailyWorks{
				DailyWorks: []DailyWork{
					{
						Date:  "2006-01-02",
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
			wantErr: false,
		},
		{
			name: "invalidDate",
			args: args{DailyWorks: DailyWorks{
				DailyWorks: []DailyWork{
					{
						Date:  "aaa",
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
			wantErr: true,
		},
		{
			name: "invalidStart",
			args: args{DailyWorks: DailyWorks{
				DailyWorks: []DailyWork{
					{
						Date:  "2006-01-02",
						Start: "99:00",
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
			wantErr: true,
		},
		{
			name: "invalidEnd",
			args: args{DailyWorks: DailyWorks{
				DailyWorks: []DailyWork{
					{
						Date:  "2006-01-02",
						Start: "09:00",
						End:   "99:99",
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
			wantErr: true,
		},
		{
			name: "invalidProjects",
			args: args{DailyWorks: DailyWorks{
				DailyWorks: []DailyWork{
					{
						Date:  "2006-01-02",
						Start: "09:00",
						End:   "99:99",
						Projects: []Project{
							{
								Name:     "test",
								Duration: "99:00",
							},
						},
					},
				},
			},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		err := tt.args.DailyWorks.validate()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. DailyWorksValidate() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
	}
}
