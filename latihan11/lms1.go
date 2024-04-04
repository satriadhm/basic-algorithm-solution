package main 

type tabInt [26]int

func pecahArray( T tabInt, n int, puluhan, satuan *tabInt, np,ns= *int){	
	for i := 0; i < n; i++{
		if T[i]%10==0{
			puluhan[*np]= T[i]
			*np++
		}else if T[i]%10==T[i]{
			satuan[*ns]=0
			*ns++
		}
	}
}

