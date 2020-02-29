/*
 Navicat Premium Data Transfer

 Source Server         : test
 Source Server Type    : PostgreSQL
 Source Server Version : 110001
 Source Host           : 118.24.1.111:5432
 Source Catalog        : store
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 110001
 File Encoding         : 65001

 Date: 26/02/2019 16:00:40
*/


-- ----------------------------
-- Sequence structure for users_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."users_id_seq";
CREATE SEQUENCE "public"."users_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 32767
START 1
CACHE 1;

-- ----------------------------
-- Table structure for administrators
-- ----------------------------
DROP TABLE IF EXISTS "public"."administrators";
CREATE TABLE "public"."administrators" (
  "id" int2 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
  "name" varchar(25) COLLATE "pg_catalog"."default" NOT NULL,
  "username" varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
  "password" varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
  "created_at" timestamptz(0),
  "updated_at" timestamptz(0),
  "deleted_at" timestamptz(0)
)
;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "public"."users";
CREATE TABLE "public"."users" (
  "id" int2 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
  "name" varchar(25) COLLATE "pg_catalog"."default" NOT NULL,
  "username" varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
  "password" varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
  "created_at" timestamptz(0),
  "updated_at" timestamptz(0),
  "deleted_at" timestamptz(0)
)
;

-- ----------------------------
-- Records of administrators
-- ----------------------------
INSERT INTO "public"."administrators" VALUES (1, 'admin', 'admin', 'e10adce10adc3949ba59abbe56e057f20f883e', NULL, NULL, NULL);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."users_id_seq"
OWNED BY "public"."administrators"."id";
SELECT setval('"public"."users_id_seq"', 2, true);
