package main

func lef ( []int ) int {
	
	leftsum := 0
	zongshu := 0
	for i :=0 ;  i < len(lef);i ++{
	 zongshu += lef[i]
	if  leftsum == zongshu{
		return  0
		for  y := 0 ; y < len(lef) : y++[
			leftsum += lef[y]
		 if	leftsum == zongshu - lef[y] -leftsum {
			return lef[y]
		 } 

		]
}
}
