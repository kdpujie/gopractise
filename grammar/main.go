// main
package grammar

import (
	"fmt"
	"io"
	//	"log"
	"encoding/json"
	"math"
	"math/cmplx"
	"math/rand"
	"net/http"
	"os"
	//	"reader"
	"runtime"
	"strings"
	"time"
	"unicode"
)

var c, python, java bool

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	Big   = 1 << 100
	Small = Big > 99
)

func main() {
	Test()
	fmt.Printf("My first go program!")
	fmt.Println("Hello World!")
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("Now you have %g problems", math.Sqrt(7))
	fmt.Println(math.Pi) //math.pi 小写的未被导出可见
	fmt.Println(add(43, 13))
	fmt.Println(addn(11, 22))
	a, b := swap("hello", "world") //:=简洁赋值语句
	fmt.Println(a, b)
	fmt.Println(split(17)) //裸返回，当前值
	var i int
	var str string = "i am a string"
	fmt.Println(i, c, python, java, str)
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
	fmt.Printf("%v\n%q", "string", "Pstring")
	t_int := 42
	var f_float float64 = float64(t_int)
	//var f_float_2 float64 = t_int 必须显示转换
	fmt.Println(t_int, f_float)

	//	fmt.Println(needInt(Small))
	//	fmt.Println(needFloat(Small))
	//	fmt.Println(needFloat(Big))

	for i = 1; i < 100; i++ {
		fmt.Printf("%v\t", i)
		if v := i * 2; v < 100 {
			fmt.Printf("%v\n", v)
		}
	}

	fmt.Printf("\nsqrt(%v)=%v\n", 7, sqrt(7))

	defer osifo()

	fmt.Println(time.Now())

	type coor struct {
		x int
		y int
	}
	fmt.Println(coor{100, 100})
	var v coor = coor{50, 50}
	fmt.Println(v.x)

	var p *int
	var tmp int = 100
	p = &tmp
	fmt.Println(*p)

	var pt *coor = &v
	fmt.Println(pt.x)

	var arr [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(arr, "length ", len(arr))

	sliceTest()

	var arr1 = make([]int, 10, 100)
	arr1 = arr1[10:100]
	arr1[20] = 100 //需要进行切片才能访问(空的为nil)
	//arr1 = append(arr1, 1, 2, 3)
	fmt.Println(arr1, "len=", len(arr1), "cap=", cap(arr1), "arr1[20]", arr1[20]) //数组长度为len，容量为cap，预留
	for i, v := range arr1 {
		fmt.Printf("arr[%v]=%v\n", i, v)
	}

	for _, v2 := range arr1 {
		fmt.Printf("%v", v2)
	}

	//图片的显示
	//pic.Show(getpic)

	//map的使用
	type ele struct {
		a int
		b int
		s bool
	}
	var map_ele map[string]ele = make(map[string]ele)
	var map_ele2 = map[string]ele{
		"world": {50, 50, true},
		"hello": {100, 100, true}, //注意最后的这个逗号是必须的！
	}
	map_ele["hello"] = ele{100, 100, true}
	fmt.Println(map_ele["hello"])
	fmt.Println(map_ele2)
	val, ok := map_ele2["hello"]
	fmt.Println("value:", val, "exist:", ok)
	delete(map_ele2, "hello")
	val, ok = map_ele2["hello"]
	fmt.Println("value:", val, "exist:", ok)

	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))

	jud := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Printf("Fields are: %q", strings.FieldsFunc("  foo1;bar2,baz3...", jud))

	//采用函数闭包计算斐波拉契
	foc := fibonacci()
	for i := 0; i < 100; i++ {
		fmt.Println(foc())
	}

	var ieletest = &iele{11, 11, true}
	fmt.Println("ieletest.all() ", ieletest.all(), "ieletest.off()", ieletest.off())

	var pv = strr{"hello"}
	var po intt
	po = &pv
	fmt.Println(po.getstring(), po.getstring(), po.getstring())

	//错误信息打印
	fmt.Println(Sqrt_Err(2))
	fmt.Println(Sqrt_Err(-2))

	reder := strings.NewReader("Hello, Reader!")
	buf := make([]byte, 8)
	for {
		n, err := reder.Read(buf)
		fmt.Printf("n = %v err = %v b = %q\n", n, err, buf[:n])
		if err == io.EOF {
			break
		}
	}

	//	reader.Validate(MyReader{})

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)

	//	var web HttpServer
	//	weberr := http.ListenAndServe("localhost:4000", web)
	//	if weberr != nil {
	//		log.Fatal(weberr)
	//	}

	channel := make(chan string)
	go say("world!", channel)
	//	say("hello", channel)	//管道中使用该方法会导致死锁
	go say("hello", channel)
	chanmsg1, chanmsg2 := <-channel, <-channel
	fmt.Println(chanmsg1, chanmsg2)
	close(channel)

	channel1 := make(chan string, 10)
	fmt.Println("cap(channel1)=", cap(channel1))
	channel2 := make(chan string, 10)
	go talk(channel1, cap(channel1))
	go talk(channel2, cap(channel2))
	listen(channel1, channel2)

	fmt.Println(time.Millisecond)

	timeBoom()

	jsontest()

	maptest()

	type Ad struct {
		ifo  int
		name string
	}
	type Lot struct {
		Ads  []*Ad
		size int
	}
	var ads = make([]*Ad, 10)
	ads[0] = &Ad{
		ifo:  100,
		name: "hanse",
	}

}

//==============================================================================
func add(x int, y int) int {
	return x + y
}

func addn(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

//牛顿迭代法求平方根
func sqrt(x float64) float64 {
	var t float64 = 1.0
	for i := 1; i < 10; i++ {
		t = (t + x/t) / 2
	}
	return t
}

func osifo() {
	fmt.Printf("\nGO run on ")
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Linux")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Printf("Others(%v)\n", os)
	}
}

func printBoard(s [][]string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%s\n", strings.Join(s[i], " "))
	}
}

func sliceTest() {
	txt := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	txt[0][0] = "X"
	txt[2][2] = "O"
	txt[2][0] = "X"
	txt[1][0] = "O"
	txt[0][2] = "X"
	printBoard(txt)
}

func getpic(dx, dy int) [][]uint8 {
	img := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		img[i] = make([]uint8, dx)
	}
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			img[i][j] = uint8(i*j + i ^ j + (i+j)/2)
		}
	}
	return img
}

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func() int {
	var per int = 0
	var next int = 1
	var tmp int = -1
	return func() int {
		if tmp < 1 {
			tmp++
			return tmp
		}
		tmp = per + next
		per = next
		next = tmp
		return tmp
	}
}

type iele struct {
	x int
	y int
	s bool
}

//值访问不能改变内部值（为副本访问）v iele，指针访问可改变值v *iele
func (v iele) all() int {
	return (v.x + v.y)
}

func (v *iele) off() bool {
	v.s = false
	return v.s
}

//接口
type intt interface {
	getstring() string
}

type strr struct {
	str string
}

func (s *strr) getstring() string {
	s.str = s.str + "_____"
	return s.str
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number:", float64(e))
}

func Sqrt_Err(x float64) (float64, error) {
	if x < 0 {
		var e = ErrNegativeSqrt(x)
		return x, &e
	}
	var t float64 = 1.0
	for i := 1; i < 10; i++ {
		t = (t + x/t) / 2
	}
	return t, nil
}

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (m MyReader) Read(b []byte) (int, error) {
	var len int = len(b)
	for i := 0; i < len; i++ {
		b[i] = byte('A')
	}
	return len, nil
}

type rot13Reader struct {
	r io.Reader
}

func (rot13 rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot13.r.Read(p)
	for i := 0; i < len(p); i++ {
		if (p[i] >= 'A' && p[i] < 'N') || (p[i] >= 'a' && p[i] < 'n') {
			p[i] += 13
		} else if (p[i] > 'M' && p[i] <= 'Z') || (p[i] > 'm' && p[i] <= 'z') {
			p[i] -= 13
		}
	}
	return
}

type HttpServer struct {
}

func (web HttpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "WEB Hello world!")
}

func say(s string, c chan string) {
	for i := 0; i < 10; i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Println(s)
	}
	c <- s
}

func talk(c chan string, n int) {
	for i := 0; i < n; i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Printf(".")
		c <- "word"
	}
	c <- "quit"
	close(c)
	fmt.Println("Talk OK!")
}

func listen(c1, c2 chan string) {
	time.Sleep(10 * time.Millisecond)
	for {
		//select 选择可执行的语句进行执行
		select {
		case s := <-c1:
			time.Sleep(10 * time.Millisecond)
			fmt.Println("chan c1:", s)
		case st := <-c2:
			time.Sleep(10 * time.Millisecond)
			fmt.Println("chan c2:", st)
			if st == "quit" {
				return
			}
		}
	}
}

func timeBoom() {
	tick := time.Tick(1000 * time.Millisecond)
	boom := time.After(5000 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func jsontest() {
	//`为特殊标点
	type Message struct {
		Name string `json:"msg_name"`       // 对应JSON的msg_name
		Body string `json:"body,omitempty"` // 如果为空置则忽略字段
		Time int64  `json:"-"`              // 直接忽略字段
	}
	var m = Message{
		Name: "Alice",
		Body: "",
		Time: 1294706395881547000,
	}
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	fmt.Println(string(data))
}

func maptest() {
	m := make(map[string]string)
	m["hello"] = "echo hello"
	m["world"] = "echo world"
	m["go"] = "echo go"
	m["is"] = "echo is"
	m["cool"] = "echo cool"

	for k, v := range m {
		fmt.Printf("k=%v, v=%v\n", k, v)
	}
}

const EXPIRATION_TIME = 30 * 60 //过期时间，单位秒
func Test() {

	Visit("www.baidu.com", B{s: "hello"}, A{s: "world"})

	os.Exit(0)

	type person struct {
		name string
		age  int
	}
	one := person{name: "dingding", age: 100}
	two := one
	two.name = "dongdong"
	fmt.Println(one)
	fmt.Println(two) // 值语义

	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays // 这个常量没有导出
	)
	fmt.Println(Saturday) //Saturday=6

	var times uint32
	now := time.Now().Unix()
	times = uint32(now + EXPIRATION_TIME)
	fmt.Println(times)

	//1467109709
	//1467106601			百度
	//1467107109178			java
	//1467107006098864300	go 纳秒

	var arr1 = make([]int, 10, 100)
	arr1 = arr1[10:100]
	arr1[20] = 100 //需要进行切片才能访问(空的为nil)
	arr1 = append(arr1, 1, 2, 3)
	fmt.Println(arr1, "len=", len(arr1), "cap=", cap(arr1), "arr1[20]", arr1[20]) //数组长度为len，容量为cap，预留
	for i, v := range arr1 {
		fmt.Printf("arr[%v]=%v\n", i, v)
	}

	//	var arr2 = make([]byte, -1)
	//	for i, v := range arr2 {
	//		fmt.Printf("arr[%v]=%v\n", i, v)
	//	}

	os.Exit(0)
}

func Visit(url string, a ...interface{}) (string, int) {
	return doVisit(url, a)
}
func doVisit(url string, a []interface{}) (string, int) {
	var s string
	var e int

	switch arg1 := a[0].(type) {
	default:
		switch arg2 := a[1].(type) {
		default:
			fmt.Println(arg1, arg2)
			//			visit(reflect.ValueOf(a[0]), reflect.ValueOf(a[1]))
		}
	}

	return s, e
}

type B struct {
	s string
}
type A struct {
	s string
}

func visit(a A, b B) {
	fmt.Println(a, b)
}
