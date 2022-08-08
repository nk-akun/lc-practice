#include <algorithm>
#include <cmath>
#include <cstdio>
#include <cstring>
#include <iostream>
#include <vector>
using namespace std;

int a[1234];
int dp[1234][2];
int cnt;

int main() {

    int x;
    while (~scanf("%d", &x)) {
        a[++cnt] = x;
    }

    for (int i = 1; i <= cnt; i++) {
        dp[i][0] = max(dp[i - 1][0], dp[i - 1][1]);
        if (i < 3) {
            dp[i][1] = a[i];
        }
        else {
            dp[i][1] = max(dp[i - 1][0], dp[i - 3][1]) + a[i];
        }
    }

    cout << max(dp[cnt][0], dp[cnt][1]) << endl;

    return 0;
}