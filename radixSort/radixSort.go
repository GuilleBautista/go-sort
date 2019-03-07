package lsd

/*import (
	"fmt"
	"math/rand"
	"time"
)*/

func maxDigits(arg0 []int) int {
	max:=0
	for _,number:=range(arg0){
		result:=number
		for i:=0;result>0;i++{
			result=int(result/10)
			if(i>max){
				max=i
			}
		}
	}
	return max+1
}

func Radix(arg0 []int) []int{
	ordered:=true
	for i:=0;ordered && i<len(arg0)-1;i++ {
		if(arg0[i]>arg0[i+1]){
			ordered=false
		}
	}
	lsd:=arg0
	maxDigits:=maxDigits(lsd)

	for i:=0;i<maxDigits;i++{
		var buckets [10][]int
		var digits = make([]int, i+1)
		for _,elem:=range(lsd){
			//fmt.Println(lsd) //Only testing purposes
			tmp:=elem
			//fmt.Println(tmp) //Only testin purposes

			/*To get the digit at position i, we create an array of
			digits and initialice it with a digit in each position.
			To do this we use a temporal variable "tmp", which
			will be reduced in one digit each time, i times. Then we
			access the i possition of the digits array. */
			for pos:=0; pos<=i; pos++{
				digits[pos]=tmp%10
				tmp=int(tmp/10)
			}
			tmp=digits[i]
			//fmt.Println("aux= ", tmp)	//only testing purposes

			//We add the number we are comparing to its "bucket"
			buckets[tmp]= append(buckets[tmp], elem)
		}
		lsd=nil
		for _,bucket:=range(buckets){
			for _,elem:=range(bucket){
				lsd=append(lsd, elem)
			}
		}
	}
	return lsd
}
