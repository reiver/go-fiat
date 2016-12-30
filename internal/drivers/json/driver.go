package verboten

import (
	"bytes"
	"encoding/json"
	"strings"
)

type internalDriver struct {}

func (receiver internalDriver) Eval(code string, ctx map[string][]interface{}) (interface{}, error) {

	m := map[string]interface{}{}

	fieldNameList := strings.Split(code, ",")
	fieldNames := map[string]struct{}{}
	for _, fieldName := range fieldNameList {
		fieldNames[fieldName] = struct{}{}
	}

	for k, vs := range ctx {

		if _, ok := fieldNames[k]; !ok {
			continue
		}

		if 0 >= len(vs) {

		}
		v := vs[0]

		m[k] = v
	}


	var buffer bytes.Buffer

	jsonEncoder := json.NewEncoder(&buffer)

	if err := jsonEncoder.Encode(&m); nil != err {
		return nil, err
	}

	s := buffer.String()

	return s, nil
}
