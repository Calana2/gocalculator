package utils

import (
 "testing"
 "utils"
)

func TestEvaluateRPN_DivByZero(t *testing.T){
 rpn := []string{"10","0"}
 result := utils.EvaluateRPN(rpn)

 if result == 0 {
  t.Log("Error: Expected result to be 0")
  t.Fail()
 }
}

func TestInfixToRPN_NotAllowedExpression(t *testing.T){
 _,err1 := utils.InfixToRPN("2$ + 121.3%")
 _,err2 := utils.InfixToRPN("bla bla bla")

 if err1 == nil || err2 == nil {
  t.Log("Expected error code to be -1")
  t.Fail()
 }
}

func TestInfixToRPN_FirstNegative(t *testing.T){
 exp := " -2 + 4"
 rpn,_ := utils.InfixToRPN(exp)

 if rpn[0] != "-2" {
  t.Log("Expected result to be 2")
  t.Fail()
 }
}

func TestEvaluateRPN_BadOperator(t *testing.T){
 exp := "2 ++ 1"
 rpn,_ := utils.InfixToRPN(exp)
 result := utils.EvaluateRPN(rpn)

 if result != 0 {
  t.Log("Expected result to be 0")
  t.Fail()
 }
}



