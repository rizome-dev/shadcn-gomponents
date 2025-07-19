#!/bin/bash
set -euo pipefail

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Test configuration
BUILD_DIR="./build"
TEST_DIR="./test-e2e"
SHARED_LIB=""

# Detect platform and set shared library name
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    SHARED_LIB="$BUILD_DIR/signer-amd64.so"
elif [[ "$OSTYPE" == "darwin"* ]]; then
    SHARED_LIB="$BUILD_DIR/signer-arm64.dylib"
else
    echo -e "${RED}Unsupported platform: $OSTYPE${NC}"
    exit 1
fi

echo -e "${YELLOW}Running E2E tests for shared library...${NC}"

# Check if shared library exists
if [ ! -f "$SHARED_LIB" ]; then
    echo -e "${RED}Error: Shared library not found at $SHARED_LIB${NC}"
    echo "Please build the shared library first using 'just build' or 'make sharedlib-all'"
    exit 1
fi

# Create test directory
mkdir -p "$TEST_DIR"

# Create C test program
cat > "$TEST_DIR/test_sharedlib.c" << 'EOF'
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <dlfcn.h>

// Function pointers for shared library functions
typedef char* (*GetVersion_func)();
typedef char* (*GetPlatform_func)();
typedef char* (*Sign_func)(const char*);
typedef int (*Verify_func)(const char*, const char*);
typedef void (*FreeString_func)(char*);

int main() {
    void *handle;
    char *error;
    int test_passed = 1;

    // Load the shared library
    #ifdef __APPLE__
        handle = dlopen("./build/signer-arm64.dylib", RTLD_LAZY);
    #else
        handle = dlopen("./build/signer-amd64.so", RTLD_LAZY);
    #endif

    if (!handle) {
        fprintf(stderr, "Cannot open shared library: %s\n", dlerror());
        return 1;
    }

    // Clear any existing error
    dlerror();

    // Load functions
    GetVersion_func GetVersion = (GetVersion_func) dlsym(handle, "GetVersion");
    GetPlatform_func GetPlatform = (GetPlatform_func) dlsym(handle, "GetPlatform");
    Sign_func Sign = (Sign_func) dlsym(handle, "Sign");
    Verify_func Verify = (Verify_func) dlsym(handle, "Verify");
    FreeString_func FreeString = (FreeString_func) dlsym(handle, "FreeString");

    if ((error = dlerror()) != NULL) {
        fprintf(stderr, "Error loading functions: %s\n", error);
        dlclose(handle);
        return 1;
    }

    // Test GetVersion
    printf("Testing GetVersion...\n");
    char *version = GetVersion();
    printf("Version: %s\n", version);
    if (strcmp(version, "1.0.0") != 0) {
        fprintf(stderr, "ERROR: Expected version '1.0.0', got '%s'\n", version);
        test_passed = 0;
    }
    FreeString(version);

    // Test GetPlatform
    printf("\nTesting GetPlatform...\n");
    char *platform = GetPlatform();
    printf("Platform: %s\n", platform);
    FreeString(platform);

    // Test Sign and Verify
    printf("\nTesting Sign and Verify...\n");
    const char *test_data = "Hello, World!";
    char *signature = Sign(test_data);
    printf("Data: %s\n", test_data);
    printf("Signature: %s\n", signature);

    // Verify with correct signature
    int verify_result = Verify(test_data, signature);
    if (verify_result != 1) {
        fprintf(stderr, "ERROR: Verification failed for correct signature\n");
        test_passed = 0;
    } else {
        printf("Verification: PASSED\n");
    }

    // Verify with incorrect signature
    int verify_fail = Verify(test_data, "invalid_signature");
    if (verify_fail != 0) {
        fprintf(stderr, "ERROR: Verification passed for incorrect signature\n");
        test_passed = 0;
    } else {
        printf("Invalid signature test: PASSED\n");
    }

    FreeString(signature);

    // Test NULL handling
    printf("\nTesting NULL handling...\n");
    char *null_sign = Sign(NULL);
    if (strlen(null_sign) != 0) {
        fprintf(stderr, "ERROR: Sign(NULL) should return empty string\n");
        test_passed = 0;
    }
    FreeString(null_sign);

    int null_verify = Verify(NULL, NULL);
    if (null_verify != 0) {
        fprintf(stderr, "ERROR: Verify(NULL, NULL) should return 0\n");
        test_passed = 0;
    }

    // Close the library
    dlclose(handle);

    if (test_passed) {
        printf("\n✅ All tests passed!\n");
        return 0;
    } else {
        printf("\n❌ Some tests failed!\n");
        return 1;
    }
}
EOF

# Create Python test program
cat > "$TEST_DIR/test_sharedlib.py" << 'EOF'
#!/usr/bin/env python3
import ctypes
import platform
import sys
import os

def main():
    # Load the shared library
    if platform.system() == "Darwin":
        lib_path = "./build/signer-arm64.dylib"
    else:
        lib_path = "./build/signer-amd64.so"
    
    if not os.path.exists(lib_path):
        print(f"Error: Shared library not found at {lib_path}")
        return 1
    
    try:
        lib = ctypes.CDLL(lib_path)
    except Exception as e:
        print(f"Error loading shared library: {e}")
        return 1
    
    # Define function signatures
    lib.GetVersion.restype = ctypes.c_char_p
    lib.GetPlatform.restype = ctypes.c_char_p
    lib.Sign.argtypes = [ctypes.c_char_p]
    lib.Sign.restype = ctypes.c_char_p
    lib.Verify.argtypes = [ctypes.c_char_p, ctypes.c_char_p]
    lib.Verify.restype = ctypes.c_int
    lib.FreeString.argtypes = [ctypes.c_char_p]
    
    test_passed = True
    
    # Test GetVersion
    print("Testing GetVersion...")
    version = lib.GetVersion().decode('utf-8')
    print(f"Version: {version}")
    if version != "1.0.0":
        print(f"ERROR: Expected version '1.0.0', got '{version}'")
        test_passed = False
    
    # Test GetPlatform
    print("\nTesting GetPlatform...")
    platform_str = lib.GetPlatform().decode('utf-8')
    print(f"Platform: {platform_str}")
    
    # Test Sign and Verify
    print("\nTesting Sign and Verify...")
    test_data = b"Hello, World!"
    signature = lib.Sign(test_data).decode('utf-8')
    print(f"Data: {test_data.decode('utf-8')}")
    print(f"Signature: {signature}")
    
    # Verify with correct signature
    verify_result = lib.Verify(test_data, signature.encode('utf-8'))
    if verify_result != 1:
        print("ERROR: Verification failed for correct signature")
        test_passed = False
    else:
        print("Verification: PASSED")
    
    # Verify with incorrect signature
    verify_fail = lib.Verify(test_data, b"invalid_signature")
    if verify_fail != 0:
        print("ERROR: Verification passed for incorrect signature")
        test_passed = False
    else:
        print("Invalid signature test: PASSED")
    
    if test_passed:
        print("\n✅ All Python tests passed!")
        return 0
    else:
        print("\n❌ Some Python tests failed!")
        return 1

if __name__ == "__main__":
    sys.exit(main())
EOF

# Run C tests
echo -e "\n${YELLOW}Running C tests...${NC}"
gcc -o "$TEST_DIR/test_sharedlib" "$TEST_DIR/test_sharedlib.c" -ldl
if "$TEST_DIR/test_sharedlib"; then
    echo -e "${GREEN}C tests passed!${NC}"
else
    echo -e "${RED}C tests failed!${NC}"
    exit 1
fi

# Run Python tests if Python is available
if command -v python3 &> /dev/null; then
    echo -e "\n${YELLOW}Running Python tests...${NC}"
    chmod +x "$TEST_DIR/test_sharedlib.py"
    if python3 "$TEST_DIR/test_sharedlib.py"; then
        echo -e "${GREEN}Python tests passed!${NC}"
    else
        echo -e "${RED}Python tests failed!${NC}"
        exit 1
    fi
else
    echo -e "${YELLOW}Python not found, skipping Python tests${NC}"
fi

# Clean up
rm -rf "$TEST_DIR"

echo -e "\n${GREEN}✅ All E2E tests completed successfully!${NC}"