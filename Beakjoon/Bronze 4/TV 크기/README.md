```go
package main

import (
	"fmt"
	"math"
)

func main(){
	var a,b,c float64
	fmt.Scanln(&a,&b,&c)
	d := math.Sqrt(a*a/(b*b+c*c))
	fmt.Println(math.Floor(b*d), math.Floor(c*d))
}
```
첫번째 코드다.  
836KB 4ms 174B

```go
package main

import (
	"fmt"
	"math"
)

func main(){
	var a,b,c float64
	fmt.Scanln(&a,&b,&c)
	d := math.Sqrt(math.Pow(a,2)/(math.Pow(b,2)+math.Pow(c,2)))
	fmt.Println(int(b*d), int(c*d))
}
```
두번째 코드다.  
820KB 4ms 190B  
math.Floor로 내림했지만 int로 바꿔줬다. 

```go
package main

import (
	"fmt"
	"math"
)

func main(){
	var a,b,c float64
	fmt.Scanln(&a,&b,&c)
	d := math.Sqrt(a*a/(b*b+c*c))
	fmt.Println(int(b*d), int(c*d))
}
```
세번째 코드다.  
832KB 4ms 160B  
pow랑 연산때리는것 중에서 어떤게 메모리 덜 먹을지 궁금해서 시도해봤다. 

```go
package main

import (
	"fmt"
	"math"
)

func main(){
	var a,b,c float64
	fmt.Scanf("%f %f %f",&a,&b,&c)
	d := math.Sqrt(math.Pow(a,2)/(math.Pow(b,2)+math.Pow(c,2)))
	fmt.Println(int(b*d), int(c*d))
}
```
네번째 코드다.  
828KB 4ms 200B  
Scanln대신 Scanf를 써봤다.  
8KB 차이가 났다.