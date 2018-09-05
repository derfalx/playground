package main

import ("fmt"
        "strings")


// Sample project to learn the basics of Go
// Goal is to implement an Enigma machine to en-
// and decrypt text

type rotor struct {
    counter  int
    alphabet []byte
    offset   int
}

type enigma struct {
    rotors      []rotor
    breadboard  map[byte]byte
}

func getAsciiCylinder() *rotor {
    rot := rotor{counter: 0,
                    offset: 0,
                    alphabet: []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}}
    return &rot
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

func (rot *rotor) encryptChar(increment bool) (bool, byte) {
    index := rot.offset % 26
    if increment {
        rot.counter++
        rot.offset++
    }
    overflow := rot.counter % 26 == 0
    char := rot.alphabet[index]
    return overflow, char
}

func (machine *enigma) encryptString(clearText *string){
    var cryptText strings.Builder
    rotorsReverse := make([]rotor, len(machine.rotors))
    for ichar := range *clearText {
        char := byte(ichar)
        char = machine.breadboard[char]
        _ = rotorsReverse
        cryptText.WriteByte(char)
    }
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
    rotA := getAsciiCylinder()
    rotB := getAsciiCylinder()
    rotC := getAsciiCylinder()
    rotD := getAsciiCylinder()
    machine := enigma{rotors: []rotor{ *rotA, *rotB, *rotC, *rotD},
                     breadboard: breadboard}
    machine.print(false)
}
