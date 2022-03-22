#include <algorithm>
#include <cstdio>
#include <cstring>
#include <iostream>
#include <map>
#include <stack>
#include <vector>
#define mem(a, b) memset(a, b, sizeof(a))
#define mod 1000000007
using namespace std;
typedef long long ll;
const int maxn = 1e5 + 5;
const double eps = 1e-12;
const int inf = 0x3f3f3f3f;

int main() {
    int n;
    cin >> n;
    string s;
    cin >> s;
    string graph[n];
    for (int i = 0; i < n; i++) {
        cin >> graph[i];
    }
    int ans = 0;
    for (int i = 0; i < n; i++)
        for (int j = 0; j < n; j++) {
            int k;
            //往右走
            if (n - j >= s.length()) {
                for (k = j; k < j + s.length(); k++) {
                    if (graph[i][k] != s[k - j])
                        break;
                }
                if (k >= j + s.length())
                    ans++;
            }
            //往下走
            if (n - i >= s.length()) {
                for (k = i; k < i + s.length(); k++) {
                    if (graph[k][j] != s[k - i])
                        break;
                }
                if (k >= i + s.length())
                    ans++;
            }
        }
    cout << ans << endl;
    return 0;
}