-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               8.0.29 - MySQL Community Server - GPL
-- Server OS:                    Win64
-- HeidiSQL Version:             12.4.0.6659
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Dumping database structure for ngs_db
CREATE DATABASE IF NOT EXISTS `ngs_db` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `ngs_db`;

-- Dumping structure for table ngs_db.mt202
CREATE TABLE IF NOT EXISTS `mt202` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `SenderBIC` char(12) NOT NULL DEFAULT '0' COMMENT 'The Logical Termial (LT)\r\nAddress is a 12-character FIN\r\naddress. It is the address of the\r\nsending LT for input messages\r\nor of the receiving LT for output\r\nmessages, and includes the\r\nBranch Code. It consists of: -\r\nthe BIC 8 CODE (8 characters)\r\n- the Logical Terminal Code (1\r\nupper case alphabetic\r\ncharacter) - the BIC Branch\r\nCode (3 characters)',
  `ReceiverBIC` char(12) NOT NULL DEFAULT '0' COMMENT 'This address\r\nis the 12-\r\ncharacter\r\nSWIFT\r\naddress of\r\nthe receiver\r\nof the\r\nmessage. It\r\ndefines the\r\ndestination to\r\nwhich the\r\nmessage\r\nshould be\r\nsent.',
  `Direction` char(1) NOT NULL DEFAULT '0' COMMENT 'Input: I, Output: O',
  `MTType` char(3) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0' COMMENT 'The Message\r\nType consists\r\nof 3 digits\r\nwhich define\r\nthe MT\r\nnumber of the\r\nmessage\r\nbeing input',
  `UETR` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0' COMMENT 'Field 121:\r\n\r\nThis field provides an end-to-end reference across a payment transaction. The format of field 121 is xxxxxxxx-xxxx-4xxx-yxxxxxxxxxxxxxxx where x is any hexadecimal character (lower case only) and y is one of 8, 9, a, or b. Field 121 is mandatory on all MTs 103, 103 STP, 103 REMIT, 202, 202 COV, 205, and 205 COV. See the FIN Service Description for additional information.\r\n\r\nField 121 can be present without field 111. Field 111 can only be present if field 121 is also present.',
  `F20` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '(A-Z, a-z, 0-9) ',
  `F21` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0' COMMENT '(A-Z, a-z, 0-9) ',
  `F32A_ValueDate` date NOT NULL DEFAULT '0000-00-00' COMMENT 'F32A YYMMDD (0-9)',
  `F32A_Currency` char(3) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'F32A Currency (A-Z, a-z)',
  `F32A_Amount` decimal(20,5) NOT NULL DEFAULT '0.00000' COMMENT 'F32A Amount',
  `F52a` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT 'The value of Field Tag :52A: or :52D:. Concatenate all data lines using the delimiter value full stop “.”',
  `F57a` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT 'The value of Field Tag :57A: or :57B: or :57D:. Concatenate all data lines using the delimiter value full stop “.”',
  `F58a` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'The value of Field Tag :58A: or :58D:. Concatenate all data lines using the delimiter value full stop “.”',
  `RawData` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'The whole content of the payment message',
  `created_at` timestamp NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=117 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Data exporting was unselected.

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
