## 函数声明
```go
 func name(parameterList)(returnList){
 	body
 }
```
go中进行参数传递是值传递，调用函数时传递的是拷贝对象，只有传递引用类型才会在修改参数的基础上对参数值进行修改。