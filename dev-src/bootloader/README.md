- Bootloader

    - Bootloader

        - boot.s

            - first line is to define 16-bit mode

            - next declare the global mode with the defined label visible to the linker.

            -  define label, give a jmp command to create an infinite loop. fill 510 bytes of zeroes
            and the last two bytes are the magic numbers. (add zeros to 510 bytes, .fill, count, size value)
            ie . `.fill 510-(-.label),1,0`

            - .word 0xaa55 (little endian)

        ```
        # Setting the 16-bit mode and define the global label "tiger"
            .code16
            .global tiger

            tiger:
                jmp tiger

            .fill 510-(.-tiger),1,0
            .word 0xaa55
        ```

        This is a basic bootloader component.

    - Compile with GNU Assembler (GAS).
        
        `as -o boot.o boot.s`

    - Use the linker to produce the final code.
        `ld -o boot.bin --oformat binary -e tiger -Ttext 0x7c00 boot.o`

    - Setup Qemu on System, run the binary file in Qemu.
        `qemu-system-x86_64 -drive file=boot.bin,index=0,media=disk,format=raw`

        If the Windows setup is followed, and on WSL2 or Windows Terminal, use 
        `qemu-system-x86_64.exe -drive file=boot.bin,index=0,media=disk,format=raw`

- Generate your own patterns from here http://www.patorjk.com/software/taag/