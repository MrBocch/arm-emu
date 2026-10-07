# arm-emu

A [armlite](https://peterhigginson.co.uk/ARMlite/) emulator.

## example

Here is an example of a program that prints out a multiplication table. 

[showcase](https://github.com/user-attachments/assets/9bab4dbe-cf60-47bc-8aef-318b0f6faa94)

Using a 'sys' is a work around for not knowing what interupts are, and its just ugly.
Im rewriting it.

## Notes

Comments 

```C
// this is a comment
;; this is also a comment
<!-- /* Multiline -->
 * comments also.
*/
```

No doubly nesting comments. 

## Implementation (✅/❌)

| opcode   | implemented | 
|----------|-------------|
| mov      |     ✅      |
| add      |     ✅      |
| sub      |     ✅      | 

....
