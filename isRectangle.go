package main

import (
	"errors"
	"fmt"
)

type Point struct {
	x float32
	y float32
}

type Shape struct {
	points []Point
}
func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
func Cosine_of_the_angle_between_two_points(points []Point) (error, float32) {
	if len(points) != 2 {
		return errors.New("The points list must consist of two elements"), 1.1
	}
	return nil, (points[0].x*points[1].x + points[0].y*points[1].y) / (points[0].pythogram() * points[1].pythogram())
}
func max(list []float32) float32 {
	result := list[0]
	for _, element := range list {
		if element > result {
			result = element
		}
	}
	return result
}
func min(list []float32) float32 {
	result := list[0]
	for _, element := range list {
		if element < result {
			result = element
		}
	}
	return result
}
func pow(x float32, y int) float32 {
	result := float32(1)
	if y == 0 {
		return result
	}
	for i := 0; i < abs(y); i++ {
		result *= x
	}
	if y < 0 {
		return 1 / result
	} else {
		return result
	}
}
func square(x float32, n float32) float32 {
	return x*x - n
}
func squarePrime(x float32) float32 {
	return 2 * x
}

func sqrt(num float32) float32 {
	result := float32(1.0)
	for i := 0; i < 100000; i++ {
		result -= square(result, num) / squarePrime(num)
	}
	return result
}
func (p *Point) add(p2 Point) {
	(*p).x += (p2).x
	(*p).y += (p2).y
}
func (p *Point) sub(p2 Point) {
	(*p).x -= (p2).x
	(*p).y -= (p2).y
}
func (p *Point) pythogram() float32 {
	return sqrt(pow((*p).x, 2) + pow((*p).y, 2))
}

func (sh *Shape) getXCoordinates() []float32 {
	result := []float32{}
	for _, p := range (*sh).points {
		result = append(result, p.x)
	}
	return result
}
func (sh *Shape) getYCoordinates() []float32 {
	result := []float32{}
	for _, p := range (*sh).points {
		result = append(result, p.y)
	}
	return result
}
func (sh *Shape) getHighest() []Point {
	ys := (*sh).getYCoordinates()
	maxy := max(ys)
	result := []Point{}
	for _, p := range (*sh).points {
		if p.y == maxy {
			result = append(result, p)
		}
	}
	return result
}
func (sh *Shape) getLowest() []Point {
	ys := (*sh).getYCoordinates()
	miny := min(ys)
	result := []Point{}
	for _, p := range (*sh).points {
		if p.y == miny {
			result = append(result, p)
		}
	}
	return result
}
func (sh *Shape) getRighests() []Point {
	ys := (*sh).getXCoordinates()
	maxx := max(ys)
	result := []Point{}
	for _, p := range (*sh).points {
		if p.x == maxx {
			result = append(result, p)
		}
	}
	return result
}
func (sh *Shape) getLeftest() []Point {
	ys := (*sh).getXCoordinates()
	minx := min(ys)
	result := []Point{}
	for _, p := range (*sh).points {
		if p.x == minx {
			result = append(result, p)
		}
	}
	return result
}
func (sh *Shape) group_all_points() (Point, Point, []Point) {
	lowest := (*sh).getLowest()
	if len(lowest) != 1 {
		lowestShape := &Shape{lowest}
		lowest = lowestShape.getLeftest()
	}
	highest := (*sh).getHighest()
	if len(highest) != 1 {
		highestShape := &Shape{highest}
		highest = highestShape.getRighests()
	}
	PointList := []Point{}
	for _, p := range (*sh).points {
		if p != highest[0] && p != lowest[0] {
			PointList = append(PointList, p)
		}
	}
	return lowest[0], highest[0], PointList
}
func isRectangle(sh Shape) bool {
	lowest, highest, others := (&sh).group_all_points()
	lowestPointer := &lowest
	highestPointer := &highest
	Cosangles := []float32{}
	for i, _ := range others {
		(&others[i]).sub(lowest)
	}
	_, ang := Cosine_of_the_angle_between_two_points(others)
	Cosangles = append(Cosangles, ang)
	for i, _ := range others {
		(&others[i]).add(lowest)
		(&others[i]).sub(highest)
	}
	_, ang = Cosine_of_the_angle_between_two_points(others)
	Cosangles = append(Cosangles, ang)
	for i, _ := range others {
		(&others[i]).add(highest)
	}
	for _, p := range others {
		(lowestPointer).sub(p)
		(highestPointer).sub(p)
		_, ang = Cosine_of_the_angle_between_two_points([]Point{lowest, highest})
		Cosangles = append(Cosangles, ang)
		(lowestPointer).add(p)
		(highestPointer).add(p)
	}
	for _, i := range Cosangles {
		if i != 0 {
			return false
		}
	}
	return true
}

func main() {
	p1 := Point{0, 2}
	p2 := Point{1, 1}
	p3 := Point{3, 3}
	p4 := Point{2, 4}
	ps := Shape{[]Point{p1, p2, p3, p4}}
	fmt.Println(isRectangle(ps))
}
