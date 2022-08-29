#include <algorithm>
#include <cstdio>
#include <cstring>
#include <iostream>
using namespace std;

int n;
long long k;
struct node {
    long long cost;
    long long val;
} a[200005];

bool cmp(node &x, node &y) {
    return x.cost < y.cost;
}

int main() {

    cin >> n >> k;
    for (int i = 1; i <= n; i++) {
        scanf("%lld %lld", &a[i].cost, &a[i].val);
    }

    sort(a + 1, a + n + 1, cmp);

    int l = 1, r = 1;
    long long sum = a[1].val, ans = 0;
    while (l <= n && r <= n) {
        while (a[r].cost - a[l].cost < k && r <= n) {
            ans = max(ans, sum);
            r++;
            sum += a[r].val;
        }
        sum -= a[l].val;
        l++;
    }

    cout << ans << endl;

    return 0;
}

/*

5 3
1 3
2 1
5 2
3 1
4 3


ACG CGC UCG U

*/