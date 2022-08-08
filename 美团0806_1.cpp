#include <algorithm>
#include <cmath>
#include <cstdio>
#include <cstring>
#include <iostream>
#include <vector>
using namespace std;

long long x, y;

int main() {
    int t;
    cin >> t;
    while (t--) {
        cin >> x >> y;
        cout << min(min(x, y), (x + y) / 3) << endl;
    }
}