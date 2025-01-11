package main

import (
	_ "embed"
	"log/slog"
	"sort"
	"strconv"
	"strings"
	"time"
)

//go:embed inp
var inp string

func main() {
	slog.Info("Starting day")

	start := time.Now()
	got1 := new(solver).prep(inp).part1()
	slog.Info("Part 1", "time", time.Since(start).String(), "Ans", got1)

	start2 := time.Now()
	got2 := new(solver).prep(inp).part2()
	slog.Info("Part 2", "time", time.Since(start2).String(), "Ans", got2)
}

type gate struct {
	op  string
	in1 string
	in2 string
	out string
}

type solver struct {
	initialValues map[string]int
	gates         []gate
	wireValues    map[string]int
}

func (s *solver) prep(inp string) *solver {
	s.initialValues = make(map[string]int)
	s.gates = make([]gate, 0)
	s.wireValues = make(map[string]int)

	lines := strings.Split(strings.TrimSpace(inp), "\n")
	parsingGates := false

	for _, line := range lines {
		if line == "" {
			parsingGates = true
			continue
		}

		if !parsingGates {
			// Parse initial values
			parts := strings.Split(line, ": ")
			wire := parts[0]
			value := 0
			if parts[1] == "1" {
				value = 1
			}
			s.initialValues[wire] = value
			s.wireValues[wire] = value
		} else {
			// Parse gates
			parts := strings.Split(line, " -> ")
			outWire := parts[1]
			inParts := strings.Split(parts[0], " ")

			if len(inParts) == 3 {
				s.gates = append(s.gates, gate{
					op:  inParts[1],
					in1: inParts[0],
					in2: inParts[2],
					out: outWire,
				})
			}
		}
	}

	return s
}

func (s *solver) simulateGate(g gate) bool {
	val1, ok1 := s.wireValues[g.in1]
	val2, ok2 := s.wireValues[g.in2]

	if !ok1 || !ok2 {
		return false
	}

	var result int
	switch g.op {
	case "AND":
		result = val1 & val2
	case "OR":
		result = val1 | val2
	case "XOR":
		result = val1 ^ val2
	}

	s.wireValues[g.out] = result
	slog.Info("[DEBUG_LOG] Gate operation",
		"gate", g.op,
		"in1", g.in1,
		"val1", val1,
		"in2", g.in2,
		"val2", val2,
		"out", g.out,
		"result", result)
	return true
}

func (s *solver) simulate() {
	// Keep simulating until no more changes
	for {
		progress := false
		for _, g := range s.gates {
			if _, exists := s.wireValues[g.out]; !exists {
				if s.simulateGate(g) {
					progress = true
				}
			}
		}
		if !progress {
			break
		}
	}
}

func (s *solver) part1() int {
	s.simulate()
	return s.getZNumber()
}

func (s *solver) getBinaryNumber(prefix string) int {
	var wires []string
	for wire := range s.initialValues {
		if strings.HasPrefix(wire, prefix) {
			wires = append(wires, wire)
		}
	}

	// Sort wires by their number in ascending order (LSB to MSB)
	sort.Slice(wires, func(i, j int) bool {
		ni, _ := strconv.Atoi(strings.TrimPrefix(wires[i], prefix))
		nj, _ := strconv.Atoi(strings.TrimPrefix(wires[j], prefix))
		return ni < nj
	})

	// Build binary number from LSB to MSB
	result := 0
	for _, wire := range wires {
		bitPos, _ := strconv.Atoi(strings.TrimPrefix(wire, prefix))
		val := s.initialValues[wire]
		result = result | (val << bitPos)
		slog.Info("[DEBUG_LOG] Building number",
			"prefix", prefix,
			"wire", wire,
			"bitPos", bitPos,
			"val", val,
			"current_result", result)
	}

	return result
}

func (s *solver) getZNumber() int {
	var zWires []string
	for wire := range s.wireValues {
		if strings.HasPrefix(wire, "z") {
			zWires = append(zWires, wire)
		}
	}

	// Sort z wires by their number
	sort.Slice(zWires, func(i, j int) bool {
		ni, _ := strconv.Atoi(strings.TrimPrefix(zWires[i], "z"))
		nj, _ := strconv.Atoi(strings.TrimPrefix(zWires[j], "z"))
		return ni < nj
	})

	// Build binary number from LSB to MSB
	result := 0
	for _, wire := range zWires {
		bitPos, _ := strconv.Atoi(strings.TrimPrefix(wire, "z"))
		val := s.wireValues[wire]
		result = result | (val << bitPos)
		slog.Info("[DEBUG_LOG] Building z number",
			"wire", wire,
			"bitPos", bitPos,
			"val", val,
			"current_result", result)
	}

	return result
}

func (s *solver) part2() string {
	return ""
}
