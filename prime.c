#include<stdio.h>
#include<math.h>

int prime(unsigned long n){
	if(n <= 3) return 0;
	
	unsigned long i;
	for(i=3;i< (unsigned long)sqrt(n);i+=2){
		if(0 == n % i) return 0;
		}
	return 1;
}


unsigned long M = 10000000;

int main(int argc, char *argv[]){
	unsigned long i;
	if(M <=3){
		printf("1,2,3:是素数\n");
		return 0;
		}

	for(i=3;i<=M;i+=2){
		if(prime(i));
			//printf("%ld:是素数\n",i);
	}
printf("run done C lang");
return 0;
}

