untuk projek ini, setiap module punya go.mod berbeda.

jadi go get di folder yang sedang dikerjakan dan cd keluar untuk aksi git

contoh: 

1. Open folder client/server di vscode
2. lakukan `go get [package url]` untuk menambahkan package
3. edit code di folder
4. setelah selesai, jalankan:
```
    cd ..
    git pull
    git add .
    git commit -m "nama commit"
    git push
```
5. lalu tutup terminal atau jalankan `cd .\[folder]\` sesuai dengan module mana yang ingin diedit selanjutnya