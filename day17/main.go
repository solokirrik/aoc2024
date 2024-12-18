package main

import (
	_ "embed"
	"errors"
	"log/slog"
	"math"
	"strconv"
	"strings"

	"github.com/solokirrik/aoc2024/pkg"
)

//go:embed inp
var inp string

func main() {
	slog.Info("Part 1", "Ans", new(solver).prep(inp).part1())
	slog.Info("Part 2", "Ans", new(solver).prep(inp).part2())
}

type solver struct {
	regA  uint64
	regB  uint64
	regC  uint64
	instr []uint64
}

func (s *solver) prep(inp string) *solver {
	parts := strings.Split(inp, "\n\n")

	regs := strings.Split(parts[0], "\n")
	s.regA, _ = strconv.ParseUint(strings.Split(regs[0], ": ")[1], 10, 64)
	s.regB, _ = strconv.ParseUint(strings.Split(regs[1], ": ")[1], 10, 64)
	s.regC, _ = strconv.ParseUint(strings.Split(regs[2], ": ")[1], 10, 64)

	instr := strings.Split(strings.Split(parts[1], ": ")[1], ",")
	s.instr = make([]uint64, 0, len(instr))
	for i := range instr {
		val, err := strconv.ParseUint(instr[i], 10, 64)
		pkg.PanicOnErr(err)
		s.instr = append(s.instr, val)
	}

	return s
}

func (s *solver) part1() string {
	instrPtr := 0
	var out []string
	var newOut string

	for instrPtr != -1 && instrPtr+1 < len(s.instr) {
		opcode := s.instr[instrPtr]
		comboOperandPtr := instrPtr + 1
		comboOperand := s.instr[comboOperandPtr]

		switch opcode {
		case 0:
			instrPtr = s.adv(comboOperandPtr, instrPtr, comboOperand)
		case 1:
			instrPtr = s.bxl(comboOperandPtr, instrPtr, comboOperand)
		case 2:
			instrPtr = s.bst(comboOperandPtr, instrPtr, comboOperand)
		case 3:
			instrPtr = s.jnz(instrPtr, comboOperand)
		case 4:
			instrPtr = s.bxc(instrPtr)
		case 5:
			instrPtr, newOut = s.out(comboOperandPtr, instrPtr, comboOperand)
			if instrPtr != -1 {
				out = append(out, newOut)
			}
		case 6:
			instrPtr = s.bdv(comboOperandPtr, instrPtr, comboOperand)
		case 7:
			instrPtr = s.cdv(comboOperandPtr, instrPtr, comboOperand)
		}
	}

	return strings.Join(out, ",")
}

func (s *solver) part2() uint64 {
	return 42
}

var errNoOperand = errors.New("no operand")

func (s *solver) operandValue(op uint64) uint64 {
	/*
		Combo operands 0 through 3 represent literal values 0 through 3.
		Combo operand 4 represents the value of register A.
		Combo operand 5 represents the value of register B.
		Combo operand 6 represents the value of register C.
		Combo operand 7 is reserved and will not appear in valid programs.
	*/
	switch op {
	case 0, 1, 2, 3:
		return op
	case 4:
		return s.regA
	case 5:
		return s.regB
	case 6:
		return s.regC
	case 7:
		panic(errNoOperand)
	default:
		panic(errNoOperand)
	}
}

func (s *solver) adv(comboOperandPtr, instrPtr int, comboOperand uint64) (newinstrPtr int) {
	/*
		The adv instruction (opcode 0) performs division. The numerator is the value in the A register.
		The denominator is found by raising 2 to the power of the instruction's combo operand.
		(So, an operand of 2 would divide A by 4 (2^2); an operand of 5 would divide A by 2^B.)
		The result of the division operation is truncated to an integer and then written to the A register.
	*/
	if comboOperandPtr >= len(s.instr) {
		return -1
	}
	pow := s.operandValue(comboOperand)
	numerator := float64(s.regA)
	denominator := math.Pow(2.0, float64(pow))

	s.regA = uint64(numerator / denominator)
	return instrPtr + 2
}

func (s *solver) bxl(comboOperandPtr, instrPtr int, comboOperand uint64) (newinstrPtr int) {
	/*
		The bxl instruction (opcode 1) calculates the bitwise XOR of register B and
		the instruction's literal operand, then stores the result in register B.
	*/
	if comboOperandPtr >= len(s.instr) {
		return -1
	}
	s.regB = s.regB ^ comboOperand
	return instrPtr + 2
}

func (s *solver) bst(comboOperandPtr, instrPtr int, comboOperand uint64) (newinstrPtr int) {
	/*
		The bst instruction (opcode 2) calculates the value of its combo operand modulo 8
		(thereby keeping only its lowest 3 bits), then writes that value to the B register.
	*/
	if comboOperandPtr >= len(s.instr) {
		return -1
	}
	s.regB = s.operandValue(comboOperand) % 8
	return instrPtr + 2
}

func (s *solver) jnz(instrPtr int, comboOperand uint64) (newinstrPtr int) {
	/*
		The jnz instruction (opcode 3) does nothing if the A register is 0.
		However, if the A register is not zero, it jumps by setting the instruction pointer to
		the value of its literal operand; if this instruction jumps, the instruction pointer is not
		increased by 2 after this instruction.
	*/
	if s.regA == 0 {
		return instrPtr + 2
	}
	return int(s.operandValue(comboOperand))
}

func (s *solver) bxc(instrPtr int) (newinstrPtr int) {
	/*
		The bxc instruction (opcode 4) calculates the bitwise XOR of register B and register C,
		then stores the result in register B. (For legacy reasons, this instruction reads an operand
		but ignores it.)
	*/
	s.regB = s.regB ^ s.regC
	return instrPtr + 2
}

func (s *solver) out(comboOperandPtr, instrPtr int, comboOperand uint64) (newinstrPtr int, out string) {
	/*
		The out instruction (opcode 5) calculates the value of its combo operand modulo 8,
		then outputs that value. (If a program outputs multiple values, they are separated by commas.)
	*/
	if comboOperandPtr >= len(s.instr) {
		return -1, ""
	}
	mod8 := s.operandValue(comboOperand) % 8

	return instrPtr + 2, strconv.FormatUint(mod8, 10)
}

func (s *solver) bdv(comboOperandPtr, instrPtr int, comboOperand uint64) (newinstrPtr int) {
	/*
		The bdv instruction (opcode 6) works exactly like the adv instruction except that
		the result is stored in the B register. (The numerator is still read from the A register.)
	*/
	if comboOperandPtr >= len(s.instr) {
		return -1
	}
	pow := s.operandValue(comboOperand)
	s.regB = uint64(math.Floor(float64(s.regA) / math.Pow(2.0, float64(pow))))
	return instrPtr + 2
}

func (s *solver) cdv(comboOperandPtr, instrPtr int, comboOperand uint64) (newinstrPtr int) {
	/*
		The cdv instruction (opcode 7) works exactly like the adv instruction except that the result
		is stored in the C register. (The numerator is still read from the A register.)
	*/
	if comboOperandPtr >= len(s.instr) {
		return -1
	}
	pow := s.operandValue(comboOperand)
	s.regC = uint64(math.Floor(float64(s.regA) / math.Pow(2.0, float64(pow))))
	return instrPtr + 2
}
