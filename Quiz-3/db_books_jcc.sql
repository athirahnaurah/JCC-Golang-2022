-- phpMyAdmin SQL Dump
-- version 5.0.3
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 27 Mar 2022 pada 22.54
-- Versi server: 10.4.14-MariaDB
-- Versi PHP: 7.4.11

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_books_jcc`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `book`
--

CREATE TABLE `book` (
  `id` int(11) NOT NULL,
  `title` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `image_url` varchar(500) NOT NULL,
  `release_year` int(11) NOT NULL,
  `price` varchar(255) NOT NULL,
  `total_page` int(11) NOT NULL,
  `thickness` varchar(10) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `category_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `book`
--

INSERT INTO `book` (`id`, `title`, `description`, `image_url`, `release_year`, `price`, `total_page`, `thickness`, `created_at`, `updated_at`, `category_id`) VALUES
(31, 'Anchika', 'Dia yang bersamaku tahun 1995', 'https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.gramedia.com%2Fblog%2Fancika-sosok-tambatan_hati_dilan-setelah-milea%2F&psig=AOvVaw1Cw4K20pu-izKxFJgLlL0R&ust=1648493292980000&source=images&cd=vfe&ved=0CAsQjRxqFwoTCPiWrJv65vYCFQAAAAAdAAAAABAD', 2021, '100.000', 337, 'tebal', '2022-03-28 01:48:28', '2022-03-28 01:48:28', 1),
(33, 'Rindu', 'Lima kisah dalam sebuah perjalanan panjang kerinduan', 'https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.goodreads.com%2Fbook%2Fshow%2F23288914-rindu&psig=AOvVaw2KQVV-WH2hkrMH4UNlLWxo&ust=1648493877004000&source=images&cd=vfe&ved=0CAsQjRxqFwoTCLjOgbL85vYCFQAAAAAdAAAAABAD', 2014, '90.000', 544, 'tebal', '2022-03-28 01:58:39', '2022-03-28 01:58:39', 1),
(34, 'Full House', 'Komik Full House Jilid 16', 'https://www.google.com/url?sa=i&url=https%3A%2F%2Fen.wikipedia.org%2Fwiki%2FFull_House_(manhwa)&psig=AOvVaw1mJVyoDUEBHTPNeP_61nLF&ust=1648494864996000&source=images&cd=vfe&ved=0CAsQjRxqFwoTCMjYwoqA5_YCFQAAAAAdAAAAABAD', 1999, '50.000', 100, 'tipis', '2022-03-28 02:14:44', '2022-03-28 02:14:44', 3);

-- --------------------------------------------------------

--
-- Struktur dari tabel `category`
--

CREATE TABLE `category` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `category`
--

INSERT INTO `category` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'novel', '2022-03-27 22:51:07', '2022-03-27 22:51:07'),
(2, 'biografi', '2022-03-27 22:51:51', '2022-03-27 22:53:01'),
(3, 'komik', '2022-03-27 22:51:58', '2022-03-27 22:51:58');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `book`
--
ALTER TABLE `book`
  ADD PRIMARY KEY (`id`),
  ADD KEY `category_id` (`category_id`);

--
-- Indeks untuk tabel `category`
--
ALTER TABLE `category`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `book`
--
ALTER TABLE `book`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=35;

--
-- AUTO_INCREMENT untuk tabel `category`
--
ALTER TABLE `category`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `book`
--
ALTER TABLE `book`
  ADD CONSTRAINT `book_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
