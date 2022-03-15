#include <algorithm>
#include <cstdio>
#include <cstring>
#include <iostream>
#include <map>
#include <stack>
#include <vector>
using namespace std;
const int maxn = 1 * (1e6 + 5);

// 未测试

int n, m;
int a[maxn];
long long k;

int main() {
    cin >> n >> m >> k;

    for (int i = 1; i <= n; i++) {
        scanf("%d", &a[i]);
    }

    long long sum = a[1];
    int num = 0;

    vector<long long> ans;
    ans.clear();
    ans.push_back(sum);

    int l = 1, r = 1;
    while (l < n && r < n) {
        while (r < n && num <= m) {
            r++;
            if (a[r] < 0) {
                num++;
                if (num > m) {
                    num--;
                    r--;
                    break;
                }
            }
            sum += a[r];
            if (num <= m && l <= n && r <= n) {
                ans.push_back(sum);
            }
        }
        while (l < n && l < r && num >= m || r == n) {
            sum -= a[l];
            if (a[l] < 0) {
                num--;
            }
            if (num <= m) {
                ans.push_back(sum);
            }
            l++;
        }
    }

    long long result = -1e16;
    for (int i = 0; i < ans.size(); i++) {
        // cout << ans[i] << endl;
        if (ans[i] <= k && ans[i] > result) {
            result = ans[i];
        }
    }

    cout << result << endl;

    return 0;
}

/*

5 2 15
3 -7 8 -5 9

*/