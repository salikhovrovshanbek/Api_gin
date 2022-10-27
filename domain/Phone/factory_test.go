package Phone

import (
	"Api_With_AbdulHamid/pkg/id"
	"reflect"
	"testing"
)

func TestFactory_NewPhone(t *testing.T) {
	type fields struct {
		idGenerator id.IGenerator
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Phone
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Factory{
				idGenerator: tt.fields.idGenerator,
			}
			got, err := f.NewPhone(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPhone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPhone() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFactory(t *testing.T) {
	type args struct {
		ig id.IGenerator
	}
	tests := []struct {
		name string
		args args
		want Factory
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFactory(tt.args.ig); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}
