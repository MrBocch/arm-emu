

mov r0, #0
mov r1, #0
mov r2, #40
mov r3, #end_of_code
input_array:
	ldr r4, .InputNum
	str r4, [r3 + r0]
	add r0, r0, #4
	cmp r0, r2
	blt input_array
	mov r0, #0
	b selection_sort

selection_sort:
	cmp r0, r2
	beq break
	cmp r1, r2
	beq reset_j

	ldr r11, [r3 + r0]
	ldr r12, [r3 + r1]
	cmp r11, r12
	bgt swap
after_swap:
	add r1, r1, #4
	b selection_sort

reset_j:
	add r0, r0, #4
	mov r1, r0
	b selection_sort

swap:
	str r11, [r3 + r1]
	str r12, [r3 + r0]
	b after_swap

break:
	halt

before_sort: .asciz "Before sorting\n"
after_sort: .asciz  "\nAfter sorting\n"

end_of_code:
