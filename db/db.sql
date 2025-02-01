-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               8.0.30 - MySQL Community Server - GPL
-- Server OS:                    Win64
-- HeidiSQL Version:             12.1.0.6537
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Dumping database structure for lkpd
CREATE DATABASE IF NOT EXISTS `lkpd` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `lkpd`;

-- Dumping structure for table lkpd.apbd
CREATE TABLE IF NOT EXISTS `apbd` (
  `id` varchar(40) COLLATE utf8mb4_unicode_ci NOT NULL,
  `kode_rekening` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `nama_rekening` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `uraian` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `sub_uraian` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `jml_anggaran_sebelum` bigint DEFAULT NULL,
  `jml_anggaran_setelah` bigint DEFAULT NULL,
  `selisih_anggaran` bigint DEFAULT NULL,
  `persen` varchar(3) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `user_id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tahun_anggaran` varchar(4) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `apbd_kode_rekening_index` (`kode_rekening`),
  KEY `apbd_user_id_index` (`user_id`),
  CONSTRAINT `apbd_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table lkpd.bidang
CREATE TABLE IF NOT EXISTS `bidang` (
  `id` varchar(40) COLLATE utf8mb4_unicode_ci NOT NULL,
  `kode_bidang` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `nama_bidang` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `alias_bidang` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table lkpd.failed_jobs
CREATE TABLE IF NOT EXISTS `failed_jobs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `uuid` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `connection` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `queue` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `payload` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `exception` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `failed_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `failed_jobs_uuid_unique` (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table lkpd.file_iku
CREATE TABLE IF NOT EXISTS `file_iku` (
  `id` varchar(40) COLLATE utf8mb4_unicode_ci NOT NULL,
  `nama_file` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `sub_kegiatan_iku_id` varchar(40) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `file_iku_realisasi_capaian_id_index` (`sub_kegiatan_iku_id`),
  CONSTRAINT `file_iku_realisasi_subkegiatan_iku_id_foreign` FOREIGN KEY (`sub_kegiatan_iku_id`) REFERENCES `sub_kegiatan_iku` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table lkpd.formulasi
CREATE TABLE IF NOT EXISTS `formulasi` (
  `id` varchar(40) COLLATE utf8mb4_unicode_ci NOT NULL,
  `indikator_kinerja_id` varchar(40) COLLATE utf8mb4_unicode_ci NOT NULL,
  `formulasi` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `tipe_penghitungan` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL,
  `divisi_id` varchar(40) COLLATE utf8mb4_unicode_ci NOT NULL,
  `alasan` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `formulasi_indikator_kinerja_id_index` (`indikator_kinerja_id`),
  KEY `formulasi_divisi_id_index` (`divisi_id`),
  CONSTRAINT `formulasi_divisi_id_foreign` FOREIGN KEY (`divisi_id`) REFERENCES `bidang` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `formulasi_indikator_kinerja_id_foreign` FOREIGN KEY (`indikator_kinerja_id`) REFERENCES `indikator_kinerja` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table lkpd.golongan
CREATE TABLE IF NOT EXISTS `golongan` (
  `id` varchar(40) COLLATE utf8mb4_unicode_ci NOT NULL,
  `nama_golongan` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table lkpd.iku_realisasi
CREATE TABLE IF NOT EXISTS `iku_realisasi` (
  `id` varchar(40) COLLATE utf8mb4_unicode_ci NOT NULL,
  `kode_iku` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `sasaran_strategis_id` varchar(40) COLLATE utf8mb4_unicode_ci NOT NULL,
  `indikator_kinerja_id` varchar(40) COLLATE utf8mb4_unicode_ci NOT NULL,
  `formula_id` varchar(40) COLLATE utf8mb4_unicode_ci NOT NULL,
  `target` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `target_tercapai` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0',
  `user_id` varchar(40) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `iku_realisasi_sasaran_strategis_id_index` (`sasaran_strategis_id`),
  KEY `iku_realisasi_indikator_kinerja_id_index` (`indikator_kinerja_id`),
  KEY `iku_realisasi_formula_id_index` (`formula_id`),
  KEY `iku_realisasi_user_id_index` (`user_id`),
  CONSTRAINT `iku_realisasi_formula_id_foreign` FOREIGN KEY (`formula_id`) REFERENCES `formulasi` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `iku_realisasi_indikator_kinerja_id_foreign` FOREIGN KEY (`indikator_kinerja_id`) REFERENCES `indikator_kinerja` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `iku_realisasi_sasaran_strategis_id_foreign` FOREIGN KEY (`sasaran_strategis_id`) REFERENCES `sasaran_strategis` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `iku_realisasi_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table lkpd.indikator_kinerja
CREATE TABLE IF NOT EXISTS `indikator_kinerja` (
  `id` varchar(40) COLLATE utf8mb4_unicode_ci NOT NULL,
  `indikator_kinerja` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `kode_indikator` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table lkpd.kegiatan_iku
CREATE TABLE IF NOT EXISTS `kegiatan_iku` (
  `id` varchar(40) COLLATE utf8mb4_unicode_ci NOT NULL,
  `nama_kegiatan` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `kode_kegiatan` varchar(15) COLLATE utf8mb4_unicode_ci NOT NULL,
  `divisi_id` varchar(40) COLLATE utf8mb4_unicode_ci NOT NULL,
  `program_iku_id` varchar(40) COLLATE utf8mb4_unicode_ci NOT NULL,
  `indikator_kinerja_id` varchar(40) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tahun` varchar(4) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `kegiatan_iku_realisasi_capaian_id_index` (`divisi_id`),
  KEY `program_iku_id` (`program_iku_id`),
  KEY `indikator_kinerja_id` (`indikator_kinerja_id`),
  CONSTRAINT `kegiatan_iku_divisi_id_foreign` FOREIGN KEY (`divisi_id`) REFERENCES `bidang` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `kegiatan_iku_indikator_kinerja_id_foreign` FOREIGN KEY (`indikator_kinerja_id`) REFERENCES `indikator_kinerja` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `kegiatan_iku_program_iku_id_foreign` FOREIGN KEY (`program_iku_id`) REFERENCES `program_anggaran_iku` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table lkpd.kode_rekening
CREATE TABLE IF NOT EXISTS `kode_rekening` (
  `id` varchar(40) COLLATE utf8mb4_unicode_ci NOT NULL,
  `nama_rekening` varchar(250) COLLATE utf8mb4_unicode_ci NOT NULL,
  `kode_rekening` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `ref` varchar(40) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `kode_rekening_kode_rekening_unique` (`kode_rekening`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table lkpd.migrations
CREATE TABLE IF NOT EXISTS `migrations` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `migration` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `batch` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table lkpd.pangkat
CREATE TABLE IF NOT EXISTS `pangkat` (
  `id` varchar(40) COLLATE utf8mb4_unicode_ci NOT NULL,
  `nama_pangkat` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table lkpd.password_resets
CREATE TABLE IF NOT EXISTS `password_resets` (
  `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `token` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `token_use` enum('0','1') COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  KEY `password_resets_email_index` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table lkpd.pegawai
CREATE TABLE IF NOT EXISTS `pegawai` (
  `id` varchar(40) COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_id` varchar(40) COLLATE utf8mb4_unicode_ci NOT NULL,
  `nip` bigint DEFAULT NULL,
  `jabatan` varchar(250) COLLATE utf8mb4_unicode_ci NOT NULL,
  `masa_kerja_golongan` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `diklat` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `pendidikan` varchar(250) COLLATE utf8mb4_unicode_ci NOT NULL,
  `usia` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  `jenis_kelamin` enum('laki-laki','perempuan') COLLATE utf8mb4_unicode_ci NOT NULL,
  `agama` enum('islam','hindu','budha','kristen','konghucu') COLLATE utf8mb4_unicode_ci NOT NULL,
  `kenaikan_pangkat_tahun_berikutnya` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `batas_pensiun` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `golongan_id` varchar(40) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `pangkat_id` varchar(40) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `bidang_id` varchar(40) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `pegawai_user_id_index` (`user_id`),
  KEY `pegawai_golongan_id_index` (`golongan_id`),
  KEY `pegawai_pangkat_id_index` (`pangkat_id`),
  KEY `bidang_id` (`bidang_id`),
  CONSTRAINT `pegawai_golongan_id_foreign` FOREIGN KEY (`golongan_id`) REFERENCES `golongan` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `pegawai_ibfk_1` FOREIGN KEY (`bidang_id`) REFERENCES `bidang` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `pegawai_pangkat_id_foreign` FOREIGN KEY (`pangkat_id`) REFERENCES `pangkat` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table lkpd.personal_access_tokens
CREATE TABLE IF NOT EXISTS `personal_access_tokens` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `tokenable_type` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tokenable_id` bigint unsigned NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `token` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `abilities` text COLLATE utf8mb4_unicode_ci,
  `last_used_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `personal_access_tokens_token_unique` (`token`),
  KEY `personal_access_tokens_tokenable_type_tokenable_id_index` (`tokenable_type`,`tokenable_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table lkpd.program_anggaran_iku
CREATE TABLE IF NOT EXISTS `program_anggaran_iku` (
  `id` varchar(40) COLLATE utf8mb4_unicode_ci NOT NULL,
  `program` varchar(150) COLLATE utf8mb4_unicode_ci NOT NULL,
  `kode_program` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  `anggaran` bigint NOT NULL,
  `anggaran_terpakai` bigint NOT NULL,
  `persentase_anggaran` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `keterangan` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table lkpd.realisasi_anggaran
CREATE TABLE IF NOT EXISTS `realisasi_anggaran` (
  `id` varchar(40) COLLATE utf8mb4_unicode_ci NOT NULL,
  `kode_rekening` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `anggaran_terealisasi` bigint NOT NULL,
  `tahun_anggaran` varchar(4) COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_id` varchar(40) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `realisasi_anggaran_kode_rekening_index` (`kode_rekening`),
  KEY `realisasi_anggaran_user_id_foreign` (`user_id`),
  CONSTRAINT `realisasi_anggaran_kode_rekening_foreign` FOREIGN KEY (`kode_rekening`) REFERENCES `apbd` (`kode_rekening`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `realisasi_anggaran_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table lkpd.sasaran_strategis
CREATE TABLE IF NOT EXISTS `sasaran_strategis` (
  `id` varchar(40) COLLATE utf8mb4_unicode_ci NOT NULL,
  `sasaran_strategis` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table lkpd.schedule
CREATE TABLE IF NOT EXISTS `schedule` (
  `id` varchar(40) COLLATE utf8mb4_unicode_ci NOT NULL,
  `jenis_acara` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tgl_acara` date NOT NULL,
  `jam_acara` time NOT NULL,
  `lokasi_acara` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `redaksi_acara` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `acara_dari` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_id` varchar(40) COLLATE utf8mb4_unicode_ci NOT NULL,
  `status` enum('0','1') COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `schedule_user_id_foreign` (`user_id`),
  CONSTRAINT `schedule_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table lkpd.sub_kegiatan_iku
CREATE TABLE IF NOT EXISTS `sub_kegiatan_iku` (
  `id` varchar(40) COLLATE utf8mb4_unicode_ci NOT NULL,
  `sub_kegiatan` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `indikator_kinerja` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `target_kinerja` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `kode_kegiatan_iku` varchar(15) COLLATE utf8mb4_unicode_ci NOT NULL,
  `persentase` double DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table lkpd.users
CREATE TABLE IF NOT EXISTS `users` (
  `id` varchar(40) COLLATE utf8mb4_unicode_ci NOT NULL,
  `nama` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `username` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `foto` text COLLATE utf8mb4_unicode_ci,
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `jenis_pegawai` enum('admin','pimpinan','pegawai') COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'pegawai',
  `bidang_id` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `remember_token` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `users_username_unique` (`username`),
  UNIQUE KEY `users_email_unique` (`email`),
  KEY `users_divisi_id_index` (`bidang_id`) USING BTREE,
  CONSTRAINT `users_divisi_id_foreign` FOREIGN KEY (`bidang_id`) REFERENCES `bidang` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data exporting was unselected.

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
