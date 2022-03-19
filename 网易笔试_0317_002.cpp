#include <algorithm>
#include <cstdio>
#include <cstring>
#include <iostream>
#include <map>
#include <stack>
#include <vector>
using namespace std;
const int maxn = 1 * (1e5 + 5);

const int inf = 1e9 + 7;
const double eps = 1e-8;
char mp[60][60];
int dir[8][2] = {{-1, -1}, {-1, 0}, {-1, 1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {0, -1}};
int mk[2][60][60];
int main() {
    int t, m, n;
    scanf("%d", &t);
    while (t--) {
        memset(mk, 0, sizeof(mk));
        int cur = 1;
        scanf("%d%d", &m, &n);
        for (int i = 1; i <= m; i++)
            scanf("%s", mp[i] + 1);
        for (int i = 0; i < n; i++) {
            int x, y;
            scanf("%d%d", &x, &y);
            x++, y++;
            mp[x][y] = cur ? 'w' : 'b';
            cur ^= 1;
            memset(mk[cur], 0, sizeof(mk[cur]));
            for (int k = 0; k < 8; k++) {
                bool ok = false;
                int tx = x, ty = y;
                int ansx = -1, ansy = -1;
                while (1) {
                    tx = tx + dir[k][0];
                    ty = ty + dir[k][1];
                    if (tx > m || tx < 1 || ty > m || ty < 1) {
                        break;
                    }
                    if (mp[tx][ty] == '-') {
                        break;
                    }
                    else if (mp[x][y] != mp[tx][ty]) {
                        ok = true;
                        ansx = tx, ansy = ty;
                    }
                }
                tx = x, ty = y;
                while (ok) {
                    tx = tx + dir[k][0];
                    ty = ty + dir[k][1];
                    if (ansx == tx && ansy == ty) {
                        break;
                    }
                    if (mk[cur ^ 1][tx][ty])
                        continue;
                    if (mp[tx][ty] == 'w') {
                        mp[tx][ty] = 'b';
                        mk[cur][tx][ty] = 1;
                    }
                    else if (mp[tx][ty] == 'b') {
                        mp[tx][ty] = 'w';
                        mk[cur][tx][ty] = 1;
                    }
                }
            }
        }
        for (int i = 1; i <= m; i++) {
            printf("%s\n", mp[i] + 1);
        }
        printf("END\n");
    }
    return 0;
}
/*
2
6 4
------
---b--
-wwwb-
-bbw--
-ww---
-ww---
1 2
0 2
2 0
2 5
6 4
------
--b---
--wbb-
--www-
--b-b-
--w---
1 4
1 1
5 4
0 2

*/