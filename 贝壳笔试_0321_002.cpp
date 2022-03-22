#include <algorithm>
#include <cstdio>
#include <cstring>
#include <iostream>
#include <map>
#include <stack>
#include <vector>
#define mem(a, b) memset(a, b, sizeof(a))
#define mod 1000000007
using namespace std;
typedef long long ll;
const int maxn = 1e5 + 5;
const double eps = 1e-12;
const int inf = 0x3f3f3f3f;

int a[1000010];

void init() {
    for (int i = 1; i <= 1000000; i++) {
        int sum = 0, tmp = i;
        while (tmp) {
            sum += tmp % 10;
            tmp /= 10;
        }
        if (i % sum == 1) {
            a[i] = 1;
        }
    }
    for (int i = 1; i <= 1000000; i++) {
        a[i] += a[i - 1];
    }
}

int main() {
    init();
    int t, l, r;
    cin >> t;

    while (t--) {
        scanf("%d %d", &l, &r);
        printf("%d\n", a[r] - a[l - 1]);
    }

    return 0;
}