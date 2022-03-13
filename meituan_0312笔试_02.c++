#include <algorithm>
#include <cstdio>
#include <cstring>
#include <iostream>
#include <map>
#include <stack>
#include <vector>
using namespace std;
const int maxn = 1 * (1e5 + 5);

int a[10002];

int main() {
    int n;
    cin >> n;
    for (int i = 1; i <= n; i++) {
        scanf("%d", &a[i]);
    }

    int ans = 0;
    for (int i = 1; i <= n; i++) {
        int num = 0;
        for (int j = i; j <= n; j++) {
            if (a[j] < 0) {
                num++;
            }
            if (!(num & 1)) {
                ans++;
            }
        }
    }
    printf("%d\n", ans);
    return 0;
}