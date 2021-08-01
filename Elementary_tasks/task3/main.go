package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type triangle struct {
	name string
	s1   float64
	s2   float64
	s3   float64
}

func (t triangle) area() float64 {
	p := (t.s1 + t.s2 + t.s3) / 2
	return math.Round(math.Sqrt(p*(p-t.s1)*(p-t.s2)*(p-t.s3))*100) / 100
}

type shapeList []triangle

func (shl shapeList) Len() int           { return len(shl) }
func (shl shapeList) Less(i, j int) bool { return shl[i].area() > shl[j].area() }
func (shl shapeList) Swap(i, j int)      { shl[i], shl[j] = shl[j], shl[i] }

func main() {
	err := Task3(flag.Bool("help", false, `triangles help`), "")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func Task3(hf *bool, proposed string) error {
	helpFlag := hf
	flag.Parse()
	if *helpFlag {
		fmt.Println(`Input format (separator - comma):
		<Name>, <side length>, <side length>, <side length>`)
		return nil
	}
	trianglesText := fillTriangles(proposed)
	triangles, err := convertShapes(trianglesText)
	if err != nil {
		return err
	}
	sortShapes(triangles)
	fmt.Println(printSortedShapes(triangles, "Triangle"))
	return nil
}

func fillTriangles(proposed string) string {
	if len(proposed) != 0 {
		return proposed
	}
	res := ""
	fmt.Printf("\nPlease fill in triangles: \n")
	scanner := bufio.NewScanner(os.Stdin)
	for {

		fmt.Print("Fill in triangle : ")
		scanner.Scan()
		text := scanner.Text()
		res += text + "\n"
		desc := strings.Split(strings.Replace(text, " ", "", -1), ",")
		if len(desc) != 4 {
			fmt.Println("Wrong formated input.\nShould be: <Name>, <side length>, <side length>, <side length>")
			os.Exit(1)
		}
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
	return res
}

func convertShapes(shapesString string) (shapeList, error) {
	triangles := shapeList{}
	stringSlice := strings.Split(shapesString, "\n")
	for _, v := range stringSlice[:len(stringSlice)-1] {
		desc := strings.Split(strings.Replace(v, " ", "", -1), ",")
		name := desc[0]
		s1, err := strconv.ParseFloat(desc[1], 32)
		if err != nil {
			return shapeList{}, errors.New("wrong first side entered")
		}
		s2, err := strconv.ParseFloat(desc[2], 32)
		if err != nil {
			return shapeList{}, errors.New("wrong second side entered")
		}
		s3, err := strconv.ParseFloat(desc[3], 32)
		if err != nil {
			return shapeList{}, errors.New("wrong third side entered")
		}

		if s1 >= s2+s3 || s2 >= s1+s3 || s3 >= s1+s2 {
			return shapeList{}, errors.New("the entered sides cannot be combined into a triangle")
		}
		triangles = append(triangles, triangle{name, s1, s2, s3})
	}
	return triangles, nil
}

func sortShapes(shl shapeList) shapeList {
	sort.Sort(shl)
	return shl
}

func printSortedShapes(sh shapeList, name string) string {
	res := ""
	res += fmt.Sprintf("\n============= %s list: ===============\n", name)
	for i, tr := range sh {
		res += fmt.Sprintf("%d. [%s %s]: %.2f cm^2\n", i+1, name, tr.name, tr.area())
	}
	return res
}
