#include <bits/stdc++.h>
using namespace std;
map<int, int> idx;
struct node {
    int val, zheng, fu;
    node() {
        val = 0;
        zheng = 0;
        fu = 0;
    }
    node(int _val, int _zheng, int _fu) {
        val = _val;
        zheng = _zheng;
        fu = _fu;
    }
};
const int INF = 0x3f3f3f3f;
vector<node> num;
vector<int> z;
int main() {
    num.push_back(node());
    int n;
    scanf("%d", &n);
    for (int i = 0; i < n; i++) {
        int tmp;
        scanf("%d", &tmp);
        z.push_back(tmp);
        if (idx[tmp])
            num[idx[tmp]].zheng++;
        else {
            idx[tmp] = num.size();
            num.push_back(node(tmp, 1, 0));
        }
    }
    for (int i = 0; i < n; i++) {
        int tmp;
        scanf("%d", &tmp);
        if (tmp != z[i]) {
            if (idx[tmp])
                num[idx[tmp]].fu++;
            else {
                idx[tmp] = num.size();
                num.push_back(node(tmp, 0, 1));
            }
        }
    }
    int ans = INF;
    int half = n / 2 + (n & 1);
    for (int i = 0; i < num.size(); i++) {
        node &p = num[i];
        if (p.zheng + p.fu >= half) {
            int cost = half - p.zheng;
            ans = max(min(ans, cost), 0);
        }
    }
    if (ans == INF)
        ans = -1;
    printf("%d\n", ans);
    return 0;
}
/*
5
1 5 7 5 5
10 5 5 9 10
3
1 2 3
4 4 1
*/