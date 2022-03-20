#include <algorithm>
#include <cstdio>
#include <cstring>
#include <iostream>
#include <map>
#include <stack>
#include <vector>
#define mem(a, b) memset(a, b, sizeof(a))
#define mod 1000000007
using namespace std;
typedef long long ll;
const int maxn = 1e5 + 5;
const double eps = 1e-12;
const int inf = 0x3f3f3f3f;

struct edge {
    int to;
    int w;
    int ne;
} e[1000010];

int n, m, len;
int head[1010];
bool vis[1010];

void add(int u, int v, int w) {
    e[len].to = v;
    e[len].w = w;
    e[len].ne = head[u];
    head[u] = len++;
}

void dfs(int now, int value) {
    vis[now] = true;
    for (int i = head[now]; i != -1; i = e[i].ne) {
        int id = e[i].to;
        if (vis[id] || e[i].w < value) {
            continue;
        }
        dfs(id, value);
    }
}

bool judge(int value) {
    memset(vis, 0, sizeof(vis));
    dfs(1, value);
    for (int i = 1; i <= n; i++) {
        if (!vis[i]) {
            return false;
        }
    }
    return true;
}

int main() {
    memset(head, -1, sizeof(head));
    cin >> n >> m;

    int u, v, w, maxv = -1;
    for (int i = 1; i <= m; i++) {
        scanf("%d %d %d", &u, &v, &w);
        add(u, v, w);
        add(v, u, w);
        maxv = max(maxv, w);
    }

    int l = 1, r = maxv;
    while (l <= r) {
        int mid = (l + r) >> 1;
        if (judge(mid)) {
            l = mid + 1;
        }
        else {
            r = mid - 1;
        }
    }

    cout << r << endl;

    return 0;
}

/*

3 3
1 2 3
1 3 4
2 3 5

6 7
1 2 3
2 3 5
2 4 4
4 5 4
3 6 3
1 5 4
4 6 4

6 7
1 2 3
2 3 5
2 4 4
4 5 4
3 6 3
1 5 4
4 6 3

6 7
1 2 5
2 3 5
2 4 4
4 5 5
3 6 5
1 5 5
4 6 3

*/