#include <algorithm>
#include <cmath>
#include <cstdio>
#include <cstring>
#include <iostream>
#include <map>
#include <stack>
#include <vector>
using namespace std;
const int maxn = 1 * (1e5 + 5);

int cost[10005][15];

int calMinEnerge(int n, int m, vector<int> &bomb) {
    int inf = 0x3f3f3f3f;
    for (int i = 1; i <= m; i++) {
        for (int j = 1; j <= n; j++) {
            cost[i][j] = inf;
        }
    }

    cost[1][1] = 0;
    for (int i = 2; i <= m; i++) {
        for (int j = 1; j <= n; j++) {
            for (int k = 1; k <= n; k++) {
                cost[i][j] = min(cost[i][j], cost[i - 1][k] + (k != j));
            }
        }
        cost[i][bomb[i - 1]] = inf;
    }
    int ans = inf;
    for (int i = 1; i <= n; i++) {
        ans = min(ans, cost[m][i]);
    }
    return ans;
}

int main() {
    int n, m;
    vector<int> bomb;
    scanf("%d%d", &n, &m);
    for (int i = 1; i <= m; i++) {
        int tmp;
        scanf("%d", &tmp);
        bomb.push_back(tmp);
    }
    printf("%d\n", calMinEnerge(n, m, bomb));

    return 0;
}