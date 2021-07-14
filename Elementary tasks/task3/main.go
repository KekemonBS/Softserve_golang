package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type shape interface {
	area()
}

type triangle struct {
	name string
	s1   float64
	s2   float64
	s3   float64
}

func (t triangle) area() float64 {
	p := (t.s1 + t.s2 + t.s3) / 2
	return math.Sqrt(p * (p - t.s1) * (p - t.s2) * (p - t.s3))
}

type shapePointerList []*triangle

func (shl shapePointerList) Len() int           { return len(shl) }
func (shl shapePointerList) Less(i, j int) bool { return shl[i].area() > shl[j].area() }
func (shl shapePointerList) Swap(i, j int)      { shl[i], shl[j] = shl[j], shl[i] }

func main() {
	helpFlag := flag.Bool("help", false, `triangles help`)
	flag.Parse()
	if *helpFlag {
		fmt.Println(`Input format (separator - comma):
		<Name>, <side length>, <side length>, <side length>`)
		os.Exit(1)
	}

	triangles := make(shapePointerList, 8, 16)
	triangles = fillTriangles()
	sort.Sort(triangles)
	printSortedShapes(triangles, "Triangle")
}

func fillTriangles() shapePointerList {
	fmt.Printf("\nPlease fill in triangles: \n")

	triangles := shapePointerList{}

	scanner := bufio.NewScanner(os.Stdin)
	for true {

		fmt.Print("Fill in triangle : ")
		scanner.Scan()
		text := scanner.Text()
		desc := strings.Split(strings.Replace(text, " ", "", -1), ",")

		name := desc[0]
		s1, _ := strconv.ParseFloat(desc[1], 32)
		s2, _ := strconv.ParseFloat(desc[2], 32)
		s3, _ := strconv.ParseFloat(desc[3], 32)

		triangles = append(triangles, &triangle{name, s1, s2, s3})

		fmt.Print("Continue ? [Y/n] : ")
		scanner.Scan()
		ans := scanner.Text()
		if ans == "Y" || ans == "y" {
			continue
		} else if ans == "N" || ans == "n" {
			break
		} else {
			fmt.Println("Wrong letter")
		}
	}
	return triangles
}

func printSortedShapes(sh shapePointerList, name string) {
	fmt.Printf("\n============= %s list: ===============\n", name)
	for i, tr := range sh {
		fmt.Printf("%d. [%s %s]: %.2f cm^2\n", i+1, name, tr.name, tr.area())
	}
}
