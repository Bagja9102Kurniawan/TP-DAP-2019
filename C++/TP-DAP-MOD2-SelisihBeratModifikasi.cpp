/*	NIM  : 1301194020
	NAMA : BAGJA 9102 KURNIAWAN
	Program ini adalah program pengecekan selisih 
	berat benda yang akan dibawa pada motor.
	motor akan oleng ke satu bagian jika 
	selisihnya lebih dari sama dengan 9kg dengan output
	bernilai true(motor oleng), 
	program akan berhenti jika isi kedua kantong melebihi 150kg 
	atau salah satu kantong bernilai negatif.
*/
#include <iostream>
using namespace std;
float i,c;
int main (){
	while ((i+c<150) &&  (i>=0&&c>=0))
		{
		cout << "Masukkan berat belanjaan di kedua kantong : \n";
		cin >> i >> c;
		cout <<"\nSepeda motor pak Andi akan oleng :" << (i-c>=9||c-i>=9);
	}
	cout << "Proses selesai.";
	return 0;
}
