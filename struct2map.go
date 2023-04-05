package struct2map

import (
	"reflect"
	"strings"
)

func Struct2Map(inStruct interface{}, tag string, tagNameParser func(string) string, omitEmpty bool, fieldExcluded ...string) map[string]interface{} {
	ret := make(map[string]interface{})

	inStructValue := reflect.ValueOf(inStruct)

	// 结构体指针
	if inStructValue.Kind() == reflect.Ptr {
		inStructValue = reflect.Indirect(inStructValue)
	}

	// 必须为结构体
	if !inStructValue.IsValid() || inStructValue.IsZero() || inStructValue.Kind() != reflect.Struct {
		return ret
	}
	inStructType := inStructValue.Type()

	// 空间换时间
	var fieldsExcludedMap map[string]struct{}
	if len(fieldExcluded) != 0 {
		fieldsExcludedMap = make(map[string]struct{})
		for _, key := range fieldExcluded {
			fieldsExcludedMap[key] = struct{}{}
		}
	}

	numField := inStructValue.NumField()
	for i := 0; i < numField; i++ {
		structFiled := inStructType.Field(i)
		// 非导出
		if !structFiled.IsExported() {
			continue
		}

		// 指定tag
		tagValue, ok := structFiled.Tag.Lookup(tag)
		if !ok {
			continue
		}

		tagName := tagNameParser(tagValue)
		if tagName == "" {
			continue
		}

		// 需要排除
		if _, ok := fieldsExcludedMap[tagName]; ok {
			continue
		}

		// 填充值
		structFiledValue := inStructValue.Field(i)
		if structFiledValue.IsZero() {
			if !omitEmpty {
				ret[tagName] = structFiledValue.Interface()
			}
		} else {
			ret[tagName] = structFiledValue.Interface()
		}
	}

	return ret
}

/*
JsonTagNameParser 从json tag中提取名字

json:"authKey,omitempty"
*/
func JsonTagNameParser(tagValue string) string {
	if tagValue == "" {
		return ""
	}

	tagNameAndOptions := strings.Split(tagValue, ",")
	for _, val := range tagNameAndOptions {
		if val == "-" {
			return ""
		}

		if val != "" {
			return val
		}

		break
	}

	return ""
}

/*
ProtobufTagNameParser 从protobuf tag中提取名字

`protobuf:"bytes,1,opt,name=authKey,proto3"`
*/
func ProtobufTagNameParser(tagValue string) string {
	if tagValue == "" {
		return ""
	}

	tagNameAndOptions := strings.Split(tagValue, ",")
	for _, val := range tagNameAndOptions {
		if strings.HasPrefix(val, "name=") {
			return strings.Split(val, "=")[1]
		}
	}

	return ""
}
