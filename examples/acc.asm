
mov r0, #0

loop:
	cmp r0, #101
	beq exit
	sys #0
	sys #4 
	add r0, r0, #1
	b loop


exit:
	HALT
