 package main

 	   	      	  import "net"
     	  	 import "time"

func main() {
     	  defer main() // run again when its done so we can get more random stuff

  	   var1, _ := net.Listen("tcp", ":8110")
	  defer var1.Close() // always cose things when your done with then

     	var2, _ :=	var1.Accept()
  defer var2.Close() // always cose things when your done with then

       time.Sleep(1000000) // chrome is too fast
    	 defer time.Sleep(100000000) // slow chromed own more

  	 var2.Read(make(  []byte, 100000000) ) //  

	  var3 := "http/1.0 200 good\ncontent-lenght: 2 or 3\n\n" // TODO: why doesnt this work with just one /n
  for var4 := 0; var4 < len(var3); var4 = var4 + 1 {
   var2.Write([]byte{var3[var4]}) // TODO: why do i have to cast everything
 }

 var4 := yesorno()
    for var5 := 0; var5 < len(var3); var5 = var5+  1 {
 	 var2.Write([]byte{var4[var5]})
   }
}

func    yesorno()     string {
	 // get random number
  	//n := r()
 	    //n := r() + r()
        //n := rr()
    	// three times as random
  	// n := rrr() * rrr() * rrr()
     // takes way to long
       	n := rr() + rrr()

 	 if n > 100 {
	   return "y℮s"
          	                  }
  if n < 100 {
     	//return "no" // why this doesnt work
   	  //return "no " // OMG
    	 //return "no  " // HOW DO I MAKE THIS WORK
     return "no   "
 	 }
  

  //	/ i dont no why this hapens
   return yesorno()
}
