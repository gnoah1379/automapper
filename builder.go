package automapper

import (
	"reflect"
)

type Builder interface {
	Set(fieldName string, transform TransformHandler) Builder
	Build() AutoMapper
}

type builder struct {
	mapper AutoMapper
}

func (b *builder) Set(fieldName string, transform TransformHandler) Builder {
	b.mapper.transforms[fieldName] = transform
	return b
}

func (b *builder) Build() AutoMapper {
	return b.mapper
}

func NewBuilder(srcTemplate interface{}, dstTemplate interface{}) Builder {
	result := new(builder)
	result.mapper.srcTemplateType = reflect.TypeOf(srcTemplate)
	result.mapper.dstTemplateType = reflect.TypeOf(dstTemplate)
	result.mapper.transforms = make(map[string]TransformHandler)
	for i := 0; i < result.mapper.dstTemplateType.NumField(); i++ {
		fieldName := result.mapper.dstTemplateType.Field(i).Tag.Get(TAG)
		if len(fieldName) < 1 {
			continue
		}
		result.mapper.transforms[fieldName] = Default
	}
	return result
}
