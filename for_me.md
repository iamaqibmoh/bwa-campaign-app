Rangkuman :
-
Ada 4 layer untuk Clean Architecture
> 1. View (Form HTML : Input User)
 
> 2. Handler / Controller
* Fungsinya untuk mapping Input User -> struct Input ranah Service

> 3. Service
* Fungsinya untuk mapping struct input -> struct User ranah Repository

> 4. Repository
* Untuk insert data (struct User) -> Database


Nah melalui 4 layer tersebut baru deh masuk ke database datanya :)

Tahapan upload file ke database
-
1. Ambil file dari request form-data menggunakan key-file nya
2. Save file asli di server (path sudah ditentukan)
3. Yang disimpan ke database string file-path nya