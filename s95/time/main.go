package main

import "time"

//type MYDuration = time.Duration
type MYDuration  time.Duration

func (m MYDuration) simple()  {  // cannot define new methods on non-local type time.Duration

}
func main()  {

}
