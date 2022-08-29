#include <algorithm>
#include <cmath>
#include <cstdio>
#include <cstring>
#include <iostream>
#include <map>
#include <vector>
using namespace std;

int n;
int a[1000005];

bool cmp(pair<int, int> x, pair<int, int> y) {
    return x.second > y.second;
}

int solve() {
    map<int, int> fi;
    map<int, int> se;
    fi.clear();
    se.clear();

    for (int i = 1; i <= n; i += 2) {
        fi[a[i]]++;
    }
    for (int i = 2; i <= n; i += 2) {
        se[a[i]]++;
    }

    vector<pair<int, int>> ff;
    vector<pair<int, int>> ss;

    ff.clear();
    ss.clear();

    for (map<int, int>::iterator it = fi.begin(); it != fi.end(); it++) {
        ff.push_back(make_pair(it->first, it->second));
    }
    for (map<int, int>::iterator it = se.begin(); it != se.end(); it++) {
        ss.push_back(make_pair(it->first, it->second));
    }

    sort(ff.begin(), ff.end(), cmp);
    sort(ss.begin(), ss.end(), cmp);

    ff.push_back(make_pair(0, 0));
    ss.push_back(make_pair(0, 0));

    if (ff[0].first != ss[0].first) {
        return n - (ff[0].second + ss[0].second);
    }
    else {
        return n - max(ff[0].second + ss[1].second, ff[1].second + ss[0].second);
    }
}

int main() {

    cin >> n;
    for (int i = 1; i <= n; i++) {
        scanf("%d", &a[i]);
    }

    if (n <= 2) {
        cout << 0 << endl;
    }
    else {
        cout << solve() << endl;
    }

    return 0;
}