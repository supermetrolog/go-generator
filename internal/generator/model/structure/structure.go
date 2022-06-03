package structure

import "errors"

type structure struct {
	name   string
	fields []*field
}

type field struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

func (f *field) Validate() (bool, error) {
	if len(f.Name) == 0 {
		return false, errors.New("Name is not be blank")
	}
	if len(f.Type) == 0 {
		return false, errors.New("Name is not be blank")
	}
	return true, nil
}

func NewStructureField(name, Type, tag, comment string) (*field, error) {
	field := &field{
		Name:    name,
		Type:    Type,
		Tag:     tag,
		Comment: comment,
	}

	ok, err := field.Validate()
	if ok {
		return field, err
	}
	return nil, err
}

func (s *structure) AddField(fields ...*field) error {
	if fields == nil {
		return errors.New("field is not be <nil>")
	}
	for _, field := range fields {
		s.fields = append(s.fields, field)
	}
	return nil
}
