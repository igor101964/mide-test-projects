#include <iostream>
#include <vector>
int main() {
    std::cout << "Hello from C++!" << std::endl;
    std::vector<int> v = {1, 2, 3};
    for (auto x : v)
        std::cout << "  item: " << x << std::endl;
    return 0;
}
