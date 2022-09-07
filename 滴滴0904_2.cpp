#include <algorithm>
#include <cstdio>
#include <cstring>
#include <iostream>
using namespace std;

int n;
int l[100005], r[100005], q[100005], ans[100005];
int sum[70005][16];

void init() {
    for (int i = 0; i <= 70003; i++) {
        int t = i;
        int x = 0;
        while (t) {
            x ^= t % 10;
            t /= 10;
        }
        sum[i][x] = 1;
    }

    for (int i = 0; i < 16; i++) {
        for (int j = 0; j <= 70003; j++) {
            if (j == 0) {
                continue;
            }
            sum[j][i] += sum[j - 1][i];
        }
    }
}

int main() {
    init();
    cin >> n;
    for (int i = 1; i <= n; i++) {
        cin >> l[i];
    }
    for (int i = 1; i <= n; i++) {
        cin >> r[i];
    }
    for (int i = 1; i <= n; i++) {
        cin >> q[i];
    }

    for (int i = 1; i <= n; i++) {
        if (q[i] > 15) {
            ans[i] = 0;
            continue;
        }
        if (l[i] == 0) {
            ans[i] = sum[r[i]][q[i]];
        }
        else {
            ans[i] = sum[r[i]][q[i]] - sum[l[i] - 1][q[i]];
        }
    }

    cout << ans[1];
    for (int i = 2; i <= n; i++) {
        cout << ' ' << ans[i];
    }
    cout << endl;

    return 0;
}