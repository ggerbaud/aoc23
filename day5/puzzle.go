package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

const day = "5"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	alm := toAlmanac(lines)
	total := part1(alm)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(alm)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(alm almanac) int {
	minLoc := -1
	for _, seed := range alm.seedList {
		location := transformSeed(seed, alm)
		if minLoc == -1 || location < minLoc {
			minLoc = location
		}
	}
	return minLoc
}

func part2(alm almanac) int {
	minLoc := 0
	for i := 0; true; i++ {
		cLoc := minLoc + i
		seed := transformSeedBack(cLoc, alm)
		if alm.seeds.isValid(seed) {
			return cLoc
		}
	}
	// unreachable
	return -1
}

func transformSeed(seed int, alm almanac) int {
	soil := alm.seedToSoilMap.toDestination(seed)
	fertilizer := alm.soilToFertilizerMap.toDestination(soil)
	water := alm.fertilizerToWater.toDestination(fertilizer)
	light := alm.waterToLight.toDestination(water)
	temperature := alm.lightToTemperature.toDestination(light)
	humidity := alm.temperatureToHumidity.toDestination(temperature)
	return alm.humidityToLocation.toDestination(humidity)
}

func transformSeedBack(location int, alm almanac) int {
	humidity := alm.humidityToLocation.toSource(location)
	temperature := alm.temperatureToHumidity.toSource(humidity)
	light := alm.lightToTemperature.toSource(temperature)
	water := alm.waterToLight.toSource(light)
	fertilizer := alm.fertilizerToWater.toSource(water)
	soil := alm.soilToFertilizerMap.toSource(fertilizer)
	return alm.seedToSoilMap.toSource(soil)
}

func toAlmanac(lines []string) almanac {
	alm := almanac{}
	for idx := 0; idx < len(lines); idx++ {
		line := lines[idx]
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "seeds:") {
			alm.seedList = utils.ListOfNumbers(strings.TrimPrefix(line, "seeds:"), " ")
			alm.seeds = make([]seedRange, 0)
			for i := 0; i < len(alm.seedList); i += 2 {
				alm.seeds = append(alm.seeds, seedRange{alm.seedList[i], alm.seedList[i+1]})
			}
			continue
		}
		if line == "seed-to-soil map:" {
			idx, alm.seedToSoilMap = handleMapper(lines, idx+1)
			continue
		}
		if line == "soil-to-fertilizer map:" {
			idx, alm.soilToFertilizerMap = handleMapper(lines, idx+1)
			continue
		}
		if line == "fertilizer-to-water map:" {
			idx, alm.fertilizerToWater = handleMapper(lines, idx+1)
			continue
		}
		if line == "water-to-light map:" {
			idx, alm.waterToLight = handleMapper(lines, idx+1)
			continue
		}
		if line == "light-to-temperature map:" {
			idx, alm.lightToTemperature = handleMapper(lines, idx+1)
			continue
		}
		if line == "temperature-to-humidity map:" {
			idx, alm.temperatureToHumidity = handleMapper(lines, idx+1)
			continue
		}
		if line == "humidity-to-location map:" {
			idx, alm.humidityToLocation = handleMapper(lines, idx+1)
			continue
		}
	}
	return alm
}

func handleMapper(lines []string, idx int) (int, std) {
	mappers := make([]mapper, 0)
	for j := idx; j < len(lines); j++ {
		line := lines[j]
		if line == "" {
			idx = j
			break
		}
		ranges := utils.ListOfNumbers(line, " ")
		mappers = append(mappers, mapper{ranges[0], ranges[1], ranges[2]})
	}
	return idx, mappers
}

func (sr seedRanges) isValid(v int) bool {
	for _, s := range sr {
		if s.IsValid(v) {
			return true
		}
	}
	return false
}

func (r seedRange) IsValid(v int) bool {
	return v >= r.start && v < r.start+r.length
}

func (s std) toDestination(source int) int {
	for _, m := range s {
		ok, d := m.toDestination(source)
		if ok {
			return d
		}
	}
	return source
}

func (s std) toSource(d int) int {
	if d == -1 {
		return -1
	}
	for _, m := range s {
		ok, s := m.toSource(d)
		if ok {
			return s
		}
	}
	return d
}

func (m mapper) toDestination(s int) (bool, int) {
	if s >= m.source && s < m.source+m.length {
		return true, s + m.destination - m.source
	}
	return false, -1
}

func (m mapper) toSource(d int) (bool, int) {
	if d >= m.destination && d < m.destination+m.length {
		return true, d + m.source - m.destination
	}
	return false, -1
}

type (
	almanac struct {
		seedList              []int
		seeds                 seedRanges
		seedToSoilMap         std
		soilToFertilizerMap   std
		fertilizerToWater     std
		waterToLight          std
		lightToTemperature    std
		temperatureToHumidity std
		humidityToLocation    std
	}

	seedRanges []seedRange

	seedRange struct {
		start  int
		length int
	}

	std []mapper

	mapper struct {
		destination int
		source      int
		length      int
	}
)
