package main
import (
  "fmt" //input and output package
  "time"//to retrieve the date and time package
  "math/rand"//to get the random integernumber
  "math"
  "runtime" //math operations
)
 var avenger, dc,starwar string = "moon","superman","valder"
 //not seperate declaration and initialization are allowed outside of function
 /*
 bool ->false(default)
 int -> 0(default)
 */
 
func add(age int,number int){
  fmt.Println(age+number)
}
func requires(quick int,wicket int){
  fmt.Println(quick-wicket)
}
func multi(height float64,weight float64) float64{
  return height * weight
}
func wanted(Variable1 ,Variable2 int) int{ //function continued
  return Variable1+Variable2
 }
 func result(i,j string) (string,string){ //multiple return values
   return i,j
 }
 //named return values
 func display(num int)(x,y int){
   x = num*num
   y=num+num
   return
 }
 func loop(sun int){
  sum:=0
  for i:=0; i<sun;i++{
  sum+=i
  }
   fmt.Println(sum)
  }
  
  func forcontinued(){
    naturals:=1
    for ; naturals<10;{ //not initialization needed and increment are also not needed

      naturals =naturals+naturals
    }
    fmt.Println(naturals)
  }
  //while loop:
  func whilefor(){ //more same like forcontinued
    cricket :=1
    for cricket<10{
      cricket +=cricket
    }
    fmt.Println(cricket)
  }
  func sqrt(x float64) string { //if statement
    if (x < 0) {
      return sqrt(-x) //negative go to if and continue function again
    }
    return fmt.Sprint(math.Sqrt(x)) //positive no go to return statement
  }
  func pow(x, n, lim float64) float64 { //upgraded version of if condition
    if v := math.Pow(x, n); v < lim { //lightly have the touch of for loop
      return v
    }
    return lim
  }
  func show(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
      return v
    } else{
    fmt.Printf("%v >= %v \n",v,lim)
    } 
    return lim
  }
type Moonknight struct{
  color int
  length int
}
var(
  v1 = Moonknight{1,2}
  v2 = Moonknight{color : 90}
  v3 = Moonknight{}
  v4 = &Moonknight{12,9}
)  

func main(){
  //Zero Values:
  var value string
  var phone int
  var laptop float32
  var control bool
  fmt.Println(value,phone,laptop,control)
  var c,d = true,false //specific type of variable declaration
  fmt.Println(c,d)
  var a,b int = 23,22
  karthik :="Omkar"
  fmt.Println(karthik)
  school,college,office := 12,"degree","lifelong"
  fmt.Println(school,college,office)
  fmt.Println(avenger,dc,starwar) //variables
  fmt.Println(a,b)
  fmt.Println("Hello World!")
  fmt.Println(time.Now())
  fmt.Println("the random numbe is ",rand.Intn(100))
  fmt.Println(rand.Intn(10))
  fmt.Println(math.Sqrt(9))
  fmt.Println(math.Pi)
  add(1212,12)
  requires(1892,176)
  fmt.Println(multi(12.342,122.3122))
  fmt.Println(wanted(11,11))
  fmt.Println(result("Hello","World"))
  //Type conversion:
  var time float32=99.386
  fmt.Println(time)
  var water int = int(time)
  fmt.Println(water)   //float to int

  var drink int = 22
  fmt.Println(drink)
  var cool float32 =float32(drink)
  fmt.Println(cool)

  //type identification
  fmt.Printf("%T\n",time )
  fmt.Printf("%T",cool)

  //constant variable :
  const PIPE = "Ronaldo"
  fmt.Println(PIPE)
  //PIPE = "virat"("the error may occur because const donot allow modification")
  //const variable can put either in local or global declaration
   
  //Loop:FOR
  loop(10)
  //for continued
  forcontinued()
  //whilefor loop
  whilefor()
  fmt.Println(sqrt(2), sqrt(-4)) //if
  fmt.Println(  pow(3, 2, 10), pow(3, 3, 20)) //enhanced if
  fmt.Println(show(3,2,10),show(3,3,20)) //if and else

  //switch case example:

    fmt.Print("Go runs on ")
	switch os := runtime.GOOS
	os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
	fmt.Println(runtime.GOOS) //linux
  
  // Another Example:
  num :=5
switch num{
case 1:
fmt.Println("1")
case 2:
fmt.Println("2")
case 3:
fmt.Println("3")
case 4:
fmt.Println("4")
case 5:
fmt.Println("5")
default :
fmt.Println("NO")
}
/* SWITCH CASE WITH NO CONDITION:
t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
*/
/*
fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

*/

//DEFER
defer fmt.Println("world")
fmt.Println("hello")
//STACKING DEFER
fmt.Println("counting")
for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
fmt.Println("done")

//Pointer in Go:
Dish := "Elon"
fmt.Println(Dish) //String
Dish2 := &Dish
fmt.Println(Dish2)
fmt.Println(*Dish2)

//Structs: (User Defined Datatype)
fmt.Println(Moonknight{12,12})
fmt.Println(Moonknight{9,18})
 
//Struct Fields:
//call the struct :
Pammal := Moonknight{12,4}
fmt.Println("The struct variable color is",Pammal.color)
fmt.Println("The struct variable length",Pammal.length)

//Pointer to struct:
PonniNagar := Moonknight{345,18}
Ponni := &PonniNagar 
Ponni.color = 74 //not the variable creation 
fmt.Println(Ponni.color) //avoid the (*) symbol

//Struct Literals:
fmt.Println(v1,v2,v3,v4)
}






	

