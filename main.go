package main

import (
	"fmt"
	"github.com/go-macaron/inject"
	_ "github.com/go-macaron/inject"
	"kubepack.dev/lib-helm/pkg/repo"
	_ "kubepack.dev/lib-helm/pkg/repo"
)

func main() {
	reg := repo.NewDiskCacheRegistry()

	i := inject.New()
	i.Map(reg)
	i.MapTo(reg, (*repo.IRegistry)(nil)) // also works without this line

	i.Invoke(check)

	fmt.Println("hello")
}

func check(r1 *repo.Registry, r2 repo.IRegistry) {
	fmt.Println(r1 != nil)
	fmt.Println(r2 != nil)
}
