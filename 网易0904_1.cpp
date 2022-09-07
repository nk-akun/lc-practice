#include <algorithm>
#include <cstdio>
#include <cstring>
#include <iostream>
using namespace std;

int n;
int a[200005];

int main() {
    cin >> n;

    for (int i = 1; i <= n; i++) {
        cin >> a[i];
    }

    sort(a + 1, a + n + 1);

    int minSub = 0x3f3f3f3f;
    int ans1, ans2;
    for (int i = 2; i <= n; i++) {
        minSub = min(a[i] - a[i - 1], minSub);
    }

    cout << minSub << endl;

    return 0;
}