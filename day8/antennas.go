package day8

import (
	"aoc2024/utils"
	"fmt"
	"slices"
)

type Location struct {
	v, h int
}

type Mapping map[string][]Location

func GetAntennaLocations(matrix [][]string) Mapping {
	locations := make(map[string][]Location)
	for v, line := range matrix {
		for h, char := range line {
			if char != "." {
				locations[char] = append(locations[char], Location{v, h})
			}
		}
	}
	return locations
}

func FindAntinodes(matrix [][]string, mapping Mapping) []Location {

	visited := []Location{}
	for _, value := range mapping {
		for _, iloc := range value {
			for _, jloc := range value {
				if iloc != jloc {
					dv := iloc.v - jloc.v
					dh := iloc.h - jloc.h
					newLoc1 := Location{iloc.v + dv, iloc.h + dh}
					newLoc2 := Location{iloc.v - dv, iloc.h - dh}

					if utils.IsInside(newLoc1.v, newLoc1.h, len(matrix), len(matrix[0])) && !slices.Contains(visited, newLoc1) && !slices.Contains(value, newLoc1) {
						visited = append(visited, newLoc1)
					}
					if utils.IsInside(newLoc2.v, newLoc2.h, len(matrix), len(matrix[0])) && !slices.Contains(visited, newLoc2) && !slices.Contains(value, newLoc2) {
						visited = append(visited, newLoc1)
					}
				}
			}
		}
	}
	return visited
}

func removeDuplicates(visited []Location) []Location {
	allKeys := make(map[Location]bool)
	list := []Location{}
	for _, item := range visited {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func populate(matrix [][]string, visited []Location) {
	fmt.Println("---------------")
	for _, loc := range visited {
		matrix[loc.v][loc.h] = "#"
	}
	for _, line := range matrix {
		fmt.Println(line)
	}

}

func FindGridAntinodes(matrix [][]string, mapping Mapping) []Location {

	visited := []Location{}
	populate(matrix, visited)
	for _, value := range mapping {
		for i := 0; i < len(value); i++ {
			for j := i + 1; j < len(value); j++ {

				dv := value[i].v - value[j].v
				dh := value[i].h - value[j].h
				newLoc1 := Location{value[i].v + dv, value[i].h + dh}
				newLoc2 := Location{value[i].v - dv, value[i].h - dh}

				for utils.IsInside(newLoc1.v, newLoc1.h, len(matrix), len(matrix[0])) {
					if !slices.Contains(visited, newLoc1) && !slices.Contains(value, newLoc1) {
						visited = append(visited, newLoc1)
					}
					newLoc1 = Location{newLoc1.v + dv, newLoc1.h + dh}
				}
				for utils.IsInside(newLoc2.v, newLoc2.h, len(matrix), len(matrix[0])) {
					if !slices.Contains(visited, newLoc2) && !slices.Contains(value, newLoc2) {
						visited = append(visited, newLoc2)
					}
					newLoc2 = Location{newLoc2.v - dv, newLoc2.h - dh}
				}
			}
		}
	}
	populate(matrix, visited)
	return visited
}

func CountAntinodes(matrix [][]string, mapping Mapping) int {
	visited := FindGridAntinodes(matrix, mapping)
	for _, locations := range mapping {
		visited = append(visited, locations...)
	}
	return len(removeDuplicates(visited))
}
