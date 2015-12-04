# lazycombCompilerGo
This is a simple compiler for lazyComb written in golang.

### Compiling .lazy programs
The compilation process is made in two steps:
  - Compile .lazy file into .c file
  - Compile .c file to generate the executable

To execute both steps just run 

    $ ./compileAndExecute.sh lazyComb_examples/example1.lazy

As seen in the above example, all samples in lazy are located at lazyCom_examples 

### Cleaning *.o *.d and binary files generate

To remove all files generated during compilation, execute

    $ make clean
    
### Misc
  - On the first step of the compilation, the output file is called out.c. It is generated on vm/
  - The executable will be created at the project root
