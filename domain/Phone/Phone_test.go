package Phone

import (
	"github.com/google/uuid"
	"reflect"
	"testing"
)

func TestPhone_Id(t *testing.T) {
	type fields struct {
		id   uuid.UUID
		name string
	}
	tests := []struct {
		name   string
		fields fields
		want   uuid.UUID
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Phone{
				id:   tt.fields.id,
				name: tt.fields.name,
			}
			if got := p.Id(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Id() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPhone_Name(t *testing.T) {
	type fields struct {
		id   uuid.UUID
		name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Phone{
				id:   tt.fields.id,
				name: tt.fields.name,
			}
			if got := p.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPhone_validate(t *testing.T) {
	type fields struct {
		id   uuid.UUID
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Phone{
				id:   tt.fields.id,
				name: tt.fields.name,
			}
			if err := p.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
