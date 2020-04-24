/*
 Navicat MySQL Data Transfer

 Source Server         : 201neiwan
 Source Server Type    : MySQL
 Source Server Version : 50725
 Source Host           : 192.168.3.201
 Source Database       : blog

 Target Server Type    : MySQL
 Target Server Version : 50725
 File Encoding         : utf-8

 Date: 04/24/2020 14:45:45 PM
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `category`
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(60) DEFAULT NULL COMMENT '名称',
  `order_num` int(11) NOT NULL DEFAULT '0' COMMENT '排序值（越小越靠前）',
  `slug` varchar(60) DEFAULT NULL,
  `desc` text,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8 COMMENT='文章分类';

-- ----------------------------
--  Records of `category`
-- ----------------------------
BEGIN;
INSERT INTO `category` VALUES ('23', '首页焦点1', '50', 'jiaodian1', null), ('24', '首页焦点2', '50', 'jiaodian2', null), ('27', 'php/mysql', '0', 'php-mysql', null), ('28', 'linux', '1', 'linux', null), ('29', '前端', '2', 'frontend', null), ('30', '转载', '5', 'trans', null), ('31', 'java', '4', 'java', null), ('32', 'python', '4', 'python', null), ('33', 'golang', '4', 'golang', null), ('34', 'test', '0', 'test', null);
COMMIT;

-- ----------------------------
--  Table structure for `post`
-- ----------------------------
DROP TABLE IF EXISTS `post`;
CREATE TABLE `post` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `post_title` varchar(255) NOT NULL DEFAULT '' COMMENT '标题',
  `post_excerpt` text,
  `post_content` longtext COMMENT '内容',
  `tags` text COMMENT '标签',
  `categorys` text COMMENT '分类',
  `is_published` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否已发布',
  `poster` varchar(255) DEFAULT NULL COMMENT '海报',
  `created_at` datetime DEFAULT NULL,
  `modified_at` datetime DEFAULT NULL,
  `published_at` datetime DEFAULT NULL COMMENT '发布时间',
  `views` int(11) DEFAULT NULL COMMENT '阅读数',
  `post_author` int(11) DEFAULT NULL COMMENT '发布者',
  `like_count` int(11) NOT NULL DEFAULT '0',
  `like_uids` text,
  `like_ips` text,
  `coins` text COMMENT '相关项目',
  `is_dujia` tinyint(4) NOT NULL DEFAULT '0',
  `is_shoufa` tinyint(4) NOT NULL DEFAULT '0',
  `words_num` int(11) NOT NULL DEFAULT '0',
  `push_status` tinyint(4) DEFAULT '0' COMMENT '推送状态',
  `is_broadcast` int(11) DEFAULT '0',
  `is_special_column` tinyint(1) NOT NULL DEFAULT '0',
  `special_column_id` int(11) DEFAULT NULL,
  `comment_count` int(11) DEFAULT '0' COMMENT '评论次数',
  `is_first_submit` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否首次提交',
  `is_show_taipei` tinyint(4) DEFAULT '0' COMMENT '显示区域：中国台北，0 不显示， 1 显示',
  `is_show_mainland` tinyint(4) DEFAULT '0' COMMENT '显示区域：中国大陆， 0 不显示， 1 显示',
  `is_show_first` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否首页显示（如果是专栏文章，设置首页值为1，否则为0）改字段主要针对专栏设置到首页',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=76 DEFAULT CHARSET=utf8 COMMENT='文章';

-- ----------------------------
--  Records of `post`
-- ----------------------------
BEGIN;
INSERT INTO `post` VALUES ('65', 'test', 'test', 'content', 'tags', 'golang', '1', 'poster', '2020-04-21 09:35:02', null, null, '0', '0', '0', null, null, null, '0', '0', '0', '0', '0', '0', null, '0', '1', '0', '0', '0'), ('66', 'test', 'test', 'content', 'tags', 'golang', '1', 'poster', '2020-04-21 09:35:24', null, null, '0', '0', '0', null, null, null, '0', '0', '0', '0', '0', '0', null, '0', '1', '0', '0', '0'), ('67', '', '', 'content', 'tags', 'golang', '1', 'poster', '2020-04-21 09:35:27', null, null, '0', '0', '0', null, null, null, '0', '0', '0', '0', '0', '0', null, '0', '1', '0', '0', '0'), ('68', '', '', '', '', 'golang', '1', 'poster', '2020-04-21 09:35:33', null, null, '0', '0', '0', null, null, null, '0', '0', '0', '0', '0', '0', null, '0', '1', '0', '0', '0'), ('69', 'test', 'test', 'content', 'tags', 'golang', '1', '', '2020-04-21 09:35:42', null, null, '0', '0', '0', null, null, null, '0', '0', '0', '0', '0', '0', null, '0', '1', '0', '0', '0'), ('70', 'test', 'test', 'content', 'tags', 'golang', '1', 'poster', '2020-04-21 09:35:46', null, null, '0', '0', '0', null, null, null, '0', '0', '0', '0', '0', '0', null, '0', '1', '0', '0', '0'), ('71', 'test', 'test', 'content', 'tags', 'golang', '1', 'poster', '2020-04-21 09:36:32', null, null, '0', '0', '0', null, null, null, '0', '0', '0', '0', '0', '0', null, '0', '1', '0', '0', '0'), ('72', 'test', 'test', 'content', 'tags', 'golang', '1', 'poster', '2020-04-21 09:36:45', null, null, '0', '0', '0', null, null, null, '0', '0', '0', '0', '0', '0', null, '0', '1', '0', '0', '0'), ('73', 'test', 'test', 'content', 'tags', 'golang', '1', 'poster', '2020-04-21 10:11:54', null, null, '0', '0', '0', null, null, null, '0', '0', '0', '0', '0', '0', null, '0', '1', '0', '0', '0'), ('74', 'test', 'test', 'content', 'tags', 'golang', '1', 'poster', '2020-04-21 10:14:43', null, null, '0', '0', '0', null, null, null, '0', '0', '0', '0', '0', '0', null, '0', '1', '0', '0', '0'), ('75', 'test', 'test', 'content', 'tags', 'golang', '1', 'poster', '2020-04-23 09:07:35', null, null, '0', '0', '0', null, null, null, '0', '0', '0', '0', '0', '0', null, '0', '1', '0', '0', '0');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
