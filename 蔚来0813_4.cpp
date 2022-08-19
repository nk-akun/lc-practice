#include <algorithm>
#include <cstdio>
#include <cstring>
#include <iostream>
using namespace std;

int main() {
    int n, x;
    cin >> n >> x;

    int ans = 0;
    for (int i = 1; i <= n; i++) {
        int t = i;
        while (t) {
            if (t % 10 == x) {
                ans++;
            }
            t /= 10;
        }
    }

    cout << ans << endl;
}