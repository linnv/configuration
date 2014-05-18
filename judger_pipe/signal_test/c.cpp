#include <iostream>
#include <cstdio>
#include <algorithm>
#include <cstring>
using namespace std;
int T, h, n;
void print(int num)
{
	char a[40];
	for(int i = 0; i < n; i++)
		a[n-1-i] = '0' + (bool)(num & (1<<i));
	a[n] = 0;
	printf("%s\n", a);
}

bool check(int num)
{
	int cnt = 0;
	while(num)
	{
		if(num & 1)
			cnt++;
		num >>= 1;
	}
	return cnt == h;
}

int main()
{
	//freopen("in.txt", "r" ,stdin);
	//freopen("out.txt", "w" ,stdout);
	scanf("%d", &T);
	while(T--)
	{
		scanf("%d %d", &n, &h);
		for(int i = 1; i < (1<< n); i++)
			if(check(i))
				print(i);
		if(T)
			printf("\n");
	}
	return 0;
}
