package concurrent

// func TestConcurrent(t *testing.T) {
// 	c := New(200, true)
// 	retInt := func(i int) int {
// 		time.Sleep(time.Duration((i%3 + 1)) * time.Second)
// 		return i
// 	}
// 	ThreadCount := 1000
// 	for i := 0; i < ThreadCount; i++ {
// 		if err := c.Call(retInt, i); err != nil {
// 			panic(err)
// 		}
// 	}
// 	c.Close()
// }
