package mypackage

import "fmt"

// MyPc Car con acceso publico
type MyPc struct {
	Ram  int
	Disk int
	Band string
}

//String Output personalizado de una estructura propio
func (myPc MyPc) String() string {
	return fmt.Sprintf("Tengo %d FB RAM, %d GB Disco y es una %s", myPc.Ram, myPc.Disk, myPc.Band)
}
