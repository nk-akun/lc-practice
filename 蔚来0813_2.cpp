#include <algorithm>
#include <cstdio>
#include <cstring>
#include <iostream>
using namespace std;

int main() {

    int t, a, b;
    cin >> t;
    while (t--) {
        cin >> a >> b;
        cout << b + a * b << endl;
    }
}