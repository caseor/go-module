package go-module

import "fmt" 

// v1
func SayHiV1(name string) string {
   return fmt.Sprintf("Hi, v1.0.0 %s", name)
}

// v1
func SayHiV2(name string) string {
   return fmt.Sprintf("Hi, v2.0.0 %s", name)
}