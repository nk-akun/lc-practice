#include <algorithm>
#include <cmath>
#include <cstdio>
#include <cstring>
#include <iostream>
#include <map>
#include <stack>
#include <vector>
using namespace std;
const int maxn = 1 * (1e5 + 5);

int n;
int t[10005][2];
int a[10005];

int main() {
    memset(t, 0, sizeof(t));

    cin >> n;
    for (int i = 1; i <= n; i++) {
        scanf("%d", &a[i]);
    }

    int p;
    for (int i = 1; i <= n; i++) {
        scanf("%d", &p);
        if (p == 0) {
            continue;
        }

        t[p][a[i]]++;
    }

    int ansW = 0, ansB = 0;
    for (int i = 1; i <= n; i++) {
        if (a[i] == 0) {
            if (t[i][1] > 0 || (t[i][0] == 0 && t[i][1] == 0)) {
                ansW++;
            }
        }
        else {
            if (t[i][1] == 0) {
                ansB++;
            }
        }
    }

    printf("%d %d\n", ansW, ansB);
    return 0;
}