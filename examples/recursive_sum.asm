; input n +1 if you want, sum of n
; ans stored to r2
mov r1, #11
mov r2, #0 
bl sum
halt 

sum:
	cmp r1, #1
	beq sum_base
	push {r1, lr}
	sub r1, r1, #1
	bl sum              ; i think doing this is cheating.
						; and prob not correct.
						; once it reaches base case, it pops lr
						; and goes back to the base case;until it
						; goes back to line 5, is this a proper technique? (prob not)
						; i need a style guide on how to not write, bad assembly 
sum_base:
	add r2, r2, r1
	pop {r1, lr}
	ret 

; (this is what deepseek gave me)
; Input: n in R0
; Output: sum in R0


MOV R0, #5          
BL recursive_sum
HALT

recursive_sum:
    PUSH {R1, LR}       ; Save R1 and Link Register
    
    CMP R0, #1          ; Base case: if n <= 1
    BEQ base_case       ; theres no branch if less than or equal to unless BEQ ...\nBLT...
    
    MOV R1, R0          ; Save current n in R1
    SUB R0, R0, #1      ; n = n - 1
    BL recursive_sum    ; Recursive call: sum(n-1)
    
    ADD R0, R0, R1      ; Return sum(n-1) + n
    
    POP {R1, PC}        ; Restore R1 and return

base_case:
    MOV R0, #1          ; sum(1) = 1
    POP {R1, PC}        ; Restore R1 and return

