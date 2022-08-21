#include <algorithm>
#include <cstdio>
#include <cstring>
#include <iostream>
#include <map>
using namespace std;

string solution(string &S) {
    map<char, int> mp;
    for (int i = 0; i < S.length(); i++) {
        mp[S[i]]++;
    }

    string ans = "";
    for (char c = '9'; c > '0'; c--) {
        while (mp[c] > 1) {
            ans += c;
            mp[c] -= 2;
        }
    }
    if (mp['0'] >= 2 && ans.length() > 0) {
        while (mp['0'] > 1) {
            ans += "0";
            mp['0'] -= 2;
        }
    }

    bool flag = false;
    for (char c = '9'; c >= '0' && !flag; c--) {
        if (mp[c] > 0) {
            string tmp = ans;
            reverse(tmp.begin(), tmp.end());
            ans = ans + c + tmp;
            flag = true;
        }
    }

    if (!flag) {
        string tmp = ans;
        reverse(tmp.begin(), tmp.end());
        ans += tmp;
    }

    return ans;
}

int main() {

    string s = "1212211";
    cout << solution(s) << endl;
}