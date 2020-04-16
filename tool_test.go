package main

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"runtime"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

func TestInit(t *testing.T) {
	fmt.Println("test")
}

//func BenchmarkAlgorithmThree(B *testing.B){
//	fmt.Println(111);
//}

//func ExampleHello() {
//	fmt.Println("hello")
//	// Output: hello
//}

func TestSortInt(t *testing.T) {
	var is sort.IntSlice = sort.IntSlice{3, 4, 8, 1, 4, 5, 6, 7, 5, 4, 3, 2, 1}
	//is.Sort();
	fmt.Println(sort.IsSorted(is))
	fmt.Println(is)
	sort.Stable(is)
	//sort.Sort(is);
	fmt.Println(sort.IsSorted(is))
	fmt.Println(is)
	//ris:=
	sort.Sort(sort.Reverse(is))
	fmt.Println(is)
	//for i:=0;i<ris.Len() ;i++  {
	//	fmt.Println(ris[i])
	//}

	p := is.Search(5)
	fmt.Println(p)
	fmt.Println(sort.SearchInts(is, 5))
	fmt.Println(is.Len())
	//fmt.Println(is);

}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func Example() {
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
	// Output: people
}

func TestTcp(t *testing.T) {
	//建立连接
	d := &net.Dialer{
		Timeout:       60 * time.Second,
		Deadline:      time.Now().Add(100 * time.Second),
		LocalAddr:     nil,
		FallbackDelay: 0,
		KeepAlive:     0,
		Resolver:      nil,
		Control:       nil,
	}
	var err error
	var ctx context.Context
	ctx = context.TODO()
	ctx = context.WithValue(ctx, "method", "GET")
	var conn net.Conn
	conn, err = d.DialContext(ctx, "tcp", "127.0.0.1:80")
	//conn, err := net.Dial("tcp", "47.52.105.170:80")
	if err != nil {
		fmt.Println("net.Dialt err", err)
		return
	}
	defer conn.Close()
	msg := "GET /index.html \r\n"
	msg += "HEAD /HTTP/1.1\r\n"
	msg += "127.0.0.1:80\r\n"
	msg += "Connection: close\r\n"
	msg += "Content-Type: text/html;charset=utf-8\r\n"
	msg += "\r\n\r\n"
	//fmt.Println(msg);

	//发送文件名到接收端
	_, err = conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("conn.Write err", err)
		return
	}
	buf := make([]byte, 4096)
	//接收服务器返还的指令
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err", err)
		return
	}
	fmt.Println(string(buf[:n]))
	////返回ok，可以传输文件
	//if string(buf[:n]) == "ok" {
	//	sendFile(conn, filepath)
	//}

}

func TestParam(t *testing.T) {
	t.Log(runtime.NumGoroutine())
	t.Log(runtime.NumCPU())
	//t.Log(runtime.NumCgoCall())
}

func OneTask(once *sync.Once) int {
	var a = 1
	once.Do(func() {
		a = a + 2
		fmt.Println("aaa")
	})

	return a

}
func TestOneTask(t *testing.T) {
	var once = &sync.Once{}
	for i := 1; i <= 7; i++ {
		num := OneTask(once)
		fmt.Println(num)
	}
	time.Sleep(2 * time.Second)
}

type OTask struct {
	TaskId   int
	TaskName string
}

func OneTaskReturn(id int) *OTask {
	time.Sleep(100 * time.Millisecond)
	fmt.Println(id)
	return &OTask{TaskId: id, TaskName: fmt.Sprintf("第%d个任务", id)}
}

func TestOnTaskReturn(t *testing.T) {
	var ch = make(chan *OTask, 10) //会造成线程溢出
	//var ch =make(chan *OTask,10)
	//var task *OTask;
	t.Log("before num:", runtime.NumGoroutine())
	for i := 0; i < 10; i++ {
		go func(i int) {
			ch <- OneTaskReturn(i)
		}(i)
	}
	//var isTimeOut=false
	//t.Log(<-ch)
	select {
	case obj := <-ch:
		t.Log(obj)
	case <-time.After(time.Millisecond):
		t.Log("time out")

	}
	t.Log("middle Num:", runtime.NumGoroutine())
	time.Sleep(time.Second)
	t.Log("after Num:", runtime.NumGoroutine())
}

func TestAllTaskReturn(t *testing.T) {
	var num = 10
	var ch = make(chan *OTask) //会造成线程溢出
	//var ch =make(chan *OTask,10)
	//var task *OTask;
	t.Log("before num:", runtime.NumGoroutine())

	for i := 0; i < num; i++ {
		go func(i int) {
			ch <- OneTaskReturn(i)
		}(i)
	}
	for i := 0; i < num; i++ {
		t.Log(<-ch)
	}
	t.Log("after Num:", runtime.NumGoroutine())
}

func BenchmarkWriteString(b *testing.B) {
	var strArr []string = []string{"广东省", "广州市", "番禺区", "小谷围街道", "青蓝街", "有米科技", "三楼", "讯码科技", "后台研发部"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var str = ""
		for _, value := range strArr {
			str += value
		}
		//b.Log(str)
	}

	b.StopTimer()
}

func BenchmarkBufByte(b *testing.B) {
	var strArr []string = []string{"广东省", "广州市", "番禺区", "小谷围街道", "青蓝街", "有米科技", "三楼", "讯码科技", "后台研发部"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var str = bytes.Buffer{}
		for _, value := range strArr {
			str.WriteString(value)
		}
		//b.Log(str.String())
	}

	b.StopTimer()
}

func BenchmarkBufString(b *testing.B) {
	var strArr []string = []string{"广东省", "广州市", "番禺区", "小谷围街道", "青蓝街", "有米科技", "三楼", "讯码科技", "后台研发部"}
	//var str=bytes.Buffer{};

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var str = strings.Builder{}
		for _, value := range strArr {
			str.WriteString(value)
		}
		//b.Log(str.String())
	}

	b.StopTimer()
}

func BenchmarkStringJoin(b *testing.B) {
	var strArr []string = []string{"广东省", "广州市", "番禺区", "小谷围街道", "青蓝街", "有米科技", "三楼", "讯码科技", "后台研发部"}
	//var str=bytes.Buffer{};

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strings.Join(strArr, "")
		//b.Log(str)
	}

	b.StopTimer()
}

func BenchmarkFmtString(b *testing.B) {
	//var strArr []string=[]string{"广东省","广州市","番禺区","小谷围街道","青蓝街","有米科技","三楼","讯码科技","后台研发部"}
	//var str=bytes.Buffer{};
	var sqlStr = "select * from tb_content where id=%d and name=%s"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf(sqlStr, 1, "my")
		//b.Log(str.String())
	}

	b.StopTimer()
}

func TestRecover(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log(err)
		}
	}()
	var a [10]int
	for i := 0; i < 11; i++ {
		a[i] = i
	}
}

func TestChannel(t *testing.T) {
	var ch = make(chan int, 10)
	for i := 0; i < 10; i++ {
		go func() {
			//t.Log("start-------")
			time.Sleep(time.Second * 4)
			ch <- 1
			// t.Log("end---------")
		}()
	}
	var a = 1
	var b = 2
	var abd = 3
	var add = 4
	t.Log(a, b, abd, add)
	var stat = &runtime.MemStats{}
	runtime.ReadMemStats(stat)
	t.Log("alloc:", stat.Alloc)
	t.Log("StackInuse:", stat.StackInuse)
	t.Log("HeapIdle", stat.HeapIdle)
	time.Sleep(time.Second * 8)
	runtime.ReadMemStats(stat)
	t.Log(runtime.NumGoroutine())
	t.Log("alloc:", stat.Alloc)
	t.Log("StackInuse:", stat.StackInuse)
	t.Log("HeapIdle", stat.HeapIdle)
}

func TestAtomic(t *testing.T) {
	//增减操作
	var a int32 = 1
	t.Log("a : ", a)
	//函数名以Add为前缀，加具体类型名
	//参数一，是指针类型
	//参数二，与参数一类型总是相同
	//增操作

	t.Log("loadInt32", atomic.LoadInt32(&a))
	new_a := atomic.AddInt32(&a, 3)
	t.Log("new_a : ", new_a)
	t.Log("a", a)
	//减操作
	new_a = atomic.AddInt32(&a, -2)
	t.Log("new_a : ", new_a)
	t.Log("a", a)
}

func TestAtomicChannel(t *testing.T) {
	var a int32 = 0
	var wg sync.WaitGroup
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func(i int) {
			//time.Sleep(time.Millisecond * 300)
			//a++
			_ = atomic.AddInt32(&a, 1)
			//t.Log(new)
			//t.Log(a)
			wg.Done()
		}(i)
	}

	wg.Wait()
	//time.Sleep(time.Second * 1)
	t.Log(a)
}

func TestUnsafe(t *testing.T) {
	var a float32 = 3333
	t.Log(a)
	tmp := unsafe.Pointer(&a)
	t.Log(unsafe.Sizeof(tmp))
	t.Log(unsafe.Alignof(tmp))
	//t.Log(unsafe.Offsetof(tmp))
	var b = (*float64)(tmp)
	t.Log(*b)
	var c float64 = 6
	t.Log(*b + c)
}
