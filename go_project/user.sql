SET FOREIGN_KEY_CHECKS=0;
 
-- ----------------------------
-- Table structure for `user`
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(60) DEFAULT NULL,
  `age` int(4) DEFAULT NULL,
  `info` varchar(60) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=25 DEFAULT CHARSET=utf8;
 
-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', 'slene', '81', 'Info');
INSERT INTO `user` VALUES ('2', 'slene81', '81', 'Info81');
INSERT INTO `user` VALUES ('3', 'slene87', '87', 'Info87');
INSERT INTO `user` VALUES ('4', 'slene47', '47', 'Info47');
INSERT INTO `user` VALUES ('5', 'slene59', '59', 'Info59');
INSERT INTO `user` VALUES ('6', 'slene81', '81', 'Info81');
INSERT INTO `user` VALUES ('7', 'slene18', '18', 'Info18');
INSERT INTO `user` VALUES ('8', 'slene25', '25', 'Info25');
INSERT INTO `user` VALUES ('9', 'slene40', '40', 'Info40');
INSERT INTO `user` VALUES ('10', 'slene56', '56', 'Info56');
INSERT INTO `user` VALUES ('11', 'slene0', '0', 'Info0');
INSERT INTO `user` VALUES ('12', 'slene94', '94', 'Info94');
INSERT INTO `user` VALUES ('13', 'slene11', '11', 'Info11');
INSERT INTO `user` VALUES ('14', 'slene62', '62', 'Info62');
INSERT INTO `user` VALUES ('15', 'slene89', '89', 'Info89');
INSERT INTO `user` VALUES ('16', 'slene28', '28', 'Info28');
INSERT INTO `user` VALUES ('17', 'slene74', '74', 'Info74');
INSERT INTO `user` VALUES ('18', 'slene11', '11', 'Info11');
INSERT INTO `user` VALUES ('19', 'slene45', '45', 'Info45');
INSERT INTO `user` VALUES ('20', 'slene37', '37', 'Info37');
INSERT INTO `user` VALUES ('21', 'slene6', '6', 'Info6');
INSERT INTO `user` VALUES ('22', 'slene95', '95', 'Info95');
INSERT INTO `user` VALUES ('23', 'slene66', '66', 'Info66');
INSERT INTO `user` VALUES ('24', 'slene28', '28', 'Info28');

