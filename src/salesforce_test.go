package main

import (
	"reflect"
	"testing"
)

func Test_FindProject(t *testing.T) {
	type args struct {
		Project Project
		Repo    []projectRepo
	}
	tests := []struct {
		name       string
		args       args
		want       bool
		wantDetail projectRepo
	}{
		{
			name: "Exist.",
			args: args{
				Project: Project{
					Name:     "aaa",
					Duration: "09:00",
				},
				Repo: []projectRepo{
					{
						Name:     "IRI000/aaa",
						CssClass: "csstest",
						Number:   1,
					},
					{
						Name:     "IRI000/bbb",
						CssClass: "csstest",
						Number:   2,
					},
				},
			},
			want: true,
			wantDetail: projectRepo{
				Name:     "IRI000/aaa",
				CssClass: "csstest",
				Number:   1,
			},
		},
		{
			name: "NotExist.",
			args: args{
				Project: Project{
					Name:     "ccc",
					Duration: "09:00",
				},
				Repo: []projectRepo{
					{
						Name:     "IRI000/aaa",
						CssClass: "csstest",
						Number:   1,
					},
					{
						Name:     "IRI000/bbb",
						CssClass: "csstest",
						Number:   2,
					},
				},
			},
			want:       false,
			wantDetail: projectRepo{},
		},
	}
	for _, tt := range tests {
		got, detail := findProject(tt.args.Project, tt.args.Repo)
		if got != tt.want {
			t.Errorf("%q. findProject() got = %v, want %v", tt.name, got, tt.want)
		}
		if !reflect.DeepEqual(detail, tt.wantDetail) {
			t.Errorf("%q. findProject() = %v, want %v", tt.name, detail, tt.wantDetail)
		}
	}
}
