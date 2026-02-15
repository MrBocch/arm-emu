mov r1, #array          ; You cant a[0], thats an abstraction     
mov r3, #last           ; dealing with raw addresses here. wow.
mov r4, #array

; r1 = a[low]
; r2 = a[mid]
; r3 = a[high]

mov r12, #enter         
str r12, .WriteString   ; print("Enter number: ")
ldr r12, .InputNum       ; r12 = input() 

binary_search:

	add r2, r1, r3
	lsr r2, r2, #3
	lsl r2, r2, #2          ; mid; oh wow! 

	ldr r11, [r2]
	cmp r11, r12
	beq found

	cmp r1, r3
	beq not_found
	bgt not_found

	cmp r11, r12
	bgt greater_than
	blt less_than

found:
	mov r12, #found_s
	str r12, .WriteString
	sub r2, r2, #array
	lsr r2, r2, #2           ; divide by 4
	str r2, .WriteSignedNum
	halt

greater_than:
	; high = mid -1
	sub r3, r3, #4
	b binary_search

less_than:
	; low = mid+1
	add r1, r1, #4
	b binary_search

not_found:
	mov r11, #not_found_s
	str r11, .WriteString
	str r12, .WriteSignedNum
	halt

enter: .asciz "Enter number: "
found_s: .asciz "Index is: "
not_found_s: .asciz "Did not find: "

array: 1
2
3
4
5
6
7
8
9
10
last: 100
