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
        printf("%d ",a[i]);
        cout<<endl;
        return;
    }
    for(i=1;i<=n;i++)
    {
        if(!used[i])
        {
            used[i]=1;
            a[k+1]=i;
            dfs(k+1);
            used[i]=0;
        }
    }
}
int main()
{
    cin>>n;
    dfs(0);
    return 0;
}
