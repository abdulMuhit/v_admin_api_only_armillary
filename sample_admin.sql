-- phpMyAdmin SQL Dump
-- version 4.7.9
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1:3306
-- Generation Time: Aug 05, 2019 at 08:16 PM
-- Server version: 10.2.14-MariaDB
-- PHP Version: 7.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `bandros_vendor`
--

DELIMITER $$
--
-- Procedures
--
DROP PROCEDURE IF EXISTS `add_child_user`$$
CREATE DEFINER=`admin`@`localhost` PROCEDURE `add_child_user` (IN `iemail` VARCHAR(100), IN `istatus` CHAR(1), IN `iparent` VARCHAR(100), IN `ipassword` VARCHAR(1000), IN `_list` VARCHAR(300), IN `irole` CHAR(1), IN `inama` VARCHAR(300))  NO SQL
BEGIN

DECLARE _next TEXT DEFAULT NULL;
DECLARE _nextlen INT DEFAULT NULL;
DECLARE _value TEXT DEFAULT NULL;
    
 DECLARE errno INT;
    DECLARE EXIT HANDLER FOR SQLEXCEPTION
    BEGIN
    GET CURRENT DIAGNOSTICS CONDITION 1 errno = MYSQL_ERRNO;
    SELECT errno AS MYSQL_ERROR;
    ROLLBACK;
    END;
    
SELECT SUBSTRING(id,10) sufix FROM `pengguna`
WHERE DATE_FORMAT(tanggal,'%Y-%m-%d')=CURDATE()
ORDER BY SUFIX DESC LIMIT 1 INTO @LAST_SUFIX_ID;

SELECT CONCAT("p",DATE_FORMAT(NOW(),'%y%m%d'),IF(@LAST_SUFIX_ID IS NULL,'0001',LPAD(@LAST_SUFIX_ID+1, 4, '0'))) INTO @NEW_ID_P;
   
START TRANSACTION;

INSERT INTO `pengguna` (id, email, `status`, parent, `password`, role, nama) VALUES (@NEW_ID_P, iemail, istatus, iparent, ipassword, irole, inama);

iterator:
LOOP
  IF LENGTH(TRIM(_list)) = 0 OR _list IS NULL THEN
    LEAVE iterator;
  END IF;
  
  SET _next = SUBSTRING_INDEX(_list,',',1);
  SET _nextlen = LENGTH(_next);
  SET _value = TRIM(_next);
  
SELECT RIGHT(id,4) sufix FROM `brand_user` WHERE id_user = @NEW_ID_P ORDER BY SUFIX DESC LIMIT 1 INTO @LAST_SUFIX;
SELECT @LAST_SUFIX;

SELECT CONCAT(@NEW_ID_P,IF(@LAST_SUFIX IS NULL,'0001',LPAD(@LAST_SUFIX+1, 4, '0'))) INTO @NEW_ID;
SELECT @NEW_ID;

INSERT INTO `brand_user` (id, id_user, id_brand) VALUES (@NEW_ID, @NEW_ID_P, _next);

  SET _list = INSERT(_list,1,_nextlen + 1,'');
END LOOP;

COMMIT;
END$$

DROP PROCEDURE IF EXISTS `assign_brand`$$
CREATE DEFINER=`admin`@`localhost` PROCEDURE `assign_brand` (IN `user` VARCHAR(50), IN `stat` TINYINT(1), IN `_list` MEDIUMTEXT, IN `inama` VARCHAR(300), IN `iid` VARCHAR(100))  NO SQL
BEGIN

DECLARE _next TEXT DEFAULT NULL;
DECLARE _nextlen INT DEFAULT NULL;
DECLARE _value TEXT DEFAULT NULL;
    
 DECLARE errno INT;
    DECLARE EXIT HANDLER FOR SQLEXCEPTION
    BEGIN
    GET CURRENT DIAGNOSTICS CONDITION 1 errno = MYSQL_ERRNO;
    SELECT errno AS MYSQL_ERROR;
    ROLLBACK;
    END;
   
START TRANSACTION;

DELETE FROM `brand_user` WHERE id_user=iid;

iterator:
LOOP
  
  
  IF LENGTH(TRIM(_list)) = 0 OR _list IS NULL THEN
    LEAVE iterator;
  END IF;

  
  SET _next = SUBSTRING_INDEX(_list,',',1);

  
  
  
  SET _nextlen = LENGTH(_next);

  
  SET _value = TRIM(_next);

SELECT RIGHT(id,4) sufix FROM `brand_user` WHERE id_user = iid ORDER BY SUFIX DESC LIMIT 1 INTO @LAST_SUFIX;
SELECT @LAST_SUFIX;

SELECT CONCAT(iid,IF(@LAST_SUFIX IS NULL,'0001',LPAD(@LAST_SUFIX+1, 4, '0'))) INTO @NEW_ID;

  
  INSERT INTO `brand_user` (id, id_user, id_brand) VALUES (@NEW_ID, iid, _next);

  
  
  
  
  SET _list = INSERT(_list,1,_nextlen + 1,'');
END LOOP;




UPDATE `pengguna` SET status = stat, nama = inama WHERE id=iid;

    
COMMIT;

END$$

DROP PROCEDURE IF EXISTS `delete_by_vendor`$$
CREATE DEFINER=`admin`@`localhost` PROCEDURE `delete_by_vendor` (IN `idd` VARCHAR(100))  NO SQL
BEGIN
DECLARE _next TEXT DEFAULT NULL;
DECLARE _nextlen INT DEFAULT NULL;
DECLARE _value TEXT DEFAULT NULL;

DECLARE n INT DEFAULT 0;
DECLARE i INT DEFAULT 0;

DECLARE errno INT;
    DECLARE EXIT HANDLER FOR SQLEXCEPTION
    BEGIN
    GET CURRENT DIAGNOSTICS CONDITION 1 errno = MYSQL_ERRNO;
    SELECT errno AS MYSQL_ERROR;
    ROLLBACK;
    END;
   
START TRANSACTION;

SELECT COUNT(*) FROM pengguna WHERE parent=idd INTO n;
select n;
SET i=0;
WHILE i<n DO 

 SELECT id FROM pengguna WHERE parent=idd LIMIT i,1 INTO @VAR;
 SELECT CONCAT_WS(",", @JO, @VAR) into @JO;
 SET i = i + 1;
END WHILE;

SELECT CONCAT_WS(",",@JO,idd) into @JO;
SELECT @JO;



iterator:
LOOP
  
  
  IF LENGTH(TRIM(@JO)) = 0 OR @JO IS NULL THEN
    LEAVE iterator;
  END IF;

  
  SET _next = SUBSTRING_INDEX(@JO,',',1);

  
  
  
  SET _nextlen = LENGTH(_next);

  
  SET _value = TRIM(_next);

  
  
  
  DELETE FROM `brand_user` WHERE id_user=_value;
  SET @JO = INSERT(@JO,1,_nextlen + 1,'');

END LOOP;

DELETE FROM `pengguna` WHERE parent=idd;
DELETE FROM `pengguna` WHERE id=idd;
COMMIT;

End$$

DROP PROCEDURE IF EXISTS `insert_vendor_tran`$$
CREATE DEFINER=`admin`@`localhost` PROCEDURE `insert_vendor_tran` (IN `nama` VARCHAR(100), IN `no_hp` VARCHAR(50), IN `email` VARCHAR(100), IN `alamat` TEXT, IN `nama_brand` VARCHAR(100), IN `kota` VARCHAR(50), IN `provinsi` VARCHAR(50), IN `kodepos` VARCHAR(50), IN `website` VARCHAR(100), IN `facebook` VARCHAR(100), IN `instagram` VARCHAR(100), IN `marketplace` VARCHAR(100), IN `kategori` VARCHAR(100), IN `pic1` VARCHAR(100), IN `pic2` VARCHAR(100), IN `pic3` VARCHAR(100))  NO SQL
BEGIN

 DECLARE errno INT;
    DECLARE EXIT HANDLER FOR SQLEXCEPTION
    BEGIN
    GET CURRENT DIAGNOSTICS CONDITION 1 errno = MYSQL_ERRNO;
    SELECT errno AS MYSQL_ERROR;
    ROLLBACK;
    END;

SELECT SUBSTRING(id,8) sufix FROM `vendor`
WHERE DATE_FORMAT(tanggal,'%Y-%m-%d')=CURDATE()
ORDER BY SUFIX DESC LIMIT 1 INTO @LAST_SUFIX;

SELECT CONCAT("3",DATE_FORMAT(NOW(),'%y%m%d'),IF(@LAST_SUFIX IS NULL,'0001',LPAD(@LAST_SUFIX+1, 4, '0'))) INTO @NEW_ID;

START TRANSACTION;

INSERT INTO `vendor`(id, nama, no_hp, email, alamat, nama_brand, kota, provinsi, kodepos, website, facebook, instagram, marketplace, kategori) VALUE(@NEW_ID, nama, no_hp, email, alamat, nama_brand, kota, provinsi, kodepos, website, facebook, instagram, marketplace, kategori );

INSERT INTO `picture`(id, owner, tipe, path_url) VALUE(CONCAT(@NEW_ID,"1"), @NEW_ID, "vendorpic", pic1);

IF (pic2 IS NOT NULL AND pic2 !='') THEN

INSERT INTO `picture`(id, owner, tipe, path_url) VALUE(CONCAT(@NEW_ID,"2"), @NEW_ID, "vendorpic", pic2);
END IF;

IF (pic3 IS NOT NULL AND pic3 !='') THEN

INSERT INTO `picture`(id, owner, tipe, path_url) VALUE(CONCAT(@NEW_ID,"3"), @NEW_ID, "vendorpic", pic3);
END IF;

COMMIT WORK;

SELECT @NEW_ID new_id;

END$$

DELIMITER ;

-- --------------------------------------------------------

--
-- Table structure for table `brand_user`
--

DROP TABLE IF EXISTS `brand_user`;
CREATE TABLE IF NOT EXISTS `brand_user` (
  `id` varchar(40) NOT NULL,
  `id_user` varchar(30) NOT NULL,
  `id_brand` varchar(30) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `brand_user`
--

INSERT INTO `brand_user` (`id`, `id_user`, `id_brand`) VALUES
('p19080500010001', 'p1908050001', '109'),
('p19080500020001', 'p1908050002', '109'),
('p19080500030001', 'p1908050003', '109'),
('p19080500040001', 'p1908050004', '108'),
('p19080500040002', 'p1908050004', '107'),
('p19080500040003', 'p1908050004', '105'),
('p19080500050001', 'p1908050005', '107'),
('p19080500060001', 'p1908050006', '107'),
('p19080500060002', 'p1908050006', '108');

-- --------------------------------------------------------

--
-- Table structure for table `pengguna`
--

DROP TABLE IF EXISTS `pengguna`;
CREATE TABLE IF NOT EXISTS `pengguna` (
  `id` varchar(100) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password` varchar(1000) NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT 0,
  `parent` varchar(100) DEFAULT NULL,
  `role` char(1) NOT NULL,
  `nama` varchar(300) NOT NULL,
  `tanggal` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `pengguna`
--

INSERT INTO `pengguna` (`id`, `email`, `password`, `status`, `parent`, `role`, `nama`, `tanggal`) VALUES
('000000', 'myAdmin', '4376ef317c9b749e56d776029eac96d41db2aa105539595e80aa45146c47a35651f8619aa304520c13ea81190257e17c374748308c5de7d8f13daee14f34983b', 1, 'SU', '0', 'Super A', '2018-10-15 00:17:10'),
('p1908050001', 'rabani@email.com', '4376ef317c9b749e56d776029eac96d41db2aa105539595e80aa45146c47a35651f8619aa304520c13ea81190257e17c374748308c5de7d8f13daee14f34983b', 1, '', '1', 'rabani user', '2019-08-04 17:26:14'),
('p1908050006', 'atva2@email.com', '4376ef317c9b749e56d776029eac96d41db2aa105539595e80aa45146c47a35651f8619aa304520c13ea81190257e17c374748308c5de7d8f13daee14f34983b', 1, 'p1908050004', '2', 'atva dan skylight', '2019-08-04 18:23:59');

-- --------------------------------------------------------

--
-- Table structure for table `picture`
--

DROP TABLE IF EXISTS `picture`;
CREATE TABLE IF NOT EXISTS `picture` (
  `id` varchar(100) NOT NULL,
  `owner` varchar(30) NOT NULL,
  `tipe` varchar(30) NOT NULL,
  `path_url` varchar(1000) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `picture`
--

INSERT INTO `picture` (`id`, `owner`, `tipe`, `path_url`) VALUES
('100', '31810200001', 'vendorpic', 'vendorpic0852339788240yALBsdjS.jpg'),
('101', '31810200002', 'vendorpic', 'vendorpic089784478510ngCwFkDi.jpg'),
('102', '31810230001', 'vendorpic', 'vendorpic0856492393230BuufFMoW.png'),
('103', '31810270001', 'vendorpic', 'vendorpic0896015845390XVlBzgba.JPG'),
('104', '31810270001', 'vendorpic', 'vendorpic0896015845391XVlBzgba.JPG'),
('105', '31810270001', 'vendorpic', 'vendorpic0896015845392XVlBzgba.JPG'),
('106', '31810290001', 'vendorpic', 'vendorpic0813940262400MRAjWwhT.jpg'),
('107', '31810290001', 'vendorpic', 'vendorpic0813940262401MRAjWwhT.jpg'),
('108', '31810290001', 'vendorpic', 'vendorpic0813940262402MRAjWwhT.jpg'),
('109', '31810290002', 'vendorpic', 'vendorpic0858922473060tcuAxhxK.png'),
('110', '31811030001', 'vendorpic', 'vendorpic0812338646490DaFpLSjF.jpg'),
('111', '31811030002', 'vendorpic', 'vendorpic0813351436950XoEFfRsW.jpg'),
('112', '31811050001', 'vendorpic', 'vendorpic0853875924160LDnJObCs.jpg'),
('318110500021', '31811050002', 'vendorpic', 'vwl_3e34d22c-e41e-402e-b414-e33ba3447f54addvendor2.png'),
('318110500022', '31811050002', 'vendorpic', 'vwl_f8db3607-57d8-4f97-a3b9-e891ceb749dbalumnidicoding.png'),
('318110800011', '31811080001', 'vendorpic', 'vwl_0454b6cb-fe40-41aa-9c7f-6137dff105cdarmy_logo_SWISS.png'),
('318110800012', '31811080001', 'vendorpic', 'vwl_e00c35a1-39b9-4760-a3bb-67277cc16840LOGO_BELMONT_NEW.png'),
('318110800013', '31811080001', 'vendorpic', 'vwl_a3d96e2b-67c7-4839-9a59-9fcab35ad9adLOGO_RUGGER.png'),
('318110900011', '31811090001', 'vendorpic', 'vwl_5c82f146-8f3f-4a43-a02b-94ef6373b383IMG-20181106-WA0001.jpg'),
('318111000011', '31811100001', 'vendorpic', 'vwl_44b43a3e-66f3-4ea7-a95f-64234c159573ADVKUNING_1.jpg'),
('318111100011', '31811110001', 'vendorpic', 'vwl_dba551b5-e373-4d7e-8823-12a495642af1IMG_20181103_105223.jpg'),
('318111100012', '31811110001', 'vendorpic', 'vwl_8e3d687d-4dc0-48d1-9893-047435b44777IMG_20181103_104802.jpg'),
('318111100013', '31811110001', 'vendorpic', 'vwl_7f5b28ef-72a7-4e72-bcd6-9f97490c649bIMG_20181103_103502.jpg'),
('318111200011', '31811120001', 'vendorpic', 'vwl_4e68a6c1-8d58-4099-b2a7-d6067e347aa9IMG20181107094159.jpg'),
('318111500011', '31811150001', 'vendorpic', 'vwl_f52d6f10-17cd-4c8d-ab1c-df1dcd02445dIMG_20181115_134126.jpg'),
('318111500021', '31811150002', 'vendorpic', 'vwl_3fb808bc-5ba2-432d-81c9-f7b837e78ae2IMG_20181115_162824.jpg'),
('318111500022', '31811150002', 'vendorpic', 'vwl_2664230d-b44c-40ca-a9d4-edf619bd1082IMG_20181115_162625.jpg'),
('318111500031', '31811150003', 'vendorpic', 'vwl_6f34498e-7e77-41bb-8430-c7aae987cc6fIMG-20181112-WA0002.jpg'),
('318111500032', '31811150003', 'vendorpic', 'vwl_84d1b431-369c-4f8e-b10f-c11a166a1e20IMG-20181112-WA0003.jpg'),
('318111500033', '31811150003', 'vendorpic', 'vwl_c0ca6a09-b8ad-4fca-99fa-3e63a57fb54cIMG-20181112-WA0004.jpg'),
('318112100011', '31811210001', 'vendorpic', 'vwl_81035fc8-f312-40a2-beee-1beb153f279eIMG_20170826_164542_834.jpg'),
('318112200011', '31811220001', 'vendorpic', 'vwl_5de4feed-1f48-42a1-8117-1c58f0091c216938aee7-4e77-4e8f-86ab-5c4237406c96.jpeg'),
('318112200012', '31811220001', 'vendorpic', 'vwl_1077ae99-64c8-4ad6-b266-9302f7729cb11af08468-5a78-4c34-be8a-9c389781ad70.jpeg'),
('318112200013', '31811220001', 'vendorpic', 'vwl_8bbbe4e0-50d3-4107-9ce9-8c7b873c215a65f4804d-13c8-465a-b367-2620ed340c7b.jpeg'),
('318112300011', '31811230001', 'vendorpic', 'vwl_452c5eb0-4d9e-484e-b7c8-d8d7757b22ddd_etnik_batik_tenun_nusantara___Bp4DoFAAgv3___.jpg'),
('318112300012', '31811230001', 'vendorpic', 'vwl_95a47156-8cfa-47c0-853f-14defdaf6d83d_etnik_batik_tenun_nusantara___Bp4TuVGghXw___.jpg'),
('318112300013', '31811230001', 'vendorpic', 'vwl_86b677cd-da4b-40c4-b3ee-33d9b2ab75a4d_etnik_batik_tenun_nusantara___BqEksn7g4qT___.jpg'),
('318112300021', '31811230002', 'vendorpic', 'vwl_cde2a31e-f611-4742-a158-fb547a3e5357jashujan.jpg'),
('318112300022', '31811230002', 'vendorpic', 'vwl_2c01c8c1-9be3-405d-98ce-2470ae5227e0rca001.jpg'),
('318112700011', '31811270001', 'vendorpic', 'vwl_ab59f954-cd7c-4dbb-a0c2-01cd79adfa1bSO-01_Coklat_(1).jpg'),
('318112700012', '31811270001', 'vendorpic', 'vwl_f1234445-ce72-47fa-ab4e-727da51f4c8aSO-01_Cream_(1).jpg'),
('318112700013', '31811270001', 'vendorpic', 'vwl_23274c7e-9b66-48b1-b1f1-b0f66ad70c54SO-01.jpg'),
('318112800011', '31811280001', 'vendorpic', 'vwl_2eeb8773-7bf6-4293-9d3e-1034140ec1a3IMG-20181117-WA0018-min.jpg'),
('318112800012', '31811280001', 'vendorpic', 'vwl_3f8c2273-ab59-42d1-9ba6-bea215c53b28IMG-20181117-WA0017-min.jpg'),
('318112800021', '31811280002', 'vendorpic', 'vwl_a98e7dac-7b19-43f2-ab8c-11f2576139aaSO-01_Coklat_(1).jpg'),
('318112800022', '31811280002', 'vendorpic', 'vwl_d248d761-e0f2-4270-ba9d-d73d8f7abd5dSO-01_Cream_(1).jpg'),
('318112800023', '31811280002', 'vendorpic', 'vwl_1018c665-9911-421f-830a-b4c2097b206cSO-01.jpg'),
('318112800031', '31811280003', 'vendorpic', 'vwl_d3748a92-0b0e-49a4-b18f-9f8e08ef1cafIMG_20181126_083932_717.jpg'),
('318112800032', '31811280003', 'vendorpic', 'vwl_c1a62eee-6cfd-4a6c-af1b-55016f0026f31543195728772.jpg'),
('318112800033', '31811280003', 'vendorpic', 'vwl_cef2a10a-104c-49a4-b29d-675db4eeb84d1543194624260.jpg'),
('318112800041', '31811280004', 'vendorpic', 'vwl_54d14321-73b7-485b-910f-e54184edf5c1IMG-20181127-WA0034.jpeg'),
('318112800051', '31811280005', 'vendorpic', 'vwl_84b85693-3a87-4140-ba5a-2ad02a18a9aaIMG-20181127-WA0034.jpeg'),
('318112800052', '31811280005', 'vendorpic', 'vwl_50719eaa-b2e1-4de3-bc11-c80c8bb96f7dIMG-20181127-WA0003.jpeg'),
('318112800053', '31811280005', 'vendorpic', 'vwl_27fa6b8f-3aeb-4bbb-800f-97cfa2e81372IMG-20181127-WA0000.jpeg'),
('318112800061', '31811280006', 'vendorpic', 'vwl_06872925-7304-43e6-b396-d774ec3ab02bpp.jpg'),
('318112900011', '31811290001', 'vendorpic', 'vwl_94816eea-646f-4799-8960-d4bdf6db61d5SALMA-2-600x600_(1).jpg'),
('318112900012', '31811290001', 'vendorpic', 'vwl_2d6f847a-42ac-4f3e-b648-9a64f5b8b522riby-pants-2-768x768.jpg'),
('318112900013', '31811290001', 'vendorpic', 'vwl_844ff398-ee67-4021-8991-0c1bdd392fecLAYERED-2-768x768.jpg'),
('318120100011', '31812010001', 'vendorpic', 'vwl_2bf655ba-e681-414e-b852-1237bc07fd7aIMG_8018_resize.jpg'),
('318120300011', '31812030001', 'vendorpic', 'vwl_2554cc79-f4fa-4cbc-9e09-431bd3ce82b3DP-003.jpeg'),
('318120300012', '31812030001', 'vendorpic', 'vwl_bb9b8a23-6bca-4b7f-953e-046e6d2b3a48PK-003.jpeg'),
('318120300013', '31812030001', 'vendorpic', 'vwl_6436a717-d64b-464b-bd9f-5fe2833c1930HC-003.jpeg'),
('318120300021', '31812030002', 'vendorpic', 'vwl_1f3fa387-80b6-4f31-84b8-ad321943340bEnjoy_Orange_(2).jpg'),
('318120300022', '31812030002', 'vendorpic', 'vwl_ba71ea14-7a68-4f49-82ca-1b2e850daa4cPlastic_Cup_Hitam_(1).jpg'),
('318120300023', '31812030002', 'vendorpic', 'vwl_1fadd6ab-bc0c-4ed7-b28b-5ab34bf4dfa706_sling_bag_hitam.jpg'),
('99', '31810190001', 'vendorpic', 'vendorpic0853202187272bWGvbqzg.jpg');

-- --------------------------------------------------------

--
-- Table structure for table `vendor`
--

DROP TABLE IF EXISTS `vendor`;
CREATE TABLE IF NOT EXISTS `vendor` (
  `id` varchar(30) NOT NULL,
  `nama` varchar(100) NOT NULL,
  `no_hp` varchar(30) NOT NULL,
  `email` varchar(100) NOT NULL,
  `alamat` text NOT NULL,
  `nama_brand` varchar(100) NOT NULL,
  `kota` varchar(100) NOT NULL,
  `provinsi` varchar(30) NOT NULL,
  `kodepos` varchar(20) NOT NULL,
  `website` varchar(100) DEFAULT NULL,
  `facebook` varchar(100) DEFAULT NULL,
  `instagram` varchar(100) DEFAULT NULL,
  `marketplace` varchar(100) DEFAULT NULL,
  `kategori` varchar(100) NOT NULL,
  `tanggal` timestamp NOT NULL DEFAULT current_timestamp(),
  `status` varchar(40) DEFAULT 'waiting',
  `password` varchar(300) DEFAULT NULL,
  `role` tinyint(4) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `vendor`
--

INSERT INTO `vendor` (`id`, `nama`, `no_hp`, `email`, `alamat`, `nama_brand`, `kota`, `provinsi`, `kodepos`, `website`, `facebook`, `instagram`, `marketplace`, `kategori`, `tanggal`, `status`, `password`, `role`) VALUES
('123456', 'super admin', '', 'admin', '', '', '', '', '', '', '', '', '', '', '2018-09-04 03:48:00', 'rejected', '4376ef317c9b749e56d776029eac96d41db2aa105539595e80aa45146c47a35651f8619aa304520c13ea81190257e17c374748308c5de7d8f13daee14f34983b', 7),
('31809110001', 'Iqbal test', '4343443', 'ciebal745@gmail.com', 'Jl. test 123 no 1 blok A', 'BRANDTEST', 'BANDUNG', 'JAWA BARAT', '402354', '', 'fb.com', '', 'test.com/fb', 'Fashion,Aksesoris', '2018-09-11 06:48:51', 'waiting', NULL, 0),
('31812030002', 'Sardjono', '4343434', 'assd@gmail.com', 'Vil Sawa', 'No Brand', 'Tangerang Selatan', 'Banten', '15413', '', '', 'https://www.instagram.com/aldora_ind/', 'https://www.tokopedia.com/aldoraind', 'Minuman,Fashion', '2018-12-03 10:48:20', 'waiting', NULL, 0);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
