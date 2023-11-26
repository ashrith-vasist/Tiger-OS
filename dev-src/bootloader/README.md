- Bootloader

    - Bootloader

        - boot.s

            - first line is to define 16-bit mode

            - next declare the global mode with the defined label visible to the linker.

            -  define label, give a jmp command to create an infinite loop. fill 510 bytes of zeroes
            and the last two bytes are the magic numbers. (add zeros to 510 bytes, .fill, count, size value)
            ie . `.fill 510-(-.label),1,0`

            - .word 0xaa55 (little endian)
        
        - Compile with GNU Assembler (GAS).
        
        `as -o boot.o boot.s`

        - Use the linker to produce the final code.
        `ld -o boot.bin --oformat binary -e tiger -Ttext 0x7c00 boot.o`

        - Setup Qemu on System, run the binary file in Qemu.
        `qemu-system-x86_64 -drive file=boot.bin,index=0,media=disk,format=raw`