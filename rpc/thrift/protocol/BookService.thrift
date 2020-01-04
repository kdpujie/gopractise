namespace go thriftrpc
namespace java com.ksyun.protocol.thrift.test

struct Work {
  1: i32 num1 = 0,
  2: i32 num2,
  3: optional string comment,
}

service BookService{
    
    string readBook(1:string name,2:Work work)
}