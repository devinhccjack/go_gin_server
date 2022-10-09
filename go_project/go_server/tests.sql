/*
 Navicat MySQL Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80021
 Source Host           : localhost:3306
 Source Schema         : testdb

 Target Server Type    : MySQL
 Target Server Version : 80021
 File Encoding         : 65001

 Date: 09/10/2022 23:11:39
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tests
-- ----------------------------
DROP TABLE IF EXISTS `tests`;
CREATE TABLE `tests` (
  `id` int NOT NULL AUTO_INCREMENT,
  `testcol` varchar(255) DEFAULT NULL,
  `title` varchar(255) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `status` tinyint DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tests
-- ----------------------------
BEGIN;
INSERT INTO `tests` VALUES (1, 'testcol', 'title', 'description', 1, '2022-10-09 23:03:56', '2022-10-09 23:03:56', NULL);
INSERT INTO `tests` VALUES (2, 'testcol', 'title', 'description', 1, '2022-10-09 23:04:16', '2022-10-09 23:04:16', NULL);
INSERT INTO `tests` VALUES (3, 'testcol', 'title', 'description', 1, '2022-10-09 23:04:17', '2022-10-09 23:04:17', NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
