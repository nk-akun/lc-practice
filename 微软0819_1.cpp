#include <algorithm>
#include <cstdio>
#include <cstring>
#include <iostream>
#include <map>
using namespace std;

// int solution(int X[], int Y[], int N, int W) {
//     sort(X, X + N);

//     int idx = 0;
//     int ans = 0;

//     while (idx < N) {
//         ans++;
//         int j;
//         for (j = idx; j < N && X[idx] + W >= X[j]; j++) {
//         }
//         idx = j;
//     }

//     return ans;
// }

int solution(vector<int> &X, vector<int> &Y, int W) {
    sort(X.begin(), X.end());
    int idx = 0;
    int ans = 0;

    while (idx < X.size()) {
        ans++;
        int j;
        for (j = idx; j < X.size() && X[idx] + W >= X[j]; j++) {
        }
        idx = j;
    }

    return ans;
}

int main() {
    // int x[6] = {2, 4, 2, 6, 7, 1};
    // int y[6] = {0, 5, 3, 2, 1, 5};
    // int n = 6;
    // int w = 2;

    int x[6] = {2, 4, 2, 6, 7, 1};
    int y[6] = {0, 5, 3, 2, 1, 5};
    int n = 6;
    int w = 2;

    cout << solution(x, y, n, w) << endl;
}