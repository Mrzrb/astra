package astra

import (
	"reflect"

	"github.com/Mrzrb/astra/astTraversal"
)

func ExtractBindingTags(fields map[string]Field) (bindingTags []astTraversal.BindingTagType, uniqueBindings bool) {
	for _, field := range fields {
		var previousTag astTraversal.BindingTag
		for bindingType, bindingTag := range field.StructFieldBindingTags {
			if previousTag != (astTraversal.BindingTag{}) && !uniqueBindings {
				uniqueBindings = !reflect.DeepEqual(bindingTag, previousTag)
			}

			if !Contains(bindingTags, bindingType) {
				bindingTags = append(bindingTags, bindingType)
			}

			previousTag = bindingTag
		}
	}

	if len(bindingTags) == 0 {
		return astTraversal.BindingTags, uniqueBindings
	}

	return bindingTags, uniqueBindings
}

func ContentTypeToBindingTag(contentType string) astTraversal.BindingTagType {
	mimetypeToBindingTagMap := map[string]astTraversal.BindingTagType{
		"application/json":                  astTraversal.JSONBindingTag,
		"application/xml":                   astTraversal.XMLBindingTag,
		"application/x-www-form-urlencoded": astTraversal.FormBindingTag,
		"multipart/form-data":               astTraversal.FormBindingTag,
		"application/yaml":                  astTraversal.YAMLBindingTag,
	}

	return mimetypeToBindingTagMap[contentType]
}

func BindingTagToContentTypes(bindingTag astTraversal.BindingTagType) []string {
	bindingTagToMimetypeMap := map[astTraversal.BindingTagType][]string{
		astTraversal.JSONBindingTag: {"application/json"},
		astTraversal.XMLBindingTag:  {"application/xml"},
		astTraversal.FormBindingTag: {"application/x-www-form-urlencoded", "multipart/form-data"},
		astTraversal.YAMLBindingTag: {"application/yaml"},
	}

	return bindingTagToMimetypeMap[bindingTag]
}
