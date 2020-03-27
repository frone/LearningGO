package test

import (
	"errors"
	"testing"
)

//通过多参数返回进行错误处理
var LessThanTwoError = errors.New("parameter should be more than 2!")
var LagerThanHundredError = errors.New("parameter should be less than 100!")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, LagerThanHundredError
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(-10); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}

}
