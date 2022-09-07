#include <algorithm>
#include <cstdio>
#include <cstring>
#include <iostream>
using namespace std;

int n;
int a[200005];

int main() {
    cin >> n;
    bool flag = false;
    int downCount = 0;
    long long sum = 0;
    for (int i = 1; i <= n; i++) {
        cin >> a[i];
        if (a[i] == 0) {
            flag = true;
        }
        if (a[i] < 0) {
            downCount++;
        }
        if (a[i] < 0) {
            a[i] = abs(a[i]);
        }
        sum += abs(a[i] - 1);
    }

    if (!flag && downCount % 2 == 0) {
        flag = true;
    }

    long long ans = 1e17;
    for (int i = 1; i <= n; i++) {
        long long tmp = sum - abs(a[i] - 1) + abs(a[i] - 7);
        ans = min(ans, tmp);
    }

    if (!flag) {
        ans += 2;
    }

    cout << ans << endl;

    return 0;
}

/*

5
-6 0 2 -2 3

*/