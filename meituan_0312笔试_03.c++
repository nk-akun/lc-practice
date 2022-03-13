#include <algorithm>
#include <cstdio>
#include <cstring>
#include <iostream>
#include <map>
#include <stack>
#include <vector>
using namespace std;
const int maxn = 1 * (1e5 + 5);

int arr[100][2];
int vis[100];
int ans_count;
int num;
int maxvalue;
void dfs(int n, int ans_count) {
    if (n == num) {
        if (ans_count > maxvalue) {
            maxvalue = ans_count;
        }
        return;
    }
    int x = arr[n][0];
    int y = arr[n][1];
    for (int i = 0; i < 2; i++) {
        if (i == 0) {
            if (vis[x] == 0 && vis[y] == 0 && x != y) {
                vis[x] = 1;
                vis[y] = 1;
                dfs(n + 1, ans_count + 1);
                vis[x] = 0;
                vis[y] = 0;
            }
        }
        else {
            dfs(n + 1, ans_count);
        }
    }
}

int main() {
    int n, m;
    scanf("%d%d", &num, &m);
    for (int i = 0; i < num; i++) {
        scanf("%d%d", &arr[i][0], &arr[i][1]);
    }
    dfs(0, 0);
    printf("%d\n", maxvalue);
    return 0;
}