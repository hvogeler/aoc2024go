package cpu

import (
	"testing"
)

func Test_Adv(t *testing.T) {
	t.Run("Adv1", func(t *testing.T) {
		cpu := Cpu{
			regA: 42,
		}
		Adv(&cpu, 1)
		if cpu.regA != 21 {
			t.Errorf("Expected 21, got %d", cpu.regA)
		}
	})
	t.Run("Adv2", func(t *testing.T) {
		cpu := Cpu{
			regA: 41,
		}
		Adv(&cpu, 1)
		if cpu.regA != 20 {
			t.Errorf("Expected 20, got %d", cpu.regA)
		}
	})
	t.Run("Adv3", func(t *testing.T) {
		cpu := Cpu{
			regA: 41,
		}
		Adv(&cpu, 2)
		if cpu.regA != 10 {
			t.Errorf("Expected 10, got %d", cpu.regA)
		}
	})
	t.Run("Adv3mpx", func(t *testing.T) {
		cpu := Cpu{
			regA: 41,
		}
		OpExec[adv](&cpu, 2)
		if cpu.regA != 10 {
			t.Errorf("Expected 10, got %d", cpu.regA)
		}
	})
	t.Run("Bdv", func(t *testing.T) {
		cpu := Cpu{
			regA: 41,
		}
		OpExec[bdv](&cpu, 2)
		if cpu.regB != 10 {
			t.Errorf("Expected 10, got %d", cpu.regA)
		}
	})
	t.Run("Cdv", func(t *testing.T) {
		cpu := Cpu{
			regA: 41,
		}
		cpu.ExecInstr(NewInstruction(uint8(cdv), 2))
		if cpu.regC != 10 {
			t.Errorf("Expected 10, got %d", cpu.regA)
		}
	})
	t.Run("Bst", func(t *testing.T) {
		cpu := Cpu{
			regA: 41,
		}
		cpu.ExecInstr(NewInstruction(uint8(bst), 4))
		if cpu.regB != 1 {
			t.Errorf("Expected 1, got %d", cpu.regB)
		}
		cpu.ExecInstr(NewInstruction(uint8(bst), 3))
		if cpu.regB != 3 {
			t.Errorf("Expected 3, got %d", cpu.regB)
		}
	})
	t.Run("Bxl", func(t *testing.T) {
		cpu := Cpu{
			regB: 0b010101,
		}
		cpu.ExecInstr(NewInstruction(uint8(bxl), 0b100))
		if cpu.regB != 0b10001 {
			t.Errorf("Expected 0b10001, got %d", cpu.regB)
		}
		cpu.regB = 41
		cpu.ExecInstr(NewInstruction(uint8(bxl), 7))
		if cpu.regB != 46 {
			t.Errorf("Expected 46, got %d", cpu.regB)
		}
		cpu.regB = 538
		cpu.ExecInstr(NewInstruction(uint8(bxl), 6))
		if cpu.regB != 540 {
			t.Errorf("Expected 540, got %d", cpu.regB)
		}
	})
	t.Run("Bxc", func(t *testing.T) {
		cpu := Cpu{
			regB: 0b010101,
			regC: 4,
		}
		cpu.ExecInstr(NewInstruction(uint8(bxc), 255))
		if cpu.regB != 0b10001 {
			t.Errorf("Expected 0b10001, got %d", cpu.regB)
		}
		cpu.regB = 41
		cpu.regC = 7
		cpu.ExecInstr(NewInstruction(uint8(bxc), 255))
		if cpu.regB != 46 {
			t.Errorf("Expected 46, got %d", cpu.regB)
		}
		cpu.regB = 538
		cpu.regC = 6
		cpu.ExecInstr(NewInstruction(uint8(bxc), 255))
		if cpu.regB != 540 {
			t.Errorf("Expected 540, got %d", cpu.regB)
		}
	})
}
