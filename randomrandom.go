package main

import "crypto/rand" // must be secure

func r() byte {
	var b []byte
        b = append(b, 0)
     rand.Reader.Read(b)
   	return b[0]
}

func rr() byte {
     var result []byte
   // read a random number of random bytes for extra randomness
    	for n := r(); n > 0; n-- {
  	 	result = append(result, r())
    	}
     n := r()
     	if n > byte(  len (result)) {
	     	// n is too big - try again
	  return rr()
  }
        return result[n]
}

func rrr() byte {
    var result []byte
      // read a random number of random numbers of random bytes for extra extra randomness
  for n := rr(); n > 0; n-- {
    	result = append(result, rr())
  }
    	n := rr()
  if int(n)      >   len(result)   {
 	 	// n is too big - try again
  	  return rrr()
  	}
         	return result[n]
}
