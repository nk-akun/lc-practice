#include <algorithm>
#include <cstdio>
#include <cstring>
#include <iostream>
using namespace std;

int n, k;
long long a[200005];

bool cmp(int x, int y) {
    return x > y;
}

int main() {
    cin >> n >> k;
    for (int i = 1; i <= n; i++) {
        cin >> a[i];
    }

    sort(a + 1, a + n + 1, cmp);

    int l = 1, r = 1, ans = 1;
    long long sum = a[1];
    while (r <= n) {
        while (l < r && sum / (double)(r - l + 1) * k < (double)a[l]) {
            sum -= a[l];
            l++;
        }
        ans = max(ans, r - l + 1);
        if (r == n) {
            break;
        }
        r++;
        sum += a[r];
    }

    cout << ans << endl;

    return 0;
}