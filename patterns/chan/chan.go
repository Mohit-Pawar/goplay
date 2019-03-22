// package main

// import (
//   "fmt"
// )

// func f(left, right chan int) {
//   left<- <-right + 1
// }

// func main() {
//   n := 100
//   leftmost := make(chan int)
//   left := leftmost
//   right := leftmost
//   for i := 0; i < n; i++ {
//     right = make(chan int)
//     go f(left, right)
//     left = right
//   }
//   go func(c chan int) { c <- 1 }(right)
//   fmt.Println("Value", <-leftmost)
// }

// package main

// import (
//   "fmt"
// )

// func f1(in <-chan int) <-chan int {
//   out := make(chan int)
//   go func() {
//     val := <-in
//     out <- 1 + val
//   }()
//   return out
// }

// func f2(in <-chan int) <-chan int {
//   out := make(chan int)
//   go func() {
//     val := <-in
//     out <- 1 + val
//   }()
//   return out
// }

// func main() {
//   ch := make(chan int)
//   out := f1(ch)
//   out2 := f2(out)
//   go func(ch chan int) { ch <- 1 }(ch)
//   final := <- out2
//   fmt.Println("Output ", final)
// }


// package main

// import (
//   "fmt"
//   "bufio"
//   "os"
//   "strings"
// )

// func printer(in <-chan string, exit <-chan struct{}) <-chan struct{} {
//   done := make(chan struct{})
//   go func() {
//     for{
//       select {
//         case val := <-in:
//           fmt.Println("Value: ", val)
//         case <-exit:
//           fmt.Println("Exiting")
//           done <- struct{}{}
//           return
//       }
//     }
//   }()
//   return done
// }

// func main() {
//   ch := make(chan string)
//   exit := make(chan struct{})
//   reader := bufio.NewReader(os.Stdin)
//   done := printer(ch, exit)
//   for {
//     fmt.Println("Enter the Text")
//     text, _ := reader.ReadString('\n')
//     text = strings.Replace(text, "\n", "", -1)
//     if strings.Compare(text, "exit") == 0 {
//       exit <- struct{}{}
//       <-done
//       break
//     } else {
//       ch <- text
//     }
//   }
// }


// package main

// import (
//   "fmt"
//   "time"
//   "strconv"
// )

// type User struct {
//   Name string
// }

// var userQueue = make(chan User, 10)

// func main() {
//   num := 10
//   for i := 0; i < num; i++ {
//     go worker(i)
//   }
//   for i := 0; i < 10 * num; i++ {
//     name := strconv.Itoa(i)
//     userQueue <- User{Name: name}
//   }
//   time.Sleep(3 * time.Second)
// }

// func worker(index int) {
//   for {
//     select {
//       case user := <-userQueue:
//         fmt.Println(index, ") User: ", user.Name)
//     }
//   }
// }


// package main

// import (
//   "fmt"
//   "os/signal"
// )

// func main() {
//   fmt.Println("Ignoring all the signals")
//   signal.Ignore()
//   ch := make(chan int)
//   <-ch
// }


// package main

// import (
//   "fmt"
//   "os"
//   "os/signal"
// )

// func main() {
//   signals := make(chan os.Signal, 1)
//   signal.Notify(signals)
//   for {
//     signal := <-signals
//     fmt.Println("Got Signal: ", signal)
//   }
// }