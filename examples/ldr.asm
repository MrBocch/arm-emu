// ldr/str indirectly
mov r1, #1
mov r0, #0
// overwriting code here
str r0, [r1]
ldr r1, [r0]


halt
