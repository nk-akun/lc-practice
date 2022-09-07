#include <algorithm>
#include <cstdio>
#include <cstring>
#include <iostream>
#include <queue>
#include <vector>
using namespace std;

const long long Maxn = 2000065;
typedef long long ll;

priority_queue<pair<ll, ll>> q;
long long ans, n, m, tot, ver[Maxn], head[Maxn], Next[Maxn], edge[Maxn], d[Maxn], v[Maxn];
int U[Maxn], V[Maxn];
void add(long long x, long long y, long long z) {
    ver[++tot] = y;
    Next[tot] = head[x];
    head[x] = tot;
    edge[tot] = z;
}
long long dij(long long s, long long e, long long T) {
    memset(d, 0x7f, sizeof(d));
    memset(v, 0, sizeof(v));
    d[s] = 0;
    q.push(make_pair(0, s));
    while (q.size()) {
        long long x = q.top().second;
        q.pop();
        if (v[x])
            continue;
        v[x] = 1;
        for (int i = head[x]; i; i = Next[i]) {
            ll y = ver[i], z = edge[i];
            ll dis = 1;
            if (z > T)
                dis = 10000000000000;
            if (d[y] > d[x] + dis) {
                d[y] = d[x] + dis;
                q.push(make_pair(-d[y], y));
            }
        }
    }
    return d[e];
}
int main() {
    int n, m, k;
    cin >> n >> m >> k;
    for (int i = 1; i <= m; i++)
        cin >> U[i];
    for (int i = 1; i <= m; i++)
        cin >> V[i];
    for (int i = 1; i <= m; i++) {
        long long d;
        cin >> d;
        add(U[i], V[i], d);
        add(V[i], U[i], d);
    }
    int l = 1, r = 100000, ans = 0;
    while (l <= r) {
        int mid = (l + r) >> 1;
        int dis = dij(1, n, mid);
        if (dis <= k) {
            r = mid - 1;
            ans = mid;
        }
        else {
            l = mid + 1;
        }
    }
    cout << ans << endl;
    return 0;
}