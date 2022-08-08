#include <algorithm>
#include <cstdio>
#include <cstring>
#include <iostream>
using namespace std;

int left[300005];
int right[300005];

int solution(string &S) {
    // write your code in C++ (C++14 (g++ 6.2.0))
    int n = S.length();
    for (int i = 0; i <= n + 2; i++) {
        left[i] = 0;
        right[i] = 0;
    }
    for (int i = 1; i <= n; i++) {
        if (S[i - 1] == 'R') {
            left[i] = left[i - 1] + 1;
        }
        else {
            left[i] = left[i - 1];
        }
    }

    for (int i = n; i >= 0; i--) {
        if (S[i - 1] == 'R') {
            right[i] = right[i + 1] + 1;
        }
        else {
            right[i] = right[i + 1];
        }
    }

    int ans = 0;
    for (int i = 0; i < n; i++) {
        if (S[i] == 'R')
            continue;
        ans += min(left[i + 1], right[i + 1]);
    }
    return ans;
}

int main() {
    string s;
    cin >> s;
    cout << solution(s);
}