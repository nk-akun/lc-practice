#include <algorithm>
#include <cstdio>
#include <cstring>
#include <iostream>
using namespace std;

int n;
int a[200005];
int b[200005];

int main() {
    cin >> n;
    for (int i = 1; i <= n; i++) {
        cin >> a[i];
    }
    for (int i = 1; i <= n; i++) {
        cin >> b[i];
    }

    int ans = 0, idx = 1;
    while (idx < n) {
        int count = 0;
        while (idx < n && a[idx + 1] - a[idx] == b[idx + 1] - b[idx]) {
            count++;
            idx++;
        }
        ans = max(ans, count + 1);
        idx++;
    }

    cout << ans << endl;

    return 0;
}