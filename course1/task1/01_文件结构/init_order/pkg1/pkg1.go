package pkg1

import (
	"fmt"

	_ "github.com/Zack-Zz/go-study/course1/init_order/pkg2"
)

const PkgName string = "pkg1"

var PkgNameVar string = getPkgName()

func init() {
	fmt.Println("pkg1 init method invoked")
}

func getPkgName() string {
	fmt.Println("pkg1.PkgNameVar has been initialized")
	return PkgName
}
