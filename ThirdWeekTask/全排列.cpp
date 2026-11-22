#include<bits/stdc++.h>
using namespace std;
int n,a[100],used[100];


void dfs(int k)
{
    int i;
    if(k==n)
    {
        int i;
        for(i=1;i<=n;i++)
        printf("%d ",used[i]);
        cout<<endl;
        return;
    }
    for(i=1;i<=n;i++)
    {
        if(!a[i])
        {
            a[i]=1;
            used[k+1]=i;
            dfs(k+1);
            a[i]=0;
        }
    }
}
int main()
{
    cin>>n;
    dfs(0);
    return 0;
}
