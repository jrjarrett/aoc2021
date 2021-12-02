package application

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/jrjarrett/aoc2021/internal/day1"
	"github.com/jrjarrett/aoc2021/internal/day2"
	"log"
	"os"
	"strconv"
	"strings"
)

type App struct {
	Sonar            day1.Sonar
	NavigationSystem day2.NavigationSystem
}

func New() *App {
	application := App{Sonar: day1.Sonar{}}
	return &application
}

func (a *App) Run() {
	a.Day1_1()
	a.Day1_2()
	a.Day2_1()
}

func (a *App) Day1_1() {
	sonarSweep, err := getAOCSonarInput("testData/day1/challenge1.txt")
	if err != nil {
		log.Fatal(err)
	}
	result := a.Sonar.DepthIncrease(sonarSweep)
	fmt.Printf("Day 1.1 result is %#v\n", result)
}

func (a *App) Day1_2() {
	sonarSweep, err := getAOCSonarInput("testData/day1/challenge1.txt")
	if err != nil {
		log.Fatal(err)
	}
	result := a.Sonar.DepthIncrease(a.Sonar.CreateSlidingWindows(sonarSweep))
	fmt.Printf("Day 1.2 result is %#v\n", result)
}

func (a *App) Day2_1() {
	position, err := getAOCDirectionInput("testData/day2/day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	result, err := a.NavigationSystem.CalculatePosition(position)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Day 2.1 location is %#v. Answer is %d", result, result.Horizontal*result.Depth)
}

func getAOCSonarInput(fileName string) ([]int, error) {
	var sonarInput []int
	file, err := os.Open(fileName)
	if err != nil {
		return sonarInput, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.ParseInt(scanner.Text(), 10, 0)
		if err != nil {
			return sonarInput, err
		}
		sonarInput = append(sonarInput, int(i))
	}
	err = file.Close()
	if err != nil {
		return sonarInput, err
	}
	if scanner.Err() != nil {
		return sonarInput, scanner.Err()
	}
	return sonarInput, nil
}

func getAOCDirectionInput(fileName string) ([]day2.Navigation, error) {
	var line string
	var maneuvers []day2.Navigation
	file, err := os.Open(fileName)
	if err != nil {
		return maneuvers, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		if len(strings.Fields(line)) != 2 {
			return maneuvers, errors.New("Wrong input in file " + line)
		}
		direction := strings.Fields(line)[0]
		position, err := strconv.Atoi(strings.Fields(line)[1])
		if err != nil {
			return maneuvers, err
		}
		nav := day2.Navigation{
			Direction: direction,
			Position:  position,
		}
		maneuvers = append(maneuvers, nav)
	}
	return maneuvers, nil
}