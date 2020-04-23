use gin;

/*用户表*/
CREATE TABLE `in_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL,
  `email` varchar(64) NOT NULL,
  `password` varchar(64) NOT NULL,
  `code` varchar(64) NOT NULL DEFAULT "",
  `pwd_salt` varchar(30) NOT NULL DEFAULT "",
  `reg_ip` varchar(20) NOT NULL DEFAULT "",
  `reg_time` int(11) NOT NULL DEFAULT 0,
  `last_login_ip` varchar(20) NOT NULL DEFAULT "",
  `last_login_time` int(11) NOT NULL DEFAULT "0",
  `avatar` varchar(128) NOT NULL  DEFAULT "",
  `stat` tinyint(3) NOT NULL DEFAULT '-1' COMMENT '0-是正常;-1-未验证;-2-冻结；-3-删除',
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
