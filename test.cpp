
#include <algorithm>
#include <cstdio>
#include <cstring>
#include <iostream>
#include <map>
#include <stack>
#include <vector>
using namespace std;
const int maxn = 1 * (1e5 + 5);

class Solution {
public:
    int singleNumber(vector<int> &nums) {
        int len = nums.size();
        int ans = 0;
        for (int bit = 0; bit < 32; bit++) {
            int count = 0;
            for (int i = 0; i < len; i++) {
                if (nums[i] & (1 << bit)) {
                    count++;
                }
            }
            if (count % 3) {
                ans |= (1 << bit);
            }
        }
        return ans;
    }
};

int main() {

    int b[10] = {-2, -2, 1, 1, 4, 1, 4, 4, -4, -2};

    Solution a;
    vector<int> nums(b, b + 10);
    int ans = a.singleNumber(nums);
    cout << ans << endl;
}