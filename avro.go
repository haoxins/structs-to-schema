package structstoschema

import (
	"errors"
	"reflect"
)

const (
	AvroTypeRecord = "record"
	AvroTypeString = "string"
	AvroTypeBool   = "boolean"
	AvroTypeInt32  = "int"
	AvroTypeInt64  = "long"
)

type Element struct {
	Name     string
	Type     string
	Children []Element
}

func ParseAvro(in any) (string, error) {
	root := Element{}
	t := reflect.TypeOf(in)

	if t.Kind() != reflect.Struct {
		return "", errors.New("Invalid type")
	}

	root.Name = t.Name()
	root.Type = AvroTypeRecord
	root.Children = ParseAvroElements(t)

	return root.toAvsc(), nil
}

func (e *Element) toAvsc() string {
	s := ""
	if len(e.Children) == 0 {
		s = `{"name":"` + e.Name + `","type":"` + e.Type + `"}`
	} else {
		fields := ""
		for _, child := range e.Children {
			fields += "," + child.toAvsc()
		}
		fields = fields[1:]
		s = `{"name":"` + e.Name + `","type":"` + e.Type + `","fields":[` + fields + `]}`
	}

	return s
}

func ParseAvroElements(t reflect.Type) []Element {
	elements := []Element{}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Type.Kind() != reflect.Struct {
			elements = append(elements, Element{
				Name: field.Name,
				Type: castTypeToAvro(field.Type.Kind()),
			})
		} else {
			elements = append(elements, Element{
				Name:     field.Name,
				Type:     AvroTypeRecord,
				Children: ParseAvroElements(field.Type),
			})
		}
	}

	return elements
}

func castTypeToAvro(k reflect.Kind) string {
	switch k {
	case reflect.String:
		return AvroTypeString
	case reflect.Bool:
		return AvroTypeBool
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		return AvroTypeInt32
	case reflect.Int64:
		return AvroTypeInt64
	// TODO: Add more types
	default:
		panic(errors.New("Invalid type"))
	}
}
