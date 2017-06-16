package main

import (
	"fmt"
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

func main(){
	values := []float64{12,23,123,451,126,634,3}
	fmt.Println(mean(values))
}