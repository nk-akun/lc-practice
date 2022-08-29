#include <algorithm>
#include <cmath>
#include <cstdio>
#include <cstring>
#include <iostream>
#include <vector>
using namespace std;

int n;
char s[200050];

void solve(int x) {
    vector<int> ans;
    ans.clear();

    int len = strlen(s);
    for (int i = 0; i < len; i++) {
        if (i == 0) {
            ans.push_back(1);
        }
        else {
            if (s[i] > s[i - 1]) {
                ans.push_back(ans[ans.size() - 1] + 1);
            }
            else {
                ans.push_back(1);
            }
        }
    }

    printf("Case #%d:", x);
    for (int i = 0; i < ans.size(); i++) {
        printf(" %d", ans[i]);
    }
    cout << endl;
}

int main() {
    int t;
    cin >> t;
    for (int i = 1; i <= t; i++) {
        cin >> n;
        cin >> s;
        solve(i);
    }
}