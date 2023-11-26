.code16
.global tiger

tiger:
    jmp tiger

.fill 510-(.-tiger),1,0
.word 0xaa55
