#include <algorithm>
#include <cstdio>
#include <cstring>
#include <iostream>
using namespace std;

int a[200005];
int cnt;

int solution(string &S, int B) {
    cnt = 0;
    int last = 0;
    if (S[0] == 'x') {
        last = 1;
    }
    else {
        last = 0;
    }
    S.append(".");
    for (int i = 1; i < S.length(); i++) {
        if (S[i] == '.' && S[i - 1] == 'x') {
            a[cnt] = last;
            cnt++;
            last = 0;
        }
        else if (S[i] == 'x') {
            last++;
        }
    }

    int ans = 0;
    sort(a, a + cnt);
    for (int i = cnt - 1; i >= 0 && B > 0; i--) {
        if (a[i] + 1 <= B) {
            ans += a[i];
            B -= (a[i] + 1);
        }
        else {
            ans += B - 1;
            break;
        }
    }

    return ans;
}

int main() {

    string s = "...xxx..x....xxx.";
    cout << solution(s, 7) << endl;

    s = "..xxxxx";
    cout << solution(s, 4) << endl;

    s = "x.x.xxx...x";
    cout << solution(s, 0) << endl;

    s = "xxxxx..xxx";
    cout << solution(s, 8) << endl;
}