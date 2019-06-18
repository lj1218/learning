// 12.7. Accessing Struct Field Tags
package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/search", search)
	_ = http.ListenAndServe(":8000", nil)
}

// search implements the /search URL endpoint.
func search(resp http.ResponseWriter, req *http.Request) {
	var data struct {
		Labels     []string `http:"l"`
		MaxResults int      `http:"max"`
		Exact      bool     `http:"x"`
	}
	data.MaxResults = 10 // set default
	if err := Unpack(req, &data); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest) // 400
		return
	}

	// ...rest of handler...
	_, _ = fmt.Fprintf(resp, "Search: %+v\n", data)
}

// Unpack populates the fields of the struct pointed to by ptr
// from the HTTP request parameters in req.
func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	// Build map of fields keyed by effective name.
	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = v.Field(i)
	}

	// update struct field for each parameter in the request.
	for name, values := range req.Form {
		f := fields[name]
		if !f.IsValid() {
			continue // ignore unrecognized HTTP parameters
		}
		for _, value := range values {
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}

// Output:
/*
http://localhost:8000/search
Search: {Labels:[] MaxResults:10 Exact:false}

http://localhost:8000/search?l=golang&l=programming
Search: {Labels:[golang programming] MaxResults:10 Exact:false}

http://localhost:8000/search?l=golang&l=programming&max=100
Search: {Labels:[golang programming] MaxResults:100 Exact:false}

http://localhost:8000/search?x=true&l=golang&l=programming
Search: {Labels:[golang programming] MaxResults:10 Exact:true}

http://localhost:8000/search?q=hello&x=123
x: strconv.ParseBool: parsing "123": invalid syntax

http://localhost:8000/search?q=hello&max=lots
max: strconv.ParseInt: parsing "lots": invalid syntax
*/
