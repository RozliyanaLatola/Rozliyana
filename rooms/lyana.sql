-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Oct 09, 2020 at 04:25 AM
-- Server version: 10.4.14-MariaDB
-- PHP Version: 7.4.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `lyana`
--

-- --------------------------------------------------------

--
-- Table structure for table `ruangs`
--

CREATE TABLE `ruangs` (
  `RuangID` varchar(50) NOT NULL,
  `RuangName` varchar(50) NOT NULL,
  `PasienName` varchar(50) NOT NULL,
  `Biaya` varchar(50) NOT NULL,
  `LamaNginap` varchar(50) NOT NULL,
  `Penyakit` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `ruangs`
--

INSERT INTO `ruangs` (`RuangID`, `RuangName`, `PasienName`, `Biaya`, `LamaNginap`, `Penyakit`) VALUES
('R001', 'BIG BOSS', 'Chika', '1500000', '5', 'DBD'),
('R002', 'ICU', 'Chaca', '850000', '3', 'MAAG'),
('R003', 'INTEL', 'LALA', '700000', '3', 'USUS BUNTU'),
('R004', 'CEMARA', 'LUNG LUNG', '200000', '1', 'FLU');
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
