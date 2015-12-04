INCLUDE_PATH="./vm/"

C_FILES := $(wildcard ./vm/*.c)
O_FILES := $(C_FILES:.c=.o)
C_DEPS := $(C_FILES:.c=.d)

ifneq ($(strip $(C_DEPS)),)
-include $(C_DEPS)
endif

# Compile all files from vm/
vm/%.o: vm/%.c
	@echo ' '
	@echo '[INFO] Building file: $<'
	@echo '[INFO] Invoking: Cross GCC Compiler'
	gcc -I$(INCLUDE_PATH) -O0 -g3 -Wall -c -fmessage-length=0 -MMD -MP -MF"$(@:%.o=%.d)" -MT"$(@)" -o "$@" "$<"
	@echo '[INFO] Finished building: $<'
	@echo ' '

# Run Target
run: compiler
	@echo '[INFO] Executing generated binary ./lazy'
	./lazy

# Compiler Target
compiler: $(O_FILES)
	@echo '[INFO] Building target: $@'
	@echo '[INFO] Invoking: Cross GCC Linker'
	gcc  -o "lazy" $(O_FILES)
	@echo '[INFO] Finished building target: $@'
	@echo ' '
	@echo ' '

# Clean Target
clean:
	-$(RM) $(O_FILES) $(C_DEPS) lazy
	-@echo ' '

.PHONY: all clean
