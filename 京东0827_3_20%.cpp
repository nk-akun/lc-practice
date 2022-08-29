#include <iostream>
#include <vector>
using namespace std;
const long long mod = 1e9 + 7;
const int maxn = 1e6 + 9;
long long fact[maxn + 1], inv[maxn + 1];

long long qpow(long long a, long long b) {
    long long ans;
    if (b == 0)
        return 1;
    if (b % 2)
        ans = a * qpow((a * a) % mod, b / 2);
    else
        ans = qpow((a * a) % mod, b / 2);
    return ans % mod;
}

void init() {
    fact[0] = 1;
    for (int i = 1; i <= maxn; ++i)
        fact[i] = (fact[i - 1] * i) % mod;
    inv[maxn] = qpow(fact[maxn], mod - 2);
    for (int i = maxn - 1; i >= 0; i--) {
        inv[i] = inv[i + 1] * (i + 1);
        inv[i] %= mod;
    }
}
long long C(int n, int m) {
    return ((fact[n] * inv[m]) % mod * (inv[n - m])) % mod;
}

long long count(long long n, long long red) {
    if (n < red * 3)
        return 0;
    n -= red * 3;
    return (qpow(26, n) * C(n + red, red)) % mod;
}
int main() {
    long long n;
    cin >> n;
    init();
    long long ans = 0;
    int pos = 1;
    for (int i = 2; i <= n; i++) {
        if (i * 3 > n)
            break;
        if (pos) {
            ans += count(n, i);
        }
        else {
            ans = (ans + mod - count(n, i));
        }
        ans %= mod;
        pos ^= 1;
    }
    cout << ans << endl;
    return 0;
}

/*

*/