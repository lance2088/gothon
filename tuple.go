package gothon

import (
	"encoding/binary"
)

type Tuple []Object

func (tuple *Tuple) Read(r *Reader, t byte) {
	size := tupleSize(r, t)

	//fmt.Printf("Tuple of size: %d\n", size)

	*tuple = make([]Object, size)

	for i := 0; i < int(size); i++ {
		(*tuple)[i] = r.ReadObject()
	}
}

func tupleSize(r *Reader, t byte) int {
	if t == '(' {
		var size int32
		binary.Read(r, binary.LittleEndian, &size)
		return int(size)
	} else if t == ')' {
		var size int8
		binary.Read(r, binary.LittleEndian, &size)
		return int(size)
	} else {
		panic("Cannot interpred tuple size!")
	}
}
