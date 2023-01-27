/* Copyright (C) 2016 Krešimir Nesek
 *
 * This software may be modified and distributed under the terms
 * of the MIT license. See the LICENSE file for details.
 */
package main

import (
	"testing"

	"github.com/solidgate-tech/elktail/testutils"
)

func TestResolveField(t *testing.T) {
	model1 := map[string]interface{}{
		"@timestamp": 3711,
		"message":    2138,
		"map": map[string]interface{}{
			"test": "test",
		},
	}
	testutils.AssertEqualsString(t, "2138", eval(model1, "message"))
	testutils.AssertEqualsString(t, "test", eval(model1, "map.test"))
	testutils.AssertEqualsString(t, "", eval(model1, "map.foo"))
	testutils.AssertEqualsString(t, "", eval(model1, "bar"))
}

func eval(model interface{}, expr string) string {
	result, _ := EvaluateExpression(model, expr)
	return result
}
