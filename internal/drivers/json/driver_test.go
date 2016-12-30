package verboten

import (
	"testing"
)

func TestInternalDriverEval(t *testing.T) {

	tests := []struct{
		Code string
		CTX map[string][]interface{}
		Expected string
	}{
		{
			Code: "",
			CTX: map[string][]interface{}{
				"apple":  []interface{}{"ONE"},
				"banana": []interface{}{"TWO"},
				"cherry": []interface{}{"THREE"},
				"grape":  []interface{}{"FOUR"},
			},
			Expected: `{}`+"\n",
		},



		{
			Code: "apple",
			CTX: map[string][]interface{}{
				"apple":  []interface{}{"ONE"},
				"banana": []interface{}{"TWO"},
				"cherry": []interface{}{"THREE"},
				"grape":  []interface{}{"FOUR"},
			},
			Expected: `{"apple":"ONE"}`+"\n",
		},
		{
			Code: "banana",
			CTX: map[string][]interface{}{
				"apple":  []interface{}{"ONE"},
				"banana": []interface{}{"TWO"},
				"cherry": []interface{}{"THREE"},
				"grape":  []interface{}{"FOUR"},
			},
			Expected: `{"banana":"TWO"}`+"\n",
		},
		{
			Code: "cherry",
			CTX: map[string][]interface{}{
				"apple":  []interface{}{"ONE"},
				"banana": []interface{}{"TWO"},
				"cherry": []interface{}{"THREE"},
				"grape":  []interface{}{"FOUR"},
			},
			Expected: `{"cherry":"THREE"}`+"\n",
		},



		{
			Code: "apple,banana,cherry",
			CTX: map[string][]interface{}{
				"apple":  []interface{}{"ONE"},
				"banana": []interface{}{"TWO"},
				"cherry": []interface{}{"THREE"},
				"grape":  []interface{}{"FOUR"},
			},
			Expected: `{"apple":"ONE","banana":"TWO","cherry":"THREE"}`+"\n",
		},
	}


	for testNumber, test := range tests {

		var driver internalDriver

		value, err := driver.Eval(test.Code, test.CTX)
		if nil != err {
			t.Errorf("For test #%d, did not expect to get an error, but actually got one: (%T) %v", testNumber)
			continue
		}

		if expected, actual := test.Expected, value; expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q.", testNumber, expected, actual)
			continue
		}
	}
}
