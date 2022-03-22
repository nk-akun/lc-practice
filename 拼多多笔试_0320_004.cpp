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

/*

题意：输入n,k和长度为n的字符串s，输出字典序大于等于s的字符串，保证出现的每一种字母的数量%k=0.
不存在则输出-1.

*/

int n, k;
char s[100010];
int num[30];
map<char, int> numMap;

bool judge(int sum) {
    memset(num, 0, sizeof(num));
    for (char c = 'a'; c <= 'z'; c++) {
        int need = (k - numMap[c] % k) % k;
        sum -= need;
        if (sum < 0) {
            return false;
        }
        num[c - 'a'] = need;
    }
    if (sum % k == 0) {
        num[0] += sum;
        return true;
    }
    return false;
}

int main() {
    int t;
    cin >> t;
    while (t--) {
        scanf("%d %d", &n, &k);
        scanf(" %s", s);
        int length = strlen(s);

        numMap.clear();
        for (int i = 0; i < length; i++) {
            numMap[s[i]]++;
        }

        int step;
        bool succ = false;
        for (int i = length - 1; i >= 0; i--) {
            numMap[s[i]]--;
            for (i == length - 1 ? s[i] = s[i] : s[i] = s[i] + 1; s[i] <= 'z'; s[i]++) {
                numMap[s[i]]++;
                if (judge(length - 1 - i)) {
                    step = i;
                    succ = true;
                    break;
                }
                numMap[s[i]]--;
            }
            if (succ) {
                break;
            }
        }
        if (succ) {
            for (char c = 'a'; c <= 'z'; c++) {
                while (num[c - 'a']--) {
                    s[++step] = c;
                }
            }
            printf("%s\n", s);
        }
        else {
            printf("-1\n");
        }
    }
    return 0;
}

/*


12
5 1
ababa
ababa
5 1
aaaab
aaaab
6 2
abcabc
abcabc
6 2
aaaccb
aaadad
6 2
abbcca
abbcca
6 2
abcdef
abdabd
12 4
aabccdcacbbc
aabdaabbbddd
12 4
aabccbccbaab
aabccbccbaab
1 1
a
a
1 1
b
b
1 2
b
-1
2 1
ba
ba


*/