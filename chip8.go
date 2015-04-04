package chip8

import "fmt"

var chip8Fontset = [80]byte{
	// Extra spooky magic numbers
	0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
	0x20, 0x60, 0x20, 0x20, 0x70, // 1
	0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
	0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
	0x90, 0x90, 0xF0, 0x10, 0x10, // 4
	0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
	0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
	0xF0, 0x10, 0x20, 0x40, 0x40, // 7
	0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
	0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
	0xF0, 0x90, 0xF0, 0x90, 0x90, // A
	0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
	0xF0, 0x80, 0x80, 0x80, 0xF0, // C
	0xE0, 0x90, 0x90, 0x90, 0xE0, // D
	0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
	0xF0, 0x80, 0xF0, 0x80, 0x80, // F
}


type Chip8 struct {
	// Opcode
	opcode uint16

	// Memory
	// TODO: Make this private and create loadMemory func
	Memory [4096]byte

	// Registers (V0 to VF)
	v [16]byte

	// Index Register
	indexRegister uint16

	// Program Counter
	programCounter uint16

	// Graphics
	gfx [64 * 32]byte

	// Timers
	delayTimer byte
	soundTimer byte

	// Stack
	stack [16]uint16
	stackPointer uint16

	// Key Input
	key [16]byte
}

func NewChip8() *Chip8 {
	c8 := new(Chip8)
	// Go already initializes all number types to 0, so not much to initialize here.
	
	// Initialize pc and fontset
	c8.programCounter = 0x200
	for i := 0; i < 80; i++ {
		c8.Memory[i] = chip8Fontset[i]
	}
	return c8
}

func (c8* Chip8) EmulateCycle() {
	/*
	Fetch opcode
	Each opcode is two bytes, but since each section of memory is one byte we need to
	combine two sections of memory and increment the programCounter by two for each
	instruction.
	*/
	c8.opcode = uint16(c8.Memory[c8.programCounter]) << 8 | uint16(c8.Memory[c8.programCounter + 1])

	/*
	The following is an explanation for each opcode.  Taken from the Chip8 wikipedia page.

	0NNN	Calls RCA 1802 program at address NNN.
	00E0	Clears the screen.
	00EE	Returns from a subroutine.
	1NNN	Jumps to address NNN.
	2NNN	Calls subroutine at NNN.
	3XNN	Skips the next instruction if VX equals NN.
	4XNN	Skips the next instruction if VX doesn't equal NN.
	5XY0	Skips the next instruction if VX equals VY.
	6XNN	Sets VX to NN.
	7XNN	Adds NN to VX.
	8XY0	Sets VX to the value of VY.
	8XY1	Sets VX to VX or VY.
	8XY2	Sets VX to VX and VY.
	8XY3	Sets VX to VX xor VY.
	8XY4	Adds VY to VX. VF is set to 1 when there's a carry, and to 0 when there isn't.
	8XY5	VY is subtracted from VX. VF is set to 0 when there's a borrow, and 1 when there isn't.
	8XY6	Shifts VX right by one. VF is set to the value of the least significant bit of VX before the shift.
	8XY7	Sets VX to VY minus VX. VF is set to 0 when there's a borrow, and 1 when there isn't.
	8XYE	Shifts VX left by one. VF is set to the value of the most significant bit of VX before the shift.
	9XY0	Skips the next instruction if VX doesn't equal VY.
	ANNN	Sets I to the address NNN.
	BNNN	Jumps to the address NNN plus V0.
	CXNN	Sets VX to a random number, masked by NN.
	DXYN	Sprites stored in memory at location in index register (I), maximum 8bits wide. Wraps around the screen. If when drawn, clears a pixel, register VF is set to 1 otherwise it is zero. All drawing is XOR drawing (i.e. it toggles the screen pixels)
	EX9E	Skips the next instruction if the key stored in VX is pressed.
	EXA1	Skips the next instruction if the key stored in VX isn't pressed.
	FX07	Sets VX to the value of the delay timer.
	FX0A	A key press is awaited, and then stored in VX.
	FX15	Sets the delay timer to VX.
	FX18	Sets the sound timer to VX.
	FX1E	Adds VX to I.
	FX29	Sets I to the location of the sprite for the character in VX. Characters 0-F (in hexadecimal) are represented by a 4x5 font.
	FX33	Stores the Binary-coded decimal representation of VX, with the most significant of three digits at the address in I, the middle digit at I plus 1, and the least significant digit at I plus 2. (In other words, take the decimal representation of VX, place the hundreds digit in memory at location in I, the tens digit at location I+1, and the ones digit at location I+2.)
	FX55	Stores V0 to VX in memory starting at address I.
	FX65	Fills V0 to VX with values from memory starting at address I.
	*/
	switch c8.opcode & 0xF000 {
	case 0x0000:
		switch c8.opcode {
		case 0x00E0:
			//TODO
			//00E0	Clears the screen.
		case 0x00EE:
			//TODO
			//00EE	Returns from a subroutine.
		default:
			fmt.Printf("%X not implemeted\n", c8.opcode)
		}
	case 0x1000:
		//TODO
		//1NNN	Jumps to address NNN.
	case 0x2000:
		//TODO
		//2NNN	Calls subroutine at NNN.
	case 0x3000:
		//TODO
		//3XNN	Skips the next instruction if VX equals NN.
	case 0x4000:
		//TODO
		//4XNN	Skips the next instruction if VX doesn't equal NN.
	case 0x5000:
		switch c8.opcode & 0x000F {
		case 0x0000:
			//TODO
			//5XY0	Skips the next instruction if VX equals VY.
		default:
			fmt.Printf("%X not implemeted\n", c8.opcode)
		}
	case 0x6000:
		//TODO
		//6XNN	Sets VX to NN.
	case 0x7000:
		//TODO
		//7XNN	Adds NN to VX.
	case 0x8000:
		switch c8.opcode & 0x000F {
		case 0x0000:
			//TODO
			//8XY0	Sets VX to the value of VY.
		case 0x0001:
			//TODO
			//8XY1	Sets VX to VX or VY.
		case 0x0002:
			//TODO
			//8XY2	Sets VX to VX and VY.
		case 0x0003:
			//TODO
			//8XY3	Sets VX to VX xor VY.
		case 0x0004:
			//TODO
			//8XY4	Adds VY to VX. VF is set to 1 when there's a carry, and to 0 when there isn't.
		case 0x0005:
			//TODO
			//8XY5	VY is subtracted from VX. VF is set to 0 when there's a borrow, and 1 when there isn't.
		case 0x0006:
			//TODO
			//8XY6	Shifts VX right by one. VF is set to the value of the least significant bit of VX before the shift.
		case 0x0007:
			//TODO
			//8XY7	Sets VX to VY minus VX. VF is set to 0 when there's a borrow, and 1 when there isn't.
		case 0x000E:
			//TODO
			//8XYE	Shifts VX left by one. VF is set to the value of the most significant bit of VX before the shift.
		default:
			fmt.Printf("%X not implemeted\n", c8.opcode)
		}
	case 0x9000:
		switch c8.opcode & 0x000F {
		case 0x0000:
			//TODO
			//9XY0	Skips the next instruction if VX doesn't equal VY.
		default:
			fmt.Printf("%X not implemeted\n", c8.opcode)
		}
	case 0xA000:
		//TODO
		//ANNN	Sets I to the address NNN.
	case 0xB000:
		//TODO
		//BNNN	Jumps to the address NNN plus V0.
	case 0xC000:
		//TODO
		//CXNN	Sets VX to a random number, masked by NN.
	case 0xD000:
		//TODO
		//DXYN	Sprites stored in memory at location in index register (I), maximum 8bits wide. Wraps around the screen. If when drawn, clears a pixel, register VF is set to 1 otherwise it is zero. All drawing is XOR drawing (i.e. it toggles the screen pixels)
	case 0xE000:
		switch c8.opcode & 0x00FF {
		case 0x009E:
			//TODO
			//EX9E	Skips the next instruction if the key stored in VX is pressed.
		case 0x00A1:
			//TODO
			//EXA1	Skips the next instruction if the key stored in VX isn't pressed.
		default:
			fmt.Printf("%X not implemeted\n", c8.opcode)
		}
	case 0xF000:
		switch c8.opcode & 0x00FF {
		case 0x0007:
			//TODO
			//FX07	Sets VX to the value of the delay timer.
		case 0x00A:
			//TODO
			//FX0A	A key press is awaited, and then stored in VX.
		case 0x015:
			//TODO
			//FX15	Sets the delay timer to VX.
		case 0x0018:
			//TODO
			//FX18	Sets the sound timer to VX.
		case 0x001E:
			//TODO
			//FX1E	Adds VX to I.
		case 0x0029:
			//TODO
			//FX29	Sets I to the location of the sprite for the character in VX. Characters 0-F (in hexadecimal) are represented by a 4x5 font.
		case 0x0033
			//TODO
			//FX33	Stores the Binary-coded decimal representation of VX, with the most significant of three digits at the address in I, the middle digit at I plus 1, and the least significant digit at I plus 2. (In other words, take the decimal representation of VX, place the hundreds digit in memory at location in I, the tens digit at location I+1, and the ones digit at location I+2.)
		case 0x0055:
			//TODO
			//FX55	Stores V0 to VX in memory starting at address I.
		case 0x0065:
			//TODO
			//FX65	Fills V0 to VX with values from memory starting at address I.
		default:
			fmt.Printf("%X not implemeted\n", c8.opcode)
		}	
	default:	
		fmt.Printf("%X not implemeted\n", c8.opcode)
	}

}