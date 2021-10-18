package models

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestColumnsMapAppend(t *testing.T) {
	type fields struct {
		Items []ColumnMap
	}
	type args struct {
		item ColumnMap
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "TestAdd",
			fields: fields{Items: []ColumnMap{
				{
					Name:   "Id",
					Label:  "Ид",
					Format: "string",
				},
				{Name: "Name",
					Label:  "Наименование",
					Format: "string",
				}}},
			args: args{item: ColumnMap{
				Name:   "Created_at",
				Label:  "Время создания",
				Format: "datetime",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cm := &columnsMap{
				Items: tt.fields.Items,
			}
			cm.Append(tt.args.item)
		})
	}
}

func TestColumnsMapGetColumn(t *testing.T) {
	type fields struct {
		Items []ColumnMap
	}
	type args struct {
		columnName string
	}
	columnMapData := []ColumnMap{
		{
			Name:   "created_at",
			Label:  "Дата создания",
			Format: "datetime",
		},
		{Name: "Name",
			Label:  "Наименование",
			Format: "string",
		}}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []interface{}
	}{
		{
			name:   "TestName",
			fields: fields{Items: columnMapData},
			args:   args{columnName: "Name"},
			want:   []interface{}{"created_at", "Name"},
		},
		{
			name:   "TestFormat",
			fields: fields{Items: columnMapData},
			args:   args{columnName: "Format"},
			want:   []interface{}{"datetime", "string"},
		},
		{
			name:   "TestLabel",
			fields: fields{Items: columnMapData},
			args:   args{columnName: "Label"},
			want:   []interface{}{"Дата создания", "Наименование"},
		},
		{
			name:   "TestErr",
			fields: fields{Items: columnMapData},
			args:   args{columnName: "Labl"},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cm := &columnsMap{
				Items: tt.fields.Items,
			}
			if got := cm.GetColumn(tt.args.columnName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColumnMapGetField(t *testing.T) {
	test := ColumnMap{
		Name:   "Id",
		Label:  "Ид",
		Format: "integer",
	}
	firstTest, ok := test.getField("Name")
	assert.Equal(t, true, ok)
	assert.Equal(t, "Id", firstTest)

	secondTest, ok := test.getField("Label")
	assert.Equal(t, true, ok)
	assert.Equal(t, "Ид", secondTest)

	ThreeTest, ok := test.getField("Format")
	assert.Equal(t, true, ok)
	assert.Equal(t, "integer", ThreeTest)
}
