#include "test.hpp"
#include <iostream>

void print_hello_from_test() {
  Test test = Test();
  test.print_hello_from_test();
}

void Test::print_hello_from_test() { std::cout << "TEST\n"; }
