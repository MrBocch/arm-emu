
mov r0, #msg
str r0, .WriteString
halt

msg: .ASCIZ "hello world!"
