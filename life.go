package main

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"strconv"
	"time"
)

func initialize(width int, height int) [][]bool {
	frame := [][]bool{}
	for row := 0; row < height; row++ {
		frame = append(frame, []bool{})
		for col := 0; col < width; col++ {
			frame[row] = append(frame[row], false)
		}
	}
	for row := 1; row < (height - 2); row++ {
		for col := 1; col < (width - 1); col++ {
			var num uint8
			binary.Read(rand.Reader, binary.LittleEndian, &num)
			cellState := num%2 == 0
			frame[row][col] = cellState
		}
	}
	return frame
}

func printFrame(frame [][]bool, gen int) {
	var buffer bytes.Buffer
	buffer.WriteString("\033[42m")
	for i := 0; i < 40; i++ {
		buffer.WriteString("\n")
	}
	buffer.WriteString(strconv.Itoa(gen))
	buffer.WriteString("\n")
	for row := 0; row < len(frame); row++ {
		for col := 0; col < len(frame[row]); col++ {
			if frame[row][col] {
				buffer.WriteString("█")
			} else {
				buffer.WriteString(" ")
			}
		}
		buffer.WriteString("\n")
	}
	fmt.Print(buffer.String())
	fmt.Println("\033[0m")
}

func generate(source [][]bool, target [][]bool, gen *int) [][]bool {
	for row := 1; row < len(source)-2; row++ {
		for col := 1; col < len(source[row])-1; col++ {
			liveNeighbors := 0
			if source[row-1][col-1] {
				liveNeighbors += 1
			}
			if source[row][col-1] {
				liveNeighbors += 1
			}
			if source[row+1][col-1] {
				liveNeighbors += 1
			}
			if source[row-1][col] {
				liveNeighbors += 1
			}
			if source[row+1][col] {
				liveNeighbors += 1
			}
			if source[row-1][col+1] {
				liveNeighbors += 1
			}
			if source[row][col+1] {
				liveNeighbors += 1
			}
			if source[row+1][col+1] {
				liveNeighbors += 1
			}
			target[row][col] = false
			if source[row][col] {
				if liveNeighbors == 2 || liveNeighbors == 3 {
					target[row][col] = true
				}
			} else {
				if liveNeighbors == 3 {
					target[row][col] = true
				}
			}
		}
	}
	*gen++
	return target
}

func main() {
	width := 80
	height := 40
	gen := 0
	frame1 := initialize(width, height)
	frame2 := initialize(width, height)
	printFrame(frame1, gen)
	for {
		if gen%2 == 0 {
			generate(frame1, frame2, &gen)
			printFrame(frame2, gen)
		} else {
			generate(frame2, frame1, &gen)
			printFrame(frame1, gen)
		}
		time.Sleep(150 * time.Millisecond)
	}
}
