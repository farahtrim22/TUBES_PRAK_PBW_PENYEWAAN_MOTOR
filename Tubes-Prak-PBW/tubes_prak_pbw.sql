-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 02 Jul 2024 pada 13.49
-- Versi server: 10.4.32-MariaDB
-- Versi PHP: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `tubes_prak_pbw`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `customer`
--

CREATE TABLE `customer` (
  `id_customer` int(10) UNSIGNED NOT NULL,
  `nama_lengkap` varchar(300) NOT NULL,
  `nik` varchar(16) NOT NULL,
  `jenis_kelamin` tinyint(1) NOT NULL,
  `alamat` text NOT NULL,
  `nomor_telepon` varchar(15) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `customer`
--

INSERT INTO `customer` (`id_customer`, `nama_lengkap`, `nik`, `jenis_kelamin`, `alamat`, `nomor_telepon`) VALUES
(34, 'Nadia Ayu Rahmawati', '1111111111111111', 2, 'Bekasi Timur Regency', '082110111111'),
(35, 'Afni Puspita Sari', '2222222222222222', 2, 'Cibubur', '082210222222'),
(36, 'Ramadhani', '3333333333333333', 1, 'Depok Baru', '0811233746392'),
(37, 'Salwa Khairu Mista', '4444444444444444', 2, 'Cilodong', '08123765432'),
(39, 'Siti gumai', '6666666666666666', 2, 'Gondangdia', '0863283728326');

-- --------------------------------------------------------

--
-- Struktur dari tabel `motor`
--

CREATE TABLE `motor` (
  `id_motor` int(10) UNSIGNED NOT NULL,
  `merek` varchar(300) NOT NULL,
  `tipe` varchar(300) NOT NULL,
  `jenis_motor` varchar(300) NOT NULL,
  `tahun_produksi` int(11) NOT NULL,
  `warna` varchar(300) NOT NULL,
  `stok` int(11) NOT NULL,
  `harga` decimal(10,0) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `motor`
--

INSERT INTO `motor` (`id_motor`, `merek`, `tipe`, `jenis_motor`, `tahun_produksi`, `warna`, `stok`, `harga`) VALUES
(13, 'Honda', 'Beat', 'Matic', 2022, 'Hitam', 1, 120000),
(14, 'Yamaha', 'NMAX', 'Matic', 2022, 'Hitam', 21, 140000),
(15, 'Honda ', 'Scoopy', 'Matic', 2021, 'Merah', 0, 120000),
(16, 'Vespa', 'Primavera 150', 'Matic', 2022, 'Biru', 3, 150000),
(17, 'Suzuki', 'Nex II 113 F', 'Matic', 2018, 'Merah', 14, 60000),
(18, 'Honda', 'CB150R', 'Sport', 2019, 'Hitam', 6, 140000),
(20, 'Suzuki', 'Nex I', 'Matic', 2017, 'putih', 18, 50000);

-- --------------------------------------------------------

--
-- Struktur dari tabel `sewa`
--

CREATE TABLE `sewa` (
  `id_sewa` int(10) UNSIGNED NOT NULL,
  `id_customer` int(10) UNSIGNED DEFAULT NULL,
  `id_motor` int(10) UNSIGNED DEFAULT NULL,
  `tanggal_sewa` date NOT NULL,
  `tanggal_kembali` date NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `sewa`
--

INSERT INTO `sewa` (`id_sewa`, `id_customer`, `id_motor`, `tanggal_sewa`, `tanggal_kembali`) VALUES
(27, 34, 13, '2024-07-02', '2024-07-03'),
(28, 36, 15, '2024-07-03', '2024-07-04');

-- --------------------------------------------------------

--
-- Struktur dari tabel `user`
--

CREATE TABLE `user` (
  `id_user` int(10) UNSIGNED NOT NULL,
  `nama_lengkap` varchar(300) NOT NULL,
  `email` varchar(300) NOT NULL,
  `username` varchar(50) NOT NULL,
  `password` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `user`
--

INSERT INTO `user` (`id_user`, `nama_lengkap`, `email`, `username`, `password`) VALUES
(10, 'Daffa Abraar Sajuti', 'daffasaj@gmail.com', 'daffa123', '$2a$10$HJDbc9s4AluKgeVTV3lFCuxzDtjX.rU9Zpz/LaQdN1aD0.01as9va'),
(11, 'Nadia Ayu', 'nayu@gmail.com', 'nadiayu', '$2a$10$E/bmZqeR.LWv23Vvltgb/.cg7iaFtfD0hIhw0bBR83S5opho1R1u.'),
(13, 'Farah Tri Mahardini', 'farahtrim@gmail.com', 'farah123', '$2a$10$zUGyYUWZRt0UAQpBI53RGu.qnH7ENHp3xDzQpE61CsFtU/n2l1tHm'),
(15, 'daffa', 'daffa@gamil.com', 'daffa', '$2a$10$3PF/czbNYF/2jY6vNbmmh.uKXnzksMftj9avHcXqqKNhS7rgxFq9q');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `customer`
--
ALTER TABLE `customer`
  ADD PRIMARY KEY (`id_customer`);

--
-- Indeks untuk tabel `motor`
--
ALTER TABLE `motor`
  ADD PRIMARY KEY (`id_motor`);

--
-- Indeks untuk tabel `sewa`
--
ALTER TABLE `sewa`
  ADD PRIMARY KEY (`id_sewa`),
  ADD KEY `id_customer` (`id_customer`,`id_motor`),
  ADD KEY `id_motor` (`id_motor`);

--
-- Indeks untuk tabel `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id_user`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `customer`
--
ALTER TABLE `customer`
  MODIFY `id_customer` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=40;

--
-- AUTO_INCREMENT untuk tabel `motor`
--
ALTER TABLE `motor`
  MODIFY `id_motor` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=21;

--
-- AUTO_INCREMENT untuk tabel `sewa`
--
ALTER TABLE `sewa`
  MODIFY `id_sewa` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=30;

--
-- AUTO_INCREMENT untuk tabel `user`
--
ALTER TABLE `user`
  MODIFY `id_user` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `sewa`
--
ALTER TABLE `sewa`
  ADD CONSTRAINT `sewa_ibfk_1` FOREIGN KEY (`id_customer`) REFERENCES `customer` (`id_customer`),
  ADD CONSTRAINT `sewa_ibfk_2` FOREIGN KEY (`id_motor`) REFERENCES `motor` (`id_motor`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
