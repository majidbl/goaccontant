package util

import (
  "unicode"
  "bytes"
  )
  
  
func IsIntFloat(s string) bool {
      var isNumber = true
      var pointFloat = 0
      if len(s) == 0 {
        isNumber = false
      }
	    for _, c := range s {
	        if !unicode.IsDigit(c){
	          if c == '.' && pointFloat < 1 {
	            isNumber = true
	            pointFloat++
	          }else{
	            isNumber = false
	          }
	        }
	        
	    }
	    return isNumber
}

func FloatPoint(sr string)(int,bool, bool){
  if !IsIntFloat(sr){
    return -1,false,false
  }
  for index, c := range sr{
    if c == '.'{
      return index, true, false
    }
  }
    return -1, false, true
}

func FormatAmount(sr string)(string,bool,bool){
  var h bytes.Buffer
  n, f, i := FloatPoint(sr)
  var count = 0
  if f{
    var rev bytes.Buffer
    if n > 3{
      for i := n-1 ; i > -1; i--{
        if count%3 == 0 && count != 0{
          h.WriteString(",")
        }
        h.WriteString(string(sr[i]))
        count++
      }
      com := h.String()
      for d := len(com)-1; d > -1; d--{
        rev.WriteString(string(com[d]))
      }
      for t := n ; t < len(sr); t++{
        rev.WriteString(string(sr[t]))
      }
      //fmt.Println(rev.String())
    return rev.String(),true,false
    }else{
      return sr, true,false
    }
  }
    // This part excute when input in integet
    if i{
      var rev bytes.Buffer
      if len(sr) > 3{
        for i := len(sr)-1 ; i > -1; i--{
          if count%3 == 0 && count != 0{
            h.WriteString(",")
          }
          h.WriteString(string(sr[i]))
          count++
        }
        com := h.String()
        for t := len(com)-1 ; t > -1; t--{
           rev.WriteString(string(com[t]))
        }
      return rev.String(),false,true
      }
      
    }
  
  return sr,false,false
}