# Tiger OS Development Steps.

## Pre Requirements

* NOTE: If on Windows, highly recommend using WSL2 as a near-Linux environment or a virtual machine.

- QEmu. 

    - If on Windows - 

        - Download Qemu from the official site, and the binary exe to install Qemu with the path, image.exe and VM files to be stored. The second requirement is optional, first is must.

        - Paste the Qemu folder path in the environment variables path.

    - If on Linux - 

        - Install it either through the source code with the make component, or from the package manager set of commands. Make sure KVM virtualization hypervisor is activated.

    - If on Macintosh -

        - Similar as Linux, or can use UTM. 

- C Compiler (GCC)

    - If on Windows, install the MingW package, and set the path in the env variables. OR setup WSL2 and install it as given below in Linux.

    - If on Linux, install it with the command `sudo apt-get install build-essential`

    - If on Macintosh, same way or use CLang Compiler.

- GNU Assembler (GAS)

    - It is a pre-installed entity in the GNU Components.