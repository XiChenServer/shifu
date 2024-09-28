package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func fetchMeasurements() ([][]float64, error) {
	resp, err := http.Get("http://deviceshifu-plate-reader.deviceshifu.svc.cluster.local/get_measurement")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println("Response body:", string(body))

	// 处理空格分隔的数字
	lines := strings.Split(string(body), "\n")
	var measurements [][]float64

	for _, line := range lines {
		if line == "" {
			continue // 跳过空行
		}
		var row []float64
		values := strings.Fields(line)
		for _, value := range values {
			num, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return nil, err
			}
			row = append(row, num)
		}
		measurements = append(measurements, row)
	}
	return measurements, nil
}

func calculateAverage(measurements [][]float64) float64 {
	sum := 0.0
	count := 0

	for _, row := range measurements {
		for _, value := range row {
			sum += value
			count++
		}
	}

	if count == 0 {
		return 0
	}
	return sum / float64(count)
}

func main() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		measurements, err := fetchMeasurements()
		if err != nil {
			log.Println("Error fetching data:", err)
			continue
		}
		average := calculateAverage(measurements)
		fmt.Println("Average measurement:", average)
	}
}

//  docker build --tag measurement:v0.0.1 .
// sudo kind load docker-image measurement:v0.0.1
// sudo kubectl run measurement --image=measurement:v0.0.1
// sudo kubectl logs measurement -f
//  sudo kubectl delete pod measurement
