package packageA

var ThisWillBeExported
var thisWillNotBeExported

package packageB

func AFunction() {
	b := packageA.ThisWillBeExported
}
