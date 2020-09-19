package loader

import "C"
import (
	"fmt"
)

//export v0
func v0() uintptr {
	return 0
}

//export v1
func v1(uintptr) uintptr {
	fmt.Println("ssssssssssssssssssssssss")
	return 0
}

//export v2
func v2(uintptr, uintptr) uintptr {
	return 0
}

//export v3
func v3(uintptr, uintptr, uintptr) uintptr {
	return 0
}

//export v4
func v4(uintptr, uintptr, uintptr, uintptr) uintptr {
	return 0
}

//export v5
func v5(uintptr, uintptr, uintptr, uintptr, uintptr) uintptr {
	return 0
}

//export v6
func v6(uintptr, uintptr, uintptr, uintptr, uintptr, uintptr) uintptr {
	return 0
}

//export v7
func v7(uintptr, uintptr, uintptr, uintptr, uintptr, uintptr, uintptr) uintptr {
	return 0
}

//export v8
func v8(uintptr, uintptr, uintptr, uintptr, uintptr, uintptr, uintptr, uintptr) uintptr {
	return 0
}

//export v9
func v9(uintptr, uintptr, uintptr, uintptr, uintptr, uintptr, uintptr, uintptr, uintptr) uintptr {
	return 0
}

//export v10
func v10(uintptr, uintptr, uintptr, uintptr, uintptr, uintptr, uintptr, uintptr, uintptr, uintptr) uintptr {
	return 0
}
