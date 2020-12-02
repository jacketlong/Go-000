学习笔记
1.Go语言是多返回值的，通常使用bool或者error来确定函数调用成功或者失败
bool：当错误原因只有一种，或者上层应用不关心错误原因，只关心结果时使用
error：错误的原因有很多，且调用者关心错误原因时使用
2.error 本质
// error 是一个普通的接口
type error interface {
	Error() string
}
