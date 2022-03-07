
#include <algorithm>
#include <cstdio>
#include <cstring>
#include <iostream>
#include <map>
#include <stack>
#include <vector>
using namespace std;
const int maxn = 1 * (1e5 + 5);

// 扫雷

char mp[10][10];

int main() {
    for (int i = 1; i <= 4; i++) {
        scanf(" %s", mp[i]);
    }

    for (int i = 1; i <= 4; i++) {
        cout << mp[i] << endl;
    }
}
