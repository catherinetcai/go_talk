package packageA

var ThisWillBeExported
var thisWillNotBeExported

package packageB

import "packageA"

func AFunction() {
	b := packageA.ThisWillBeExported
}
