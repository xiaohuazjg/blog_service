CREATE TABLE blog_service.blog_tag (
	`tag_id` INT(10) auto_increment NOT NULL,
	`tag_name` varchar(100) DEFAULT '' NOT NULL  COMMENT '��ǩ����',
	`created_on` int(10) unsigned DEFAULT '0' COMMENT '����ʱ��',
    `created_by` varchar(100) DEFAULT '' COMMENT '������',
    `modified_on` int(10) unsigned DEFAULT '0' COMMENT '�޸�ʱ��',
   `modified_by` varchar(100) DEFAULT '' COMMENT '�޸���',
   `deleted_on` int(10) unsigned DEFAULT '0' COMMENT 'ɾ��ʱ��',
   `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '�Ƿ�ɾ�� 0Ϊδɾ����1Ϊ��ɾ��',
   `state` tinyint(3) unsigned DEFAULT '1' COMMENT '״̬ 0Ϊ���á�1Ϊ����',
	 PRIMARY KEY (`tag_id`)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci;

use databases blog_service ;

CREATE TABLE blog_service.blog_auth (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `app_key` varchar(20) DEFAULT '' COMMENT 'Key',
  `app_secret` varchar(50) DEFAULT '' COMMENT 'Secret',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '�½�ʱ��',
  `created_by` varchar(100) DEFAULT '' COMMENT '������',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '�޸�ʱ��',
  `modified_by` varchar(100) DEFAULT '' COMMENT '�޸���',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT 'ɾ��ʱ��',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '�Ƿ�ɾ�� 0Ϊδɾ����1Ϊ��ɾ��',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='��֤����';

CREATE TABLE blog_service.blog_article_tag (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `article_id` int(11) NOT NULL COMMENT '����ID',
  `tag_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '��ǩID',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '����ʱ��',
  `created_by` varchar(100) DEFAULT '' COMMENT '������',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '�޸�ʱ��',
  `modified_by` varchar(100) DEFAULT '' COMMENT '�޸���',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT 'ɾ��ʱ��',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '�Ƿ�ɾ�� 0Ϊδɾ����1Ϊ��ɾ��',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='���±�ǩ����';

CREATE TABLE blog_service.blog_article (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT '' COMMENT '���±���',
  `desc` varchar(255) DEFAULT '' COMMENT '���¼���',
  `cover_image_url` varchar(255) DEFAULT '' COMMENT '����ͼƬ��ַ',
  `content` longtext COMMENT '��������',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '�½�ʱ��',
  `created_by` varchar(100) DEFAULT '' COMMENT '������',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '�޸�ʱ��',
  `modified_by` varchar(100) DEFAULT '' COMMENT '�޸���',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT 'ɾ��ʱ��',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '�Ƿ�ɾ�� 0Ϊδɾ����1Ϊ��ɾ��',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '״̬ 0Ϊ���á�1Ϊ����',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='���¹���';