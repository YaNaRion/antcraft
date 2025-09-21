# Makefile for C++ Project with Subdirectories + pre-generated Protobuf in proto/

# ============== Configuration ==============
CXX = g++
CXXFLAGS = -Wall -Wextra -g -std=c++23

# Include paths: manually written + proto-generated headers
CXXFLAGS += -Iinclude -Iproto

# Use pkg-config for protobuf if available
PROTOBUF_CFLAGS := $(shell pkg-config --cflags protobuf 2>/dev/null)
PROTOBUF_LIBS   := $(shell pkg-config --libs protobuf 2>/dev/null)

ifeq ($(strip $(PROTOBUF_LIBS)),)
    PROTOBUF_CFLAGS =
    PROTOBUF_LIBS = -lprotobuf
endif

CXXFLAGS += $(PROTOBUF_CFLAGS)

# Directory structure
SRC_DIR = src
PROTO_DIR = proto         # Directory containing *.pb.cc / *.pb.h
BUILD_DIR = build
BIN_DIR = bin

# Raylib settings
RAYLIB_INCLUDE_DIR = ./raylib-5.5_linux_amd64/include
RAYLIB_LIB_DIR     = ./raylib-5.5_linux_amd64/lib/

TARGET = myprogram

RAYLIB_STATIC = $(RAYLIB_LIB_DIR)/libraylib.a
LIBS = $(RAYLIB_STATIC) -lm -ldl -lpthread $(PROTOBUF_LIBS)

UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
    LIBS += -lX11 -lGL -lopenal
endif
ifeq ($(UNAME_S),Darwin)
    LIBS += -framework OpenGL -framework Cocoa -framework IOKit -framework CoreAudio -framework CoreVideo
endif

CXXFLAGS += -I$(RAYLIB_INCLUDE_DIR)

# ============== Source file discovery ==============
SRC_SRCS := $(shell find $(SRC_DIR) -name '*.cpp' -o -name '*.cc')
PROTO_SRCS := $(shell find $(PROTO_DIR) -name '*.pb.cc')
# Also .cc files in proto (if you have .cc that are not .pb.cc) optionally:
PROTO_CC_SRCS := $(shell find $(PROTO_DIR) -name '*.cc' -not -name '*.pb.cc')

MAIN_SRC := main.cpp

# Object files
OBJS := \
    $(patsubst $(SRC_DIR)/%.cpp,$(BUILD_DIR)/$(SRC_DIR)/%.o,$(filter %.cpp,$(SRC_SRCS))) \
    $(patsubst $(SRC_DIR)/%.cc,$(BUILD_DIR)/$(SRC_DIR)/%.o,$(filter %.cc,$(SRC_SRCS))) \
    $(patsubst $(PROTO_DIR)/%.pb.cc,$(BUILD_DIR)/$(PROTO_DIR)/%.pb.o,$(PROTO_SRCS)) \
    $(patsubst $(PROTO_DIR)/%.cc,$(BUILD_DIR)/$(PROTO_DIR)/%.o,$(PROTO_CC_SRCS))

MAIN_OBJ := $(BUILD_DIR)/$(MAIN_SRC:.cpp=.o)

OBJ_DIRS := $(sort $(dir $(OBJS)) $(dir $(MAIN_OBJ)))

# ============== Build Rules ==============
all: $(BIN_DIR)/$(TARGET)

$(BIN_DIR)/$(TARGET): $(MAIN_OBJ) $(OBJS) $(RAYLIB_STATIC)
	mkdir -p $(BIN_DIR)
	$(CXX) $(CXXFLAGS) $(MAIN_OBJ) $(OBJS) -o $@ $(LIBS)

$(BUILD_DIR)/%.o: %.cpp | $(OBJ_DIRS)
	mkdir -p $(dir $@)
	$(CXX) $(CXXFLAGS) -c $< -o $@

$(BUILD_DIR)/$(SRC_DIR)/%.o: $(SRC_DIR)/%.cpp | $(OBJ_DIRS)
	mkdir -p $(dir $@)
	$(CXX) $(CXXFLAGS) -c $< -o $@

$(BUILD_DIR)/$(SRC_DIR)/%.o: $(SRC_DIR)/%.cc | $(OBJ_DIRS)
	mkdir -p $(dir $@)
	$(CXX) $(CXXFLAGS) -c $< -o $@

$(BUILD_DIR)/$(PROTO_DIR)/%.pb.o: $(PROTO_DIR)/%.pb.cc | $(OBJ_DIRS)
	mkdir -p $(dir $@)
	$(CXX) $(CXXFLAGS) -c $< -o $@

$(BUILD_DIR)/$(PROTO_DIR)/%.o: $(PROTO_DIR)/%.cc | $(OBJ_DIRS)
	mkdir -p $(dir $@)
	$(CXX) $(CXXFLAGS) -c $< -o $@

$(BUILD_DIR)/%.o: %.cpp | $(OBJ_DIRS)
	mkdir -p $(dir $@)
	$(CXX) $(CXXFLAGS) -c $< -o $@

# Create directories
$(OBJ_DIRS):
	mkdir -p $@

# Clean
clean:
	rm -rf $(BUILD_DIR) $(BIN_DIR)

run: all
	./$(BIN_DIR)/$(TARGET)

debug: all
	gdb --args ./$(BIN_DIR)/$(TARGET)

.PHONY: all clean run debug

