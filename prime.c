#include<stdio.h>
#include<math.h>

int prime(unsigned long n){
	if(n <= 2) return 1;
	
	unsigned long i;
	for(i=3;i< (unsigned long)sqrt(n);i+=2){
		if(0 == n % i) return 0;
		}
	return 1;
}


unsigned long M = 500000;

int main(int argc, char *argv[]){
	unsigned long i;

	for(i=3;i<=M;i+=1){
		if(prime(i));
			//printf("%ld:是素数\n",i);
	}
printf("run done C lang");
return 0;
}

