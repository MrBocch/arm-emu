;; i dont like all the magic number stuff
;; also, I had to painstakingly try to find the
;; last pixel on bottom right, instead of just calculating it. very amateur


      mov r12, #0x0
      mov r0, #.Pixel0
      mov r1, #0xFF0000
loop:
      str r1, [r0 + r12]
      adds r12, r12, #0x4
      cmp r12, #0xC00
      beq reset
      b loop
reset:
      mov r12, #0x0
      add r1, r1, #0xFF
      cmp r1, #0x1000000
      BGT color
      b loop
color:
      mov r1, #0xFF0000
      b loop
