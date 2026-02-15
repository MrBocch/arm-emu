; r0 for printing
; R1 for storing tries
; R2 for storing guess
; r12 contains the number
setup:
    ldr r12, .Random
    and r12, r12, #127
    add r12, r12, #1
    mov r1, #0
    mov r0, #guessNum
    str r0, .WriteString
    B gameLoop
gameLoop:
      add r1, r1, #1
      cmp r1, #5

      BEQ lostGame
      ldr r2, .InputNum
      cmp r2, r12
      BGT pTooHigh
      BLT pTooLow
      BEQ wonGame
wonGame:
      mov r0, #wonStr
      str r0, .WriteString
      HALT
lostGame:
      mov r0, #lostStr
      str r0, .WriteString
      mov r0, #theNumStr
      str r0, .WriteString
      str r12, .WriteUnsignedNum
      HALT
pTooHigh:
      mov r0, #tooHigh
      str r0, .WriteString
      B gameLoop
pTooLow:
      mov r0, #tooLow
      str r0, .WriteString
      B gameLoop

guessNum: .ASCIZ "Guess a number between 1 and 128\n"
tooLow: .ASCIZ "Too low!\n"
tooHigh: .ASCIZ "Too high!\n"
wonStr: .ASCIZ "You won!\n"
lostStr: .ASCIZ "You lost!\n"
theNumStr: .ASCIZ "The number was: "

1|      mov r12, #0x0
2|      mov r0, #.Pixel0
3|      mov r1, #0xFF0000
4|loop:
5|      str r1, [r0 + r12]
6|      adds r12, r12, #0x4
7|      cmp r12, #0xC00
8|      beq reset
9|      b loop
10|reset:
11|      mov r12, #0x0
12|      add r1, r1, #0xFF
13|      cmp r1, #0x1000000
14|      BGT color
15|      b loop
16|color:
17|      mov r1, #0xFF0000
18|      b loop
