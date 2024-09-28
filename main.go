package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func getMeasurement() ([]float64, error) {
	cmd := exec.Command("sudo", "kubectl", "exec", "nginx", "--", "curl", "http://deviceshifu-plate-reader.deviceshifu.svc.cluster.local/get_measurement")

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("error: %s: %s", err, stderr.String())
	}

	body := out.String()

	// 解析返回值为浮点数切片
	var measurements []float64
	for _, s := range strings.Fields(body) {
		value, err := strconv.ParseFloat(s, 64)
		if err == nil {
			measurements = append(measurements, value)
		}
	}
	return measurements, nil
}

func main() {
	ticker := time.NewTicker(10 * time.Second) // 默认每10秒请求一次
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			measurements, err := getMeasurement()
			if err != nil {
				fmt.Printf("Error fetching data: %v\n", err)
				continue
			}

			// 计算平均值
			var sum float64
			fmt.Println(measurements)
			for _, m := range measurements {
				sum += m
			}
			if len(measurements) > 0 {
				avg := sum / float64(len(measurements))
				fmt.Printf("Average measurement: %.2f\n", avg)
			} else {
				fmt.Println("No measurements found.")
			}
		}
	}
}
