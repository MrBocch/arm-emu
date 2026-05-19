
// print "Multiplication table"
// output M
mov r0, #77
sys #3
// output u
mov r0, #117
sys #3
// output l
mov r0, #108
sys #3
// output t
mov r0, #116
sys #3
// output i
mov r0, #105
sys #3
// output p
mov r0, #112
sys #3
// output l
mov r0, #108
sys #3
// output i
mov r0, #105
sys #3
// output c
mov r0, #99
sys #3
// output a
mov r0, #97
sys #3
// output t
mov r0, #116
sys #3
// output i
mov r0, #105
sys #3
// output o
mov r0, #111
sys #3
// output n
mov r0, #110
sys #3
// output
mov r0, #32
sys #3
// output T
mov r0, #84
sys #3
// output a
mov r0, #97
sys #3
// output b
mov r0, #98
sys #3
// output l
mov r0, #108
sys #3
// output e
mov r0, #101
sys #3
sys #4

mov r1, #1 // i
mov r2, #1 // j 
mov r3, #0 // res

// for i = range 1, 12
outer:
	cmp r1, #13
	beq exit
inner:
	cmp r2, #13
	beq inner_exit
	add r3, r3, r1
	// inner printing logic 
	// print i 
	mov r0, r1
	sys #1
	// print " "
	mov r0, #32
	sys #3
	// print "x"
	mov r0, #120
	sys #3
	// print " "
	mov r0, #32
	sys #3

	// print j 
	mov r0, r2
	sys #1 

	// print " "
	mov r0, #32
	sys #3
	// print "="
	mov r0, #61
	sys #3
	// print " "
	mov r0, #32
	sys #3

	// print res 
	mov r0, r3
	sys #1


	// newline 
	sys #4 
	
	add r2, r2, #1
	b inner
inner_exit:
	// newlien 
	mov r3, #0
	sys #4 
	mov r2, #1
	add r1, r1, #1
	b outer


exit:
	HALT
