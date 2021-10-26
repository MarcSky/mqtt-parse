package main

import "testing"

func TestReal(t *testing.T) {
	type test struct {
		subscribe string
		topic     string
		result    bool
	}

	var tests = []test{
		{
			"leaf2/leaf1/leaf3/leaf4/leaf5/+/leaf7/#", "leaf7", false,
		},
		{
			"leaf", "leaf/leaf2/leaf3", false,
		},
		{
			"leaf2/+/leaf3", "leaf2/hidden/leaf3", true,
		},
		{
			"leaf2/+/leaf3/#", "leaf2/hidden/leaf3/leaf4/leaf5", true,
		},
		{
			"leaf2/leaf5/#", "leaf2/leaf5", false,
		},
		{
			"leaf2/leaf5/#", "leaf2/leaf5/leaf7/leaf9", true,
		},
		{
			"#", "leaf2/leaf5", true,
		},
		{
			"", "leaf2/leaf5", false,
		},
		{
			"leaf2/leaf5", "", false,
		},
		{
			"leaf2+/leaf5", "leaf2more/leaf5", true,
		},
		{
			"+leaf2/leaf5", "moreleaf2/leaf5", true,
		},
		{
			"leaf2/+leaf1/leaf3/leaf4/leaf5+/+/leaf7/#", "leaf2/testleaf1/leaf3/leaf4/leaf5more/data/leaf7/caramba", true,
		},
	}

	for i := range tests {
		r := parse(tests[i].subscribe, tests[i].topic)
		if r != tests[i].result {
			t.Errorf("problem with test %d. We expected %v but got %v", i, tests[i].result, r)
		}
	}
}
