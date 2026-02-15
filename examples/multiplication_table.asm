
MOV r1,  #1        ; i
MOV r2,  #1        ; j
MOV r3,  #0        ; sum
mov r0, #linez     ; -----------
B loop

loop:
    CMP r1, #13
    BEQ end
    ; sum
    ADD r3, r3, r1

    STR r1, .WriteSignedNum

    PUSH {R0}
    MOV r0, #x
    BL printString

    STR r2, .WriteSignedNum

    PUSH {r0}
    MOV r0, #equals
    BL printString
    ; str sum
    STR r3, .WriteSignedNum

    PUSH {r0}
    MOV r0, #nline
    BL printString
    ADD r2, r2, #1

    CMP r2, #13
    BEQ setinnerloop
    b loop

setinnerloop:
    ADD r1, r1, #1
    MOV r2, #1
    MOV r3, #0
    STR r0, .WriteString
    B loop

end:
    HALT

printString:
  str r0, .WriteString
  pop {r0}
  ret
; i x j = (i*j)
x: .ASCIZ " x "
equals: .ASCIZ " = "
nline: .ASCIZ "\n"
linez: .ASCIZ "--------------\n"
