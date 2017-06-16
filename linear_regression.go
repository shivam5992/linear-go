package main

import (
	"fmt"
	"math"
)


func mean(values []float64) float64{
	sum := 0.0
	for _, v := range values{
		sum += v
	}
	return sum / float64(len(values))
}


func covariance(x []float64, mean_x float64, y []float64, mean_y float64) float64{
	covar := 0.0

	i := 0
	for _, x_val := range(x){
		covar_prod := (x_val - mean_x) * (y[i] - mean_y)
		covar += covar_prod
		i += 1
	}
	return covar
}


func variance(values []float64, mean_value float64) float64{
	variance_sum := 0.0
	for _, v := range values{
		abs := v - mean_value
		true_abs := abs*abs
		variance_sum += true_abs
	}
	return variance_sum
}




func rmse(actual []float64, predicted []float64) float64{
	sum_error := 0.0
	i := 0
	for _, value := range actual{
		err := predicted[i] - value
		sum_error += err*err
		i += 1 
	}
	mean_error := sum_error / float64(len(actual))
	return math.Sqrt(mean_error)
}

func coefficients(pred_vars []float64, target []float64) []float64 {
	x_mean := mean(pred_vars)
	y_mean := mean(target)

	b1 := covariance(pred_vars, x_mean, target, y_mean) / variance(pred_vars, x_mean)
	b0 := y_mean - (b1 * x_mean)

	coff := []float64{b0, b1}
	return coff
}

func linear_regression(pred_vars []float64, target []float64, test_vars [] float64) []float64{
	var predictions []float64

	coff := coefficients(pred_vars, target)
	for _, row := range test_vars{
		y_pred := coff[0] + coff[1] * row
		predictions = append(predictions, y_pred)
	}
	return predictions
}


func main(){
	pred_vars := []float64{1,2,4,3,5}
	target := []float64{1,3,3,2,5}
	
	test_vars := []float64{5,6}
	actual := []float64{6,7}

	predicted := linear_regression(pred_vars, target, test_vars)
	fmt.Println(predicted)

	rmse := rmse(actual, predicted)
	fmt.Println(rmse)
}