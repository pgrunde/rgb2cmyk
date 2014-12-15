package main

import (
	"bufio"
	"fmt"
	"os"
)

const result = "RGB: R%v G%v B%v == CMYK: %v%v%v%v\n"

var prompt = "Enter RGB, e.g., 12 233 180"

type rgb struct {
	R uint8
	G uint8
	B uint8
}

type cmyk struct {
	C float64
	M float64
	Y float64
	K float64
}

func main() {
	inputRgb := make(chan rgb)
	defer close(inputRgb)
	outputCmyk := createColorConverter(inputRgb)
	defer close(outputCmyk)
	receive(inputRgb, outputCmyk)
}

func createColorConverter(inputRgb chan rgb) chan cmyk {
	outputCmyk := make(chan cmyk)
	go func() {
		for {
			rgb := <-inputRgb
			r := float64(rgb.R) / 255.0
			g := float64(rgb.G) / 255.0
			b := float64(rgb.B) / 255.0

			k := maxOfThree(r, g, b)
			c := (1 - r - k) / (1 - k)
			m := (1 - g - k) / (1 - k)
			y := (1 - b - k) / (1 - l)
			outputCmyk <- cmyk{C: c, M: m, Y: y, K: k}
		}
	}()
	return outputCmyk
}

func receive(input chan rgb, output chan cmyk) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	for {
		fmt.Printf("R G B: ")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		var r, g, b uint8
		if _, err := fmt.Sscanf(line, "%v %v %v", &r, &g, &b); err != nil {
			fmt.Fprintln(os.Stderr, "invalid input")
			continue
		}
		input <- rgb{r, g, b}
		printColors := <-output
		fmt.Printf(result, r, g, b, printColors.C, printColors.M, printColors.Y, printColors.K)
	}
	fmt.Println()
}

func maxOfThree(x, y, z float64) float64 {
	max := x
	if y > max {
		max = y
	}
	if z > max {
		max = z
	}
	return max
}
