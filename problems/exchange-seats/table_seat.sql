/*
 Navicat Premium Data Transfer

 Source Server         : root@localhost
 Source Server Type    : MySQL
 Source Server Version : 50724
 Source Host           : localhost:3306
 Source Schema         : leetcode

 Target Server Type    : MySQL
 Target Server Version : 50724
 File Encoding         : 65001

 Date: 14/02/2019 11:42:34
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for seat
-- ----------------------------
DROP TABLE IF EXISTS `seat`;
CREATE TABLE `seat` (
  `id`      int(11)                                 DEFAULT NULL,
  `student` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

-- ----------------------------
-- Records of seat
-- ----------------------------
BEGIN;
INSERT INTO `seat` VALUES (1, 'Abbot');
INSERT INTO `seat` VALUES (2, 'Doris');
INSERT INTO `seat` VALUES (3, 'Emerson');
INSERT INTO `seat` VALUES (4, 'Green');
INSERT INTO `seat` VALUES (5, 'Jeames');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
