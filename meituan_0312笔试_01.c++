#include <algorithm>
#include <cstdio>
#include <cstring>
#include <iostream>
#include <map>
#include <stack>
#include <vector>
using namespace std;
const int maxn = 1 * (1e5 + 5);

// 判断一个数是否是11的倍数或数字中包含两个1

int main() {
    int n, x;
    cin >> n;

    while (n--) {
        scanf("%d", &x);
        if (x % 11 == 0) {
            printf("yes\n");
            continue;
        }

        int count = 0;
        while (x) {
            if (x % 10 == 1) {
                count++;
            }
            x /= 10;
        }
        if (count < 2) {
            printf("no\n");
        }
        else {
            printf("yes\n");
        }
    }

    return 0;
}