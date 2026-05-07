mov r0, #1
mov r1, #0
loop:
    cmp r1, #32
    beq exit
    lsl r0, r0, #1
    add r1, r1, #1
    b loop
exit:
    halt
