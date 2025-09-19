# Makefile for C++ Project with Subdirectories

# ============== Configuration ==============
# Compiler settings
CXX = g++
CXXFLAGS = -Wall -Wextra -g -std=c++23 -I. -Iinclude

# Use pkg-config for protobuf if available
PROTOBUF_CFLAGS := $(shell pkg-config --cflags protobuf 2>/dev/null)
PROTOBUF_LIBS   := $(shell pkg-config --libs protobuf 2>/dev/null)

ifeq ($(strip $(PROTOBUF_LIBS)),)
    # Fallback if pkg-config not available
    PROTOBUF_CFLAGS =
    PROTOBUF_LIBS = -lprotobuf
endif

CXXFLAGS += $(PROTOBUF_CFLAGS)

# Directory structure
SRC_DIR = src
INCLUDE_DIR = include
BUILD_DIR = build
BIN_DIR = bin

# Raylib paths (adjust these if needed)
RAYLIB_INCLUDE_DIR = ./raylib-5.5_linux_amd64/include
RAYLIB_LIB_DIR = ./raylib-5.5_linux_amd64/lib/

# Target executable name
TARGET = myprogram

# Raylib static library and dependencies
RAYLIB_STATIC = $(RAYLIB_LIB_DIR)/libraylib.a
LIBS = $(RAYLIB_STATIC) -lm -ldl -lpthread $(PROTOBUF_LIBS)

# Platform-specific libraries
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
    LIBS += -lX11 -lGL -lopenal
endif
ifeq ($(UNAME_S),Darwin)
    LIBS += -framework OpenGL -framework Cocoa -framework IOKit -framework CoreAudio -framework CoreVideo
endif

# Add Raylib include path
CXXFLAGS += -I$(RAYLIB_INCLUDE_DIR)

# ============== Automatic File Discovery ==============
# Find all source files recursively (including .pb.cc in src/)
SRCS = $(shell find $(SRC_DIR) -name '*.cpp' -o -name '*.cc')
# Main.cpp is treated separately since it's in root
MAIN_SRC = main.cpp

# Generate object file paths
OBJS = $(patsubst $(SRC_DIR)/%.cpp,$(BUILD_DIR)/$(SRC_DIR)/%.o,$(filter %.cpp,$(SRCS))) \
       $(patsubst $(SRC_DIR)/%.cc,$(BUILD_DIR)/$(SRC_DIR)/%.o,$(filter %.cc,$(SRCS)))
MAIN_OBJ = $(BUILD_DIR)/$(MAIN_SRC:.cpp=.o)

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

# Link all objects into executable with Raylib + Protobuf
$(BIN_DIR)/$(TARGET): $(MAIN_OBJ) $(OBJS) $(RAYLIB_STATIC)
	$(CXX) $(CXXFLAGS) $(MAIN_OBJ) $(OBJS) -o $@ $(LIBS)

# Compile main.cpp
$(BUILD_DIR)/%.o: %.cpp | $(OBJ_DIRS)
	$(CXX) $(CXXFLAGS) -c $< -o $@

# Compile src files (.cpp and .cc, including .pb.cc)
$(BUILD_DIR)/$(SRC_DIR)/%.o: $(SRC_DIR)/%.cpp | $(OBJ_DIRS)
	$(CXX) $(CXXFLAGS) -c $< -o $@
$(BUILD_DIR)/$(SRC_DIR)/%.o: $(SRC_DIR)/%.cc | $(OBJ_DIRS)
	$(CXX) $(CXXFLAGS) -c $< -o $@

# Check if Raylib static library exists
$(RAYLIB_STATIC):
	@if [ ! -f "$(RAYLIB_STATIC)" ]; then \
		echo "Error: Raylib static library not found at $(RAYLIB_STATIC)"; \
		echo "Please verify the path or download Raylib"; \
		exit 1; \
	fi

# ============== Utility Targets ==============
clean:
	rm -rf $(BUILD_DIR) $(BIN_DIR)

run: all
	./$(BIN_DIR)/$(TARGET)

debug: all
	gdb --args ./$(BIN_DIR)/$(TARGET)

.PHONY: all clean run debug

