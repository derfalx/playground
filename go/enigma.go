package main

import ("fmt"
        "strings")

// ==================================================================
// Sample project to learn the basics of Go
// Goal is to implement an Enigma machine to en-
// and decrypt text
// ==================================================================


// ==================================================================
//
// Debuging Stuff
//
// ==================================================================
const DEBUG = true
func println(param ...interface{})(n int, err error){
    if !DEBUG {
        return 0, nil
    }
    return fmt.Println(param...)
}

func printf(s string, param ...interface{})(n int, err error){
    if !DEBUG {
        return 0, nil
    }
    return fmt.Printf(s, param...)
}

// ==================================================================
//
// Default Rotor Configurations
//
// ==================================================================
const I = map[byte]byte{'A': 'E', 'B': 'K', 'C': 'M', 'D': 'F',
                        'E': 'L', 'F': 'G', 'G': 'D', 'H': 'Q',
                        'I': 'V', 'J': 'Z', 'K': 'N', 'L': 'T',
                        'M': 'O', 'N': 'W', 'O': 'Y', 'P': 'H',
                        'Q': 'X', 'R': 'U', 'S': 'S', 'T': 'P',
                        'U': 'A', 'V': 'I', 'W': 'B', 'X': 'R',
                        'Y': 'C', 'Z': 'J'}

const II = map[byte]byte{'A': 'A', 'B': 'J', 'C': 'D', 'D': 'K',
                         'E': 'S', 'F': 'I', 'G': 'R', 'H': 'U',
                         'I': 'X', 'J': 'B', 'K': 'L', 'L': 'H',
                         'M': 'W', 'N': 'T', 'O': 'M', 'P': 'C',
                         'Q': 'Q', 'R': 'G', 'S': 'Z', 'T': 'N',
                         'U': 'P', 'V': 'Y', 'W': 'F', 'X': 'V',
                         'Y': 'O', 'Z': 'E'}

const II = map[byte]byte{'A': 'B', 'B': 'D', 'C': 'F', 'D': 'H',
                         'E': 'J', 'F': 'L', 'G': 'C', 'H': 'P',
                         'I': 'R', 'J': 'T', 'K': 'X', 'L': 'V',
                         'M': 'Z', 'N': 'N', 'O': 'Y', 'P': 'E',
                         'Q': 'I', 'R': 'W', 'S': 'G', 'T': 'A',
                         'U': 'K', 'V': 'M', 'W': 'U', 'X': 'S',
                         'Y': 'Q', 'Z': 'O'}

const UKW_A = map[byte]byte{'A': 'E', 'B': 'J', 'C': 'M', 'D': 'Z',
                            'F': 'L', 'G': 'Y', 'H': 'X', 'I': 'V',
                            'K': 'W', 'N': 'R', 'O': 'Q', 'P': 'U',
                            'S': 'T'}

const UKW_B = map[byte]byte{'A': 'Y', 'B': 'R', 'C': 'U', 'D': 'H',
                            'E': 'Q', 'F': 'S', 'G': 'L', 'I': 'P',
                            'J': 'X', 'K': 'N', 'M': 'O', 'T': 'Z',
                            'V': 'W'}

// ==================================================================
//
// Structures Needed
//
// ==================================================================

type rotor struct {
    counter  int
    alphabet map[byte]byte
    offset   int
}

type enigma struct {
    rotors      []rotor
    breadboard  map[byte]byte
}

// ==================================================================
//
// Util Functions
//
// ==================================================================

func reverse(rots []rotor) []rotor {
    rev := make([]rotor, len(rots))
	for i := 0; i < len(rots)/2; i++ {
		j := len(rots) - i - 1
		rev[i], rev[j] = rots[j], rots[i]
	}
	return rev
}

func (machine *enigma) print(rotorOnly bool){
    fmt.Println("**************************")
    fmt.Println("****** Enigma State ******")
    fmt.Println("===== Breadboard =====")
    if !rotorOnly{
        for k, v := range machine.breadboard {
            fmt.Printf("%s is connected to %s\n", string(k), string(v))
        }
    }
    for i, rot := range machine.rotors {
        fmt.Printf("===== Cylinder #%d =====\n", i)
        fmt.Printf("counter: %d\n", rot.counter)
        fmt.Printf("offset: %d\n", rot.offset)
    }
    fmt.Println("******* End State ********")
    fmt.Println("**************************")
}

func getAsciiCylinder(cabling map[byte]byte) *rotor {
    rot := rotor{counter: 0,
                    offset: 0,
                    alphabet: cabling}
    return &rot
}


// ==================================================================
//
// Main Functions
//
// ==================================================================

func(rot *rotor) rotate() {
    rot.counter = rot.counter + 1
    rot.offset = rot.offset + 1
}


func (rot *rotor) encryptChar(increment bool, char byte) (bool, byte) {
    var t_char byte
    t_char = char
    char = char - 'A'
    index := byte((int(char) + rot.offset) % 26)
    if increment {
        t_counter, t_offset := rot.counter, rot.offset
        rot.counter++
        rot.offset++
        printf("Changed offset from %d to %d.\n", t_offset, rot.offset)
        printf("Changed counter from %d to %d.\n", t_counter, rot.counter)
    }
    overflow := rot.counter % 26 == 0
    println("index = ", index)
    char = rot.alphabet[index]
    printf("Shifted %s to %s.\n", string(t_char), string(char))
    return overflow, char
}

func (machine *enigma) encryptString(clearText string) string{
    printf("String to encrypt: %s\n", clearText)
    var cryptText strings.Builder
    rotorsReverse := reverse(machine.rotors)
    for _, ichar := range clearText {
        char := byte(ichar)
        printf("Encrypt char: %s\n", string(char))
        char = machine.breadboard[char]
        overflow := true
        for i, _ := range machine.rotors {
            overflow, char = machine.rotors[i].encryptChar(overflow, char)
        }
        overflow = false
        for i, _ := range rotorsReverse {
            overflow, char = machine.rotors[i].encryptChar(false, char)
        }
        printf("Encrypted to char: %s\n", string(char))
        cryptText.WriteByte(char)
    }
    return cryptText.String()
}

func main(){
    breadboard := map[byte]byte {'A' : 'A', 'B' : 'B', 'C' : 'C',
                                'D' : 'D', 'E' : 'E', 'F' : 'F',
                                'G' : 'G', 'H' : 'H', 'I' : 'I',
                                'J' : 'J', 'K' : 'K', 'L' : 'L',
                                'M' : 'M', 'N' : 'N', 'O' : 'P',
                                'Q' : 'Q', 'R' : 'R', 'S' : 'S',
                                'T' : 'T', 'U' : 'U', 'V' : 'V',
                                'W' : 'W', 'X' : 'X', 'Y' : 'Y',
                                'Z': 'Z',}
    //rotA := getAsciiCylinder()
    //rotB := getAsciiCylinder()
    //rotC := getAsciiCylinder()
    rotD := getAsciiCylinder()
    rotD.offset = 0
    machine := enigma{rotors: []rotor{ /**rotA, *rotB, *rotC,*/ *rotD},
                     breadboard: breadboard}
    machine.print(false)
    fmt.Println(machine.encryptString("AAA"))
}
