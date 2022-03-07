#include <algorithm>
#include <cstdio>
#include <cstring>
#include <iostream>
#include <map>
#include <stack>
#include <vector>
using namespace std;
const int maxn = 1 * (1e5 + 5);

// 题意：正n边形内部可以连出多少个等腰三角形

long long n;

int main() {
    cin >> n;

    long long ans = 0;
    if (n % 2 == 1) {
        ans = (n + 2) / 4;
    }
    else {
        if (n % 4 == 0) {
            ans = n / 4 - 1;
        }
        else {
            ans = n / 4;
        }
    }
    ans = ans * n;

    if (n % 3 == 0) {
        ans = ans - n / 3 * 2;
    }
    cout << ans << endl;

    return 0;
}