package main

import (
	"github.com/wejick/poc_tego/package1"
	"github.com/wejick/poc_tego/package2"
)

func main() {
	function2 := func(interface1 package1.Interface1) {
	}
	function2(package2.Struct1{})

	// not sure which one was used, but this one would error
	// ./app.go:11: cannot use package2.Struct1 literal (type *package2.Struct1) as type *package1.Interface1 in argument to
	//  function2:
	//         *package1.Interface1 is pointer to interface, not interface
	// normally we don't create pointer to interface since interface itself is referenced object just like slice or map

	// function2 := func(interface1 *package1.Interface1) {
	// }
	// function2(&package2.Struct1{})
}
