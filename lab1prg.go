package main


func main() {
              var prev,current,sum int
              prev=0
              current=1
              sum=0
              for prev+current<40{
              
                  next:=prev+current
				  
                  if next%2==0{
                  
                      sum+=next
                  }
                  prev=current
                  current=next
				}
        println("the sum of the even  valued terms is",sum)
}