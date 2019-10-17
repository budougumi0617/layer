package layer

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLayer(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string

		in        []byte
		want      *Layer
		wantError bool
	}{
		{
			name: "simple",
			in:   []byte("[\"hoge\",\"bar\"]"),
			want: &Layer{
				Packages: []string{"hoge", "bar"},
				Inside:   nil,
				Raw:      []interface{}{"hoge", "bar"},
			},
		},
		{
			name: "nested",
			in:   []byte("[\"hoge\",\"bar\",12,[\"nesthoge\",\"nestbar\"]]"),
			want: &Layer{
				Packages: []string{"hoge", "bar"},
				Inside: &Layer{
					Packages: []string{"nesthoge", "nestbar"},
					Inside:   nil,
					Raw:      []interface{}{"nesthoge", "nestbar"},
				},
				Raw: []interface{}{"hoge", "bar", float64(12), []interface{}{"nesthoge", "nestbar"}},
			},
		},
		{
			name:      "error",
			in:        []byte("[\"hoge\",\"bar\"}]"),
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := &Layer{}
			if err := json.Unmarshal(tt.in, got); tt.wantError && err == nil {
				t.Fatalf("want error, but not error")
			} else if !tt.wantError && err != nil {
				t.Fatalf("want no err, but has error %#v", err)
			}

			if !tt.wantError {
				if diff := cmp.Diff(got, tt.want); diff != "" {
					fmt.Printf("got = %+v\n", got)
					fmt.Printf("tt.want = %+v\n", tt.want)
					t.Fatalf("diff from want\n%s\n", diff)
				}
			}
		})
	}
}
