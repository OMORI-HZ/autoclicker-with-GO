package main

import (
    "fmt"
    "time"

    "github.com/go-vgo/robotgo"
)

func main() {
    var clickInterval float64
    var numClicks int

    // Get user input for the interval and number of clicks
    fmt.Print("Enter the interval between clicks (in seconds): ")
    fmt.Scanf("%f", &clickInterval)

    fmt.Print("Enter the number of clicks: ")
    fmt.Scanf("%d", &numClicks)

    fmt.Println("Autoclicker will start in 5 seconds...")
    time.Sleep(5 * time.Second)

    // Convert the interval to a duration
    interval := time.Duration(clickInterval * float64(time.Second))

    for i := 0; i < numClicks; i++ {
        robotgo.Click()
        time.Sleep(interval)
    }

    fmt.Println("Autoclicking completed.")
}
