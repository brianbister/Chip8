package chip8

import (
	"testing"
)

func TestOpcodeANNN(t *testing.T) {
	c8 := NewChip8()
	c8.Memory[0x200] = 0xA1
	c8.Memory[0x201] = 0x23
	c8.Memory[0x202] = 0xA0
	c8.Memory[0x203] = 0x00
	
	c8.EmulateCycle()
	if c8.IndexRegister != 0x0123 {
		t.Logf("%X != 0x0123", c8.IndexRegister)
		t.Fail()
	}
	c8.EmulateCycle()
	if c8.IndexRegister != 0x0000 {
		t.Logf("%X != 0x0000", c8.IndexRegister)
		t.Fail()
	}
}

func TestOpcode8XY0(t *testing.T) {
	c8 := NewChip8()
	c8.V[1] = 0x05
	c8.V[5] = 0x10
	c8.Memory[0x200] = 0x80
	c8.Memory[0x201] = 0x10
	c8.Memory[0x202] = 0x85
	c8.Memory[0x203] = 0x60
	
	c8.EmulateCycle()
	if c8.V[0] != 0x05 {
		t.Logf("%X != 0x05", c8.V[0])
		t.Fail()
	}
	c8.EmulateCycle()
	if c8.V[5] != 0x10 {
		t.Logf("%X != 0x10", c8.V[5])
		t.Fail()
	}
}