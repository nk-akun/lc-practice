#include <algorithm>
#include <cstdio>
#include <cstring>
#include <iostream>
using namespace std;

int n;
int m;
long long k;
int a[1234];

int main() {
    cin >> n >> m >> k;
    for (int i = 1; i <= n; i++) {
        cin >> a[i];
    }
    int pos = k % (2 * n);
    if (pos == 0) {
        pos = n;
    }

    if (pos <= n) {
        cout << a[pos] << endl;
    }
    else {
        cout << a[n - (pos - n) + 1] << endl;
    }

    return 0;
}