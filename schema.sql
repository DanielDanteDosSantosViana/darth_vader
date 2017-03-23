/*
DROP TABLE IF EXISTS `sequence`;
CREATE TABLE `sequence` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `stub` varchar(1) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `stub` (`stub`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `url_encurtada`;
CREATE TABLE `url_encurtada` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `url_longa` varchar(10240) NOT NULL,
  `alias` varchar(11) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `tempo_operacao` varchar(11) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `acessos` int(10) NOT NULL DEFAULT 0,
 PRIMARY KEY (`id`),
  UNIQUE KEY `alias` (`alias`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

*/

DROP TABLE IF EXISTS `area`;
CREATE TABLE `area` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `descricao` varchar(10240) NOT NULL,
 PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
