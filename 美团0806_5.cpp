#include <algorithm>
#include <cmath>
#include <cstdio>
#include <cstring>
#include <iostream>
#include <vector>
using namespace std;

int main() {

    string s1 = "MeiTuannauTieMwow";
    string s2 = "wowMeiTuannauTieMwow";

    int t;
    cin >> t;
    while (t--) {
        long long x;
        cin >> x;
        if (x - 1 < s1.length()) {
            cout << s1[x - 1] << endl;
        }
        else {
            cout << s2[(x - 1 - s1.length()) % s2.length()] << endl;
        }
    }
}