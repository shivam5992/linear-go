package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinearRegression(t *testing.T) {

	pred_vars := []float64{1,2,4,3,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30}
	target := []float64{1,3,3,2,5,7,3,6,9,9,5,2,4,6,0,9,7,6,5,4,4,3,3,2,4,5,5,6,6,7}
	
	test_vars := []float64{5,6}
	actual := []float64{6,7}

	predicted := linear_regression(pred_vars, target, test_vars)
	fmt.Println(predicted)

	rmse := rmse(actual, predicted)
	fmt.Println(rmse)

	assert := assert.New(t)
	assert.InDelta(predicted, 3, -2, "yture")
	assert.InDelta(rmse, -0.262323, 10E-12, "")

}