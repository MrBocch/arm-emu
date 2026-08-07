# arm-emu Assembler

## Comments

Nothing special about comments.

```asm
;  single line comment
// single line comment

/* multi
 * line
 * comment
*/
```

Note on nesting comments. You can't nest multiple multiline comments
together.
  
```asm
/* We can
 * // nest comments like this
*/

// or /* like */ this

// [DO NOT DO THIS!]
/* But not like
  /* this */
*/
```

## Case Sensitivity

Instructions and registers can be written upper or lowercase but can't be mixed.

```asm
mov R0, #41
ADD r0, r0, #1
```

Labels are case sensitive.

## Labels

Labels are case sensitive.

Name collisions result in a error. 

They are defined like this
```asm
loop:
  add r0, r0, #1
  cmp r0, #100
  bne loop
  
```

## Assembly Directives 

Assembly directives are instructions for the assembler.
