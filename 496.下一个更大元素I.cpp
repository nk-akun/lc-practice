#include <algorithm>
#include <cstdio>
#include <cstring>
#include <iostream>
#include <stack>
#include <vector>
using namespace std;
const int maxn = 1 * (1e5 + 5);

// 单调递减栈

class Solution {
public:
    vector<int> nextGreaterElement(vector<int> &nums1, vector<int> &nums2) {
        int record[10005];
        stack<int> s;
        while (!s.empty())
            s.pop();
        for (int i = 0; i < nums2.size(); i++) {
            if (s.empty()) {
                s.push(i);
                continue;
            }
            if (nums2[s.top()] < nums2[i]) {
                while (!s.empty() && nums2[s.top()] < nums2[i]) {
                    int temp = nums2[s.top()];
                    s.pop();
                    record[temp] = nums2[i];
                }
            }
            s.push(i);
        }
        while (!s.empty()) {
            record[nums2[s.top()]] = -1;
            s.pop();
        }

        vector<int> ans;
        for (int i = 0; i < nums1.size(); i++) {
            ans.push_back(record[nums1[i]]);
        }
        return ans;
    }
};

int main() {
    Solution a;
    vector<int> nums1;
    vector<int> nums2;
    nums1.push_back(4);
    nums1.push_back(1);
    nums1.push_back(2);
    nums2.push_back(1);
    nums2.push_back(3);
    nums2.push_back(4);
    nums2.push_back(2);

    vector<int> ans = a.nextGreaterElement(nums1, nums2);
    for (int i = 0; i < ans.size(); i++) {
        printf("%d ", ans[i]);
    }
    cout << endl;
}