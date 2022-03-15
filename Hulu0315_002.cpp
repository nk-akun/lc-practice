#include <algorithm>
#include <cstdio>
#include <cstring>
#include <iostream>
#include <map>
#include <stack>
#include <vector>
using namespace std;
const int maxn = 1 * (1e5 + 5);

// AC

int n, m, a, b;
char mp[206][206];
char pic[25][25];

int judge(int sx, int sy) {
    if (sx + a - 1 >= n || sy + b - 1 >= m) {
        return 0;
    }
    for (int i = 0; i < a; i++) {
        for (int j = 0; j < b; j++) {
            if (mp[sx + i][sy + j] != pic[i][j]) {
                return 0;
            }
        }
    }
    return 1;
}

int main() {
    cin >> n >> m;
    for (int i = 0; i < n; i++) {
        scanf(" %s", mp[i]);
    }

    cin >> a >> b;
    for (int i = 0; i < a; i++) {
        scanf(" %s", pic[i]);
    }

    int ans = 0;
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < m; j++) {
            ans += judge(i, j);
        }
    }

    printf("%d\n", ans);
    return 0;
}

/*

2 6
*****_
_*****
1 2
*_

*/