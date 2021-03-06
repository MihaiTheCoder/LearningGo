package main

import (
    "fmt"
)

func main()  {
    plantCapacities := []float64 {30, 30, 30, 60, 60, 100}
    capacity := plantCapacities[0] + plantCapacities[1] + plantCapacities[2] + plantCapacities[3] +
    plantCapacities[4] + plantCapacities[5]

    gridLoad := 75.

    utilization:= gridLoad/capacity

    fmt.Printf("%-20s%.0f\n", "Capacity: ",capacity)
    fmt.Printf("%-20s%.0f\n","Load: ",gridLoad)
    fmt.Printf("%-20s%.1f%%\n","utilization: ",utilization * 100)
}