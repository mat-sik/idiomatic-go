package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
    input := Test{
        text: "hello",
    }

    out, err := validateStructStringFields(input)
    if err != nil {
        panic(err)
    }

    fmt.Println(out)
}

func validateStructStringFields(input any) (bool, error) {
    inputValue := reflect.ValueOf(input)
    inputType := inputValue.Type()

    inputKind := inputType.Kind()
    if inputKind != reflect.Struct {
        return false, fmt.Errorf("input is of kind: %s, but should be struct", inputKind.String())
    }

    for i := 0; i < inputType.NumField(); i++ {
        field := inputType.Field(i)

        tagValue, ok := field.Tag.Lookup(tagMaxStrLen)
        if !ok {
            continue
        }

        fieldType := field.Type
        fieldKind := fieldType.Kind()
        if fieldKind != reflect.String {
            return false, fmt.Errorf("tagged field is of kind: %s, but should be string", fieldKind.String())
        }

        maxStrLen, err := strconv.Atoi(tagValue)
        if err != nil {
            return false, err
        }

        fieldValue := inputValue.Field(i)
        stringField := fieldValue.String()

        strLen := len(stringField)

        if strLen > maxStrLen {
            return false, nil
        }
    }
    

    return true, nil;
}

const tagMaxStrLen = "maxStrLen"

type Test struct {
    text string `maxStrLen:"5"`
}
