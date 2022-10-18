#include <algorithm>
#include <cmath>
#include <cstdio>
#include <cstring>
#include <iostream>
#include <sstream>
#include <vector>
using namespace std;

char s[12345];
char t[12345];

bool judge1(int st, int t) {
    for (int i = st + 1; i <= t; i++) {
        if (s[i] != s[i - 1]) {
            return false;
        }
    }
    return true;
}

bool judge2(int st1, int st2, int l) {
    for (int i = 0; i < l; i++) {
        if (s[st1 + i] != s[st2 + i]) {
            return false;
        }
    }
    return true;
}

int main() {
    cin >> s;
    int len = strlen(s);

    for (int i = len; i >= len / 2; i--) {
        if (i > len) {
            i = len;
        }
        int cnt = 0;
        int j;
        for (j = 0; j <= len - i; j++) {
            if (judge1(j, j + i - 1)) {
                int tmp = i;
                char num[10] = {};
                stringstream p;
                p << tmp;
                p >> num;
                t[cnt++] = '\0';
                strcat(t, num);
                cnt = strlen(t);
                t[cnt++] = '(';
                t[cnt++] = s[j];
                t[cnt++] = ')';
                j = j + i - 1;
            }
            else {
                t[cnt++] = s[j];
            }
        }

        for (; j < len; j++) {
            t[cnt++] = s[j];
        }
        t[cnt] = '\0';
        swap(s, t);
        len = strlen(s);
    }

    len = strlen(s);
    for (int i = len / 2; i > 0; i--) {
        int cnt = 0;
        int j;
        for (j = 0; j <= len - i; j++) {
            int st = j + i;
            int po = 0;
            while (st <= len - i && judge2(j, st, i)) {
                po++;
                st += i;
            }
            if (po > 0) {
                po++;
                char num[10] = {};
                stringstream p;
                p << po;
                p >> num;
                t[cnt++] = '\0';
                strcat(t, num);
                cnt = strlen(t);
                t[cnt++] = '(';
                for (int k = j; k < j + i; k++) {
                    t[cnt++] = s[k];
                }
                t[cnt++] = ')';
                j = j + po * i - 1;
            }
            else {
                t[cnt++] = s[j];
            }
        }
        for (; j < len; j++) {
            t[cnt++] = s[j];
        }
        t[cnt] = '\0';
        swap(s, t);
        len = strlen(s);
    }

    cout << s << endl;

    return 0;
}