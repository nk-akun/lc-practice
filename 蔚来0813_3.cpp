#include <algorithm>
#include <cstdio>
#include <cstring>
#include <iostream>
using namespace std;

int main() {

    int t, ans = 0;
    cin >> t;
    while (t--) {
        int tmp;
        cin >> tmp;
        while (tmp) {
            if (tmp % 10 != 0) {
                ans++;
            }
            tmp /= 10;
        }
    }

    cout << ans << endl;
    return 0;
}