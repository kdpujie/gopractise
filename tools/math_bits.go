package main

import (
	"fmt"
	"github.com/willf/bitset"
	"os"
	"strings"
)

func main() {
	var b, b1, b2 bitset.BitSet
	b.Set(1).Set(2).Set(3)
	b1.Set(3).Set(4).Set(5).Set(100)
	inter := b.Intersection(&b1)
	u := b.Union(&b1)
	fmt.Fprintf(os.Stdout, "%s\n%s\n是否存在4：%v\n", b.DumpAsBits(), b.String(), b.Test(3))

	fmt.Fprintf(os.Stdout, "b和b1交集：%s,len=%d,len(b2)=%d \n", inter.String(), len(inter.String()), len(b2.String()))
	fmt.Fprintf(os.Stdout, "b和b1并集：%s len=%d \n", u.String(), len(u.String()))
	s := u.String()
	noPrefix := strings.TrimPrefix(s, "{")
	noSuffix := strings.TrimSuffix(noPrefix, "}")
	v := strings.Split(noSuffix, ",")
	fmt.Fprintf(os.Stdout, "去掉前缀：%s，去掉后缀：%s，转化为切片：%v \n", noPrefix, noSuffix, v)

}
