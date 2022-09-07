#include <algorithm>
#include <iostream>
#include <vector>
using namespace std;

int a[1000005], sum[1000005];

int main() {
    int w, n, cur = 0;
    cin >> w >> n;
    for (int i = 1; i <= n; i++) {
        cin >> a[i];
        sum[i] = sum[i - 1] + a[i];
        if (i != 1) {
            sum[i]++;
        }
    }
    vector<int> vc;
    for (int i = 1; i <= w; i++) {
        int pos = lower_bound(sum, sum + n + 1, i) - sum;
        pos--;
        int res = sum[n] - sum[pos];
        if (pos != 0) {
            res--;
        }
        if (w - i < res) {
            vc.push_back(i);
        }
    }
    cout << vc.size() << endl;
    for (int i = 0; i < vc.size(); i++) {
        if (i == vc.size() - 1) {
            cout << vc[i] << endl;
        }
        else {
            cout << vc[i] << " ";
        }
    }
    return 0;
}
