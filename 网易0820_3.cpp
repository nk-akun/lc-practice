#include <algorithm>
#include <cstdio>
#include <cstring>
#include <iostream>
using namespace std;

int n, k;
int a[200005];
bool del[200005];

int main() {

    cin >> n >> k;
    for (int i = 1; i <= n; i++) {
        scanf("%d", &a[i]);
    }

    int num = n - k;
    for (int m = 31; m >= 0 && num > 0; m--) {
        int cnt = 0;
        for (int i = 1; i <= n; i++) {
            if (!del[i] && !((1 << m) & a[i])) {
                cnt++;
            }
        }
        if (cnt > num) {
            continue;
        }

        for (int i = 1; i <= n; i++) {
            if (!del[i] && !((1 << m) & a[i])) {
                del[i] = true;
            }
        }

        num -= cnt;
    }

    int ans = 0;
    for (int m = 31; m >= 0; m--) {
        bool flag = true;
        for (int i = 1; i <= n; i++) {
            if (!del[i] && !((1 << m) & a[i])) {
                flag = false;
                break;
            }
        }
        if (flag) {
            ans |= (1 << m);
        }
    }

    cout << ans << endl;

    return 0;
}

/*

5 2
10 6 12 18 14

result: 12

5 2
2 3 4 6 5

result: 4

3 1
2 6 5

result: 6

3 2
2 6 5

result: 4

3 2
2 7 5

result: 5

4 2
9 5 10 14

result: 10

*/