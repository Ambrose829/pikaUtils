package util

import "strconv"

/**
  @Author: pika
  @Date: 2022/4/1 10:40 上午
  @Description: 基础类型转换
*/



func Str2F64(str string) float64 {
	f64, _ := strconv.ParseFloat(str, 64)
	return f64
}