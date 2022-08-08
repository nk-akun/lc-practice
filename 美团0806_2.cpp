#include <algorithm>
#include <cmath>
#include <cstdio>
#include <cstring>
#include <iostream>
using namespace std;

int n;
int a[200005];
int l[200005];
int r[200005];

int main() {

    cin >> n;
    for (int i = 1; i <= n; i++) {
        scanf("%d", &a[i]);
    }

    int sum = 0;
    for (int i = 1; i <= n; i++) {
        if (a[i] >= 0) {
            sum++;
        }
        l[i] = sum;
    }

    sum = 0;
    for (int i = n; i >= 0; i--) {
        r[i] = sum;
        if (a[i] <= 0) {
            sum++;
        }
    }

    int ans = 0x3f3f3f3f;
    for (int i = 0; i <= n; i++) {
        ans = min(ans, l[i] + r[i]);
    }

    cout << ans << endl;
}