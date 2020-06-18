/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50644
 Source Host           : localhost:3306
 Source Schema         : heimdallr

 Target Server Type    : MySQL
 Target Server Version : 50644
 File Encoding         : 65001

 Date: 18/06/2020 17:02:29
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for hd_backend_user
-- ----------------------------
DROP TABLE IF EXISTS `hd_backend_user`;
CREATE TABLE `hd_backend_user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `real_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `user_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `user_pwd` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `is_super` tinyint(1) NOT NULL DEFAULT 0,
  `status` int(11) NOT NULL DEFAULT 0,
  `mobile` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `email` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `avatar` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of hd_backend_user
-- ----------------------------
INSERT INTO `hd_backend_user` VALUES (1, '超级管理员', 'admin', 'e10adc3949ba59abbe56e057f20f883e', 1, 1, '', '', '/static/upload/admin.png');
INSERT INTO `hd_backend_user` VALUES (2, '张三', 'zhangsan', 'e10adc3949ba59abbe56e057f20f883e', 0, 1, '', '', '');
INSERT INTO `hd_backend_user` VALUES (3, '李四', 'lisi', 'e10adc3949ba59abbe56e057f20f883e', 0, 0, '', '', '');
INSERT INTO `hd_backend_user` VALUES (4, '只读用户', 'readonly', '336ebbb2179beaa7340a4f1620f3af40', 0, 1, '', '', '/static/upload/readonly.png');

-- ----------------------------
-- Table structure for hd_course
-- ----------------------------
DROP TABLE IF EXISTS `hd_course`;
CREATE TABLE `hd_course`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `short_name` varchar(8) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `price` double NOT NULL DEFAULT 0,
  `real_price` double NOT NULL DEFAULT 0,
  `img` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `start_time` datetime(0) NOT NULL,
  `end_time` datetime(0) NOT NULL,
  `seq` int(11) NOT NULL DEFAULT 0,
  `creator_id` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for hd_env
-- ----------------------------
DROP TABLE IF EXISTS `hd_env`;
CREATE TABLE `hd_env`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `env_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `env_descript` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `ipaddr` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `port` int(5) NOT NULL,
  `token` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `relation_path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'ng_conf',
  `status` tinyint(1) NOT NULL DEFAULT 1,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of hd_env
-- ----------------------------
INSERT INTO `hd_env` VALUES (1, 'DEV环境', '开发人员测试环境', '192.168.43.151', 12321, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1OTE5MzczNDYsInVzZXJfaWQiOjEsInBhc3N3b3JkIjoiQnVsdGdhbmciLCJ1c2VybmFtZSI6ImhlaW1kYWxsIiwiZnVsbF9uYW1lIjoiaGVpbWRhbGwiLCJwZXJtaXNzaW9ucyI6W119.D2f8tPjzPeJYAZfhVcFox57Q0oQXtTyxedLGxUlTmSQ', 'ng_conf1', 1);
INSERT INTO `hd_env` VALUES (2, 'SIT环境', '系统单元测试环境', '192.168.43.151', 12321, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1OTE5MzczNDYsInVzZXJfaWQiOjEsInBhc3N3b3JkIjoiQnVsdGdhbmciLCJ1c2VybmFtZSI6ImhlaW1kYWxsIiwiZnVsbF9uYW1lIjoiaGVpbWRhbGwiLCJwZXJtaXNzaW9ucyI6W119.D2f8tPjzPeJYAZfhVcFox57Q0oQXtTyxedLGxUlTmSQ', 'ng_conf2', 1);

-- ----------------------------
-- Table structure for hd_resource
-- ----------------------------
DROP TABLE IF EXISTS `hd_resource`;
CREATE TABLE `hd_resource`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `rtype` int(11) NOT NULL DEFAULT 0,
  `name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `parent_id` int(11) NULL DEFAULT NULL,
  `seq` int(11) NOT NULL DEFAULT 0,
  `icon` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `url_for` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 47 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of hd_resource
-- ----------------------------
INSERT INTO `hd_resource` VALUES (7, 1, '权限管理', 8, 200, 'fa fa-balance-scale', '');
INSERT INTO `hd_resource` VALUES (8, 0, '系统菜单', NULL, 200, '', '');
INSERT INTO `hd_resource` VALUES (9, 1, '资源管理', 7, 300, '', 'ResourceController.Index');
INSERT INTO `hd_resource` VALUES (12, 1, '角色管理', 7, 200, '', 'RoleController.Index');
INSERT INTO `hd_resource` VALUES (13, 1, '用户管理', 7, 100, '', 'BackendUserController.Index');
INSERT INTO `hd_resource` VALUES (14, 1, 'Heimdallr信息管理', 8, 100, 'fa fa-gears', '');
INSERT INTO `hd_resource` VALUES (23, 1, '日志管理(空)', 14, 400, '', '');
INSERT INTO `hd_resource` VALUES (25, 2, '编辑', 9, 100, 'fa fa-pencil', 'ResourceController.Edit');
INSERT INTO `hd_resource` VALUES (26, 2, '编辑', 13, 100, 'fa fa-pencil', 'BackendUserController.Edit');
INSERT INTO `hd_resource` VALUES (27, 2, '删除', 9, 100, 'fa fa-trash', 'ResourceController.Delete');
INSERT INTO `hd_resource` VALUES (29, 2, '删除', 13, 100, 'fa fa-trash', 'BackendUserController.Delete');
INSERT INTO `hd_resource` VALUES (30, 2, '编辑', 12, 100, 'fa fa-pencil', 'RoleController.Edit');
INSERT INTO `hd_resource` VALUES (31, 2, '删除', 12, 100, 'fa fa-trash', 'RoleController.Delete');
INSERT INTO `hd_resource` VALUES (32, 2, '分配资源', 12, 100, 'fa fa-th', 'RoleController.Allocate');
INSERT INTO `hd_resource` VALUES (38, 0, '信息页面', NULL, 100, '', '');
INSERT INTO `hd_resource` VALUES (39, 1, '统计概览', 38, 100, 'fa fa-tachometer', 'StatisticalController.Index');
INSERT INTO `hd_resource` VALUES (40, 1, '配置详情', 38, 200, 'fa fa-list-alt', 'DetailController.Index');
INSERT INTO `hd_resource` VALUES (41, 1, 'Server 管理', 14, 100, 'fa fa-circle-o', '');
INSERT INTO `hd_resource` VALUES (42, 1, 'Location 管理', 14, 200, 'fa fa-circle-o', '');
INSERT INTO `hd_resource` VALUES (43, 1, 'Stream 管理', 14, 300, 'fa fa-circle-o', '');
INSERT INTO `hd_resource` VALUES (44, 1, '环境管理', 7, 400, '', 'EnvController.Index');
INSERT INTO `hd_resource` VALUES (45, 2, '编辑', 44, 100, 'fa fa-pencil', 'EnvController.Edit');
INSERT INTO `hd_resource` VALUES (46, 2, '删除', 44, 100, 'fa fa-trash', 'EnvController.Delete');

-- ----------------------------
-- Table structure for hd_role
-- ----------------------------
DROP TABLE IF EXISTS `hd_role`;
CREATE TABLE `hd_role`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `seq` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 103 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of hd_role
-- ----------------------------
INSERT INTO `hd_role` VALUES (100, '超级管理员', 100);
INSERT INTO `hd_role` VALUES (101, '角色管理员', 101);
INSERT INTO `hd_role` VALUES (102, '只读角色', 102);

-- ----------------------------
-- Table structure for hd_role_backenduser_rel
-- ----------------------------
DROP TABLE IF EXISTS `hd_role_backenduser_rel`;
CREATE TABLE `hd_role_backenduser_rel`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL,
  `backend_user_id` int(11) NOT NULL,
  `created` datetime(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of hd_role_backenduser_rel
-- ----------------------------
INSERT INTO `hd_role_backenduser_rel` VALUES (1, 100, 1, '2020-06-16 03:10:41');
INSERT INTO `hd_role_backenduser_rel` VALUES (2, 101, 2, '2020-06-16 03:11:07');
INSERT INTO `hd_role_backenduser_rel` VALUES (3, 102, 4, '2020-06-18 07:53:33');

-- ----------------------------
-- Table structure for hd_role_resource_rel
-- ----------------------------
DROP TABLE IF EXISTS `hd_role_resource_rel`;
CREATE TABLE `hd_role_resource_rel`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL,
  `resource_id` int(11) NOT NULL,
  `created` datetime(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 65 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of hd_role_resource_rel
-- ----------------------------
INSERT INTO `hd_role_resource_rel` VALUES (21, 101, 38, '2020-06-16 03:11:31');
INSERT INTO `hd_role_resource_rel` VALUES (22, 101, 39, '2020-06-16 03:11:31');
INSERT INTO `hd_role_resource_rel` VALUES (23, 101, 40, '2020-06-16 03:11:31');
INSERT INTO `hd_role_resource_rel` VALUES (24, 101, 8, '2020-06-16 03:11:31');
INSERT INTO `hd_role_resource_rel` VALUES (25, 101, 14, '2020-06-16 03:11:31');
INSERT INTO `hd_role_resource_rel` VALUES (26, 101, 41, '2020-06-16 03:11:31');
INSERT INTO `hd_role_resource_rel` VALUES (27, 101, 42, '2020-06-16 03:11:31');
INSERT INTO `hd_role_resource_rel` VALUES (28, 101, 43, '2020-06-16 03:11:31');
INSERT INTO `hd_role_resource_rel` VALUES (29, 101, 23, '2020-06-16 03:11:31');
INSERT INTO `hd_role_resource_rel` VALUES (30, 101, 7, '2020-06-16 03:11:31');
INSERT INTO `hd_role_resource_rel` VALUES (31, 101, 13, '2020-06-16 03:11:31');
INSERT INTO `hd_role_resource_rel` VALUES (32, 101, 26, '2020-06-16 03:11:31');
INSERT INTO `hd_role_resource_rel` VALUES (33, 101, 29, '2020-06-16 03:11:31');
INSERT INTO `hd_role_resource_rel` VALUES (34, 101, 12, '2020-06-16 03:11:31');
INSERT INTO `hd_role_resource_rel` VALUES (35, 101, 30, '2020-06-16 03:11:31');
INSERT INTO `hd_role_resource_rel` VALUES (36, 101, 31, '2020-06-16 03:11:31');
INSERT INTO `hd_role_resource_rel` VALUES (37, 101, 32, '2020-06-16 03:11:31');
INSERT INTO `hd_role_resource_rel` VALUES (38, 101, 9, '2020-06-16 03:11:31');
INSERT INTO `hd_role_resource_rel` VALUES (39, 101, 25, '2020-06-16 03:11:31');
INSERT INTO `hd_role_resource_rel` VALUES (40, 101, 27, '2020-06-16 03:11:31');
INSERT INTO `hd_role_resource_rel` VALUES (41, 100, 38, '2020-06-16 06:46:10');
INSERT INTO `hd_role_resource_rel` VALUES (42, 100, 39, '2020-06-16 06:46:10');
INSERT INTO `hd_role_resource_rel` VALUES (43, 100, 40, '2020-06-16 06:46:10');
INSERT INTO `hd_role_resource_rel` VALUES (44, 100, 8, '2020-06-16 06:46:10');
INSERT INTO `hd_role_resource_rel` VALUES (45, 100, 14, '2020-06-16 06:46:10');
INSERT INTO `hd_role_resource_rel` VALUES (46, 100, 41, '2020-06-16 06:46:10');
INSERT INTO `hd_role_resource_rel` VALUES (47, 100, 42, '2020-06-16 06:46:10');
INSERT INTO `hd_role_resource_rel` VALUES (48, 100, 43, '2020-06-16 06:46:10');
INSERT INTO `hd_role_resource_rel` VALUES (49, 100, 23, '2020-06-16 06:46:10');
INSERT INTO `hd_role_resource_rel` VALUES (50, 100, 7, '2020-06-16 06:46:10');
INSERT INTO `hd_role_resource_rel` VALUES (51, 100, 13, '2020-06-16 06:46:10');
INSERT INTO `hd_role_resource_rel` VALUES (52, 100, 26, '2020-06-16 06:46:10');
INSERT INTO `hd_role_resource_rel` VALUES (53, 100, 29, '2020-06-16 06:46:10');
INSERT INTO `hd_role_resource_rel` VALUES (54, 100, 12, '2020-06-16 06:46:10');
INSERT INTO `hd_role_resource_rel` VALUES (55, 100, 30, '2020-06-16 06:46:10');
INSERT INTO `hd_role_resource_rel` VALUES (56, 100, 31, '2020-06-16 06:46:10');
INSERT INTO `hd_role_resource_rel` VALUES (57, 100, 32, '2020-06-16 06:46:10');
INSERT INTO `hd_role_resource_rel` VALUES (58, 100, 9, '2020-06-16 06:46:10');
INSERT INTO `hd_role_resource_rel` VALUES (59, 100, 25, '2020-06-16 06:46:10');
INSERT INTO `hd_role_resource_rel` VALUES (60, 100, 27, '2020-06-16 06:46:10');
INSERT INTO `hd_role_resource_rel` VALUES (61, 100, 44, '2020-06-16 06:46:10');
INSERT INTO `hd_role_resource_rel` VALUES (62, 102, 38, '2020-06-18 07:53:23');
INSERT INTO `hd_role_resource_rel` VALUES (63, 102, 39, '2020-06-18 07:53:23');
INSERT INTO `hd_role_resource_rel` VALUES (64, 102, 40, '2020-06-18 07:53:23');

SET FOREIGN_KEY_CHECKS = 1;
