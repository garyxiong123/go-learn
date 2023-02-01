package panic

import "testing"

//空指针引用
func Test_Panic_Scenario_NullPoint(t *testing.T) {

}

//下标越界
func Test_Panic_Scenario_OutOfIndex(t *testing.T) {
	//testOutOfIndex(3)

}

func testOutOfIndex(i int) {
	//var balance = [2]int{1, 2}
	//println(balance)
	//m := balance[i]
	//println(m)
}

//除数为0
func Test_Panic_Scenario_Div_Zero(t *testing.T) {
	//m := 1 / 0
	//println(m)
}

//不应该出现的分支，比如default

//输入不应该引起函数错误
