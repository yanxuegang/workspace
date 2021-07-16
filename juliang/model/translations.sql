--
-- Table structure for table `translations`
--

DROP TABLE IF EXISTS `translations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `translations` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `aid` varchar(64) NOT NULL DEFAULT '',
  `convert_id` varchar(64) NOT NULL DEFAULT '',
  `request_id` varchar(128) NOT NULL DEFAULT '',
  `imei` varchar(128) NOT NULL DEFAULT '',
  `idfa` varchar(128) NOT NULL DEFAULT '',
  `androidid` varchar(128) NOT NULL DEFAULT '',
  `oaid` varchar(128) NOT NULL DEFAULT '',
  `oaid_md5` varchar(128) NOT NULL DEFAULT '',
  `os` int(11) NOT NULL DEFAULT 0,
  `mac` varchar(128) NOT NULL DEFAULT '',
  `mac1` varchar(128) NOT NULL DEFAULT '',
  `ip` varchar(128) NOT NULL DEFAULT '',
  `ua` varchar(128) NOT NULL DEFAULT '',
  `geo` varchar(128) NOT NULL DEFAULT '',
  `ts` datetime DEFAULT NULL,
  `callback_url` varchar(255) NOT NULL DEFAULT '',
  `callback` varchar(255) NOT NULL DEFAULT '',
  `model` varchar(64) NOT NULL DEFAULT '',
  `status` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;