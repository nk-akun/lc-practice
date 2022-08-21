#include <algorithm>
#include <cstdio>
#include <cstring>
#include <iostream>
#include <map>
#include <vector>
using namespace std;

const int maxn = 100009;
struct edge {
    int to;
    int w;
    int ne;
} e[maxn];

int n, m, len;
int head[maxn], dis[maxn], vis[maxn];

void add(int u, int v, int w) {
    e[len].to = v;
    e[len].w = w;
    e[len].ne = head[u];
    head[u] = len++;
}

void init() {
    len = 0;
    memset(dis, 0, sizeof(dis));
    memset(vis, 0, sizeof(vis));
    memset(head, -1, sizeof(head));
}

void dfs(int x) {
    vis[x] = 1;
    dis[x] = 1;
    for (int i = head[x]; i != -1; i = e[i].ne) {
        int id = e[i].to;
        if (vis[id]) {
            continue;
        }
        dfs(id);
        dis[x] += dis[id];
    }
}

int solution(vector<int> &A, vector<int> &B) {
    init();
    n = A.size() + 1;
    m = A.size();

    for (int i = 0; i < A.size(); i++) {
        add(A[i], B[i], 1);
        add(B[i], A[i], 1);
    }

    dfs(0);

    int ans = 0;
    for (int i = 1; i <= n; i++) {
        ans += (dis[i] + 4) / 5;
    }

    return ans;
}

int main() {

    vector<int> a;
    vector<int> b;
    // a.push_back(0);
    // a.push_back(1);
    // a.push_back(1);
    // b.push_back(1);
    // b.push_back(2);
    // b.push_back(3);

    a.push_back(0);
    b.push_back(1);

    cout << solution(a, b) << endl;
}
