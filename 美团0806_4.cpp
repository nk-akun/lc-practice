#include <algorithm>
#include <cmath>
#include <cstdio>
#include <cstring>
#include <iostream>
#include <vector>
using namespace std;

int n, k;
int cnt[20005];

struct node {
    int type;
    int pos;
} a[20005];

bool cmp(node x, node y) {
    if (x.type != y.type) {
        return x.type < y.type;
    }
    return x.pos < y.pos;
}

int main() {
    cin >> n >> k;

    int t;
    for (int i = 1; i <= n; i++) {
        scanf("%d", &t);
        a[i].pos = i;
        a[i].type = t;
        cnt[t]++;
    }

    sort(a + 1, a + n + 1, cmp);

    vector<int> ans1;
    vector<int> ans2;
    for (int i = 1; i <= n; i++) {
        if (a[i].type != a[i - 1].type) {
            int t = a[i].type;
            int xnum = (cnt[t] + 1) / 2;
            int j;
            for (j = i; j < i + xnum; j++) {
                ans1.push_back(a[j].pos);
            }
            for (j = i + xnum; j <= n && a[j].type == t; j++) {
                ans2.push_back(a[j].pos);
            }
            i = j - 1;
        }
    }

    sort(ans1.begin(), ans1.end());
    sort(ans2.begin(), ans2.end());

    if (ans1.size() > 0) {
        printf("%d", ans1[0]);
    }
    for (int i = 1; i < ans1.size(); i++) {
        printf(" %d", ans1[i]);
    }
    cout << endl;
    if (ans2.size() > 0) {
        printf("%d", ans2[0]);
    }
    for (int i = 1; i < ans2.size(); i++) {
        printf(" %d", ans2[i]);
    }
    cout << endl;
}