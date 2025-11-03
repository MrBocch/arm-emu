; input 10 numbers
; selection sorts them
; make a print array function how?
; parameterize everything

mov r0, #0          ; i
mov r1, #0          ; j
mov r2, #40         ; size = (sizeof(i32) * n) ; n = 10 ; sizeof(int) = 4 ; word*n
mov r3, #end_of_code; &arr

input_array:
	ldr r4, .InputNum
	str r4, [r3 + r0]
	add r0, r0, #4              ; add by word size
	cmp r0, r2                  ; size of array
	blt input_array
	mov r0, #0
	b selection_sort

selection_sort:
	cmp r0, r2               
	beq break                   ; break if i == size 
	cmp r1, r2
	beq reset_j                 ; if j == size { i++; j = i }

	ldr r11, [r3 + r0]          ; a[i]
	ldr r12, [r3 + r1]          ; a[j]
	cmp r11, r12 	
	bgt swap 
after_swap: 
	add r1, r1, #4
	b selection_sort 

reset_j:
	add r0, r0, #4             ; i++
	mov r1, r0                 ; j = i
	b selection_sort

swap:
	str r11, [r3 + r1]
	str r12, [r3 + r0]
	b after_swap
 
break:
	halt

before_sort: .asciz "Before sorting\n"
after_sort: .asciz  "\nAfter sorting\n" 

;; the array will be "allocated" after all the code
;; the ideal location.
;; but, what if you had to arrays? if they are static, no problem arr2 = (arr1*size)+1 ?
;; 
end_of_code:

