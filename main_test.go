package main

import (
	"strings"
	"testing"
)

func TestReal(t *testing.T) {
	type test struct {
		subscribe string
		topic     string
		result    bool
	}

	var tests = []test{
		{
			"tp/binance/trading/lea#f2/+leaf1/leaf3/leaf4/leaf5+/+/leaf7/leaf8",
			"tp/binance/trading/leaf2/AAleaf1/leaf3/leaf4/leaf5BB/leaf6/leaf7/CCC/DDD",
			false,
		},
		{
			"tp/binance/trading/leaf2/+leaf1/leaf3/leaf4/leaf5+/+/leaf7/CCC/DDD",
			"tp/binance/trading/leaf2/AAleaf1/leaf3/leaf4/leaf5BB/leaf6/leaf7/CCC/DDD",
			true,
		},
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
			"leaf2/+/leaf3/#/leaf4", "leaf2/hidden/leaf3/leaf4/leaf5", false,
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
		r := parse(strings.Split(tests[i].subscribe, "/"), strings.Split(tests[i].topic, "/"))
		if r != tests[i].result {
			t.Errorf("problem with test %d. We expected %v but got %v", i, tests[i].result, r)
		}
	}
}
