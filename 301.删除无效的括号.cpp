#include <algorithm>
#include <cstdio>
#include <cstring>
#include <iostream>
#include <map>
#include <stack>
#include <vector>
using namespace std;
const int maxn = 1 * (1e5 + 5);

// 深搜无剪枝

class Solution {
public:
    vector<string> removeInvalidParentheses(string s) {
        stack<char> st;
        while (!st.empty()) {
            st.pop();
        }

        for (int i = 0; i < s.length(); i++) {
            if (s[i] == '(') {
                st.push(s[i]);
            }
            else if (s[i] == ')') {
                if (!st.empty() && st.top() == '(') {
                    st.pop();
                }
                else {
                    st.push(s[i]);
                }
            }
        }

        int left = 0, right = 0;
        while (!st.empty()) {
            if (st.top() == '(')
                left++;
            else
                right++;
            st.pop();
        }

        memset(state, false, sizeof(state));
        dfs(s, 0, left, right);
        return ans;
    }

    bool state[30];
    int sum[30];
    vector<string> ans;
    map<string, bool> exist;
    void dfs(string &s, int pos, int left, int right) {
        if (pos == s.length()) {
            judge(s);
            return;
        }
        if (s[pos] == '(') {
            // delete
            if (left > 0) {
                state[pos] = true;
                dfs(s, pos + 1, left - 1, right);
                state[pos] = false;
            }
        }
        else if (s[pos] == ')') {

            // delete
            if (right > 0) {
                state[pos] = true;
                dfs(s, pos + 1, left, right - 1);
                state[pos] = false;
            }
        }
        dfs(s, pos + 1, left, right);
    }

    void judge(string &s) {
        string str = "";
        int code = 0;
        for (int i = 0; i < s.length(); i++) {
            if (state[i]) {
                continue;
            }
            str += s[i];
            if (s[i] == '(') {
                code++;
            }
            else if (s[i] == ')') {
                code--;
            }
            if (code < 0) {
                return;
            }
        }
        if (code == 0 && !exist[str]) {
            ans.push_back(str);
            exist[str] = true;
        }
    }
};

int main() {
    Solution a;
    vector<string> ans = a.removeInvalidParentheses(")))))))()))())((()))))");
    cout << ans.size() << endl;
    for (int i = 0; i < ans.size(); i++) {
        cout << ans[i] << endl;
    }

    return 0;
}

/*

执行用时：456 ms, 在所有 C++ 提交中击败了20.72%的用户
内存消耗：33.7 MB, 在所有 C++ 提交中击败了24.65%的用户
通过测试用例：
127 / 127

*/