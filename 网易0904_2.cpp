#include <algorithm>
#include <cmath>
#include <cstdio>
#include <cstring>
#include <iostream>
using namespace std;

int n;

int gcdMath(int x, int y) {
    return x % y == 0 ? y : gcdMath(y, x % y);
}

int main() {
    cin >> n;

    if (n <= 4) {
        cout << n << ' ' << 1 << endl;
        return 0;
    }

    int i, last = 0;
    for (i = 2; i < n; i++) {
        double tmp = (double)n / i;
        double now = 1.0;
        for (int j = 1; j <= i; j++) {
            now *= tmp;
        }
        if (now < last) {
            i--;
            break;
        }
        else {
            last = now;
        }
    }

    int up = 1, down = 1;
    for (int k = 1; k <= i; k++) {
        up *= n;
    }
    for (int k = 1; k <= i; k++) {
        down *= i;
    }

    int gcd = gcdMath(up, down);
    up /= gcd;
    down /= gcd;

    cout << up << ' ' << down << endl;

    return 0;
}
