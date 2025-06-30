package schema_test

import (
	"testing"

	namespace "github.com/server/internal/core/namespace/domain"
)

func TestFQPermissionNameFromNamespace(t *testing.T) {
	type args struct {
		namespace namespace.NameSpaceName
		verb      string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basic namespace and verb",
			args: args{
				namespace: "app/user",
				verb:      "delete",
			},
			want: "app_user_delete",
		},
		{
			name: "namespace using alias",
			args: args{
				namespace: "user",
				verb:      "delete",
			},
			want: "app_user_delete",
		},
		{
			name: "namespace without resource",
			args: args{
				namespace: "hello",
				verb:      "delete",
			},
			want: "hello_default_delete",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.namespace.BuildRelation(tt.args.verb); got != tt.want {
				t.Errorf("tt.args.namespace.BuildRelation() = %v, want %v", got, tt.want)
			}
		})
	}
}
