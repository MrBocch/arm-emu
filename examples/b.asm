
mov r0, #0
mov r1, #10
loop:
	add r0, r0, #1
	// cmp r0, #10
	cmp r0, r1
	blt loop
	beq exit
exit:
    halt
