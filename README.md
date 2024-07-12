## Penjelasan Analisis Komplesitas dari Kode

1. Time Complexity
   - Loop fungsi utama mengulangi setiap karakter dalam input berupa string.
   - Di dalam loop, operasi waktu konstan dilakukan:
     - Memeriksa tipe karakter apakah sesuai opening dan closing menggunakan switch.
     - Push atau Pop elemen dari Array (operasi append dan slicing).
     - Menggunakan fungsi isMatching, yang juga memiliki kompleksitas waktu yang konstan.
   - Maka dari itu, Time Complexity dari fungsi isBalanced menggunakan looping, sehingga menghasilkan O(n).

2. Space Complexity
   - Code tersebut menggunakan array untuk menyimpan tanda kurung buka.
   - Ukuran maksimum array bisa menjadi n dalam case ini.
