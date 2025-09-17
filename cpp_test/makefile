# Makefile for C++ Project with Subdirectories

# ============== Configuration ==============
# Compiler settings
CXX = g++
CXXFLAGS = -Wall -Wextra -g -std=c++23 -I. -Iinclude
LDFLAGS = 

# Directory structure
SRC_DIR = src
INCLUDE_DIR = include
BUILD_DIR = build
BIN_DIR = bin

# Target executable name
TARGET = myprogram

# Additional libraries (e.g., -lm for math)
LIBS = -lm -lssl -lcrypto

# ============== Automatic File Discovery ==============
# Find all source files recursively
SRCS = $(shell find $(SRC_DIR) -name '*.cpp')
# Main.cpp is treated separately since it's in root
MAIN_SRC = main.cpp

# Generate object file paths
OBJS = $(patsubst $(SRC_DIR)/%.cpp,$(BUILD_DIR)/%.o,$(SRCS))
MAIN_OBJ = $(BUILD_DIR)/$(MAIN_SRC:.cpp=.o)

# Add include paths
CXXFLAGS += -I$(INCLUDE_DIR)

# Add libraries to linker flags
LDFLAGS += $(LIBS)

# Get all directories needed for object files
OBJ_DIRS = $(sort $(dir $(OBJS)) $(dir $(MAIN_OBJ)))

# ============== Build Rules ==============
# Default target
all: $(BUILD_DIR) $(BIN_DIR) $(OBJ_DIRS) $(BIN_DIR)/$(TARGET)

# Create necessary directories
$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

$(BIN_DIR):
	mkdir -p $(BIN_DIR)

# Rule to create all object directories
$(OBJ_DIRS):
	mkdir -p $@

# Link all objects into executable
$(BIN_DIR)/$(TARGET): $(MAIN_OBJ) $(OBJS)
	$(CXX) $(CXXFLAGS) $^ -o $@ $(LDFLAGS)

# Compile main.cpp (special case since it's in root)
$(BUILD_DIR)/%.o: %.cpp | $(OBJ_DIRS)
	$(CXX) $(CXXFLAGS) -c $< -o $@

# Compile src files
$(BUILD_DIR)/%.o: $(SRC_DIR)/%.cpp | $(OBJ_DIRS)
	$(CXX) $(CXXFLAGS) -c $< -o $@

# ============== Utility Targets ==============
clean:
	rm -rf $(BUILD_DIR) $(BIN_DIR)

run: all
	./$(BIN_DIR)/$(TARGET)

debug: all
	gdb --args ./$(BIN_DIR)/$(TARGET)

.PHONY: all clean run debug

