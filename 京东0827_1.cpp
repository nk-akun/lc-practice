#include <algorithm>
#include <cmath>
#include <cstdio>
#include <cstring>
#include <iostream>
#include <vector>
using namespace std;

int n, k;
char s[200005];

int main() {

    cin >> n >> k;
    scanf(" %s", s + 1);

    for (int i = 1; i <= k; i++) {
        if (s[i] >= 'a' && s[i] <= 'z') {
            s[i] -= 32;
        }
    }

    for (int i = k + 1; i <= n; i++) {
        if (s[i] >= 'A' && s[i] <= 'Z') {
            s[i] += 32;
        }
    }

    cout << s + 1 << endl;

    return 0;
}