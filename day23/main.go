package main

import (
	_ "embed"
	"log/slog"
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

type link [2]string

type solver struct {
	conns []link
	adj   map[string]map[string]bool
}

func (s *solver) prep(inp string) *solver {
	lines := strings.Split(inp, "\n")
	s.conns = make([]link, len(lines))
	s.adj = make(map[string]map[string]bool)

	for i, line := range lines {
		if line == "" {
			continue
		}
		a, b, _ := strings.Cut(line, "-")
		s.conns[i] = [2]string{a, b}

		// Initialize maps if they don't exist
		if s.adj[a] == nil {
			s.adj[a] = make(map[string]bool)
		}
		if s.adj[b] == nil {
			s.adj[b] = make(map[string]bool)
		}

		// Add bidirectional connections
		s.adj[a][b] = true
		s.adj[b][a] = true
	}

	return s
}

func (s *solver) part1() int {
	// Get all unique computers
	computers := make([]string, 0, len(s.adj))
	for comp := range s.adj {
		computers = append(computers, comp)
	}

	count := 0
	// Try all possible combinations of 3 computers
	for i := 0; i < len(computers); i++ {
		for j := i + 1; j < len(computers); j++ {
			for k := j + 1; k < len(computers); k++ {
				comp1, comp2, comp3 := computers[i], computers[j], computers[k]

				// Check if they form a triangle (all connected to each other)
				if s.adj[comp1][comp2] && s.adj[comp2][comp3] && s.adj[comp1][comp3] {
					// Check if at least one computer starts with 't'
					if strings.HasPrefix(comp1, "t") ||
						strings.HasPrefix(comp2, "t") ||
						strings.HasPrefix(comp3, "t") {
						count++
					}
				}
			}
		}
	}

	return count
}

func (s *solver) isClique(computers []string) bool {
	for i := 0; i < len(computers); i++ {
		for j := i + 1; j < len(computers); j++ {
			if !s.adj[computers[i]][computers[j]] {
				return false
			}
		}
	}
	return true
}

func (s *solver) findMaxClique(current []string, candidates map[string]bool) []string {
	if len(candidates) == 0 {
		return current
	}

	maxClique := current

	// Try each candidate
	for candidate := range candidates {
		// Find new candidates that are connected to all current nodes
		newCandidates := make(map[string]bool)
		for next := range candidates {
			if next <= candidate { // Skip already processed nodes
				continue
			}
			isConnected := true
			for _, node := range current {
				if !s.adj[next][node] {
					isConnected = false
					break
				}
			}
			if isConnected && s.adj[next][candidate] {
				newCandidates[next] = true
			}
		}

		// Add candidate to current set and recurse
		newCurrent := append(append([]string{}, current...), candidate)
		result := s.findMaxClique(newCurrent, newCandidates)

		if len(result) > len(maxClique) {
			maxClique = result
		}
	}

	return maxClique
}

func (s *solver) part2() string {
	// Get all unique computers
	candidates := make(map[string]bool)
	for comp := range s.adj {
		candidates[comp] = true
	}

	// Find the maximum clique
	maxClique := s.findMaxClique([]string{}, candidates)

	// Sort the result alphabetically
	for i := 0; i < len(maxClique)-1; i++ {
		for j := i + 1; j < len(maxClique); j++ {
			if maxClique[i] > maxClique[j] {
				maxClique[i], maxClique[j] = maxClique[j], maxClique[i]
			}
		}
	}

	return strings.Join(maxClique, ",")
}
