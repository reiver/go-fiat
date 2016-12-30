package fiat

import (
	"testing"
)

func TestExtractFollowingLines(t *testing.T) {

	tests := []struct{
		Tag string
		EndNeedle string
		Expected string
	}{
		{
			Tag: "apple:\"ONE\" banana:\"TWO\" cherry:\"THREE\"\n1\n2\n3\n",
			EndNeedle: "",
			Expected: "1\n2\n3\n",
		},
		{
			Tag: "apple:\"ONE\" banana:\"TWO\" cherry:\"THREE\"\r\n1\r\n2\r\n3\r\n",
			EndNeedle: "",
			Expected: "1\r\n2\r\n3\r\n",
		},



		{
			Tag: "apple:\"ONE\" banana:\"TWO\" cherry:\"THREE\"\n1\n2\n3\n<<<END>>>THERE IS MORE AFTER THIS!",
			EndNeedle: "<<<END>>>",
			Expected: "1\n2\n3\n",
		},
		{
			Tag: "apple:\"ONE\" banana:\"TWO\" cherry:\"THREE\"\r\n1\r\n2\r\n3\r\n<<<END>>>THERE IS MORE AFTER THIS!",
			EndNeedle: "<<<END>>>",
			Expected: "1\r\n2\r\n3\r\n",
		},



		{
			Tag: `apple:"ONE" banana:"TWO" cherry:"THREE"
1
2
3
`,
			EndNeedle: "",
			Expected: `1
2
3
`,
		},



		{
			Tag: `apple:"ONE" banana:"TWO" cherry:"THREE"
1
2
3
<<<END>>>
THERE IS MORE AFTER THIS!`,
			EndNeedle: "<<<END>>>",
			Expected: `1
2
3
`,
		},
	}


	for testNumber, test := range tests {

		actual, err := extractFollowingLines(test.Tag, test.EndNeedle)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}
		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q.", testNumber, expected, actual)
			continue
		}

	}
}
