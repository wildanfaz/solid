# SOLID Principle

Implementasi SOLID Principle dalam Bahasa Pemrograman Go

## 1. SRP : Single Responsibility Principle

Dalam konteks CRUD, implementasi yang tepat dari SRP dapat dilakukan dengan memisahkan tanggung jawab antara lapisan repository dan lapisan service

```bash
  make start_srp
```

## 2. OCP : Open/Closed Principle

Menambahkan implementasi baru tanpa mengubah code yang sudah ada

```bash
  make start_ocp

```

## 3. LSP : Liskov Substitution Principle

Kelas turunan harus dapat digunakan sebagai pengganti kelas induknya

```bash
  make start_lsp

```

## 4. ISP : Interface Segregation Principle

Tidak boleh dipaksa untuk mengimplementasikan metode yang tidak digunakan dan interface yang spesifik harus dipilih daripada satu interface umum

```bash
  make start_isp

```

## 5. DIP : Dependency Inversion Principle

Modul tingkat tinggi tidak boleh bergantung pada modul tingkat rendah. Keduanya seharusnya bergantung pada abstraksi dan abstraksi tidak boleh bergantung pada detail. Detail seharusnya bergantung pada abstraksi

```bash
  make start_dip

```