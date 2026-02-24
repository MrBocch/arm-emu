
mov r0, #0
loop:
	add r0, #1
	b loop
	halt // this will never halt
