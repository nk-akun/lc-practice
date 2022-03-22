#include <algorithm>
#include <cstdio>
#include <cstring>
#include <iostream>
using namespace std;
const int inf = 10000000;
char s[3005];
int dp[3005];
int sum[33][3005];

int main() {
    int n;
    scanf("%d", &n);
    scanf("%s", s + 1);
    for (int i = 1; i <= n; i++) {
        sum[s[i] - 'a'][i] = 1;
        dp[i] = -inf;
    }
    for (int i = 1; i <= n; i++) {
        for (int j = 0; j < 26; j++) {
            sum[j][i] += sum[j][i - 1];
        }
    }
    dp[1] = -1;
    dp[0] = 0;
    for (int i = 1; i <= n; i++) {
        for (int j = 1; j <= i; j++) {
            int tmp = 0;
            for (int k = 0; k < 26; k++) {
                if (sum[k][i] - sum[k][i - j] == 0)
                    continue;
                tmp += (sum[k][i] - sum[k][i - j]) % 2 ? -1 : 1;
            }
            dp[i] = max(dp[i], dp[i - j] + tmp);
        }
    }
    printf("%d\n", dp[n]);
    return 0;
}