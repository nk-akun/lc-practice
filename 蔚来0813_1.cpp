#include <algorithm>
#include <cstdio>
#include <cstring>
#include <iostream>
using namespace std;

int n, m;
long long c[100][100];

void init() {
    for (int i = 1; i <= 50; i++) {
        c[i][0] = 1;
    }
    c[1][1] = 1;

    for (int n = 2; n <= 50; n++) {
        for (int m = 1; m <= 50; m++) {
            c[n][m] = c[n - 1][m - 1] + c[n - 1][m];
        }
    }
}

int main() {

    init();

    int n, m;
    cin >> n >> m;
    double ans = c[n][m];
    for (int i = 1; i <= m; i++) {
        ans *= 0.8;
    }
    for (int i = 1; i <= n - m; i++) {
        ans *= 0.2;
    }

    printf("%.4f\n", ans);
}