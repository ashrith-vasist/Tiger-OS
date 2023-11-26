# Setting the 16-bit mode and define the global label "tiger"
.code16
.global tiger

tiger:
    mov $message, %si
    mov $0xe, %ah

print_message: 
    lodsb
    cmp $0, %al
    je finish
    int $0x10
    jmp print_message

finish:
    hlt

message: .asciz "Tiger OS"

.fill 510-(.-tiger),1,0
.word 0xaa55
