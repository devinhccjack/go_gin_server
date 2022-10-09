/*
 Navicat MySQL Data Transfer

 Source Server         : mysql
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : localhost:3306
 Source Schema         : testdb

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 09/10/2022 11:21:25
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tests
-- ----------------------------
DROP TABLE IF EXISTS `tests`;
CREATE TABLE `tests` (
  `id` int NOT NULL AUTO_INCREMENT,
  `testcol` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=401 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of tests
-- ----------------------------
BEGIN;
INSERT INTO `tests` VALUES (4, 'testcol');
INSERT INTO `tests` VALUES (5, 'testcol');
INSERT INTO `tests` VALUES (6, 'testcol');
INSERT INTO `tests` VALUES (7, 'testcol');
INSERT INTO `tests` VALUES (8, 'testcol');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
