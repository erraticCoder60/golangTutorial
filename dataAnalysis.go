package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"

	"github.com/go-gota/gota/dataframe"
	"github.com/wcharczuk/go-chart"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat"
)

func main() {
	// Step 1: Reading Data from CSV
	file, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Print the first few rows of the dataframe
	df := dataframe.ReadCSV(file)
	fmt.Println(df)

	// Step 2: Data Transformation
	// Assume the data has a column "value" which we want to filter and convert to float64
	values := df.Col("value").Float()

	// Filter out missing or invalid data
	var cleanValues []float64
	for _, v := range values {
		if !isNaN(v) {
			cleanValues = append(cleanValues, v)
		}
	}

	// Step 3: Statistical Analysis using Gonum
	mean := stat.Mean(cleanValues, nil)
	stdDev := stat.StdDev(cleanValues, nil)
	// Sort the data
	sort.Float64s(cleanValues)
	median := stat.Quantile(0.5, stat.Empirical, cleanValues, nil)
	min, max := minMax(cleanValues)
	fmt.Printf("Mean: %v, Standard Deviation: %v, Median: %v, Min: %v, Max: %v\n",
		mean, stdDev, median, min, max)

	// Using Gonum matrix for additional operations
	dataMatrix := mat.NewDense(len(cleanValues), 1, cleanValues)
	fmt.Printf("Data Matrix:\n%v\n", mat.Formatted(dataMatrix, mat.Prefix(" "), mat.Excerpt(0)))

	// // Calculating Covariance Matrix (though with single variable it's just variance)
	// covMatrix := stat.CovarianceMatrix(nil, dataMatrix, nil)
	// fmt.Printf("Covariance Matrix:\n%v\n", mat.Formatted(covMatrix, mat.Prefix(" "), mat.Excerpt(0)))

	// Additional Gonum operations: Calculating correlation
	var corr float64
	if len(cleanValues) > 1 {
		corr = stat.Correlation(cleanValues, cleanValues, nil)
	}
	fmt.Printf("Correlation: %v\n", corr)

	// Step 4: Data Visualization
	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				Name:    "Data Values",
				XValues: makeRange(len(cleanValues)),
				YValues: cleanValues,
			},
		},
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)

}

// Helper function to check for NaN values
func isNaN(f float64) bool {
	return math.IsNaN(f)
}

// Helper function to calculate min and max values
func minMax(data []float64) (float64, float64) {
	min, max := data[0], data[0]
	for _, v := range data {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

// Helper function to create a range of integers for X axis
func makeRange(length int) []float64 {
	r := make([]float64, length)
	for i := range r {
		r[i] = float64(i)
	}
	return r
}
