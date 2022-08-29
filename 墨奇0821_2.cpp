#include <bits/stdc++.h>
using namespace std;
const int MAX = 100;
struct dir {
    int up, down, left, right;
    dir() {}
};
int cnt = 0;
void dfs(vector<string> &G, vector<vector<dir>> &board, vector<vector<bool>> &vis, int x, int y) {
    //	cout<<x<<" "<<y<<endl;
    int n = G.size();
    int m = G[0].length();

    int nx = x, ny = (board[x][y].left + board[x][y].right - y);
    //	cout<<nx<<" "<<ny<<endl;
    if (!vis[nx][ny] && G[nx][ny] == '.') {
        cnt++;
        G[nx][ny] = G[x][y];
        vis[nx][ny] = true;
        dfs(G, board, vis, nx, ny);
    }

    nx = ((board[x][y].up + board[x][y].down) - x), ny = y;
    if (!vis[nx][ny] && G[nx][ny] == '.') {
        cnt++;
        G[nx][ny] = G[x][y];
        vis[nx][ny] = true;
        dfs(G, board, vis, nx, ny);
    }
}

int main() {
    int T, Case = 1;
    scanf("%d", &T);
    while (T--) {
        cnt = 0;
        vector<string> G;
        int n, m;
        scanf("%d%d", &n, &m);
        vector<vector<dir>> board(n, vector<dir>(m));
        vector<vector<bool>> vis(n, vector<bool>(m));
        for (int i = 0; i < n; i++) {
            string s;
            cin >> s;
            G.push_back(s);
        }
        for (int i = 0; i < n; i++) {
            board[i][0].left = G[i][0] == '#' ? 0 : -1;
            board[i][m - 1].right = G[i][m - 1] == '#' ? m - 1 : m;
            for (int j = 1; j < m; j++) {
                if (G[i][j] == '#')
                    board[i][j].left = j;
                else
                    board[i][j].left = board[i][j - 1].left;
            }
            for (int j = m - 2; j >= 0; j--) {
                if (G[i][j] == '#')
                    board[i][j].right = j;
                else
                    board[i][j].right = board[i][j + 1].right;
            }
        }

        for (int i = 0; i < m; i++) {
            board[0][i].up = G[0][i] == '#' ? 0 : -1;
            board[n - 1][i].down = G[n - 1][i] == '#' ? n - 1 : n;
            for (int j = 1; j < n; j++) {
                if (G[j][i] == '#')
                    board[j][i].up = j;
                else
                    board[j][i].up = board[j - 1][i].up;
            }
            for (int j = n - 2; j >= 0; j--) {
                if (G[j][i] == '#')
                    board[j][i].down = j;
                else
                    board[j][i].down = board[j + 1][i].down;
            }
        }
        //		for(int i=0; i<n; i++)
        //		for(int j=0; j<m; j++)
        //			cout<<i<<" "<<j<<": "<<board[i][j].up<<" "<<board[i][j].down<<" "<<board[i][j].left<<" "<<board[i][j].right<<endl;
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                if (G[i][j] != '.' && G[i][j] != '#' && !vis[i][j]) {
                    vis[i][j] = true;
                    dfs(G, board, vis, i, j);
                }
            }
        }
        printf("Case #%d: %d\n", Case, cnt);
        Case++;
        for (int i = 0; i < n; i++)
            cout << G[i] << endl;
    }

    return 0;
}